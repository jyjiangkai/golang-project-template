package utils

import (
	"golang-project-template/pkg/client"
	"math"
)

var MUL float64 = 1

func Sum(num int) int {
	sum := 0
	for i := 0; i < num; i++ {
		sum = sum + i
	}
	return sum
}

func Sqrt(num float64) float64 {
	num = MUL * num
	return math.Sqrt(float64(num))
}

func Show(c client.ClientInterface) string {
	return c.Get()
}

func Modify(c client.ClientInterface, name string) string {
	return c.Update(name)
}

func IsEqual(a, b string) bool {
	return a == b
}
