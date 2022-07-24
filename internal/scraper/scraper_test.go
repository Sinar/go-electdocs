package scraper

import (
	"strings"
	"testing"
)

func Test_json2csv(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"happy #1", args{
			input: `[{
"id":"20"
}]`}, true},
		{"bad #1", args{
			input: `[{
id:  20, name: "min"
},
{
id : 42, name: "boo" ,
kerja : "mata mata",
},
]`}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := json2csv(strings.NewReader(tt.args.input)); (err != nil) != tt.wantErr {
				t.Errorf("json2csv() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
