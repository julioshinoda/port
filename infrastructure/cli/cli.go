package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/julioshinoda/port/interfaces/json"
	"github.com/spf13/cobra"
)

var parserCmd = &cobra.Command{
	Use:   "parser",
	Short: "Parse json file ",
	Long:  `Parse json file `,
	Run: func(cmd *cobra.Command, args []string) {
		if filename == "" {
			fmt.Printf("invalid file")
			return
		}
		ctx := context.Background()
		json.ParseJson(ctx, filename, os.Getenv("DATABASE_URL"))
	},
}

func init() {
	rootCmd.AddCommand(parserCmd)
}
