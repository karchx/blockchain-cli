package prompts

import (
	"blockchain-cli/blockchain"
	"blockchain-cli/wallet"
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
)

func PromptAddress(text string) (string, error) {
	wallets, _ := wallet.CreateWallet()
	addrs := wallets.GetAllAddresses()
	if len(addrs) == 0 {
		return "", errors.New("you must create a wallet first")
	}

	bcExists := blockchain.BlockChainExists()

	items := []string{}

	for _, addr := range addrs {
		w := wallets.Wallets[addr]

		if !bcExists {
			if len(w.Alias) > 0 {
				items = append(items, fmt.Sprintf("%s %s", w.Alias, addr))
			} else {
				items = append(items, addr)
			}

			continue
		}

		bal := wallet.GetBalance(addr)
		if len(w.Alias) > 0 {
			items = append(items, fmt.Sprintf("%s %s => %d", w.Alias, addr, bal))
		} else {
			items = append(items, fmt.Sprintf("%s => %d", addr, bal))
		}
	}

	prompt := promptui.Select{
		Label: text,
		Items: items,
	}

	i, _, err := prompt.Run()
	if err != nil {
		fmt.Println("Failed to choose address")
		return "", err
	}

	return addrs[i], nil
}
