package main

import (
	"damagecalculator/domain/ability"
	"damagecalculator/domain/item"
	"damagecalculator/domain/move"
	"damagecalculator/domain/pokenames"
	"damagecalculator/domain/species"
	"damagecalculator/infra/local"
	"damagecalculator/presentation/server"
	"log"
)

func Repositories() (pokenames.Repository, species.Repository, move.Repository, ability.Repository, item.Repository) {
	return local.Names(), local.Species(), local.Move(), local.Ability(), local.Item()
}

// 古いビルドのキャッシュが残っていた？
// \node_modules\@vue\cli-plugin-router\generator\template\src\views
func main() {
	s := server.NewServer(Repositories())
	err := s.Serve()
	if err != nil {
		log.Fatal("server run failed ", err)
	}
}

/*
	attacker
	nature
	[6]basePoints
	rank
	ability
	item

	move

	defender
	nature
	[6]basePoints
	rank
	ability
	item

	options
		やけど
		てだすけ
		リフレクター
		ひかりのかべ

	天候
	フィールド

	-> 全部一度に渡すのではなく
	一つ渡して計算できるなら結果を取得

	Attacker(Defender)の状態を保存
*/
