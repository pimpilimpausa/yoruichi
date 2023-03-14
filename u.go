package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
	path, _ := os.Executable()
	path = regexp.MustCompile(`\w+$`).ReplaceAllString(path, "")

	dat, err := os.ReadFile(path + "sessionid.txt")
	if err != nil {
		panic(err)
	}
	return string(dat[:])
}
