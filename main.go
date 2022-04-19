package main

import (
	"fmt"

	"gitlab.inlive7.com/crypto/trongridv1/pkg/api"
)

func main() {
	c := api.New(api.Network_Shasta)
	fmt.Println(c.GetTransactionsByAddress("TK7q7c6RRSjTvuzmVmZNgq18nQrmx1UZtc", api.GetTransactionsByAddressRequest{
		Limit: "2",
		TRC20: true,
	}))
}
