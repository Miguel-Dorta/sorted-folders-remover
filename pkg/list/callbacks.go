package list

import (
	"fmt"
	"io/ioutil"
)

// listYear is the callback for the list function when listing years
func (dr *dirRemover) listYear(yearPath string, yearInt int) (bool, error) {
	dr.d.Y = yearInt
	monthDirs, err := ioutil.ReadDir(yearPath)
	if err != nil {
		return false, fmt.Errorf("error listing year directory \"%s\": %w", yearPath, err)
	}

	allMonthsRemoved, err := list(yearPath, monthDirs, dr.listMonth)
	if err != nil {
		return false, err
	}
	if !allMonthsRemoved {
		return false, nil
	}
	if err := dr.removeAll(yearPath); err != nil {
		return false, fmt.Errorf("error removing year directory \"%s\": %w", yearPath, err)
	}
	return true, nil
}

// listMonth is the callback for the list function when listing months
func (dr *dirRemover) listMonth(monthPath string, monthInt int) (bool, error) {
	dr.d.M = monthInt
	dayDirs, err := ioutil.ReadDir(monthPath)
	if err != nil {
		return false, fmt.Errorf("error listing month directory \"%s\": %w", monthPath, err)
	}

	allDaysRemoved, err := list(monthPath, dayDirs, dr.listDay)
	if err != nil {
		return false, err
	}
	if !allDaysRemoved {
		return false, nil
	}
	if err := dr.removeAll(monthPath); err != nil {
		return false, fmt.Errorf("error removing month directory \"%s\": %w", monthPath, err)
	}
	return true, nil
}

// listDay is the callback for the list function when listing days
func (dr *dirRemover) listDay(dayPath string, dayInt int) (bool, error) {
	dr.d.D = dayInt
	if !dr.opts.SkipSize {
		size, err := getDirSize(dayPath)
		if err != nil {
			return false, err
		}
		dr.d.Size = size
	}
	if !dr.removeFunc(dr.d) {
		return false, nil
	}
	if err := dr.removeAll(dayPath); err != nil {
		return false, fmt.Errorf("error removing day directory \"%s\": %w", dayPath, err)
	}
	return true, nil
}
