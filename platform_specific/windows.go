// +build !target_linux target_windows

package platform_specific

import (
	"fmt"
	"os"
	"os/exec"
	"github.com/herrBez/remote_controller/core"
)

// Backend function that increases the volume of the computer of the specified
// percent value. 65535 is the maximal value.
func IncreaseVolume(percent int) {
	delta := 65535*percent/100
	parameter := fmt.Sprintf("%d", delta)
	cmd := exec.Command("cmd", "/C", "nircmd.exe", "changesysvolume", parameter)
	working_directory, _ := os.Getwd()
	cmd.Dir = working_directory
	core.ExecuteCommand(cmd)
}

func MuteVolume() {
  parameter := fmt.Sprintf("%d", 2)
  cmd := exec.Command("cmd", "/C", "nircmd.exe", "mutesysvolume", parameter)
  core.ExecuteCommand(cmd)
}

func Pause() {
	parameter := fmt.Sprintf("0x20")
	cmd := exec.Command("cmd", "/C", "nircmd.exe", "sendkey", parameter, "press")
	core.ExecuteCommand(cmd)
}

func SendKeyBoardInput(key string) {
	cmd := exec.Command("cmd", "/C", "nircmd.exe", "sendkey", key, "press")
	core.ExecuteCommand(cmd)
}

func SwitchApp() {
	cmd := exec.Command("cmd", "/C", "nircmd.exe", "sendkeypress", "alt+tab")
	core.ExecuteCommand(cmd)
}