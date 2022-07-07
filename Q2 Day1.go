package main

import "fmt"

type node struct {
	val   string
	left  *node
	right *node
}

func buildtree(s string) *node {
	if len(s) == 0 {
		return nil
	}
	if len(s) == 1 {
		var head node
		head.val = s[0:1]
		return &head
	} else {
		var head, l node
		head.left = &l
		l.val = s[0:1]
		head.val = s[1:2]
		head.right = buildtree(s[2:])
		return &head
	}
}

func printpreorder(head *node, res *string) {
	*res += head.val
	if head.left != nil {
		printpreorder(head.left, res)
	}
	if head.right != nil {
		printpreorder(head.right, res)
	}
}

func printpostorder(head *node, res *string) {
	if head == nil {
		return
	} else {
		printpostorder(head.left, res)
		printpostorder(head.right, res)
		*res += head.val
	}
}

func main() {
	var s, pre, post string
	var head *node
	s = "a+b-c"
	head = buildtree(s)
	buildtree(s)
	printpreorder(head, &pre)
	printpostorder(head, &post)
	fmt.Println(pre)
	fmt.Println(post)
}
