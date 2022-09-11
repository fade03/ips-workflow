package main

import (
	"os"
	"path/filepath"
	"strings"
)

type Item struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Arg      string `json:"arg"`
}

type Items []*Item

func parseItems(a *Application, items []*Item) []*Item {
	for _, option := range a.Component.Options {
		if option.Name == "lastOpenedProject" {
			if realPath, isExist := getRealPath(option.Value); isExist {
				items = append(items, &Item{
					Title: "Last Opened Project",
					Subtitle: realPath,
					Arg: realPath,
				})	
			}
		} else {
			for _, entry := range option.Maps.Entries {
				if realPath, isExist := getRealPath(entry.Key); isExist {
					_, projectName := filepath.Split(realPath)
					items = append(items, &Item{
						Title: projectName,
						Subtitle: realPath,
						Arg: realPath,
					})
				}
			}
		}
	}

	return items
}

func getRealPath(path string) (string, bool) {
	realPath := strings.ReplaceAll(path, "$USER_HOME$", homeDir)
	if _, err := os.Stat(realPath); os.IsNotExist(err) {
		return "", false
	}
	
	return realPath, true
}