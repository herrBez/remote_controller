// +build !target_linux target_windows

package platform_specific

import (
	"fmt"
	"os/exec"
	"github.com/herrBez/remote_controller/core"
)

// https://docs.microsoft.com/en-us/windows/win32/inputdev/virtual-key-codes
var vk map[string]string = map[string]string {
	",": "0xBC",
	" ": "0x20",
}

// Backend function that increases the volume of the computer of the specified
// percent value. 65535 is the maximal value.
func IncreaseVolume(percent int) {
	delta := 65535*percent/100
	parameter := fmt.Sprintf("%d", delta)
	cmd := exec.Command("cmd", "/C", "nircmd.exe", "changesysvolume", parameter)
	core.ExecuteCommand(cmd)
}

func SetAbsoluteVolume(percent int) {
	fmt.Println("Set absolute volume")
	delta := 65535*percent/100
	parameter := fmt.Sprintf("%d", delta)
	cmd := exec.Command("cmd", "/C", "nircmd.exe", "setsysvolume", parameter)
	core.ExecuteCommand(cmd)
}

func MuteVolume() {
  parameter := fmt.Sprintf("%d", 2)
  cmd := exec.Command("cmd", "/C", "nircmd.exe", "mutesysvolume", parameter)
  core.ExecuteCommand(cmd)
}

func Pause() {
	cmd := exec.Command("cmd", "/C", "nircmd.exe", "sendkey", "0x20", "press")
	core.ExecuteCommand(cmd)
}

func SendSingleKeyBoardInput(key string) {
	symbol, ok := vk[key]
	fmt.Println(string(symbol) + "//" + string(key))
	if !ok {
		fmt.Println(key + " not present")
	} else {
		fmt.Println(key + " present")
		key = symbol
	}
	fmt.Println(string(symbol) + "==" + string(key))
	cmd := exec.Command("cmd", "/C", "nircmd.exe", "sendkey", key, "press")
	core.ExecuteCommand(cmd)
}


func SendKeyBoardInput(key string) {
	fmt.Println("I am going to print '" + key + "'")
	var i = 0
	for i = 0; i < len(key); i++ {
		SendSingleKeyBoardInput(string(key[i]))
	}
}

func SwitchApp() {
	cmd := exec.Command("cmd", "/C", "nircmd.exe", "sendkeypress", "alt+tab")
	core.ExecuteCommand(cmd)
}

// Move the cursor into another position and back to original position
func MoveCursor() {
	cmd1 := exec.Command("cmd", "/C", "nircmd.exe", "movecursor", "10", "10")
	cmd2 := exec.Command("cmd", "/C", "nircmd.exe", "movecursor", "-10", "-10")
	core.ExecuteCommand(cmd1)
	core.ExecuteCommand(cmd2)
}