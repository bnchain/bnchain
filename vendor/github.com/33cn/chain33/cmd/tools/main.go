// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"

	"github.com/bnchain/bnchain/cmd/tools/commands"
	"github.com/bnchain/bnchain/common/log"
	"github.com/spf13/cobra"
)

//var (
//	mlog = log15.New("module", "tools")
//)

func main() {
	log.SetLogLevel("debug")
	runCommands()
}

func addCommands(rootCmd *cobra.Command) {
	rootCmd.AddCommand(
		commands.ImportCmd(),
		commands.UpdateInitCmd(),
		commands.CreatePluginCmd(),
		commands.GenDappCmd(),
	)
}

func runCommands() {
	rootCmd := &cobra.Command{
		Use:   "tools",
		Short: "bnchain tools",
	}
	addCommands(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		//mlog.Error("Execute command failed.", "error", err)
		os.Exit(1)
	}
}
