package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	homeDir    string
	nameSearch string
)

func init() {
	homeDir, _ = os.UserHomeDir()
	flag.StringVar(&nameSearch, "name", "", "...")
	flag.Parse()
}

func main() {
	if nameSearch == "" {
		parseXmlAndOutput()
	} else {
		searchFromRecords(nameSearch)
	}

}

func parseXmlAndOutput() {
	wildcardPath := filepath.Join(homeDir, "Library/Application Support/JetBrains/IntelliJIdea*/options/recentProjects.xml")

	xmlPaths, err := filepath.Glob(wildcardPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(xmlPaths) <= 0 {
		return
	}

	file, err := os.OpenFile("./opened.records", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	var items Items
	for _, xmlPath := range xmlPaths {
		a, err := parseXML(xmlPath)
		if err != nil {
			fmt.Println(err)
			return
		}

		items = parseAndSave(a, items, writer)
	}
	writer.Flush()

	if len(items) <= 0 {
		return
	}

	bytes, _ := json.Marshal(items)
	fmt.Printf("{\"items\":%s}", string(bytes))
}

func searchFromRecords(nameSearch string) {
	var items Items

	file, err := os.OpenFile("./opened.records", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()
		_, projectName := filepath.Split(line)
		if strings.Contains(strings.ToLower(projectName), strings.ToLower(nameSearch)) {
			items = append(items, &Item{
				Title:    projectName,
				Subtitle: line,
				Arg:      line,
			})
		}
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	if len(items) <= 0 {
		return
	}

	bytes, _ := json.Marshal(items)
	fmt.Printf("{\"items\":%s}", string(bytes))
}
