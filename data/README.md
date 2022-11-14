# Data

Put raw data here ...

## SPR Data

Election Commission (SPR) data 

Latest candidates at the top is --> https://dashboard.spr.gov.my/js/penamaan.js

otherwise it will be a form of the event e.g.  https://dashboard.spr.gov.my/sbh16/js/penamaan.js

DUN + Party data can be found here for the latest general --> https://dashboard.spr.gov.my/js/data.min.js

whereas for the specific one will be like e.g. https://dashboard.spr.gov.my/sbh16/js/data.min.js?234077=

NOTE: SPR data for a General Election seems to be different

Candidates --> https://dashboard.spr.gov.my/pru14/js/penamaan.js
Party + DUN + PAR --> https://dashboard.spr.gov.my/pru14/js/data.js

### HOWTO Obtain SPR Data

Use https://www.convertcsv.com/json-to-csv.htm; the most robust against

Converted data should be stored under the appropriate folder with the following:
- Candidates + Results - candidates-results.csv
- DUNs - dun.csv
- PARs - par.csv
- Party - party.csv


## Gazette Data

### HOWTO Obtain

### HOWTO Clean JSON

- https://jsonformatter.curiousconcept.com/# <-- auto-fix format
- https://www.convertcsv.com/json-to-csv.htm
 
### HOWTO Clean

- Use the tooling under ./tools .. Tabula
- If all else fails; e.g.
  - Open it in Google Doc to get the raw text version in drive.google.com; download as text
  - Leave a Google doc shared link with Viewer-only permission; as it has spatial structure easier to copy
