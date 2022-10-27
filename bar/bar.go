package main

import (
	"fmt"

	"github.com/li-go/go-workspace-test/hoge"

	"golang.org/x/example/stringutil"
)

func Bar() {
	fmt.Println(stringutil.Reverse("Hello Gopher! 中华人民共和国"))
	hoge.Hoge()
}
