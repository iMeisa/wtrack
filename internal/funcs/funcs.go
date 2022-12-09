package funcs

import (
	"fmt"
	"github.com/iMeisa/weed/internal/tools"
	"html/template"
)

var Functions = template.FuncMap{
	"add":       add,
	"contains":  tools.Contains,
	"format":    format,
	"makeRange": makeRange,
	"subtract":  subtract,
}

func add(num1, num2 int) int {
	return num1 + num2
}

func format(f float32) string {
	return fmt.Sprintf("%.3f", f)
}

func makeRange(count int) []int {
	var seq []int

	for i := 0; i < count; i++ {
		seq = append(seq, i)
	}

	return seq
}

func subtract(num1, num2 int) int {
	return num1 - num2
}
