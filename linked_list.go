package main

import(
  "fmt"
)

type node struct {
  val int
  data string
  next *node
}

func insert(n *node, data string) {
  for {
    newNode := node{val: n.val+1, data: data}
    if n.next == nil{
      n.next = &newNode
      return
    }
    n = n.next
  }
}

func main(){
  n3 := node{val: 2, data: "i'm third"}
  n2 := node{val: 1, data: "i'm second", next: &n3}
  root := node{val: 0, data: "i'm root",  next: &n2}

  insert(&root, "i'm fourth")

  fmt.Println(root.next.next.next.val, root.next.next.next.data)
}
