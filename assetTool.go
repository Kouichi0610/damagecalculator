package main

import (
	"damagecalculator/infra/pokeapi"
	"io/ioutil"
	"os"

	"github.com/jteeuwen/go-bindata"
)

const tmpDir = "dcalc"
const tmpSpecies = "species.txt"
const tmpMoves = "moves.txt"

func makeAssets() {

	dir, err := ioutil.TempDir("", tmpDir)
	defer os.RemoveAll(dir)
	if err != nil {
		panic(err)
	}
	file, err := ioutil.TempFile(dir, tmpSpecies)
	if err != nil {
		panic(err)
	}
	pokeapi.SaveSpecies(file)

	file, err = ioutil.TempFile(dir, tmpMoves)
	if err != nil {
		panic(err)
	}
	pokeapi.SaveMoves(file)

	saveAssets(dir)
}

func saveAssets(dir string) {
	c := bindata.NewConfig()
	c.Input = []bindata.InputConfig{
		//		{Path: "./data/", Recursive: true},
		{Path: dir, Recursive: true},
	}
	c.Package = "assets"
	c.Recursive = true
	c.Output = "./assets/assets.go"
	err := bindata.Translate(c)
	if err != nil {
		panic(err)
	}
}
