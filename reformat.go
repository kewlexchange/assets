package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func convertToLower(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if info.IsDir() {
		// Klasör adını küçük harfe çevir
		newPath := strings.ToLower(path)

		// Klasör adını değiştir
		err := os.Rename(path, newPath)
		if err != nil {
			fmt.Println(err)
		}
	}

	return nil
}

func main() {
	// Klasörünüzün yolunu belirtin
	rootFolder := "/Users/ersanyakit/Documents/GitHub/nfts"

	// Walk fonksiyonunu kullanarak tüm alt klasörleri dolaşın
	err := filepath.Walk(rootFolder, convertToLower)
	if err != nil {
		fmt.Println(err)
	}
}
