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
}

func SendSingleKeyBoardInput(key string) {
	fmt.Println("SendSingleKeyBoardInput")
}


func SendKeyBoardInput(key string) {
	fmt.Println("SendSingleKeyBoardInput")

}

func SwitchApp() {
	fmt.Println("SwitchApp")

}

// Move the cursor into another position and back to original position
func MoveCursor() {
	fmt.Println("MoveCursor")

}



