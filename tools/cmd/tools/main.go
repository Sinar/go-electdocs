package main

import (
	"fmt"
	"github.com/leowmjw/pdf"
	"log"
	"os/exec"
)

func main() {
	fmt.Println("Toolz ...")
	extractTextWithTabula("/Users/leow/GOMOD/go-electdocs/internal/pdf/testdata/pub_20190304_PUB120_2019.pdf")
}

func extractTextWithTabula(pdfPath string) error {
	fmt.Println("In extractTextWithTabula ..")
	f, r, err := pdf.Open(pdfPath)
	// remember close file
	defer f.Close()
	if err != nil {
		return err
	}
	totalPage := r.NumPage()

	// DEBUG
	//	fmt.Println(totalPage)
	// Option #3: Run extenal
	cmdToRun := fmt.Sprintf("java -jar /Users/leow/DATA/TINDAKMSIA/tabula-java/target/tabula-1.0.5-SNAPSHOT-jar-with-dependencies.jar -t -p1-%d %s", totalPage, pdfPath)
	// DEBUG
	fmt.Println(cmdToRun)

	cmd := exec.Command("sh", "-c", cmdToRun)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)

	return nil
}
