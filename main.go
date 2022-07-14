package main

import (
	"fmt"
	"math"
)

func main() {

	/*
	for i := 0; i <= 1000000; i++ {
		fmt.Println(i, spell(i))
	}
	*/

	i := 767
	fmt.Println(i, spell(i))

}

func pow(i int, p int) int {
	return int(math.Pow(1000, float64(p)))
}

func spell(n int) string {
	to19 := []string{
		"як",
		"ду",
		"се",
		"чор",
		"панч",
		"шаш",
		"хафт",
		"хашт",
		"нух",
		"дах",
		"езда",
		"дувозда",
		"сензда",
		"чорда",
		"понзда",
		"шонзда",
		"хабда",
		"хаштда",
		"нузда",
	}

	tens := []string{
		"бист",
		"си",
		"чил",
		"панчо",
		"шаст",
		"хафтод",
		"хаштод",
		"навад",
	}



	if n == 0 {
		return ""
	}

	if n < 20 {
		return to19[n - 1]
	}

	if n < 100 {
		if n % 10 == 0 {
			return tens[n / 10 - 2] + " " + spell(n % 10)
		}
		return tens[n / 10 - 2] + "у " + spell(n % 10)
	}

	if n < 1000 {
		if n % 100 == 0 {
			return to19[n / 100 - 1] + "сад " + spell(n % 100)
		}
		return to19[n / 100 - 1] + "саду " + spell(n % 100)
	}

	for idx, w := range []string{"хазор", "миллион", "миллиард"} {
		p := idx + 1
		if n < pow(1000, p + 1) {
			if n % pow(1000, p) == 0 {
				return spell(n / pow(1000, p)) + " " + w + " " + spell(n % pow(1000, p))
			}
			return spell(n / pow(1000, p)) + " " + w + "у " + spell(n % pow(1000, p))
		}
	}

	return "error"
}
