package model

type Creature struct {
	CreatureID   int    `json:"creature_id"`
	CreatureName string `json:"creature_name"`
	NickName     string `json:"nickname"`
	Species      string `json:"species"`
	ImageUrl     string `json:"imageurl"`
	Abilities    string `json:"abilities"`
}
