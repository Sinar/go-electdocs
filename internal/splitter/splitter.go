package splitter

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

//type mydata struct {
//	v []map[string]any
//}

func extractJSON(input io.Reader, output map[string]string) error {

	b, rerr := io.ReadAll(input)
	if rerr != nil {
		return rerr
	}
	// Start the loop with the inital data
	startContent := string(b)
loop:
	// Split before first '['; this will be the indicator; filename?
	before, after, found := strings.Cut(startContent, "[")
	if !found {
		fmt.Println("INVALID OPENING ... exiting ..")
		//return errors.New("Found no opening JSON")
		return nil
	}

	// Left over; split at ']'; this will be the content
	wholeJSON, remaining, found := strings.Cut(after, "]")
	if !found {
		fmt.Println("INVALID CLOSING ... exiting ...")
		return nil
		//return errors.New("Found no closing JSON")
	}
	// Send it back out??
	// Write only if has JSON in ID
	for _, id := range strings.Split(before, " ") {
		if strings.Contains(id, "JSON") || strings.Contains(id, "data") {
			// DEBUG
			fmt.Println("ID: ", id)
			output[id] = fmt.Sprintf("[%s]", wholeJSON)
		}
	}

	// Remaining will go on the next loop after index increased ..
	// DEBUG
	//fmt.Println("LEFTOVERS ..")
	//fmt.Println(remaining)

	if len(remaining) > 0 {
		startContent = remaining
		goto loop
	}

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
	fmt.Println("=========== CSV HEADER!!!!")
	// TODO: BUG map order NOT guaranteed!
	// FIrst row to determine the the column name ..
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
			case nil:
				rowContent = ""
			default:
				rowContent = fmt.Sprintf("%v", d)
			}
			row = append(row, rowContent)
		}
		// Add newline; not needed!! is mesed uop due to map being random ,..
		row = append(row, " ")
		w := csv.NewWriter(os.Stdout)
		w.UseCRLF = false
		werr := w.Write(row)
		if werr != nil {
			return werr
		}
		// Write any buffered data to the underlying writer (standard output).
		w.Flush()
		if err := w.Error(); err != nil {
			return err
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
