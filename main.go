package main

import (
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"os"
	"time"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"

	"github.com/PuerkitoBio/goquery"
)

const privatekeyConstant = "115792089237316195423570985008687907853269984665640564039457584007913129639936"

func main() {

	// rng from unix time
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	// const string to big.Int
	maxprivatekeys := new(big.Int)
	maxprivatekeys.SetString(privatekeyConstant, 10)
	for {
		// random number from 0 to privatekeyConstant
		pktocheck := new(big.Int)
		pktocheck.Rand(rng, maxprivatekeys)

		pktodisplay := make([]byte, 32)

		copy(pktodisplay[32-len(pktocheck.Bytes()):], pktocheck.Bytes())

		_, public := btcec.PrivKeyFromBytes(btcec.S256(), pktodisplay)

		// caddr, _ := btcutil.NewAddressPubKey(public.SerializeCompressed(), &chaincfg.MainNetParams)
		uaddr, _ := btcutil.NewAddressPubKey(public.SerializeUncompressed(), &chaincfg.MainNetParams)

		doc, err := goquery.NewDocument("https://blockchain.info/address/" + uaddr.EncodeAddress())
		if err != nil {
			fmt.Printf("Error connecting to blockchain.info: %s", err)
		}
		doc.Find("#total_received").Each(func(i int, s *goquery.Selection) {
			complete := s.Find("span").Text()
			// final := s.Find("span").Text()
			if complete == "0 BTC" {
				fmt.Printf("no prior transactions by pk %x\n", pktodisplay)
			} else {
				fmt.Printf("possible wallet found in pk %x, added to needles.txt\n", pktodisplay)
				file, err := os.OpenFile("needles.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					fmt.Printf("Error pushing data to file: %s", err)
				}
				if _, err := file.WriteString(uaddr.EncodeAddress() + " " + string(pktodisplay) + "\n"); err != nil {
					log.Fatal(err)
				}
			}
			// fmt.Printf("%x %34s %34s\n", pktodisplay, uaddr.EncodeAddress(), caddr.EncodeAddress())
		})
	}
}
