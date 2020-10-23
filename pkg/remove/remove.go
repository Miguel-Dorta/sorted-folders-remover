package remove

import (
	"github.com/Miguel-Dorta/sorted-folders-remover/pkg/list"
	"time"
)

// UpToSize preserves the files up to the size specified, then removes them
func UpToSize(path string, size int64, dryRun bool) error {
	totalSize := int64(0)
	currentDate := currentDateInt()
	return list.ListAndRemove(path, &list.Options{
		DryRun:   dryRun,
		SkipSize: false,
	}, func(day list.Day) bool {
		if dateInt(day.Y, day.M, day.D) >= currentDate {
			return false
		}
		totalSize += day.Size
		return totalSize > size
	})
}

// PreservingNDays preserves the files up to n existing days, then removes them
func PreservingNDays(path string, n int, dryRun bool) error {
	totalDays := 0
	currentDate := currentDateInt()
	return list.ListAndRemove(path, &list.Options{
		DryRun:   dryRun,
		SkipSize: true,
	}, func(day list.Day) bool {
		if dateInt(day.Y, day.M, day.D) >= currentDate {
			return false
		}
		totalDays++
		return totalDays > n
	})
}

// currentDateInt gets the dateInt of now
func currentDateInt() int {
	now := time.Now()
	return dateInt(now.Year(), int(now.Month()), now.Day())
}

// dateInt combines the year, month and day of a date in the last bits of an int. The result will be, in reverse order:
// DDDD DMMM MYYY YYYY ...
func dateInt(y, m, d int) int {
	y <<= 9
	y += m << 5
	return y + d
}
