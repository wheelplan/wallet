package main

import (
	"fmt"
	"github.com/foxnut/go-hdwallet"
	"log"
	"runtime"
	"strings"
)

func main() {

	log.Printf("START 🐻‍❄️...\n\n")

	numCPU := runtime.NumCPU()

	done := make(chan int, 1)

	for i := 0; i < numCPU; i++ {
		go task()
	}

	<-done

}

func task() {

	for {

		mnemonic, _ := hdwallet.NewMnemonic(12, "")
		master, err := hdwallet.NewKey(
			hdwallet.Mnemonic(mnemonic),
		)
		if err != nil {
			panic(err)
		}

		// btc
		wallet, _ := master.GetWallet(hdwallet.Purpose(hdwallet.ZeroQuote+49), hdwallet.CoinType(hdwallet.BTC), hdwallet.AddressIndex(0))
		btcAddr, _ := wallet.GetKey().AddressP2WPKHInP2SH()
		btcWif, err := wallet.GetKey().PrivateWIF(true)
		if err != nil {
			panic(err)
		}

		// eth
		wallet, _ = master.GetWallet(hdwallet.CoinType(hdwallet.ETH))
		ethAddr, _ := wallet.GetAddress()
		ethPriv := wallet.GetKey().PrivateHex()

		btcKeys := []string{"MyCoin", "MyLove", "MyDream", "MyBTC", "China", "Bitcoin", "LiuCan", "Lucky"}

		btcCheck := false
		for _, value := range btcKeys {
			btcCheck = strings.HasSuffix(btcAddr, value)
			if btcCheck {
				break
			}
		}

		ethKeys := []string{"00000000", "00000001", "00000003", "00000006", "00000007", "00000008", "00000009", "00001314", "00002020"}
		ethAddrCut := ethAddr[2:6] + ethAddr[38:]
		ethCheck := false
		for _, value := range ethKeys {
			ethCheck = strings.HasSuffix(ethAddrCut, value)
			if ethCheck {
				break
			}
		}

		if btcCheck || ethCheck {

			log.Println("")
			fmt.Println(mnemonic)

			fmt.Println("BTC PrivateKey ：", btcWif)
			fmt.Println("BTC Address : ", btcAddr)

			fmt.Println("ETH PrivateKey ：", ethPriv)
			fmt.Printf("ETH Address : %s\n\n", ethAddr)

		}
	}

}
