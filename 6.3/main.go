package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	_, _ = fmt.Scan(&str)
	str = strings.ReplaceAll(str, "--", "2")
	str = strings.ReplaceAll(str, "-.", "1")
	str = strings.ReplaceAll(str, ".", "0")
	fmt.Println(str)
}
