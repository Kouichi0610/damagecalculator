package manual

import (
	"bufio"
	"damagecalculator/domain/species"
	"damagecalculator/infra/index"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type speciesRepository struct {
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
	fp, err := os.Open("species.txt")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	reader := bufio.NewReaderSize(fp, 4096)
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