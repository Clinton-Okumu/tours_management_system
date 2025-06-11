package main

import "flag"

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Go backend server")
}
