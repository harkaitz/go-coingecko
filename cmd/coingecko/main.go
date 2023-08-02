package main

import (
	"os"
	"fmt"
	"time"
	"strings"
	"strconv"
	"errors"
	"github.com/harkaitz/go-coingecko"
	"github.com/harkaitz/go-u27"
	"github.com/pborman/getopt/v2"
)

const help string = `Usage: coingecko [-C usd] [-D DATE] [-L DAYS] ...

Fetch cryptocurrency prices from "coingecko.com".

  -l                    : List available cryptocurrencies.
  -p TICKETS...         : Print the price of tickers.
  -f FROM    VALUE...   : Convert to fiat
  -c TO      VALUE...   : Convert to crypto.
  -k FROM-TO VALUE...   : Convert crypto to crypto.
  -g COIN [-D,-L]       : Print price graph.
  -i COIN [-D]          : Print historic price.

Copyright (c) 2023 Harkaitz Agirre, harkaitz.aguirre@gmail.com`


func main() {
	var err     error
	var arg     string
	var value   float64
	var coinID  coingecko.CoinID
	var coins []coingecko.Coin
	var data    coingecko.CoinData
	var graph   coingecko.CoinGraph
	var price   float64
	var result  string
	var date    time.Time = u.Now()
	var start   time.Time
	
	
	defer func() {
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err.Error())
			os.Exit(1)
		}
		fmt.Print(result)
	}()
	
	_      = getopt.BoolLong   ("help"  , 'h')
	lFlag := getopt.BoolLong   ("list"  , 'l')
	pFlag := getopt.BoolLong   ("prices", 'p')
	fFlag := getopt.StringLong ("to-fiat"  , 'f', "")
	cFlag := getopt.StringLong ("to-crypto", 'c', "") 
	kFlag := getopt.StringLong ("convert"  , 'k', "")
	gFlag := getopt.StringLong ("graph"    , 'g', "")
	iFlag := getopt.StringLong ("history"  , 'i', "")
	
	CFlag := getopt.StringLong ("currency", 'C', "usd")
	DFlag := getopt.StringLong ("date"    , 'D', "")
	LFlag := getopt.Int64Long  ("days"    , 'L', 365)
	
	
	getopt.SetUsage(func() { fmt.Println(help) })
	getopt.Parse()
	
	if *DFlag != "" {
		date, err = time.Parse("2006-01-02", *DFlag)
		if err != nil { return }
	}
	start = date.AddDate(0, 0, -int(*LFlag))
	
	switch {
	case *lFlag:
		//
		coins, err = coingecko.GetCoinList()
		if err != nil { return }
		//
		fmt.Printf("%-20v %-10v %v\n", "ID", "Symbol", "Name")
		for _, c := range coins {
			fmt.Printf("%-20v %-10v %v\n", c.ID, c.Symbol, c.Name)
		}
	case *pFlag:
		//
		for _, arg = range getopt.Args() {
			//
			price, err = coingecko.GetCoinPrice(arg, *CFlag)
			if err != nil { return }
			//
			result += fmt.Sprintf("%-10v %v\n", arg, price)
		}
	case *fFlag != "":
		//
		price, err = coingecko.GetCoinPrice(*fFlag, *CFlag)
		if err != nil { return }
		//
		for _, arg = range getopt.Args() {
			//
			value, err = strconv.ParseFloat(arg, 64)
			if err != nil { return }
			//
			result += SprintAmount(*CFlag, value*price)
		}
	case *cFlag != "":
		//
		price, err = coingecko.GetCoinPrice(*cFlag, *CFlag)
		if err != nil { return }
		//
		for _, arg = range getopt.Args() {
			//
			value, err = strconv.ParseFloat(arg, 64)
			if err != nil { return }
			//
			result += SprintAmount(*cFlag, value/price)
		}
	case *kFlag != "":
		var names   []string = strings.Split(*kFlag, "-")
		var prices [2]float64
		//
		if len(names) != 2 {
			err = errors.New("Invalid argument for -k")
			return
		}
		//
		prices[0], err = coingecko.GetCoinPrice(names[0], "btc")
		if err != nil { return }
		prices[1], err = coingecko.GetCoinPrice(names[1], "btc")
		if err != nil { return }
		//
		for _, arg = range getopt.Args() {
			//
			value, err = strconv.ParseFloat(arg, 64)
			if err != nil { return }
			//
			result += SprintAmount(names[1], value*prices[0]/prices[1])
		}
	case *gFlag != "":
		//
		coinID, err = coingecko.GetCoinID(*gFlag)
		if err != nil { return }
		//
		graph, err = coingecko.GetCoinGraph(coinID, start, date, *CFlag)
		if err != nil { return }
		//
		graph.PrintPrices()
	case *iFlag != "":
		//
		coinID, err = coingecko.GetCoinID(*iFlag)
		if err != nil { return }
		//
		data, err = coingecko.GetCoinHistory(coinID, date)
		if err != nil { return }
		//
		price, err = data.Price(*CFlag)
		if err != nil { return }
		//
		result = SprintAmount(*iFlag, price)
	default:
		getopt.Usage()
	}
}

func SprintAmount(name string, value float64) string {
	return fmt.Sprintf("%-10v %.10f\n", name, value)
}
