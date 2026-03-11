package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type FileInfo struct {
	name string
	size int64
	kind string
	ext  string
}

const maxWidth = 20
const strFormat = "%-25s %-15s %-15s %s\n"

func formatSize(size int64) string {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
	)

	kb := float64(size) / KB
	mb := float64(size) / MB
	gb := float64(size) / GB

	if gb >= 1 {
		return fmt.Sprintf("%.2f GB", gb)
	} else if mb >= 1 {
		return fmt.Sprintf("%.2f MB", mb)
	} else if kb >= 1 {
		return fmt.Sprintf("%.2f KB", kb)
	}

	return fmt.Sprintf("%d B", size)
}

func truncate(s string, max int) string {
	if len(s) > max {
		return s[:max-3] + "..."
	}
	return s
}

func dirCount(path string) string {
	items, err := os.ReadDir(path)
	if err != nil {
		return "?"
	}

	count := len(items)
	if count == 1 {
		return fmt.Sprintf("%d item", len(items))
	}
	return fmt.Sprintf("%d items", len(items))
}

func formatHeader(nameCol, typeCol, extCol, sizeCol string, divider int) {
	fmt.Printf("\n"+strFormat, nameCol, typeCol, extCol, sizeCol)
	fmt.Println(strings.Repeat("-", divider))
}

func main() {
	path := "."

	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
		return
	}

	var fileList []FileInfo

	for _, file := range files {

		extension := filepath.Ext(file.Name())

		fileType := "file"
		iconType := "\U0000ea7b"

		if file.IsDir() {
			extension = dirCount(path + "/" + file.Name())
			fileType = "directory"
			iconType = "\U0000f114"
		}

		if extension == "" {
			extension = "unknown"
		}

		info, err := file.Info()
		if err != nil { // skip file and continue
			continue
		}

		fileList = append(fileList, FileInfo{
			name: file.Name(),
			size: info.Size(),
			kind: iconType + "  " + fileType,
			ext:  extension,
		})
	}

	// sort files based on size - largest to smallest
	sort.Slice(fileList, func(i, j int) bool {
		return fileList[i].size > fileList[j].size // i: position 0 - j: position 1
	})

	formatHeader("Name", "Type", "", "Size", 65)

	for _, file := range fileList {
		fmt.Printf(strFormat, truncate(file.name, maxWidth), file.kind, file.ext, formatSize(file.size))
	}
}
