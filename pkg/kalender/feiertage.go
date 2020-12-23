package kalender

import (
	"fmt"
	"time"
)

// Feiertag mit den Angaben zum Land und der Region und ob dieser dort als
type Feiertag struct {
	Name   string
	Land   string
	Region []string
	Datum  time.Time
	Frei   bool
}

// NewFeiertage erstellt eine Liste der Feiertage
func NewFeiertage(jahr int, land string, region string) (*[]Feiertag, error) {
	if jahr < 1900 || jahr >= 3000 {
		return nil, fmt.Errorf("Jahr '%v' ungültig. Das Jahr muss zwischen 1900 und 2999 liegen", jahr)
	}

	var definition *[]Feiertagsdefinition

	switch land {
	case "DE":
		definition = feiertagsdefinitionDE()
	default:
		return nil, fmt.Errorf("Land '%v' unbekannt", land)
	}

	ostern := Ostern(jahr)
	vierterAdvent := VierterAdvent(jahr)

	feiertage := make([]Feiertag, len(*definition), len(*definition))
	for i, d := range *definition {
		switch d.Art {
		case TagDesMonats:
			feiertage[i] = berechneFeiertagDesMonats(jahr, d)
		case BasisOstern:
			feiertage[i] = berechneFeiertagBasisdatum(ostern, d)
		case BasisVierterAdvent:
			feiertage[i] = berechneFeiertagBasisdatum(vierterAdvent, d)
		}

	}
	return &feiertage, nil
}

// String gibt die Feiertagsinformationen in Menschenlesbarer Form aus
func (f Feiertag) String() string {
	return fmt.Sprintf("%v - %v (%v) Frei:%v Regionen:%v", f.Land, f.Name, f.Datum.Format("2006-01-02"), f.Frei, f.Region)
}

func berechneFeiertagDesMonats(jahr int, definition Feiertagsdefinition) Feiertag {
	return Feiertag{
		Name:   definition.Name,
		Land:   definition.Land,
		Region: definition.Region,
		Frei:   definition.Frei,
		Datum:  time.Date(jahr, definition.Monat, definition.Offset, 0, 0, 0, 0, time.UTC),
	}
}

func berechneFeiertagBasisdatum(datum time.Time, definition Feiertagsdefinition) Feiertag {
	return Feiertag{
		Name:   definition.Name,
		Land:   definition.Land,
		Region: definition.Region,
		Frei:   definition.Frei,
		Datum:  datum.AddDate(0, 0, definition.Offset),
	}
}
