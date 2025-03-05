package conversion

import (
	"errors"
	"strconv"
)

func StringToFloat64(strings []string) ([]float64, error) {
	var floats []float64
	for _, stringval := range strings {
		val, err := strconv.ParseFloat(stringval, 64)
		if err != nil {
			return nil, errors.New("error parsing string to float64")
		}
		floats = append(floats, val)
	}

	return floats, nil

}
