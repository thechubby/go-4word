package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Программа подсчитывает количество 4х значных слов. Введите строку")

	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	line := myscanner.Text()

	var linesplit = strings.Split(line, " ")

	for i, l := range linesplit {
		var k = string(l)
		if len(k) == 4 {
			fmt.Println(i, l)
		}
	}
}
