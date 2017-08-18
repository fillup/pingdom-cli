package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

// UptimeCommand - Struct to attach cli.Command to for uptime subcommand
type UptimeCommand struct {
}

// Run -- Run uptime command
func (c *UptimeCommand) Run(args []string) int {
	var periodName string
	var tags string

	// Parse flags
	cmdFlags := flag.NewFlagSet("uptime", flag.ContinueOnError)
	cmdFlags.StringVar(&periodName, "period", "Today", fmt.Sprintf("Period for report: %s", strings.Join(c.GetPeriodOptions(), ", ")))
	cmdFlags.StringVar(&tags, "tags", "", "Comma-separated list of tags. Ex: tag1,tag2")
	cmdFlags.Usage = func() { c.Help() }
	if err := cmdFlags.Parse(args); err != nil {
		c.Help()
		return 1
	}

	uptimeByCheck := map[string]string{}

	// cmdargs := ArgMapFromArray(args, c.GetArgMap())

	checksList := ListChecks(tags)

	timePeriod := GetPeriodByName(periodName, 0)
	periodFrom, periodTo := timePeriod.String()

	fmt.Println("\nPingdom Uptime Report")
	fmt.Printf("Reporting Period (%s):\n", periodName)
	fmt.Printf("From: %s\n", periodFrom)
	fmt.Printf("To:   %s\n", periodTo)
	fmt.Println("")

	const padding = 2
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)
	fmt.Fprintln(w, "Check\tUptime\t")
	for _, v := range checksList.Checks {
		summaryAverage := GetSummaryAverage(v.ID, timePeriod.from, timePeriod.to)
		uptimeByCheck[v.Name] = CalcUptimePercent(summaryAverage.Summary.Status.Totaldown, summaryAverage.Summary.Status.Totalup)
		fmt.Fprintln(w, fmt.Sprintf("%s\t%s\t", v.Name, uptimeByCheck[v.Name]))
	}

	w.Flush()
	fmt.Println("")
	return 0
}

// Help - Return Help information
func (c *UptimeCommand) Help() string {
	helpText := `
Usage: pingdom uptime [options]
  Runs uptime report on Pingdom for specified period on all checks tagged
  with given tags.
Options:
  -period PeriodName      Default: Today
                          Options: %s
  -tags tag1,tag2         If specified, only include checks with one of these tags
`
	return strings.TrimSpace(fmt.Sprintf(helpText, strings.Join(c.GetPeriodOptions(), ", ")))
}

// Synopsis - Return summary for command
func (c *UptimeCommand) Synopsis() string {
	return "Generate uptime report based on provided arguments"
}

// CalcUptimePercent - Do math to calculate uptime as a percentage
func CalcUptimePercent(down int, up int) string {
	return fmt.Sprintf("%.3f%%", 100-(float32(down)/float32(up))*100)
}

// GetPeriodOptions - get list of options for -period flag
func (c *UptimeCommand) GetPeriodOptions() []string {
	return []string{
		"Today",
		"Yesterday",
		"ThisWeek",
		"LastWeek",
		"ThisMonth",
		"LastMonth",
		"ThisYear",
		"LastYear",
	}
}
