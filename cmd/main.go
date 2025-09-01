package main

import (
	"GoScan/report"
	scanTool "GoScan/scanner"
	"GoScan/utils"
)

func main() {
	utils := utils.NewUtils()
	scan := scanTool.NewScanFiles()
	report := report.NewReportGet(scan)
	report.GetMinecraftFileReport()
	defer func() {
		utils.LogMemoryUsage()
	}()
}