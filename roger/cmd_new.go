package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func new(c *cli.Context) error {

	if c.NArg() == 0 {
		err := fmt.Errorf(`Need args`)
		return err
	}
	projectDir := c.Args().Get(0)
	if _, err := os.Stat(projectDir); os.IsNotExist(err) {

		// project root dir
		os.Mkdir(projectDir, os.ModePerm)

		// // config dir
		// configDir := filepath.Join(projectDir, `config`)
		// os.Mkdir(configDir, os.ModePerm)

		// // copy yml
		// src := filepath.Join(configDir, `environments.yml.sample`)
		// dest := filepath.Join(configDir, `environments.yml`)
		// _ = os.Link(src, dest)

		// //
		// src := filepath.Join(projectDir, `environments.yml.sample`)
		// dest := filepath.Join(projectDir, `environments.yml`)
		// _ = os.Link(src, dest)

		return nil
	}
	return nil
}
