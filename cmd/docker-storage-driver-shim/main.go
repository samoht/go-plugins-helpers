package main

import (
	"flag"
	"fmt"
	"os"

	upstream "github.com/docker/docker/daemon/graphdriver"

	"github.com/docker/docker/daemon/graphdriver/aufs"
	"github.com/docker/docker/daemon/graphdriver/btrfs"
	"github.com/docker/docker/daemon/graphdriver/devmapper"
	"github.com/docker/docker/daemon/graphdriver/overlay"
	"github.com/docker/docker/daemon/graphdriver/vfs"
	"github.com/docker/docker/daemon/graphdriver/zfs"
	"github.com/docker/go-plugins-helpers/graphdriver/shim"
)

var (
	driver = flag.String("driver", "aufs", "The storage driver to use")
)

func main() {
	var Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	var init upstream.InitFunc

	switch *driver {
	case "aufs":
		init = aufs.Init
	case "btrfs":
		init = btrfs.Init
	case "zfs":
		init = zfs.Init
	case "devicemapper":
		init = devmapper.Init
	case "overlay":
		init = overlay.Init
	case "vfs":
		init = vfs.Init
	default:
		Usage()
		os.Exit(1)
	}
	h := shim.NewHandlerFromGraphDriver(init)
	h.ServeUnix("root", "storage-shim")

}
