# README - PRU DUN Sarawak ke-12)

## Methodology

Official Site: https://dashboard.spr.gov.my/swk12/#!/home

Raw Data (JSON) for Candidates --> https://dashboard.spr.gov.my/swk12/js/penamaan.js

Raw Data (JSON) for Party --> https://dashboard.spr.gov.my/swk12/js/data.min.js

Raw Data (CSV) for Results --> https://dashboard.spr.gov.my/swk12/js/data.min.js

Gazette for Nomination

Gazette for Result

## Steps

- Add JSON (start with [, end with ]) and paste into --> https://www.convertcsv.com/json-to-csv.htm
- Download csv with headers; very fast + good!!

- Take JSON from https://dashboard.spr.gov.my/swk12/js/penamaan.js (under `dataPenamaan=`) and clean it using https://ryanmarcus.github.io/dirty-json/; copy result into file raw-candidates.json
- Take JSON from https://dashboard.spr.gov.my/swk12/js/data.min.js (under `partiJSON=`) and clean it using https://ryanmarcus.github.io/dirty-json/; copy result into file raw-party.json
- Take the string from `const seatsCsv=` and echo it and pipe into the raw-dun.csv file

## CReate Pivots for major canidates; split independents ..
