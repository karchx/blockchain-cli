package wallet

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"encoding/json"
	"log"
	"os"
)

const WALLET_FILE = "./tmp/wallets.data"

type P struct {
	Name string
}

func (ws *Wallets) SaveFile() {
	data, err := json.Marshal(ws)
	if err != nil {
		log.Printf("Error json Marshal: %s", err.Error())
	}

	err = os.WriteFile(WALLET_FILE, data, 0644)
	if err != nil {
		log.Panicln(err)
	}
}

func (ws *Wallets) loadFile() error {
	if _, err := os.Stat("./tmp"); os.IsNotExist(err) {
		os.Mkdir("tmp", os.ModeDevice)
	}

	if _, err := os.Stat(WALLET_FILE); os.IsNotExist(err) {
		return err
	}

	var wallets Wallets
	fileContent, err := os.ReadFile(WALLET_FILE)
	if err != nil {
		return err
	}

	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&wallets)
	if err != nil {
		return err
	}

	ws.Wallets = wallets.Wallets
	return nil
}

func CreateWallet() (*Wallets, error) {
	wallets := Wallets{
		Wallets: make(map[string]*Wallet),
	}

	err := wallets.loadFile()
	return &wallets, err
}

func (ws *Wallets) AddWallet(alias string) string {
	wallet := makeWallet(alias)
	address := string(wallet.address())

	ws.Wallets[address] = wallet
	return address
}

func (ws *Wallets) GetAllAddresses() []string {
	var addresses []string

	for address := range ws.Wallets {
		addresses = append(addresses, address)
	}

	return addresses
}
