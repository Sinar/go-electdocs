# Preparing Data

- Unzip the CSV for state into the data/pru-14-2018/<STATE> folder
- Import each PAR CSV file into Google Sheet; select "Insert New Sheet"; and ensure NO convert date, number
- Format Saluran as Plain text; add 1 for UNDI POS
- Other fields with ballot number formatted as Number without decimals; to double check against PDF

## Double check count against

## Download verified raw data

- Download into testdata folder under the pattern P078.csv; note absence of '.' and always 3 digits for PAR
- Make copy of cmd for STATE
- Run dummy check against 1 PAR

## Download RAW Candidates

- Go to tab RAW_CANDIDATES_PAR; filter by the PAR (e.g. 07800 -> 09100 for PAHANG)
- Copy into a new tab called CANDIDATES_<STATE>
- For field DUN_ID; format as Number; with the pattern '00000'
- Filter the data; sort first by DUN_ID increasing, then BALLOT_ID increasing; with 1st for as Header
- Use formula UNIQUE to determine unique PARTY + COALITION; see RAW_PARTY for lookup
- If needed; add IND + PARTY lookup into the lookup.go file
- Download as <STATE>.csv and put into testdata