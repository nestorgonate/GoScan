package report

import (
	scanTool "GoScan/scanner"
	"fmt"
	"strings"
)

type IReport interface {
	GetMinecraftFileReport()
	GetPrefetchFileReport()
}
type Report struct {
	hackName []string
	scan     scanTool.IScanFiles
	hackVerified []string
	suspiciousFile []string
}

func NewReportGet(scan scanTool.IScanFiles) *Report {
	return &Report{
		scan:     scan,
	}
}

func (r *Report) GetMinecraftFileReport(){
	r.hackName = append(r.hackName,
		"fly",
		"KillAura",
		"AutoClicker",
		"XRay",
		"NoClip",
		"Speed",
		"ESP",
		"WallHack",
		"TriggerBot",
		"FastPlace",
		"Wrust",
		"autoclick",
	)
	dirsToScan, filesToScan, err := r.scan.ScanFilesMinecraftFunc()
	if err != nil {
		fmt.Printf("The .minecraft was not scanned: %v", err)
		return
	}
	fmt.Printf("Doing the .minecraft scan...\n")
	//Iterar carpetas
	for _, dir := range *dirsToScan{
		for _, hack := range r.hackName{
			if strings.Contains(strings.ToLower(dir), strings.ToLower(hack)){
				r.hackVerified = append(r.hackVerified, dir)
			}
		}
	}

	//Iterar archivos
	for _, file := range *filesToScan{
		for _, hack := range r.hackName{
			if strings.Contains(strings.ToLower(file.(string)), strings.ToLower(hack)){
				r.hackVerified = append(r.hackVerified, file.(string))
			}
		}
	}

	if len(r.hackVerified) > 0{
		fmt.Printf("Hacks found during the .minecraft scan:\n")
		for _, hack := range r.hackVerified{
			fmt.Printf("%v\n", hack)
		}
	}
	fmt.Printf("No cheats were found during the scan\n")
	fmt.Printf("Analayzed files:\n")
	for _, file := range *filesToScan{
		fmt.Printf("%v\n", file)
	}
	fmt.Printf("=====The .minecraft is complete=====\n")
}

func (r *Report) GetPrefetchFileReport(){
	r.hackVerified = nil
	r.suspiciousFile = append(r.suspiciousFile, "Logitech", "Razer", "Autoclick", "Chrome")
	prefetchScan, err := r.scan.ScanPrefetch()
	if err != nil{
		fmt.Printf("The prefetch file was not scanned: %v\n", err)
		return
	}
	fmt.Printf("Doing the prefetch scan...\n")
	for _, file := range *prefetchScan{
		for _, suspicious := range r.suspiciousFile{
			if strings.Contains(strings.ToLower(file.Name), strings.ToLower(suspicious)){
				r.hackVerified = append(r.hackVerified, file.Name)
			}
		}
	}
	if len(r.hackVerified) > 0{
		for _, hack := range r.hackVerified{
			fmt.Printf("Suspicious file was found during the scan: %v\n", hack)
		}
	}
	fmt.Printf("No suspicious files were found during the scan\n")
	fmt.Printf("Analyzed files:\n")
	for _, file := range *prefetchScan{
		fmt.Printf("%v\n", file)
	}
	fmt.Printf("=====The prefetch scan is complete=====\n")
}
