package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/halfbreak/trade-cli/model"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets info on a currency pair",
	Long:  `Gets all the current information from the given currency pair, the following one are allowed: btcusd, btceur, btcgbp, btcpax, gbpusd, gbpeur, eurusd, xrpusd, xrpeur, xrpbtc, xrpgbp, xrppax, ltcusd, ltceur, ltcbtc, ltcgbp, ethusd, etheur, ethbtc, ethgbp, ethpax, bchusd, bcheur, bchbtc, bchgbp, paxusd, paxeur, paxgbp`,
	Run: func(cmd *cobra.Command, args []string) {

		exchange := &model.BitStamp{}
		getCurrency(args[0], exchange)
	},
	Args: cobra.MinimumNArgs(1),
}

func init() {
	RootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getCurrency(currencyPair string, exchange model.Exchange) {

	if exchange.IsInvalidCurrencyPair(currencyPair) {
		fmt.Println("Currency pair " + currencyPair + " not recognized")
		os.Exit(1)
	}

	response, err := http.Get(exchange.GetCurrencyPairURL(currencyPair))

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
}
