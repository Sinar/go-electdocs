package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// List of N from 11 to 20
	ns := []string{"11", "12", "13", "14", "15", "16", "17", "18", "19", "20"}
	for _, n := range ns {
		transformN(n)
	}
}

func transformN(n string) {
	inputFile := fmt.Sprintf("DATA/CSV/N.%s %s.csv", n, getName(n))
	outputFile := fmt.Sprintf("DATA/OUTPUT/Sarawak-N.%s.csv", n)
	// Open input
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}
	// Header
	header := records[0]
	// Find candidate columns
	candidateCols := []int{}
	totalCol := -1
	rejectedCol := -1
	unreturnedCol := -1
	jumlahCol := -1
	koddmCol := -1
	namadmCol := -1
	namaPusatCol := -1
	nomborCol := -1
	for i, h := range header {
		if strings.Contains(h, "Bilangan Undian Oleh Pemilih Bagi Setiap Orang Calon Yang Bertanding (B)") && !strings.Contains(h, "Jumlah") {
			candidateCols = append(candidateCols, i)
		} else if strings.Contains(h, "Jumlah Undian Oleh Pemilih") {
			totalCol = i
		} else if strings.Contains(h, "Bilangan Kertas Undi Yang Ditolak") {
			rejectedCol = i
		} else if strings.Contains(h, "Jumlah Kertas Undi Yang Dikeluarkan Kepada Pengundi Tetapi Tidak Dimasukkan Ke Dalam Peti Undi") {
			unreturnedCol = i
		} else if strings.Contains(h, "Jumlah Kertas Undi Yang Patut Berada Di Dalam Peti Undi") {
			jumlahCol = i
		} else if h == "KODDM" {
			koddmCol = i
		} else if h == "NAMADM" {
			namadmCol = i
		} else if h == "Nama Pusat Mengundi" {
			namaPusatCol = i
		} else if h == "Nombor Tempat Mengundi (saluran)" {
			nomborCol = i
		}
	}
	// Party map for this N
	partyMap := getPartyMap(n, header, candidateCols)
	// Output header
	outputHeader := []string{
		"UNIQUE CODE", "STATE", "BALLOT TYPE", "PARLIAMENTARY CONSTITUENCY CODE", "PARLIAMENTARY CONSTITUENCY NAME",
		"STATE CONSTITUENCY CODE", "STATE CONSTITUENCY NAME", "POLLING DISTRICT CODE", "POLLING DISTRICT NAME",
		"POLLING CENTRE", "VOTING CHANNEL NUMBER", "TOTAL BALLOTS ISSUED",
		"BN", "BN CANDIDATE", "BN CANDIDATE SEX", "BN CANDIDATE AGE", "BN VOTE",
		"PH (1)", "PH (1) CANDIDATE", "PH (1) CANDIDATE SEX", "PH (1) CANDIDATE AGE", "PH (1) VOTE",
		"PH (2)", "PH (2) CANDIDATE", "PH (2) CANDIDATE SEX", "PH (2) CANDIDATE AGE", "PH (2) VOTE",
		"PAS", "PAS CANDIDATE", "PAS CANDIDATE SEX", "PAS CANDIDATE AGE", "PAS VOTE",
		"STAR", "STAR CANDIDATE", "STAR CANDIDATE SEX", "STAR CANDIDATE AGE", "STAR VOTE",
		"PBDSB", "PBDSB CANDIDATE", "PBDSB CANDIDATE SEX", "PBDSB CANDIDATE AGE", "PBDSB VOTE",
		"INDEPENDENT 1", "INDEPENDENT 1 CANDIDATE", "INDEPENDENT 1 CANDIDATE SEX", "INDEPENDENT 1 CANDIDATE AGE", "INDEPENDENT 1 VOTE",
		"INDEPENDENT 2", "INDEPENDENT 2 CANDIDATE", "INDEPENDENT 2 CANDIDATE SEX", "INDEPENDENT 2 CANDIDATE AGE", "INDEPENDENT 2 VOTE",
		"TOTAL VALID VOTES", "TOTAL REJECTED VOTES", "TOTAL UNRETURNED BALLOTS", "CHECKER (VALID VOTE)", "CHECKER (TOTAL VOTE ISSUED)",
	}
	// Output records
	var outputRecords [][]string
	outputRecords = append(outputRecords, outputHeader)
	for _, record := range records[1:] {
		outputRow := make([]string, len(outputHeader))
		// UNIQUE CODE
		koddm := record[koddmCol]
		if koddm == "" {
			koddm = "_"
		} else {
			koddm = strings.ReplaceAll(koddm, "/", "_")
		}
		pCode := getPCode(n)
		outputRow[0] = fmt.Sprintf("%s_N.%s_%s_%s", pCode, n, koddm, record[nomborCol])
		// STATE
		outputRow[1] = "SARAWAK"
		// BALLOT TYPE
		namadm := record[namadmCol]
		if namadm == "UNDI POS" {
			outputRow[2] = "POSTAL VOTE"
		} else if namadm == "UNDI AWAL" {
			outputRow[2] = "EARLY VOTE"
		} else {
			outputRow[2] = "ORDINARY VOTE"
		}
		// PARLIAMENTARY CONSTITUENCY CODE
		outputRow[3] = pCode
		// PARLIAMENTARY CONSTITUENCY NAME
		outputRow[4] = getPName(pCode)
		// STATE CONSTITUENCY CODE
		outputRow[5] = fmt.Sprintf("N.%s", n)
		// STATE CONSTITUENCY NAME
		outputRow[6] = getName(n)
		// POLLING DISTRICT CODE
		outputRow[7] = record[koddmCol]
		// POLLING DISTRICT NAME
		outputRow[8] = record[namadmCol]
		// POLLING CENTRE
		outputRow[9] = record[namaPusatCol]
		// VOTING CHANNEL NUMBER
		outputRow[10] = record[nomborCol]
		// TOTAL BALLOTS ISSUED
		outputRow[11] = record[jumlahCol]
		// Votes
		votes := make(map[string]int)
		for i, col := range candidateCols {
			vote, _ := strconv.Atoi(record[col])
			party := partyMap[i]
			votes[party] += vote
		}
		// Fill party columns
		parties := []string{"BN", "PH (1)", "PH (2)", "PAS", "STAR", "PBDSB", "INDEPENDENT 1", "INDEPENDENT 2"}
		idx := 12
		for _, party := range parties {
			outputRow[idx] = ""                           // party name or blank
			outputRow[idx+1] = ""                         // candidate
			outputRow[idx+2] = ""                         // sex
			outputRow[idx+3] = ""                         // age
			outputRow[idx+4] = strconv.Itoa(votes[party]) // vote
			idx += 5
		}
		// TOTAL VALID VOTES
		outputRow[idx] = record[totalCol]
		idx++
		// TOTAL REJECTED VOTES
		outputRow[idx] = record[rejectedCol]
		idx++
		// TOTAL UNRETURNED BALLOTS
		outputRow[idx] = record[unreturnedCol]
		idx++
		// CHECKER
		outputRow[idx] = ""
		idx++
		outputRow[idx] = ""
		outputRecords = append(outputRecords, outputRow)
	}
	// Write output
	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outFile.Close()
	writer := csv.NewWriter(outFile)
	writer.WriteAll(outputRecords)
	writer.Flush()
	if err := writer.Error(); err != nil {
		fmt.Println("Error writing CSV:", err)
	}
}

func getPartyMap(n string, header []string, candidateCols []int) []string {
	partyMap := make([]string, len(candidateCols))
	for i, col := range candidateCols {
		candidateName := strings.TrimPrefix(header[col], "Bilangan Undian Oleh Pemilih Bagi Setiap Orang Calon Yang Bertanding (B) :")
		partyMap[i] = getParty(n, candidateName)
	}
	return partyMap
}

func getParty(n, candidate string) string {
	// Hardcode based on known
	switch n {
	case "11":
		switch candidate {
		case "SIH HUA TONG":
			return "BN"
		case "LINA SOO":
			return "PH (1)"
		case "SEE CHEE HOW":
			return "PH (2)"
		}
	case "12":
		switch candidate {
		case "CHONG CHIENG JEN":
			return "PH (2)"
		case "YAP YAU SIN":
			return "BN"
		}
	case "13":
		switch candidate {
		case "VOON SHIAK NI":
			return "INDEPENDENT 1"
		case "SULAIMAN BIN KADIR":
			return "INDEPENDENT 2"
		case "DATO' SERI HJ. OTHMAN BIN HJ. BOJENG":
			return "BN"
		case "LO KHERE CHIANG":
			return "BN"
		case "ABDUL AZIZ ISA BIN MARINDO":
			return "PH (2)"
		}
	case "14":
		switch candidate {
		case "LIU THIAN LEONG":
			return "INDEPENDENT 1"
		case "DR. SIM KUI HIAN":
			return "BN"
		case "CHIEW WANG SEE":
			return "PH (2)"
		}
	case "15":
		switch candidate {
		case "ABDUL KARIM RAHMAN HAMZAH":
			return "BN"
		case "MAHMUD EPAH":
			return "PH (1)"
		case "ISHAK BUJI":
			return "INDEPENDENT 1"
		case "MOHAMAD MAHDEEN SAHARUDDIN":
			return "INDEPENDENT 2"
		}
	case "16":
		switch candidate {
		case "IDRIS BUANG":
			return "BN"
		case "DAUD EALI":
			return "PH (1)"
		case "YAKUP KHALID":
			return "INDEPENDENT 1"
		case "SIGANDAM SULAMAN":
			return "INDEPENDENT 2"
		case "HIPNI SULAIMAN":
			return "INDEPENDENT 1"
		case "ISMAWI MUHAMMAD":
			return "INDEPENDENT 2"
		}
	case "17":
		switch candidate {
		case "HAMZAH BRAHIM":
			return "BN"
		case "LESLIE TING XIANG ZHI":
			return "PH (2)"
		case "GEORGE YOUNG SI RICORD JUNIOR":
			return "INDEPENDENT 1"
		case "ATET DEGO":
			return "INDEPENDENT 2"
		}
	case "18":
		switch candidate {
		case "MIRO SIMUH":
			return "BN"
		case "MICHAEL SAWING":
			return "PH (1)"
		case "BULN RIBOS":
			return "INDEPENDENT 1"
		case "IANA AKAM":
			return "INDEPENDENT 2"
		case "JECKY MISIENG":
			return "INDEPENDENT 1"
		}
	case "19":
		switch candidate {
		case "JERIP SUSIL":
			return "BN"
		case "CHAN HON HIUNG":
			return "PH (2)"
		case "CHONG SIEW HUNG":
			return "INDEPENDENT 1"
		case "SANJAN DAIK":
			return "INDEPENDENT 2"
		case "JOSHUA ROMAN":
			return "INDEPENDENT 1"
		}
	case "20":
		switch candidate {
		case "ROLAND SAGAH WEE INN":
			return "BN"
		case "CHRISTO MICHAEL":
			return "PH (1)"
		case "DADI TIAP JUUL":
			return "INDEPENDENT 1"
		case "EDISON JAMANG":
			return "INDEPENDENT 2"
		case "BAI DUNG AK":
			return "INDEPENDENT 1"
		}
	}
	return "INDEPENDENT 1" // default
}

func getName(n string) string {
	switch n {
	case "11":
		return "BATU LINTANG"
	case "12":
		return "KOTA SENTOSA"
	case "13":
		return "BATU KITANG"
	case "14":
		return "BATU KAWAH"
	case "15":
		return "ASAJAYA"
	case "16":
		return "MUARA TUANG"
	case "17":
		return "STAKAN"
	case "18":
		return "SEREMBU"
	case "19":
		return "MAMBONG"
	case "20":
		return "TARAT"
	}
	return ""
}

func getPCode(n string) string {
	switch n {
	case "11", "12", "13", "14":
		return "P.192"
	case "15", "16", "17", "18":
		return "P.193"
	case "19", "20":
		return "P.194"
	}
	return ""
}

func getPName(p string) string {
	switch p {
	case "P.192":
		return "MAS GADING"
	case "P.193":
		return "SANTUBONG"
	case "P.194":
		return "PETRA JAYA"
	}
	return ""
}
