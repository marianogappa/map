package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

const htmlTemplate = `<!DOCTYPE html>
<html lang="en">

<head>
    <base target="_top">
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <title>{{.MainTitle}}</title>

    <link rel="shortcut icon" type="image/x-icon" href="docs/images/favicon.ico" />

    <link rel="stylesheet" href="https://unpkg.com/leaflet@1.9.4/dist/leaflet.css"
        integrity="sha256-p4NxAoJBhIIN+hmNHrzRCf9tD/miZyoHS5obTRR9BMY=" crossorigin="" />
    <script src="https://unpkg.com/leaflet@1.9.4/dist/leaflet.js"
        integrity="sha256-20nQCchB9co0qIjJZRGuk2/Z9VM+kNiyxNV1lvTlZBo=" crossorigin=""></script>

    <style>
        html,
        body {
            height: 100%;
            margin: 0;
        }

        .leaflet-container {
            height: 100%;
            width: 100%;
        }
    </style>


</head>

<body>

    <div id='map'></div>

    <script>
        const map = L.map('map').setView([20, 20], 2.5);

        L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
            attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
        }).addTo(map);

        {{range .Marker}}
        L.marker([{{.Lat}}, {{.Lng}}], {}).bindPopup('{{.Label}}').addTo(map);
        {{end}}

    </script>

</body>

</html>
`

type Marker struct {
	Lat   float64
	Lng   float64
	Label string
}

type TemplateData struct {
	MainTitle string
	Marker    []Marker
}

func main() {
	// Parse command line flags
	var mainTitle string
	var separator string
	flag.StringVar(&mainTitle, "title", "World Map", "Main title for the HTML file")
	flag.StringVar(&separator, "separator", "tab", "Separator for input data (tab or comma)")
	flag.Parse()

	// Validate separator flag
	if separator != "tab" && separator != "comma" {
		log.Fatal("Invalid separator. Please use 'tab' or 'comma'")
	}

	// Determine the correct command to open the browser based on the OS
	var openCmd string
	switch runtime.GOOS {
	case "darwin":
		openCmd = "open"
	case "linux":
		openCmd = "xdg-open"
	case "windows":
		openCmd = "cmd /c start"
	default:
		log.Fatal("Unsupported operating system")
	}

	// Read input from stdin
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal("Failed to read input:", err)
	}

	// Split input into lines and parse into markers
	lines := strings.Split(string(input), "\n")
	var markers []Marker
	for _, line := range lines {
		var parts []string
		if separator == "tab" {
			parts = strings.Split(line, "\t")
		} else {
			parts = strings.Split(line, ",")
		}
		if len(parts) == 3 {
			lat, lng := parseCoords(parts[0], parts[1])
			markers = append(markers, Marker{
				Lat:   lat,
				Lng:   lng,
				Label: parts[2],
			})
		}
	}

	// Prepare data for HTML template
	data := TemplateData{
		MainTitle: mainTitle,
		Marker:    markers,
	}

	// Execute HTML template
	tmpl := template.Must(template.New("map").Parse(htmlTemplate))
	tmpfile, err := ioutil.TempFile("", "map-*.html")
	if err != nil {
		log.Fatal("Failed to create temporary file:", err)
	}
	defer tmpfile.Close()
	err = tmpl.Execute(tmpfile, data)
	if err != nil {
		log.Fatal("Failed to execute template:", err)
	}

	// Open HTML file in default browser
	cmd := exec.Command(openCmd, tmpfile.Name())
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to open browser:", err)
	}
}

func parseCoords(latStr, lngStr string) (float64, float64) {
	lat := parseFloat(latStr)
	lng := parseFloat(lngStr)
	return lat, lng
}

func parseFloat(s string) float64 {
	val := strings.TrimSpace(s)
	if strings.Contains(val, ",") {
		val = strings.ReplaceAll(val, ",", ".")
	}
	f, _ := strconv.ParseFloat(val, 64)
	return f
}
