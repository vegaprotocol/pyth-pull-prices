# Pyth-pull-prices

## Usage

```
git clone https://github.com/vegaprotocol/pyth-pull-prices.git
cd pyth-pull-prices

go build -o pyth-pull-prices ./

./pyth-pull-prices --endpoint https://gnosis-mainnet.public.blastapi.io --config ./config.json --port 8080

curl localhost:8080/avgPrice?symbol=LDO/USD
```
