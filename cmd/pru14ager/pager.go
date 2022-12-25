package main

import (
	"fmt"
	"github.com/bitfield/script"
	"github.com/davecgh/go-spew/spew"
	"net/url"
	"os"
	"regexp"
	"strings"
)

// Map[KeyID] --> Candidate
// KeyID: PAR_ID
// mapCandidate takes from result and assembles
var mapCandidate map[string][]candidate

type candidate struct {
	name        string
	age         string // After transformation
	matchURL    string
	matchedName string
	matchRawAge string
}

func checkPAR() {
	// From top level PAR
	p := script.NewPipe()
	defer p.Close()
	data := url.Values{}
	// __EVENTTARGET=ctl00$ContentPlaceHolder1$uscKeputusanParlimen$rdoPru$3
	//__EVENTARGUMENT
	//__LASTFOCUS
	//__VSTATE=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
	//__VIEWSTATE
	//__SCROLLPOSITIONX=0
	//__SCROLLPOSITIONY=601
	//__EVENTVALIDATION=/wEdABKvVXD1oYELeveMr0vHCmYPfjiX7FFf0MJkoirIsIEK3E9Tycr57kum+sDlA7xMrwRXpgzk1qeIbbs7UICXjXiK1nV2ohCvKna6jgbyA5oyp8B20IVHcLGGpeHVvNzxaZX94t14uLkyce/U+ST7y7+R+lUWuAch62jh/2FBC2Fnd8T6HbMFOjyOD/wCSwNpSxnUcrI8f+67haW8kYJzzOPnuDqKweB2slFuiUnzXJep9VFeKBQGCETivikDUr9a+pnti7fcgdy5I/l0+1ILAMWVMYdPUi2QWAkG1OdXxHiJfmapT5EEJ94JFq1Ypzsoo81ucM11TtFdoOjYvAWY62Bf1RbTu6V3+Zd595bdsH87vQixpP01nhd1C/6/xFGZH0F8WJw0RDP+pSgK6uxQnMLifEQKfsiKMzVNpGrA0XM26w==
	//ctl00$ContentPlaceHolder1$uscKeputusanParlimen$HiddenField2=melaka
	//ctl00$ContentPlaceHolder1$uscKeputusanParlimen$HiddenField3=Masjid Tanah
	//ctl00$ContentPlaceHolder1$uscKeputusanParlimen$rdoPru=6
	//ctl00$ContentPlaceHolder1$uscKeputusanParlimen$rdoParlimenNav=5
	//ctl00$ContentPlaceHolder1$HiddenField1
	data.Set("ctl00$ContentPlaceHolder1$uscKeputusanParlimen$HiddenField2", "melaka")
	data.Set("ctl00$ContentPlaceHolder1$uscKeputusanParlimen$HiddenField3", "Masjid Tanah")
	data.Set("ctl00$ContentPlaceHolder1$uscKeputusanParlimen$rdoPru", "6")
	data.Set("ctl00$ContentPlaceHolder1$uscKeputusanParlimen$rdoParlimenNav", "5")
	data.Set("ctl00$ContentPlaceHolder1$HiddenField1", "")
	p = p.WithReader(strings.NewReader(data.Encode()))
	p = p.Post("https://pru.sinarharian.com.my/undian/melaka/masjid-tanah")
	// Filter out: ContentPlaceHolder1_uscKeputusanParlimen_rptParlimen_rptParlimenCalon_0_lnkToCalon*
	// for each candidate ..
	//p = p.Match("ContentPlaceHolder1_uscKeputusanParlimen_rptParlimen_rptParlimenCalon_0_lnkToCalon")
	i, err := p.Stdout()
	if err != nil {
		panic(err)
	}
	fmt.Println("READ:", i)

}

func examplePAR() {
	p := script.NewPipe()
	defer p.Close()
	// Raw data looks like this ..
	// {
	//	"__EVENTTARGET": "ctl00$ContentPlaceHolder1$uscKeputusanParlimen$rdoPru$3",
	//	"__EVENTARGUMENT": "",
	//	"__LASTFOCUS": "",
	//	"__VSTATE": "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",
	//	"__VIEWSTATE": "",
	//	"__SCROLLPOSITIONX": "0",
	//	"__SCROLLPOSITIONY": "555",
	//	"__EVENTVALIDATION": "/wEdABKvVXD1oYELeveMr0vHCmYPfjiX7FFf0MJkoirIsIEK3E9Tycr57kum+sDlA7xMrwRXpgzk1qeIbbs7UICXjXiK1nV2ohCvKna6jgbyA5oyp8B20IVHcLGGpeHVvNzxaZX94t14uLkyce/U+ST7y7+R+lUWuAch62jh/2FBC2Fnd8T6HbMFOjyOD/wCSwNpSxnUcrI8f+67haW8kYJzzOPnuDqKweB2slFuiUnzXJep9VFeKBQGCETivikDUr9a+pnti7fcgdy5I/l0+1ILAMWVMYdPUi2QWAkG1OdXxHiJfmapT5EEJ94JFq1Ypzsoo81ucM11TtFdoOjYvAWY62Bf1RbTu6V3+Zd595bdsH87vQixpP01nhd1C/6/xFGZH0F8WJw0RDP+pSgK6uxQnMLifEQKfsiKMzVNpGrA0XM26w==",
	//	"ctl00$ContentPlaceHolder1$uscKeputusanParlimen$HiddenField2": "melaka",
	//	"ctl00$ContentPlaceHolder1$uscKeputusanParlimen$HiddenField3": "jasin",
	//	"ctl00$ContentPlaceHolder1$uscKeputusanParlimen$rdoPru": "6",
	//	"ctl00$ContentPlaceHolder1$uscKeputusanParlimen$rdoParlimenNav": "5",
	//	"ctl00$ContentPlaceHolder1$HiddenField1": ""
	//}
	// curl: curl 'https://pru.sinarharian.com.my/undian/melaka/jasin' -X POST -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:108.0) Gecko/20100101 Firefox/108.0' -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8' -H 'Accept-Language: en-US,en;q=0.5' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/x-www-form-urlencoded' -H 'Origin: https://pru.sinarharian.com.my' -H 'Alt-Used: pru.sinarharian.com.my' -H 'Connection: keep-alive' -H 'Referer: https://pru.sinarharian.com.my/undian/melaka/jasin' -H 'Cookie: _pbjs_userid_consent_data=3524755945110770; _ga=GA1.3.1394770160.1668526747; __gads=ID=98aa206fd2884176-22fa26aa6fd80093:T=1668526748:S=ALNI_MZrtF2acNw0zgHLoj7HHIA1S6badA; __gpi=UID=00000b7cdf373dec:T=1668526748:RT=1671951221:S=ALNI_Masg_Ml3sLJulJaAcL0MXq5KKCxPA; _cc_id=f9ff4cd9448e991ce6c18291993e570; cto_bundle=JiPrqV82dDNiU2MxRlZuNyUyQnBQdmJHWkFaJTJCdnFlb0x5bnR1JTJCUXFNdm5xTHNvdVB3azJBUDNIT0Jjb25mWTZ1RWdrVFJmNkJsUGQlMkJFSFNuRm9vdWw1TGVsT3NpYWdYcFV1cCUyQjVRSENFJTJCdXlRdTlPTXJNU1V2blY4S09NWjVhT2NuQW9qcUFoZHlJM2lHZSUyQldSVWxaTkt4WWk5dyUzRCUzRA; __atuvc=0%7C48%2C0%7C49%2C0%7C50%2C3%7C51%2C1%7C52; ___nrbi=%7B%22firstVisit%22%3A1668573680%2C%22userId%22%3A%2255f51141-0f27-4cbf-aa65-9b51475093f6%22%2C%22userVars%22%3A%5B%5D%2C%22futurePreviousVisit%22%3A1668573680%2C%22timesVisited%22%3A1%7D; compass_uid=55f51141-0f27-4cbf-aa65-9b51475093f6; _gcl_au=1.1.298270835.1668573681; _ga_4Y13Y7JWS5=GS1.1.1668573681.1.1.1668574227.60.0.0; _ga_Y3X3QKNCMC=GS1.1.1668573681.1.1.1668574227.60.0.0; _ga_NJ6319DRLB=GS1.1.1668573681.1.1.1668574227.60.0.0; _ga_YJQ7PQLFWV=GS1.1.1668573681.1.1.1668574227.60.0.0; _ga_JFLF6LQ1DE=GS1.1.1668573681.1.1.1668574227.60.0.0; _ga_0BGT5MH2FW=GS1.1.1668573681.1.1.1668574227.60.0.0; iUUID=52b49239a2026d8cbccc55a3a9021fb2; _fbp=fb.2.1668573682436.1095233389; _poool=958dfac0-bf02-46a8-8926-0d1a7736a1f9; _clck=d9gygo|1|f6m|0; _ga_QXGHHKXNYZ=GS1.1.1668573791.1.1.1668574227.60.0.0; __atssc=google%3B11; _ga_E6VCT2NCRB=GS1.1.1668591503.2.0.1668591664.60.0.0; am_FPID=bd791cc8-587f-427d-8ed7-7c16acfed598; _gid=GA1.3.1265016033.1671951220; panoramaId_expiry=1672037624372; ASP.NET_SessionId=0uywchgsicey0bfire13tt3c; __atuvs=63a7f38d73d4fc3c000; _gat_UA-6733299-4=1; FCNEC=%5B%5B%22AKsRol9gArGZ_QWyLpx7OdtG4RRaHeq2Xhb7Ark0jr3Ri8Sla8sfNidJmwu-1I2tyiY_heftwLiCbSCibxjKlDp-GQMBrR67ei7PZTI4oPylBtsM7HawRbnMtIkJIexSWelrtjvgZByBiJBPudZkVGeWwf_ykOwCmA%3D%3D%22%5D%2Cnull%2C%5B%5D%5D' -H 'Upgrade-Insecure-Requests: 1' -H 'Sec-Fetch-Dest: document' -H 'Sec-Fetch-Mode: navigate' -H 'Sec-Fetch-Site: same-origin' -H 'Sec-Fetch-User: ?1' -H 'TE: trailers' --data-raw '__EVENTTARGET=ctl00%24ContentPlaceHolder1%24uscKeputusanParlimen%24rdoPru%243&__EVENTARGUMENT=&__LASTFOCUS=&__VSTATE=H4sIAAAAAAAEAO1aW3Pb1hE2IYK6i7Jj0ddaJ6lTtY0gCBRv6tiZISXXkmVxOKKUTJ48hwJEQgRBDi5x5f%2FQX9CXzPShP6JvmXamz%2F1D7e7B7QAEFYlOHjq1M2IOsLvf7jl7zu6eJf%2BTyT8Uxdq5Y2xvP98bmo5mOi2DnmsHQ0PVLOW5a58faSPXcW1qtqhl6APNfP7a0tVvdO29ku%2Bq4heThY%2BpbiJfPq9eqOLOzbTsuwkFU8gVUW7aWaFwXlyUlFK1VN1VqlUlrxaEC%2FwQsvCZEzLwmRVm8wXh2eK7zw8dbbA3dE1HWC0Qjy2Tzxeyzxab9Hu9Sx3tzDLEFdk1VZ2a8kgDffaz7Kn2J0fMtdiTqjJMX3A9Iy4H3H1Npb11QRSPcAB8QoxvNeIzqOlQE1jnjvwxcM%2FEuO8F3I5maWa3S00X%2BBdOwyeQyMYk7odWuwZ1pZFuUrMLMkstfCYt9gxSYrr1MFfaR%2BtbOAC%2BXIwvWhPa83BzLTYCztn0edo4t%2B7Qwnm2%2FTFwz6Vb3XepQSXDHYxclFg6wmfylj2D1HxM6m40V8eil%2FSKgsh8K3gA%2FoUY%2F4OA39S6mqVLtjbo6AZzQb7JXpG2%2FwpkF9NnPoA59FFR7piNgHMpfS0vhz02bfENDoBvOR3RoB2XGZF7y0bAuZKOaNOOt7faOAC%2BfIwvH%2FFZ9D3z42zbG7L5FIT1O8Jy4XG05QWQKoysvu9PcH8fxM97uqmzE7SWudc6OSL7Z03S3CrukD1G4bY%2BAnyOAIGb5Z5ruBKOqQMv4Zhe6T0fa43DKuFSe0TuhCDckxhcB%2F6nwycs9RC22ThQlTQCInd4QrsSrpYtGFBXtvC8uT7cZzG4E4%2FEnSsE%2BzJm1UhzqKGbXQn3GbxmGhyd%2BoCFCHCnSNpsXyGVO3WI%2BTyG2R86VMKdK9surJ4ugStUiDQe5OMIsrRL2oyDHHkc3CFF2KfMoXh%2BZYgpl7Ay0gCsHcg2fPhwSwFcG99xZxcBHjG7cIuhB0dDy6Ryx%2B3ikRpbsEqFNDwSd6QjK7xtdU4HmjU0pZ7e7eHhCib1HHFa9ZO3h8evmqS1tV2tkT2PlxyEvNypR2CS5lYw0pFU%2FbxvD00ffD0BroArWsBG9n02LjqE54CdWbZssMIQPak%2BCa1SJqceGwRVYOPiBaLdjdawrw%2BoGXckD1StkCOfg4skiHGP8wOsBMUN4YE8SYLUKqQdsMAxwEQ3%2B41u6x1D62Hqm%2FEOfzzfxQI6S2ssoH9KcJ8S3P9ighM%2FJbhPCe5Tgvs%2FSHC4y3LBIQ8SW%2BLQJjIbJkHvkjd3OKBddsO7L%2BuDrnxh0K7sJb2tkdl9tlw3IBmZcAv0LnyjIB8KYFeQHRNHOp4cQ13rWfGzSAfLl6hiXRTFfpg81zJ%2BJk2c7rFEysMWeFiP7CPP9aNci%2BBR7k0c%2BJTUy2t4GGmIsrGvY8Hh8zNq4TN2IgikJmxeU8wNSPS1LMUEPD3xLJ%2BIC%2FEkP8ELLO8HXhj5RQDzglcRJEJFoiCYZDcj%2Bqi5UVAxeDvGe0pEj7HyYYJ3g4oi8K4dVRjMu2HFkQgiqQUHr%2BMRt4OQxStBgqWPSflLHytVErFlvFLhdT2IdIX0rcsRapofcdUMquHKm0TEmVTd8JqeRpq8gieod%2FyJ5ZOyntKx6igRqBLF0YRd4NVLwS4YBNUT2wV%2BLcWVSMJYKTVhx7LqKtixl36pxXasV3dxxZQwXnZNsNUj%2Bk7IGUFdxmz1qzSu%2BBLGirQJtrK6LbDV9os4ZqtX0QmrMcxkQcejrvGojOrjztphyYfIYQXIEkAuiPJ7tr1nUNsWZ%2FcMyIea%2Biz77vN2QxBU74KUXc%2BJs4emDmWLsT7LXmPcygZpRS3MhQGTpZhcYekiHxkoiNmWslPCei6%2BOeQBtS91VYKQS3twD8vvDQeQ%2FdS61XWhIHTErKLUlMAIuAIcM35MtlD7%2FhTc%2BlwgnxkzpoylTUKaGkNL6sJ5imSLke6FOpDJayTfWFYY01sB2ccJWTAXsoHUoY4bCu9EihdPGR1KW8e9jXRSdRWEf5UQxnArOS7tseI1lC9F2lcOgIWcAgt5g5eq22Bkx2yopUyA1breOBQuc9M%2FAjrxYsLNpcUx1bt8pPeFL6mtm6FYJVIqvkHKTSRy%2BdV8IfdseZ86FIugP%2BqaoYpL%2B7o9MujVuyZUsM9WkPgNNVzNowqH%2B9hbx7eNIeB31dU1uGWfnCkK%2Byyyzx32WWKf5TUxo2SKmVKmIijK%2Fa9mxC7%2BK2SELGbLVVUtZC78RHMRtvGFcONHgUTwFJXZ0WaHNeNzLHjNkElxmnUX1gR2kmPnMLxNZgtzfMckJy6MTOm9bhL4Y2EDC9E1QRiZQqupRlczuBIYQ1Mu1mp4gCXNGujaFXUg89CB7ap6UOgWjutt8urk%2BPDVd%2FXTQ9KuH7fP9g%2BbfhMH78zF8mZlu8T1YwKVHVNocCp%2F7aus7io7Mu2oriH1KJTmkjd%2BT3u66it9UG%2Fsn70lB%2FWjw2Pijb%2BtHxzu81qVTWV3h%2BvuhBPtCa2DSOvjSGtJHrh4a%2BxIrtMbhKX8yvHZaf3tYYOcnR4c17mZidXNUqnMtYR8FTNdh868Pq2nTq2KR1OFgK9j0Ddhci6EiT61gqkd1Jv7J%2FVvD%2BGv3oTJnZFG%2Fah%2BEqmdKW9XWTfH2xdendvRTBXKP7k3HGhyq%2BnvjfnA0tJmSVHUeNU%2F02Mpwx9yO%2FLFV3fuCpnVwh38b%2Bn%2BV3fuRFNMo81HcJGLVy%2BwhyeyAxbWEqAfOJfEJT4n4BbGt1pUDrOd%2B7RQi%2BerWdzlUlOBBf3tNTkmqBjh9t%2FVvahQLJa3uYTlF4KMYVqwTJppRUD7%2FXVowXW%2Fo6uQ9EM8LpuunHospMFYpgcU0gzcAbzfXIdHrzQLZjugbogVS7dAhoUD8rRAKUaVWLy5BsuAzR3BcFlYfIuUW4tn06wo%2F9ROcGhn2JPg3gZ7IUTjkvLSKTKQJmOYFkwcM60sNbE%2BIRNrG9Yx0yPfc2k6d8JItxXPpRmBlcr6ZJQuVbV%2BCMJn7ddIuaXwbJoFWKd8ORlkQM97dMTVOwBV5UqWY0YPCqbbA82l2bSb4mgOSnUtIEAhhuWgEaLVuPO%2Bz1igmGMs08LNx4yb84xTtq%2F3PLXxah%2FC7HIbp85ItxVfGLOiAlZgwN6YXCAH3VRXpdFZqPDx2m%2BmniHDlFiLqZYVU5abRxth40mX%2BsAxtEI4LlovtxgHOWIc06ItpRqXFqvjcFdUsrCb4YRgXKxebAGdnDD6dEjLqWaVWD0zGayvGVqHX30uXmMvjxGnwFgZs6YK1mDU3rz27iOPNLPbhyRuctcxgOSvU62AJbjPTQ2ZT7WykrIzEpAsX%2FY1a9gL4crJxHuE5GmhVlMtw6j%2Bu%2BvhOm5fd6Q%2B1P5hxKlw8X2xgXRyhPSpwe6mGle74VwHQyNKH5VqctmOkTwt1L0xy2pg2W7KBuYuvbCBbT0WgmqxQ%2BARp8D4LM2a4nbKuseQcAwjuLpB8RHicfF%2Bhd3n3wYsU%2BPdT7VPSanQeDzVvRpG86xyoT%2B3z0i3ll9LtaOYkjx4nA5%2BQWNJPd3Qw%2Fhc5TtdDcZADpBhSqxCqmUY659fg%2BZoxrAvQQUZQnGRfv4UqQSu%2F9OgPBgzaBcMSusLsjYL3DQHnSjrV7nonmsw0i0kH6bqxpj%2BMBXB0gehOBe%2FZ070wU1lHqWqrEw02tasPo0Q%2BOK6zUi3kHycqjutBeghDPC7Ou4GVeXC7tyxT7yV9JNUCzDOfpE%2BB6%2Bcsih4Lrz%2BVrkIu%2BzXZieM4%2FY4c6xTtz4fb86tL2A7bn3x2h7ccuuk6bcgCd%2BLy9QSrbiFsBUX%2F1JkNWGqHPyug2%2BhpzXrVlAzfrse%2FsAk0bVbuVHXTmgqsVbENT27TmrPLtFAC1tZynZVtoa2oXutM4MGv%2B9YOcG3pB68jVpZO5vlclqPLtEtexq2BZUd%2BdI14Aow0iGB9sHLwS8XVt%2BE76EqwPeRGmWzqtTSmnLx7uOTcCalskw%2FoIbBsKdKNqVBCzBfx9eQ23sqaVOqxpRU0tpyYgdqSltsvGrU25GqR4Gq4rYiw3nVbajqRtTQDV%2FR8pH3krTYy1CNUFau78I1xrpwymZtpxL8knXJ%2B%2F4Fy2z2BLshpT%2FnzcHfXv%2BBf%2FmfvWcXLZNw4ekRCwsX%2Fub3LM%2FuFJUKt8T%2FyrxQ9e%2FJOX4%2F9XIDZr7x9QtKepZ28XJD1s2LodTRLN2hMsrJPVeXrvC3OA7F7NPzGp4O8Bh6kJskVbfhAf761AzRvn4B2MS2zl9u9Bxn9AdZHlnuFgQUavUoXoW3zoeDrcGVDLdhUAhEGS4pu%2B8MDTtooJegXnTDBpwd1em93NjZ%2FhLgZQp%2FMIWvwxaoIG7xc1IdqO81TOEbXyu7pDn8HpIX1JLF7aISSs76kn%2BNLYc3dcnRHUP7udflAKb0HYgQECEoQpgIYSJBMORFYjPNJJ1aLHJO%2FefNnVosyud6v%2BtK9IOlm5KjQRCUbLi02RDkdRve6xZYLgUp6OMcWlXe2a5NmEbCNP7iHv3LFB69%2FaLssRnVkZ0wduKzE2AnwA6uJEGej81QSHhSqe5ynvzHjT0JcrJrO%2FSDxIL5jz9cUvNSl0aa5XawuQP2%2FvgD3wL%2FyKNZZp6MEsdHO7L28ztyijU5Q3byhmMnITtBdj7Nx2Y4k3RkbZtz5L9v7sjaNmQvDCMf%2FJ2H2iV7aOtDrT80hwOcgam6qtuXot78R7qzwtzZZ7HoA0FYwmB%2Fcbf%2BbQq3TrtCRxTvVh%2F8E4pChBcigRCJvhyJzTabdPGuwrn47zd38a7Cfvgr4R506HvqWSh529M1HbDbu2J9pFerXrgFVaRHPz7W%2FpQv%2FzyFL2%2BxFPiLaHIA55Dxsf1JvCPK%2BIh3aeXn9F%2FLG6U%2FozgAAA%3D%3D&__VIEWSTATE=&__SCROLLPOSITIONX=0&__SCROLLPOSITIONY=555&__EVENTVALIDATION=%2FwEdABKvVXD1oYELeveMr0vHCmYPfjiX7FFf0MJkoirIsIEK3E9Tycr57kum%2BsDlA7xMrwRXpgzk1qeIbbs7UICXjXiK1nV2ohCvKna6jgbyA5oyp8B20IVHcLGGpeHVvNzxaZX94t14uLkyce%2FU%2BST7y7%2BR%2BlUWuAch62jh%2F2FBC2Fnd8T6HbMFOjyOD%2FwCSwNpSxnUcrI8f%2B67haW8kYJzzOPnuDqKweB2slFuiUnzXJep9VFeKBQGCETivikDUr9a%2Bpnti7fcgdy5I%2Fl0%2B1ILAMWVMYdPUi2QWAkG1OdXxHiJfmapT5EEJ94JFq1Ypzsoo81ucM11TtFdoOjYvAWY62Bf1RbTu6V3%2BZd595bdsH87vQixpP01nhd1C%2F6%2FxFGZH0F8WJw0RDP%2BpSgK6uxQnMLifEQKfsiKMzVNpGrA0XM26w%3D%3D&ctl00%24ContentPlaceHolder1%24uscKeputusanParlimen%24HiddenField2=melaka&ctl00%24ContentPlaceHolder1%24uscKeputusanParlimen%24HiddenField3=jasin&ctl00%24ContentPlaceHolder1%24uscKeputusanParlimen%24rdoPru=6&ctl00%24ContentPlaceHolder1%24uscKeputusanParlimen%24rdoParlimenNav=5&ctl00%24ContentPlaceHolder1%24HiddenField1='
	data := url.Values{}
	data.Set("ctl00$ContentPlaceHolder1$uscKeputusanParlimen$HiddenField2", "melaka")
	data.Set("ctl00$ContentPlaceHolder1$uscKeputusanParlimen$HiddenField3", "jasin")
	data.Set("ctl00$ContentPlaceHolder1$uscKeputusanParlimen$rdoPru", "6")
	data.Set("ctl00$ContentPlaceHolder1$uscKeputusanParlimen$rdoParlimenNav", "5")
	p = p.WithReader(strings.NewReader(data.Encode()))
	p = p.Post("https://pru.sinarharian.com.my/undian/melaka/jasin")
	// Filter out: ContentPlaceHolder1_uscKeputusanParlimen_rptParlimen_rptParlimenCalon_0_lnkToCalon*
	// for each candidate ..
	p = p.Match("ContentPlaceHolder1_uscKeputusanParlimen_rptParlimen_rptParlimenCalon_0_lnkToCalon")
	i, err := p.Stdout()
	if err != nil {
		panic(err)
	}
	fmt.Println("READ:", i)
	//req := http.NewRequest(http.MethodPost, "")
	//p := script.Do(&req)
	//p.Post("")
	//script.Post("")

}

func matchCandidatesName(parID string, c *[]candidate) error {
	candidateDir := fmt.Sprintf("testdata/%s", parID)
	// Load up the data parID
	// each must find in the candidate ..
	// Look for all files in dataPath
	candidateFiles, err := script.ListFiles(candidateDir).Slice()
	if err != nil {
		panic(err)
	}
	//spew.Dump(candidateFiles)
	for _, dataFilePath := range candidateFiles {
		candidateFilePath := ""
		//fmt.Println("<<<<<<<<<<<<", parC.name, "in", candidateDir, ">>>>>>>>>>")
		// DEBUG
		//fmt.Println("Look for:", safeName, "  in", dataFilePath)
		fmt.Println(dataFilePath)
		for i, candidate := range *c {
			safeName := strings.ReplaceAll(strings.ToLower(candidate.name), " ", "-")
			// Exact match ..
			if strings.Contains(dataFilePath, safeName) {
				fmt.Println("MATCH_1")
				candidateFilePath = dataFilePath
				if (*c)[i].matchedName != "" {
					return fmt.Errorf("EXISTING: %s CUR: %s", (*c)[i].matchedName, candidateFilePath)
				}
				(*c)[i].matchedName = safeName
				(*c)[i].matchURL = candidateFilePath
				break
			}

			for _, namePart := range strings.Split(safeName, "-") {
				// Skip common name like BIN BINTI A/L A/P?
				if strings.ToUpper(namePart) == "BIN" {
					fmt.Println("Skipping common namePart - BIN")
					continue
				}

				if strings.ToUpper(namePart) == "BINTI" {
					fmt.Println("Skipping common namePart - BINTI")
					continue
				}

				if strings.ToUpper(namePart) == "MOHD" {
					fmt.Println("Skipping common namePart - MOHD")
					continue
				}

				// DEBUZg
				//spew.Dump(namePart)
				if strings.Contains(dataFilePath, namePart) {
					fmt.Println("MATCH_2")
					candidateFilePath = dataFilePath
					if (*c)[i].matchedName != "" {
						return fmt.Errorf("EXISTING: %s CUR: %s", (*c)[i].matchedName, candidateFilePath)
					}
					(*c)[i].matchedName = safeName
					(*c)[i].matchURL = candidateFilePath
					break
				}

			}
		}
	}

	spew.Dump(c)

	// If already matched before; it is a FATAL error!!

	return nil
}

func ExtractCandidatePerPAR(state string, pars []string) {
	// Load all Results ..
	candidatesPAR := LookupResults(state)
	// DEBUG
	//spew.Dump(candidatesPAR)
	// maybe no need
	//mapCandidate = make(map[string][]candidate, len(pars))
	// For each PAR
	for _, par := range pars {
		// Derive PAR_ID
		parID := fmt.Sprintf("%s00", par[1:])
		fmt.Println("PAR:", parID)
		// For each ballotID; find the match first?
		// load the file; safeName is encoded ..

		candidatesInPAR := candidatesPAR[parID]
		err := matchCandidatesName(parID, &candidatesInPAR)
		if err != nil {
			panic(err)
		}
		// Then run another round?
		//var candidates []candidate
		//// Append all candidates by BallotID order
		//// Add into map by PAR_ID
		//mapCandidate[parID] = candidates
		//
		fmt.Println("After ....")
		spew.Dump(candidatesPAR[parID])
		break
	}

	// For each mapKey; dump it all out!
	//spew.Dump(mapCandidate)
}

func downloadCandidates(par string, calons []string) {
	fmt.Println("PAR:", par)
	baseURL := "https://pru.sinarharian.com.my"
	// PAR == P123 --> 123
	code := par[1:]
	parID := fmt.Sprintf("%s00", code)
	dataPath := fmt.Sprintf("testdata/%s", parID)
	/*
		([]string) (len=3 cap=4) {
		 (string) (len=153) "<a id=\"ContentPlaceHolder1_uscKeputusanParlimen_rptParlimen_rptParlimenCalon_0_lnkToCalon_0\" class=\"calon\" href=\"/calon/6151/shamsul-iskandar-mohd-akin\">",
		 (string) (len=141) "<a id=\"ContentPlaceHolder1_uscKeputusanParlimen_rptParlimen_rptParlimenCalon_0_lnkToCalon_1\" class=\"calon\" href=\"/calon/858/mohd-ali-rustam\">",
		 (string) (len=143) "<a id=\"ContentPlaceHolder1_uscKeputusanParlimen_rptParlimen_rptParlimenCalon_0_lnkToCalon_2\" class=\"calon\" href=\"/calon/2820/md-khalid-kassim\">"
		}
	*/
	// Create data folder
	merr := os.MkdirAll(dataPath, 0755)
	if merr != nil {
		panic(merr)
	}
	// Go through each candidate .. with these pattern
	rexp := regexp.MustCompile("^<a.*href=\"(.+/(.+?))\".*$")
	replaceTemplate := "$1,$2"
	for _, calon := range calons {
		s := rexp.ReplaceAllString(calon, replaceTemplate)
		// Extract out from template ..
		/*
			(string) (len=56) "/calon/288/mas-ermieyati-samsudin,mas-ermieyati-samsudin"
			(string) (len=37) "/calon/6219/nasir-othman,nasir-othman"
			(string) (len=43) "/calon/6217/sabirin-ja`afar,sabirin-ja`afar"
			(string) (len=49) "/calon/6066/mohd-redzuan-yusof,mohd-redzuan-yusof"
		*/
		c := strings.Split(s, ",")
		candidateURL := baseURL + c[0]
		safeName := c[1]
		candidatePath := fmt.Sprintf("%s/%s.html", dataPath, safeName)
		// DEBUG
		//spew.Dump(s)
		// Get file and save it ..
		// If the file exists; ignore it again??
		// testdata/<PAR>/...html
		if script.IfExists(candidatePath).Error() != nil {
			// has error; means the file does not exist!
			fmt.Println("GOTTA DOWNLOAD!!!", candidateURL, "INTO", candidatePath)
			n, err := script.Get(candidateURL).WriteFile(candidatePath)
			if err != nil {
				panic(err)
			}
			fmt.Println("N:", n)
		} else {
			fmt.Println("FOUND! at", candidatePath)
		}

	}
}

// DownloadCandidatePerPAR go from raw PAR to PAR_ID
func DownloadCandidatePerPAR(state string, pars []string) {
	// For each candidate; filter per PAR??

	// For each Daerah Mengundi saluran
	//	file: testdata/Saluran-<STATE>.csv
	// Create pru14-<state>.txt
	for _, par := range pars {
		lines, err := script.File(fmt.Sprintf("testdata/%s.html", par)).Match("/calon").Slice()
		if err != nil {
			panic(err)
		}
		// DEBUG
		//spew.Dump(lines)
		downloadCandidates(par, lines)
	}
}
