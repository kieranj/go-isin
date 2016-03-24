# ISIN

[![Build Status](https://travis-ci.org/kieranj/go-isin.svg?branch=master)](https://travis-ci.org/kieranj/go-isin)

Validate ISINs with Go

## Installation

    go get github.com/kieranj/go-isin

## Usage

    isin, err := isin.NewIsin("US037833100")

    if err != nil {
      fmt.Println(err.Msg)
    }

    isin.Valid() // true

    isin.Checksum() // 5

## License

MIT
