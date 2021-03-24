package hole

import (
	"fmt"
	"gonum.org/v1/gonum/mathext"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

func perimeter(ai, bi int) (p float64) {
	a, b := float64(ai), float64(bi)
	if a < b {
		a, b = b, a
	}
	return 4 * a * mathext.CompleteE(1-math.Pow(b, 2)/math.Pow(a, 2))
}

func getTest(rlim int) (a, b int, p string) {
	a = rand.Intn(rlim) + 1
	b = rand.Intn(rlim) + 1
	p = fmt.Sprintf("%.10f", perimeter(a, b))
	return
}

func ellipsePerimeters() (args []string, out string) {
	var outs []string

	// some random tests
	var a, b, rlim int
	var p string

	for _, n := range []float64{1, 2, 3} {
		for i := 0; i < 3; i++ {
			rlim = int(math.Pow(10.0, n))
			a, b, p = getTest(rlim)
			args = append(args, strconv.Itoa(a)+" "+strconv.Itoa(b))
			outs = append(outs, p)
		}
	}

	rand.Shuffle(len(args), func(i, j int) {
		args[i], args[j] = args[j], args[i]
		outs[i], outs[j] = outs[j], outs[i]
	})

	out = strings.Join(outs, "\n")
	return
}
