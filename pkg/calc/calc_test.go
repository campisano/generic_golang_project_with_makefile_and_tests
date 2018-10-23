package calc

import "testing"
import "github.com/stretchr/testify/assert"

func TestCalc_add(t *testing.T) {
	calc := &Calc{}

	calc.Sum(1)

	assert.Equal(t, 1.0, calc.GetResult())
	assert.Equal(t, 1, len(calc.GetHistory()))
	assert.Equal(t, "+ 1", calc.GetHistory()[0])
}

func TestCalc_sub(t *testing.T) {
	calc := &Calc{}

	calc.Sub(1)

	assert.Equal(t, -1.0, calc.GetResult())
	assert.Equal(t, 1, len(calc.GetHistory()))
	assert.Equal(t, "- 1", calc.GetHistory()[0])
}

func TestCalc_mult(t *testing.T) {
	calc := &Calc{}

	calc.Mult(1)

	assert.Equal(t, 0.0, calc.GetResult())
	assert.Equal(t, 1, len(calc.GetHistory()))
	assert.Equal(t, "* 1", calc.GetHistory()[0])
}

func TestCalc_div(t *testing.T) {
	calc := &Calc{}

	calc.Div(1)

	assert.Equal(t, 0.0, calc.GetResult())
	assert.Equal(t, 1, len(calc.GetHistory()))
	assert.Equal(t, "/ 1", calc.GetHistory()[0])
}
