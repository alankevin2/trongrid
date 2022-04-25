package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"reflect"
	"time"

	"github.com/spf13/viper"
)

func New(network Network) *TronGridV1 {
	return createInstance(network)
}

func (t *TronGridV1) GetTransactionsByAddress(address string, request GetTransactionsByAddressRequest) (result *TrongridTransactionsResp) {
	trc20 := ""
	if request.TRC20 {
		trc20 = "trc20"
	}
	tokenAddress := t.tokens[string(request.Symbol)]
	query := fmt.Sprintf("only_confirmed=true&only_to=true&limit=%s&order_by=%s&min_timestamp=%s&max_timestamp=%s&contract_address=%s",
		request.Limit, request.OrderBy, request.MinTimestamp, request.MaxTimestamp, tokenAddress)
	url := fmt.Sprintf("%s/v1/accounts/%s/transactions/%s?%s", t.baseURL, address, trc20, query)
	fmt.Println(url)
	var body bytes.Buffer
	req, err := http.NewRequest("GET", url, &body)
	if err != nil {
		return
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("TRON-PRO-API-KEY", t.apiKey) // testnet does not require one

	client := &http.Client{Timeout: 15 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	fmt.Println(string(resBody))
	err = json.Unmarshal(resBody, &result)
	if err != nil {
		return
	}
	return result
}

/** private methods **/

func createInstance(network Network) *TronGridV1 {
	var fileName string
	switch network {
	case Network_Mainnet:
		fileName = "provider-mainnet.yml"
	case Network_Shasta:
		fileName = "provider-testnet-shasta.yml"
	}

	currentPath, _ := os.Getwd()
	fullpath := path.Join(currentPath, "config", fileName)
	_, err := os.Stat(fullpath)
	// 如果找不到，代表當前執行環境不是以此pkg為主，而是被別人vendor引用
	if err != nil {
		pkgPath := reflect.TypeOf(TronGridV1{}).PkgPath() + "../../../config/"
		fullpath = path.Join(currentPath, "vendor", pkgPath, fileName)
	}
	viper.SetConfigFile(fullpath)
	viper.SetConfigType("yml")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("fatal error config file: %w", err))
	}
	return &TronGridV1{
		network: network,
		apiKey:  viper.GetString("root.api-key"),
		baseURL: viper.GetString("root.url"),
		tokens:  viper.GetStringMapString("root.tokens"),
	}
}
