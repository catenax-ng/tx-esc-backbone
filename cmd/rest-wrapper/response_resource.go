package main

// ResponseResource contains information about the issued transaction.
//
// swagger:model ResponseResource
type ResponseResource struct {
	// Height of the block containing the transaction
	Height int64 `json:"height,omitempty"`
	// TxHash is the hash of the transaction
	TxHash string `json:"txHash,omitempty"`
	// RawLog contains the raw log of the transaction.
	RawLog string `json:"rawLog,omitempty"`
}
