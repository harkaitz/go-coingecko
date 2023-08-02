package coingecko

import (
	"github.com/harkaitz/go-u27"
	"errors"
	"time"
	"log"
	"fmt"
	"strconv"
	"encoding/json"
)

type Coin struct {
	ID     CoinID `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

type CoinID string

type CoinData struct {
	ID                 CoinID `json:"id"`
	Symbol             string `json:"symbol"`
	Name               string `json:"name"`
	BlockTimeInMinutes int    `json:"block_time_in_minutes"`
	Links struct {
		Homepage []string `json:"homepage"`
	} `json:"links"`
	Image struct {
		Thumb string `json:"thumb"`
		Small string `json:"small"`
		Large string `json:"large"`
	} `json:"image"`
	MarketData struct {
		CurrentPrice map[string]float64 `json:"current_price"`
	} `json:"market_data"`
}

type CoinGraph struct {
	Prices       [][2]float64
	MarketCaps   [][2]float64
	TotalVolumes [][2]float64
}

var Coingecko u.RPC = u.RPC {
	URL: "https://api.coingecko.com",
}

func GetCoinList() (coins []Coin, err error) {
	if !u.GetFromCache("COINLIST", 3600, &coins) { 
		err = Coingecko.SimQuery("/api/v3/coins/list", "GET", &coins)
		if err != nil { return }
		u.SaveToCache("COINLIST", coins)
	}
	return
}

func GetCoinData(id CoinID) (data CoinData, err error) {
	cache := "COINPRICE." + string(id)
	url := "/api/v3/coins/" + string(id) + "?localization=false"
	if !u.GetFromCache(cache, 20, &data) {
		err = Coingecko.SimQuery(url, "GET", &data)
		u.SaveToCache(cache, data)
	}
	return
}

func GetCoinID(ticker string) (id CoinID, err error) {
	var coins []Coin
	var coin    Coin
	coins, err = GetCoinList()
	if err != nil { return }
	for _, coin = range coins {
		if string(coin.ID) == ticker { return coin.ID, nil }
	}
	for _, coin = range coins {
		if coin.Symbol == ticker { return coin.ID, nil }
	}
	return "", errors.New("Coin " + ticker + " not found")
}

func GetCoinPrice(ticker, currency string) (price float64, err error) {
	var id   CoinID
	var data CoinData
	id, err = GetCoinID(ticker)
	if err != nil { return }
	data, err = GetCoinData(id)
	if err != nil { return }
	price, err = data.Price(currency)
	if err != nil { return }
	return
}

func GetCoinGraph(id CoinID, start, end time.Time, currency string) (graph CoinGraph, err error) {
	var url string
	
	url += "/api/v3/coins/" + string(id) + "/market_chart/range"
	url += "?" + "vs_currency=" + currency
	url += "&" + "from=" + strconv.FormatInt(start.Unix(), 10)
	url += "&" + "to="   + strconv.FormatInt(end.Unix()  , 10)
	
	err = Coingecko.SimQuery(url, "GET", &graph)
	return
}

func GetCoinHistory(id CoinID, day time.Time) (data CoinData, err error) {
	var url  string
	
	url += "/api/v3/coins/" + string(id) + "/history"
	url += "?" + "date=" + day.Format("02-01-2006")
	url += "&" + "localization=false"
	
	err = Coingecko.SimQuery(url, "GET", &data)
	return
}

func (g *CoinGraph) PrintPrices() {
	var row [2]float64
	
	for _, row = range g.Prices {
		fmt.Printf("%-15.0f %-15.10f\n", row[0]/1000, row[1])
	}
}

func (d *CoinData) Print() {
	b, err := json.Marshal(d)
	if err != nil {
		log.Panic(err)
	}
	fmt.Print(string(b))
}

func (d *CoinData) Price(currency string) (price float64, err error) {
	var found bool
	price, found = d.MarketData.CurrentPrice[currency]
	if !found {
		err = errors.New("Currency not supported")
		return
	}
	return
}
