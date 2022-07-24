package splitter

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_json2csv(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"happy #1", args{
			input: strings.NewReader(`[{
"id":20,
"nama" : "bob"
},{
"id":"30",
"nama" : "min",
kerja: "nelayan"
}]`)}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := json2csv(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("json2csv() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_extractJSON(t *testing.T) {
	type args struct {
		input  io.Reader
		output map[string]string
	}

	//bufio.NewReader()
	f, oerr := os.Open("../../data/pru-14-2018/0-calon.txt")
	if oerr != nil {
		t.Fatal(oerr)
	}
	// mutliple entirees ..
	mf, moerr := os.Open("../../data/pru-14-2018/0-data.txt")
	if moerr != nil {
		t.Fatal(oerr)
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"happy #1", args{
			input:  f,
			output: make(map[string]string),
		},
			true},
		{
			"multiple #1", args{
			input:  mf,
			output: make(map[string]string),
		},
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := extractJSON(tt.args.input, tt.args.output); (err != nil) != tt.wantErr {
				t.Errorf("extractJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			//spew.Dump(tt.args.output)
			for _, content := range tt.args.output {
				jerr := json2csv(strings.NewReader(content))
				if jerr != nil {
					t.Fatal(jerr)
				}

			}
		})
	}
}
