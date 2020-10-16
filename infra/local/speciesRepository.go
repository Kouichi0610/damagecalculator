package local

import (
	"bufio"
	"bytes"
	"damagecalculator/assets"
	"damagecalculator/domain/species"
	"damagecalculator/infra/index"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type speciesRepository struct {
}

func (s *speciesRepository) GetAll() []species.Species {
	res := make([]species.Species, 0)
	names := index.IndexArray()
	for _, name := range names {
		sp, ok := speciesMap[name]
		if !ok {
			continue
		}
		res = append(res, *sp)
	}
	return res
}
func (s *speciesRepository) Get(name string) (*species.Species, error) {
	jname, err := index.Index(name)
	if err != nil {
		return nil, err
	}
	res, ok := speciesMap[jname]
	if !ok {
		return nil, fmt.Errorf("%s not found.", name)
	}
	return res, nil
}

var speciesMap map[string]*species.Species

func init() {
	speciesMap = make(map[string]*species.Species, 0)
	path := ""
	for _, assetName := range assets.AssetNames() {
		if strings.Contains(assetName, "species") {
			path = assetName
			break
		}
	}
	res, err := assets.Asset(path)
	if err != nil {
		panic(err)
	}
	rd := bytes.NewReader(res)
	reader := bufio.NewReaderSize(rd, 4096)
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

		sp := &species.Species{}
		err = json.Unmarshal(buf, sp)

		ename := sp.Name
		jname, err := index.JName(sp.Name)
		sp.Name = jname
		if err != nil {
			panic(err)
		}
		speciesMap[ename] = sp
	}

}
