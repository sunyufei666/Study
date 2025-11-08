package main

import (
	"DappTask1/count"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/86a446a0a7ce4a1b825ca9779d77bcd9")
	if err != nil {
		panic(err)
	}

	// getBlockInfo(client)
	// createTransaction(client)
	// deployCountNumContract(client)
	execCountContract(client)
}

// 实现查询指定区块号的区块信息，包括区块的哈希、时间戳、交易数量等
func getBlockInfo(client *ethclient.Client) {
	header, err := client.HeaderByNumber(context.Background(), big.NewInt(9513630))
	if err != nil {
		panic(err)
	}
	fmt.Printf("区块的哈希值为：%s\n", header.Hash())
	fmt.Printf("区块的时间戳为：%d\n", header.Time)
	txCount, err := client.TransactionCount(context.Background(), header.Hash())
	if err != nil {
		panic(err)
	}
	fmt.Printf("区块的交易数量为：%d\n", txCount)
}


/*
- 构造一笔简单的以太币转账交易，指定发送方、接收方和转账金额。
- 对交易进行签名，并将签名后的交易发送到网络。
- 输出交易的哈希值
*/
func createTransaction(client *ethclient.Client) {
	privateKey, err := crypto.HexToECDSA("0x7f90122bf0700f9e7e1f688fe926940e8839f35347a3df58c4f45f4c1f1f7d1a")
	if err != nil {
		panic(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	publicAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 当前账户的nonce
	nonce, err := client.PendingNonceAt(context.Background(), publicAddress)
	if err != nil {
		log.Fatal(err)
	}
	// 接收方地址
	toAddress := common.HexToAddress("0x12f371af09d3c3c3acf423f8484396b205391b4d")
	// 转账金额
	value := big.NewInt(1000000000000000000) // 1 ETH
	// 燃气费上限
	gasLimit := uint64(21000)
	// 燃气价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 创建交易
	tx := types.NewTransaction(nonce,toAddress,value, gasLimit, gasPrice,nil)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 对交易进行签名
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}

// 部署计数器合约
// 合约部署地址：0x20bfdDaC15d50c7d830C923f140aa1cF31Ce1d50
// 交易哈希：0x35c21bdd4b274553a4f8d7bbce3c292582e582f5caac503af137c35185b22158
func deployCountNumContract(client *ethclient.Client) {
	privateKey, err := crypto.HexToECDSA("0158ea697ca66972afc9485343bac42c00dfeafb156b1e4c3105710cf6512441")
	if err != nil {
		panic(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	publicAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 获取当前网络的链ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		panic(err)
	}

	// 当前账户的nonce
	nonce, err := client.PendingNonceAt(context.Background(), publicAddress)
	if err != nil {
		panic(err)
	}

	// 燃气价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice
	address, tx, _, err := count.DeployCount(auth, client)
	if err != nil {
		panic(err)
	}
	fmt.Printf("合约部署地址：%s\n", address.Hex())
	fmt.Printf("交易哈希：%s\n", tx.Hash().Hex())
}

// 调用计数器合约测试
func execCountContract(client *ethclient.Client) {
	contractAddress := common.HexToAddress("0x20bfdDaC15d50c7d830C923f140aa1cF31Ce1d50")
	// 合约实例
	countContract, err := count.NewCount(contractAddress, client)
	if err != nil {
		panic(nil)
	}
	// 私钥
	privateKey, err := crypto.HexToECDSA("0158ea697ca66972afc9485343bac42c00dfeafb156b1e4c3105710cf6512441")
	if err != nil {
		panic(err)
	}
	// 获取链id
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		panic(err)
	}
	// 初始化opt交易实例
	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	tx, err := countContract.Add(opt)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Add count Transaction sent: %s\n", tx.Hash().Hex())
	// 查询count值
	callOpt := &bind.CallOpts{Context: context.Background()}
	countValue, err := countContract.Count(callOpt)
	if err != nil {
		panic(nil)
	}
	fmt.Printf("count经过加1之后的值为：%d", countValue)
}