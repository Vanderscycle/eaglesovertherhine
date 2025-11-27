package units

import (
	"fmt"

	"github.com/vanderscycle/eaglesovertherhine/lib/pops"
)

// TODO: create person class w/ rank, name,  + history (battles/medals/promotions)
type companyTemplate struct {
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

var FrenchCompanyDefault = companyTemplate{
	privates:         120,
	corporals:        8,
	corporalFourrier: 1,
	drummers:         2,
	sergeants:        4,
	sergeantMajor:    1,
	subLieutenant:    1,
	lieutenant:       1,
	captain:          2,
}

type company struct {
	name             string
	nickname         string
	privates         []pops.MilitaryPop
	corporals        int16
	corporalFourrier int16
	drummers         int16
	sergeants        int16
	sergeantMajor    int16
	subLieutenant    int16
	lieutenant       int16
	captain          int16
	template         companyTemplate
}

func NewCompany(name string) company {
	c := company{name: name, nickname: "unearned", corporals: 8, corporalFourrier: 1, drummers: 2, sergeants: 4, sergeantMajor: 1, subLieutenant: 1, lieutenant: 1, captain: 1, template: FrenchCompanyDefault}

	for range 120 {
		s := pops.NewSoldier()
		c.privates = append(c.privates, s)
	}
	return c

}

func (c company) Status() {
	fmt.Printf("The company's name: %s\nTemplate: %+v\n", c.name, c.template)
}

func (c company) ListPrivates() {
	for _, s := range c.privates { // s is a Soldier
		s.Status()
	}
}
