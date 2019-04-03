package main

import (
	"bufio"
	"fmt"
	"log"
	"crypto/rand"
	"os"
	"sync"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/qshuai/go-electrum/electrum"
)

func main() {

    padded := make([]byte, 32)
	electrum.DebugMode = false
	var wg sync.WaitGroup

	file, err := os.Open("electrum.txt")
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var nodes []string
	for scanner.Scan() {
		nodes = append(nodes, scanner.Text())
	}
	wg.Add(len(nodes))
	for _, line := range nodes {
		go func(address string) {
			defer wg.Done()
			node := electrum.NewNode()
			if err := node.ConnectTCP(address); err != nil {
				log.Fatal(err)
			}
			for {
			    rand.Read(padded)

				privkey, public := btcec.PrivKeyFromBytes(btcec.S256(), padded)
                                wifu, _ := btcutil.NewWIF(privkey, &chaincfg.MainNetParams, false)
				uaddr, _ := btcutil.NewAddressPubKey(public.SerializeUncompressed(), &chaincfg.MainNetParams)
				
				transactions, err := node.BlockchainAddressGetHistory(uaddr.EncodeAddress())
				if err != nil {
					fmt.Printf("error with address %s: %s", uaddr.EncodeAddress(), err)
				}
				if len(transactions) == 0 {
					fmt.Printf("address %s has no history..\n", uaddr.EncodeAddress())
				} else {
					fmt.Printf("address %s has use history!!!!!!!!! (private key %x)\n", uaddr.EncodeAddress(), wifu.String())
					file, err := os.OpenFile("needles.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
					if err != nil {
						fmt.Printf("Error pushing data to file: %s", err)
						os.Exit(666)
					}
					if _, err := file.WriteString(uaddr.EncodeAddress() + " " + wifu.String() + "\n"); err != nil {
						fmt.Printf("Error pushing data to file: %s", err)
						os.Exit(666)
					}
				}
			}
		}(line)
	}
	wg.Wait()
}
