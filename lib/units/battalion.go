package units

import (
	"fmt"
	"reflect"

	"github.com/vanderscycle/eaglesovertherhine/lib/pops"
)

type battalion struct {
	name                 string
	bloodname            string
	adjudantSousOfficier []pops.MilitaryPop
	adjudantMajor        []pops.MilitaryPop
	chefDeBattalion      []pops.MilitaryPop
	flagBearer           pops.MilitaryPop
	drivers              []pops.CivilianPop
	cantinieres          []pops.CivilianPop
	companies            []company
	companyTemplate      CompanyTemplate
	template             BattalionCompanyTemplate
	nation               pops.Nation
	history              string
}

func NewBattalion(name string, template BattalionCompanyTemplate, companyTemplate CompanyTemplate, nation pops.Nation) battalion {
	b := battalion{name: name, bloodname: "unearned", template: template, history: "TODO"}

	// battalion staff generation
	b.adjudantSousOfficier = append(b.adjudantSousOfficier, pops.NewSoldier(pops.AdjudantSousOfficier, nation, pops.Nco))
	b.adjudantMajor = append(b.adjudantMajor, pops.NewSoldier(pops.AdjudantMajor, nation, pops.Nco))
	b.chefDeBattalion = append(b.chefDeBattalion, pops.NewSoldier(pops.ChefDeBattalion, nation, pops.Officier))
	b.flagBearer = pops.NewSoldier(pops.Caporal, nation, pops.MusicColours)

	for i:=0; i<=4;i++{
		b.drivers = append(b.drivers,pops.NewCivilian(b.nation, pops.MilitarySupport))
		b.cantinieres = append(b.drivers,pops.NewCivilian(b.nation, pops.MilitarySupport))
	}

	// company generation
	tmplVal := reflect.ValueOf(b.template)
	tmplType := reflect.TypeOf(b.template)

	for i := 0; i < tmplVal.NumField(); i++ {
		fieldName := tmplType.Field(i).Name
		count := int(tmplVal.Field(i).Int())
		//fmt.Printf("%s\n", fieldName)
		for j := range count {
			compInt := j + 1
			// fmt.Printf("\n%d %s\n", compInt, fieldName)
			switch fieldName {
			case "Chasseurs":
				b.companies = append(b.companies, NewCompany(fmt.Sprintf("%d Chasseurs", compInt), companyTemplate, nation))
			case "Voltigeurs":
				b.companies = append(b.companies, NewCompany(fmt.Sprintf("%d Voltigeurs", compInt), companyTemplate, nation))
			case "Carabiniers":
				b.companies = append(b.companies, NewCompany(fmt.Sprintf("%d Carabiniers", compInt), companyTemplate, nation))
			case "Fusiliers":
				b.companies = append(b.companies, NewCompany(fmt.Sprintf("%d Fusiliers", compInt), companyTemplate, nation))
			case "Grenadiers":
				b.companies = append(b.companies, NewCompany(fmt.Sprintf("%d Grenadiers", compInt), companyTemplate, nation))
			}
		}
	}
	return b
}

func (b battalion) Status() {
	fmt.Printf("The battalion's name: %s, nickname: %s and country %s\nTemplate: %+v\n", b.name, b.bloodname, b.nation, b.template)
}

func (b *battalion) EarnBloodName(name string) {
	b.bloodname = name
}

// TODO: make a note about this func returning a pointer
func (b battalion) GetCompany(name string) *company {
	for i := range b.companies {
		fmt.Printf("company: %s bloodname: %s\n", b.companies[i].bloodname, b.companies[i].name)
		if b.companies[i].name == name {
			fmt.Printf("found comp's capitatine! %+v\n", b.companies[i].capitaines)
			return &b.companies[i]
		}
	}
	return &b.companies[0]
}

func (b battalion) ListCompanies() {
	for _, s := range b.companies {
		fmt.Printf("%+v\n", s)
	}
}
