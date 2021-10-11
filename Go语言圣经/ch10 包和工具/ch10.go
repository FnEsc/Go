package main

/*
构建包：

方法1
$ cd $GOPATH/src/gopl.io/ch1/helloworld
$ go build
方法2
$ cd anywhere
$ go build gopl.io/ch1/helloworld
方法3
$ cd $GOPATH
$ go build ./src/gopl.io/ch1/helloworld
不能这样
$ cd $GOPATH
$ go build src/gopl.io/ch1/helloworld
Error: cannot find package "src/gopl.io/ch1/helloworld".

*/
