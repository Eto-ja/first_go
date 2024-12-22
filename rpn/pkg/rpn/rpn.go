package rpn

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Calc(expression string) (float64, error) {
	litters := strings.Split(expression, "")
	k1 := make([]float64, 0)
	k2 := make([]string, 0)
	er := errors.New("Что-то пошло не так")

	pr := false
	k_pr := ""
	k_12 := make([]string, 0)

	if len(expression) == 0 {
		return 0, er
	}

	for i := 0; i < len(expression); i++ {
		if litters[i] == "(" {
			pr = true
		} else if litters[i] == ")" {
			pr = false
			kk, err := Calc(k_pr)
			if err != nil {
				return 0, er
			}
			k_12 = append(k_12, fmt.Sprintf("%f", kk))
			k_pr = ""
		} else if pr {
			k_pr += litters[i]
		} else {
			k_12 = append(k_12, litters[i])
		}
	}
	if len(k_12)%2 == 0 {
		return 0, er
	}

	lil, err := strconv.ParseFloat(k_12[0], 64) // lil=2
	if err != nil {
		return 0, er
	}

	for i := 1; i < len(k_12); i++ {
		if k_12[i] == "+" || k_12[i] == "-" {
			k1 = append(k1, lil)
			k2 = append(k2, k_12[i])
			lil, err = strconv.ParseFloat(k_12[i+1], 64)
			if err != nil {
				return 0, er
			}
		} else if k_12[i] == "*" {
			li, err := strconv.ParseFloat(k_12[i+1], 64)
			if err != nil {
				return 0, er
			}
			lil *= li
		} else if k_12[i] == "/" {
			li, err := strconv.ParseFloat(k_12[i+1], 64)
			if err != nil || li == 0 {
				return 0, er
			}
			lil /= li
		}
	}
	k1 = append(k1, lil)
	if len(k1)-1 != len(k2) || len(k1) == len(k2) {
		return 0, er
	}
	if len(k1) == 0 {
		return 0, er
	}
	var sc float64 = k1[0]
	k1 = k1[1:]
	for i := 0; i < len(k2); i++ {
		if k2[i] == "+" {
			sc += k1[i]
		} else {
			sc -= k1[i]
		}
	}
	return sc, nil
}
