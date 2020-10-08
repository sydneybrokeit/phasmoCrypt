package main

import (
	"flag"
	"github.com/hmschreck/phasmoCrypt/lib"
	"io/ioutil"
)

var fileNameInput = flag.String("input", "saveData.txt", "filename for input to decrypt")
var fileNameOutput = flag.String("output", "saveDataDec.txt", "filename for output after decrypting")

func main() {
	flag.Parse()
	inputFile, err := ioutil.ReadFile(*fileNameInput)
	if err != nil {
		panic("couldn't read file input")
	}
	decryptedFile := phasmalib.XorData(inputFile)
	ioutil.WriteFile(*fileNameOutput, decryptedFile, 777)
}

