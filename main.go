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

func main() {
	s := server.NewServer(Repositories())
	_, err := s.Serve()
	if err != nil {
		log.Fatal("server run failed ", err)
	}
}
