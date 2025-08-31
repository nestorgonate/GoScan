package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type Utils struct {
}

func NewUtils() *Utils {
	return &Utils{}
}

func (r *Utils) GetMinecraftPath() (string, error) {
	baseDir, err := os.UserHomeDir()
	if err != nil{
		return "", nil
	}
	//runtime.GOOS obtiene el nombre del sistema operativo en el cual se ejcuta el programa
	switch runtime.GOOS {
	case "windows":
		return filepath.Join(baseDir, "AppData", "Roaming", ".minecraft"), nil
	case "linux":
		return filepath.Join(baseDir, ".minecraft"), nil
	case "darwin": // macOS
		return filepath.Join(baseDir, "Library", "Application Support", "minecraft"), nil
	default:
		return "", fmt.Errorf("SO no soportado: %s", runtime.GOOS)
	}
}