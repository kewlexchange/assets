package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Token struct {
	ChainId  int    `json:"chainId"`
	Address  string `json:"address"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
	LogoURI  string `json:"logoURI"`
}

type SwapDefault struct {
	Name      string   `json:"name"`
	Timestamp string   `json:"timestamp"`
	Version   Version  `json:"version"`
	Tags      Tags     `json:"tags"`
	LogoURI   string   `json:"logoURI"`
	Keywords  []string `json:"keywords"`
	Tokens    []Token  `json:"tokens"`
}

type Version struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
	Patch int `json:"patch"`
}

type Tags struct{}

func main() {
	// Dosyadan JSON verisini oku

	filePath := "./avax/index.json"
	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Dosya okuma hatası:", err)
		os.Exit(1)
	}

	// JSON verisini SwapDefault struct tipine çöz
	var swapDefault SwapDefault
	err = json.Unmarshal(jsonData, &swapDefault)
	if err != nil {
		fmt.Println("JSON çözme hatası:", err)
		os.Exit(1)
	}

	// Logo URI'leri güncelle
	swapDefault.LogoURI = ""

	for i := range swapDefault.Tokens {

		var ContractAddress = strings.ToLower(swapDefault.Tokens[i].Address)
		var NewLogoURL = fmt.Sprintf("https://raw.githubusercontent.com/kewlexchange/assets/main/avax/tokens/%s/logo.svg", ContractAddress)
		swapDefault.Tokens[i].LogoURI = NewLogoURL
	}

	// Güncellenmiş JSON verisini bastır
	updatedJSON, err := json.MarshalIndent(swapDefault, "", "    ")
	if err != nil {
		fmt.Println("JSON formatlama hatası:", err)
		os.Exit(1)
	}

	// Güncellenmiş JSON verisini dosyaya yaz
	err = ioutil.WriteFile(filePath, updatedJSON, 0644)
	if err != nil {
		fmt.Println("Dosya yazma hatası:", err)
		os.Exit(1)
	}

	fmt.Println("JSON dosyası güncellendi ve yeni veri dosyaya yazıldı.")
}
