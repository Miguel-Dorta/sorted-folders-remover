package list

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

type Day struct {
	Y, M, D int
	Size    int64
}

type Options struct {
	DryRun   bool
	SkipSize bool
}

type dirRemover struct {
	d          Day
	opts       *Options
	removeFunc func(Day) bool
}

// ListAndRemove lists the path provided, that MUST contain a substructure like [path/]YYYY/MM/DD, and applies
// removeFunc to each day directory to determine if remove it. It's recommended that you set opts.SkipSize true if
// you're not going to use list.Day.Size
func ListAndRemove(path string, opts *Options, removeFunc func(Day) bool) error {
	dr := &dirRemover{
		removeFunc: removeFunc,
		opts:       opts,
	}
	yearDirs, err := ioutil.ReadDir(path)
	if err != nil {
		return fmt.Errorf("error listing path: %w", err)
	}

	if _, err := list(path, yearDirs, dr.listYear); err != nil {
		return err
	}
	return nil
}

// list iterates the files provided in reverse alphabetic order and, if they are a directory and its name can be
// converted to an int, it executes listFunc providing the file path and int. It returns if all the valid files where
// removed.
func list(path string, dirs []os.FileInfo, listFunc func(string, int) (bool, error)) (bool, error) {
	sortFilesDescending(dirs)
	allRemoved := true
	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}
		dirName := dir.Name()
		dirInt, err := strconv.Atoi(dirName)
		if err != nil {
			continue
		}

		b, err := listFunc(filepath.Join(path, dirName), dirInt)
		if err != nil {
			return false, err
		}
		if !b {
			allRemoved = false
		}
	}
	return allRemoved, nil
}
