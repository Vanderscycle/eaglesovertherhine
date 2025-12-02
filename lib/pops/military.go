package pops

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
)

type MilitaryPop struct {
	firstName string
	lastName  string
	rank      Rank
	nation    Nation
	role      Role
	// sex       Sex
	// age       int16
}

func NewSoldier(rank Rank, nation Nation, role Role) MilitaryPop {
	m := MilitaryPop{firstName: gofakeit.FirstName(), lastName: gofakeit.LastName(), rank: rank, nation: nation, role: role}
	return m
}

// TODO: eventually militaryPop will be derived from a general pop
// func (m MilitaryPop) Age(min int, max int) int16 {
// 	return int16(min + rand.Intn(max-min))
// }

func (m MilitaryPop) Status() {
	fmt.Printf("%+v\n", m)
}
