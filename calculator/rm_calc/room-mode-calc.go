package rm_calc

import (
	"errors"
	"math"
	"sort"
)

func CalcOblique(length int32, width int32, height int32) ([]float32, error) {
	err := checkInput(length, width, height)
	if err != nil {
		return nil, err
	}

	// container for result sets
	var result []float32

	i := 0
	for {
		mode := int32(i + 1)
		calculated := calcMode(mode, length, width, height)
		if calculated <= 500 {
			result = append(result, calculated)
			i++
		} else {
			break
		}
	}

	sortResult(result)
	return result, nil
}

func CalcTangential(length int32, width int32, height int32) ([]float32, error) {
	err := checkInput(length, width, height)
	if err != nil {
		return nil, err
	}

	// container for result sets
	var result []float32

	// iterate through modes
	for i := 0; i < 20; i++ {
		mode := int32(i + 1)
		for k := 0; k < 3; k++ {
			var calculated float32
			if k == 0 {
				calculated = getTangential(mode, length, width)
			} else if k == 1 {
				calculated = getTangential(mode, length, height)
			} else if k == 2 {
				calculated = getTangential(mode, width, height)
			}
			if calculated <= 500 {
				result = append(result, calculated)
			}
		}
	}

	result = sortResult(result)
	return result, nil
}

func CalcAxial(length int32, width int32, height int32) ([]float32, error) {
	err := checkInput(length, width, height)
	if err != nil {
		return nil, err
	}

	var result []float32

	for i := 0; i < 3; i++ {
		var v int32

		// iterate through l, w, h
		if i == 0 {
			v = length
		} else if i == 1 {
			v = width
		} else if i == 2 {
			v = height
		}
		// calculate 20 times with modes

		k := 0
		for {
			// 0 ~ 20
			mode := int32(k + 1)
			value := getAxial(mode, v)
			// collect value under 500 Hertz
			if value <= 500 {
				result = append(result, value)
				k++
			} else {
				break
			}
		}
	}

	// sort asc
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return result, nil
}

//wrapper for calcMode with axial
func getAxial(mode int32, value int32) float32 {
	return calcMode(mode, value, 0, 0)
}

//wrapper for calcMode with tangential
func getTangential(mode int32, firstDir int32, secondDir int32) float32 {
	return calcMode(mode, firstDir, secondDir, 0)
}

/**
Calculation is based on following formula...

f = (c/2) * sqrt((n/L)^2 + (n/W)^2 + (n/H)^2)

f = frequency of the mode
c = speed of sound in meter
n = order of the mode
L, W, H = length, width, height of the room in meter
*/
func calcMode(modeInput int32, l int32, w int32, h int32) float32 {
	//exchange cm to m
	mode := float64(modeInput)
	length := float64(l) / 100
	width := float64(w) / 100
	height := float64(h) / 100
	//speed of sound (default to 344m/s)
	c := float64(344)

	var result float64

	if w == 0 {
		// axial calculation
		result = (c / 2) * math.Sqrt(math.Pow(mode/length, 2))
	} else if h == 0 {
		// tangential calculation
		result = (c / 2) * math.Sqrt(math.Pow(mode/length, 2)+math.Pow(mode/width, 2))
	} else {
		result = (c / 2) * math.Sqrt(math.Pow(mode/length, 2)+math.Pow(mode/width, 2)+math.Pow(mode/height, 2))
	}
	//round result to two decimal places
	rounded := float32(math.Round(result*100) / 100)
	return rounded
}

//check invalid input
func checkInput(l int32, w int32, h int32) error {
	err := errors.New("invalid input detected")
	if l <= 0 || w <= 0 || h <= 0 {
		return err
	}
	return nil
}

//sort result slice in ascending order
func sortResult(source []float32) []float32 {
	result := source
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}
