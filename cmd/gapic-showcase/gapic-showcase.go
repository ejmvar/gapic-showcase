// Code generated. DO NOT EDIT.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/spf13/cobra"
)

var Verbose, OutputJSON bool
var ctx = context.Background()
var marshaler = &jsonpb.Marshaler{Indent: "  "}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Print verbose output")
	rootCmd.PersistentFlags().BoolVarP(&OutputJSON, "json", "j", false, "Print JSON output")
}

var rootCmd = &cobra.Command{
	Use:   "gapic-showcase",
	Short: "Root command of gapic-showcase",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}

func printVerboseInput(srv, mthd string, data interface{}) {
	fmt.Println("Service:", srv)
	fmt.Println("Method:", mthd)
	fmt.Print("Input: ")
	printMessage(data)
}

func printMessage(data interface{}) {
	var s string

	if OutputJSON {
		d, _ := json.MarshalIndent(data, "", "  ")
		s = string(d)
	} else if msg, ok := data.(proto.Message); ok {
		s = msg.String()
	} else if page, ok := data.(map[string]interface{}); ok {
		s = fmt.Sprintf("%v", page)
	}

	fmt.Println(s)
}
