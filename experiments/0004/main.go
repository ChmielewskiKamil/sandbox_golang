package main

import (
	"context"
	"log"
	"os/exec"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	// "math"
	// "math/big"
)

const (
	AlchemyURL = "wss://eth-mainnet.g.alchemy.com/v2/aUaZPeLamCpNvlbhw7g-_bvnHwDr91zb"
	myAddr     = "0xFA3C26C27dB82e2CdE32d8fE7F2Ae1AD2227CB68"
	port       = "8545"
)

func main() {
	err := run()
	if err != nil {
		log.Fatalf("Exiting because: %v", err)
	}
}

func run() (err error) {
	cmd, err := startAnvil(AlchemyURL, port)
	if err != nil {
		return err
	}
	defer closeAnvil(cmd)

	client, err := ethclient.Dial("http://localhost:" + port)
	if err != nil {
		return logAndErr("Failed to connect the ethclient", err)
	}

	account := common.HexToAddress(myAddr)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return logAndErr("Failed to query the balance", err)
	}
	log.Printf("Balance: %d", balance)

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		return logAndErr("Failed to get the latest block", err)
	}
	log.Printf("Number of TXs: %d", len(block.Transactions()))

	count, err := client.TransactionCount(context.Background(), block.Hash())
	if count == 0 || err != nil {
		return logAndErr("Failed to get the transaction count", err)
	}
	tx := block.Body().Transactions[0]
	log.Printf("TX Hash: %s", tx.Hash().Hex())

	sender, err := types.Sender(types.NewCancunSigner(tx.ChainId()), tx)
	if err != nil {
		return logAndErr("Failed to get the sender", err)
	}
	log.Printf("Sender: %s", sender.Hex())

	return nil
}

// startAnvil starts the local chain with the given RPC URL and port. It returns
// the handle to the process so that the caller can defer the cleanup.
func startAnvil(rpcURL, port string) (*exec.Cmd, error) {
	cmd := exec.Command("anvil", "--fork-url", rpcURL, "--port", port)
	if err := cmd.Start(); err != nil {
		return nil, logAndErr("Failed to start anvil", err)
	}
	log.Printf("Anvil process starting up with PID: %d", cmd.Process.Pid)
	log.Printf("Sleeping for 5 seconds to let Anvil start up...")
	time.Sleep(5 * time.Second)
	return cmd, nil
}

// closeAnvil kills the Anvil process. Defer this function after starting Anvil.
// It will log and exit if the process cannot be killed.
func closeAnvil(cmd *exec.Cmd) {
	if err := cmd.Process.Kill(); err != nil {
		log.Fatalf("Failed to kill Anvil process: %s", err)
	}
	log.Printf("Terminated Anvil process with PID: %d", cmd.Process.Pid)
}

// logAndErr is a helper function that logs the error and returns it
// e.g. return logAndErr("Failed to do something", err) would print:
// "Failed to do something: <err>"
func logAndErr(msg string, err error) error {
	log.Printf("%s: %v", msg, err)
	return err
}
