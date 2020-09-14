package Abstract

import (
	"fmt"
)

func main() {
	p := Newperson("Tom")
	p.SetAge(10)
	p.SetSalary(2000)
	fmt.Println(*p)
}
