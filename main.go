package main

import (
	"os"
	"fmt"
	"net/http"
//	"mime"

	"github.com/PuerkitoBio/goquery"
	"github.com/takaaki-mizuno/web-architecture-analyzer/detectors"
	"github.com/takaaki-mizuno/web-architecture-analyzer/detectors/languages"
	"github.com/takaaki-mizuno/web-architecture-analyzer/detectors/info"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please specify the URL you want to check.")
		os.Exit(0)
	}

	url := os.Args[1]
	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	doc, err := goquery.NewDocumentFromResponse(response)

	info := info.NewInfo()
	detector := detectors.Detector{response, doc, &info}

	languageDetectors := [2]detectors.DetectorInterface{&languages.PHP{&detector}, &languages.Python{&detector}}
	for _, detector := range languageDetectors {
		detector.Detect()
	}

	php := languages.PHP{&detector}
	php.Detect()

	fmt.Println("Language Name:", info.Language.Name);
	fmt.Println("Language Version:", info.Language.Version);

	/*
	doc.Find("div").Each(func(i int, s *goquery.Selection) {
		link := s.Find("a").Text()
		fmt.Printf("Link: %s\n", link)
	})



	mimetype, parameters, err := mime.ParseMediaType(response.Header["Content-Type"][0])
	fmt.Println("Media type : ", mimetype)

	for param := range parameters {
		fmt.Printf("%v = %v\n\n", param, parameters[param])
	}
	*/


}