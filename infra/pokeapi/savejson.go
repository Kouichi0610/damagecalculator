package pokeapi

import (
	"damagecalculator/infra/index"
	"encoding/json"
	"fmt"
	"os"
)

func SaveSpecies() {
	fp, err := os.Create("species.txt")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	rp := new(speciesRepository)
	indices := index.IndexArray()
	for _, name := range indices {
		sp, err := rp.Get(name)
		if err != nil {
			panic(err)
		}
		bin, err := json.Marshal(sp)
		if err != nil {
			panic(err)
		}
		str := fmt.Sprintf("%s,\n", string(bin))
		fp.WriteString(str)

		fmt.Printf("writed:%s\n", name)
	}
}

/*
	fp, err := os.Create("poke.txt")
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	for i := 1; i <= 1000; i++ {
		id := strconv.Itoa(i)
		poke, err := pokeapi.Pokemon(id)
		if err != nil {
			break
		}
		p, err := pokeapi.PokemonSpecies(id)
		if err != nil {
			break
		}
		jname := ""
		for _, n := range p.Names {
			if n.Language.Name == JpName {
				jname = n.Name
				break
			}
		}

		wstr := fmt.Sprintf("{\"%s\", \"%s\"},\n", poke.Name, jname)
		fp.WriteString(wstr)
	}
*/

/*
type hoge struct {
	Name  string
	Types []types.Type
}

func Test_Json(t *testing.T) {
	// true　を[]byteのjson形式で
	bolB, _ := json.Marshal(true)
	t.Errorf("%s", string(bolB))

	h := &hoge{
		"フライゴン",
		[]types.Type{types.Dragon, types.Ground},
	}
	hB, _ := json.Marshal(h)
	t.Errorf("%s", string(hB))

	// jsonから構造体
	h2 := new(hoge)
	json.Unmarshal(hB, h2)

	t.Errorf("%s %s", h2.Name, types.NewTypes(h2.Types...).String())
}

*/
