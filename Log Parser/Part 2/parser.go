package main

import (
	"fmt"
	"strconv"
	"strings"
)

type result struct {
	domain string
	visits int
}

type parser struct {
	sum     map[string]result
	domains []string
	total   int
	lines   int
	lerr    error
}

func newParser() *parser {
	return &parser{sum: make(map[string]result)}
}
func parse(p *parser, line string) (parsed result) {

	if p.lerr != nil {
		return
	}
	p.lines++
	fields := strings.Fields(line)
	if len(fields) != 2 {
		p.lerr = fmt.Errorf("Wrong input: %v (line #%d)", fields, p.lines)
		return
	}
	var err error

	parsed.domain = fields[0]
	parsed.visits, err = strconv.Atoi(fields[1])

	if parsed.visits < 0 || err != nil {
		p.lerr = fmt.Errorf("Wrong input: %v (line #%d)", fields, p.lines)
	}
	return
}

func update(p *parser, parsed result) {
	if p.lerr != nil {
		return
	}

	if _, ok := p.sum[parsed.domain]; !ok {
		p.domains = append(p.domains, parsed.domain)
	}

	p.total += parsed.visits

	p.sum[parsed.domain] = result{
		domain: parsed.domain,
		visits: parsed.visits + p.sum[parsed.domain].visits,
	}
}

func err(p *parser) error {
	return p.lerr
}
