package main

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestBinaryTree(t *testing.T) {
  assert := assert.New(t)

  tree:= Create(50, "head")
  assert.Equal(tree.size, 1, "tree improperly initialized")

  tree.Append(30, "first left child")
  assert.Equal(tree.head.lChild.index, 30, "append l malfunction")

  tree.Append(70, "first right child")
  assert.Equal(tree.head.rChild.index, 70, "append r malfunction")
}
