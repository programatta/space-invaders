//go:build !js && !wasm

package platform

import "os"

func ExitGame() {
	os.Exit(0)
}
