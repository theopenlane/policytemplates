package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/TylerBrock/colorjson"
	"github.com/spf13/cobra"
)

// jsonPrint prints a JSON formatted string with color
func jsonPrint(s []byte) error {
	var obj map[string]interface{}

	err := json.Unmarshal(s, &obj)
	cobra.CheckErr(err)

	f := colorjson.NewFormatter()
	f.Indent = 2

	o, err := f.Marshal(obj)
	cobra.CheckErr(err)

	fmt.Println(string(o))

	return nil
}

// jsonOutput prints the output in a JSON format
func jsonOutput(out any) error {
	s, err := json.Marshal(out)
	cobra.CheckErr(err)

	return jsonPrint(s)
}

// saveToFile saves the output to a file
func saveToFile(out any, path string) error {
	s, err := json.MarshalIndent(out, "", "    ")
	cobra.CheckErr(err)

	err = os.WriteFile(path, s, 0600) //nolint:mnd

	return err
}
