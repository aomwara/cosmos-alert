package rpc

var RPC = "http://localhost:26657"

func SetRPC(rpc string) {
	RPC = rpc
}

func GetRPC() string {
	return RPC
}
