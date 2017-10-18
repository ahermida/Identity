/*
		Handler for CLI command listaddress
*/
package cmd

import (
		"fmt"
		"log"

		"github.com/ahermida/Identity/wallet"
)

//Lists base58 encoded addresses
func (cli *CLI) listAddresses(nodeID string) {
	// Create a new wallets object by loading data from
	// 'wallet_NODEID.dat'
	wallets, err := wallet.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}

  	// Get and print the addresses from the wallets object
	addresses := wallets.GetAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}
}
