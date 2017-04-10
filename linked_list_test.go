package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryTree(t *testing.T) {
	fmt.Println("this is a tree")
}

func TestLinkedList(t *testing.T) {
	assert := assert.New(t)

	n3 := node{val: 2, data: "i'm third"}
	n2 := node{val: 1, data: "i'm second", next: &n3}
	root := node{val: 0, data: "i'm root", next: &n2}

	appendSequential(&root, "i'm fourth")
	insert(&root, 5, "i'm sixth")
	insert(&root, 4, "i'm fifth")

	assert.Equal(n3.next.val, 2, "n3.next.val is wrong")

	fmt.Println(n3.next.val, n3.next.data)
	fmt.Println(n3.next.next.val, n3.next.next.data)
	fmt.Println(n3.next.next.next.val, n3.next.next.next.data)

	fmt.Println("----------------")

	n3.next.removeNext()
	fmt.Println(n3)
	fmt.Println(n3.next)
	fmt.Println(n3.next.next)

	fmt.Println("----------------")

	root.removeN(2)
	fmt.Println(root.next)
	fmt.Println(root.next.next)

}
