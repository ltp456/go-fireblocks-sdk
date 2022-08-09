package go_fireblocks

import (
	"fmt"
	"go-fireblocks/form"
	"testing"
)

var fbClient *FBClient
var err error

func init() {
	endpoint := "https://api.fireblocks.io"
	apiKey := ""
	secretPath := ""
	fbClient, err = NewFBClient(endpoint, apiKey, secretPath)
	if err != nil {
		panic(err)
	}
}

func TestFBClient_GetVaultAccountsWithPaged(t *testing.T) {
	accountsWithPaged, err := fbClient.GetVaultAccountsWithPaged(ASC, 100, 10, 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(accountsWithPaged.Paging, accountsWithPaged.NextURL, accountsWithPaged.PreviousURL)
	for _, item := range accountsWithPaged.Accounts {
		for _, asset := range item.Assets {
			fmt.Printf("accountId: %v accontName: %v , assetName: %v asset: %v  \n", item.ID, item.Name, asset.ID, asset)
		}
	}
}

func TestFBClient_GetVaultAccounts(t *testing.T) {
	//This endpoint is unavailable. Please use /v1/vault/accounts_paged
	accounts, err := fbClient.GetVaultAccounts()
	if err != nil {
		panic(err)
	}
	for _, item := range accounts {
		for _, asset := range item.Assets {
			fmt.Printf("accountId: %v accontName: %v , assetName: %v asset: %v  \n", item.ID, item.Name, asset.ID, asset)
		}
	}

}

func TestFBClient_VaultAccountById(t *testing.T) {
	account, err := fbClient.VaultAccountById("0")
	if err != nil {
		panic(err)
	}
	for _, item := range account.Assets {
		fmt.Printf("accountId: %v,accountName: %v,assetName: %v, asset: %v \n", account.ID, account.Name, item.ID, item)
	}

}

func TestFBClient_CreateVaultAccount(t *testing.T) {
	err := fbClient.CreateVaultAccount("NFT-Wallet")
	if err != nil {
		panic(err)
	}
}

func TestFBClient_RenameVaultAccount(t *testing.T) {
	resp, err := fbClient.RenameVaultAccount("4", "NFT Wallet")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func TestFBClient_AccountAssetBalance(t *testing.T) {
	assetBalance, err := fbClient.AccountAssetBalance("0", "AVAXTEST")
	if err != nil {
		panic(err)
	}
	fmt.Printf("assetName: %v, asset: %v \n", assetBalance.ID, assetBalance)

}

func TestFBClient_CreateNewVaultWallet(t *testing.T) {
	vaultWallet, err := fbClient.CreateNewVaultWallet("4", "AVAXTEST")
	if err != nil {
		panic(err)
	}
	fmt.Printf("name: %v wallet: %v \n", vaultWallet.ID, vaultWallet)

}

func TestFBClient_RefreshAssetBalances(t *testing.T) {
	assets, err := fbClient.RefreshAssetBalances("4", "AVAXTEST")
	if err != nil {
		panic(err)
	}
	fmt.Printf("name: %v ,asset: %v \n", assets.ID, assets)
}

func TestFBClient_HiddenVaultAccount(t *testing.T) {
	status, err := fbClient.HiddenVaultAccount("4")
	if err != nil {
		panic(err)
	}
	fmt.Println(status)
}

func TestFBClient_ShowVaultAccount(t *testing.T) {
	err := fbClient.ShowVaultAccount("4")
	if err != nil {
		panic(err)
	}

}

func TestFBClient_AssetAddress2(t *testing.T) {
	assetAddress, err := fbClient.AssetAddress("4", "AVAXTEST")
	if err != nil {
		panic(err)
	}
	for _, item := range assetAddress {
		fmt.Printf("assetName: %v ,address: %v \n", item.AssetID, item)
	}
}

func TestFBClient_CreateDepositAddress2(t *testing.T) {
	// todo
	depositAddress, err := fbClient.CreateDepositAddress("4", "BNB_TEST")
	if err != nil {
		panic(err)
	}
	fmt.Println(depositAddress)

}

func TestFBClient_RenameAddress(t *testing.T) {
	// todo
	err := fbClient.RenameAddress("4", "AVAXTEST",
		"0x7ABe9Eb8a9782c1b7c82b4b8EEd026Bac98a59E5", Value{"description", "renameaddr"})
	if err != nil {
		panic(err)
	}
}

func TestFBClient_MaximumSpendableAmount(t *testing.T) {

	err := fbClient.MaximumSpendableAmount("4", "AVAXTEST")
	if err != nil {
		panic(err)
	}

}

func TestFBClient_SetAMLCustomerRefIdForVaultAccount(t *testing.T) {

	status, err := fbClient.SetAMLCustomerRefIdForVaultAccount("4", "test001")
	if err != nil {
		panic(err)
	}
	fmt.Println(status)

}

func TestFBClient_SetAMLCustomerRefIdForAddress(t *testing.T) {
	status, err := fbClient.SetAMLCustomerRefIdForAddress("4", "AVAXTEST", "0x7ABe9Eb8a9782c1b7c82b4b8EEd026Bac98a59E5", "test")
	if err != nil {
		panic(err)
	}
	fmt.Println(status)

}

func TestFBClient_UnspentInput(t *testing.T) {
	//todo
	unspentInput, err := fbClient.UnspentInput("0", "BTC_TEST")
	if err != nil {
		panic(err)
	}
	fmt.Println(unspentInput)
}

func TestFBClient_PublickeyInfo(t *testing.T) {
	publickeyInfo, err := fbClient.PublickeyInfo(MPC_EDDSA_ED25519, []int{}, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(publickeyInfo)
}

func TestFBClient_FireblocksAddressPublickeyInfo(t *testing.T) {
	info, err := fbClient.FireblocksAddressPublickeyInfo("4", "AVAXTEST", "1", "1", true)
	if err != nil {
		panic(err)
	}
	fmt.Println(info)
}

func TestFBClient_SetAutoFuel(t *testing.T) {
	status, err := fbClient.SetAutoFuel("4", false)
	if err != nil {
		panic(err)
	}
	fmt.Println(status)
}

func TestFBClient_VaultAssetsBalances(t *testing.T) {
	assetsBalances, err := fbClient.VaultAssetsBalances()
	if err != nil {
		panic(err)
	}
	for _, item := range assetsBalances {
		fmt.Printf("name: %v ,asset: %v \n", item.ID, item)
	}
}

func TestFBClient_VaultAssetBalance(t *testing.T) {
	assets, err := fbClient.VaultAssetBalance("AVAXTEST")
	if err != nil {
		panic(err)
	}
	fmt.Println(assets)
}

//----vault end --------------

//----internal wallet --------

func TestFBClient_InternalWallets(t *testing.T) {
	wallets, err := fbClient.InternalWallets()
	if err != nil {
		panic(err)
	}
	for _, item := range wallets {
		fmt.Printf("name: %v ,wallet: %v \n", item.Name, item)
	}
}

func TestFBClient_CreateInternalWallet(t *testing.T) {
	internalWallet, err := fbClient.CreateInternalWallet("wallet-group-hot-wallet")
	if err != nil {
		panic(err)
	}
	fmt.Println(internalWallet)
}

func TestFBClient_CreateInternalWalletAsset(t *testing.T) {
	internalWalletAsset, err := fbClient.CreateInternalWalletAsset("217288e9-1f83-476b-a4a9-9bdf67b18d32", "AVAXTEST", "0xce3cbCEAeCeFF12234C591672d532AE549b42C1b")
	if err != nil {
		panic(err)
	}
	fmt.Println(internalWalletAsset)
}

func TestFBClient_InternalWallet(t *testing.T) {
	wallet, err := fbClient.InternalWallet("217288e9-1f83-476b-a4a9-9bdf67b18d32")
	if err != nil {
		panic(err)
	}
	fmt.Println(wallet)
}

func TestFBClient_InternalWalletAsset(t *testing.T) {
	walletAsset, err := fbClient.InternalWalletAsset("217288e9-1f83-476b-a4a9-9bdf67b18d32", "AVAXTEST")
	if err != nil {
		panic(err)
	}
	fmt.Println(walletAsset)

}

func TestFBClient_DeleteInternalWallet(t *testing.T) {
	err := fbClient.DeleteInternalWallet("217288e9-1f83-476b-a4a9-9bdf67b18d32")
	if err != nil {
		panic(err)
	}
}

func TestFBClient_DeleteInternalAsset(t *testing.T) {
	err := fbClient.DeleteInternalAsset("217288e9-1f83-476b-a4a9-9bdf67b18d32", "AVAXTEST")
	if err != nil {
		panic(err)
	}
}

func TestFBClient_SetInternalWalletAMLCustomerRefId(t *testing.T) {
	err := fbClient.SetInternalWalletAMLCustomerRefId("217288e9-1f83-476b-a4a9-9bdf67b18d32", "wallet-test")
	if err != nil {
		panic(err)
	}
}

//---- internal wallet end ------

// -----external wallet ----------

func TestFBClient_ExternalWallets(t *testing.T) {
	wallets, err := fbClient.ExternalWallets()
	if err != nil {
		panic(err)
	}
	for _, item := range wallets {
		fmt.Printf("name: %v ,wallet: %v \n", item.Name, item)
	}
}

func TestFBClient_CreateExternalWallett(t *testing.T) {
	internalWallet, err := fbClient.CreateExternalWallet("wallet-group-external-wallet")
	if err != nil {
		panic(err)
	}
	fmt.Println(internalWallet)
}

func TestFBClient_CreateExternalWalletAsset(t *testing.T) {
	internalWalletAsset, err := fbClient.CreateExternalWalletAsset("f962d132-d573-8237-0986-cd0130b3fee4", "AVAXTEST", "0xce3cbCEAeCeFF12234C591672d532AE549b42C1b")
	if err != nil {
		panic(err)
	}
	fmt.Println(internalWalletAsset)
}

func TestFBClient_ExternalWalletAsset(t *testing.T) {
	walletAsset, err := fbClient.ExternalWalletAsset("f962d132-d573-8237-0986-cd0130b3fee4", "AVAXTEST")
	if err != nil {
		panic(err)
	}
	fmt.Println(walletAsset)

}

func TestFBClient_ExternalWallet(t *testing.T) {
	wallet, err := fbClient.ExternalWallet("f962d132-d573-8237-0986-cd0130b3fee4")
	if err != nil {
		panic(err)
	}
	fmt.Println(wallet)
}

func TestFBClient_DeleteExternalWallet(t *testing.T) {
	err := fbClient.DeleteExternalWallet("f962d132-d573-8237-0986-cd0130b3fee4")
	if err != nil {
		panic(err)
	}
}

func TestFBClient_DeleteExternalAsset(t *testing.T) {
	err := fbClient.DeleteExternalAsset("f962d132-d573-8237-0986-cd0130b3fee4", "AVAXTEST")
	if err != nil {
		panic(err)
	}
}

func TestFBClient_SetExternalWalletAMLCustomerRefId(t *testing.T) {
	err := fbClient.SetExternalWalletAMLCustomerRefId("f962d132-d573-8237-0986-cd0130b3fee4", "wallet-test")
	if err != nil {
		panic(err)
	}
}

// ---- external wallet end -----

// ----- contract wallet -------
func TestFBClient_ContractWallets(t *testing.T) {
	wallets, err := fbClient.ContractWallets()
	if err != nil {
		panic(err)
	}
	for _, item := range wallets {
		fmt.Printf("name: %v ,wallet: %v \n", item.ID, item)
	}
}

func TestFBClient_CreateContractWallet(t *testing.T) {
	internalWallet, err := fbClient.CreateContractWallet("wallet-group-hot-wallet")
	if err != nil {
		panic(err)
	}
	fmt.Println(internalWallet)
}

func TestFBClient_CreateContractWalletAsset(t *testing.T) {
	internalWalletAsset, err := fbClient.CreateContractWalletAsset("217288e9-1f83-476b-a4a9-9bdf67b18d32", "AVAXTEST", "0x...")
	if err != nil {
		panic(err)
	}
	fmt.Println(internalWalletAsset)
}

func TestFBClient_ContractWallet(t *testing.T) {
	wallet, err := fbClient.ContractWallet("217288e9-1f83-476b-a4a9-9bdf67b18d32")
	if err != nil {
		panic(err)
	}
	fmt.Println(wallet)
}

func TestFBClient_ContractWalletAsset(t *testing.T) {
	walletAsset, err := fbClient.ContractWalletAsset("217288e9-1f83-476b-a4a9-9bdf67b18d32", "AVAXTEST")
	if err != nil {
		panic(err)
	}
	fmt.Println(walletAsset)

}

func TestFBClient_DeleteContractWallet(t *testing.T) {
	err := fbClient.DeleteContractWallet("217288e9-1f83-476b-a4a9-9bdf67b18d32")
	if err != nil {
		panic(err)
	}
}

func TestFBClient_DeleteContractsAsset(t *testing.T) {
	err := fbClient.DeleteContractsAsset("217288e9-1f83-476b-a4a9-9bdf67b18d32", "AVAXTEST")
	if err != nil {
		panic(err)
	}
}

// ---- contract wallet end -----

// --- exchange accounts ------

func TestFBClient_ExchangeAccounts(t *testing.T) {
	accounts, err := fbClient.ExchangeAccounts()
	if err != nil {
		panic(err)
	}
	for _, item := range accounts {
		fmt.Printf("name: %v ,account: %v \n", item.Name, item)
	}
}

func TestFBClient_ExchangeAccountById(t *testing.T) {
	exchangeAccount, err := fbClient.ExchangeAccountById("")
	if err != nil {
		panic(err)
	}
	fmt.Println(exchangeAccount)
}

func TestFBClient_ExchangeAssetById(t *testing.T) {
	assets, err := fbClient.ExchangeAssetById("", "")
	if err != nil {
		panic(err)
	}
	fmt.Println(assets)
}

func TestFBClient_ExchangeInternalTransfer(t *testing.T) {
	source := form.TradingAccountType{}
	dest := form.TradingAccountType{}
	err := fbClient.ExchangeInternalTransfer("", "", "",
		source,
		dest)
	if err != nil {
		panic(err)
	}
}

func TestFBClient_ConvertExchangeAccountFounds(t *testing.T) {
	err := fbClient.ConvertExchangeAccountFounds("", "", "", "")
	if err != nil {
		panic(err)
	}
}

// ------exchange account end ------

// ------ fiat account ----------

func TestFBClient_FiatAccounts(t *testing.T) {
	accounts, err := fbClient.ListFiatAccounts()
	if err != nil {
		panic(err)
	}
	for _, item := range accounts {
		fmt.Printf("name: %v ,account: %v \n", item.ID, item)
	}
}

func TestFBClient_FiatAccounts2(t *testing.T) {
	account, err := fbClient.FiatAccounts("")
	if err != nil {
		panic(err)
	}
	fmt.Println(account)
}

func TestFBClient_RedeemFundsToLinkedDDA(t *testing.T) {
	err := fbClient.RedeemFundsToLinkedDDA("", "")
	if err != nil {
		panic(err)
	}
}

func TestFBClient_DepositFundsFromLinkedDDA(t *testing.T) {
	err := fbClient.DepositFundsFromLinkedDDA("", "")
	if err != nil {
		panic(err)
	}
}

// ----- fiat account end -------------

// ------- transaction ---------

func TestFBClient_Transactions(t *testing.T) {
	transactions, err := fbClient.Transactions()
	if err != nil {
		panic(err)
	}
	for _, tx := range transactions {
		fmt.Printf("txId: %v,assetId: %v ,from: %v, to: %v tx: %v \n", tx.ID, tx.AssetID, tx.Source, tx.Destination, tx)
	}

}

func TestFBClient_CreateTransaction(t *testing.T) {
	source := form.Source{
		Type: form.VaultAccount,
		ID:   "0",
	}
	dest := form.Destination{
		Type: form.OneTimeAddressType,
		OneTimeAddress: form.OneTimeAddress{
			Address: "0x5f08C57E89B8206ec629b9c10b1E4e7abC109f29",
		},
	}
	transaction, err := fbClient.CreateTransaction(
		"AVAXTEST",
		source,
		dest,
		"0.1",
		false,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(transaction)
}

func TestFBClient_TransactionById(t *testing.T) {
	transaction, err := fbClient.TransactionById("2d16bc48-b419-4c14-a340-f48c32c2c760")
	if err != nil {
		panic(err)
	}
	fmt.Println(transaction)
}

func TestFBClient_TransactionByExternalId(t *testing.T) {
	//todo
	transaction, err := fbClient.TransactionByExternalId("0xd4b7552226d637d796833fbbcaa5b0a30555f17ad65482b365abfddbbd861a37")
	if err != nil {
		panic(err)
	}
	fmt.Println(transaction)
}

func TestFBClient_CancelTransaction(t *testing.T) {
	result, err := fbClient.CancelTransaction("2d16bc48-b419-4c14-a340-f48c32c2c760")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFBClient_DropTransaction(t *testing.T) {
	result, err := fbClient.DropTransaction("2d16bc48-b419-4c14-a340-f48c32c2c760")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFBClient_FreeTransaction(t *testing.T) {
	result, err := fbClient.FreeTransaction("2d16bc48-b419-4c14-a340-f48c32c2c760")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFBClient_UnFreeTransaction(t *testing.T) {
	result, err := fbClient.UnFreeTransaction("")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFBClient_ValidateAddress(t *testing.T) {
	status, err := fbClient.ValidateAddress("ETH_TEST", "0x5f08C57E89B8206ec629b9c10b1E4e7abC109f29")
	if err != nil {
		panic(err)
	}
	fmt.Println(status)
}

func TestFBClient_NetworkFee(t *testing.T) {
	networkFee, err := fbClient.NetworkFee("AVAXTEST")
	if err != nil {
		panic(err)
	}
	fmt.Println(networkFee)
}

func TestFBClient_EstimateTransactionFee(t *testing.T) {
	source := form.Source{
		Type: form.VaultAccount,
		ID:   "0",
	}
	dest := form.Destination{
		Type: form.OneTimeAddressType,
		OneTimeAddress: form.OneTimeAddress{
			Address: "0x5f08C57E89B8206ec629b9c10b1E4e7abC109f29",
		},
	}
	estimateFee, err := fbClient.EstimateTransactionFee(
		"AVAXTEST",
		"0.1",
		source,
		dest,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(estimateFee)
}

func TestFBClient_SetConfirmationThreshold(t *testing.T) {
	threshold, err := fbClient.SetConfirmationThreshold("2d16bc48-b419-4c14-a340-f48c32c2c760", 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(threshold)
}

func TestFBClient_SetConfirmationThresholdByHash(t *testing.T) {
	txConfirm, err := fbClient.SetConfirmationThresholdByHash("0xd4b7552226d637d796833fbbcaa5b0a30555f17ad65482b365abfddbbd861a37", 3)
	if err != nil {
		panic(err)
	}
	fmt.Println(txConfirm)
}

// --------transaction end -------

// ------ network connections -------

func TestFBClient_ListNetwork(t *testing.T) {
	networkConnections, err := fbClient.ListNetwork()
	if err != nil {
		panic(err)
	}
	for _, item := range networkConnections {
		fmt.Printf("name: %v ,item: %v \n", item.ID, item)
	}
}

func TestFBClient_NetworkConnection(t *testing.T) {
	connection, err := fbClient.NetworkConnection("")
	if err != nil {
		panic(err)
	}
	fmt.Println(connection)
}

// ------- support assets ------

func TestFBClient_SupportedAssets(t *testing.T) {
	assets, err := fbClient.SupportedAssets()
	if err != nil {
		panic(err)
	}
	for _, item := range assets {
		fmt.Printf("name: %v ,item: %v \n", item.ID, item)
	}
}

// ----- gas station ----

func TestFBClient_GasStation(t *testing.T) {
	gasStation, err := fbClient.GasStation()
	if err != nil {
		panic(err)
	}
	fmt.Println(gasStation)

}

func TestFBClient_GasStationByAssetId(t *testing.T) {
	gasStation, err := fbClient.GasStationByAssetId("AVAXTEST")
	if err != nil {
		panic(err)
	}
	fmt.Println(gasStation)
}

func TestFBClient_EditGasStation(t *testing.T) {
	station, err := fbClient.EditGasStation("0.0005", "0.1", "10000000")
	if err != nil {
		panic(err)
	}
	fmt.Println(station)
}

func TestFBClient_EditGasStationByAssetId(t *testing.T) {
	gasStation, err := fbClient.EditGasStationByAssetId("AVAXTEST", "0.1", "0.1", "100000")
	if err != nil {
		panic(err)
	}
	fmt.Println(gasStation)
}

// ---- all users -----

func TestFBClient_AllUsers(t *testing.T) {
	users, err := fbClient.AllUsers()
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		fmt.Println(user)
	}
}
