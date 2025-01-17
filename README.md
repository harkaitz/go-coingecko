# GO-COINGECKO

A curated command line interface to coingecko.

## Go programs

    Usage: coingecko [-C usd] [-D DATE] [-L DAYS] ...
    
    Fetch cryptocurrency prices from "coingecko.com".
    
      -l                    : List available cryptocurrencies.
      -p TICKETS...         : Print the price of tickers.
      -f FROM    VALUE...   : Convert to fiat
      -c TO      VALUE...   : Convert to crypto.
      -k FROM-TO VALUE...   : Convert crypto to crypto.
      -g COIN [-D,-L]       : Print price graph.
      -i COIN [-D]          : Print historic price.
    
    Copyright (c) 2023 Harkaitz Agirre, harkaitz.aguirre@gmail.com

## Go documentation

    package coingecko // import "github.com/harkaitz/go-coingecko"
    
    var VerboseRPC bool = os.Getenv("VERBOSE_RPC") != ""
    func GetCoinPrice(ticker, currency string) (price float64, err error)
    func GetFromCache(id string, exp int64, out any) (found bool)
    func IsError(oB []byte) (err error)
    func SaveToCache(id string, in any)
    type Coin struct{ ... }
        func GetCoinList() (coins []Coin, err error)
    type CoinData struct{ ... }
        func GetCoinData(id CoinID) (data CoinData, err error)
        func GetCoinHistory(id CoinID, day time.Time) (data CoinData, err error)
    type CoinGraph struct{ ... }
        func GetCoinGraph(id CoinID, start, end time.Time, currency string) (graph CoinGraph, err error)
    type CoinID string
        func GetCoinID(ticker string) (id CoinID, err error)
    type RPC struct{ ... }
        var Coingecko RPC = RPC{ ... }
    type RPCRequest struct{ ... }
    type RPCResponse struct{ ... }

## Collaborating

For making bug reports, feature requests, support or consulting visit
one of the following links:

1. [gemini://harkadev.com/oss/](gemini://harkadev.com/oss/)
2. [https://harkadev.com/oss/](https://harkadev.com/oss/)
