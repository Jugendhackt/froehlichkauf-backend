package main

import (
	"errors"
	"log"
)

func umwelt(material, herkunftNutzer, herkunftProdukt string) (float32, float32, float32) {
	var verpackung float32
	var fehler error
	var herkunft float32

	switch material {
	case "plastic":
		verpackung = 0
	case "paper":
		verpackung = 2
	case "glass":
		verpackung = 3
	case "none":
		verpackung = 5

	default:
		fehler = errors.New("ungültiges material")
	}

	if fehler != nil {
		log.Println(fehler.Error())
	}

	if herkunftNutzer == herkunftProdukt {
		herkunft = 5
	} else {
		herkunft = 0
	}

	return 0.5*verpackung + 0.5*herkunft, verpackung, herkunft
}

func ethik(bedingungen string) float32 {
	var herstellungsbedingungen float32
	var fehler error

	switch bedingungen {
	case "gut":
		herstellungsbedingungen = 5
	case "schlecht":
		herstellungsbedingungen = 0

	default:
		fehler = errors.New("ungültige bedingungen")
	}

	if fehler != nil {
		log.Println(fehler.Error())
	}

	return herstellungsbedingungen
}

/*
func gesundheit(nutritions []nutrition, zutaten []string) {

}
*/
