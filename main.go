package main

import (
	"github.com/foxnut/go-hdwallet"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

func main() {

	log.Printf("%v START 🦁🦏🦛🐘🐯️...\n\n", time.Now().Format("2006-01-02 15:04:05"))

	numCPU := runtime.NumCPU()

	done := make(chan int, 1)

	for i := 0; i < numCPU; i++ {
		go task(i)
	}

	<-done

}

func task(cpuNum int) {

	total := 0
	startTime := time.Now()

	for i := 1; true; i++ {

		mnemonic, _ := hdwallet.NewMnemonic(12, "")
		master, err := hdwallet.NewKey(
			hdwallet.Mnemonic(mnemonic),
		)
		if err != nil {
			panic(err)
		}

		// btc
		wallet, _ := master.GetWallet(hdwallet.Purpose(hdwallet.ZeroQuote+49), hdwallet.CoinType(hdwallet.BTC),
			hdwallet.AddressIndex(0))
		btcAddr, _ := wallet.GetKey().AddressP2WPKHInP2SH()
		btcWif, err := wallet.GetKey().PrivateWIF(true)
		if err != nil {
			panic(err)
		}

		// eth
		wallet, _ = master.GetWallet(hdwallet.CoinType(hdwallet.ETH))
		ethAddr, _ := wallet.GetAddress()
		ethPriv := wallet.GetKey().PrivateHex()

		btcKeys := []string{"MyCoin", "MyLove", "MyDream", "MyBTC", "China", "Bitcoin", "LiuCan", "Lucky", "me"}

		btcCheck := false
		for _, value := range btcKeys {
			btcCheck = strings.HasSuffix(btcAddr, value)
			if btcCheck {
				break
			}
		}

		ethCheck := false
		if !btcCheck {
			ethKeys := []string{"00000000", "00000001", "00000003", "00000006", "00000007", "00000008",
				"00000009", "00001314", "00002020"}
			ethAddrCut := ethAddr[2:6] + ethAddr[38:]

			for _, value := range ethKeys {
				ethCheck = strings.HasSuffix(ethAddrCut, value)
				if ethCheck {
					break
				}
			}
		}

		if btcCheck || ethCheck {
			total++
			endTime := time.Since(startTime)

			log.Printf("%v  Cpu-%d  idx:%d  total:%d  rate:%d  runtime:%.6v  rate/t:%.6v",
				time.Now().Format("2006-01-02 15:04:05"), cpuNum, i, total, i/total, endTime.Abs(),
				endTime/time.Duration(total).Abs())

			log.Printf("%s\nBTC Address: %s\nBTC PrivateKey: %s\nETH Address: %s\nETH PrivateKey: %s\n\n",
				mnemonic, btcAddr, btcWif, ethAddr, ethPriv)

		}
	}

}

func init() {

	log.SetFlags(0)
	logFile, err := os.OpenFile(".keys.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panic("Open log file exception.")
	}

	log.SetOutput(logFile)

}
