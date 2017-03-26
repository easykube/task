package main

import (
	"os"
	"strings"

	"github.com/flosch/pongo2"
)

func main() {

	vars := os.Environ()
	for _, s := range vars {
		println(s)
		list := strings.Split(s, "=")
		println("===")
		println(list[0])
		println("=")
		println(list[1])
	}
	tpl, err := pongo2.FromString("cmd /k echo hello,{{name}}")
	if err != nil {
		panic(err)
	}
	out, err := tpl.Execute(pongo2.Context{"name": "tom"})
	if err != nil {
		panic(err)
	}
	println(out)
	println("======")

}
