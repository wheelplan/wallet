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

		btcCheck := strings.HasSuffix(btcAddr, "MyCoin") || strings.HasSuffix(btcAddr, "MyLove") || strings.HasSuffix(btcAddr, "MyDream") ||
			strings.HasSuffix(btcAddr, "MYBTC") || strings.HasSuffix(btcAddr, "China") || strings.HasSuffix(btcAddr, "Bitcoin") ||
			strings.HasSuffix(btcAddr, "LiuCan") || strings.HasSuffix(btcAddr, "MyBTC") || strings.HasSuffix(btcAddr, "XXXXX") ||
			strings.HasSuffix(btcAddr, "5201314") || strings.HasSuffix(btcAddr, "Lucky") || strings.HasSuffix(btcAddr, "5201314")

		ethAddrCut := ethAddr[2:6] + ethAddr[38:]
		ethCheck := strings.HasSuffix(ethAddrCut, "00000000") || strings.HasSuffix(ethAddrCut, "00000001") || strings.HasSuffix(ethAddrCut, "00000003") ||
			strings.HasSuffix(ethAddrCut, "00000006") || strings.HasSuffix(ethAddrCut, "00000007") || strings.HasSuffix(ethAddrCut, "00000008") ||
			strings.HasSuffix(ethAddrCut, "00000009") || strings.HasSuffix(ethAddrCut, "00001314") || strings.HasSuffix(ethAddrCut, "00002020")

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
