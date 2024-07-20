package main

import "fmt"

type stack []byte

func isValid(s string) bool {
	m := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	if len(s)%2 != 0 {
		return false
	}
	st := make(stack, 0)
	for i := 0; i <= len(s)-1; i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			st.Push(s[i])
		}
		if s[i] == ')' || s[i] == '}' || s[i] == ']' {

			p, _ := st.Pop()
			if p != m[s[i]] {
				return false
			}
		}
	}
	return len(st) == 0

}
func (s *stack) Push(v byte) {
	*s = append(*s, v)
}

func (s *stack) Pop() (byte, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top, true
}

func main() {
	fmt.Println(isValid("(){}[]")) // true
	fmt.Println(isValid("([{}])")) // true
	fmt.Println(isValid("({[}])")) // false

}