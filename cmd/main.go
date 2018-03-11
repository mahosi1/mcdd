package main

import (
	"fmt"
	"mcdf"
	"os"
	"strings"

	"github.com/sakeven/RbTree"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	s := "abc"
	a := s[0]
	i := strings.ToUpper(string(a))
	fmt.Println(i)
	b := make([]byte, 0, 1000)
	fmt.Println(b)
	cf := mcdf.NewCompoundFile()
	fmt.Println(cf)
	h := mcdf.NewHeader()
	f, err := os.Create("./data")
	check(err)
	defer f.Close()
	h.Write(f)
	t := rbtree.NewTree()
	fmt.Println(t)
}
