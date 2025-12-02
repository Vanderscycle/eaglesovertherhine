package nations

import (
	"github.com/vanderscycle/eaglesovertherhine/lib/units"
)

var FrenchCompanyDefault = units.CompanyTemplate{
	Soldats:         120,
	Caporals:        8,
	CaporalFourrier: 1,
	Drummers:        2,
	Sergents:        4,
	SergentMajor:    1,
	SousLieutenants: 1,
	Lieutenants:     1,
	Capitaines:      1,
}

var FrenchLineInfantryBattalion = units.BattalionCompanyTemplate{
	Voltigeurs:  1,
	Chasseurs:   0,
	Carabiniers: 0,
	Fusiliers:   4,
	Grenadiers:  1,
}

var FrenchLightInfantryBattalion = units.BattalionCompanyTemplate{
	Voltigeurs:  1,
	Chasseurs:   4,
	Carabiniers: 1,
	Fusiliers:   0,
	Grenadiers:  0,
}
