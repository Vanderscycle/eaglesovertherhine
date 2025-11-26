package lib

import (
	"fmt"
)

// TODO: create person class w/ rank, name,  + history (battles/medals/promotions)
type company struct {
	name             string
	nickname string
	privates         int16
	corporals        int16
	corporalFourrier int16
	drummers         int16
	sergeants        int16
	sergeantMajor    int16
	subLieutenant    int16
	lieutenant       int16
	captain          int16
}

func New(name string) company {
	c := company{name: name, nickname: "unearned", privates: 120, corporals: 8, corporalFourrier: 1, drummers: 2, sergeants: 4, sergeantMajor: 1, subLieutenant: 1, lieutenant: 1, captain: 1}
	return c
}

func (c company) Status() {
	fmt.Printf("The company's name: %s", c.name)
}
