package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"

	"github.com/manifoldco/promptui"
	"github.com/meinto/cobra-utils"
	"github.com/meinto/glow/pkg/cli/cmd/util"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(installCmd)
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install glow",
	Run: func(cmd *cobra.Command, args []string) {

		flist, err := fileList(".")
		util.CheckForError(err, "cannot get file list")

		index, _, err := cobraUtils.PromptSelect(
			"Select your downloaded glow file",
			flist,
		)
		util.CheckForError(err, fmt.Sprintf("cannot get path to glow file: %s", err))

		filePath, err := filepath.Abs(flist[index])
		util.CheckForError(err, "cannot get absolute file path")

		index, err = usageOptions()
		util.CheckForError(err, "cannot get usage option")

		var newFileName string
		switch index {
		case 0:
			newFileName = "/usr/local/bin/glow"
		case 1:
			newFileName = "/usr/local/bin/git-glow"
		}

		if _, err := os.Stat(newFileName); !os.IsNotExist(err) {
			replace, err := replaceFile(newFileName)
			if err != nil {
				log.Fatal(err)
			}
			if !replace {
				log.Fatal("file not replaced")
			}
		}

		err = os.Rename(filePath, newFileName)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("successfully moved glow")
	},
}

func pathToGlowFile() (string, error) {
	validate := func(input string) error {
		filePath, err := filepath.Abs(input)
		if err != nil {
			return fmt.Errorf("error while creating absolute path: %s", err.Error())
		}
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return errors.New("file does not exist")
		}
		return nil
	}

	getFileName := promptui.Prompt{
		Label:    "Name of binary: ",
		Validate: validate,
	}

	fileName, err := getFileName.Run()
	if err != nil {
		return "", err
	}

	filePath, err := filepath.Abs(fileName)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

func usageOptions() (int, error) {
	prompt := promptui.Select{
		Label: "How do you want to use glow?",
		Items: []string{"global", "git plugin"},
	}

	index, _, err := prompt.Run()
	if err != nil {
		return -1, err
	}

	return index, nil
}

func replaceFile(filePath string) (bool, error) {
	prompt := promptui.Select{
		Label: "File exists. Do you want to replace it?",
		Items: []string{"yes", "no"},
	}

	index, _, err := prompt.Run()
	if err != nil {
		return false, err
	}

	if index == 0 {
		return true, nil
	}
	return false, nil
}

func fileList(rootPath string) ([]string, error) {
	var files []string

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if !strings.HasPrefix(path, "glow_") {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		return files, errors.Wrap(err, "error creating file list")
	}

	return files, nil
}
