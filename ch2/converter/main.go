package main

import (
	"fmt"
	"gopl-exercises/ch2/converter/weightconv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		w, _ := strconv.ParseFloat(arg, 64)
		p := weightconv.Pound(w)
		k := weightconv.Kilogram(w)
		fmt.Printf("%s = %s, %s = %s\n",
			p, weightconv.PToK(p), k, weightconv.KToP(k))
	}
}
