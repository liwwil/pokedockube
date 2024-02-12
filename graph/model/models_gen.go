// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Mutation struct {
}

type Pokemon struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Description *string          `json:"description,omitempty"`
	TypeID      string           `json:"typeID"`
	CategoryID  string           `json:"categoryID"`
	AbilityID   string           `json:"abilityID"`
	Type        *PokemonType     `json:"type"`
	Category    *PokemonCategory `json:"category"`
	Ability     *PokemonAbility  `json:"ability"`
}

type PokemonAbility struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PokemonCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PokemonType struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Query struct {
}

type CreatePokemon struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	TypeID      string  `json:"typeID"`
	CategoryID  string  `json:"categoryID"`
	AbilityID   string  `json:"abilityID"`
}

type NewAbility struct {
	Name string `json:"name"`
}

type NewCategory struct {
	Name string `json:"name"`
}

type NewType struct {
	Name string `json:"name"`
}

type UpdatePokemon struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	TypeID      string  `json:"typeID"`
	CategoryID  string  `json:"categoryID"`
	AbilityID   string  `json:"abilityID"`
}
