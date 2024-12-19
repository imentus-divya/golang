package main

import (
	"fmt"

	"github.com/aptos-labs/aptos-go-sdk/bcs"
	"github.com/aptos-labs/aptos-go-sdk"
	"github.com/aptos-labs/aptos-go-sdk/crypto"
)

func get_account(private_key string) *aptos.Account {
	OwnerprivateKeyHex := private_key
	privateKey := &crypto.Ed25519PrivateKey{}
	err := privateKey.FromHex(OwnerprivateKeyHex)
	if err != nil {
		panic("Failed to parse OwnerprivateKeyHex key:" + err.Error())
	}
	authKey := privateKey.AuthKey()
	ownerAccount, err := aptos.NewAccountFromSigner(privateKey, *authKey)
	return ownerAccount
}

func main() {
	client, err := aptos.NewClient(aptos.TestnetConfig)
	if err != nil {
		panic("Failed to create client:" + err.Error())
	}

	ownerAccount := get_account("0x7bb9d6ca22703cbfa939c28908b123227073f60ddb1b889d813cf17222e22536")
	userAccount := get_account("0x1c45103d1c8e03dd535d89dbee4aff7b3c0cdf3a7996e099d035b9793a708a74")
	sponsorAccount := get_account("0xa4e62aa10073df1902aadbf16fd7d1e2bd390130adecc360972f3fffdf4f379c")

	// contract invocation - KGenPoGNFT_V1

	// =========get_admin=========
	KGenPoGNFT_V1 := aptos.ModuleId{Address: ownerAccount.Address, Name: "KGenPoGNFT_V1"}
	var noTypeTags []aptos.TypeTag
	viewResponse, err := client.View(&aptos.ViewPayload{
		Module:   KGenPoGNFT_V1,
		Function: "get_admin",
		ArgTypes: noTypeTags,
		Args:     [][]byte{},
	})
	if err != nil {
		panic("Failed to view fa address:" + err.Error())
	}
	fmt.Println("-----------get_admin response-----------", viewResponse)

	// ==========================MINT=============================
	// admin secondary signer
	// user signer
	// fee payer
	// Build transaction
	// Prepare the arguments for mint_player_nft

	// playerUsername := "player123"
	// avatarCid := "QmSomeAvatarCid"
	// kgenCommunityMemberBadge := "badge1"
	// proofOfHumanBadge := "1" // Using uint8 for proof badges
	// proofOfPlayBadge := "1"
	// proofOfSkillBadge := "1"
	// proofOfCommerceBadge := "1"
	// proofOfSocialBadge := "1"
	// pohScoreData := "42cb8933d3f1cdf454ffb0d13eb3711e"
	// popScoreData := "42cb8933d3f1cdf454ffb0d13eb3711e"
	// poskScoreData := "42cb8933d3f1cdf454ffb0d13eb3711e"
	// pocScoreData := "42cb8933d3f1cdf454ffb0d13eb3711e"
	// posScoreData := "42cb8933d3f1cdf454ffb0d13eb3711e"
	// pogScoreData := "42cb8933d3f1cdf454ffb0d13eb3711e"


	// Construct the function call arguments
	// args := [][]byte{
	// 	[]byte(playerUsername),
	// 	[]byte(avatarCid),
	// 	[]byte(kgenCommunityMemberBadge),
	// 	[]byte(proofOfHumanBadge),
	// 	[]byte(proofOfPlayBadge),
	// 	[]byte(proofOfSkillBadge),
	// 	[]byte(proofOfCommerceBadge),
	// 	[]byte(proofOfSocialBadge),
	// 	[]byte(pohScoreData),
	// 	[]byte(popScoreData),
	// 	[]byte(poskScoreData),
	// 	[]byte(pocScoreData),
	// 	[]byte(posScoreData),
	// 	[]byte(pogScoreData),
	// }

	

	serializedUsername, err := bcs.Serialize("usernameBytes")



	// Construct the EntryFunction payload
	payload := &aptos.EntryFunction{
		Module: aptos.ModuleId{
			Address: ownerAccount.Address, // Use player's address for module (adjust as necessary)
			Name:    "KGenPoGNFT_V1",      // Replace with your actual module name
		},
		Function: "mint_player_nft",
		ArgTypes: []aptos.TypeTag{},
		// Args: args,
		Args: 

	}

	rawTxn, err := client.BuildTransactionMultiAgent(
		userAccount.Address,
		aptos.TransactionPayload{
			Payload: payload,
		},

		aptos.FeePayer(&sponsorAccount.Address),
		aptos.AdditionalSigners([]aptos.AccountAddress{ownerAccount.Address}),
		// aptos.AdditionalSigners(&ownerAccount.Address)

	)
	if err != nil {
		panic("Failed to build transaction:" + err.Error())
	}

	// Sign transaction
	userAuth, err := rawTxn.Sign(userAccount)
	ownerAuth, err := rawTxn.Sign(ownerAccount)
	sponsorAuth, err := rawTxn.Sign(sponsorAccount)

	signedFeePayerTxn, ok := rawTxn.ToFeePayerSignedTransaction(
		userAuth,
		sponsorAuth,
		[]crypto.AccountAuthenticator{
			*ownerAuth, // Adding the extra signer to the list of additional signers
		},
	)
	if !ok {
		panic("Failed to build fee payer signed transaction")
	}
	// Submit and wait for it to complete
	submitResult, err := client.SubmitTransaction(signedFeePayerTxn)
	if err != nil {
		panic("Failed to submit transaction:" + err.Error())
	}
	txnHash := submitResult.Hash
	println("Submitted transaction hash:", txnHash)

	// Wait for the transaction


	_, err = client.WaitForTransaction(txnHash)
	if err != nil {
		panic("Failed to wait for transaction:" + err.Error())
	}
	// Wait for the transaction to complete and get the full response
	txnResponse, err := client.WaitForTransaction(txnHash)
	if err != nil {
		panic("Failed to wait for transaction:" + err.Error())
	}
	println("Submitted transaction hash:", txnResponse)

}
