package zoom

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestParseOpenHours(t *testing.T) {
	jsonStr := "{\"mon\":[\"08:00 AM - 01:00 PM\", \"02:00 PM - 04:00 PM\"],\"tue\":[\"08:00 AM - 04:00 PM\"],\"wed\":[\"08:00 AM - 04:00 PM\"],\"thu\":[\"08:00 AM - 04:00 PM\"],\"fri\":[\"08:00 AM - 04:00 PM\"],\"sat\":[\"09:00 AM - 01:00 PM\"],\"sun\":\"0\"}"
	expected := OpenHours{
		Mon: []HourRange{{Start: "08:00 AM", End: "01:00 PM"}, {Start: "02:00 PM", End: "04:00 PM"}},
		Tue: []HourRange{{Start: "08:00 AM", End: "04:00 PM"}},
		Wed: []HourRange{{Start: "08:00 AM", End: "04:00 PM"}},
		Thu: []HourRange{{Start: "08:00 AM", End: "04:00 PM"}},
		Fri: []HourRange{{Start: "08:00 AM", End: "04:00 PM"}},
		Sat: []HourRange{{Start: "09:00 AM", End: "01:00 PM"}},
		Sun: nil,
	}

	result, err := parseOpenHours(jsonStr)
	if err != nil {
		t.Errorf("ParseOpenHours returned an unexpected error: %v", err)
	}

	if diff := cmp.Diff(expected.Mon, result.Mon); diff != "" {
		t.Errorf("ParseOpenHours() Mon = %v, want %v", result.Mon, expected.Mon)
	}

	if diff := cmp.Diff(expected.Tue, result.Tue); diff != "" {
		t.Errorf("ParseOpenHours() Tue = %v, want %v", result.Tue, expected.Tue)
	}

	if diff := cmp.Diff(expected.Wed, result.Wed); diff != "" {
		t.Errorf("ParseOpenHours() Wed = %v, want %v", result.Wed, expected.Wed)
	}

	if diff := cmp.Diff(expected.Thu, result.Thu); diff != "" {
		t.Errorf("ParseOpenHours() Thu = %v, want %v", result.Thu, expected.Thu)
	}

	if diff := cmp.Diff(expected.Fri, result.Fri); diff != "" {
		t.Errorf("ParseOpenHours() Fri = %v, want %v", result.Fri, expected.Fri)
	}

	if diff := cmp.Diff(expected.Sat, result.Sat); diff != "" {
		t.Errorf("ParseOpenHours() Sat = %v, want %v", result.Sat, expected.Sat)
	}

	if diff := cmp.Diff(expected.Sun, result.Sun); diff != "" {
		t.Errorf("ParseOpenHours() Sun = %v, want %v", result.Sun, expected.Sun)
	}
}
