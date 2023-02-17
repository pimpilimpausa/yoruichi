package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func i_reader(o string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(o)

	i, _ := reader.ReadString('\n')
	return strings.TrimSuffix(i, "\n")
}

func f_reader(d string) string {
	fmt.Println(d)

	dat, err := os.ReadFile("/Users/jsu/Desktop/euspeji/jsu/sessionid.txt")
	if err != nil {
		panic(err)
	}
	return string(dat[:])
}
