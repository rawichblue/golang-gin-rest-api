package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// NotReqArgs Not required areuments
func NotReqArgs(cmd *cobra.Command, args []string) error {
	if len(args) != 0 {
		return fmt.Errorf("Not required areuments")
	}
	return nil
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func must[T any](t T, err error) T {
	panicErr(err)
	return t
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getError(_ any, err error) error {
	return err
}
