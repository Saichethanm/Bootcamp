package main

import "fmt"

func main() {

	s := make([]string, 3)
	fmt.Println(s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)
	fmt.Println(s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "f")

	s = append(s, "g", "h", "i")

	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)

	l := c[2:5]
	fmt.Println(l)

	l = c[:5]
	fmt.Println(l)

	l = c[1:]
	fmt.Println(l)

	t := []string{"g", "h", "i"}
	fmt.Println(t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, 1)
		twoD[i][0] = i + innerLen
	}

	fmt.Println(twoD)
}
