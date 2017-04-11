package main

import(
  "fmt"
)

type Node struct {
  parent *Node
  rChild *Node
  lChild *Node
  index int
  data string
}

type Tree struct {
  head *Node
  size int
}

func Create(index int, data string) Tree{
  n := Node{index: index, data: data}
  t:= Tree{&n, 1}
  return t
}

func (t *Tree) Append(index int, data string){
  t.size += 1
  n := Node{index: index, data: data}
  nextNode := t.head
  pNode := t.head.parent
  for{
    if n.index >= nextNode.index{
      pNode = nextNode
      nextNode = nextNode.rChild
      if nextNode == nil{
        n.parent = pNode
        pNode.rChild = &n
        return
      }
    } else if n.index < nextNode.index{
      pNode = nextNode
      nextNode = nextNode.lChild
      if nextNode == nil{
        n.parent = pNode
        pNode.lChild = &n
        return
      }
    }
  }
}

func main(){
  fmt.Println("pls test")
}
