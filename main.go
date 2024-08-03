package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"

	"code.vegaprotocol.io/pyth-pull-prices/pythmcapindex"
	"code.vegaprotocol.io/pyth-pull-prices/pythrebase"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	endpoint      string
	port          uint
	configPath    string
	pythRebase    = common.HexToAddress("0x2ba6654fdb605637994047eca217a02e224c4828")
	pythMcapIndex = common.HexToAddress("0xF60DA8e2DACfc5677F11c370F6C9d595f13C13Dc")

	pairs = map[string]string{
		"PEPE/USD":   "d69731a2e74ac1ce884fc3890f7ee324b6deb66147055249568869ed700882e4",
		"BTC/USD":    "e62df6c8b4a85fe1a67db44dc12de5db330f7ac66b72dc658afedf0f4a415b43",
		"SOL/USD":    "ef0d8b6fda2ceba41da15d4095d1da392a0d2f8ed0c6c7bc0f4cfac8c280b56d",
		"ETH/USD":    "ff61491a931112ddf1bd8147cd1b641375f79f5825126d665480874634fd0ace",
		"USDT/USD":   "2b89b9dc8fdf9f34709a5b106b472f0f39bb6ca9ce04b0fd7f2e971688e2e53b",
		"INJ/USD":    "7a5bc1d2b56ad029048cd63964b3ad2776eadf812edc1a43a31406cb54bff592",
		"SNX/USD":    "39d020f60982ed892abbcd4a06a276a9f9b7bfbce003204c110b6e488f502da3",
		"LDO/USD":    "c63e2a7f37a04e5e614c07238bedb25dcc38927fba8fe890597a593c0b2fa4ad",
		"LINK/USD":   "8ac0c70fff57e9aefdf5edf44b51d62c2d433653cbb2cf5cc06bb115af04d221",
		"ENA/USD":    "b7910ba7322db020416fcac28b48c01212fd9cc8fbcbaf7d30477ed8605f6bd4",
		"SUI/USD":    "23d7315113f5b1d3ba7a83604c44b94d79f4fd69af77f804fc7f920a6dc65744",
		"TIA/USD":    "09f7c1d7dfbb7df2b8fe3d3d87ee94a2259d212da4f30c1f0540d066dfa44723",
		"ARB/USD":    "3fa4252848f9f0a1480be62745a4629d9eb1322aebab8a791e344b3b9c1adcf5",
		"DOGE/USD":   "dcef50dd0a4cd2dcc17e45df1676dcb336a11a61c69df7a0299b0150c672d25c",
		"SHIB/USD":   "f0d57deca57b3da2fe63a493f4c25925fdfd8edf834b20f93e1f84dbd1504d4a",
		"WIF/USD":    "4ca4beeca86f0d164160323817a4e42b10010a724c2217c6ee41b54cd4cc61fc",
		"FLOKI/USD":  "6b1381ce7e874dc5410b197ac8348162c0dd6c0d4c9cd6322672d6c2b1d58293",
		"NEON/USD":   "d82183dd487bef3208a227bb25d748930db58862c5121198e723ed0976eb92b7",
		"GMCI30/USD": "b3bd2a89c026fe7873da91b39bf9347fdea24e5b588330b4899788cf48878133",
	}

	indexes = map[string][]pairMcap{
		"MEME5/USD": {
			{ /**/ "FLOKI/USD", big.NewInt(9690710434697)},
			{ /* */ "PEPE/USD", big.NewInt(420690000000000)},
			{ /*  */ "WIF/USD", big.NewInt(998905927)},
			{ /* */ "DOGE/USD", big.NewInt(144596276384)},
			{ /* */ "SHIB/USD", big.NewInt(589519957587269)},
		}, //                              9223372036854775808
	}
)

type config struct {
	Pairs   map[string]string     `json:"pairs"`
	Indexes map[string][]pairMcap `json:"indexes"`
}

type pairMcap struct {
	Pair string   `json:"pair"`
	MCap *big.Int `json:"m_cap"`
}

type priceProvider struct {
	ethClient     *ethclient.Client
	pythRebase    *pythrebase.PythrebaseCaller
	pythMcapIndex *pythmcapindex.PythmcapindexCaller

	cache map[string]*big.Int
}

func (p *priceProvider) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	sym := values.Get("symbol")
	if len(sym) <= 0 {
		http.Error(w, "missing symbol query parameter", http.StatusBadRequest)
		return
	}

	if pairId, ok := pairs[sym]; ok {
		id, err := hex.DecodeString(pairId)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		p.doPairPrice(w, r, id)
		return

	} else if indexCfg, ok := indexes[sym]; ok {
		p.doPythMcapPrice(w, r, indexCfg, sym)
		return
	}

	http.Error(w, "unsupported pair symbol", http.StatusBadRequest)
}

func (p *priceProvider) doPythMcapPrice(w http.ResponseWriter, r *http.Request, pairCaps []pairMcap, sym string) {
	indexAssets := make([]pythmcapindex.IndexAsset, 0, len(pairCaps))
	for _, v := range pairCaps {
		pairId, ok := pairs[v.Pair]
		if !ok {
			http.Error(w, fmt.Sprintf("internal server error: %v", "bad pair in index"), http.StatusInternalServerError)
			return
		}

		id, err := hex.DecodeString(pairId)
		if err != nil {
			http.Error(w, "internal server error: couldn't get bytes for pairId", http.StatusInternalServerError)
			return
		}

		var id32 [32]byte
		copy(id32[:], id)

		indexAssets = append(indexAssets, pythmcapindex.IndexAsset{
			PythId:      id32,
			TotalSupply: big.NewInt(0).Set(v.MCap),
		})
	}

	cachedPrice := p.cache[sym]

	price, err := p.pythMcapIndex.GetIndexPrice(&bind.CallOpts{
		Context: r.Context(),
	}, indexAssets)
	if err != nil {
		if cachedPrice == nil {
			http.Error(w, fmt.Sprintf("internal server error: %v", err), http.StatusInternalServerError)
			return
		}

		price = cachedPrice
	}

	p.cache[sym] = price
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"price": "%v"}`, price.String())
}

func (p *priceProvider) doPairPrice(w http.ResponseWriter, r *http.Request, id []byte) {
	var id32 [32]byte
	copy(id32[:], id)

	cachedPrice := p.cache[string(id)]

	price, err := p.pythRebase.GetPrice0(&bind.CallOpts{
		Context: r.Context(),
	}, id32)
	if err != nil {
		if cachedPrice == nil {
			http.Error(w, fmt.Sprintf("internal server error: %v", err), http.StatusInternalServerError)
			return
		}

		price = cachedPrice
	}

	p.cache[string(id)] = price

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"price": "%v"}`, price.String())
}

func main() {
	flag.Parse()

	if len(endpoint) <= 0 {
		log.Fatalf("missing required parameter endpoint")
	}

	if len(configPath) > 0 {
		loadCustomConfig()
	}

	ethClient, err := ethclient.DialContext(context.Background(), endpoint)
	if err != nil {
		log.Fatalf("could not start eth client: %v", err)
	}

	pythRebase, err := pythrebase.NewPythrebaseCaller(pythRebase, ethClient)
	if err != nil {
		log.Fatalf("could not instantiate pyth rebase caller: ", err)
	}

	pythMcapIndex, err := pythmcapindex.NewPythmcapindexCaller(pythMcapIndex, ethClient)
	if err != nil {
		log.Fatalf("could not instantiate pyth mcap index caller: ", err)
	}

	cache := make(map[string]*big.Int)

	pp := &priceProvider{ethClient, pythRebase, pythMcapIndex, cache}

	http.Handle("/avgPrice", pp)

	httpBind := fmt.Sprintf(":%v", port)
	log.Printf("starting HTTP API on %v", httpBind)
	log.Fatal(http.ListenAndServe(httpBind, nil))
}

func loadCustomConfig() {
	content, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("couldn't read configuration file: %v", err)
	}

	c := config{}
	err = json.Unmarshal(content, &c)
	if err != nil {
		log.Fatalf("couldn't unmarshal configuration file: %v", err)
	}

	// all good here, just set the global variable
	pairs = c.Pairs
	indexes = c.Indexes
}

func init() {
	flag.StringVar(&endpoint, "endpoint", "", "the gnosis RPC endpoint")
	flag.StringVar(&configPath, "config", "", "The path to a custom config, if none set use default values")
	flag.UintVar(&port, "port", 8080, "the HTTP API port")
}
