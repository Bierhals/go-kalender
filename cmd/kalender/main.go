package main

import (
	"fmt"

	"github.com/bierhals/kalender/pkg/feiertag"
)

func main() {
	fmt.Println("Neujahr", feiertag.Neujahr(2023))
	fmt.Println("Heilige drei Könige", feiertag.HeiligeDreiKönige(2023))
	fmt.Println("Weiberfastnacht", feiertag.Weiberfastnacht(2023))
	fmt.Println("Fastnachtssamstag", feiertag.Fastnachtssamstag(2023))
	fmt.Println("Valentinstag", feiertag.Valentinstag(2023))
	fmt.Println("Fastnachtssonntag", feiertag.Fastnachtssonntag(2023))
	fmt.Println("Rosenmontag", feiertag.Rosenmontag(2023))
	fmt.Println("Fastnacht", feiertag.Fastnacht(2023))
	fmt.Println("Aschermittwoch", feiertag.Aschermittwoch(2023))
	fmt.Println("Internationaler Frauentag", feiertag.InternationalerFrauentag(2023))
	fmt.Println("Josefstag", feiertag.Josefstag(2023))
	fmt.Println("Palmsonntag", feiertag.Palmsonntag(2023))
	fmt.Println("Gründonnerstag", feiertag.Gruendonnerstag(2023))
	fmt.Println("Karfreitag", feiertag.Karfreitag(2023))
	fmt.Println("Karsamstag", feiertag.Karsamstag(2023))
	fmt.Println("Ostersonntag", feiertag.Ostern(2023))
	fmt.Println("Ostermontag", feiertag.Ostermontag(2023))
	fmt.Println("Walpurgisnacht", feiertag.Walpurgisnacht(2023))
	fmt.Println("Tag der Arbeit", feiertag.TagDerArbeit(2023))
	fmt.Println("Muttertag", feiertag.Muttertag(2023))
	fmt.Println("Christi Himmelfahrt", feiertag.ChristiHimmelfahrt(2023))
	fmt.Println("Vatertag", feiertag.Vatertag(2023))
	fmt.Println("Pfingstsonntag", feiertag.Pfingstsonntag(2023))
	fmt.Println("Pfingstmontag", feiertag.Pfingstmontag(2023))
	fmt.Println("Fronleichname", feiertag.Fronleichnam(2023))
	fmt.Println("Johannistag", feiertag.Johannistag(2023))
	fmt.Println("PeterUndPaul", feiertag.PeterUndPaul(2023))
	fmt.Println("Augsburger Friedensfest", feiertag.AugsburgerFriedensfest(2023))
	fmt.Println("Mariä Himmelfahrt", feiertag.MariaeHimmelfahrt(2023))
	fmt.Println("Weltkindertag", feiertag.Weltkindertag(2023))
	fmt.Println("Tag der Deutschen Einheit", feiertag.TagDerDeutschenEinheit(2023))
	fmt.Println("Reformationstag", feiertag.Reformationstag(2023))
	fmt.Println("Allerheiligen", feiertag.Allerheiligen(2023))
	fmt.Println("Allerseelen", feiertag.Allerseelen(2023))
	fmt.Println("Martinstag", feiertag.Martinstag(2023))
	fmt.Println("Volkstrauertag", feiertag.Volkstrauertag(2023))
	fmt.Println("Buß- und Bettag", feiertag.BussUndBettag(2023))
	fmt.Println("Totensonntag", feiertag.Totensonntag(2023))
	fmt.Println("Barbaratag", feiertag.Barbaratag(2023))
	fmt.Println("Nikolaus", feiertag.Nikolaus(2023))
	fmt.Println("Erster Advent", feiertag.ErsterAdvent(2023))
	fmt.Println("Zweiter Advent", feiertag.ZweiterAdvent(2023))
	fmt.Println("Dritter Advent", feiertag.DritterAdvent(2023))
	fmt.Println("Vierter Advent", feiertag.VierterAdvent(2023))
	fmt.Println("Heiligabend", feiertag.Heiligabend(2023))
	fmt.Println("Erster Weihnachtsfeiertag", feiertag.ErsterWeihnachtsfeiertag(2023))
	fmt.Println("Zweiter Weihnachtsfeiertag", feiertag.ZweiterWeihnachtsfeiertag(2023))
	fmt.Println("Silvester", feiertag.Silvester(2023))

	if feiertage, err := feiertag.NewFeiertage(2023, "DE", ""); err == nil {
		for _, f := range *feiertage {
			fmt.Println(f)
		}
	}
}
