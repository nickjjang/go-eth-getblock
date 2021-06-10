package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func _hex2int(hexStr string) (int64, error) {
	// remove 0x suffix if found in the input string
	cleaned := strings.Replace(hexStr, "0x", "", -1)

	// base 16 for hexadecimal
	return strconv.ParseInt(cleaned, 16, 64)
}

func _int2hex(value int64) string {
	ret := strings.ToLower(fmt.Sprintf("0x%x", value))
	return ret
}

func EthBlockNumber() int64 {
	param := make(map[string]interface{})
	param["id"] = 1
	param["method"] = "eth_blockNumber"
	param["jsonrpc"] = "2.0"
	body, _ := json.Marshal(param)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://cloudflare-eth.com", bytes.NewReader(body))
	req.Header.Add("Content-Type", "application/json")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return -1
	}
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return -1
	}
	value, err := _hex2int(result["result"].(string))
	if err != nil {
		return -1
	}
	fmt.Println(value)
	return value
}

func EthGetBlockByNumber(number int64) interface{} {
	param := make(map[string]interface{})
	param["id"] = 1
	param["method"] = "eth_getBlockByNumber"
	param["jsonrpc"] = "2.0"
	param["params"] = []interface{}{_int2hex(number), true}
	body, _ := json.Marshal(param)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://cloudflare-eth.com", bytes.NewReader(body))
	req.Header.Add("Content-Type", "application/json")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil
	}
	result = result.(map[string]interface{})["result"]
	return result
}
