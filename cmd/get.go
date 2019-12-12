package cmd

import (
	"fmt"
	"log"
	"regexp"

	"github.com/pbs/gorson/internal/gorson/io"
	"github.com/pbs/gorson/internal/gorson/json"
	"github.com/pbs/gorson/internal/gorson/util"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var format string

func get(path string) {
	p := util.NewParameterStorePath(path)
	pms := io.ReadFromParameterStore(*p)
	if format == "yaml" || format == "yml" {
		serialized, err := yaml.Marshal(pms)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(serialized))

	} else if format == "env" {
		serialized, err := yaml.Marshal(pms)
		if err != nil {
			log.Fatal(err)
		}
		var re = regexp.MustCompile(`(\w*):\s`)
		fmt.Println(re.ReplaceAllString(string(serialized), "${1}="))

	} else if format == "json" {
		marshalled := json.Marshal(pms)
		fmt.Println(marshalled)
	} else {
		log.Fatal("No proper format requested. (yaml, env, json allowed)")
	}
}

func init() {
	cmd := &cobra.Command{
		Use:   "get /a/parameter/store/path",
		Short: "Get parameters from a parameter store path",
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			get(path)
		},
		Args: cobra.ExactArgs(1),
	}
	cmd.Flags().StringVarP(&format, "format", "f", "json", "the format of gorson get output.")
	rootCmd.AddCommand(cmd)
}
