package main

import (
	"fmt"
	"os"
	"time"
)

func SoalSatu(input []string) string {
	var result string
	var word string

	strCache := make(map[string]int)

	for i := 0; i < len(input); i++ {
		if word == "" {
			if _, ok := strCache[input[i]]; !ok {
				strCache[input[i]] = i
			} else {
				word = input[i]
				result = fmt.Sprintf("%d %d", strCache[input[i]]+1, i+1)
			}
		} else if word == input[i] {
			result = fmt.Sprintf("%s %d", result, i+1)
		}
	}

	if word == "" {
		return "false"
	}

	return result
}

func SoalDua(totalBelanja int64, uangPembeli int64) {
	//sorted uang pecahan
	pecahan := []int64{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}
	seriKembalian := make([]int, len(pecahan))

	// kurang bayar
	if uangPembeli < totalBelanja {
		fmt.Println("False, kurang bayar")
		return
	}

	kembalian := uangPembeli - totalBelanja

	// loop only when kembalian lebih besar dari pecahan terkecil
	for kembalian > pecahan[len(pecahan)-1] {
		for i, v := range pecahan {
			if v < kembalian {
				seriKembalian[i]++
				kembalian = kembalian - v
				break
			}
		}
	}

	// print kembalian
	for i, v := range seriKembalian {
		if v > 0 {
			fmt.Printf("%d lembar %d\n", v, pecahan[i])
		}
	}
}

func SoalTiga(input string) bool {
	// to track opening brackets
	brackets := map[rune]int{
		'<': 0, '{': 0, '[': 0,
	}

	closing := map[rune]rune{
		'>': '<', '}': '{', ']': '[',
	}

	for _, v := range input {
		// if opening bracket
		if _, ok := brackets[v]; ok {
			brackets[v]++
			continue
		}

		// if no opening bracket available for this bracket
		if brackets[closing[v]] <= 0 {
			return false
		}

		brackets[closing[v]]--
	}

	// check any left over opening brackets
	for _, v := range brackets {
		if v > 0 {
			return false
		}
	}

	return true
}

func SoalEmpat(cutiBersama int, tanggalJoin string, tanggalCuti string, durasiCuti int) (bool, string) {
	cutiKantor := 14
	var message string

	jumlahCuti := cutiKantor - cutiBersama

	join, err := time.Parse("2006-01-02", tanggalJoin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cuti, err := time.Parse("2006-01-02", tanggalCuti)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	endOfTheYear := time.Date(cuti.Year(), 12, 31, 0, 0, 0, 0, cuti.Location())
	validCuti := join.Add(180 * 24 * time.Hour)
	totalDays := endOfTheYear.Sub(validCuti)

	// check 180 days validation
	if validCuti.After(cuti) {
		return false, "Alasan: Karena belum 180 hari sejak tanggal join karyawan"
	}

	var days = totalDays.Hours() / 24

	// employee already passed 1 year
	if days > 365 {
		days = 365
	}

	cutiCanBeTaken := days / 365 * float64(jumlahCuti)

	if durasiCuti > int(cutiCanBeTaken) {
		return false, fmt.Sprintf("Alasan: Karena hanya boleh mengambil %.0f hari cuti", cutiCanBeTaken)
	}

	return true, message
}
