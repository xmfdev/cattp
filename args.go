package main

import (
	"github.com/xmfdev/targs"
	"os"
)

const (
	flagPortShort           = "-p"
	flagPortLong            = "--port="
	flagAddressShort        = "-a"
	flagAddressLong         = "--address="
	flagFileSystemRootShort = "-r"
	flagFileSystemRootLong  = "--root="
	flagExtraLogsShort      = "-l"
	flagExtraLogsLong       = "--logs"
)

func handleAddress(address string) {
	SERVER_ADDRESS = address
}

func handlePort(port string) {
	SERVER_PORT = port
}

func handleFileSystemRoot(root string) {
	SERVER_FILE_SYSTEM_ROOT = root
}

func handleExtraLogs(_ string) {
	SERVER_ENABLE_EXTRA_LOGS = true
}

func createOptions() []targs.Option {
	return []targs.Option{
		{
			Options:     []string{flagAddressShort, flagAddressLong},
			Handler:     handleAddress,
			HasExtraArg: true,
		},
		{
			Options:     []string{flagPortShort, flagPortLong},
			Handler:     handlePort,
			HasExtraArg: true,
		},
		{
			Options:     []string{flagFileSystemRootShort, flagFileSystemRootLong},
			Handler:     handleFileSystemRoot,
			HasExtraArg: true,
		},
		{
			Options: []string{flagExtraLogsShort, flagExtraLogsLong},
			Handler: handleExtraLogs,
			HasExtraArg: false,
		},
	}
}

func handleArgs() {
	targs.HandleArgs(os.Args[1:], createOptions(), nil)
}
