package pokeapi

import (
	"damagecalculator/infra/index"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func SaveMoves(fp *os.File) {
	defer fp.Close()
	rp := new(movesRepository)

	for i := 1; i <= 2000; i++ {
		id := strconv.Itoa(i)
		move, err := rp.Get(id)
		if err != nil {
			break
		}
		bin, err := json.Marshal(move)
		if err != nil {
			panic(err)
		}
		str := fmt.Sprintf("%s\n", string(bin))
		fp.WriteString(str)
		fmt.Printf("%s\n", move.Name)
	}
}

func SaveSpecies(fp *os.File) {
	defer fp.Close()

	rp := new(speciesRepository)
	mr := new(movesRepository)
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

		for i := 0; i < len(sp.Moves); i++ {
			m, _ := mr.Get(sp.Moves[i])
			jname := m.Name
			sp.Moves[i] = jname
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
