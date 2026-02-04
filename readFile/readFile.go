package readFile

import (
	"bufio"
	"fmt"
	"os"
)

func FileReader() {
	file, err := os.Open("Theo.Txt")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
