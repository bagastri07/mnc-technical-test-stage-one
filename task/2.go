package task

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func formatCurrency(amount int) string {
	result := ""
	for i, c := range strconv.Itoa(amount) {
		if i > 0 && (len(strconv.Itoa(amount))-i)%3 == 0 {
			result += "."
		}
		result += string(c)
	}
	return result
}

func calculateChange(totalBelanja, uangDibayar int) string {
	if uangDibayar < totalBelanja {
		return "False, kurang bayar"
	}

	// Rupiah denominations
	denominations := []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}

	// Calculate the change and round down to the nearest 100
	kembalian := uangDibayar - totalBelanja
	kembalianRounded := int(math.Floor(float64(kembalian)/100) * 100)

	// Calculate the denominations required
	result := fmt.Sprintf("Kembalian yang harus diberikan kasir: %s, dibulatkan menjadi %s\nPecahan uang:\n", formatCurrency(kembalian), formatCurrency(kembalianRounded))
	changeDetails := make(map[int]int)

	for _, denom := range denominations {
		if kembalianRounded >= denom {
			quantity := kembalianRounded / denom
			changeDetails[denom] = quantity
			kembalianRounded -= quantity * denom
		}
	}

	// Sort denominations in descending order for consistent output
	keys := make([]int, 0, len(changeDetails))
	for denom := range changeDetails {
		keys = append(keys, denom)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	// Format the output with details of each denomination
	for _, denom := range keys {
		quantity := changeDetails[denom]
		unit := "lembar"
		if denom < 1000 {
			unit = "koin"
		}
		result += fmt.Sprintf("%d %s %s\n", quantity, unit, formatCurrency(denom))
	}

	return strings.TrimSpace(result)
}
