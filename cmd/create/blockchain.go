package create

import "github.com/spf13/cobra"

var createAddress string = ""

// createBlockchainCmd represents the createBlockchain command
var createBlockchainCmd = &cobra.Command{
  Use: "blockchain",
  Short: "create a new blockchain",
  Long: `create a new blockchain and award a given address for the creation of the genesis block`,
  Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
  CreateCmd.AddCommand(createBlockchainCmd)

  createBlockchainCmd.PersistentFlags().StringVarP(&createAddress, "address", "a", "", "Address to send initial reward to")
}

func createBlockchain(address string) {}