package main

import (
	"flag"
	server "goblockchain/server"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	port := flag.Uint("port", 5001, "TCP Port Number for Blockchain Server")
	flag.Parse()
	app := server.NewServer(uint16(*port))
	app.Run()
}
