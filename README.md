# Taiwan ID Validator (twid)

[![Go Test](https://github.com/guanting112/taiwan-id-validator/actions/workflows/go.yml/badge.svg)](https://github.com/guanting112/taiwan-id-validator/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/guanting112/taiwan-id-validator.svg)](https://pkg.go.dev/github.com/guanting112/taiwan-id-validator)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

`twid` is a lightweight and zero-dependency Go package for validating Taiwan National IDs, Alien Resident Certificates (ARC), and Unified Business Numbers (UBN).

It supports:
- **Taiwan National ID** (中華民國身分證字號)
- **New Alien Resident Certificate** (新式外來人口統一證號) (Since 2021)
- **Old Alien Resident Certificate** (舊式外來人口統一證號)
- **Unified Business Number (UBN)**: Validates company tax IDs, fully supporting the latest logic (including the 7th-digit rule). (新舊格式的統一編號)

## Installation

```bash
go get github.com/guanting112/taiwan-id-validator
```

## Usage

```go
package main

import (
	"fmt"

	twid "github.com/guanting112/taiwan-id-validator"
)

func main() {
	// National ID
	fmt.Println(twid.Validate("A123456789"))         // true
	fmt.Println(twid.ValidateNationId("A123456789")) // true
	fmt.Println(twid.ValidateNationId("A123456700")) // false

	// Taiwan Company UBN
	fmt.Println(twid.ValidateUbn("84149961")) // true
	fmt.Println(twid.ValidateUbn("22099131")) // true
	fmt.Println(twid.ValidateUbn("!@)(#&6y01)")) // false
	fmt.Println(twid.ValidateUbn("25317520")) // false
	fmt.Println(twid.ValidateUbn("00000000")) // false

	// New ARC
	fmt.Println(twid.Validate("A800000014"))      // true
	fmt.Println(twid.ValidateArcId("A800000014")) // true
	fmt.Println(twid.ValidateArcId("A800000015")) // false

	// Old ARC
	fmt.Println(twid.Validate("AC01234567"))      // true
	fmt.Println(twid.ValidateArcId("AC01234567")) // true

	// Invalid
	fmt.Println(twid.Validate("A123456788")) // false
}
```
