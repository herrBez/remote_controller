// +build target_linux !target_windows

package platform_specific

import (
	"fmt"
)

func Foo() {
	fmt.Println("Linux")
}