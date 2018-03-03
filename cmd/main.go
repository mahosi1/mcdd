package main

import (
	"mcdf"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	h := mcdf.NewHeader()
	f, err := os.Create("./data")
	check(err)
	defer f.Close()
	h.Write(f)
}
