package main

import "testing"

func TestEarthToMarsDate(t *testing.T) {
	dates := []struct {
		in  string
		out MarsDate
	}{
		{"2019-12-27T15:22:22Z", MarsDate{MSD: "51,896.44035", MTC: "10:34:06"}},
		{"2006-01-02T15:04:05Z", MarsDate{MSD: "46,926.06938", MTC: "01:39:54"}},
		{"2019-12-27T15:04:03", MarsDate{MSD: "51,896.42797", MTC: "10:16:17"}},
	}
	for _, date := range dates {
		got, err := EarthToMarsDate(date.in)
		if err != nil {

		}

		if got.MSD != date.out.MSD && got.MTC != date.out.MTC {
			t.Errorf("EarthToMarsDate(%q) == %q, want %q", date.in, got, date.out)
		}
	}

	failed := []struct {
		in  string
		out ResultError
	}{
		{"2006-01-02 15:04:05", ResultError{Message: "Invalid DateTime format provided. Use RFC3339 for DateTime format: \"2019-12-27T15:04:03\""}},
	}
	for _, fail := range failed {
		_, err := EarthToMarsDate(fail.in)
		if err.Message != "Invalid DateTime format provided. Use RFC3339 for DateTime format: \"2019-12-27T15:04:03\"" {
			t.Errorf("EarthToMarsDate(%q) == %q, want %q", fail.in, err, fail.out)
		}
	}
}
