package lom

import "github.com/bitfield/script"

// Laws of Malaysia Monitoring ...
// Look at current month from the last run
// Rollup to week + months? Start new child workflow?

func downloadPUA() {
	p := script.Post("")
	if err := p.Error(); err != nil {
		panic(err)
	}
}
