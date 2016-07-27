package main

import (
	"flag"
	"os"

	"github.com/WuHan0608/micadvisor_open/g"
)

var CadvisorPort = "18080"

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	flag.Parse()

	g.ParseConfig(*cfg)

	LogRun("sys start")

	pushData()
	iAmAlive()
}

func iAmAlive() {
	f, _ := os.OpenFile("test.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0660)
	defer f.Close()

	f.Write([]byte("b"))
}
