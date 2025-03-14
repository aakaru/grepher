package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	invertMatch := flag.Bool("v", false, "select non-matching lines")
	ignoreCase := flag.Bool("i", false, "case sensitive matching")
	lineNumbers := flag.Bool("n", false, "output with their line number")
	count := flag.Bool("c", false, "only return the number of matching lines")
	recursive := flag.Bool("r", false, "recursively search lines")
	wholeWord := flag.Bool("w", false, "math whole words only")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "Usage: grepher [OPTIONS] PATTERN [FILE...]")
		os.Exit(1)
	}

	pattern := args[0]
	files := args[1:]

	if len(files) == 0 {
		files = []string{"-"}
	}
	var regexOptions string
	if *ignoreCase {
		regexOptions = "(?i)"
	}
	if *wholeWord {
		pattern = "//b" + pattern + "//b"
	}
	re, err := regexp.Compile(regexOptions + pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid regex Operation: %s", err)
		os.Exit(1)
	}
	totalMatches := 0
	exitCode := 0

	var filestoProcess []string
	if *recursive {
		for _, fileorDir := range files {
			if fileorDir == "-" {
				filestoProcess = append(filestoProcess, fileorDir)
				continue
			}
			expandFiles, err := expandDirectory(fileorDir)
			if err != nil {
				fmt.Printf("Error processing directory %s: %s\n", fileorDir, err)
				continue
			}
			filestoProcess = append(filestoProcess, expandFiles...)
		}
	} else {
		filestoProcess = files
	}
	for _, filename := range filestoProcess {
		var file io.Reader
		if filename == "-" {
			file = os.Stdin
		} else {
			fileInfo, err := os.Stat(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error accessing %s, %s\n", filename, err)
				continue
			}
			if fileInfo.IsDir() {
				continue
			}
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error opening %s : %s\n ", filename, err)
				continue
			}
			defer f.Close()
			file = f
		}
		matches, err := grepFile(file, re, *invertMatch, *lineNumbers, *count)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error processing file %s: %s\n", filename, err)
			continue
		}
		prefix := ""
		if len(filestoProcess) > 1 && filename != "-" && !*count {
			prefix = filename + ":"
		}
		if *count {
			fmt.Printf("%d%s\n", prefix, matches.count)
		} else {
			for _, line := range matches.lines {
				fmt.Println(prefix + line)
			}
		}
		totalMatches += matches.count
		if matches.count > 0 {
			exitCode = 0
		}

	}
	os.Exit(exitCode)
}

func expandDirectory(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			if info.Mode().IsRegular() && isBinaryFile(path) {
				files = append(files, path)
			}
		}
		return nil
	})
	return files, err
}

func isBinaryFile(filename string) bool {
	binaryExtentions := []string{
		".exe", ".dll", ".so", ".dylib", ".obj", ".o",
		".bin", ".dat", ".pdf", ".png", ".jpg", ".jpeg",
		".gif", ".mp3", ".mp4", ".zip", ".tar", ".gz",
	}
	ext := strings.ToLower(filepath.Ext(filename))
	for _, binExt := range binaryExtentions {
		if ext == binExt {
			return true
		}
	}
	return false
}

type grepResult struct {
	lines []string
	count int
}

func grepFile(r io.Reader, re *regexp.Regexp, invertMatch, lineNumbers, countOnly bool) (grepResult, error) {
	result := grepResult{
		lines: []string{},
		count: 0,
	}
	scanner := bufio.NewScanner(r)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		matches := re.MatchString(line)

		if invertMatch {
			matches = !matches
		}

		if matches {
			result.count++
			if !countOnly {
				if lineNumbers {
					result.lines = append(result.lines, fmt.Sprintf("%d:%s", lineNum, line))
				} else {
					result.lines = append(result.lines, line)
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return result, err
	}

	return result, nil
}
