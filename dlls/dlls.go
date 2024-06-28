package dlls

import (
	"fmt"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

func GetActiveWindowTitle() (string, error) {
	//hwnd := win.GetForegroundWindow()
	user32 := syscall.NewLazyDLL("user32.dll")
	// pega a função "GetForegroundWindow" na dll
	procGetForegroundWindow := user32.NewProc("GetForegroundWindow")
	hwnd, _, _ := procGetForegroundWindow.Call()

	if hwnd == 0 {
		return "", fmt.Errorf("nenhuma janela ativa encontrada")
	}
	return getWindowText(hwnd), nil
}

func getWindowText(hwnd uintptr) string {
	// Define o tamanho máximo do título da janela
	const maxCount = 255
	buf := make([]uint16, maxCount)

	// Carrega a função GetWindowTextW da DLL user32
	user32 := syscall.NewLazyDLL("user32.dll")
	getWindowText := user32.NewProc("GetWindowTextW")

	// Chama a função GetWindowTextW
	_, _, _ = getWindowText.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&buf[0])), uintptr(maxCount))

	// Converte UTF-16 para string
	return string(utf16.Decode(buf))
}