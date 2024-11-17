/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"my/scene-switcher/api"
	"my/scene-switcher/scene"

	"github.com/spf13/cobra"
)

// sceneCmd represents the scene command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		schedulerToSynchronizerChannel := make(chan string)
		syncer := scene.DummySynchronizer{
			BaseSynchronizer: scene.BaseSynchronizer{
				Sync: schedulerToSynchronizerChannel}}
		go syncer.Run()

		r := api.SetupRouter()
		r = api.SetupSceneEndpoint(r, schedulerToSynchronizerChannel)
		api.SetupDeviceEndpoint(r)
		mcScheduler := scene.MusicCastSceneScheduler{}
		mcScheduler.Run()
		r.Run(":8088")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sceneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sceneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
