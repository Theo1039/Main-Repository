package appendFile

import (
	"fmt"
	"os"
)

func FileAppend() {
	file, err := os.OpenFile("Theo.Txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("error opening file:", err)
		return
	}
	defer file.Close()
	_, err = file.Write([]byte("this is an appended file: Hello users!"))
	if err != nil {
		fmt.Println("error writing to file:", err)
		return
	}
	fmt.Println("file appended successfully")

}
