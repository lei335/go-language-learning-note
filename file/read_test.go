package file

import (
	"fmt"
	"time"
)

func test() {
	file := "~/learn/go-mefs-v2/mefs-user"

	start := time.Now()

	read0(file)
	t0 := time.Now()
	fmt.Println("Cost time ", t0.Sub(start))

	read1(file)
	t1 := time.Now()
	fmt.Println("Cost time ", t1.Sub(t0))

	read2(file)
	t2 := time.Now()
	fmt.Println("Cost time ", t2.Sub(t1))

	read3(file)
	t3 := time.Now()
	fmt.Println("Cost time ", t3.Sub(t2))
}
