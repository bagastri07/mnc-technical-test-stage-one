package task

import (
	"fmt"
	"math"
	"time"
)

func canTakePersonalLeave(joinDateStr string, jointLeaveDays int, leaveStartDateStr string, leaveDuration int) (bool, string) {
	// Parse the dates
	joinDate, err := time.Parse("2006-01-02", joinDateStr)
	if err != nil {
		return false, "Tanggal bergabung tidak valid."
	}
	leaveStartDate, err := time.Parse("2006-01-02", leaveStartDateStr)
	if err != nil {
		return false, "Tanggal rencana cuti tidak valid."
	}

	// Calculate eligibility start date (180 days after joining)
	eligibilityStartDate := joinDate.AddDate(0, 0, 180)

	// Check if the employee is eligible for personal leave
	if leaveStartDate.Before(eligibilityStartDate) {
		return false, "Karena belum 180 hari sejak tanggal join karyawan"
	}

	// Calculate the number of available days until the end of the year
	endOfYear := time.Date(joinDate.Year(), 12, 31, 0, 0, 0, 0, joinDate.Location())
	availableDays := int(endOfYear.Sub(eligibilityStartDate).Hours()/24) + 1 // Include the last day

	// Calculate the maximum personal leave days that can be taken
	maxLeaveDays := int(math.Floor(float64(availableDays) / 365 * float64(jointLeaveDays)))

	// Check leave duration limits
	if leaveDuration > maxLeaveDays {
		return false, fmt.Sprintf("Karena hanya boleh mengambil %d hari cuti, tetapi mencoba mengambil %d hari.", maxLeaveDays, leaveDuration)
	}
	if leaveDuration > 3 {
		return false, "Durasi cuti pribadi maksimal 3 hari berturut-turut."
	}

	return true, "Cuti pribadi dapat diambil."
}
