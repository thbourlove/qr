package main

import (
	"flag"
	"os"

	"github.com/mdp/qrterminal"
)

var (
	text string
)

func main() {
	flag.Parse()
	text = flag.Arg(0)
	config := qrterminal.Config{
		Level:     qrterminal.M,
		Writer:    os.Stdout,
		BlackChar: qrterminal.BLACK,
		WhiteChar: qrterminal.WHITE,
	}
	qrterminal.GenerateWithConfig(text, config)
}
