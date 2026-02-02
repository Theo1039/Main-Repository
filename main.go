package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readAmount(prompt string) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println(prompt)
		amount, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error", err)
			continue
		}
		amount = strings.TrimSpace(amount)
		amountInt, err := strconv.Atoi(amount)
		if err != nil {
			fmt.Println("error", err)
			continue
		}
		return amountInt
	}
}

func deposit() int {
	depositAmount := readAmount("enter deposit amount")
	balance := depositAmount
	return balance
}
func withDrawal(balance int, amount int) int {
	return balance - amount
}

func checkBalance(b int) {
	fmt.Println("balance:", b)
}

func main() {

	amount := readAmount("enter deposit amount: ")
	balance := amount
	fmt.Println("initial balance:", balance)
	withDraw := readAmount("enter withdrawal amount: ")
	if withDraw > balance {
		fmt.Println("insufficient fund")
		return
	}
	balance = withDrawal(balance, withDraw)
	fmt.Println("current balance:", balance)
	fmt.Println("press 1, to check balance")
	press := readAmount("to check balance press 1: ")
	if press == 1 {
		checkBalance(balance)
	}

}
