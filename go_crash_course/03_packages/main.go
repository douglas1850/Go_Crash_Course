package main

import (
	"fmt"
	"math"

	"github.com/douglas1850/go_crash_course/03_packages/strutil"
) //unused packages go away when you save. Can import multiple packages, seperate on new lines

func main() {
	fmt.Println(math.Floor(2.3))
	fmt.Println(math.Ceil(2.3))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse("Hello"))
}
