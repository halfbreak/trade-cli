package model

import (
	mapset "github.com/deckarep/golang-set"
)

// Exchange to query for trades
type Exchange struct {
	BaseURL       string
	CurrencyPairs []interface{}
}

// GetCurrencyPairURL - Get URL to query for the information of a currency pair
func (e *Exchange) GetCurrencyPairURL(currencyPair string) string {
	return e.BaseURL + currencyPair
}

// IsInvalidCurrencyPair - Checks if given currency pair is valid
func (e *Exchange) IsInvalidCurrencyPair(currencyPair string) bool {
	currencyPairSet := mapset.NewSetFromSlice(e.CurrencyPairs)
	return !currencyPairSet.Contains(currencyPair)
}
