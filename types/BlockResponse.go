package types

type BlockResponse struct {
	JsonRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		BlockID struct {
			Hash  string `json:"hash"`
			Parts struct {
				Total int    `json:"total"`
				Hash  string `json:"hash"`
			}
		}
		Block struct {
			Header struct {
				Version struct {
					Block string `json:"block"`
				}
				ChainID     string `json:"chain_id"`
				Height      string `json:"height"`
				Time        string `json:"time"`
				LastBlockID struct {
					Hash  string `json:"hash"`
					Parts struct {
						Total int    `json:"total"`
						Hash  string `json:"hash"`
					}
				}
				LastCommitHash     string `json:"last_commit_hash"`
				DataHash           string `json:"data_hash"`
				ValidatorsHash     string `json:"validators_hash"`
				NextValidatorsHash string `json:"next_validators_hash"`
				ConsensusHash      string `json:"consensus_hash"`
				AppHash            string `json:"app_hash"`
				LastResultsHash    string `json:"last_results_hash"`
				EvidenceHash       string `json:"evidence_hash"`
				ProposerAddress    string `json:"proposer_address"`
			}
			Data struct {
				Txs []string `json:"txs"`
			}
			Evidence struct {
				Evidence []string `json:"evidence"`
			}
			Last_Commit struct {
				Height  string `json:"height"`
				Round   int    `json:"round"`
				BlockID struct {
					Hash  string `json:"hash"`
					Parts struct {
						Total int    `json:"total"`
						Hash  string `json:"hash"`
					}
				}
				Signatures []struct {
					BlockIDFlag      int    `json:"block_id_flag"`
					ValidatorAddress string `json:"validator_address"`
					Timestamp        string `json:"timestamp"`
					Signature        string `json:"signature"`
				}
			}
		}
	}
}
