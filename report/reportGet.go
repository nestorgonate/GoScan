package report

import (
	scanTool "GoScan/scanner"
	"fmt"
	"strings"
)

type IReport interface {
	GetMinecraftFileReport()
}
type Report struct {
	hackName []string
	scan     scanTool.IScanFiles
	hackVerified []string
}

func NewReportGet(scan scanTool.IScanFiles) *Report {
	return &Report{
		scan:     scan,
	}
}

func (r *Report) GetMinecraftFileReport() {
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
		fmt.Printf("Error al escanear la carpeta .minecraft: %v", err)
	}
	fmt.Printf("Generando reporte del .minecraft...\n")
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
			if strings.Contains(strings.ToLower(file), strings.ToLower(hack)){
				r.hackVerified = append(r.hackVerified, file)
			}
		}
	}

	if len(r.hackVerified) > 0{
		fmt.Printf("Hacks encontrados mediante un escaneo de nombres genericos")
		for _, hack := range r.hackVerified{
			fmt.Printf("Hack encontrado en la carpeta .minecraft: %v", hack)
		}
	}
	fmt.Printf("No se han encontrado hacks en archivos y carpetas de .minecraft mediante un escaneo de nombres genericos\n")
	fmt.Printf("=====Escaneo de .minecraft terminado=====")
}
