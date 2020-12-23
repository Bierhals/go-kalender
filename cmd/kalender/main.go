package main

import (
	"fmt"

	"github.com/bierhals/kalender/pkg/kalender"
)

func main() {
	fmt.Println("Neujahr", kalender.Neujahr(2023))
	fmt.Println("Heilige drei Könige", kalender.HeiligeDreiKönige(2023))
	fmt.Println("Weiberfastnacht", kalender.Weiberfastnacht(2023))
	fmt.Println("Fastnachtssamstag", kalender.Fastnachtssamstag(2023))
	fmt.Println("Valentinstag", kalender.Valentinstag(2023))
	fmt.Println("Fastnachtssonntag", kalender.Fastnachtssonntag(2023))
	fmt.Println("Rosenmontag", kalender.Rosenmontag(2023))
	fmt.Println("Fastnacht", kalender.Fastnacht(2023))
	fmt.Println("Aschermittwoch", kalender.Aschermittwoch(2023))
	fmt.Println("Internationaler Frauentag", kalender.InternationalerFrauentag(2023))
	fmt.Println("Josefstag", kalender.Josefstag(2023))
	fmt.Println("Palmsonntag", kalender.Palmsonntag(2023))
	fmt.Println("Gründonnerstag", kalender.Gruendonnerstag(2023))
	fmt.Println("Karfreitag", kalender.Karfreitag(2023))
	fmt.Println("Karsamstag", kalender.Karsamstag(2023))
	fmt.Println("Ostersonntag", kalender.Ostern(2023))
	fmt.Println("Ostermontag", kalender.Ostermontag(2023))
	fmt.Println("Walpurgisnacht", kalender.Walpurgisnacht(2023))
	fmt.Println("Tag der Arbeit", kalender.TagDerArbeit(2023))
	fmt.Println("Muttertag", kalender.Muttertag(2023))
	fmt.Println("Christi Himmelfahrt", kalender.ChristiHimmelfahrt(2023))
	fmt.Println("Vatertag", kalender.Vatertag(2023))
	fmt.Println("Pfingstsonntag", kalender.Pfingstsonntag(2023))
	fmt.Println("Pfingstmontag", kalender.Pfingstmontag(2023))
	fmt.Println("Fronleichname", kalender.Fronleichnam(2023))
	fmt.Println("Johannistag", kalender.Johannistag(2023))
	fmt.Println("PeterUndPaul", kalender.PeterUndPaul(2023))
	fmt.Println("Augsburger Friedensfest", kalender.AugsburgerFriedensfest(2023))
	fmt.Println("Mariä Himmelfahrt", kalender.MariaeHimmelfahrt(2023))
	fmt.Println("Weltkindertag", kalender.Weltkindertag(2023))
	fmt.Println("Tag der Deutschen Einheit", kalender.TagDerDeutschenEinheit(2023))
	fmt.Println("Reformationstag", kalender.Reformationstag(2023))
	fmt.Println("Allerheiligen", kalender.Allerheiligen(2023))
	fmt.Println("Allerseelen", kalender.Allerseelen(2023))
	fmt.Println("Martinstag", kalender.Martinstag(2023))
	fmt.Println("Volkstrauertag", kalender.Volkstrauertag(2023))
	fmt.Println("Buß- und Bettag", kalender.BussUndBettag(2023))
	fmt.Println("Totensonntag", kalender.Totensonntag(2023))
	fmt.Println("Barbaratag", kalender.Barbaratag(2023))
	fmt.Println("Nikolaus", kalender.Nikolaus(2023))
	fmt.Println("Erster Advent", kalender.ErsterAdvent(2023))
	fmt.Println("Zweiter Advent", kalender.ZweiterAdvent(2023))
	fmt.Println("Dritter Advent", kalender.DritterAdvent(2023))
	fmt.Println("Vierter Advent", kalender.VierterAdvent(2023))
	fmt.Println("Heiligabend", kalender.Heiligabend(2023))
	fmt.Println("Erster Weihnachtsfeiertag", kalender.ErsterWeihnachtsfeiertag(2023))
	fmt.Println("Zweiter Weihnachtsfeiertag", kalender.ZweiterWeihnachtsfeiertag(2023))
	fmt.Println("Silvester", kalender.Silvester(2023))

	if feiertage, err := kalender.NewFeiertage(2023, "DE", ""); err == nil {
		for _, f := range *feiertage {
			fmt.Println(f)
		}
	}
}
