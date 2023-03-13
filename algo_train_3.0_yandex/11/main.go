package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	stack := NewStack()
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _, err := reader.ReadLine()
		if err != nil {
			panic(err)
		}
		inputArr := strings.Split(string(input), " ")

		switch inputArr[0] {
		case "push":
			n, _ := strconv.Atoi(inputArr[1])
			fmt.Println(stack.Push(n))
		case "pop":
			num, err := stack.Pop()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(num)
		case "size":
			fmt.Println(stack.Size())
		case "clear":
			fmt.Println(stack.Clear())
		case "exit":
			fmt.Println("bye")
			break
		case "back":
			num, err := stack.Back()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(num)
		}
	}
}
