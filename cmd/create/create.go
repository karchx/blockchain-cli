package create

import "github.com/spf13/cobra"

// CreateCmd represents the create command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new object related to the blockchain",
	Long:  `See "blockchain-cli --help" for sub-commands`,
}

func init() {}
