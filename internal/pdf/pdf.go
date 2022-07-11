package pdf

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/leowmjw/pdf"
	"strings"
)

type PDFPage struct {
	PageNo           int
	PDFPlainText     string
	PDFTxtSameLines  []string // combined content with same line .. proxy for changes
	PDFTxtSameStyles []string // combined content with same style .. proxy for changes
}

func Run() {
	loadPDF("/Users/leow/GOMOD/go-electdocs/internal/pdf/testdata/pub_20190304_PUB120_2019.pdf", 0, 0)
}

// loadPDF will extract out from startPage to endPage
func loadPDF(pdfPath string, startPage, endPage int) error {
	fmt.Println("In loadPDF ..")
	// NOTE: Starts from page 1 by default ..
	f, r, err := pdf.Open(pdfPath)
	// remember close file
	defer f.Close()
	if err != nil {
		return err
	}
	totalPage := r.NumPage()
	//totalPage = 10
	//startPage = 1

	fmt.Println(totalPage)
	// opton #2: Use reformed PDF extractor
	for pageIndex := startPage; pageIndex <= totalPage; pageIndex++ {
		//fmt.Println(pageIndex)
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}
		newPageProcessed := p.ExtractSameLineContent()
		spew.Dump(newPageProcessed)
		// If OK, append them ..
		//pdfDoc.Pages = append(pdfDoc.Pages, newPageProcessed)
		//extractSameLineContent(p.Content().Text)
		//spew.Dump(p.Content())
	}
	// Option #3: Run extenal
	//cmdToRun := fmt.Sprintf("java -jar /Users/leow/DATA/TINDAKMSIA/tabula-java/target/tabula-1.0.5-SNAPSHOT-jar-with-dependencies.jar -t -p1-%d %s", totalPage, "/Users/leow/GOMOD/go-electdocs/internal/pdf/"+pdfPath)
	//// DEBUG
	////fmt.Println(cmdToRun)
	//
	//cmd := exec.Command("sh", "-c", cmdToRun)
	//stdoutStderr, err := cmd.CombinedOutput()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("%s\n", stdoutStderr)

	//for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
	//	p := r.Page(pageIndex)
	//	if p.V.IsNull() {
	//		continue
	//	}
	//	texts := p.Content().Text
	//	xerr := extractSameLineContent(texts)
	//	if xerr != nil {
	//		panic(xerr)
	//	}
	//}
	return nil
}

func extractSameLineContent(pdfContentTxt []pdf.Text) error {
	var numValidLineCounted int
	var currentLineNumber float64
	var currentContent string

	var pdfTxtSameLine []string

	// DEBUG
	//spew.Dump(pdfContentTxt)

	for _, v := range pdfContentTxt {

		// Guard function .. what is it?
		//if strings.TrimSpace(v.S) == "" {
		//	fmt.Println("Skipping blank line / content ..")
		//	continue
		//}

		if currentLineNumber == 0 {
			currentLineNumber = v.Y
			// DEBUG
			//fmt.Println("Set first line to ", currentLineNumber)
			currentContent += v.S
			continue
		}

		// Happy path ..
		// DEBUG
		//fmt.Println("Append CONTENT: ", currentContent, " X: ", v.X, " Y: ", v.Y)
		// number of valid line increase when new valid line ..
		if currentLineNumber != v.Y {
			if strings.TrimSpace(currentContent) != "" {
				// trim new lines ..
				currentContent = strings.ReplaceAll(currentContent, "\n", "")
				// DEBUG
				//fmt.Println("NEW Line ... collected: ", currentContent)
				pdfTxtSameLine = append(pdfTxtSameLine, currentContent)
				numValidLineCounted++
			}
			currentContent = v.S // reset .. after append
			currentLineNumber = v.Y
		} else {
			// If on the same line, just build up the content ..
			currentContent += v.S
		}

		// NOTE: Only get MaxLineProcessed lines ..
		if numValidLineCounted > 10 {
			break
		}

	}
	// All the left over, do one more final check ...
	if strings.TrimSpace(currentContent) != "" {
		// trim new lines ..
		currentContent = strings.ReplaceAll(currentContent, "\n", "")
		// DEBUG
		//fmt.Println("NEW Line ... collected: ", currentContent)
		pdfTxtSameLine = append(pdfTxtSameLine, currentContent)
	}

	fmt.Println(pdfTxtSameLine)

	return nil
}

func extractTxtSameLine(ptrTxtSameLine *[]string, pdfContentTxt []pdf.Text) error {

	var numValidLineCounted int
	var currentLineNumber float64
	var currentContent string

	var pdfTxtSameLine []string

	// DEBUG
	//spew.Dump(pdfContentTxt)

	for _, v := range pdfContentTxt {

		// Guard function .. what is it?
		//if strings.TrimSpace(v.S) == "" {
		//	fmt.Println("Skipping blank line / content ..")
		//	continue
		//}

		if currentLineNumber == 0 {
			currentLineNumber = v.Y
			// DEBUG
			//fmt.Println("Set first line to ", currentLineNumber)
			currentContent += v.S
			continue
		}

		// Happy path ..
		// DEBUG
		//fmt.Println("Append CONTENT: ", currentContent, " X: ", v.X, " Y: ", v.Y)
		// number of valid line increase when new valid line ..
		if currentLineNumber != v.Y {
			if strings.TrimSpace(currentContent) != "" {
				// trim new lines ..
				currentContent = strings.ReplaceAll(currentContent, "\n", "")
				// DEBUG
				//fmt.Println("NEW Line ... collected: ", currentContent)
				pdfTxtSameLine = append(pdfTxtSameLine, currentContent)
				numValidLineCounted++
			}
			currentContent = v.S // reset .. after append
			currentLineNumber = v.Y
		} else {
			// If on the same line, just build up the content ..
			currentContent += v.S
		}

		// NOTE: Only get MaxLineProcessed lines ..
		//if numValidLineCounted > MaxLineProcessed {
		//	break
		//}

	}
	// All the left over, do one more final check ...
	if strings.TrimSpace(currentContent) != "" {
		// trim new lines ..
		currentContent = strings.ReplaceAll(currentContent, "\n", "")
		// DEBUG
		//fmt.Println("NEW Line ... collected: ", currentContent)
		pdfTxtSameLine = append(pdfTxtSameLine, currentContent)
	}

	*ptrTxtSameLine = pdfTxtSameLine
	//spew.Dump(ptrTxtSameLine)
	return nil
}
