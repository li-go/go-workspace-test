package main

import (
	"fmt"

	"foo/hoge"
	"golang.org/x/example/stringutil"
)

func Bar() {
	fmt.Println(stringutil.Reverse("Hello Gopher! 中华人民共和国"))
	hoge.Hoge()
}
