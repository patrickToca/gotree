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
	bi/		# Binary Tree (Simplified)
	bst/	# Binary Search Tree
	red/	# Red Black Tree
	btree/	# B-Tree

example/		# Example Code

viz/		# Tree Visualization (Graphviz)
	dot/	# Convert JSON tree data to DOT
```

Example
==========
```go

```

To-Do-List
==========
**Non-Committal on a Timeline**

- More Tree Data Structures
