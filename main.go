package main

import (
	"math"
	"os"
	"fmt"
	"strings"
	"net/http"
//	"mime"

	ui "github.com/gizak/termui"

	"github.com/PuerkitoBio/goquery"
	"github.com/takaaki-mizuno/web-architecture-analyzer/detectors"
	"github.com/takaaki-mizuno/web-architecture-analyzer/detectors/languages"
	"github.com/takaaki-mizuno/web-architecture-analyzer/detectors/frameworks"
	"github.com/takaaki-mizuno/web-architecture-analyzer/detectors/servers"
	"github.com/takaaki-mizuno/web-architecture-analyzer/detectors/packages"
	"github.com/takaaki-mizuno/web-architecture-analyzer/info"
	"github.com/takaaki-mizuno/web-architecture-analyzer/detectors/isps"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please specify the URL you want to check.")
		os.Exit(0)
	}

	url := os.Args[1]
	info := detect(url)

	err := ui.Init()
	if err != nil {
		panic(err)
	}
	buildUI(info)
	defer ui.Close()

}

func detect(url string) info.Info {
	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	doc, err := goquery.NewDocumentFromResponse(response)

	info := info.NewInfo()
	info.SetBasicInfo(url, response)

	detector := detectors.Detector{response, doc, &info}

	languageDetectors := [3]detectors.DetectorInterface{
		&languages.PHP{&detector}, &languages.Python{&detector}, &languages.Ruby{&detector}}
	for _, detector := range languageDetectors {
		result := detector.Detect()
		if result == true {
			break
		}
	}

	frameworkDetectors := [9]detectors.DetectorInterface{
		&frameworks.Laravel{&detector}, &frameworks.CakePHP{&detector}, &frameworks.Play{&detector},
		&frameworks.FuelPHP{&detector}, &frameworks.Codeigniter{&detector},
        &frameworks.JBoss{&detector}, &frameworks.Servlet{&detector}, &frameworks.Rails{&detector},
		&frameworks.Default{&detector}}
	for _, detector := range frameworkDetectors {
		result := detector.Detect()
		if result == true {
			break
		}
	}

	serverDetectors := [3]detectors.DetectorInterface{
		&servers.Coyote{&detector}, &servers.Apache{&detector}, &servers.Nginx{&detector}}
	for _, detector := range serverDetectors {
		result := detector.Detect()
		if result == true {
			break
		}
	}

	packageDetectors := [2]detectors.DetectorInterface{
		&packages.WordPress{&detector},&packages.Drupal{&detector}}
	for _, detector := range packageDetectors {
		result := detector.Detect()
		if result == true {
			break
		}
	}

	ispDetectors := [1]detectors.DetectorInterface{&isps.AWS{&detector}}
	for _, detector := range ispDetectors {
		result := detector.Detect()
		if result == true {
			break
		}
	}

	return info
}

func buildUI(info info.Info) {

	urlLabel := ui.NewPar(fmt.Sprintf("%s", info.Url))
	urlLabel.Height = 3
	urlLabel.TextFgColor = ui.ColorWhite
	urlLabel.BorderLabel = "URL"
	urlLabel.BorderFg = ui.ColorCyan

	basicInfoStrings := []string{
		fmt.Sprintf("IP: %s", strings.Join(info.Ip, ",")),
		fmt.Sprintf("Host: %s",  strings.Join(info.RealHost, ",")),
		fmt.Sprintf("Distribution: %s",  info.Distribution),
	}

	basicInfoList := ui.NewList()
	basicInfoList.Items = basicInfoStrings
	basicInfoList.ItemFgColor = ui.ColorYellow
	basicInfoList.BorderLabel = "BasicInfo"
	basicInfoList.Height = 15
	basicInfoList.Width = 25
	basicInfoList.Y = 0

	serverLabel := ui.NewPar(fmt.Sprintf("Name: %s\nVersion: %s", info.Server.Name, info.Server.Version))
	serverLabel.Height = 4
	serverLabel.TextFgColor = ui.ColorWhite
	serverLabel.BorderLabel = "Server"
	serverLabel.BorderFg = ui.ColorCyan

	languageLabel := ui.NewPar(fmt.Sprintf("Name: %s\nVersion: %s", info.Language.Name, info.Language.Version))
	languageLabel.Height = 4
	languageLabel.TextFgColor = ui.ColorWhite
	languageLabel.BorderLabel = "Language"
	languageLabel.BorderFg = ui.ColorCyan

	frameworkLabel := ui.NewPar(fmt.Sprintf("Name: %s\nVersion: %s", info.Framework.Name, info.Framework.Version))
	frameworkLabel.Height = 4
	frameworkLabel.TextFgColor = ui.ColorWhite
	frameworkLabel.BorderLabel = "Framework"
	frameworkLabel.BorderFg = ui.ColorCyan

	packageLabel := ui.NewPar(fmt.Sprintf("Name: %s\nVersion: %s", info.Package.Name, info.Package.Version))
	packageLabel.Height = 4
	packageLabel.TextFgColor = ui.ColorWhite
	packageLabel.BorderLabel = "Package"
	packageLabel.BorderFg = ui.ColorCyan

	ispLabel := ui.NewPar(fmt.Sprintf("Name: %s\nRegion: %s", info.ISP.Name, info.ISP.Region))
	ispLabel.Height = 4
	ispLabel.TextFgColor = ui.ColorWhite
	ispLabel.BorderLabel = "ISP"
	ispLabel.BorderFg = ui.ColorCyan

	var headerStrings []string
	for key, values := range info.RawHeaders {
		headerStrings = append(headerStrings, fmt.Sprintf("%s: %s", key, strings.Join(values[:], ",")))
	}

	headerList := ui.NewList()
	headerList.Items = headerStrings
	headerList.ItemFgColor = ui.ColorYellow
	headerList.BorderLabel = "Raw Headers"
	headerList.Height = 25
	headerList.Y = 0

	var cookieStrings []string
	for key, values := range info.Cookies {
		cookieStrings = append(cookieStrings, fmt.Sprintf("%s: %s", key, values))
	}

	cookieList := ui.NewList()
	cookieList.Items = cookieStrings
	cookieList.ItemFgColor = ui.ColorYellow
	cookieList.BorderLabel = "Cookies"
	cookieList.Height = 25
	cookieList.Y = 0

	/* demo */

	sinps := (func() []float64 {
		n := 400
		ps := make([]float64, n)
		for i := range ps {
			ps[i] = 1 + math.Sin(float64(i)/5)
		}
		return ps
	})()
	sinpsint := (func() []int {
		ps := make([]int, len(sinps))
		for i, v := range sinps {
			ps[i] = int(100*v + 10)
		}
		return ps
	})()


	spark := ui.Sparkline{}
	spark.Height = 12
	spdata := sinpsint
	spark.Data = spdata[:100]
	spark.LineColor = ui.ColorCyan
	spark.TitleColor = ui.ColorWhite

	sp := ui.NewSparklines(spark)
	sp.Height = 15
	sp.BorderLabel = "Sparkline"

	g1 := ui.NewGauge()
	g1.Percent = 30
	g1.Height = 4
	g1.Y = 6
	g1.BorderLabel = "Progress"
	g1.PercentColor = ui.ColorYellow
	g1.BarColor = ui.ColorGreen
	g1.BorderFg = ui.ColorWhite
	g1.BorderLabelFg = ui.ColorMagenta

	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(12, 0, urlLabel),
		),
		ui.NewRow(
			ui.NewCol(6, 0, basicInfoList),
			ui.NewCol(6, 0, sp),
		),
		ui.NewRow(
			ui.NewCol(4, 0, serverLabel),
			ui.NewCol(4, 0, languageLabel),
			ui.NewCol(4, 0, frameworkLabel),
		),
		ui.NewRow(
			ui.NewCol(4, 0, packageLabel),
			ui.NewCol(4, 0, ispLabel),
			ui.NewCol(4, 0, g1),
		),
		ui.NewRow(
			ui.NewCol(6, 0, headerList),
			ui.NewCol(6, 0, cookieList),
		),
	)
	ui.Body.Align()

    ui.Render(ui.Body)

//	redraw := make(chan bool)

    ui.Handle("/sys/kbd/q", func(ui.Event) {
        // press q to quit
        ui.StopLoop()
    })

    ui.Handle("/sys/wnd/resize", func(ui.Event) {
        // press q to quit
        ui.Body.Width = ui.TermWidth()
        ui.Body.Align()
    })


    ui.Handle("/timer/1s", func(e ui.Event) {
        ui.Body.Align()
    })

    ui.Loop()
}
