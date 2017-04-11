package main

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
  assert := assert.New(t)

  ll := Create("heady")

  assert.Equal(ll.head.data, "heady", "failed to initialize ll")

  ll.Append("second")

  assert.Equal(ll.head.next.data, "second", "failed to append")

  ll.Append("third")

  ll.Append("fourth")

  ll.Insert(2, "inserted third")

  assert.Equal(ll.head.next.next.data, "inserted third", "failed to insert")

  assert.Equal(ll.head.next.next.next.index, 3, "failed to update index")

  assert.Equal(ll.tail.index, 4, "failed to update index")

  ll.Remove(2)

  assert.Equal(ll.head.next.next.data, "third", "failed to remove")

  assert.Equal(ll.tail.index, 3, "failed to update index")

  ll.Set(2, "suck it")

  assert.Equal(ll.head.next.next.data, "suck it", "failed to set")

  assert.Equal(ll.head.next.next.data, ll.Get(2), "failed to get")

  ll.Remove(3)

  assert.Equal(ll.tail.index, 2, "failed to remove tail")

  ll.Remove(0)

  assert.Equal(ll.tail.index, 1, "failed to remove head")

}
