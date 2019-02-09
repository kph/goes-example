// Copyright © 2015-2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package main

import (
	"github.com/platinasystems/goes"
	"github.com/platinasystems/goes/cmd"
	"github.com/platinasystems/goes/cmd/bang"
	"github.com/platinasystems/goes/cmd/cat"
	"github.com/platinasystems/goes/cmd/cd"
	"github.com/platinasystems/goes/cmd/chmod"
	"github.com/platinasystems/goes/cmd/cli"
	"github.com/platinasystems/goes/cmd/cmdline"
	"github.com/platinasystems/goes/cmd/cp"
	"github.com/platinasystems/goes/cmd/daemons"
	"github.com/platinasystems/goes/cmd/dmesg"
	"github.com/platinasystems/goes/cmd/echo"
	"github.com/platinasystems/goes/cmd/elsecmd"
	"github.com/platinasystems/goes/cmd/env"
	"github.com/platinasystems/goes/cmd/exec"
	"github.com/platinasystems/goes/cmd/exit"
	"github.com/platinasystems/goes/cmd/export"
	"github.com/platinasystems/goes/cmd/falsecmd"
	"github.com/platinasystems/goes/cmd/femtocom"
	"github.com/platinasystems/goes/cmd/ficmd"
	"github.com/platinasystems/goes/cmd/function"
	"github.com/platinasystems/goes/cmd/grub"
	"github.com/platinasystems/goes/cmd/hdel"
	"github.com/platinasystems/goes/cmd/hdelta"
	"github.com/platinasystems/goes/cmd/hexists"
	"github.com/platinasystems/goes/cmd/hget"
	"github.com/platinasystems/goes/cmd/hgetall"
	"github.com/platinasystems/goes/cmd/hkeys"
	"github.com/platinasystems/goes/cmd/hset"
	"github.com/platinasystems/goes/cmd/hwait"
	"github.com/platinasystems/goes/cmd/ifcmd"
	"github.com/platinasystems/goes/cmd/iminfo"
	"github.com/platinasystems/goes/cmd/insmod"
	"github.com/platinasystems/goes/cmd/install"
	"github.com/platinasystems/goes/cmd/ip"
	"github.com/platinasystems/goes/cmd/kexec"
	"github.com/platinasystems/goes/cmd/keys"
	"github.com/platinasystems/goes/cmd/kill"
	"github.com/platinasystems/goes/cmd/ln"
	"github.com/platinasystems/goes/cmd/ls"
	"github.com/platinasystems/goes/cmd/lsmod"
	"github.com/platinasystems/goes/cmd/mkdir"
	"github.com/platinasystems/goes/cmd/mknod"
	"github.com/platinasystems/goes/cmd/mount"
	"github.com/platinasystems/goes/cmd/ping"
	"github.com/platinasystems/goes/cmd/ps"
	"github.com/platinasystems/goes/cmd/pwd"
	"github.com/platinasystems/goes/cmd/reboot"
	"github.com/platinasystems/goes/cmd/redisd"
	"github.com/platinasystems/goes/cmd/reload"
	"github.com/platinasystems/goes/cmd/restart"
	"github.com/platinasystems/goes/cmd/rm"
	"github.com/platinasystems/goes/cmd/rmmod"
	"github.com/platinasystems/goes/cmd/slashinit"
	"github.com/platinasystems/goes/cmd/sleep"
	"github.com/platinasystems/goes/cmd/source"
	"github.com/platinasystems/goes/cmd/start"
	"github.com/platinasystems/goes/cmd/stop"
	"github.com/platinasystems/goes/cmd/stty"
	"github.com/platinasystems/goes/cmd/subscribe"
	"github.com/platinasystems/goes/cmd/sync"
	"github.com/platinasystems/goes/cmd/testcmd"
	"github.com/platinasystems/goes/cmd/thencmd"
	"github.com/platinasystems/goes/cmd/truecmd"
	"github.com/platinasystems/goes/cmd/umount"
	"github.com/platinasystems/goes/cmd/uninstall"
	"github.com/platinasystems/goes/cmd/uptimed"
	"github.com/platinasystems/goes/cmd/wget"
	"github.com/platinasystems/goes/lang"
)

var Goes = &goes.Goes{
	NAME: "goes-" + name,
	APROPOS: lang.Alt{
		lang.EnUS: "an example goes machine",
	},
	ByName: map[string]cmd.Cmd{
		"!":        bang.Command{},
		"cli":      &cli.Command{},
		"cat":      cat.Command{},
		"cd":       &cd.Command{},
		"chmod":    chmod.Command{},
		"cp":       cp.Command{},
		"daemons":  daemons.Status{},
		"dmesg":    dmesg.Command{},
		"echo":     echo.Command{},
		"else":     &elsecmd.Command{},
		"env":      &env.Command{},
		"exec":     exec.Command{},
		"exit":     exit.Command{},
		"export":   export.Command{},
		"false":    falsecmd.Command{},
		"femtocom": femtocom.Command{},
		"fi":       &ficmd.Command{},
		"function": &function.Command{},
		"goes-daemons": &daemons.Server{
			Init: [][]string{
				[]string{"redisd"},
				[]string{"uptimed"},
			},
		},
		"grub":    &grub.Command{},
		"hdel":    hdel.Command{},
		"hdelta":  &hdelta.Command{},
		"hexists": hexists.Command{},
		"hget":    hget.Command{},
		"hgetall": hgetall.Command{},
		"hkeys":   hkeys.Command{},
		"hset":    hset.Command{},
		"hwait":   hwait.Command{},
		"if":      &ifcmd.Command{},
		"insmod":  insmod.Command{},
		"install": &install.Command{},
		"ip":      ip.Goes,
		"kexec":   kexec.Command{},
		"keys":    keys.Command{},
		"kill":    kill.Command{},
		"ln":      ln.Command{},
		"log":     daemons.Log{},
		"ls":      ls.Command{},
		"lsmod":   lsmod.Command{},
		"mkdir":   mkdir.Command{},
		"mknod":   mknod.Command{},
		"mount":   mount.Command{},
		"ping":    ping.Command{},
		"ps":      ps.Command{},
		"pwd":     pwd.Command{},
		"reboot":  reboot.Command{},
		"redisd": &redisd.Command{
			Devs:    []string{"lo"},
			Machine: "example",
		},
		"reload":  reload.Command{},
		"restart": &restart.Command{},
		"rm":      rm.Command{},
		"rmmod":   rmmod.Command{},
		"show": &goes.Goes{
			NAME:  "show",
			USAGE: "show OBJECT",
			APROPOS: lang.Alt{
				lang.EnUS: "print stuff",
			},
			ByName: map[string]cmd.Cmd{
				"cmdline": cmdline.Command{},
				"iminfo":  iminfo.Command{},
			},
		},
		"/init":     &slashinit.Command{},
		"sleep":     sleep.Command{},
		"source":    &source.Command{},
		"start":     &start.Command{},
		"stop":      &stop.Command{},
		"stty":      stty.Command{},
		"subscribe": subscribe.Command{},
		"sync":      sync.Command{},
		"[":         testcmd.Command{},
		"then":      &thencmd.Command{},
		"true":      truecmd.Command{},
		"umount":    umount.Command{},
		"uninstall": &uninstall.Command{},
		"uptimed":   uptimed.Command(make(chan struct{})),
		"wget":      wget.Command{},
	},
}
