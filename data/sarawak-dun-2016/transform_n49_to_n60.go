package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// List of N from 49 to 60
	ns := []string{"49", "50", "51", "52", "53", "54", "55", "56", "57", "58", "59", "60"}
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
	case "21":
		switch candidate {
		case "SIMON SINANG BADA":
			return "BN"
		case "SENIOR WILLIAM RADE":
			return "PH (1)"
		case "CHEYNE KAMBENG":
			return "INDEPENDENT 1"
		case "JONATHAN LANTIK":
			return "INDEPENDENT 2"
		case "ROLAND BANGU":
			return "INDEPENDENT 1"
		}
	case "22":
		switch candidate {
		case "MACLAINE BEN":
			return "BN"
		case "LAERRY JABONG":
			return "PH (2)"
		case "DOMINIC DADO SAGIN":
			return "INDEPENDENT 1"
		case "STEPHEN MORGAN SUGAN":
			return "INDEPENDENT 2"
		}
	case "23":
		switch candidate {
		case "JOHN ILUS":
			return "BN"
		case "BROLIN NICHOLSON BENEDICT ACHUNG":
			return "PH (2)"
		case "ELSIY TINGANG":
			return "INDEPENDENT 1"
		case "EDWARD ANDREW LUAK":
			return "INDEPENDENT 2"
		}
	case "24":
		switch candidate {
		case "AIDEL LARIWOO":
			return "BN"
		case "PIEE BIN LING":
			return "PH (1)"
		case "NUR KHAIRUNISA ABDULLAH":
			return "INDEPENDENT 1"
		case "JOLHI BEE":
			return "INDEPENDENT 2"
		}
	case "25":
		switch candidate {
		case "AWLA IDRIS":
			return "BN"
		case "HAPENI FADIL":
			return "PH (1)"
		case "RAILY ALI":
			return "INDEPENDENT 1"
		case "SAHARUDDIN ABDULLAH":
			return "INDEPENDENT 2"
		}
	case "26":
		switch candidate {
		case "ABANG ABDUL RAHMAN ZOHARI ABANG OPENG":
			return "BN"
		case "KAMAL BUJANG":
			return "PH (1)"
		case "MOHAMAD SOFIAN FARIZ SHARBINI":
			return "INDEPENDENT 1"
		case "TOMSON ANGO":
			return "INDEPENDENT 2"
		}
	case "27":
		switch candidate {
		case "JULAIHI NARawi":
			return "BN"
		case "WEL @ MAXWEL ROJIS":
			return "PH (1)"
		case "WAN CHEE WAN MAHJAR":
			return "INDEPENDENT 1"
		}
	case "28":
		switch candidate {
		case "DAYANG NOORAZAH AWANG SOHOR":
			return "BN"
		case "ABANG ABDUL KASIM ABANG BUJANG":
			return "PH (1)"
		case "WAN ABDILLAH EDRUCE WAN ABDUL RAHMAN":
			return "INDEPENDENT 1"
		case "BAHA IMAN":
			return "INDEPENDENT 2"
		case "ABANG AHMAD ABANG SUNI":
			return "INDEPENDENT 1"
		case "MOHD SEPIAN ABANG DAUD":
			return "INDEPENDENT 2"
		}
	case "29":
		switch candidate {
		case "RAZAILI GAPOR":
			return "BN"
		case "ABANG ZULKIFLI ABANG ENGKEH":
			return "PH (1)"
		case "JACKIE CHIEW":
			return "INDEPENDENT 1"
		case "MOHAMMAD ARIFIRIAZUL PAIIO":
			return "PAS"
		case "SAFIUDIN MATSAH":
			return "INDEPENDENT 2"
		}
	case "30":
		switch candidate {
		case "SNOWDAN LAWAN":
			return "BN"
		case "MASIR KUJAT":
			return "INDEPENDENT 1"
		case "KASIM MANA":
			return "INDEPENDENT 2"
		}
	case "31":
		switch candidate {
		case "MONG DAGANG":
			return "BN"
		case "ENTUSA IMAN":
			return "PBDSB"
		case "NORINA UMOI UTOT":
			return "INDEPENDENT 1"
		case "WINTON LANGGANG":
			return "INDEPENDENT 2"
		}
	case "32":
		switch candidate {
		case "FRANCIS HARDEN HOLLIS":
			return "BN"
		case "LEON JIMAT DONALD":
			return "PH (2)"
		case "WILSON ENTABANG":
			return "INDEPENDENT 1"
		case "PELI ARON":
			return "INDEPENDENT 2"
		}
	case "33":
		switch candidate {
		case "DESMOND SATENG SANJAN":
			return "BN"
		case "JOHNICAL RAYONG NGIPA":
			return "INDEPENDENT 1"
		case "STEL DATU":
			return "INDEPENDENT 2"
		case "GEMONG BATU":
			return "INDEPENDENT 1"
		}
	case "34":
		switch candidate {
		case "MALCOM MUSSEN LAMOH":
			return "BN"
		case "WILLIAM NYALLAU BADAK":
			return "INDEPENDENT 1"
		case "USUP ASUN":
			return "INDEPENDENT 2"
		case "JOHN LINANG MEREEJON":
			return "INDEPENDENT 1"
		}
	case "35":
		switch candidate {
		case "RICKY SITAM":
			return "BN"
		case "PATEK KAMIS":
			return "PH (1)"
		case "MELAINI BOLHASSAN":
			return "INDEPENDENT 1"
		case "SIM MIN LEONG":
			return "INDEPENDENT 2"
		case "KURNAEN BOBEN":
			return "INDEPENDENT 1"
		}
	case "36":
		switch candidate {
		case "GERALD RENTAP JABU":
			return "BN"
		case "ISIK UTAU":
			return "INDEPENDENT 1"
		}
	case "37":
		switch candidate {
		case "DOUGLAS UGGah EMBAS":
			return "BN"
		case "MIKAIL MATHEW ABDULLAH":
			return "PH (1)"
		case "ANDRIA GELAYAN DUNDANG":
			return "INDEPENDENT 1"
		}
	case "38":
		switch candidate {
		case "MOHAMAD DURI":
			return "BN"
		case "JOHN ANTAU LINGGANG":
			return "INDEPENDENT 1"
		case "LINANG CHAPUM":
			return "INDEPENDENT 2"
		}
	case "39":
		switch candidate {
		case "FRIDAY BELIK":
			return "BN"
		case "MUSA DINGGAT":
			return "INDEPENDENT 1"
		case "DANNY KUAN SAN SUI":
			return "INDEPENDENT 2"
		}
	case "40":
		switch candidate {
		case "MOHD CHEE KADIR":
			return "BN"
		case "HUD ANDRI ZULKARNAIN":
			return "PH (1)"
		case "WAN MOHAMAD MADEHI WAN ALI":
			return "INDEPENDENT 1"
		case "MOHAMMAD ASRI KASSIM":
			return "INDEPENDENT 2"
		}
	case "41":
		switch candidate {
		case "LEN TALIF SALLEH":
			return "BN"
		case "ABANG ADITAJAYA ABANG ALWI":
			return "INDEPENDENT 1"
		case "ABDUL MUTALIP ABDULLAH":
			return "INDEPENDENT 2"
		case "WONG CHING LING":
			return "INDEPENDENT 1"
		}
	case "42":
		switch candidate {
		case "ABDULLAH SAIDOL":
			return "BN"
		case "MOHAMAD FADILLAH SABALI":
			return "PH (1)"
		case "ABDUL RAAFIDIN MAJIDI":
			return "INDEPENDENT 1"
		case "JENNY WONG KHING LING":
			return "INDEPENDENT 2"
		case "MOHD ADNAN JULKEPPIL":
			return "INDEPENDENT 1"
		}
	case "43":
		switch candidate {
		case "SAFIEE AHMAD":
			return "BN"
		case "TING ING HUA":
			return "INDEPENDENT 1"
		case "JAMAL IBRAHIM":
			return "INDEPENDENT 2"
		}
	case "44":
		switch candidate {
		case "JUANDA JAYA":
			return "BN"
		case "ZAINAB SUHAILI":
			return "PH (1)"
		case "OSMAN RAFAIE":
			return "INDEPENDENT 1"
		}
	case "45":
		switch candidate {
		case "HUANG TIONG SII":
			return "BN"
		case "PHILIP WONG PACK MING":
			return "PH (2)"
		case "WONG CHIN KING":
			return "INDEPENDENT 1"
		case "WONG KUNG KING":
			return "INDEPENDENT 2"
		}
	case "46":
		switch candidate {
		case "DING KUONG HIING":
			return "BN"
		case "YONG SIEW WEI":
			return "PH (2)"
		case "CHRIS HII RU YEE":
			return "INDEPENDENT 1"
		case "MOH HIONG KING":
			return "INDEPENDENT 2"
		}
	case "47":
		switch candidate {
		case "WILLIAM MAWAN IKOM":
			return "BN"
		case "HEREWARD GRAMONG JOSEPH ALLEN":
			return "INDEPENDENT 1"
		case "JEMELI KERAH":
			return "INDEPENDENT 2"
		case "TEDONG GUNDA":
			return "INDEPENDENT 1"
		case "BRAWI ANGGUONG":
			return "INDEPENDENT 2"
		}
	case "48":
		switch candidate {
		case "ROLLAND DUAT JUBIN":
			return "BN"
		case "ELLY LAWAI NGALAI":
			return "INDEPENDENT 1"
		case "ABDUL HAMID SIONG":
			return "INDEPENDENT 2"
		}
	case "49":
		switch candidate {
		case "ANYI JANA":
			return "BN"
		case "SATU ANCHOM":
			return "PH (1)"
		case "LEO BUNSU":
			return "PBDSB"
		case "JOSEPH JAWA KENDAWANG":
			return "INDEPENDENT 1"
		case "CHARLIE GENAM":
			return "INDEPENDENT 2"
		}
	case "50":
		switch candidate {
		case "ALLAN SIDEN GRAMONG":
			return "BN"
		case "MUHAMMAD FAUZI JOSEPH USIT":
			return "PH (1)"
		case "NGELAYANG UNAU":
			return "INDEPENDENT 1"
		case "MADANG DIMBAB":
			return "INDEPENDENT 2"
		case "MARY RITA MATHIAS":
			return "INDEPENDENT 1"
		}
	case "51":
		switch candidate {
		case "JOSEPH CHIENG JIN EK":
			return "BN"
		case "IRENE MARY CHANG OI LING":
			return "PH (2)"
		case "JESS LAU KI MING":
			return "INDEPENDENT 1"
		case "TING KEE NGUAN":
			return "INDEPENDENT 2"
		case "PRISCILLA LAU":
			return "INDEPENDENT 1"
		case "HII TIONG HUAT":
			return "INDEPENDENT 2"
		}
	case "52":
		switch candidate {
		case "TIONG KING SING":
			return "BN"
		case "PAUL LING":
			return "PH (2)"
		case "JOSEPHINE LAU KIEW PENG":
			return "INDEPENDENT 1"
		case "WONG HUI PING":
			return "INDEPENDENT 2"
		case "JANE LAU SING YEE":
			return "INDEPENDENT 1"
		case "JULIUS ENCHANA":
			return "PBDSB"
		case "FADHIL MOHD ISA":
			return "INDEPENDENT 2"
		case "ENGGA UNCHAT":
			return "INDEPENDENT 1"
		}
	case "53":
		switch candidate {
		case "ROBERT LAU HUI YEW":
			return "BN"
		case "AMY LAU BIK YIN":
			return "PH (2)"
		case "WONG SOON KOH":
			return "INDEPENDENT 1"
		case "MICHELLE LING SHYAN MIH":
			return "INDEPENDENT 2"
		case "RICKY ENTERI":
			return "INDEPENDENT 1"
		}
	case "54":
		switch candidate {
		case "MICHAEL TIANG MING TEE":
			return "BN"
		case "DAVID WONG KEE WOAN":
			return "PH (2)"
		case "JANET LOH WUI PING":
			return "INDEPENDENT 1"
		case "LOW CHONG NGUAN":
			return "INDEPENDENT 2"
		case "JAMIE TIEW YEN HOUNG":
			return "INDEPENDENT 1"
		}
	case "55":
		switch candidate {
		case "ANNUAR RAPaee":
			return "BN"
		case "INTANURAZEAN WAN SAPUAN DAUD":
			return "INDEPENDENT 1"
		case "OLIVIA LIM WEN SIA":
			return "INDEPENDENT 2"
		}
	case "56":
		switch candidate {
		case "FATIMAH ABDULLAH":
			return "BN"
		case "SALLEH MAHALI":
			return "INDEPENDENT 1"
		}
	case "57":
		switch candidate {
		case "ROYSTON VALENTINE":
			return "BN"
		case "MOHD ARWIN ABDULLAH":
			return "PH (1)"
		case "SAIT JUNAIDI":
			return "INDEPENDENT 1"
		case "ZAINUDDIN BUDUG":
			return "INDEPENDENT 2"
		}
	case "58":
		switch candidate {
		case "ABDUL YAKUB ARBI":
			return "BN"
		case "ABDUL JALIL BUJANG":
			return "PH (1)"
		case "YUSUF ABDUL RAHMAN":
			return "INDEPENDENT 1"
		}
	case "59":
		switch candidate {
		case "CHRISTOPHER GIRA SAMBANG":
			return "BN"
		case "JOSEPH ENTULU BELAUN":
			return "INDEPENDENT 1"
		}
	case "60":
		switch candidate {
		case "JOHN SIKIE TAYAI":
			return "BN"
		case "JOSHUA JABENG":
			return "PH (1)"
		case "PETER TUAN":
			return "INDEPENDENT 1"
		case "PHILIP KELANANG":
			return "INDEPENDENT 2"
		case "UGIK SELIPEH":
			return "INDEPENDENT 1"
		case "TIUN KANUN":
			return "INDEPENDENT 2"
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
	case "21":
		return "TEBEDU"
	case "22":
		return "KEDUP"
	case "23":
		return "BUKIT SEMUJA"
	case "24":
		return "SADONG JAYA"
	case "25":
		return "SIMUNJAN"
	case "26":
		return "GEDONG"
	case "27":
		return "SEBUYAU"
	case "28":
		return "LINGGA"
	case "29":
		return "BETING MARO"
	case "30":
		return "BALAI RINGIN"
	case "31":
		return "BUKIT BEGUNAN"
	case "32":
		return "SIMANGGANG"
	case "33":
		return "ENGKILILI"
	case "34":
		return "BATANG AI"
	case "35":
		return "SARIBAS"
	case "36":
		return "LAYAR"
	case "37":
		return "BUKIT SABAN"
	case "38":
		return "KALAKA"
	case "39":
		return "KRIAN"
	case "40":
		return "KABONG"
	case "41":
		return "KUALA RAJANG"
	case "42":
		return "SEMOP"
	case "43":
		return "DARO"
	case "44":
		return "JEMORENG"
	case "45":
		return "REPOK"
	case "46":
		return "MERADONG"
	case "47":
		return "PAKAN"
	case "48":
		return "MELUAN"
	case "49":
		return "NGEMAH"
	case "50":
		return "MACHAN"
	case "51":
		return "BUKIT ASSEK"
	case "52":
		return "DUDONG"
	case "53":
		return "BAWANG ASSAN"
	case "54":
		return "PELAWAN"
	case "55":
		return "NANGKA"
	case "56":
		return "DALAT"
	case "57":
		return "TELLIAN"
	case "58":
		return "BALINGIAN"
	case "59":
		return "TAMIN"
	case "60":
		return "KAKUS"
	}
	return ""
}

func getPCode(n string) string {
	switch n {
	case "11", "12", "13", "14":
		return "P.192"
	case "15", "16", "17", "18":
		return "P.193"
	case "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32":
		return "P.194"
	case "33", "34":
		return "P.199"
	case "35", "36", "37":
		return "P.203"
	case "38", "39", "40":
		return "P.205"
	case "41", "42":
		return "P.206"
	case "43", "44":
		return "P.207"
	case "45", "46":
		return "P.208"
	case "47", "48":
		return "P.209"
	case "49", "50":
		return "P.210"
	case "51", "52":
		return "P.211"
	case "53", "54", "55":
		return "P.212"
	case "56", "57":
		return "P.213"
	case "58", "59", "60":
		return "P.214"
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
	case "P.199":
		return "SERIAN"
	case "P.203":
		return "LUBOK ANTU"
	case "P.205":
		return "SARATOK"
	case "P.206":
		return "TANJONG MANIS"
	case "P.207":
		return "IGAN"
	case "P.208":
		return "SARIKEI"
	case "P.209":
		return "JULAU"
	case "P.210":
		return "KANOWIT"
	case "P.211":
		return "LANANG"
	case "P.212":
		return "SIBU"
	case "P.213":
		return "MUKAH"
	case "P.214":
		return "SELANGAU"
	}
	return ""
}
