package splitter

import (
	"io"
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
