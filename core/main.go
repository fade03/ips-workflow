package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

var homeDir string

func init() {
	homeDir, _ = os.UserHomeDir()
}

func main() {
	wildcardPath := filepath.Join(homeDir, "/Library/Application Support/JetBrains/IntelliJIdea*/options/recentProjects.xml")

	xmlPaths, err := filepath.Glob(wildcardPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	items := make([]*Item, 0)
	for _, xmlPath := range xmlPaths {
		a, err := parseXML(xmlPath)
		if err != nil {
			fmt.Println(err)
			return
		}

		items = parseItems(a, items)
	}

	bytes, err := json.Marshal(items)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("{\"items\":%s}", string(bytes))
}
