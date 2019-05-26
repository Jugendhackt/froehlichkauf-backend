package main

import (
	"errors"
	"fmt"
	"log"
)

func umwelt(material, herkunftNutzer, herkunftProdukt string) (float32, float32, float32, error) {
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

	return 0.5*verpackung + 0.5*herkunft, verpackung, herkunft, fehler
}

/*func ethik(bedingungen string) (float32, error) {
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

	return herstellungsbedingungen, fehler
}*/

func gesundheit(nutritions []nutrition, zutaten []string) (float32, float32, float32) {

	var ingredients []nutrition
	parseJSONFile("ingredients.json", &ingredients)
	fmt.Printf("%+v\n", ingredients)
	fmt.Printf("%+v\n", nutritions)

	var punktzahlInhalt float32
	var zähler int

	for i := range zutaten {
		for j := range ingredients {
			if zutaten[i] == ingredients[j].Name {
				log.Println("Zutat gefunden:" + zutaten[i])
				punktzahlInhalt += ingredients[j].Value
				zähler++
				break
			}
		}
	}

	gesamtpunktzahlInhalt := punktzahlInhalt / float32(zähler)

	var nährwertpunktzahl float32
	var zähler2 int

	for i := range nutritions {
		switch nutritions[i].Name {
		case "sugar":
			if nutritions[i].Value < 0.05 {
				nährwertpunktzahl += 5
				log.Printf("sugar found: +5")
			} else if nutritions[i].Value < 0.225 {
				nährwertpunktzahl += 2
				log.Printf("sugar found: +5")
			} else {
				log.Printf("lisugarpides found: +0")
			}
			zähler2++

		case "salt":
			if nutritions[i].Value < 0.003 {
				nährwertpunktzahl += 5
				log.Printf("salt found: +5")
			} else if nutritions[i].Value < 0.015 {
				nährwertpunktzahl += 2
				log.Printf("salt found: +2")
			} else {
				log.Printf("salt found: +0")
			}
			zähler2++

		case "lipides":
			if nutritions[i].Value < 0.03 {
				nährwertpunktzahl += 5
				log.Printf("lipides found: +5")
			} else if nutritions[i].Value < 0.175 {
				nährwertpunktzahl += 2
				log.Printf("lipides found: +2")
			} else {
				log.Printf("lipides found: +0")
			}
			zähler2++
		}
	}

	gesamtpunktzahlNähwert := nährwertpunktzahl / float32(zähler2)
	ergebnis := 0.2*gesamtpunktzahlInhalt + 0.8*gesamtpunktzahlNähwert

	return ergebnis, gesamtpunktzahlInhalt, gesamtpunktzahlNähwert

}
