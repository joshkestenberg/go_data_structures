package main

import(
  "fmt"
)

type node struct {
  val int
  data string
  next *node
}

func appendSequential(n *node, data string) {
  for {
    newNode := node{val: n.val+1, data: data}
    if n.next == nil{
      n.next = &newNode
      return
    }
    n = n.next
  }
}

func insert(n *node, val int, data string) {
  newNode := node{val: val, data: data}
  nextNode := n.next
  for {
    if nextNode == nil{
      n.next = &newNode
      return
    } else if newNode.val > n.val && newNode.val < n.next.val {
      newNode.next = nextNode
      n.next = &newNode
      return
    } else if newNode.val > n.val && newNode.val == n.next.val {
      fmt.Println("The node value you've input already exists")
      return
    }
    n = n.next
    nextNode = nextNode.next
  }
}

func (n *node) remove() {
  if n.next == nil {
    n = nil
  } else {
    n = n.next
  }
}

func main(){
  n3 := node{val: 2, data: "i'm third"}
  n2 := node{val: 1, data: "i'm second", next: &n3}
  root := node{val: 0, data: "i'm root",  next: &n2}

  appendSequential(&root, "i'm fourth")
  insert(&root, 5, "i'm sixth")
  insert(&root, 4, "i'm fifth")

  fmt.Println(n3.next.val, n3.next.data)
  fmt.Println(n3.next.next.val, n3.next.next.data)
  fmt.Println(n3.next.next.next.val, n3.next.next.next.data)

  n3.next.remove()
  fmt.Println(n3.next)
}
