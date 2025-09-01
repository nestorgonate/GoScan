package scanTool

import (
	"GoScan/utils"
	"fmt"
	"io/fs"
	"path/filepath"
	"regexp"
	"strings"
)

type IScanFiles interface {
	ScanFilesMinecraftFunc() ([]string, []string, error)
}
type ScanFiles struct {
	utils *utils.Utils
	//Indica el nombre de las carpetas de interes en el .minecraft
	dirsToGet []string
	//Almacena el nombre de las carpetas a escanear
	dirsToScan []string
	re          regexp.Regexp
	//Indica la extension de los archivos de interes en el .minecraft
	filesToGet []string
	//Almacena el nombre de los archivos a escanear
	filesToScan []string
}

func NewScanFiles() *ScanFiles {
	return &ScanFiles{
		//Slice vacio
		re:         *regexp.MustCompile(`^(.+)\.[^\.]+$`),
	}
}

func (r *ScanFiles) ScanFilesMinecraftFunc() ([]string, []string, error) {
	//Agregar al slice los archivos de interes
	r.dirsToGet = append(r.dirsToGet, "resourcepacks", "versions", "mods")
	fmt.Printf("Las carpetas en .minecraft a escanear: %v\n", r.dirsToGet)
	//Agregar al slice las extensiones de archivos de interes
	r.filesToGet = append(r.filesToGet, ".jar", ".exe", ".dll")
	//Obtener la ruta del .minecraft en base al sistema operativo
	minecraftPath, err := r.utils.GetMinecraftPath()
	if err != nil {
		fmt.Printf("Error al obtener el path de Minecraft\n")
		return nil, nil, err
	}
	fmt.Printf("Path de minecraft: %v\n", minecraftPath)

	//Iterar la carpeta .minecraft con carpetas seleccionadas
	err = filepath.WalkDir(minecraftPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Error al acceder a la carpeta .minecraft %q: %v\n", path, err)
			return err
		}
		//Si es una carpeta
		if d.IsDir() {
			nameDir := d.Name()
			//Verificar si el directorio de .minecraft es de interes para escanear
			dirIsInDirsToGet := false
			for _, dir := range r.dirsToGet {
				if strings.EqualFold(dir, nameDir) {
					dirIsInDirsToGet = true
					fmt.Printf("Carpeta a escanear: %v\n", dir)
					//Iterar la carpeta de interes, al ser ya una carpeta de interes no verifica un nombre en especifico
					err = filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
						if d.IsDir() {
								r.dirsToScan = append(r.dirsToScan, d.Name())
								fmt.Printf("Slice dirsToScan: %v\n", r.dirsToScan)
							}else{
								//filepath.ext retorna al extension del archivo que obtiene de d.Name que retorna el nombre completo del archivo o carpeta
								fileExt := strings.ToLower(filepath.Ext(d.Name()))
								for _, ext := range r.filesToGet{
									if strings.TrimPrefix(ext, "*") == fileExt{
										cleanedNameFile := r.re.FindStringSubmatch(d.Name())
										if len(cleanedNameFile) > 1{
											r.filesToScan = append(r.filesToScan, cleanedNameFile[1])
										}
									}
								}
							}
						return nil
					})
					if err != nil {
						fmt.Printf("Error al acceder a la carpeta .minecraft %q: %v\n", path, err)
						return nil
					}
					break
				}
			}
			//Si el directorio no es el path de minecraft y no esta en dirsToGet no escanear
			if !dirIsInDirsToGet && path != minecraftPath {
				return filepath.SkipDir
			}
			return nil
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error durante el recorrido del directorio: %v", err)
		return nil, nil, err
	}
	return r.dirsToScan, r.filesToScan, nil
}
