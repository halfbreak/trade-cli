package cmd

import (
	"fmt"
	"os"

	"github.com/halfbreak/trade-cli/model"
	"github.com/halfbreak/trade-cli/services"
	"github.com/spf13/cobra"
)

var availableCommands []string = []string{"currencyPair"}

// listCmd represents the listCmd command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists options",
	Long:  `This subcommand lists different options`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Printf("No command to list, choose one of: %v\n", availableCommands)
			os.Exit(1)
		}
		var itemTolist string = args[0]
		listCommands(itemTolist)
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// byeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// byeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	listCmd.Flags().StringVarP(&Output, "output", "o", services.OutputType[0], "Chooses the output for the Get command")
}

func listCommands(command string) {
	switch command {
	case "currencyPair":
		var outputService services.OutputService = services.GetOutput(Output)
		outputService.Write(fmt.Sprintf("Available currencies are: %v", model.CurrencyPairs))
	default:
		fmt.Printf("Unknown command, choose one of: %v\n", availableCommands)
		os.Exit(1)
	}
}
