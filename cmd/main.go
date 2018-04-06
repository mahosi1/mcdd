package main

import (
	"fmt"
	"github.com/mahosi1/mcdf"
	"os"
	"strings"

	"github.com/sakeven/RbTree"
)

func check(e error) {

}

func main() {
	s := "abc"
	a := s[0]
	i := strings.ToUpper(string(a))
	fmt.Println(i)
	b := make([]byte, 0, 1000)
	fmt.Println(b)
	cf := mcdf.NewCompoundFile()
	mystream := cf.RootStorage.AddStream("somestream")
	fmt.Println(mystream)

	fmt.Println(cf)
	h := mcdf.NewHeader()
	f, err := os.Create("./data")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	h.Write(f)
	t := rbtree.NewTree()
	fmt.Println(t)
}
