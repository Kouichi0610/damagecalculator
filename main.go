package main

import (
	//"damagecalculator/assets"
	"damagecalculator/infra/pokeapi"
	"fmt"

	"github.com/jteeuwen/go-bindata"
)

func main() {
	fmt.Println("Hello")

	saveAssets()

	/*
		names := assets.AssetNames()
		for _, name := range names {
			fmt.Printf("%s\n", name)
		}

		res, err := assets.Asset("assets/moves.txt")
		if err != nil {
			fmt.Printf("%s\n", err.Error())
		}
		fmt.Printf("%d\n", len(res))
	*/
}

// TODO:アセット作成のための
func saveSpecies() {
	pokeapi.SaveSpecies()

}

func saveAssets() {
	c := bindata.NewConfig()
	// バイナリデータ
	c.Input = []bindata.InputConfig{
		{Path: "./data/", Recursive: true},
	}
	c.Package = "assets"
	c.Recursive = true
	c.Output = "./assets/assets.go"

	//c.Recursive
	err := bindata.Translate(c)

	if err != nil {
		panic(err)
	}
}

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
