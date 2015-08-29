package info

type InfoEntity struct {
	Name    string
	Version string
}

type Info struct {
	Language  InfoEntity
	Framework InfoEntity
	Server    InfoEntity
	Package   InfoEntity
	OS        InfoEntity

	Country   string
	Ip        string
	Host      string
	Place     string
}

func NewInfo() Info {
	language := InfoEntity{"", ""}
	framework := InfoEntity{"", ""}
	server := InfoEntity{"", ""}
	pkg := InfoEntity{"", ""}
	os := InfoEntity{"", ""}

	info := Info{language, framework, server, pkg, os, "", "", "", ""}
	return info
}