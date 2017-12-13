package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/atotto/clipboard"
)

func main() {
	now := time.Now().UTC()
	fmt.Printf("[%s]: %s\n", now.Format("03:04 PM"), "Starting, simply select all (Ctrl-A) on the Wormhole description and copy it (Ctrl-C), this log is in EVE Time.")
	for {
		// Refresh our time each loop - is there a better way to do this?
		now = time.Now().UTC()
		clipboardLoop(now)
		time.Sleep(500 * time.Millisecond)
	}
}

func clipboardLoop(theTime time.Time) {
	cb, err := clipboard.ReadAll()
	// Sometimes the handle will just give out, there's little harm in restarting our loop with the delay
	if err != nil {
		return
	}
	lines := strings.Split(cb, "\n")
	finishedString := ""

	// Unelegant confirmations that we're probably reading the right thing
	if len(lines) != 7 {
		return
	}
	if !strings.Contains(lines[0], "An unstable wormhole") {
		return
	}

	finishedString += "LIFE: " + checkLife(lines[4]) + " | "
	finishedString += "MASS: " + checkMass(lines[5]) + " | "
	finishedString += "SIZE: " + checkShips(lines[6])

	clipboard.WriteAll(finishedString)

	fmt.Printf("[%s]: %s\n", theTime.Format("03:04 PM"), finishedString)
}

func checkLife(line string) string {
	if strings.Contains(line, "not yet begun") {
		return ">24 hrs"
	}
	if strings.Contains(line, "beginning to decay") {
		return "4-24 hrs"
	}
	if strings.Contains(line, "reaching the end") {
		return "<4hrs"
	}
	if strings.Contains(line, "on the verge") {
		return "<15m"
	}
	return ""
}

func checkMass(line string) string {
	if strings.Contains(line, "not yet") {
		return ">50%"
	}
	if strings.Contains(line, "not to a critical degree") {
		return "10-50%"
	}
	if strings.Contains(line, "stability critically disrupted") {
		return "<10%"
	}
	return ""
}

func checkShips(line string) string {
	if strings.Contains(line, "Very large") {
		return "Very Large"
	}
	if strings.Contains(line, "Larger ships") {
		return "Large"
	}
	if strings.Contains(line, "medium size") {
		return "Medium"
	}
	if strings.Contains(line, "smallest ships") {
		return "Smallest"
	}
	return ""
}
