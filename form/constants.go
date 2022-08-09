package form

type PeerTypes string

const VaultAccount PeerTypes = "VAULT_ACCOUNT"
const OneTimeAddressType PeerTypes = "ONE_TIME_ADDRESS"
const ExchangeAccount PeerTypes = "EXCHANGE_ACCOUNT"
const InternalWallet PeerTypes = "INTERNAL_WALLET"
const ExternalWallet PeerTypes = "EXTERNAL_WALLET"
const UnknownPeer PeerTypes = "UNKNOWN_PEER"
const FiatAccount PeerTypes = "FIAT_ACCOUNT"
const NetworkConnection PeerTypes = "NETWORK_CONNECTION"
const COMPOUND PeerTypes = "COMPOUND"

type TransactionTypes string

const TransactionTransfer TransactionTypes = "TRANSFER"
const TransactionMint TransactionTypes = "MINT"
const TransactionBurn TransactionTypes = "BURN"
const TransactionSupplyToCompound TransactionTypes = "SUPPLY_TO_COMPOUND"
const TransactionRedeemFromCompound TransactionTypes = "REDEEM_FROM_COMPOUND"
const RAW TransactionTypes = "RAW"
const ContractCall TransactionTypes = "CONTRACT_CALL"
const OneTimeAddressTransfer TransactionTypes = "ONE_TIME_ADDRESS"
