package domain

import "testing"

func TestNormalizeAddress(t *testing.T) {
	// NormalizeAddress lowercases, trims, and guarantees a single 0x prefix.
	cases := map[string]string{
		"0xAbCdEf0000000000000000000000000000000001": "0xabcdef0000000000000000000000000000000001",
		"0xabcdef0000000000000000000000000000000001": "0xabcdef0000000000000000000000000000000001",
		"0XABCDEF0000000000000000000000000000000001": "0xabcdef0000000000000000000000000000000001",
		"abcdef0000000000000000000000000000000001":   "0xabcdef0000000000000000000000000000000001",
		"  0xAbC  ": "0xabc",
	}

	for in, want := range cases {
		if got := NormalizeAddress(in); got != want {
			t.Errorf("NormalizeAddress(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestChecksumAddress(t *testing.T) {
	// ChecksumAddress returns the EIP-55 mixed-case form; empty input stays empty.
	cases := map[string]string{
		"0x5aaeb6053f3e94c9b9a09f33669435e7ef1beaed": "0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed",
		"0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed": "0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed",
		"0X5AAEB6053F3E94C9B9A09F33669435E7EF1BEAED": "0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed",
		"":    "",
		"   ": "",
	}

	for in, want := range cases {
		if got := ChecksumAddress(in); got != want {
			t.Errorf("ChecksumAddress(%q) = %q, want %q", in, got, want)
		}
	}
}
