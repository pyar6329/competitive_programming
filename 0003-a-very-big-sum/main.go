package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	stdin.Scan()
	text := stdin.Text()
	numbers := strings.Split(text, " ")
	var totalNum int64
	for _, str := range numbers {
		i, _ := strconv.Atoi(str)
		totalNum += int64(i)
	}
	fmt.Println(totalNum)
}
