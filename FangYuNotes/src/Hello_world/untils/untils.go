package untils

import "fmt"

func Ppp(n int) {
	if n > 2 {
		fmt.Println("n = ", n)
		n--
		Ppp(n)
	} else {
		fmt.Println("final n = ", n)
	}

}
