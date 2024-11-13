package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func main() {
	pdfs := flag.String("pdfs", "", "Comma-separated list of PDF files to merge")
	output := flag.String("output", "", "Name of merged file")
	flag.Parse()

	inputFiles := strings.Split(*pdfs, ",")
	if len(inputFiles) == 0 || *pdfs == "" {
		fmt.Println("Please provide PDF files using -pdfs flag, e.g., combiner -pdfs=\"file1.pdf,file2.pdf\"")
		os.Exit(1)
	}

	outputDir := filepath.Dir(inputFiles[0])
	outputFile := filepath.Join(outputDir, *output+".pdf")

	err := api.MergeCreateFile(inputFiles, outputFile, false, nil)
	if err != nil {
		fmt.Printf("Error merging PDFs: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("PDFs merged successfully into %s\n", outputFile)
}
