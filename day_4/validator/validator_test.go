package validator

import "testing"

func TestSimpleValidate(t *testing.T) {
	useCases := map[string]struct {
		passphrase string
		isValid    bool
	}{
		"aa bb cc dd ee is valid": {
			passphrase: "aa bb cc dd ee",
			isValid:    true,
		},
		"aa bb cc dd aa is not valid": {
			passphrase: "aa bb cc dd aa",
			isValid:    false,
		},
		"aa bb cc dd aaa is valid": {
			passphrase: "aa bb cc dd aaa",
			isValid:    true,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			result := SimpleValidate(uc.passphrase)
			if result != uc.isValid {
				t.Fatalf("Expected result was `%t`, got `%t`", uc.isValid, result)
			}
		})
	}
}

func TestComplexValidate(t *testing.T) {
	useCases := map[string]struct {
		passphrase string
		isValid    bool
	}{
		"abcde fghij is valid": {
			passphrase: "abcde fghij",
			isValid:    true,
		},
		"abcde xyz ecdab is not valid": {
			passphrase: "abcde xyz ecdab",
			isValid:    false,
		},
		"a ab abc abd abf abj is valid": {
			passphrase: "a ab abc abd abf abj",
			isValid:    true,
		},
		"iiii oiii ooii oooi oooo is valid": {
			passphrase: "iiii oiii ooii oooi oooo",
			isValid:    true,
		},
		"oiii ioii iioi iiio is not valid": {
			passphrase: "oiii ioii iioi iiio",
			isValid:    false,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			result := ComplexValidate(uc.passphrase)
			if result != uc.isValid {
				t.Fatalf("Expected result was `%t`, got `%t`", uc.isValid, result)
			}
		})
	}
}
