package main

import (
	"github.com/containerd/containerd/cmd/containerd/app"
	_ "github.com/containerd/containerd/cmd/containerd/builtins"

	// custom plugins: aufs, zfs
	_ "github.com/containerd/aufs"
	_ "github.com/containerd/zfs"
)

func main() {
	app.Main()
}
