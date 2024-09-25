package proto

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (
	Prefix = `syntax = "proto3";`
)

func Merge(dir, pbName, target string, prefixs ...string) {
	var prefix string
	if len(prefixs) > 0 {
		prefix = prefixs[0]
	} else {
		prefix = Prefix
	}

	// dir, _ := os.Getwd()
	folderPath := path.Join(dir, pbName, "proto")    // File path for .proto files
	outputFilePath := path.Join(dir, pbName, target) // New file path

	files, err := os.ReadDir(folderPath)
	if err != nil {
		panic(err)
	}

	var fileContents []byte
	for index, file := range files {
		if !strings.HasSuffix(file.Name(), ".proto") {
			continue
		}

		fileName := filepath.Join(folderPath, file.Name())
		fileData, err := os.ReadFile(fileName)
		if err != nil {
			panic(err)
		}

		if index > 0 {
			// Remove the syntax line for all but the first file
			fileDataStr := strings.Replace(string(fileData), prefix, "", 1)
			fileContents = append(fileContents, []byte(fileDataStr)...)
		} else {
			// Keep the syntax line for the first file
			fileContents = append(fileContents, fileData...)
		}
	}

	// Define the new prefix to be placed at the top of the final merged file
	var targetPrefix = `syntax = "proto3";
package ` + pbName + `;
option go_package = "./` + pbName + `";
`

	// Replace the original prefix with the new targetPrefix in the merged contents
	finalContents := strings.Replace(string(fileContents), prefix, targetPrefix, 1)

	// Write the modified content to the output file
	err = os.WriteFile(outputFilePath, []byte(finalContents), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Merged files to", outputFilePath)
}
