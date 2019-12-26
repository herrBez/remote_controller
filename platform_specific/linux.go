// +build target_linux,!target_windows

package platform_specific

import (
	"fmt"
	"os/exec"
	"github.com/herrBez/remote_controller/core"
)


/// Syntax: amixer set Master <percent>%<+|->
func IncreaseVolume(percent int) {
	fmt.Println("Increase Volume")
	append_symbol := "+"
	if percent < 0 {
		append_symbol = "-"
	}
	parameter := fmt.Sprintf("%d%s%s", percent, "%", append_symbol)
	cmd := exec.Command("amixer", "set", "Master", parameter)
	core.ExecuteCommand(cmd)
}

func SetAbsoluteVolume(percent int) {
	fmt.Println("Set absolute volume")
	parameter := fmt.Sprintf("%d%s", percent,"%")
	fmt.Println(parameter)
	cmd := exec.Command("amixer", "set", "Master", parameter)
	core.ExecuteCommand(cmd)
}

func MuteVolume() {
	fmt.Println("Mute volume")
	cmd := exec.Command("amixer", "Master", "toggle")
	core.ExecuteCommand(cmd)
}


func Pause() {
	fmt.Println("Pause")
	cmd := exec.Command("xdotool", "key", "space")
	core.ExecuteCommand(cmd)
}

func SendSingleKeyBoardInput(key string) {
	fmt.Println("SendSingleKeyBoardInput")
}

// Usage: xdotool key [A-Z ,;](\s+[A-Z ,;])*
func SendKeyBoardInput(key string) {
	fmt.Println("SendKeyBoardInput")
	fmt.Println("I am going to print '" + key + "'")
	var i = 0
	n := len(key)
	parameter := make([]string, n+1)
	parameter[0] = "key"
	for i = 0; i < n; i++ {
		switch key[i] {
		case 0x20:
			parameter[i+1] = "space"
		default:
			parameter[i+1] = string(key[i])
		}
	}
	fmt.Println(parameter)
	cmd := exec.Command("xdotool", parameter...)
	core.ExecuteCommand(cmd)
}

func SwitchApp() {
	fmt.Println("SwitchApp")
	cmd := exec.Command("xdotool", "key", "alt+Tab")
	core.ExecuteCommand(cmd)
}

// Move the cursor into another position and back to original position
func MoveCursor() {
	fmt.Println("MoveCursor")
	cmd := exec.Command("xdotool", "mousemove", "10", "10", "mousemove", "restore")
	core.ExecuteCommand(cmd)
}



