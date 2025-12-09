package twid

import (
	"regexp"
	"strings"
)

var (
	ubnWeights    = []int{1, 2, 1, 2, 1, 2, 4, 1}
	letterWeights = map[rune]int{
		'A': 10, 'B': 11, 'C': 12, 'D': 13, 'E': 14, 'F': 15, 'G': 16,
		'H': 17, 'J': 18, 'K': 19, 'L': 20, 'M': 21, 'N': 22, 'P': 23,
		'Q': 24, 'R': 25, 'S': 26, 'T': 27, 'U': 28, 'V': 29, 'X': 30,
		'Y': 31, 'W': 32, 'Z': 33, 'I': 34, 'O': 35,
	}
	nationalIdRegExp = regexp.MustCompile(`^[A-Z][12]\d{8}$`)
	newArcIdRegExp   = regexp.MustCompile(`^[A-Z][89]\d{8}$`)
	oldArcIdRegExp   = regexp.MustCompile(`^[A-Z][A-Z]\d{8}$`)
	ubnRegExp        = regexp.MustCompile(`^\d{8}$`)
)

// ValidateArcId checks if the given ID is a valid Taiwan Alien Resident Certificate number
func ValidateArcId(id string) bool {
	id = strings.ToUpper(id)
	switch {
	case newArcIdRegExp.MatchString(id):
		return validateModernIdCardFormat(id)
	case oldArcIdRegExp.MatchString(id):
		return validateOldArcIdFormat(id)
	default:
		return false
	}
}

// ValidateNationId checks if the given ID is a valid Taiwan National
// ID.
func ValidateNationId(id string) bool {
	id = strings.ToUpper(id)
	if nationalIdRegExp.MatchString(id) {
		return validateModernIdCardFormat(id)
	} else {
		return false
	}
}

// Validate checks if the given ID is a valid Taiwan National
// ID or Alien Resident Certificate number.
//
// It supports:
//  1. National ID (ex: A123456789)
//  2. New Alien Resident Certificate (ex: A800000014)
//  3. Old Alien Resident Certificate (ex: AC01234567)
func Validate(id string) bool {
	id = strings.ToUpper(id)
	switch {
	case nationalIdRegExp.MatchString(id):
		return validateModernIdCardFormat(id)
	case newArcIdRegExp.MatchString(id):
		return validateModernIdCardFormat(id)
	case oldArcIdRegExp.MatchString(id):
		return validateOldArcIdFormat(id)
	default:
		return false
	}
}

func validateModernIdCardFormat(id string) bool {
	sum := 0

	areaWeight, ok := letterWeights[rune(id[0])]
	if !ok {
		return false
	}

	sum += (areaWeight/10)*1 + (areaWeight%10)*9
	weights := []int{8, 7, 6, 5, 4, 3, 2, 1}

	for i, char := range id[1:9] {
		num := int(char - '0')
		sum += num * weights[i]
	}

	checkDigit := int(id[9] - '0')
	sum += checkDigit

	return sum%10 == 0
}

func validateOldArcIdFormat(id string) bool {
	sum := 0

	areaWeight, ok := letterWeights[rune(id[0])]
	if !ok {
		return false
	}

	sum += (areaWeight/10)*1 + (areaWeight%10)*9

	genderWeight, ok := letterWeights[rune(id[1])]
	if !ok {
		return false
	}
	sum += (genderWeight % 10) * 8

	weights := []int{7, 6, 5, 4, 3, 2, 1}

	for i, char := range id[2:9] {
		num := int(char - '0')
		sum += num * weights[i]
	}

	checkDigit := int(id[9] - '0')
	sum += checkDigit

	return sum%10 == 0
}

// ValidateUbn checks if the given UBN is a valid Taiwan Company Tax ID.
func ValidateUbn(ubn string) bool {
	if !ubnRegExp.MatchString(ubn) {
		return false
	}

	sum1 := 0
	sum2 := 0

	for i, weight := range ubnWeights {
		if i == 6 && ubn[i] == '7' {
			sum1 += 1
			sum2 += 0
		} else {
			digit := int(ubn[i] - '0')
			product := digit * weight
			sum1 += (product / 10) + (product % 10)
			sum2 += (product / 10) + (product % 10)
		}
	}

	if sum1%5 == 0 || sum2%5 == 0 {
		return true
	}

	return false
}
