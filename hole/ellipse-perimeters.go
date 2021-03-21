package hole

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
)

func perimeter(ai, bi int) (p float64) {
	var n float64
	a, b := float64(ai), float64(bi)
	h := math.Pow(a-b, 2) / math.Pow(a+b, 2)
	var bin float64
	for ni := 0; ni < 100; ni++ {
		n = float64(ni)
		bin = math.Gamma(1.5) / (math.Gamma(1.0+n) * math.Gamma(1.5-n))
		p += math.Pow(bin, 2) * math.Pow(h, n)
	}
	p *= math.Pi * (a + b)
	return
}
func ellipsePerimeters() (args []string, out string) {
	var outs []string

	// some random tests
	var a, b int
	var p float64
	var ps string
	for i := 0; i < 10; i++ {
		a = rand.Intn(15) + 5
		b = rand.Intn(5) + 1
		args = append(args, strconv.Itoa(a)+" "+strconv.Itoa(b))

		p = perimeter(a, b)
		ps = strconv.FormatFloat(p, 'f', 40, 64)[0:2]
		outs = append(outs, ps)
	}

	rand.Shuffle(len(args), func(i, j int) {
		args[i], args[j] = args[j], args[i]
		outs[i], outs[j] = outs[j], outs[i]
	})

	out = strings.Join(outs, "\n")
	return
}
