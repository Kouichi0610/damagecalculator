package main

import (
	"damagecalculator/infra/pokeapi"
	"io/ioutil"
	"os"

	"github.com/jteeuwen/go-bindata"
)

/*
go-bindata
バイナリデータをソースコードに埋め込むためのツール。
ファイル経由でアクセスするとパスの都合でテストが通らないなどの問題があって採用。

インストール
> go get github.com/jteeuwen/go-bindata/...

使い方
> go-bindata -pkg=server -o=server/assets.go ./assets/...

-pkg パッケージ名を指定
-o リソース出力先
最後の引数 対象ディレクトリ(/...で再帰的に処理)

(./assets/... のファイルを　server/assets.goとして生成)
*/
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
