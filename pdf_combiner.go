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
	// Define and parse the flag for PDF files
	pdfs := flag.String("pdfs", "", "Comma-separated list of PDF files to merge")
	output := flag.String("output", "", "Name of merged file")
	flag.Parse()

	// Split the input PDFs and check if the flag was provided
	inputFiles := strings.Split(*pdfs, ",")
	if len(inputFiles) == 0 || *pdfs == "" {
		fmt.Println("Please provide PDF files using -pdfs flag, e.g., combiner -pdfs=\"file1.pdf,file2.pdf\"")
		os.Exit(1)
	}

	// Determine the directory of the first input file for storing the merged output
	outputDir := filepath.Dir(inputFiles[0])
	outputFile := filepath.Join(outputDir, *output+".pdf")

	// Merge the PDF files, passing `nil` for the configuration
	err := api.MergeCreateFile(inputFiles, outputFile, false, nil)
	if err != nil {
		fmt.Printf("Error merging PDFs: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("PDFs merged successfully into %s\n", outputFile)
}
