package main

import (
	"GoScan/report"
	scanTool "GoScan/scanner"
)

func main() {
	scan := scanTool.NewScanFiles()
	report := report.NewReportGet(scan)
	report.GetMinecraftFileReport()
}