package report

import (
	scanTool "GoScan/scanner"
	"fmt"
	"strings"
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
		"fly",
		"KillAura",
		"AutoClicker",
		"X-Ray",
		"NoClip",
		"Speed",
		"ESP",
		"WallHack",
		"TriggerBot",
		"FastPlace",
		"Wrust",
		"X-ray",
		"Xray",
	)
	dirsToScan, filesToScan, err := r.scan.ScanFilesMinecraftFunc()
	if err != nil {
		fmt.Printf("Error al escanear la carpeta .minecraft: %v", err)
		return nil, err
	}
	//Iterar carpetas
	for _, dir := range dirsToScan{
		for _, hack := range r.hackName{
			if strings.EqualFold(hack, dir){
				fmt.Printf("Hack encontrado en carpetas: %v\n", dir)
			}
		}
		fmt.Printf("Carpeta encontrada: %v\n", dir)
	}

	//Iterar archivos
	for _, file := range filesToScan{
		for _, hack := range r.hackName{
			if strings.EqualFold(hack, file){
				fmt.Printf("Hack encontrado en archivo: %v\n", file)
			}
		}
		fmt.Printf("Archivo encontrado: %v\n", file)
	}
	return dirsToScan, nil
}
