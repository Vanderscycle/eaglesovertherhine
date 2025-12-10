package units

import (
	"fmt"
	"reflect"

	"github.com/vanderscycle/eaglesovertherhine/lib/pops"
)

// TODO: create person class w/ rank, name,  + history (battles/medals/promotions)
type company struct {
	name            string
	bloodname       string
	soldats         []pops.MilitaryPop
	caporals        []pops.MilitaryPop
	caporalFourrier []pops.MilitaryPop
	drummers        []pops.MilitaryPop
	sergents        []pops.MilitaryPop
	sergentMajor    []pops.MilitaryPop
	sousLieutenants []pops.MilitaryPop
	lieutenants     []pops.MilitaryPop
	capitaines      []pops.MilitaryPop
	template        CompanyTemplate
	nation          pops.Nation
	history         string
}

func NewCompany(name string, template CompanyTemplate, nation pops.Nation) company {
	c := company{name: name, bloodname: "unearned", template: template, history: "TODO"}

	tmplVal := reflect.ValueOf(c.template)
	tmplType := reflect.TypeOf(c.template)

	for i := 0; i < tmplVal.NumField(); i++ {
		fieldName := tmplType.Field(i).Name
		count := int(tmplVal.Field(i).Int())
		//fmt.Printf("%s\n", fieldName)
		for j := 0; j <= count; j++ {
			//fmt.Printf("%d\n", j)
			switch fieldName {
			case "Soldats":
				c.soldats = append(c.soldats, pops.NewSoldier(pops.Soldat, nation, pops.Soldier))
			case "Caporals":
				c.caporals = append(c.caporals, pops.NewSoldier(pops.Caporal, nation, pops.Soldier))
			case "CaporalFourrier":
				c.caporalFourrier = append(c.caporalFourrier, pops.NewSoldier(pops.CaporalFourrier, nation, pops.Logistics))
			case "Drummers":
				c.drummers = append(c.drummers, pops.NewSoldier(pops.Soldat, nation, pops.MusicColours))
			case "Sergents":
				c.sergents = append(c.sergents, pops.NewSoldier(pops.Sergent, nation, pops.Nco))
			case "SergentMajor":
				c.sergentMajor = append(c.sergentMajor, pops.NewSoldier(pops.SergentMajor, nation, pops.Nco))
			case "SousLieutenants":
				c.sousLieutenants = append(c.sousLieutenants, pops.NewSoldier(pops.SousLieutenant, nation, pops.Officier))
			case "Lieutenants":
				c.lieutenants = append(c.lieutenants, pops.NewSoldier(pops.Lieutenant, nation, pops.Officier))
			case "Capitaines":
				c.capitaines = append(c.capitaines, pops.NewSoldier(pops.Capitaine, nation, pops.Officier))
			}
		}
	}
	return c
}

func (c company) Status() {
	fmt.Printf("The company's name: %s and country %s\nTemplate: %+v\n", c.name, c.nation, c.template)
}

func (c *company) EarnBloodName(name string) {
	c.bloodname = name
}

func (c company) ListSoldiers() {
	for i := range c.soldats {
		c.soldats[i].Status()
	}
}
