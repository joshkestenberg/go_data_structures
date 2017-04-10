package main

import(
  "fmt"
)

type Node struct {
  index int
  data string
  next *Node
  prev *Node
}

type LinkList struct {
  head *Node
  tail *Node
}

func Create(data string) *LinkList {
  ll := new(LinkList)
  node := new(Node)
  node.index = 0
  node.data = data
  ll.head = node
  ll.tail = node
  return ll
}

func (ll *LinkList) Append(data string) {
  index := ll.tail.index + 1
  newNode := Node{index: index, data: data, prev: ll.tail}
  ll.tail.next = &newNode
  ll.tail = &newNode
}

func (ll *LinkList) Insert(index int, data string) {
  newNode := Node{index: index, data: data}
  nextNode := ll.head
  prevNode := nextNode.prev
  if index > ll.tail.index{
    newNode.prev = ll.tail
    ll.tail.next = &newNode
    ll.tail = &newNode
  } else if index == 0 {
    newNode.next = nextNode
    nextNode.prev = &newNode
    nextNode.index = 1
    nextNode = nextNode.next
    for{
      if nextNode != nil{
        nextNode.index += 1
      } else {
        return
      }
      nextNode = nextNode.next
    }
  } else {
    for {
      if nextNode.index == index{
        newNode.prev = prevNode
        newNode.next = nextNode
        nextNode.prev = &newNode
        prevNode.next = &newNode
        for{
          if nextNode != nil{
            nextNode.index += 1
          } else {
            return
          }
          nextNode = nextNode.next
        }
      } else {
        nextNode = nextNode.next
        prevNode = nextNode.prev
      }
    }
  }
}

func (ll *LinkList) Remove(index int){
  nextNode := ll.head

  if index == ll.head.index {
    nextNode = nextNode.next
    ll.head = nextNode
    for {
      if nextNode != nil{
        nextNode.index -= 1
      } else {
        return
      }
      nextNode = nextNode.next
    }
  } else if index == ll.tail.index {
    ll.tail = ll.tail.prev
    ll.tail.next = nil
  } else {
    for {
      nextNode = nextNode.next
      if nextNode.index == index {
        nextNode.prev.next = nextNode.next
        nextNode.next.prev = nextNode.prev
        nextNode = nextNode.next
        for {
          if nextNode != nil{
            nextNode.index -= 1
          } else {
            return
          }
          nextNode = nextNode.next
        }
        nextNode = nextNode.next
      }
    }
  }
}

func (ll *LinkList) Get(index int) string {
  nextNode := ll.head
  if index > ll.tail.index {
    return "invalid entry"
  }
  for {
    if index == nextNode.index{
      return nextNode.data
    } else {
      nextNode = nextNode.next
    }
  }
}

func (ll *LinkList) Set(index int, data string) string {
  nextNode := ll.head
  if index > ll.tail.index {
    return "invalid entry"
  }
  for {
    if index == nextNode.index{
      nextNode.data = data
      return data
    } else {
      nextNode = nextNode.next
    }
  }
}

func main(){

   ll := Create("heady")

   ll.Append("second")

   ll.Append("third")

   ll.Append("fourth")

   ll.Insert(2, "inserted third")

   ll.Remove(2)

  //  ll.Set(2, "suck it")

   data := ll.Get(2)
   fmt.Println(data)

}
