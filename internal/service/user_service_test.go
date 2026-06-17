package service

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	today := time.Now()

	tests := []struct {
		name     string
		dob      time.Time
		expected int
	}{
		{
			name: "birthday is today",
			dob: time.Date(
				today.Year()-26, // 26 years ago
				today.Month(),   // same month
				today.Day(),     // same day
				0, 0, 0, 0,
				today.Location(),
			),
			expected: 26,
		},
		{
			name: "birthday was yesterday",
			dob: time.Date(
				today.Year()-26,
				today.Month(),
				today.Day()-1, // one day earlier
				0, 0, 0, 0,
				today.Location(),
			),
			expected: 26,
		},
		{
			name: "birthday is tomorrow",
			dob: time.Date(
				today.Year()-26,
				today.Month(),
				today.Day()+1, // one day later
				0, 0, 0, 0,
				today.Location(),
			),
			expected: 25,
		},
		{
			name: "newborn born today",
			dob: time.Date(
				today.Year(),
				today.Month(),
				today.Day(),
				0, 0, 0, 0,
				today.Location(),
			),
			expected: 0,
		},
		{
			name: "person born 2000-01-01",
			dob:  time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: func() int {
				age := today.Year() - 2000
				birthday := time.Date(
					today.Year(), 1, 1,
					0, 0, 0, 0,
					today.Location(),
				)
				if today.Before(birthday) {
					age--
				}
				return age
			}(),
		},
		{
			name: "person born 2000-06-16",
			dob:  time.Date(2000, 6, 16, 0, 0, 0, 0, time.UTC),
			expected: func() int {
				age := today.Year() - 2000
				birthday := time.Date(
					today.Year(), 6, 16,
					0, 0, 0, 0,
					today.Location(),
				)
				if today.Before(birthday) {
					age--
				}
				return age
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateAge(tt.dob)

			if got != tt.expected {
				t.Errorf(
					"calculateAge(%s) = %d, want %d",
					tt.dob.Format("2006-01-02"),
					got,
					tt.expected,
				)
			}
		})
	}
}

func TestCalculateAgeNeverNegative(t *testing.T) {
	futureDate := time.Now().AddDate(10, 0, 0)
	age := calculateAge(futureDate)

	if age < 0 {
		t.Errorf(
			"calculateAge returned negative age %d for future date %s",
			age,
			futureDate.Format("2006-01-02"),
		)
	}
}

func TestCalculateAgeOlderPerson(t *testing.T) {
	today := time.Now()

	dob := time.Date(
		today.Year()-60,
		today.Month(),
		today.Day(),
		0, 0, 0, 0,
		today.Location(),
	)

	got := calculateAge(dob)
	if got != 60 {
		t.Errorf("calculateAge() = %d, want 60", got)
	}
}

func TestDateFormatRoundTrip(t *testing.T) {
	cases := []string{
		"2000-06-16",
		"1990-01-01",
		"1995-12-31",
		"2005-02-28",
	}

	for _, dateStr := range cases {
		t.Run(dateStr, func(t *testing.T) {
			// Parse the string into time.Time
			parsed, err := time.Parse("2006-01-02", dateStr)
			if err != nil {
				t.Fatalf("time.Parse(%q) failed: %v", dateStr, err)
			}

			formatted := parsed.Format("2006-01-02")

			if formatted != dateStr {
				t.Errorf(
					"round trip failed: input=%q output=%q",
					dateStr,
					formatted,
				)
			}
		})
	}
}
