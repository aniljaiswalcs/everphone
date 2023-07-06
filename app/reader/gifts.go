package reader

import (
	"encoding/json"
	"io/ioutil"
)

type Gift struct {
	Name       string
	Categories []string
}

type GiftRepository interface {
	GetAllGifts() ([]Gift, error)
}

func LoadGifts(filename string) []Gift {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var gifts []Gift
	err = json.Unmarshal(data, &gifts)
	if err != nil {
		panic(err)
	}

	return gifts
}
