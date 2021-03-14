package model

import (
	mapset "github.com/deckarep/golang-set"
)

var CurrencyPairs []interface{} = []interface{}{"btcusd", "btceur", "btcgbp", "btcpax", "gbpusd", "gbpeur", "eurusd", "xrpusd", "xrpeur", "xrpbtc", "xrpgbp", "xrppax", "ltcusd", "ltceur", "ltcbtc", "ltcgbp", "ethusd", "etheur", "ethbtc", "ethgbp", "ethpax", "bchusd", "bcheur", "bchbtc", "bchgbp", "paxusd", "paxeur", "paxgbp"}

// Exchange to query for trades
type Exchange interface {
	GetCurrencyPairURL(currencyPair string) string
	IsInvalidCurrencyPair(currencyPair string) bool
}

// BitStamp Exchange to query for trades
type BitStamp struct {
}

// GetCurrencyPairURL - Get URL to query for the information of a currency pair
func (e *BitStamp) GetCurrencyPairURL(currencyPair string) string {
	return "https://www.bitstamp.net/api/v2/ticker/" + currencyPair
}

// IsInvalidCurrencyPair - Checks if given currency pair is valid
func (e *BitStamp) IsInvalidCurrencyPair(currencyPair string) bool {
	currencyPairSet := mapset.NewSetFromSlice(CurrencyPairs)
	return !currencyPairSet.Contains(currencyPair)
}
