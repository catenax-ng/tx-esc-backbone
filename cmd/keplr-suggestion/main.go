// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/catenax/esc-backbone/app"
)

type Bip44 struct {
	CoinType uint `json:"coinType"`
}

type Bech32Config struct {
	Bech32PrefixAccAddr  string `json:"bech32PrefixAccAddr"`
	Bech32PrefixAccPub   string `json:"bech32PrefixAccPub"`
	Bech32PrefixValAddr  string `json:"bech32PrefixValAddr"`
	Bech32PrefixValPub   string `json:"bech32PrefixValPub"`
	Bech32PrefixConsAddr string `json:"bech32PrefixConsAddr"`
	Bech32PrefixConsPub  string `json:"bech32PrefixConsPub"`
}

type Currency struct {
	// Name displayed in keplr
	CoinDenom string `json:"coinDenom"`
	// The actual currency name.
	CoinMinimalDenom string `json:"coinMinimalDenom"`
	CoinDecimals     uint8  `json:"coinDecimals"`
}

type GasPriceStep struct {
	Low     float64 `json:"low"`
	Average float64 `json:"average"`
	High    float64 `json:"high"`
}
type KeplrSuggestion struct {
	//See https://github.com/chainapsis/keplr-example/commit/6ad5e36dcdc73669d1fd40b2a74760e1694ad85c#
	ChainId       string       `json:"chainId"`
	ChainName     string       `json:"chainName"`
	Rpc           string       `json:"rpc"`
	Rest          string       `json:"rest"`
	StakeCurrency Currency     `json:"stakeCurrency"`
	Bip44         Bip44        `json:"bip44"`
	Bech32Config  Bech32Config `json:"bech32Config"`
	Currencies    []Currency   `json:"currencies"`
	FeeCurrencies []Currency   `json:"feeCurrencies"`
	CoinType      uint         `json:"coinType"`
	GasPriceStep  GasPriceStep `json:"gasPriceStep"`
	Features      []string     `json:"features"`
}

func toPrefix(power uint8) string {
	switch power {
	case 1:
		return "d"
	case 2:
		return "c"
	case 3:
		return "m"
	case 4:
		fallthrough
	case 5:
		fallthrough
	case 6:
		return "u" // micro
	case 7:
		fallthrough
	case 8:
		fallthrough
	case 9:
		return "n"
	case 10:
		fallthrough
	case 11:
		fallthrough
	case 12:
		return "p"
	case 13:
		fallthrough
	case 14:
		fallthrough
	case 15:
		return "f"
	case 16:
		fallthrough
	case 17:
		fallthrough
	case 18:
		return "a"
	default:
		panic("Powers higher than 18 not supported.")
	}
}
func newBechConfig(addrPrefix string) Bech32Config {

	return Bech32Config{
		Bech32PrefixAccAddr:  addrPrefix,
		Bech32PrefixAccPub:   addrPrefix + "pub",
		Bech32PrefixValAddr:  addrPrefix + "valoper",
		Bech32PrefixValPub:   addrPrefix + "valoperpub",
		Bech32PrefixConsAddr: addrPrefix + "valcons",
		Bech32PrefixConsPub:  addrPrefix + "valconspub",
	}
}
func newCurrency(minimalDenom string, decimals uint8) Currency {
	uppercaseDenom := strings.ToUpper(minimalDenom)
	minimalDenom = toPrefix(decimals) + strings.ToLower(minimalDenom)
	return Currency{
		CoinDenom:        uppercaseDenom,
		CoinMinimalDenom: minimalDenom,
		CoinDecimals:     decimals,
	}
}

func main() {

	chainId := "catenax-testnet-1"
	chainName := "Catena-X Testnet"
	bech32prefix := app.AccountAddressPrefix
	bip44 := uint(118)
	rest := "http://0.0.0.0:1317/"
	rpc := "http://0.0.0.0:26657/"
	caxCurrency := newCurrency("caxdemo", 9)
	voucherCurrency := newCurrency("voucher", 2)
	suggestion := KeplrSuggestion{
		ChainId:       chainId,
		ChainName:     chainName,
		Rpc:           rpc,
		Rest:          rest,
		StakeCurrency: caxCurrency,
		Bip44:         Bip44{bip44},
		Bech32Config:  newBechConfig(bech32prefix),
		Currencies:    []Currency{caxCurrency, voucherCurrency},
		CoinType:      bip44,
		FeeCurrencies: []Currency{caxCurrency},
		GasPriceStep: GasPriceStep{
			Low:     0.01,
			Average: 0.025,
			High:    0.04,
		},
		Features: []string{"stargate"},
	}
	bytes, err := json.Marshal(suggestion)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = os.WriteFile("./web/public/chain/"+chainId+"-suggestion.json", bytes, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
