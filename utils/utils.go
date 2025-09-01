package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"

	"golang.org/x/sys/windows"
)

type Utils struct {
}

func NewUtils() *Utils {
	return &Utils{}
}

func (r *Utils) GetMinecraftPath() (string, error) {
	baseDir, err := os.UserHomeDir()
	if err != nil {
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

func (r *Utils) AmAdmin() bool {
	var sid *windows.SID
	//SID de administrador S-1-5-32-544
	sid, err := windows.CreateWellKnownSid(windows.WinBuiltinAdministratorsSid)
	if err != nil {
		return false
	}
	//Token del proceso acutal 0
	token := windows.Token(0)
	isMember, err := token.IsMember(sid)
	if err != nil {
		return false
	}
	return isMember
}

func (r *Utils) RunWithAdmin() {
	verb := "runas"
	exe, _ := os.Executable()
	cwd, _ := os.Getwd()
	args := strings.Join(os.Args[1:], " ")

	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	var showCmd int32 = 1 //SW_NORMAL

	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
	if err != nil {
		fmt.Println(err)
	}
}

func (r *Utils) WaitUntilEnter(msg string){
	fmt.Println(msg)
	fmt.Scanln()
}
