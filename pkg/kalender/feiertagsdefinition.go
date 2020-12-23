package kalender

import "time"

// Berechnungsart stellt die verschiedenen Berechnungsgrundlage für Feiertage dar
type Berechnungsart int

const (
	// TagDesMonats berechnetr den Feiertag anhand des voreingestellten Monats
	TagDesMonats Berechnungsart = iota
	// BasisOstern berechnet den Feiertag anhand des Datums von Ostersonntag
	BasisOstern
	// BasisVierterAdvent berechnet den Feiertag anhand des Datums vom vierten Advent
	BasisVierterAdvent
)

// Feiertagsdefinition stellt die Berechnung dar, wie das Datum des Feiertags zu berechnen ist
type Feiertagsdefinition struct {
	Name   string
	Art    Berechnungsart
	Land   string
	Region []string
	Monat  time.Month
	Offset int
	Frei   bool
}
