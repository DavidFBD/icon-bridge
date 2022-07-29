package icon_test

import (
	"context"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/icon-project/icon-bridge/cmd/e2etest/chain"
	"github.com/icon-project/icon-bridge/cmd/e2etest/chain/icon"
	"github.com/icon-project/icon-bridge/common/log"
)

const (
	NID         = "0x7.icon"
	RPC_URI     = "https://berlin.net.solidwallet.io/api/v3/icon_dex"
	GodKey      = "41528d2ae0a203914f39584c6b2ace17b61c1208492be5952177ed8b16b1b99f"
	GodAddr     = "btp://0x7.icon/hx7449df3b7c6981f0e3dbae9488b024f20f62bbed"
	DemoSrcKey  = "4670f9c2a3f8a478806c03fee796fe614f0ceb8ae063184f11a9c23653eaccfe"
	DemoSrcAddr = "btp://0x7.icon/hx6d338536ac11a0a2db06fb21fe8903e617a6764d"

	DemoDstAddr = "btp://0x61.bsc/0x54a1be6CB9260A52B7E2e988Bc143e4c66b81202"
	GodDstAddr  = "btp://0x61.bsc/0x70E789D2f5D469eA30e0525DbfDD5515d6EAd30D"
	BtsAddr     = "btp://0x7.icon/cx79ba68ea1bb4591ef2b835d8a05d4953986f2b4c"
)

// const (
// 	RPC_URI     = "http://localhost:9080/api/v3/icon"
// 	GodKey      = "c4a15fbef04e99892caaa11374b115795c182d290d5d8bd7821a9ef16f4a9bcf"
// 	GodAddr     = "btp://0x613f17.icon/hxad8eec2e167c24020600ddf1acd4d03673d3f49b"
// 	DemoSrcKey  = "f4e8307da2b4fb7ff89bd984cd0613cfcfacac53abe3a1fd5b7378222bafa5b5"
// 	DemoSrcAddr = "btp://0x613f17.icon/hx691ead88bd5945a43c8a1da331ff6dd80e2936ee"
// 	DemoDstAddr = "btp://0x61.bsc/0x0000000000000000000000000000000000000000"
// 	GodDstAddr  = "btp://0x61.bsc/0x70E789D2f5D469eA30e0525DbfDD5515d6EAd30D"
// 	NID         = "0x613f17.icon"
// )

func TestTransferIntraChain(t *testing.T) {
	// godKeyPair, err := getKeyPairFromFile("../../icon.god.wallet.json", "01fe7d1fb8593bf5")
	// if err != nil {
	// 	t.Fatal(err)
	// 	return
	// }
	// t.Log(godKeyPair)

	api, err := getNewApi()
	if err != nil {
		t.Fatal(err)
		return
	}

	amount := new(big.Int)
	amount.SetString("100000000000000", 10)

	for _, coinName := range []string{"ETH"} {
		txnHash, err := api.Transfer(coinName, GodKey, DemoSrcAddr, amount)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("Transaction Hash %v", txnHash)
		res, err := api.WaitForTxnResult(context.TODO(), txnHash)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("Receipt %+v", res)
		for _, lin := range res.ElInfo {
			t.Logf("Log %+v ", lin)
		}

		if val, err := api.GetCoinBalance(coinName, DemoSrcAddr); err != nil {
			t.Fatal(err)
		} else {
			t.Logf("Balance %v", val)
		}
	}
	return
}

func TestReclaim(t *testing.T) {
	rpi, err := getNewApi()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	hash, err := rpi.Reclaim("ETH", DemoSrcKey, big.NewInt(1333035204))
	if err != nil {
		t.Fatal(err)
	}
	res, err := rpi.WaitForTxnResult(context.TODO(), hash)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Hash %v Receipt %+v", hash, res.Raw)
	showBalance(DemoSrcAddr)
}

func TestApprove(t *testing.T) {
	showBalance(DemoSrcAddr)
	coin := "ETH"
	rpi, err := getNewApi()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	amt := new(big.Int)
	amt.SetString("1110000000000", 10)
	approveHash, err := rpi.Approve(coin, DemoSrcKey, amt)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	res, err := rpi.WaitForTxnResult(context.TODO(), approveHash)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Hash %v Receipt %+v", approveHash, res.Raw)
	showBalance(DemoSrcAddr)
}

func TestTransferInterChain(t *testing.T) {
	api, err := getNewApi()
	if err != nil {
		t.Fatal(err)
		return
	}
	coin := "ETH"
	if val, err := api.GetCoinBalance(coin, DemoSrcAddr); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Initial Balance %v", val)
	}

	amount := new(big.Int)
	amount.SetString("110000000000", 10)

	txnHash, err := api.Transfer(coin, DemoSrcKey, DemoDstAddr, amount)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Transaction Hash  %v", txnHash)

	if val, err := api.GetCoinBalance(coin, DemoSrcAddr); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Intermediate Demo Balance %v", val)
	}

	res, err := api.WaitForTxnResult(context.TODO(), txnHash)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Receipt %+v", res)
	for _, lin := range res.ElInfo {
		// seq, _ := lin.GetSeq()
		t.Logf("Log %+v and Seq", lin)
	}
	if val, err := api.GetCoinBalance(coin, DemoSrcAddr); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Final Balance %v", val)
	}
}

func TestBatchTransfer(t *testing.T) {
	rpi, err := getNewApi()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	coins := []string{"TICX", "ETH", "ICX"}
	amount := big.NewInt(100000000000000)
	largeAmt := new(big.Int)
	largeAmt.SetString("26271926117961986739", 10)
	amounts := []*big.Int{amount, largeAmt, amount}
	for i, coin := range coins {

		if coin == rpi.NativeCoin() {
			continue
		}
		approveHash, err := rpi.Approve(coin, GodKey, amounts[i])
		if err != nil {
			t.Fatalf("%+v", err)
		}
		res, err := rpi.WaitForTxnResult(context.TODO(), approveHash)
		if err != nil {
			t.Fatal(err)
		}
		if res.StatusCode != 1 {
			t.Fatalf("Approve StatusCode not 1 for %vth coin %v \n %v", i, coin, res.Raw)
		}
	}
	hash, err := rpi.TransferBatch(coins, GodKey, GodDstAddr, amounts)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Hash ", hash)
	res, err := rpi.WaitForTxnResult(context.TODO(), hash)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 1 {
		t.Fatalf("StatusCode not 1 for Batch Transfer \n %+v", res.Raw)
	}
}

func TestGetCoinBalance(t *testing.T) {
	if err := showBalance(DemoSrcAddr); err != nil {
		t.Fatalf(" %+v", err)
	}
}

func showBalance(addr string) error {
	api, err := getNewApi()
	if err != nil {
		return err
	}

	for _, coinName := range []string{"ICX", "TICX", "ETH", "BNB", "TBNB"} {
		fmt.Println(coinName)
		_, err := api.GetCoinBalance(coinName, addr)
		if err != nil {
			return err
		}
		// log.Infof("coin %v amount %v", coinName, res.String())
	}
	return nil
}

func getKeyPairFromFile(walFile string, password string) (pair [2]string, err error) {
	keyReader, err := os.Open(walFile)
	if err != nil {
		return
	}
	defer keyReader.Close()

	keyStore, err := ioutil.ReadAll(keyReader)
	if err != nil {
		return
	}
	key, err := keystore.DecryptKey(keyStore, password)
	if err != nil {
		return
	}
	privBytes := ethcrypto.FromECDSA(key.PrivateKey)
	privString := hex.EncodeToString(privBytes)
	addr := ethcrypto.PubkeyToAddress(key.PrivateKey.PublicKey)
	pair = [2]string{privString, addr.String()}
	return
}

func TestReceiver(t *testing.T) {
	recv, err := getNewApi()
	if err != nil {
		panic(err)
	}
	// recv.WatchForTransferStart(1, "ICX", 10)
	// recv.WatchForTransferReceived(1, "TONE", 8)
	// recv.WatchForTransferEnd(1, "ICX", 10)

	go func() {
		if sinkChan, errChan, err := recv.Subscribe(context.Background()); err != nil {
			panic(err)
		} else {
			for {
				select {
				case err := <-errChan:
					panic(err)
				case msg := <-sinkChan:
					t.Logf("\nMessage %+v\n", msg)
				}
			}
		}
	}()
	time.Sleep(time.Second * 3000)
}

func getNewApi() (chain.ChainAPI, error) {
	srcEndpoint := RPC_URI
	addrToName := map[chain.ContractName]string{
		chain.BTS: "cxaf53c60578dc72b3e3a93d1800be1b33349289a0",
	}
	l := log.New()
	log.SetGlobalLogger(l)
	return icon.NewApi(l, &chain.Config{
		Name:              chain.ICON,
		URL:               srcEndpoint,
		ContractAddresses: addrToName,
		NetworkID:         NID,
		NativeCoin:        "ICX",
		NativeTokens:      []string{"ETH", "TICX"},
		WrappedCoins:      []string{"BNB", "TBNB"},
		GasLimit:          8000000,
	})
}

func TestConverToZeroAddress(t *testing.T) {
	addr := DemoSrcAddr
	splits := strings.Split(addr, "/")
	if len(splits) != 4 {
		return
	}
	network := splits[2]
	networkSplits := strings.Split(network, ".")
	if len(networkSplits) != 2 {
		return
	}
	networkSplits[1] += "s"
	splits[2] = strings.Join(networkSplits, ".")
	joinStr := strings.Join(splits, "/")
	fmt.Println(joinStr)
}
