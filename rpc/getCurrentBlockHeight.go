package rpc

import (
	"alertbot/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetCurrentBlockHeight() (string, error) {

	var endpoint = RPC + "/abci_info"

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

	var abciInfo types.ABCIInfoResponse
	err = json.Unmarshal(body, &abciInfo)
	if err != nil {
		return "", err
	}

	lastBlockHeight := abciInfo.Result.Response.LastBlockHeight
	return lastBlockHeight, nil

}
