package isin

import (
	"fmt"
	"regexp"
	"strconv"
)

var (
	isinRegexp = regexp.MustCompile("(([A-Z]{2})([A-Z0-9]{9}))(\\d{1})?")
)

// Isin holds deconstructed values
type Isin struct {
	CountryCode string
	Nsin        string
	Identifier  string
	Check       int
}

// FormatError for ISIN format
type FormatError struct {
	Msg string
}

// NewIsin returns a pointer to isin or error if the supplied string doesn't
// match the ISIN format
func NewIsin(str string) (*Isin, *FormatError) {
	isin, error := parse(str)
	if error != nil {
		return nil, error
	}

	return isin, nil
}

func parse(str string) (*Isin, *FormatError) {
	if len(str) < 11 {
		return nil, &FormatError{Msg: "Invalid Format"}
	}

	matches := isinRegexp.MatchString(str)

	if matches == false {
		return nil, &FormatError{Msg: "Invalid Format"}
	}

	match := isinRegexp.FindStringSubmatch(str)

	isin := &Isin{CountryCode: match[2], Nsin: match[3], Identifier: match[1]}

	if match[4] != "" {
		Check, _ := strconv.Atoi(match[4])
		isin.Check = Check
	} else {
		isin.Check = isin.Checksum()
	}

	return isin, nil
}

// Format returns the full ISIN
func (i *Isin) Format() string {
	return fmt.Sprintf("%s%d", i.Identifier, i.Check)
}

// Valid returns true if the calculated checksum matches the supplied checksum
func (i *Isin) Valid() bool {
	return i.Checksum() == i.Check
}

func (i *Isin) digits() []int {
	var digits []int

	for _, value := range i.Identifier {
		switch {
		case string(value) >= "0" && string(value) <= "9":
			digits = append(digits, int(value)-48)
		default:
			d := int(value) - 55
			digits = append(digits, d/10)
			digits = append(digits, d%10)
		}
	}

	return digits
}

// Checksum returns the calculated checksum
func (i *Isin) Checksum() int {
	s := i.digits()

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	sum := 0

	for i, num := range s {
		if i%2 == 0 {
			d := num * 2
			sum += d / 10
			sum += d % 10
		} else {
			sum += num
		}
	}

	return (10 - (sum % 10)) % 10
}
