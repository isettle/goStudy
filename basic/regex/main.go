package main

import (
	"fmt"
	"regexp"
)

const text = `
my e-mail1 is shilei_zhang@qq.com. 
your email2 is 690613233@foxmail.com.cn
`

func main() {
	reg := regexp.MustCompile(`(\w+)@\w+\.\w+[\.\w]*`)
	match := reg.FindAllStringSubmatch(text, -1)
	fmt.Println(match)
}
