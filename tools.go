package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/mtslzr/pokeapi-go"
)

const JpName = "ja-Hrkt"

func makePokeList() {
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
}

func makeFormChanges() {
	withoutList := []string{
		"-mega",
		"pikachu-",
		"-totem",
		"castform-", // ポワルン(とくせいで対応するため無視)
	}
	fp, err := os.Create("forms.txt")
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	for i := 1; i <= 1000; i++ {
		id := strconv.Itoa(i)
		p, err := pokeapi.PokemonSpecies(id)
		if err != nil {
			break
		}
		varieties := make([]string, 0)
		for _, v := range p.Varieties {
			name := v.Pokemon.Name
			if v.IsDefault {
				//varieties = append(varieties, name)
				continue
			}
			isWithout := false
			for _, w := range withoutList {
				if strings.Contains(name, w) {
					isWithout = true
					break
				}
			}
			if !isWithout {
				varieties = append(varieties, name)
			}
		}
		if len(varieties) >= 1 {
			jname := ""
			for _, n := range p.Names {
				if n.Language.Name == JpName {
					jname = n.Name
					break
				}
			}
			for _, vname := range varieties {
				wstr := fmt.Sprintf("{\"%s\", \"%s\"},\n", vname, jname)
				fp.WriteString(wstr)

			}
		}
	}
}

func appendNames() {
	fp, err := os.Open("add.txt")
	if err != nil {
		panic(err)
	}
	wfp, err := os.Create("appends.txt")
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	defer wfp.Close()

	reader := bufio.NewReaderSize(fp, 256)
	for {
		line, isPrefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if isPrefix {
			fmt.Println("読み切れなかった")
		}
		str := string(line)
		l := strings.Split(str, ",")
		fmt.Printf("%v\n", l)

		wstr := fmt.Sprintf("{%s, \"%s\", \"%s\"},\n", l[0], l[2], l[1])
		wfp.WriteString(wstr)
	}

}

func createNames() {
	fp, err := os.Create("names.txt")
	defer fp.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 1; i <= 1000; i++ {
		id := strconv.Itoa(i)
		poke, err := pokeapi.Pokemon(id)
		if err != nil {
			break
		}
		sp, _ := pokeapi.PokemonSpecies(id)

		jpname := ""
		for _, s := range sp.Names {
			if s.Language.Name == JpName {
				jpname = s.Name
				break
			}
		}

		str := fmt.Sprintf("{\"%s\", \"%s\"},\n", poke.Name, jpname)

		fp.WriteString(str)
	}

}
