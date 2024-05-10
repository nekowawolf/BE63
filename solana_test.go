// solana_test.go

package be63

import (
	"fmt"
	"testing"
)

func TestInsertsolana(t *testing.T) {
	name := "Solana Mobile"
	network := "Helium Network 5G"
	storage := "512GB"
	ram := "12GB"
	processor := "Qualcomm Snapdragon 8+ Gen1"
	camera := 50
	display := 1080
	battery := 4110
	os := "Android 13"
	price := 1000.00

	insertedID := Insertsolana(name, network, storage, ram, processor, camera, display, battery, os, price)
	fmt.Println(insertedID)
}
