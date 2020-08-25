package main

import (
	"damagecalculator/assets"
	"fmt"
)

func main() {
	fmt.Println("Hello")
	//pokeapi.SaveSpecies()

	names := assets.AssetNames()
	for _, name := range names {
		fmt.Printf("%s\n", name)
	}

	res, err := assets.Asset("assets/moves.txt")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	fmt.Printf("%d\n", len(res))
}
