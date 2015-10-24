package info

import (
	"net/http"
	urlLib "net/url"
	"net"
)

type InfoEntity struct {
	Name    string
	Version string
}

type ISPEntity struct {
	Name   string
	Region string
}

type Info struct {

	Url          string

	Language     InfoEntity
	Framework    InfoEntity
	Server       InfoEntity
	Package      InfoEntity
	OS           InfoEntity

	Distribution string
	Country      string
	Host         string
	Place        string

	Ip           []string
	RealHost     []string

	ISP          ISPEntity

	RawHeaders   map[string][]string
	Cookies      map[string]*http.Cookie
}

func NewInfo() Info {
	language := InfoEntity{"", ""}
	framework := InfoEntity{"", ""}
	server := InfoEntity{"", ""}
	pkg := InfoEntity{"", ""}
	os := InfoEntity{"", ""}
	isp := ISPEntity{"", ""}
	info := Info{
		"", language, framework, server, pkg, os, "", "", "", "", make([]string, 5), make([]string, 5), isp,
		make(map[string][]string), make(map[string]*http.Cookie)}
	return info
}

func (info *Info) SetBasicInfo(url string, response *http.Response) {
	info.Url = url
	for key, values := range response.Header {
		info.RawHeaders[key] = values
	}
	cookies := response.Cookies()
	for _, cookie := range cookies {
		info.Cookies[cookie.Name] = cookie
	}
	u, err := urlLib.Parse(url)
	if err != nil {
		panic(err)
	}
	info.Host = u.Host
	ips, err := net.LookupHost(u.Host)
	if err == nil {
		info.Ip = ips
		for _, ip := range info.Ip {
			hosts, err := net.LookupAddr(ip)
			if err == nil {
				info.RealHost = hosts
			}
		}

	}
}