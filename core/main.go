package main

import (
	"encoding/json"
	"flag"
	"fmt"
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
	flag.StringVar(&nameSearch, "name", "", "-")
	flag.Parse()
}

func main() {
	wildcardPath := filepath.Join(homeDir, "Library/Application Support/JetBrains/IntelliJIdea*/options/recentProjects.xml")

	xmlPaths, err := filepath.Glob(wildcardPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(xmlPaths) <= 0 {
		return
	}

	var items Items
	for _, xmlPath := range xmlPaths {
		a, err := parseXML(xmlPath)
		if err != nil {
			fmt.Println(err)
			return
		}

		items = parseItems(a, items)
	}

	if len(items) <= 0 {
		return
	}

	doOutput(items, nameSearch)
}

func doOutput(items Items, nameSearch string) {
	var targetItems Items

	if nameSearch == "" {
		targetItems = items
	} else {
		for _, item := range items[:len(items)-1] {
			if strings.Contains(strings.ToLower(item.Title), strings.ToLower(nameSearch)) {
				targetItems = append(targetItems, item)
			}
		}
	}

	if len(targetItems) > 0 {
		bytes, _ := json.Marshal(targetItems)
		fmt.Printf("{\"items\":%s}", string(bytes))
	}
}
