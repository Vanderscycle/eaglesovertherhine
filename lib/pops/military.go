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
}

func NewSoldier() MilitaryPop{
m := MilitaryPop{firstName: gofakeit.FirstName(),lastName: gofakeit.LastName(),rank: Soldat,nation: France}
	return m
}

func (m MilitaryPop) Status(){
fmt.Printf("the %s name is: %s %s\n", m.rank,m.firstName,m.lastName)
}
