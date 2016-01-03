package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/ghodss/yaml"
)

// FormatOutput autmatically formats json based output based on user choice.
// when selected formatter is "pretty", call prettyFormatter callback.
func FormatOutput(data []byte, prettyFormatter func([]byte)) {
	switch Format {
	case "pretty":
		prettyFormatter(data)
	case "json":
		jsonFormatter(data)
	case "yaml":
		yamlFormatter(data)
	default:
		fmt.Fprintf(os.Stderr, "Invalid formater %s. Use one of 'pretty', 'json', 'yaml'\n", Format)
		return
	}
}

// FormatOutputDef autmatically formats json based output based on user choice.
// uses yamlFormatter as pretty formatter.
func FormatOutputDef(data []byte) {
	FormatOutput(data, yamlFormatter)
}

func jsonFormatter(data []byte) {
	var out bytes.Buffer
	json.Indent(&out, data, "", "  ")
	fmt.Println(out.String())
}

func yamlFormatter(data []byte) {
	out, err := yaml.JSONToYAML(data)
	Check(err)
	fmt.Print(string(out))
}
