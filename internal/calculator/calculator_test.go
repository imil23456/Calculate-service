package calculator

import "testing"

func TestCalc(t *testing.T) {
	testCases := []struct {
		expression string
		expected   float64
	}{
		{"1+1", 2},
		{"3+3*6", 21},
		{"1+8/2*4", 17},
		{"(1+1) *2", 4},
		{"((1+4) * (1+2) +10) *4", 84},
		{"(4+3+2)/(1+2) * 10 / 3", 10},
		{"(70/7) * 10 /((3+2) * (3+7)) -2", -1},
		{"-1", -1},
		{"10*(10/10*-10)", -100},
	}

	for _, tc := range testCases {
		result, err := Calc(tc.expression)
		if err != nil {
			t.Errorf("Error evaluating %s: %v", tc.expression, err)
		}
		if result != tc.expected {
			t.Errorf("For %s, expected %v but got %v", tc.expression, tc.expected, result)
		}
	}
}
