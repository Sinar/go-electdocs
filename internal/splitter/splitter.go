package splitter

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

//type mydata struct {
//	v []map[string]any
//}

func extractJSON(input io.Reader, output io.Writer) error {

	// Split before first '['; this will be the indicator; filename?

	// Left over; split at ']'; this will be the content

	// Remaining will go on the next loop after index increased ..

	// For test out; go in stdout ..
	return nil
}

func json2csv(input io.Reader) error {
	//b, err := io.ReadAll(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(b))

	nd := json.NewDecoder(input)
	//content := mydata{}
	//for nd.More() {
	var v []map[string]any
	nerr := nd.Decode(&v)
	if nerr != nil {
		panic(nerr)
	}
	//}
	//spew.Dump(content.v)
	fmt.Println("=========== IDEE!!!!")
	for _, record := range v {
		var row []string
		var rowContent string
		for _, d := range record {
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
	//var v any
	//d := json.NewDecoder(input)
	//for d.More() {
	//	derr := d.Decode(&v)
	//	if derr != nil {
	//		panic(derr)
	//	}
	//	spew.Dump(v)
	//
	//	//for _, r := range v.() {
	//	//
	//	//}
	//}

	return nil
}
