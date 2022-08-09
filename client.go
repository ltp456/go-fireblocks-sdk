package go_fireblocks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-fireblocks/form"
	"go-fireblocks/types"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
)

type FBClient struct {
	tokenProvider *TokenProvider
	imp           *http.Client
	endpoint      string
}

func NewFBClient(endpoint, apiKey, secretPath string) (*FBClient, error) {
	secretBytes, err := os.ReadFile(secretPath)
	if err != nil {
		return nil, err
	}
	fbClient := &FBClient{
		tokenProvider: NewTokenProvider(apiKey, secretBytes),
		endpoint:      endpoint,
		imp:           http.DefaultClient,
	}
	return fbClient, nil
}

// vault endpoint

func (fb *FBClient) GetVaultAccounts(options ...Value) ([]types.VaultAccounts, error) {
	var vaultAccounts []types.VaultAccounts
	err := fb.Get("/v1/vault/accounts", Params{}, &vaultAccounts, options...)
	if err != nil {
		return nil, err
	}
	return vaultAccounts, nil
}

func (fb *FBClient) GetVaultAccountsWithPaged(orderBy PaseSort, limit, mbaiu, mbcaiu int, optValues ...Value) (*types.VaultAccountsPaged, error) {
	values := Params{}
	values.SetValue("orderBy", orderBy)
	values.SetValue("limit", limit)
	values.SetValue("maxBip44AddressIndexUsed", mbaiu)
	values.SetValue("maxBip44ChangeAddressIndexUsed", mbcaiu)
	vaultAccountsPaged := &types.VaultAccountsPaged{}
	err := fb.Get("/v1/vault/accounts_paged", values, vaultAccountsPaged, optValues...)
	if err != nil {
		return nil, err
	}
	return vaultAccountsPaged, nil
}

func (fb *FBClient) CreateVaultAccount(accountName string, options ...Value) error {
	values := Params{}
	values.SetValue("name", accountName)
	var resp string
	err := fb.Post("/v1/vault/accounts", values, resp, options...)
	if err != nil {
		return err
	}
	return nil
}

func (fb *FBClient) VaultAccountById(vaultAccountId string) (*types.Accounts, error) {
	result := &types.Accounts{}
	err := fb.Get(fmt.Sprintf("/v1/vault/accounts/%s", vaultAccountId), nil, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) RenameVaultAccount(vaultAccountId, accountName string) (*types.ReAccountName, error) {
	params := Params{}
	params.SetValue("name", accountName)
	result := &types.ReAccountName{}
	err := fb.Put(fmt.Sprintf("/v1/vault/accounts/%s", vaultAccountId), params, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) CreateNewVaultWallet(vaultAccountId, assetId string, options ...Value) (*types.VaultAssetResp, error) {
	result := &types.VaultAssetResp{}
	err := fb.Post(fmt.Sprintf("/v1/vault/accounts/%s/%s", vaultAccountId, assetId), Params{}, result, options...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) AccountAssetBalance(vaultAccountId, assertId string) (*types.AccountAssert, error) {
	assertBalance := &types.AccountAssert{}
	err := fb.Get(fmt.Sprintf("/v1/vault/accounts/%s/%s", vaultAccountId, assertId), nil, assertBalance)
	if err != nil {
		return nil, err
	}
	return assertBalance, nil
}

func (fb *FBClient) AssetById(assetId string) (*types.Assets, error) {
	result := &types.Assets{}
	err := fb.Get(fmt.Sprintf("/v1/vault/assets/%s", assetId), nil, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) RefreshAssetBalances(vaultAccountId, assetId string) (*types.VaultAssets, error) {
	result := &types.VaultAssets{}
	err := fb.Post(fmt.Sprintf("/v1/vault/accounts/%s/%s/balance", vaultAccountId, assetId), Params{}, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) HiddenVaultAccount(vaultAccountId string) (*types.StatusResp, error) {
	result := &types.StatusResp{}
	err := fb.Post(fmt.Sprintf("/v1/vault/accounts/%s/hide", vaultAccountId), Params{}, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) ShowVaultAccount(vaultAccountId string) error {
	var result string
	err := fb.Post(fmt.Sprintf("/v1/vault/accounts/%s/unhide", vaultAccountId), Params{}, &result)
	if err != nil {
		return err
	}
	return nil
}

func (fb *FBClient) CreateDepositAddress(vaultAccountId, assetId string, options ...Value) (*types.CreateAddress, error) {
	resp := &types.CreateAddress{}
	err := fb.Post(fmt.Sprintf("/v1/vault/accounts/%s/%s/addresses", vaultAccountId, assetId), Params{}, resp, options...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (fb *FBClient) AssetAddress(vaultAccountId, assetId string) ([]types.AssetAddress, error) {
	var result []types.AssetAddress
	err := fb.Get(fmt.Sprintf("/v1/vault/accounts/%s/%s/addresses", vaultAccountId, assetId), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) RenameAddress(vaultAccountId, assetId, addressId string, options ...Value) error {
	var result string
	err := fb.Put(fmt.Sprintf("/v1/vault/accounts/%s/%s/addresses/%s", vaultAccountId, assetId, addressId), Params{}, &result, options...)
	if err != nil {
		return err
	}
	return nil
}

// MaximumSpendableAmount 仅与utxo相关
func (fb *FBClient) MaximumSpendableAmount(vaultAccountId, assetId string, options ...Value) error {
	var result string
	err := fb.Put(fmt.Sprintf("/v1/vault/accounts/%s/%s/max_spendable_amount", vaultAccountId, assetId), Params{}, &result, options...)
	if err != nil {
		return err
	}
	return nil
}

func (fb *FBClient) SetAMLCustomerRefIdForVaultAccount(vaultAccountId, customerRefId string) (*types.StatusResp, error) {
	params := Params{}
	params.SetValue("customerRefId", customerRefId)
	result := &types.StatusResp{}
	err := fb.Post(fmt.Sprintf("/v1/vault/accounts/%s/set_customer_ref_id", vaultAccountId), Params{}, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) SetAMLCustomerRefIdForAddress(vaultAccountId, assetId, addressId, customerRefId string) (*types.StatusResp, error) {
	params := Params{}
	params.SetValue("customerRefId", customerRefId)
	result := &types.StatusResp{}
	err := fb.Post(fmt.Sprintf("/v1/vault/accounts/%s/%s/addresses/%s/set_customer_ref_id", vaultAccountId, assetId, addressId), Params{}, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) UnspentInput(vaultAccountId, assetId string) ([]types.UnspentInputData, error) {
	var result []types.UnspentInputData
	err := fb.Get(fmt.Sprintf("/v1/vault/accounts/%s/%s/unspent_inputs", vaultAccountId, assetId), Params{}, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) PublickeyInfo(algorithm SigningAlgorithm, derivationPath []int, compressed bool) (*types.PublicKeyInfo, error) {
	params := Params{}
	params.SetValue("algorithm", algorithm)
	params.SetValue("derivationPath", derivationPath)
	params.SetValue("compressed", compressed)
	result := &types.PublicKeyInfo{}
	err := fb.Get("/v1/vault/public_key_info", params, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) FireblocksAddressPublickeyInfo(vaultAccountId, assetId, change, addressIndex string, compressed bool) (*types.PublicKeyInfo, error) {
	params := Params{}
	params.SetValue("compressed", compressed)
	result := &types.PublicKeyInfo{}
	err := fb.Get(fmt.Sprintf("/v1/vault/accounts/%v/%v/%v/%v/public_key_info", vaultAccountId, assetId, change, addressIndex), params, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) SetAutoFuel(vaultAccountId string, autoFuel bool) (*types.StatusResp, error) {
	params := Params{}
	params.SetValue("autoFuel", autoFuel)
	result := &types.StatusResp{}
	err := fb.Post(fmt.Sprintf("/v1/vault/accounts/%s/set_auto_fuel", vaultAccountId), params, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) VaultAssetBalance(assetId string) (*types.Assets, error) {
	result := &types.Assets{}
	err := fb.Get(fmt.Sprintf("/v1/vault/assets/%s", assetId), nil, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) VaultAssetsBalances(options ...Value) ([]types.Assets, error) {
	var result []types.Assets
	err := fb.Get("/v1/vault/assets", Params{}, &result, options...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// contract wallets  endpoint

func (fb *FBClient) ContractWallets() ([]types.WalletAssets, error) {
	var result []types.WalletAssets
	err := fb.Get("/v1/contracts", Params{}, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) CreateContractWallet(name string, options ...Value) (*types.Wallet, error) {
	result := &types.Wallet{}
	params := Params{}
	params.SetValue("name", name)
	err := fb.Post("/v1/contracts", params, result, options...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) CreateContractWalletAsset(walletId, assetId, address string, options ...Value) (*types.AssetInfo, error) {
	params := Params{}
	params.SetValue("address", address)
	result := &types.AssetInfo{}
	err := fb.Post(fmt.Sprintf("/v1/contracts/%s/%s", walletId, assetId), params, result, options...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) ContractWallet(walletId string) (*types.Wallet, error) {
	result := &types.Wallet{}
	err := fb.Get(fmt.Sprintf("/v1/contracts/%s", walletId), nil, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) ContractWalletAsset(walletId, assetId string) (*types.WalletAssets, error) {
	result := &types.WalletAssets{}
	err := fb.Get(fmt.Sprintf("/v1/contracts/%s/%s", walletId, assetId), nil, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) DeleteContractWallet(walletId string) error {
	var result string
	err := fb.Get(fmt.Sprintf("/v1/contracts/%s", walletId), nil, &result)
	if err != nil {
		return err
	}
	return nil

}

func (fb *FBClient) DeleteContractsAsset(walletId, assetId string) error {
	var result string
	err := fb.Delete(fmt.Sprintf("/v1/contracts/%s/%s", walletId, assetId), nil, &result)
	if err != nil {
		return err
	}
	return nil

}

// external wallet  endpoint

func (fb *FBClient) ExternalWallets() ([]*types.Wallet, error) {
	var result []*types.Wallet
	err := fb.Get("/v1/external_wallets", nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func (fb *FBClient) ExternalWallet(walletId string) (*types.Wallet, error) {
	result := &types.Wallet{}
	err := fb.Get(fmt.Sprintf("/v1/external_wallets/%s", walletId), nil, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) ExternalWalletAsset(walletId, assetId string) (*types.WalletAssets, error) {
	result := &types.WalletAssets{}
	err := fb.Get(fmt.Sprintf("/v1/external_wallets/%s/%s", walletId, assetId), nil, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) CreateExternalWallet(name string, options ...Value) (*types.Wallet, error) {
	result := &types.Wallet{}
	params := Params{}
	params.SetValue("name", name)
	err := fb.Post("/v1/external_wallets", params, result, options...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) CreateExternalWalletAsset(walletId, assetId, address string, options ...Value) (*types.AssetInfo, error) {
	params := Params{}
	params.SetValue("address", address)
	result := &types.AssetInfo{}
	err := fb.Post(fmt.Sprintf("/v1/external_wallets/%s/%s", walletId, assetId), params, result, options...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) DeleteExternalWallet(walletId string) error {
	var result string
	err := fb.Get(fmt.Sprintf("/v1/external_wallets/%s", walletId), nil, &result)
	if err != nil {
		return err
	}
	return nil

}

func (fb *FBClient) DeleteExternalAsset(walletId, assetId string) error {
	var result string
	err := fb.Delete(fmt.Sprintf("/v1/external_wallets/%s/%s", walletId, assetId), nil, &result)
	if err != nil {
		return err
	}
	return nil

}

func (fb *FBClient) SetExternalWalletAMLCustomerRefId(walletId, customerRefId string) error {
	params := Params{}
	params.SetValue("customerRefId", customerRefId)
	var result string
	err := fb.Post(fmt.Sprintf("/v1/external_wallets/%s/set_customer_ref_id", walletId), params, &result)
	if err != nil {
		return err
	}
	return nil
}

// internal wallet endpoint

func (fb *FBClient) InternalWallets() ([]*types.Wallet, error) {
	var result []*types.Wallet
	err := fb.Get("/v1/internal_wallets", Params{}, &result)
	if err != nil {
		return nil, err
	}
	return result, nil

}
func (fb *FBClient) InternalWallet(walletId string) (*types.Wallet, error) {
	result := &types.Wallet{}
	err := fb.Get(fmt.Sprintf("/v1/internal_wallets/%s", walletId), Params{}, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) InternalWalletAsset(walletId, assetId string) (*types.WalletAssets, error) {
	result := &types.WalletAssets{}
	err := fb.Get(fmt.Sprintf("/v1/internal_wallets/%s/%s", walletId, assetId), Params{}, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) CreateInternalWallet(name string, options ...Value) (*types.Wallet, error) {
	result := &types.Wallet{}
	params := Params{}
	params.SetValue("name", name)
	err := fb.Post("/v1/internal_wallets", params, result, options...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) CreateInternalWalletAsset(walletId, assetId, address string, options ...Value) (*types.AssetInfo, error) {
	params := Params{}
	params.SetValue("address", address)
	result := &types.AssetInfo{}
	err := fb.Post(fmt.Sprintf("/v1/internal_wallets/%s/%s", walletId, assetId), params, result, options...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) DeleteInternalWallet(walletId string) error {
	var result string
	err := fb.Get(fmt.Sprintf("/v1/internal_wallets/%s", walletId), Params{}, &result)
	if err != nil {
		return err
	}
	return nil

}

func (fb *FBClient) DeleteInternalAsset(walletId, assetId string) error {
	var result string
	err := fb.Delete(fmt.Sprintf("/v1/internal_wallets/%s/%s", walletId, assetId), Params{}, &result)
	if err != nil {
		return err
	}
	return nil

}

func (fb *FBClient) SetInternalWalletAMLCustomerRefId(walletId, customerRefId string) error {
	params := Params{}
	params.SetValue("customerRefId", customerRefId)
	var result string
	err := fb.Post(fmt.Sprintf("/v1/internal_wallets/%s/set_customer_ref_id", walletId), params, &result)
	if err != nil {
		return err
	}
	return nil
}

// exchange accounts endpoint

func (fb *FBClient) ExchangeAccounts() ([]types.ExchangeAccount, error) {
	var result []types.ExchangeAccount
	err := fb.Get("/v1/exchange_accounts", nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) ExchangeAccountById(eaId string) (*types.ExchangeAccount, error) {
	result := &types.ExchangeAccount{}
	err := fb.Get(fmt.Sprintf("/v1/exchange_accounts/%s", eaId), nil, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) ExchangeAssetById(exchangeAccountId, assetId string) (*types.Assets, error) {
	result := &types.Assets{}
	err := fb.Get(fmt.Sprintf("/v1/exchange_accounts/%s/%s", exchangeAccountId, assetId), nil, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) ExchangeInternalTransfer(exchangeAccountId, asset, amount string, source, dest form.TradingAccountType) error {
	params := Params{}
	params.SetValue("asset", asset)
	params.SetValue("amount", amount)
	params.SetValue("source", source)
	params.SetValue("dest", dest)
	var result string
	err := fb.Post(fmt.Sprintf("/v1/exchange_accounts/%s/internal_transfer", exchangeAccountId), params, &result)
	if err != nil {
		return err
	}
	return nil
}

func (fb *FBClient) ConvertExchangeAccountFounds(exchangeAccountId, srcAsset, destAsset, amount string) error {
	params := Params{}
	params.SetValue("srcAsset", srcAsset)
	params.SetValue("destAsset", destAsset)
	params.SetValue("amount", amount)
	var result string
	err := fb.Post(fmt.Sprintf("/v1/exchange_accounts/%s/convert", exchangeAccountId), params, &result)
	if err != nil {
		return err
	}
	return nil
}

// Fiat accounts endpoint

func (fb *FBClient) ListFiatAccounts() ([]types.FiatAccount, error) {
	var result []types.FiatAccount
	err := fb.Get("/v1/fiat_accounts", nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) FiatAccounts(accountId string) (*types.FiatAccount, error) {
	result := &types.FiatAccount{}
	err := fb.Get(fmt.Sprintf("/v1/fiat_accounts/%s", accountId), nil, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) RedeemFundsToLinkedDDA(accountId, amount string) error {
	params := Params{}
	params.SetValue("amount", amount)
	var result string
	err := fb.Post(fmt.Sprintf("/v1/fiat_accounts/%s/redeem_to_linked_dda", accountId), params, &result)
	if err != nil {
		return err
	}
	return nil
}

func (fb *FBClient) DepositFundsFromLinkedDDA(accountId, amount string) error {
	params := Params{}
	params.SetValue("amount", amount)
	var result string
	err := fb.Post(fmt.Sprintf("/v1/fiat_accounts/%s/deposit_from_linked_dda", accountId), params, &result)
	if err != nil {
		return err
	}
	return nil
}

// network connections endpoint

func (fb *FBClient) ListNetwork() ([]types.NetworkConnections, error) {
	var result []types.NetworkConnections
	err := fb.Get("/v1/network_connections", nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) NetworkConnection(connectionId string) (*types.NetworkConnections, error) {
	result := &types.NetworkConnections{}
	err := fb.Get(fmt.Sprintf("/v1/network_connections/%s", connectionId), nil, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// users endpoint

func (fb *FBClient) AllUsers() ([]types.AllUsers, error) {
	var result []types.AllUsers
	err := fb.Get("/v1/users", nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// gas endpoint

func (fb *FBClient) EditGasStationByAssetId(assetId, gasThreshold, gasCap, maxGasPrice string) (*types.StatusResp, error) {
	values := Params{}
	values.SetValue("gasThreshold", gasThreshold)
	values.SetValue("gasCap", gasCap)
	values.SetValue("maxGasPrice", maxGasPrice)
	result := &types.StatusResp{}
	err := fb.Put(fmt.Sprintf("/v1/gas_station/configuration/%s", assetId), values, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) EditGasStation(gasThreshold, gasCap, maxGasPrice string) (*types.StatusResp, error) {
	values := Params{}
	values.SetValue("gasThreshold", gasThreshold)
	values.SetValue("gasCap", gasCap)
	values.SetValue("maxGasPrice", maxGasPrice)
	result := &types.StatusResp{}
	err := fb.Put("/v1/gas_station/configuration", values, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) GasStation() (*types.GasStation, error) {
	result := &types.GasStation{}
	err := fb.Get("/v1/gas_station", Params{}, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) GasStationByAssetId(assetId string) (*types.GasStation, error) {
	result := &types.GasStation{}
	err := fb.Get(fmt.Sprintf("/v1/gas_station/%s", assetId), nil, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// support assets endpoint

func (fb *FBClient) SupportedAssets() ([]types.SupportedAssets, error) {
	var result []types.SupportedAssets
	err := fb.Get("/v1/supported_assets", nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// transactions endpoint

func (fb *FBClient) SetConfirmationThreshold(txId string, confirms int) (*types.TxConfirm, error) {
	params := Params{}
	params.SetValue("numOfConfirmations", confirms)
	result := &types.TxConfirm{}
	err := fb.Post(fmt.Sprintf("/v1/transactions/%s/set_confirmation_threshold", txId), params, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) SetConfirmationThresholdByHash(txHash string, confirms int) (*types.TxConfirm, error) {
	params := Params{}
	params.SetValue("numOfConfirmations", confirms)
	result := &types.TxConfirm{}
	err := fb.Post(fmt.Sprintf("/v1/txHash/%s/set_confirmation_threshold", txHash), params, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) EstimateTransactionFee(assetId, amount string, source form.Source,
	destination form.Destination, options ...Value) (*types.EstimateFee, error) {
	values := Params{}
	values.SetValue("assetId", assetId)
	values.SetValue("amount", amount)
	values.SetValue("source", source)
	values.SetValue("destination", destination)
	result := &types.EstimateFee{}
	err := fb.Post("/v1/transactions/estimate_fee", values, result, options...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) NetworkFee(assetId string) (*types.EstimateFee, error) {
	values := Params{}
	values.SetValue("assetId", assetId)
	result := &types.EstimateFee{}
	err := fb.Get("/v1/estimate_network_fee", values, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) CreateTransaction(assetId string, source form.Source, destination form.Destination,
	amount string, treatAsGrossAmount bool, options ...Value) (*types.Status, error) {
	values := Params{}
	values.SetValue("assetId", assetId)
	values.SetValue("source", source)
	values.SetValue("destination", destination)
	values.SetValue("amount", amount)
	values.SetValue("treatAsGrossAmount", treatAsGrossAmount)
	status := &types.Status{}
	err := fb.Post("/v1/transactions", values, status, options...)
	if err != nil {
		return nil, err
	}
	return status, nil
}

func (fb *FBClient) CancelTransaction(fbId string) (*types.OperationStatus, error) {
	result := &types.OperationStatus{}
	err := fb.Post(fmt.Sprintf("/v1/transactions/%s/cancel", fbId), Params{}, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) DropTransaction(fbId string, options ...Value) (*types.OperationStatus, error) {
	result := &types.OperationStatus{}
	err := fb.Post(fmt.Sprintf("/v1/transactions/%s/drop", fbId), Params{}, result, options...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) FreeTransaction(fbId string) (*types.OperationStatus, error) {
	result := &types.OperationStatus{}
	err := fb.Post(fmt.Sprintf("/v1/transactions/%s/freeze", fbId), Params{}, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) UnFreeTransaction(fbId string) (*types.OperationStatus, error) {
	result := &types.OperationStatus{}
	err := fb.Post(fmt.Sprintf("/v1/transactions/%s/unfreeze", fbId), Params{}, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ValidateAddress XRP, DOT, XLM, EOS.
func (fb *FBClient) ValidateAddress(assetId, address string) (*types.AddressStatus, error) {
	result := &types.AddressStatus{}
	err := fb.Get(fmt.Sprintf("/v1/transactions/validate_address/%s/%s", assetId, address), nil, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) TransactionById(fireblocksId string) (*types.Transaction, error) {
	txInfo := &types.Transaction{}
	err := fb.Get(fmt.Sprintf("/v1/transactions/%s", fireblocksId), nil, txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

func (fb *FBClient) TransactionByExternalId(txId string) (*types.Transaction, error) {
	txInfo := &types.Transaction{}
	err := fb.Get(fmt.Sprintf("/v1/transactions/external_tx_id/%s", txId), nil, txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

func (fb *FBClient) Transactions(options ...Value) ([]types.Transaction, error) {
	var result []types.Transaction
	err := fb.Get("/v1/transactions", Params{}, &result, options...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fb *FBClient) Post(method string, param Params, value interface{}, options ...Value) error {
	for _, v := range options {
		if param == nil {
			break
		}
		param.SetValue(v.Key, v.Value)
	}
	return fb.httpReq(http.MethodPost, method, param, value)
}

func (fb *FBClient) Put(method string, param Params, value interface{}, options ...Value) error {
	for _, v := range options {
		if param == nil {
			break
		}
		param.SetValue(v.Key, v.Value)
	}
	return fb.httpReq(http.MethodPut, method, param, value)
}

func (fb *FBClient) Delete(method string, param Params, value interface{}, options ...Value) error {
	for _, v := range options {
		if param == nil {
			break
		}
		param.SetValue(v.Key, v.Value)
	}
	return fb.httpReq(http.MethodDelete, method, param, value)
}

func (fb *FBClient) Get(path string, params Params, value interface{}, options ...Value) error {
	for _, v := range options {
		if params == nil {
			break
		}
		params.SetValue(v.Key, v.Value)
	}
	return fb.get(fmt.Sprintf("%v?%v", path, params.Encode()), value)
}

func (fb *FBClient) newRequest(httpMethod, url string, reqData []byte) (*http.Request, error) {
	req, err := http.NewRequest(httpMethod, fmt.Sprintf("%s%s", fb.endpoint, url), bytes.NewReader(reqData))
	if err != nil {
		return nil, err
	}
	token, err := fb.tokenProvider.SignJwt(url, reqData)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", fb.tokenProvider.GetApiKey())
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	return req, nil
}

func (fb *FBClient) httpReq(httpMethod, path string, param interface{}, value interface{}) (err error) {
	vi := reflect.ValueOf(value)
	if vi.Kind() != reflect.Ptr {
		return fmt.Errorf("value must be pointer")
	}
	requestData, err := json.Marshal(param)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	if Debug {
		log.Printf("httpReq request: %v  %v \n", path, string(requestData))
	}
	req, err := fb.newRequest(httpMethod, path, requestData)
	if err != nil {
		return err
	}

	resp, err := fb.imp.Do(req)
	if err != nil {
		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}

		return err
	}
	if resp == nil || resp.StatusCode < http.StatusOK || resp.StatusCode > 300 {
		data, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("response err: %v %v %v", resp.StatusCode, resp.Status, string(data))
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if Debug {
		log.Printf("httpReq response: %v %v \n", path, string(data))
	}
	if len(data) != 0 {
		err = json.Unmarshal(data, value)
		if err != nil {
			return fmt.Errorf("%s%s", path, string(data))
		}
	}
	return nil

}

func (fb *FBClient) get(path string, value interface{}) (err error) {
	vi := reflect.ValueOf(value)
	if vi.Kind() != reflect.Ptr {
		return fmt.Errorf("value must be pointer")
	}

	if Debug {
		log.Printf("get request: %v \n", path)
	}

	req, err := fb.newRequest("GET", path, nil)
	if err != nil {
		return err
	}
	resp, err := fb.imp.Do(req)
	if err != nil {
		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}

		return err
	}

	if resp == nil || resp.StatusCode < http.StatusOK || resp.StatusCode > 300 {
		data, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("response err: %v  %v %v", resp.StatusCode, resp.Status, string(data))
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if Debug {
		log.Printf("get response: %v \n %v \n\n", path, string(data))
	}
	if len(data) > 0 {
		err = json.Unmarshal(data, value)
		if err != nil {
			return fmt.Errorf("%s%s", path, string(data))
		}
	}
	return nil
}
