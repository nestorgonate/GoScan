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
	ScanFilesMinecraftFunc() ([]string, error)
}
type ScanFiles struct {
	utils               *utils.Utils
	dirsToGet           []string
	filesExtensionToGet []string
	filesToScan         []string
	re                  regexp.Regexp
}

func NewScanFiles() *ScanFiles {
	return &ScanFiles{
		//Slice vacio
		dirsToGet:           []string{},
		filesExtensionToGet: []string{},
		re:                  *regexp.MustCompile(`[^\\/:]+\.jar$`),
	}
}

func (r *ScanFiles) ScanFilesMinecraftFunc() ([]string, error) {
	//Agregar al slices los archivos de interes
	r.dirsToGet = append(r.dirsToGet, "resourcepacks", "versions", "mods")
	r.filesExtensionToGet = append(r.filesExtensionToGet, "*.jar", "*.zip", "*.dll", "*.exe")

	fmt.Printf("Las carpetas en .minecraft a escanear: %v\n", r.dirsToGet)
	fmt.Printf("Las extensiones de archivos en .minecraft a escanear: %v\n", r.filesExtensionToGet)
	//Obtener la ruta del .minecraft en base al sistema operativo
	minecraftPath, err := r.utils.GetMinecraftPath()
	if err != nil {
		fmt.Printf("Error al obtener el path de Minecraft\n")
		return nil, err
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
						//Si es un archivo verificar si la extension es de interes para escanear
						if !d.IsDir() {
							//filepath.ext retorna al extension del archivo que obtiene de d.Name que retorna el nombre completo del archivo o carpeta
							extension := strings.ToLower(filepath.Ext(d.Name()))
							for _, fileExtension := range r.filesExtensionToGet {
								if strings.TrimPrefix(fileExtension, "*") == extension {
									cleanNameFile := r.re.FindString(path)
									if cleanNameFile != "" {
										r.filesToScan = append(r.filesToScan, cleanNameFile)
										fmt.Printf("Archivo a escanear: %v\n", cleanNameFile)
									}
									break
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
			//Si el directorio es la raiz o no esta en dirsToGet no escanearlo
			if !dirIsInDirsToGet && path != minecraftPath {
				return filepath.SkipDir
			}
			return nil
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error durante el recorrido del directorio: %v", err)
		return nil, err
	}
	return r.filesToScan, nil
}
