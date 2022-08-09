package go_fireblocks

const Debug = false

type SigningAlgorithm string

const MPC_ECDSA_SECP256K1 SigningAlgorithm = "MPC_ECDSA_SECP256K1"
const MPC_ECDSA_SECP25MPC_ECDSA_SECP256R16K1 SigningAlgorithm = "MPC_ECDSA_SECP256R1"
const MPC_EDDSA_ED25519 SigningAlgorithm = "MPC_EDDSA_ED25519"

type PaseSort string

const ASC PaseSort = "ASC"
const DESC PaseSort = "DESC"

type TransferStatus string

const SUBMITTED TransferStatus = "SUBMITTED"
const QUEUED TransferStatus = "QUEUED"
const PENDING_AUTHORIZATION TransferStatus = "PENDING_AUTHORIZATION"
const PENDING_SIGNATURE TransferStatus = "PENDING_SIGNATURE"
const BROADCASTING TransferStatus = "BROADCASTING"
const PENDING_3RD_PARTY_MANUAL_APPROVAL TransferStatus = "PENDING_3RD_PARTY_MANUAL_APPROVAL"
const PENDING_3RD_PARTY TransferStatus = "PENDING_3RD_PARTY"
const CONFIRMING TransferStatus = "CONFIRMING"
const PARTIALLY_COMPLETED TransferStatus = "PARTIALLY_COMPLETED"
const PENDING_AML_SCREENING TransferStatus = "PENDING_AML_SCREENING"
const COMPLETED TransferStatus = "COMPLETED"
const CANCELLED TransferStatus = "CANCELLED"
const REJECTED TransferStatus = "REJECTED"
const BLOCKED TransferStatus = "BLOCKED"
const FAILED TransferStatus = "FAILED"
