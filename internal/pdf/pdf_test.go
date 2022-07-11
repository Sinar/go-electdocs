package pdf

import (
	"testing"
)

func Test_loadPDF(t *testing.T) {
	type args struct {
		pdfPath   string
		startPage int
		endPage   int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"case #1 happy ..",
			args{
				pdfPath:   "./testdata/PUB166_2022.pdf",
				startPage: 0,
				endPage:   0,
			}, true,
		},
		{
			"case sarawak",
			args{
				pdfPath:   "./testdata/pub_20190415_PUB187.pdf",
				startPage: 0,
				endPage:   0,
			}, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := loadPDF(tt.args.pdfPath, tt.args.startPage, tt.args.endPage); (err != nil) != tt.wantErr {
				t.Errorf("loadPDF() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
