package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

func calculate(s string) int {
    ans, stack := 0, []int{}
    op, cur := '+', 0
    for i := 0; i < len(s); i++ {
        if s[i] >= '0' && s[i] <= '9' { cur = cur * 10 + int(s[i] - '0') }
        if i + 1 == len(s) || (s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/') {
            if op == '+' {
                stack = append(stack, cur)
            } else if op == '-' {
                stack = append(stack, -cur)
            } else if op == '*' {
                pop := stack[len(stack) - 1]
                stack = stack[:len(stack) - 1]
                stack = append(stack, pop * cur)
            } else if op == '/' {
                pop := stack[len(stack) - 1]
                stack = stack[:len(stack) - 1]
                stack = append(stack, pop / cur)
            }
            op = rune(s[i])
            cur = 0
        }
    }
    for _, v := range stack { ans += v }
    return ans
}

func main() {

  reader := bufio.NewReader(os.Stdin)
  fmt.Println("Simple Shell")
  fmt.Println("---------------------")

  for {
    fmt.Print("-> ")
    text, _ := reader.ReadString('\n')
    // convert CRLF to LF
    text = strings.Replace(text, "\r\n", "", -1)

	// process text and get output and print
	// evaluate string and extract nums and signs
	// 
	fmt.Println(calculate(text))
    if strings.Compare("hi", text) == 0 {
      fmt.Println("invalid")
    }

  }

}