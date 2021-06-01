package chain

import (
	"fmt"

	"github.com/iotaledger/wasp/packages/registry_pkg/chain_record"

	"github.com/iotaledger/wasp/tools/wasp-cli/config"
	"github.com/iotaledger/wasp/tools/wasp-cli/log"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List deployed chains",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client := config.WaspClient()
		chains, err := client.GetChainRecordList()
		log.Check(err)
		log.Printf("Total %d chain(s) in wasp node %s\n", len(chains), client.BaseURL())
		showChainList(chains)
	},
}

func showChainList(chains []*chain_record.ChainRecord) {
	header := []string{"chainid", "active"}
	rows := make([][]string, len(chains))
	for i, chain := range chains {
		rows[i] = []string{
			chain.ChainAddr.Base58(),
			fmt.Sprintf("%v", chain.Active),
		}
	}
	log.PrintTable(header, rows)
}
