gotree [![Build Status](https://travis-ci.org/gyuho/gotree.png?branch=master)](https://travis-ci.org/gyuho/gotree) [![GoDoc](https://godoc.org/github.com/gyuho/gotree?status.png)](http://godoc.org/github.com/gyuho/gotree) [![Project Stats](http://www.ohloh.net/p/714469/widgets/project_thin_badge.gif)](http://www.ohloh.net/p/714469)
==========

gotree provides tree visualizing tools and algorithm implementations.

Getting Started
==========
- [godoc.org](http://godoc.org/github.com/gyuho/gotree)
- [gowalker.org](http://gowalker.org/github.com/gyuho/gotree#_index)

```go
// to install, in the command line
mkdir $HOME/go
export GOPATH=$HOME/go
go get github.com/gyuho/gotree

// to include, in the code
import "github.com/gyuho/gotree"

// to call the function, in the code
[package_name].[function]

// to execute
go install
// or
go build
```


Package Hierarchy
==========
```go
tree/		# Tree Data Structure
	bst/	# Binary Search Tree
	rbt/	# Red Black Tree
	bt/		# B-Tree
	avl/	# AVL Tree

example/	# Example Code

viz/		# Tree Visualization (Graphviz)
```

Example (Binary Search Tree)
==========
```go
func Test_Show1(test *testing.T) {
	tr := bst.NewTree(5)
	for i := 0; i < 10; i++ {
		if i != 5 {
			tr = tr.Insert(int64(i))
		}
	}
	slice := []string{}
	Scan(tr, &slice)
	fmt.Println(slice)
	Show(tr, "tree1.dot")
}

func Test_Show2(test *testing.T) {
	tr := bst.NewTree(5)
	tr.Insert(int64(7))
	tr.Insert(int64(8))
	tr.Insert(int64(5))
	tr.Insert(int64(4))
	tr.Insert(int64(2))
	tr.Insert(int64(1))
	tr.Insert(int64(6))
	tr.Insert(int64(3))
	slice := []string{}
	Scan(tr, &slice)
	fmt.Println(slice)
	Show(tr, "tree2.dot")
}
```

<img src="./files/tree01.png" alt="tree01" width="350px" height="450px"/>

<img src="./files/tree02.png" alt="tree02" width="350px" height="450px"/>


To-Do-List
==========
**Non-Committal on a Timeline**

- More Tree Data Structures
