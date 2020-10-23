package list

import (
	"fmt"
	log "github.com/Miguel-Dorta/logolang"
	"os"
	"path/filepath"
	"sort"
)

// removeAll is equal to os.RemoveAll except when the Option.DryRun is active
func (dr *dirRemover) removeAll(path string) error {
	if dr.opts.DryRun {
		log.Infof("Dry run: skip removeAll for path %s", path)
		return nil
	}
	return os.RemoveAll(path)
}

// getDirSize iterates recursively a directory and returns the sum of the size of the files it contains
func getDirSize(path string) (int64, error) {
	total := int64(0)
	if err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			total += info.Size()
		}
		return nil
	}); err != nil {
		return -1, fmt.Errorf("error getting dir \"%s\" size: %w", path, err)
	}
	return total, nil
}

// sortFilesDescending sorts an slice of os.FileInfo in descending order of their name
func sortFilesDescending(f []os.FileInfo) {
	sort.Slice(f, func(i, j int) bool {
		return f[i].Name() > f[j].Name()
	})
}
