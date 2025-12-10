package pops

import (
	"encoding/json"
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
)

type MilitaryPop struct {
	FirstName string
	LastName  string
	Rank      Rank
	Nation    Nation
	Role      MilitaryRole
	// sex       Sex
	// age       int16
}

func NewSoldier(rank Rank, nation Nation, role MilitaryRole) MilitaryPop {
	m := MilitaryPop{FirstName: gofakeit.FirstName(), LastName: gofakeit.LastName(), Rank: rank, Nation: nation, Role: role}
	return m
}

// TODO: eventually militaryPop will be derived from a general pop
// func (m MilitaryPop) Age(min int, max int) int16 {
// 	return int16(min + rand.Intn(max-min))
// }

func (m MilitaryPop) Status() {
	b, _ := json.Marshal(m)
	fmt.Println(string(b))
}
