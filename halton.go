package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/barnex/bruteray/random"
)

var (
	a    = flag.Int("a", 2, "a")
	b    = flag.Int("b", 3, "b")
	s    = flag.Int("s", 1, "s")
	i    = flag.Int("i", 0, "i")
	n    = flag.Int("n", 300, "n")
	disk = flag.Bool("disk", false, "disk mapping")
	rnd  = flag.Bool("rand", false, "random instead")
)

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	i0 := *i
	max := *s * (i0 + *n)
	for i := i0; i < max; i += *s {
		u := random.Halton(*a, i)
		v := random.Halton(*b, i)
		if *rnd {
			u = rand.Float64()
			v = rand.Float64()
		}
		if *disk {
			u, v = random.UniformDisk(u, v)
			u = u/2 + 0.5
			v = v/2 + 0.5
		}
		fmt.Println(u, "\t", v)
	}

}
