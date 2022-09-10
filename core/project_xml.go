package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Application struct {
	XMLName   xml.Name   `xml:"application"`
	Component *Component `xml:"component"`
}

type Component struct {
	Options []*Option `xml:"option"`
}

type Option struct {
	Omaps *_Map `xml:"map"`
	Name  string  `xml:"name,attr"`
	Value string  `xml:"value,attr"`
}

type _Map struct {
	Entries []*Entry `xml:"entry"`
}

type Entry struct {
	Key string `xml:"key,attr"`
}

func parseXML(xmlPath string) (*Application, error) {
	bytes, err := os.ReadFile(xmlPath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	a := &Application{}
	err = xml.Unmarshal(bytes, a)

	return a, err
}
