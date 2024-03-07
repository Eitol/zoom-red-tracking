package zoom

import (
	"encoding/json"
	"fmt"
	"strings"
)

type OpenHours struct {
	Mon []HourRange
	Tue []HourRange
	Wed []HourRange
	Thu []HourRange
	Fri []HourRange
	Sat []HourRange
	Sun []HourRange // Asumimos que queremos mantener la estructura, aunque "sun" pueda no tener horarios.
}

type HourRange struct {
	Start string
	End   string
}

// Función auxiliar para parsear un string de horario a una estructura HourRange.
func parseHourRange(hourString string) []HourRange {
	if hourString == "0" {
		// Si el día está cerrado, devolvemos un slice vacío.
		return []HourRange{}
	}

	// Dividimos los horarios por coma, en caso de que haya más de un rango de horario por día.
	ranges := strings.Split(hourString, ",")
	var hourRanges []HourRange

	for _, r := range ranges {
		times := strings.Split(r, " - ")
		if len(times) == 2 {
			hourRanges = append(hourRanges, HourRange{Start: times[0], End: times[1]})
		}
	}
	return hourRanges
}

// Función que recibe el string JSON y retorna una estructura OpenHours.
func parseOpenHours(jsonStr string) (*OpenHours, error) {
	var rawHours map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &rawHours)
	if err != nil {
		return nil, err
	}

	openHours := OpenHours{}

	// Iteramos sobre cada día de la semana.
	for day, value := range rawHours {
		var hourStrings []string
		// Comprobamos el tipo de dato para el día actual.
		switch v := value.(type) {
		case []interface{}:
			for _, hour := range v {
				hourStrings = append(hourStrings, hour.(string))
			}
		case string:
			hourStrings = append(hourStrings, v)
		}

		// Parseamos los horarios para cada día.
		for _, hourString := range hourStrings {
			switch day {
			case "mon":
				openHours.Mon = append(openHours.Mon, parseHourRange(hourString)...)
			case "tue":
				openHours.Tue = append(openHours.Tue, parseHourRange(hourString)...)
			case "wed":
				openHours.Wed = append(openHours.Wed, parseHourRange(hourString)...)
			case "thu":
				openHours.Thu = append(openHours.Thu, parseHourRange(hourString)...)
			case "fri":
				openHours.Fri = append(openHours.Fri, parseHourRange(hourString)...)
			case "sat":
				openHours.Sat = append(openHours.Sat, parseHourRange(hourString)...)
			case "sun":
				openHours.Sun = append(openHours.Sun, parseHourRange(hourString)...)
			default:
				return nil, fmt.Errorf("unexpected day: %s", day)
			}
		}
	}

	return &openHours, nil
}
