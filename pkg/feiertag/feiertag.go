package feiertag

import (
	"fmt"
	"time"
)

// Feiertag mit den Angaben zum Land und der Region und ob dieser dort als
type Feiertag struct {
	Name       string
	Land       string
	Region     *[]string
	Datum      time.Time
	Frei       bool
}

// NewFeiertage erstellt eine Liste der Feiertage
func NewFeiertage(jahr int, land string, region string) (*[]Feiertag, error) {
	if (jahr < 1900 || jahr >= 3000) {
		return nil, fmt.Errorf("Jahr '%v' ungültig. Das Jahr muss zwischen 1900 und 2999 liegen", jahr)
	}

	switch land {
		case "DE":
			return newFeiertageDE(jahr, region), nil
		default:
			return nil, fmt.Errorf("Land '%v' unbekannt", land)
	}
}

func newFeiertageDE(jahr int, region string) *[]Feiertag {
	
	feiertage := make([]Feiertag, 3)
	feiertage[0] = Feiertag{
		Name: "Neujahr",
		Land: "DE",
		Datum: Neujahr(jahr),
		Frei: true,
	}
	feiertage[1] = Feiertag{
		Name: "Ostern",
		Land: "DE",
		Datum: Ostern(jahr),
		Frei: true,
	}
	feiertage[2] = Feiertag{
		Name: "Silvester",
		Land: "DE",
		Datum: Silvester(jahr),
		Frei: true,
	}

	return &feiertage
}

func newFeiertageAT(jahr int, region string) *[]Feiertag {
	return nil
}

func (f Feiertag) String() string {
	return fmt.Sprintf("%v - %v (%v)", f.Land, f.Name, f.Datum.Format("2006-01-02"))
}

// Ostern berechnet das Datum vom Ostersonntag für das übegebene Jahr
func Ostern(jahr int) time.Time {
	k := jahr / 100
	m := 15 + (3*k+3)/4 - (8*k+13)/25
	s := 2 - (3*k+3)/4
	a := jahr % 19
	d := (19*a + m) % 30
	r := (d + a/11) / 29
	og := 21 + d - r
	sz := 7 - (jahr+jahr/4+s)%7
	oe := 7 - (og-sz)%7
	os := og + oe

	day := os % 31
	month := os/31 + 3

	return time.Date(jahr, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

// Neujahr berechnet das Datum von Neujahr für das übergebene Jahr
func Neujahr(jahr int) time.Time {
	return time.Date(jahr, time.January, 1, 0, 0, 0, 0, time.UTC)
}

// HeiligeDreiKönige berechnet das Datum von "Heilige drei Könige" für das übergebene Jar
func HeiligeDreiKönige(jahr int) time.Time {
	return time.Date(jahr, time.January, 6, 0, 0, 0, 0, time.UTC)
}

// Weiberfastnacht berechnet das Datum von Weiberfastnacht für das übergebene Jahr
func Weiberfastnacht(jahr int) time.Time {
	return Ostern(jahr).AddDate(0, 0, -52)
}

// Fastnachtssamstag berechnet das Datum vom Fastnachtssamstag für das übergebene Jahr
func Fastnachtssamstag(jahr int) time.Time {
	return Ostern(jahr).AddDate(0, 0, -50)
}

// Valentinstag berechnet das Datum vom Valentinstag für das übergebene Jahr
func Valentinstag(jahr int) time.Time {
	return time.Date(jahr, time.February, 14, 0, 0, 0, 0, time.UTC)
}

// Fastnachtssonntag berechnet das Datum vom Fastnachtssonntag für das übergebene Jahr
func Fastnachtssonntag(jahr int) time.Time {
	return Ostern(jahr).AddDate(0, 0, -49)
}

// Rosenmontag berechnet das Datum vom Rosenmontag für das übergebene Jahr
func Rosenmontag(jahr int) time.Time {
	return Ostern(jahr).AddDate(0, 0, -48)
}

// Fastnacht berechnet das Datum von Fastnacht für das übergebene Jahr
func Fastnacht(jahr int) time.Time {
	return Ostern(jahr).AddDate(0, 0, -47)
}

// Aschermittwoch berechnet das Datum von Aschermittwoch für das übergebene Jahr
func Aschermittwoch(jahr int) time.Time {
	return Ostern(jahr).AddDate(0, 0, -46)
}

// InternationalerFrauentag berechnet das Datum von "Internationaler Frauentag" für das übergebene Jar
func InternationalerFrauentag(jahr int) time.Time {
	return time.Date(jahr, time.March, 8, 0, 0, 0, 0, time.UTC)
}

// Josefstag berechnet das Datum von Josefstag für das übergebene Jahr
func Josefstag(jahr int) time.Time {
	return time.Date(jahr, time.March, 19, 0, 0, 0, 0, time.UTC)
}

// Palmsonntag berechnet das Datum von Palmsonntag für das übergebene Jahr
func Palmsonntag(jahr int) time.Time {
	return Ostern(jahr).AddDate(0, 0, -7)
}

// Gruendonnerstag berechnet das Datum von Gründonnerstag für das übergebene Jahr
func Gruendonnerstag(jahr int) time.Time {
	return Ostern(jahr).AddDate(0, 0, -3)
}

// Karfreitag berechnet das Datum von Karfreitag für das übergebene Jahr
func Karfreitag(jahr int) time.Time {
	return Ostern(jahr).AddDate(0, 0, -2)
}

// Karsamstag berechnet das Datum von Karsamstag für das übergebene Jahr
func Karsamstag(jahr int) time.Time {
	return Ostern(jahr).AddDate(0, 0, -1)
}

// Ostermontag berechnet das Datum von Ostermontag für das übergebene Jahr
func Ostermontag(jahr int) time.Time {
	return Ostern(jahr).AddDate(0, 0, 1)
}

// Walpurgisnacht berechnet das Datum von Walpurgisnacht für das übergebene Jahr
func Walpurgisnacht(jahr int) time.Time {
	return time.Date(jahr, time.April, 30, 0, 0, 0, 0, time.UTC)
}

// TagDerArbeit berechnet das Datum von "Tag der Arbeit" für das übergebene Jar
func TagDerArbeit(jahr int) time.Time {
	return time.Date(jahr, time.May, 1, 0, 0, 0, 0, time.UTC)
}

// Muttertag berechnet das Datum von Muttertag für das übergebene Jahr
func Muttertag(jahr int) time.Time {
	d := time.Date(jahr, time.May, 1, 0, 0, 0, 0, time.UTC)
	d = d.AddDate(0, 0, (7-int(d.Weekday())%7)+7)
	return d
}

// ChristiHimmelfahrt berechnet das Datum von "Christi Himmelfahrt" für das übergebene Jar
func ChristiHimmelfahrt(jahr int) time.Time {
	return Ostern(jahr).AddDate(0, 0, 39)
}

// Vatertag berechnet das Datum von Vatertag für das übergebene Jahr
func Vatertag(jahr int) time.Time {
	return ChristiHimmelfahrt(jahr)
}

// Pfingstsonntag berechnet das Datum von Pfingstsonntag für das übergebene Jahr
func Pfingstsonntag(jahr int) time.Time {
	return Ostern(jahr).AddDate(0, 0, 49)
}

// Pfingstmontag berechnet das Datum von Pfingstmontag für das übergebene Jahr
func Pfingstmontag(jahr int) time.Time {
	return Ostern(jahr).AddDate(0, 0, 50)
}

// Fronleichnam berechnet das Datum von Fronleichnam für das übergebene Jahr
func Fronleichnam(jahr int) time.Time {
	return Ostern(jahr).AddDate(0, 0, 60)
}

// Johannistag berechnet das Datum von Johannistag für das übergebene Jahr
func Johannistag(jahr int) time.Time {
	return time.Date(jahr, time.June, 24, 0, 0, 0, 0, time.UTC)
}

// PeterUndPaul berechnet das Datum von "Peter und Paul" für das übergebene Jar
func PeterUndPaul(jahr int) time.Time {
	return time.Date(jahr, time.June, 29, 0, 0, 0, 0, time.UTC)
}

// AugsburgerFriedensfest berechnet das Datum von "Augsburger Friedensfest" für das übergebene Jr
func AugsburgerFriedensfest(jahr int) time.Time {
	return time.Date(jahr, time.August, 8, 0, 0, 0, 0, time.UTC)
}

// MariaeHimmelfahrt berechnet das Datum von "Mariä Himmelfahrt für" das übergebene Jar
func MariaeHimmelfahrt(jahr int) time.Time {
	return time.Date(jahr, time.August, 15, 0, 0, 0, 0, time.UTC)
}

// Weltkindertag berechnet das Datum vom Weltkindertag für das übergebene Jahr
func Weltkindertag(jahr int) time.Time {
	return time.Date(jahr, time.September, 20, 0, 0, 0, 0, time.UTC)
}

// TagDerDeutschenEinheit berechnet das Datum vom "Tag der Deutschen Einheit" für das übergebene Jhr
func TagDerDeutschenEinheit(jahr int) time.Time {
	return time.Date(jahr, time.October, 3, 0, 0, 0, 0, time.UTC)
}

// Reformationstag berechnet das Datum vom Reformationstag für das übergebene Jah
func Reformationstag(jahr int) time.Time {
	return time.Date(jahr, time.October, 31, 0, 0, 0, 0, time.UTC)
}

// Allerheiligen berechnet das Datum von Allerheiligen für das übergebene Jah
func Allerheiligen(jahr int) time.Time {
	return time.Date(jahr, time.November, 1, 0, 0, 0, 0, time.UTC)
}

// Allerseelen berechnet das Datum von Allerseelen für das übergebene Jar
func Allerseelen(jahr int) time.Time {
	return time.Date(jahr, time.November, 2, 0, 0, 0, 0, time.UTC)
}

// Martinstag berechnet das Datum vom Martinstag für das übergebene Jahr
func Martinstag(jahr int) time.Time {
	return time.Date(jahr, time.November, 11, 0, 0, 0, 0, time.UTC)
}

// Volkstrauertag berechnet das Datum vom Volkstrauertag für das übergebene Jah
func Volkstrauertag(jahr int) time.Time {
	return VierterAdvent(jahr).AddDate(0, 0, -35)
}

// BussUndBettag berechnet das Datum vom "Buß- und Bettag" für das übergebeneJahr
func BussUndBettag(jahr int) time.Time {
	return VierterAdvent(jahr).AddDate(0, 0, -32)
}

// Totensonntag berechnet das Datum vom Totensonntag für das übergebene Jah
func Totensonntag(jahr int) time.Time {
	return VierterAdvent(jahr).AddDate(0, 0, -28)
}

// Barbaratag berechnet das Datum vom Barbaratag für das übergebene Jah
func Barbaratag(jahr int) time.Time {
	return time.Date(jahr, time.December, 4, 0, 0, 0, 0, time.UTC)
}

// Nikolaus berechnet das Datum von Nikolaus für das übergebene Jahr
func Nikolaus(jahr int) time.Time {
	return time.Date(jahr, time.December, 6, 0, 0, 0, 0, time.UTC)
}

// ErsterAdvent berechnet das Datum vom "erster Advent" für das übergebene Jhr
func ErsterAdvent(jahr int) time.Time {
	return VierterAdvent(jahr).AddDate(0, 0, -21)
}

// ZweiterAdvent berechnet das Datum vom "zweiter Advent" für das übergebene Jar
func ZweiterAdvent(jahr int) time.Time {
	return VierterAdvent(jahr).AddDate(0, 0, -14)
}

// DritterAdvent berechnet das Datum vom "dritter Advent "für das übergebene Jar
func DritterAdvent(jahr int) time.Time {
	return VierterAdvent(jahr).AddDate(0, 0, -7)
}

// VierterAdvent berechnet das Datum vom "vierter Advent" für das übergebene Jar
func VierterAdvent(jahr int) time.Time {
	d := time.Date(jahr, time.December, 2, 0, 0, 0, 0, time.UTC)
	return d.AddDate(0, 0, -1*((7+int(d.Weekday()))%7))
}

// Heiligabend berechnet das Datum von Heiligabend für das übergebene Jahr
func Heiligabend(jahr int) time.Time {
	return time.Date(jahr, time.December, 24, 0, 0, 0, 0, time.UTC)
}

// ErsterWeihnachtsfeiertag berechnet das Datum vom "erster Weihnachtsfeiertag" für das übergebene Jar
func ErsterWeihnachtsfeiertag(jahr int) time.Time {
	return time.Date(jahr, time.December, 25, 0, 0, 0, 0, time.UTC)
}

// ZweiterWeihnachtsfeiertag berechnet das Datum vom "zweiten Weihnachtfeiertag" für das übergebene Jar
func ZweiterWeihnachtsfeiertag(jahr int) time.Time {
	return time.Date(jahr, time.December, 26, 0, 0, 0, 0, time.UTC)
}

// Silvester berechnet das Datum von Silvester für das übergebene Jahr
func Silvester(jahr int) time.Time {
	return time.Date(jahr, time.December, 31, 0, 0, 0, 0, time.UTC)
}
