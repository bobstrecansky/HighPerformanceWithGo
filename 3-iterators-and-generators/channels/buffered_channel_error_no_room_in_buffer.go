package main

func main() {

ch := make(chan string, 1)
ch <- "foo"
ch <- "bar"
}
