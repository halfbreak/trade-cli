package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	mapset "github.com/deckarep/golang-set"
	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets info on a currency pair",
	Long:  `Gets all the current information from the given currency pair, the following one are allowed: btcusd, btceur, btcgbp, btcpax, gbpusd, gbpeur, eurusd, xrpusd, xrpeur, xrpbtc, xrpgbp, xrppax, ltcusd, ltceur, ltcbtc, ltcgbp, ethusd, etheur, ethbtc, ethgbp, ethpax, bchusd, bcheur, bchbtc, bchgbp, paxusd, paxeur, paxgbp`,
	Run: func(cmd *cobra.Command, args []string) {

		currencyPairSlice := []interface{}{"btcusd", "btceur", "btcgbp", "btcpax", "gbpusd", "gbpeur", "eurusd", "xrpusd", "xrpeur", "xrpbtc", "xrpgbp", "xrppax", "ltcusd", "ltceur", "ltcbtc", "ltcgbp", "ethusd", "etheur", "ethbtc", "ethgbp", "ethpax", "bchusd", "bcheur", "bchbtc", "bchgbp", "paxusd", "paxeur", "paxgbp"}
		currencyPairSet := mapset.NewSetFromSlice(currencyPairSlice)

		if !currencyPairSet.Contains(args[0]) {
			fmt.Println("Currency pair " + args[0] + " not recognized")
			os.Exit(1)
		}

		fmt.Println("Getting " + strings.Join(args, " "))
		fmt.Println("https://www.bitstamp.net/api/v2/ticker/" + args[0])
		response, err := http.Get("https://www.bitstamp.net/api/v2/ticker/" + args[0])

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(responseData))
	},
	Args: cobra.MinimumNArgs(1),
}

func init() {
	RootCmd.AddCommand(helloCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
