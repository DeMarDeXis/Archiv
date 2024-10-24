package cmd

import (
	"Archivator_Go/lib/vlc"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var vlcPackCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Pack file using variable-length code",
	Run:   pack,
}

const packedExtension = "vlc"

var ErrEmptyPath = errors.New("path to file is not specified")

func pack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handleErr(ErrEmptyPath)
	}

	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleErr(err)
	}
	defer r.Close()

	data, err := io.ReadAll(r) //считываем файл в переменную
	if err != nil {
		handleErr(err)
	}

	// data -> Encode(data)

	packed := vlc.Encode(string(data))

	err = os.WriteFile(packedFileName(filePath), packed, 0644)
	if err != nil {
		handleErr(err)
	}
}

func packedFileName(path string) string {
	// /path/to/file/myFile.vlc
	// fileName := filepath.Base(path) //myFile.txt -> myFile
	// ext := filepath.Ext(fileName) //myFile.txt -> .txt
	// baseName := strings.TrimSuffix(fileName, ext) // 'myFile.txt' - '.txt' = 'myFyle'

	// return baseName + "." + packedExtension

	fileName := filepath.Base(path)

	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + packedExtension
}

func init() {
	packCmd.AddCommand(vlcPackCmd)
}
