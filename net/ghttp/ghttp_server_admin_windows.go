//go:build windows
// +build windows

package ghttp

import (
	"context"
	"os"

	"github.com/xrc360/golang/os/gproc"
)

// handleProcessSignal handles all signals from system in blocking way.
func handleProcessSignal() {
	var ctx = context.TODO()
	gproc.AddSigHandlerShutdown(func(sig os.Signal) {
		shutdownWebServersGracefully(ctx, sig)
	})

	gproc.Listen()
}
