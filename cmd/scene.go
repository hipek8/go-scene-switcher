/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"my/scene-switcher/scene"
	"time"

	"github.com/spf13/cobra"
)

// sceneCmd represents the scene command
var sceneCmd = &cobra.Command{
	Use:   "scene",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		println("Running scene")
		println(cmd)
		println(args[0])
		sync := scene.DummySynchronizer{}
		sync.SetScene(args[0], nil, nil)
		time.Sleep(time.Second * 5)
	},
}

func init() {
	rootCmd.AddCommand(sceneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sceneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sceneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
