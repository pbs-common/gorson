package cmd

import (
	"github.com/pbs/gorson/internal/gorson/io"
	"github.com/pbs/gorson/internal/gorson/util"
	"github.com/spf13/cobra"
)

var filename string

func put(path string, parameters map[string]string) {
	p := util.NewParameterStorePath(path)
	io.WriteToParameterStore(parameters, *p)
}

func init() {
	cmd := &cobra.Command{
		Use:   "put /a/parameter/store/path --file /path/to/a/file",
		Short: "write parameters to a parameter store path",
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			parameters := io.ReadJSONFile(filename)
			put(path, parameters)
		},
		Args: cobra.ExactArgs(1),
	}
	cmd.Flags().StringVarP(&filename, "file", "f", "", "json file to read key/value pairs from")
	cmd.MarkFlagRequired("file")
	rootCmd.AddCommand(cmd)
}
