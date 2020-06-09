package main

func main() {
	go AuthServer()
	go CollectorServer()
	go RankerServer()
	go TokenServer()
}
