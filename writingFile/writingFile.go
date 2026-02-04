package writingFile

import (
	"bufio"
	"fmt"
	"os"
)

func FileWriting() {
	text := "My name is Theophilus, I am learning how to write to file with package\n"

	file, err := os.Create("Theo.txt")
	if err != nil {
		fmt.Println("error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	if _, err := writer.WriteString(text); err != nil {
		fmt.Println("writing to file error:", err)
		return
	}
	writer.Flush()

	fmt.Println("writing to file successful")
}
