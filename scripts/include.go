package main

import (
	"io"
	"os"
	"strings"
)

// Reads all .txt files in the current folder
// and encodes them as strings literals in textfiles.go
func main() {
	resources := []string{"settings.yml", ".default.denvignore"}
	out, _ := os.Create("api/resources.go")
	out.Write([]byte("package api \n\nconst (\n"))
	for _, resource := range resources {
		out.Write([]byte(strings.Replace(resource, ".", "_", -1) + " = `"))
		f, _ := os.Open(resource)
		io.Copy(out, f)
		out.Write([]byte("`\n"))
	}
	out.Write([]byte(")\n"))
}
