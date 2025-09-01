package main

import (
	"GoScan/report"
	scanTool "GoScan/scanner"
	"GoScan/utils"
)

func main() {
	utils := utils.NewUtils()
	
	amAdmin := utils.AmAdmin()
	if !amAdmin{
		utils.RunWithAdmin()
	}

	scan := scanTool.NewScanFiles()
	report := report.NewReportGet(scan)
	reports := []func(){
		report.GetMinecraftFileReport,
		report.GetPrefetchFileReport,
	}

	for i, report := range reports{
		report()
		if i < len(reports)-1{
			utils.WaitUntilEnter("Press enter to continue the scan")
		}
	}

	defer func() {
		utils.LogMemoryUsage()
		utils.WaitUntilEnter("Press enter to exit")
	}()
}