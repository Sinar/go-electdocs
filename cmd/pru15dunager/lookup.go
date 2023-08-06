package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"os"
	"strconv"
	"strings"
)

type Candidate struct {
	KodDUN          string `json:"Kod DUN"`
	NamaDUN         string `json:"Nama DUN"`
	BilanganPemilih int    `json:"Bilangan Pemilih"`
	NoCalon         int    `json:"No Calon"`
	Nama            string `json:"Nama"`
	Jantina         string `json:"Jantina"`
	Umur            int    `json:"Umur"`
	Parti           string `json:"Parti"`
}

// LoadRawData ..

// Lookup Functions

func MapPARDUNs(state string) map[string]string {
	pd := make(map[string]string, 0)

	// Iterate through ..
	// Action if NOT exist; otherwise skip?
	return pd
}

// LookupCandidatesByDUN extracts Candidate based on PAR, DUN NAME
func LookupCandidatesByDUN(state, par, dun string) {
	// Example: https://pru.sinarharian.com.my/prn15/negeri-sembilan/tampin/repah

}

// LoadTSVCandidates load from the raw data from Bernama
func LoadTSVCandidates(state string) []Candidate {
	candidates := make([]Candidate, 0)
	fmt.Println("Parsing data for State:", state)
	filePath := fmt.Sprintf("testdata/%s.csv", state) // Update the file path with the actual file name

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return candidates
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comma = '\t'

	records, err := r.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return candidates
	}

	for i, row := range records {
		if i == 0 {
			// Skip the header row
			continue
		}

		candidate := Candidate{
			KodDUN:          strings.TrimSpace(row[0]),
			NamaDUN:         strings.TrimSpace(row[1]),
			BilanganPemilih: parseInt(strings.ReplaceAll(row[2], ",", "")),
			NoCalon:         parseInt(row[3]),
			Nama:            strings.TrimSpace(strings.ToUpper(row[4])),
			Jantina:         strings.TrimSpace(row[5]),
			Umur:            parseInt(row[6]),
			Parti:           strings.TrimSpace(row[7]),
		}

		candidates = append(candidates, candidate)
	}

	// DEBUG
	jsonData, err := json.MarshalIndent(candidates, "", "    ")
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return candidates
	}

	fmt.Println(string(jsonData))

	spew.Dump(candidates)
	return candidates
}

func parseInt(s string) int {

	var num int
	_, err := fmt.Sscanf(s, "%d", &num)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	}
	return num
}

// LookUpN9Candidate has ugly hard code
func LookUpN9Candidate() map[string][]string {
	mapN9 := map[string][]string{
		"N01": {"https://pru.sinarharian.com.my/calon/3723/loke-siew-fook", "https://pru.sinarharian.com.my/calon/8694/rosmadi-arif"},
		"N02": {"https://pru.sinarharian.com.my/calon/1165/jalaluddin-alias", "https://pru.sinarharian.com.my/calon/8703/amirudin-hassan"},
	}

	mapN9["N03"] = []string{"https://pru.sinarharian.com.my/calon/1166/mohd-razi-mohd-ali", "https://pru.sinarharian.com.my/calon/8707/mohd-nordin-hasim"}
	mapN9["N04"] = []string{"https://pru.sinarharian.com.my/calon/6384/bakri-sawir", "https://pru.sinarharian.com.my/calon/8708/danni-rais", "https://pru.sinarharian.com.my/calon/8791/angah-saiful"}
	mapN9["N05"] = []string{"https://pru.sinarharian.com.my/calon/8558/muhamad-zamri-omar", "https://pru.sinarharian.com.my/calon/5917/mohd-fairuz-mohd-isa"}
	mapN9["N06"] = []string{"https://pru.sinarharian.com.my/calon/5909/mustapha-nagoor", "https://pru.sinarharian.com.my/calon/8695/noor-azuan-parmin"}
	mapN9["N07"] = []string{"https://pru.sinarharian.com.my/calon/8561/mohd-zaidy-abdul-kadir", "https://pru.sinarharian.com.my/calon/5924/surash-sreenivasan"}
	mapN9["N08"] = []string{"https://pru.sinarharian.com.my/calon/3658/teo-kok-seong", "https://pru.sinarharian.com.my/calon/5914/kumar-s-paramasivam"}
	mapN9["N09"] = []string{"https://pru.sinarharian.com.my/calon/8563/mohd-asna-amin", "https://pru.sinarharian.com.my/calon/7806/fadli-che-me", "https://pru.sinarharian.com.my/calon/8792/zul-azki-mat-sulop"}
	mapN9["N10"] = []string{"https://pru.sinarharian.com.my/calon/3867/arul-kumar-al-jambunathan", "https://pru.sinarharian.com.my/calon/1582/gan-chee-biow", "https://pru.sinarharian.com.my/calon/8794/omar-bin-mohd-isa", "https://pru.sinarharian.com.my/calon/8795/yessu"}
	mapN9["N11"] = []string{"https://pru.sinarharian.com.my/calon/3866/chew-seh-yong", "https://pru.sinarharian.com.my/calon/8704/ng-soon-lean"}
	mapN9["N12"] = []string{"https://pru.sinarharian.com.my/calon/3869/ng-chin-tsai", "https://pru.sinarharian.com.my/calon/8502/qusyairi-ahmad", "https://pru.sinarharian.com.my/calon/8702/chua-eng-pu"}
	mapN9["N13"] = []string{"https://pru.sinarharian.com.my/calon/3440/aminuddin-harun", "https://pru.sinarharian.com.my/calon/8701/ahmad-raihan-muhamad-hilal", "https://pru.sinarharian.com.my/calon/8801/mohamed-hafiz-tok-merah", "https://pru.sinarharian.com.my/calon/4123/bujang-abu"}
	mapN9["N14"] = []string{"https://pru.sinarharian.com.my/calon/3442/tengku-zamrah-tengku-sulaiman", "https://pru.sinarharian.com.my/calon/5791/mohamad-rafie-ab-malek", "https://pru.sinarharian.com.my/calon/8700/muhammad-ghazali-zainal-abidin"}
	mapN9["N15"] = []string{"https://pru.sinarharian.com.my/calon/540/bibi-sharliza-mohd-khalid", "https://pru.sinarharian.com.my/calon/5802/eddin-syazlee-shith"}
	mapN9["N16"] = []string{"https://pru.sinarharian.com.my/calon/8566/muhammad-sufian-maradzi", "https://pru.sinarharian.com.my/calon/5812/jamali-salam"}
	mapN9["N17"] = []string{"https://pru.sinarharian.com.my/calon/4899/ismail-lasim", "https://pru.sinarharian.com.my/calon/8684/amrina-mohd-khalid"}
	mapN9["N18"] = []string{"https://pru.sinarharian.com.my/calon/8586/nur-zunita-begum", "https://pru.sinarharian.com.my/calon/5800/rafiei-mustapha"}
	mapN9["N19"] = []string{"https://pru.sinarharian.com.my/calon/5823/saiful-yazan-sulaiman", "https://pru.sinarharian.com.my/calon/5824/kamaruddin-md-tahir"}
	mapN9["N20"] = []string{"https://pru.sinarharian.com.my/calon/5839/ismail-ahmad-labu", "https://pru.sinarharian.com.my/calon/8699/mohamad-hanifah-abu-baker"}
	mapN9["N21"] = []string{"https://pru.sinarharian.com.my/calon/5848/tan-lee-koon", "https://pru.sinarharian.com.my/calon/8698/p-subramaniam", "https://pru.sinarharian.com.my/calon/8803/ahmad-zamali-mohamad"}
	mapN9["N22"] = []string{"https://pru.sinarharian.com.my/calon/8697/lee-boon-shian", "https://pru.sinarharian.com.my/calon/8594/siau-meow-kong"}
	mapN9["N23"] = []string{"https://pru.sinarharian.com.my/calon/3871/yap-yew-weng", "https://pru.sinarharian.com.my/calon/8693/n-satesh-kutmar", "https://pru.sinarharian.com.my/calon/8805/kumaravel-al-ramiah"}
	mapN9["N24"] = []string{"https://pru.sinarharian.com.my/calon/3724/p-gunasegaran", "https://pru.sinarharian.com.my/calon/8696/lee-ban-fatt"}
	mapN9["N25"] = []string{"https://pru.sinarharian.com.my/calon/7742/norwani-ahmat", "https://pru.sinarharian.com.my/calon/2809/kamarol-ridzuan-mohd-zain", "https://pru.sinarharian.com.my/calon/8807/syakir-fitri"}
	mapN9["N26"] = []string{"https://pru.sinarharian.com.my/calon/1179/zaifulbahri-idris", "https://pru.sinarharian.com.my/calon/8691/bakly-baba"}
	mapN9["N27"] = []string{"https://pru.sinarharian.com.my/calon/1180/mohamad-hasan", "https://pru.sinarharian.com.my/calon/8690/rozmal-malakan"}
	mapN9["N28"] = []string{"https://pru.sinarharian.com.my/calon/8568/suhaimi-aini", "https://pru.sinarharian.com.my/calon/8706/ahmad-shukri-abdul-shukor"}
	mapN9["N29"] = []string{"https://pru.sinarharian.com.my/calon/8584/yew-boon-lye", "https://pru.sinarharian.com.my/calon/8705/tang-jay-son"}
	mapN9["N30"] = []string{"https://pru.sinarharian.com.my/calon/6412/choo-ken-hwa", "https://pru.sinarharian.com.my/calon/8689/ragu-subramaniam"}
	mapN9["N31"] = []string{"https://pru.sinarharian.com.my/calon/8571/mohd-najib-mohd-isa", "https://pru.sinarharian.com.my/calon/8688/abd-fatah-zakaria"}
	mapN9["N32"] = []string{"https://pru.sinarharian.com.my/calon/544/mohd-faizal-ramli", "https://pru.sinarharian.com.my/calon/8692/zamri-md-said"}
	mapN9["N33"] = []string{"https://pru.sinarharian.com.my/calon/8587/dr-raja-sekaran", "https://pru.sinarharian.com.my/calon/8769/zabidi-ariffin"}
	mapN9["N34"] = []string{"https://pru.sinarharian.com.my/calon/1184/abd-razak-ab-said", "https://pru.sinarharian.com.my/calon/8686/ridzuan-ahmad"}
	mapN9["N35"] = []string{"https://pru.sinarharian.com.my/calon/8572/suhaimizan-bizar", "https://pru.sinarharian.com.my/calon/8687/tengku-abdullah-tengku-rakman"}
	mapN9["N36"] = []string{"https://pru.sinarharian.com.my/calon/5995/s-veerapan", "https://pru.sinarharian.com.my/calon/8685/yong-li-yi"}

	// DEBUG
	//spew.Dump(mapN9)
	return mapN9
}

// EnrichCandidateData tries to find party ..
func EnrichCandidateData(state string, candidates []Candidate) {
	// State will fix the PAR
	//var mapCandidate map[string][]string
	var sourceURL string

	switch state {
	case "N9":
		//mapCandidate = LookUpN9Candidate()
		sourceURL = "https://prn.bernama.com/penamaan/index-en.php?n=05"
	case "PENANG":

	case "SELANGOR":

	default:
		panic("INVALID STATE!!")

	}
	// For each DUN
	//spew.Dump(mapCandidate)
	// Start with the headers first ..
	candidateData := make([][]string, 0)
	// Header ..
	row := make([]string, 6)
	row[0] = "MATCH_URL"
	row[1] = "NAME"
	row[2] = "AGE"
	row[3] = "SINARHARIAN_AGE"
	row[4] = "BERNAMA_AGE"
	row[5] = "GENDER"
	candidateData = append(candidateData, row)

	for i, candidate := range candidates {
		dunID := candidate.KodDUN
		fmt.Println("LINE:", i, "KOD", dunID)
		// DEBUG
		//spew.Dump(mapCandidate[dunID])
		//if i >= 1 {
		//	break
		//}
		gender := ""
		switch candidate.Jantina {
		case "LELAKI":
			gender = "MALE"
		case "PEREMPUAN":
			gender = "FEMALE"
		default:
			gender = "INVALID"
		}
		row := make([]string, 6)
		row[0] = sourceURL
		row[1] = candidate.Nama
		row[4] = strconv.Itoa(candidate.Umur)
		row[5] = gender
		candidateData = append(candidateData, row)
	}

	// DEBUG ..
	//spew.Dump(candidateData)
	// Output with filename suffix -age-party ..
	outputCSV(fmt.Sprintf("testdata/%s-candidates-age-party.csv", state), candidateData)

}
