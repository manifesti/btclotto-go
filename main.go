package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/d4l3k/go-electrum/electrum"
)

const privatekeyConstant = "115792089237316195423570985008687907853269984665640564039457584007913129639936"

func main() {
	var wg sync.WaitGroup
	// rng from unix time
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	// const string to big.Int
	maxprivatekeys := new(big.Int)
	maxprivatekeys.SetString(privatekeyConstant, 10)

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
		go func(address string, rng *rand.Rand, maxprivatekeys *big.Int) {
			defer wg.Done()
			node := electrum.NewNode()
			if err := node.ConnectTCP(address); err != nil {
				log.Fatal(err)
			}
			for {
				pktocheck := new(big.Int)
				pktocheck.Rand(rng, maxprivatekeys)

				pktodisplay := make([]byte, 32)

				copy(pktodisplay[32-len(pktocheck.Bytes()):], pktocheck.Bytes())

				_, public := btcec.PrivKeyFromBytes(btcec.S256(), pktodisplay)

				// caddr, _ := btcutil.NewAddressPubKey(public.SerializeCompressed(), &chaincfg.MainNetParams)
				uaddr, _ := btcutil.NewAddressPubKey(public.SerializeUncompressed(), &chaincfg.MainNetParams)

				transactions, err := node.BlockchainAddressGetHistory(uaddr.EncodeAddress())
				if err != nil {
					fmt.Printf("error with address %s: %s", uaddr.EncodeAddress(), err)
				}
				if len(transactions) == 0 {
					fmt.Printf("address %s has no history..\n", uaddr.EncodeAddress())
				} else {
					fmt.Printf("address %s has use history!!!!!!!!! (private key %x)\n", uaddr.EncodeAddress(), pktodisplay)
					file, err := os.OpenFile("needles.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
					if err != nil {
						fmt.Printf("Error pushing data to file: %s", err)
						os.Exit(666)
					}
					if _, err := file.WriteString(uaddr.EncodeAddress() + " " + string(pktodisplay) + "\n"); err != nil {
						fmt.Printf("Error pushing data to file: %s", err)
						os.Exit(666)
					}
				}
			}
		}(line, rng, maxprivatekeys)
	}
	wg.Wait()
}
