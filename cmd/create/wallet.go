package create

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var walletAmount int
var walletAlias string

var createWalletCmd = &cobra.Command{
	Use:   "wallet",
	Short: "Create a new wallet",
	Long:  `Create a new wallet and add it to your collection of wallets`,
	Run: func(cmd *cobra.Command, args []string) {
		if walletAmount < 1 {
			prompt := promptui.Prompt{
				Label: "Enter a valid number of wallets to create (>= 1)",
				Validate: func(s string) error {
					v, err := strconv.ParseInt(s, 10, 64)
					if err != nil || v < 1 {
						return errors.New("invalid number")
					}

					return nil
				},
			}

			result, _ := prompt.Run()
			number, _ := strconv.ParseInt(result, 10, 64)
			walletAmount = int(number)
		}

		for i := 0; i < walletAmount; i++ {
			createWallet()
		}
	},
}

func init() {
	CreateCmd.AddCommand(createWalletCmd)

	createWalletCmd.PersistentFlags().IntVarP(&walletAmount, "amount", "n", 1, "The amount of new wallets to make")
	createWalletCmd.PersistentFlags().StringVarP(&walletAlias, "alias", "a", "", "A contextual name for a wallet")
}

func createWallet() {
	fmt.Printf("Wallet created with address %s and alias '%s'\n", "test", walletAlias)
}
