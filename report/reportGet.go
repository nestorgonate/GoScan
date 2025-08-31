package report

import (
	scanTool "GoScan/scanner"
	"fmt"
)

type IReport interface {
	GetMinecraftFileReport() ([]string, error)
}
type Report struct {
	hackName []string
	scan     scanTool.IScanFiles
}

func NewReportGet(scan scanTool.IScanFiles) *Report {
	return &Report{
		hackName: []string{},
		scan:     scan,
	}
}

func (r *Report) GetMinecraftFileReport() ([]string, error) {
	r.hackName = append(r.hackName,
		"Fly",
		"KillAura",
		"AutoClicker",
		"X-Ray",
		"NoClip",
		"Speed",
		"ESP",
		"WallHack",
		"TriggerBot",
		"FastPlace",
	)
	filesToScan, err := r.scan.ScanFilesMinecraftFunc()
	if err != nil {
		fmt.Printf("Error al escanear la carpeta .minecraft: %v", err)
		return nil, err
	}
	fmt.Printf("filesToScan en GetMinecraftFileReport: %v", filesToScan)
	return filesToScan, nil
}
