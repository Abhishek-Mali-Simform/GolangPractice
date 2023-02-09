package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	p := newParser()

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		parsed := parse(p, in.Text())
		update(p, parsed)
	}
	summarize(p)
	dumpErrs([]error{in.Err(), err(p)})
}

func summarize(p *parser) {
	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))
	sort.Strings(p.domains)
	for _, domain := range p.domains {
		fmt.Printf("%-30s %10d\n", domain, p.sum[domain].visits)
	}
	fmt.Printf("\n%-30s %10d\n", "TOTAL", p.total)
}
func dumpErrs(errs []error) {
	for _, err := range errs {
		if err != nil {
			fmt.Println("> Err:", err)
		}
	}
}
