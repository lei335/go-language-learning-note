package slice

import "fmt"

func modifySlice(s []string) []string {
	i := 0
	for _, e := range s {
		if e != "" {
			s[i] = e // 同时更改了引用数组
			i++
		}
	}
	return s[:i]
}

func TestSlice() {
	nums := [5]string{"one", "", "three", "", "five"}
	s := nums[:3]
	r := modifySlice(s)

	fmt.Println("slice s: ", s)
	fmt.Println("slice r: ", r)
	fmt.Println("nums: ", nums)
}
