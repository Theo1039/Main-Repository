package stdInput

import (
	"bufio"
	"fmt"
	"strings"
)

func GetInput(reader *bufio.Reader, prompt string) string {
	fmt.Println(prompt)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Input Error:", err)
	}
	return strings.TrimSpace(text)
}