package main

import (
	"fmt"
	"io"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"
	"github.com/huda7077/gin-and-gorm-boilerplate/models"
)

func main() {
	// More info on https://atlasgo.io/guides/orms/gorm/program
	stmts, err := gormschema.New("postgres").Load(models.AllModels...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, stmts)
}
