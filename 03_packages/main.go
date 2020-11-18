package main

import (
	"fmt"
	"math"

	"github.com/tamudashe/go_basics/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(strutil.Reverse("tamudashe"))
}
