package graph

import (
	"strconv"

	"github.com/TanyaEIEI/pokedex/graph/model"
)

type Resolver struct{
}

func getCurrentId(list interface{})(int, error) {
	switch list.(type) {
	case []*model.PokemonAbility :
		newItem := list.([]*model.PokemonAbility)
		return strconv.Atoi(newItem[len(newItem)-1].ID)
	case []*model.PokemonType :
		newItem := list.([]*model.PokemonType)
		return strconv.Atoi(newItem[len(newItem)-1].ID)
	case []*model.PokemonCategory :
		newItem := list.([]*model.PokemonCategory)
		return strconv.Atoi(newItem[len(newItem)-1].ID)
	case []*model.Pokemon :
		newItem := list.([]*model.Pokemon)
		return strconv.Atoi(newItem[len(newItem)-1].ID)
	default:
		panic("Unsupported type")
	}
}