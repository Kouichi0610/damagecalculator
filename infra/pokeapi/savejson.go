package pokeapi

import (
	"damagecalculator/infra/index"
	"encoding/json"
	"fmt"
	"os"
)

// TODO:LoadSpecies()

func SaveSpecies() {
	fp, err := os.Create("species.txt")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	rp := new(speciesRepository)
	indices := index.IndexArray()
	cnt := len(indices)
	i := 0
	for _, name := range indices {
		sp, err := rp.Get(name)
		if err != nil {
			i++
			fmt.Printf("failed %s %f\n", name, float64(float64(i)/float64(cnt)))
			continue
		}
		bin, err := json.Marshal(sp)
		if err != nil {
			panic(err)
		}
		str := fmt.Sprintf("%s\n", string(bin))
		fp.WriteString(str)

		i++
		fmt.Printf("%s %f\n", name, float64(float64(i)/float64(cnt)))

	}
}
