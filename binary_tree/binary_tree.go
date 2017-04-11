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
    if n.index > nextNode.index{
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
    } else {
      fmt.Println("index must be a unique identifier")
    }
  }
}

func (t *Tree) Get(index int) string{
  nextNode := t.head
  for{
    if index > nextNode.index{
      nextNode = nextNode.rChild
    } else if index < nextNode.index{
      nextNode = nextNode.lChild
    } else if index == nextNode.index {
      return nextNode.data
    } else if nextNode == nil {
      return "index not present in tree"
    }
  }
}

func (t *Tree) Set(index int, data string) string{
  nextNode := t.head
  for{
    if index > nextNode.index{
      nextNode = nextNode.rChild
    } else if index < nextNode.index{
      nextNode = nextNode.lChild
    } else if index == nextNode.index {
      nextNode.data = data
      return nextNode.data
    } else if nextNode == nil {
      return "index not present in tree"
    }
  }
}

func main(){
  fmt.Println("pls test")
}
