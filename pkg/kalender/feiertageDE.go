package kalender

import (
	"time"
)

func feiertagsdefinitionDE() *[]Feiertagsdefinition {
	definition := make([]Feiertagsdefinition, 19, 19)
	definition[0] = Feiertagsdefinition{
		Name:   "Neujahr",
		Land:   "DE",
		Art:    TagDesMonats,
		Frei:   true,
		Monat:  time.January,
		Offset: 1,
	}
	definition[1] = Feiertagsdefinition{
		Name:   "Heilige drei Könige",
		Land:   "DE",
		Region: []string{"DE-BW", "DE-BY", "DE-ST"},
		Art:    TagDesMonats,
		Frei:   true,
		Monat:  time.January,
		Offset: 6,
	}
	definition[2] = Feiertagsdefinition{
		Name:   "Weiberfastnacht",
		Land:   "DE",
		Art:    BasisOstern,
		Frei:   false,
		Monat:  time.January,
		Offset: -52,
	}
	definition[3] = Feiertagsdefinition{
		Name:   "Fastnachtssamstag",
		Land:   "DE",
		Art:    BasisOstern,
		Frei:   false,
		Offset: -50,
	}
	definition[4] = Feiertagsdefinition{
		Name:   "Fastnachtssonntag",
		Land:   "DE",
		Art:    BasisOstern,
		Frei:   false,
		Offset: -49,
	}
	definition[5] = Feiertagsdefinition{
		Name:   "Valentinstag",
		Land:   "DE",
		Art:    TagDesMonats,
		Frei:   false,
		Monat:  time.February,
		Offset: 14,
	}
	definition[6] = Feiertagsdefinition{
		Name:   "Rosenmontag",
		Land:   "DE",
		Art:    BasisOstern,
		Frei:   false,
		Offset: -48,
	}
	definition[7] = Feiertagsdefinition{
		Name:   "Fastnacht",
		Land:   "DE",
		Art:    BasisOstern,
		Frei:   false,
		Offset: -47,
	}
	definition[8] = Feiertagsdefinition{
		Name:   "Aschermittwoch",
		Land:   "DE",
		Art:    BasisOstern,
		Frei:   false,
		Offset: -46,
	}
	definition[9] = Feiertagsdefinition{
		Name:   "Internationaler Frauentag",
		Land:   "DE",
		Region: []string{"DE-BE"},
		Art:    TagDesMonats,
		Frei:   true,
		Monat:  time.March,
		Offset: 8,
	}
	definition[10] = Feiertagsdefinition{
		Name:   "Josefstag",
		Land:   "DE",
		Region: []string{"DE-BY"},
		Art:    TagDesMonats,
		Frei:   false,
		Monat:  time.March,
		Offset: 19,
	}
	definition[11] = Feiertagsdefinition{
		Name:   "Palmsonntag",
		Land:   "DE",
		Art:    BasisOstern,
		Frei:   false,
		Offset: -7,
	}
	definition[12] = Feiertagsdefinition{
		Name:   "Gründonnerstag",
		Land:   "DE",
		Art:    BasisOstern,
		Frei:   false,
		Offset: -3,
	}
	definition[13] = Feiertagsdefinition{
		Name:   "Karfreitag",
		Land:   "DE",
		Art:    BasisOstern,
		Frei:   true,
		Offset: -2,
	}
	definition[14] = Feiertagsdefinition{
		Name:   "Karsamstag",
		Land:   "DE",
		Art:    BasisOstern,
		Frei:   true,
		Offset: -1,
	}
	definition[15] = Feiertagsdefinition{
		Name:   "Ostersonntag",
		Land:   "DE",
		Region: []string{"DE-BB"},
		Art:    BasisOstern,
		Frei:   true,
		Offset: 0,
	}
	definition[16] = Feiertagsdefinition{
		Name:   "Ostermontag",
		Land:   "DE",
		Art:    BasisOstern,
		Frei:   true,
		Offset: 1,
	}
	definition[17] = Feiertagsdefinition{
		Name:   "erster Advent",
		Land:   "DE",
		Art:    BasisVierterAdvent,
		Frei:   false,
		Monat:  time.December,
		Offset: -21,
	}
	definition[18] = Feiertagsdefinition{
		Name:   "Weihnachten",
		Land:   "DE",
		Art:    TagDesMonats,
		Frei:   false,
		Monat:  time.December,
		Offset: 24,
	}

	return &definition
}
