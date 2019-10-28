package media

import (
	"fmt"
	"math/rand"
)

func ExampleRandExpInterval() {
	rng := rand.New(rand.NewSource(1))

	fmt.Println(randExpInterval(rng.Float64(), .1, 2))
	fmt.Println(randExpInterval(rng.Float64(), .1, 2))
	fmt.Println(randExpInterval(rng.Float64(), .1, 2))
	fmt.Println(randExpInterval(rng.Float64(), .1, 2))

}
