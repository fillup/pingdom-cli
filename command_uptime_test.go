package main

import (
	"strings"
	"testing"
)

func TestHelp(t *testing.T) {
	c := UptimeCommand{}
	help := c.Help()
	if !strings.Contains(help, "Usage: pingdom uptime") {
		t.Error("Help did not return a string")
	}
}

func TestSynopsys(t *testing.T) {
	c := UptimeCommand{}
	help := c.Synopsis()
	if !strings.Contains(help, "Generate uptime report") {
		t.Error("Synopsis did not return a string")
	}
}

func TestCalcUptimePercent(t *testing.T) {
	uptime := CalcUptimePercent(10, 90)
	if uptime != "88.889%" {
		t.Errorf("Uptime was not calculated as expected. Got: %s", uptime)
	}
}
