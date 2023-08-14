package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func outputCSV(fileName string, records [][]string) {
	// DEBUG
	//records = [][]string{
	//	{"first_name", "last_name", "username"},
	//	{"Rob", "Pike", "rob"},
	//	{"Ken", "Thompson", "ken"},
	//	{"Robert", "Griesemer", "gri"},
	//}
	var iow io.Writer
	var oerr error
	if fileName == "DEBUG" {
		iow = os.Stdout
	} else {
		iow, oerr = os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0744)
		if oerr != nil {
			panic(oerr)
		}
	}
	w := csv.NewWriter(iow)
	w.WriteAll(records) // calls Flush internally

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}
