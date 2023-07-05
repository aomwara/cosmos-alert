package rpc

import (
	"alertbot/types"
	"alertbot/validator"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetSignature(height int) (string, error) {

	var endpoint = RPC + "/block?height=" + strconv.Itoa(height)

	response, err := http.Get(endpoint)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP request failed with status: %s", response.Status)
	}

	var block types.BlockResponse
	err = json.Unmarshal(body, &block)
	if err != nil {
		return "", err
	}

	for _, signature := range block.Result.Block.Last_Commit.Signatures {
		if signature.ValidatorAddress == validator.GetValidator() {
			return signature.Signature, nil
		}

	}

	return "", nil
}
