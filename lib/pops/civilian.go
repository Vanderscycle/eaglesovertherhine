package pops

import (
	"encoding/json"
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
)

type CivilianPop struct {
	FirstName string
	LastName  string
	Nation    Nation
	Role      CivilianRole
	// sex       Sex
	// age       int16
}

func NewCivilian( nation Nation, role CivilianRole) CivilianPop {
	c := CivilianPop{FirstName: gofakeit.FirstName(), LastName: gofakeit.LastName(), Nation: nation, Role: role}
	return c
}

func (c CivilianPop) Status() {
    b, _ := json.Marshal(c)
    //b, _ := json.MarshalIndent(c, "", "  ")
    fmt.Println(string(b))
}
