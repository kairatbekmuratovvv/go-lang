package main

import (
	"fmt"
	"strconv"
)

func main() {
	hello := len("Hello")
	me := len("Я")
	smile := len("😄")
	fmt.Printf("%T", hello)
	fmt.Println(me)
	fmt.Println(smile)

	r := "Я"

	fmt.Println(r)

	s := "ЯA"
	for i, r := range s {
		fmt.Printf("%d: %c\n", i, r)
	}
	fmt.Println(s[0])

	s1 := "cat"
	//s1[0] = 'b'
	fmt.Println(s1)

	s2 := "ЯA"
	letters := []rune(s2)
	letters[0] = 'Б'

	fmt.Println(s2)
	fmt.Println(string(letters))

	n := 87

	a := string(n)
	b := strconv.Itoa(n)
	fmt.Println(a)

	fmt.Println(len(a), len(b))

	n, err := strconv.Atoi("abs")
	// n == 42
	// err == nil
	fmt.Println(n, err)

	s4 := "A🙂Я"
	fmt.Println(s4[1:5])
	a1 := "\u042F"
	b1 := `\u042F`

	fmt.Println(len(a1), len(b1))
	var text string

	n1, err1 := strconv.Atoi(text)
	fmt.Println(n1, err1 == nil)
	_, err2 := strconv.Atoi("3000000000")
	fmt.Println(err2 == nil)
	n4 := int64(9_007_199_254_740_993)

	f := float64(n)
	restored := int64(f)

	fmt.Println(n4 == restored)

}
