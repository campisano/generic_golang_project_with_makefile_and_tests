package calc

import "strconv"

// Calc is a calculator implementation for test purposes
type Calc struct {
	result  float64
	history []string
}

func (c *Calc) GetResult() float64 {
	return c.result
}

func (c *Calc) GetHistory() []string {
	return c.history
}

func (c *Calc) historyPush(oper string, num float64) {
	strNum := strconv.FormatFloat(num, 'f', -1, 64)
	c.history = append(c.history, oper+" "+strNum)
}

func (c *Calc) Sum(num float64) {
	c.result += num
	c.historyPush("+", num)
}

func (c *Calc) Sub(num float64) {
	c.result -= num
	c.historyPush("-", num)
}

func (c *Calc) Mult(num float64) {
	c.result *= num
	c.historyPush("*", num)
}

func (c *Calc) Div(num float64) {
	c.result /= num
	c.historyPush("/", num)
}
