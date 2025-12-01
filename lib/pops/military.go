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
	role Role
}

func NewSoldier(rank Rank, nation Nation, role Role) MilitaryPop {
	m := MilitaryPop{firstName: gofakeit.FirstName(), lastName: gofakeit.LastName(), rank: rank, nation: nation, role: role }
	return m
}

func (m MilitaryPop) Status() {
	fmt.Printf("%+v\n", m)
}
