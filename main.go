package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/atotto/clipboard"
)

func main() {
	for {
		clipboardLoop()
		time.Sleep(500 * time.Millisecond)
	}
}

func clipboardLoop() {
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

	finishedString += checkLife(lines[4]) + " / "
	finishedString += checkMass(lines[5]) + " / "
	finishedString += checkShips(lines[6])

	clipboard.WriteAll(finishedString)
	now := time.Now().UTC()
	fmt.Printf("[%s]: %s\n", now.Format("03:04 PM"), finishedString)
}

func checkLife(line string) string {
	if strings.Contains(line, "not yet begun") {
		return "Not yet"
	}
	if strings.Contains(line, "beginning to decay") {
		return "Beginning"
	}
	if strings.Contains(line, "reaching the end") {
		return "Reaching the end"
	}
	if strings.Contains(line, "on the verge") {
		return "On the verge"
	}
	return ""
}

func checkMass(line string) string {
	if strings.Contains(line, "not yet") {
		return "Not yet (> 50%)"
	}
	if strings.Contains(line, "not to a critical degree") {
		return "Not to critical degree (10-50%)"
	}
	if strings.Contains(line, "stability critically disrupted") {
		return "Critical (< 10%)"
	}
	return ""
}

func checkShips(line string) string {
	if strings.Contains(line, "Very large") {
		return "Very Large"
	}
	if strings.Contains(line, "Larger ships") {
		return "Larger"
	}
	if strings.Contains(line, "medium size") {
		return "Medium"
	}
	if strings.Contains(line, "smallest ships") {
		return "Smallest"
	}
	return ""
}
