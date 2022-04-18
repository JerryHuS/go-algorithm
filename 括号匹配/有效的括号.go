/**
 * @Author: alessonhu
 * @Description:
 * @File:  有效的括号
 * @Version: 1.0.0
 * @Date: 2022/4/18 16:51
 */

package main

import (
	"container/list"
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(isValid("()"))
	a := '('
	fmt.Println(reflect.TypeOf(a))
	b := "()"[0]
	fmt.Println(reflect.TypeOf(b))
}

func isValid(s string) bool {
	stack := list.New()
	for i := range s {
		switch s[i] {
		case '(', '{', '[':
			stack.PushBack(s[i])
		case ')':
			if stack.Len() == 0 {
				return false
			}
			res := stack.Remove(stack.Back())
			if res != uint8('(') {
				return false
			}
		case '}':
			if stack.Len() == 0 {
				return false
			}
			res := stack.Remove(stack.Back())
			if res != uint8('{') {
				return false
			}
		case ']':
			if stack.Len() == 0 {
				return false
			}
			res := stack.Remove(stack.Back())
			if res != uint8('[') {
				return false
			}
		}
	}
	return stack.Len() == 0
}
