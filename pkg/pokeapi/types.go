package pokeapi

type Pokemon struct {
	Id        int              `json:"id"`
	Name      string           `json:"name"`
	Types     []PokemonType    `json:"types"`
	Stats     []PokemonStat    `json:"stats"`
	Abilities []PokemonAbility `json:"abilities"`
	Height    float32          `json:"height"`
	Weight    float32          `json:"weight"`
}

type PokemonType struct {
	Slot int      `json:"slot"`
	Type TypeInfo `json:"type"`
}

type TypeInfo struct {
	Name string `json:"name"`
}

type PokemonStat struct {
	BaseStat int      `json:"base_stat"`
	Effort   int      `json:"effort"`
	Stat     StatInfo `json:"stat"`
}

type StatInfo struct {
	Name string `json:"name"`
}

type PokemonAbility struct {
	Ability AbilityInfo `json:"ability"`
	Hidden  bool        `json:"is_hidden"`
}

type AbilityInfo struct {
	Name string `json:"name"`
}
