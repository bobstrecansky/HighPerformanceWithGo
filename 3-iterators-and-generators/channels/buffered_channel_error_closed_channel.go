package main

func main() {
	ch := make(chan string, 1)
	close(ch)
	ch <- "foo"
}
