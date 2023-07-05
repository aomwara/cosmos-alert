package types

type ABCIInfoResponse struct {
	JsonRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Response struct {
			Data             string `json:"data"`
			Version          string `json:"version"`
			LastBlockHeight  string `json:"last_block_height"`
			LastBlockAppHash string `json:"last_block_app_hash"`
		}
	}
}
