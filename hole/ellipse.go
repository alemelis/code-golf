package hole

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
)

func perimeter(ai, bi int) (p float64) {
	a, b := float64(ai), float64(bi)
	h := math.Pow(a-b, 2) / math.Pow(a+b, 2)
	p = math.Pi * (a + b) * (1.0 + (3.0 * h / (10.0 + math.Sqrt(4.0-(3.0*h)))))
	return
}

func ellipse() (args []string, out string) {
	var outs []string

	// some random tests
	var a, b int
	for i := 0; i < 50; i++ {
		a = rand.Intn(50) + 1
		b = rand.Intn(50) + 1

		args = append(args, strconv.Itoa(a)+" "+strconv.Itoa(b))
		outs = append(outs, strconv.FormatFloat(perimeter(a, b), 'f', 40, 64))
	}

	rand.Shuffle(len(args), func(i, j int) {
		args[i], args[j] = args[j], args[i]
		outs[i], outs[j] = outs[j], outs[i]
	})

	out = strings.Join(outs, "\n")
	return
}
