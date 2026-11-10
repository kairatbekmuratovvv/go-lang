package main

import (
	"fmt"
	"strings"
)

type MenuList []string

func (ml MenuList) String() string {
	//c := "["
	//for i, menuItem := range ml {
	//	c = c + menuItem
	//	if i < len(ml)-1 {
	//		c = c + ", "
	//	}
	//}
	//c = c + "]"
	//
	//return c
	//return "[" + strings.Join(ml, ", ") + "]"
	return fmt.Sprintf("[%s]", strings.Join(ml, ", "))
}

func main() {
	menu := MenuList{"Coffee", "Tea", "Croissant"}
	fmt.Println("Menu:", menu)

}
