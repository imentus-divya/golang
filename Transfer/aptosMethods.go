package main
import (
	"fmt"

	"github.com/aptos-labs/aptos-go-sdk"
	"github.com/aptos-labs/aptos-go-sdk/bcs"
	
)
func main(){
	
	client, err := aptos.NewClient(aptos.DevnetConfig)
	if err != nil {
		panic("Failed to create client:" + err.Error())
	}

	// Create accounts locally for alice and bob
	sender, err := aptos.NewEd25519Account()
	if err != nil {
		panic("Failed to create sender:" + err.Error())
	}
	receiver, err := aptos.NewEd25519Account()
	if err != nil {
		panic("Failed to create receiver:" + err.Error())
	}

	fmt.Printf("\n=== Addresses ===\n")
	fmt.Printf("sender: %s\n", sender.Address.String())
	fmt.Printf("receiver:%s\n", receiver.Address.String())

	// Fund the sender with the faucet to create it on-chain
	const FundAmount = 100_000_000;
	const TransferAmount = 100_000_000;

	err =client.Fund(sender.Address,FundAmount)
	if err != nil {
		panic("Failed to fund sender:" + err.Error())
	}

	senderBalance,err := client.AccountAPTBalance(sender.Address)
	if err!=nil{
		panic("Failed to retrieve sender balance:" + err.Error())
	}

	receiverBalance,err := client.AccountAPTBalance(receiver.Address)
	if err!=nil{
		panic("Failed to retrieve sender balance:" + err.Error())
	}

	fmt.Printf("\n=== Initial Balances ===\n")
	fmt.Printf("Sneder: %d\n", senderBalance)
	fmt.Printf("Receiver:%d\n", receiverBalance)

	// BUILD TRANSACTION

	accountBytes, err := bcs.Serialize(&sender.Address)
	if err!= nil{
		panic("Failed to serialize bob's address:" + err.Error())
	}



	amountBytes, err :=bcs.SerializeU64(TransferAmount)
	if err != nil {
		panic("Failed to serialize transfer amount:" + err.Error())
	}

	rawTxn, err := client.BuildTransaction(sender.AccountAddress(), aptos.TransactionPayload{
		Payload: &aptos.EntryFunction{
			Module: aptos.ModuleId{
				Address: aptos.AccountOne,
				Name:    "aptos_account",
			},
			Function: "transfer",
			ArgTypes: []aptos.TypeTag{},
			Args: [][]byte{
				accountBytes,
				amountBytes,
			},
		}},
	)

	if err != nil {
		panic("Failed to build transaction:" + err.Error())
	}

	// SIGN TRANSACTION
	signedTxn, err := rawTxn.SignedTransaction(sender)
	if err != nil {
		panic("Failed to sign transaction:" + err.Error())
	}

	// SUBMIT TRANSACTION
	submitResult, err := client.SubmitTransaction(signedTxn)
	if err != nil {
		panic("Failed to submit transaction:" + err.Error())
	}
	txnHash := submitResult.Hash

	// WAIT FOR TXN CONFIRMATION
	_, err = client.WaitForTransaction(txnHash)
	if err != nil {
		panic("Failed to wait for transaction:" + err.Error())
	}



	senderBalanceAfter,err := client.AccountAPTBalance(sender.Address)
	if err!=nil{
		panic("Failed to retrieve sender balance:" + err.Error())
	}

	receiverBalanceAfter,err := client.AccountAPTBalance(receiver.Address)
	if err!=nil{
		panic("Failed to retrieve sender balance:" + err.Error())
	}

	fmt.Printf("\n=== After Balances ===\n")
	fmt.Printf("Sneder: %d\n", senderBalanceAfter)
	fmt.Printf("Receiver:%d\n", receiverBalanceAfter)
}