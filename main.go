package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
)

const (
	defaultZones = "Australia/Sydney,Asia/Tokyo,Europe/London,America/New_York"

	defaultLayout = "vertical"
)

func main() {
	tz := ""
	layout := ""

	flag.StringVar(&tz, "tz", defaultZones, "Comma separated list of zones. e.g Australia/Sydney,AsiaTokyo")
	flag.StringVar(&layout, "layout", defaultLayout, "Layout horizontal or vertical.")
	flag.Parse()

	if len(tz) == 0 || len(layout) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	zones := strings.Split(tz, ",")

	locations := make([]*time.Location, len(zones), len(zones))

	for i, zone := range zones {
		loc, err := time.LoadLocation(zone)
		if err != nil {
			panic(err)
		}

		locations[i] = loc
	}

	if layout == "vertical" {
		vertical(zones, locations)
		return
	}

	horizontal(zones, locations)
}

func horizontal(zones []string, locations []*time.Location) {
	for {
		// clear
		fmt.Print("\033[H\033[2J")

		// setup a table writer
		w := tablewriter.NewWriter(os.Stdout)
		w.SetHeader(zones)
		w.SetAlignment(tablewriter.ALIGN_RIGHT)

		// create the times for each timezone
		data := []string{}

		n := time.Now()

		for _, loc := range locations {
			data = append(data, n.In(loc).Format("15:04:05"))
		}

		// add the times to the table
		w.AppendBulk([][]string{data})

		// write the table
		w.Render()

		time.Sleep(1 * time.Second)
	}
}

func vertical(zones []string, locations []*time.Location) {
	// vertical clocks
	for {
		// clear
		fmt.Print("\033[H\033[2J")

		// setup a table writer
		w := tablewriter.NewWriter(os.Stdout)
		w.SetAlignment(tablewriter.ALIGN_LEFT)

		// create the times for each timezone
		lines := [][]string{}

		n := time.Now()

		for i, loc := range locations {
			line := []string{}
			line = append(line, zones[i])
			line = append(line, n.In(loc).Format("15:04:05"))

			lines = append(lines, line)
		}

		// add the times to the table
		w.AppendBulk(lines)

		// write the table
		w.Render()

		time.Sleep(1 * time.Second)
	}
}
