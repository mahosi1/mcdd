package main

import (
	"fmt"
	"mcdf"
	"os"

	"github.com/sakeven/RbTree"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	t := rbtree.NewTree()
	fmt.Println(t)
	cf := mcdf.NewCompoundFile()
	fmt.Println(cf)
	h := mcdf.NewHeader()
	f, err := os.Create("./data")
	check(err)
	defer f.Close()
	h.Write(f)
}
