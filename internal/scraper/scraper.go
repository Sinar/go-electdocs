package scraper

import (
	"encoding/csv"
	"fmt"
	"github.com/dop251/jsonx"
	"io"
	"log"
	"os"
)

func json2csv(input io.Reader) error {
	b, err := io.ReadAll(input)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(b))
	nd := jsonx.NewDecoder(b)
	//content := mydata{}
	//for nd.More() {
	var records any
	records, nerr := nd.Decode()
	if nerr != nil {
		panic(nerr)
	}
	//}
	//spew.Dump(content.v)
	fmt.Println("=========== IDEE!!!!")
	for _, record := range records.([]any) {
		var v map[string]any
		v = record.(map[string]any)
		// DEBUG
		//spew.Dump(v)
		var row []string
		var rowContent string
		for _, d := range v {
			// DEBUG
			//fmt.Println("KEY: ", k, " VAL: ", d)
			switch d.(type) {
			case string:
				rowContent = fmt.Sprintf("%s", d)
			case int:
				rowContent = fmt.Sprintf("%d", d)
			default:
				rowContent = fmt.Sprintf("%v", d)
			}
			row = append(row, rowContent)
		}
		// DEBUG
		//spew.Dump(row)
		w := csv.NewWriter(os.Stdout)
		werr := w.Write(row)
		if werr != nil {
			log.Fatal(werr)
		}
		// Write any buffered data to the underlying writer (standard output).
		w.Flush()
		if err := w.Error(); err != nil {
			log.Fatal(err)
		}

	}
	return nil
}
