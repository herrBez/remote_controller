// +build target_linux,!target_windows

package platform_specific

import (
	"fmt"
	"os/exec"
	"github.com/herrBez/remote_controller/core"
)


// Backend function that increases the volume of the computer of the specified
// percent value. 65535 is the maximal value.
func IncreaseVolume(percent int) {
	fmt.Println("Increase Volume")
}

func SetAbsoluteVolume(percent int) {
	fmt.Println("Set absolute volume")
}

func MuteVolume() {
	fmt.Println("Mute volume")
	cmd := exec.Command("echo", "hello")
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



