package game_logic

import (
	"cocogame-max/chickenroad2_srv/pb_chickenroad2"
	"github.com/shopspring/decimal"

	"strconv"
)

var (
	GameConfig       []*pb_chickenroad2.GameConfig
	CoefficientsEasy = []float64{
		1.01, 1.03, 1.06, 1.10, 1.15, 1.19, 1.24, 1.30, 1.35, 1.42, 1.48, 1.56, 1.65, 1.75, 1.85, 1.98, 2.12, 2.28, 2.47, 2.70, 2.96, 3.28, 3.70, 4.11, 4.64, 5.39, 6.50, 8.36, 12.08, 23.24,
	}
	CoefficientsMedium = []float64{
		1.08, 1.21, 1.37, 1.56, 1.78, 2.05, 2.37, 2.77, 3.24, 3.85, 4.62, 5.61, 6.91, 8.64, 10.99, 14.29, 18.96, 26.07, 37.24, 53.82, 82.36, 137.59, 265.35, 638.82, 2457.00,
	}
	CoefficientsHard = []float64{
		1.18, 1.46, 1.83, 2.31, 2.95, 3.82, 5.02, 6.66, 9.04, 12.52, 17.74, 25.80, 38.71, 60.21, 97.34, 166.87, 305.94, 595.86, 1283.03, 3267.64, 10898.54, 62162.09,
	}
	CoefficientsHardDaredevil = []float64{
		1.44, 2.21, 3.45, 5.53, 9.09, 15.30, 26.78, 48.70, 92.54, 185.08, 391.25, 894.28, 2235.72, 6096.15, 18960.33, 72432.75, 379632.82, 3608855.25,
	}
	MinBet  = decimal.NewFromFloat(0.01)
	MaxBet  = decimal.NewFromInt32(200)
	CoffMap = map[string][]float64{
		"EASY":      CoefficientsEasy,
		"MEDIUM":    CoefficientsMedium,
		"HARD":      CoefficientsHard,
		"DAREDEVIL": CoefficientsHardDaredevil,
	}
)

func toStr(f []float64) (s []string) {
	s = make([]string, len(f))
	for i, v := range f {
		s[i] = strconv.FormatFloat(v, 'f', -1, 64)
	}
	return
}

func init() {
	GameConfig = []*pb_chickenroad2.GameConfig{
		{
			Coefficients: &pb_chickenroad2.Coefficients{
				EASY:      toStr(CoefficientsEasy),
				MEDIUM:    toStr(CoefficientsMedium),
				HARD:      toStr(CoefficientsHard),
				DAREDEVIL: toStr(CoefficientsHardDaredevil),
			},
			LastWin: nil,
		},
	}
}
