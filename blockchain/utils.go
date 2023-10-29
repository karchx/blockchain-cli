package blockchain

import (
	"bytes"
	"encoding/binary"
	"log"
)

const (
	DB_PATH      string = "./tmp/blocks"
	DB_FILE      string = "./tmp/blocks/MANIFES"
	GENESIS_DATA string = "TXN #1"
)

func toHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

func handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
