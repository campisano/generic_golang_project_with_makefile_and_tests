package calc

import "testing"
import "github.com/stretchr/testify/assert"

func TestCalc_add(t *testing.T) {
	calc := &Calc{}

	calc.sum(1)

	assert.Equal(t, 1.0, calc.getResult())
	assert.Equal(t, 1, len(calc.getHistory()))
	assert.Equal(t, "+ 1", calc.getHistory()[0])
}

func TestCalc_sub(t *testing.T) {
	calc := &Calc{}

	calc.sub(1)

	assert.Equal(t, -1.0, calc.getResult())
	assert.Equal(t, 1, len(calc.getHistory()))
	assert.Equal(t, "- 1", calc.getHistory()[0])
}

func TestCalc_mult(t *testing.T) {
	calc := &Calc{}

	calc.mult(1)

	assert.Equal(t, 0.0, calc.getResult())
	assert.Equal(t, 1, len(calc.getHistory()))
	assert.Equal(t, "* 1", calc.getHistory()[0])
}

func TestCalc_div(t *testing.T) {
	calc := &Calc{}

	calc.div(1)

	assert.Equal(t, 0.0, calc.getResult())
	assert.Equal(t, 1, len(calc.getHistory()))
	assert.Equal(t, "/ 1", calc.getHistory()[0])
}
