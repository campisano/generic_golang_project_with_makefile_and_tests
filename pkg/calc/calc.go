package calc

import "strconv"

// Calc is a calculator implementation for test purposes
type Calc struct {
	result float64
	history []string
}

// // NewCalc is the Calc constructor
// func NewCalc() *Calc {
// 	return &Calc {
// 		result: 0,
// 		history: make([]string, 0),
// 	}
// }

func (c *Calc) getResult() float64 {
	return c.result
}

func (c *Calc) getHistory() []string {
	return c.history
}

func (c *Calc) historyPush(oper string, num float64) {
	strNum := strconv.FormatFloat(num, 'f', -1, 64)
	c.history = append(c.history, oper + " " + strNum)
}

func (c *Calc) sum(num float64) {
	c.result += num
	c.historyPush("+", num)
}

func (c *Calc) sub(num float64) {
	c.result -= num
	c.historyPush("-", num)
}

func (c *Calc) mult(num float64) {
	c.result *= num
	c.historyPush("*", num)
}

func (c *Calc) div(num float64) {
	c.result /= num
	c.historyPush("/", num)
}
