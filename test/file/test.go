package main

import (
	"fmt"
	"go/learning-notes/file"
	"time"
)

func main() {
	f := "~/learn/go-mefs-v2/mefs-user"

	start := time.Now()

	file.Read0(f)
	t0 := time.Now()
	fmt.Println("Cost time ", t0.Sub(start))

	file.Read1(f)
	t1 := time.Now()
	fmt.Println("Cost time ", t1.Sub(t0))

	file.Read2(f)
	t2 := time.Now()
	fmt.Println("Cost time ", t2.Sub(t1))

	file.Read3(f)
	t3 := time.Now()
	fmt.Println("Cost time ", t3.Sub(t2))
}
