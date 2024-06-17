package main

func main() {

	server := NewAPIServer(":4001")

	server.Run()
}
