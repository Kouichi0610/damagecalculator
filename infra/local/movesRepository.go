package local

import (
	"damagecalculator/domain/move"

	"bufio"
	"bytes"
	"damagecalculator/assets"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type movesRepository struct {
}

// 英語名から取得できること あるいはasset日本語対応

func (r *movesRepository) Get(name string) (*move.MoveFactory, error) {
	res, ok := movesMap[name]
	if !ok {
		return nil, fmt.Errorf("%s not found.", name)
	}
	return res, nil
}

var movesMap map[string]*move.MoveFactory

func init() {
	movesMap = make(map[string]*move.MoveFactory, 0)
	path := ""
	for _, assetName := range assets.AssetNames() {
		if strings.Contains(assetName, "moves") {
			path = assetName
			break
		}
	}
	res, err := assets.Asset(path)
	if err != nil {
		panic(err)
	}
	rd := bytes.NewReader(res)
	reader := bufio.NewReaderSize(rd, 1024)
	for {
		buf, isPrefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if isPrefix {
			panic("size over")
		}

		mv := &move.MoveFactory{}
		err = json.Unmarshal(buf, mv)

		if err != nil {
			panic(err)
		}
		movesMap[mv.Name] = mv
	}
}
