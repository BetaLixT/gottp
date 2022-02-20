package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type Location struct {
	PostCode            string  `json:"post code"`
	Country             string  `json:"country"`
	CountryAbbreviation string  `json:"country abbreviation"`
	Places              []Place `json:"places"`
}
type Place struct {
	PlaceName         string `json:"place name"`
	Longitude         string `json:"longitude"`
	State             string `json:"state"`
	StateAbbreviation string `json:"state abbreviation"`
	Latitude          string `json:"latitude"`
}

func TestGet(t *testing.T) {
	httpClient := NewClient()

	// - Testing 200
	resp, err := httpClient.Get(
		map[string]string{},
		"https://api.zippopotam.us/us/{}",
		"90210",
	)
	if err != nil {
		t.Error("Request failed")
		t.Error(err)
		t.FailNow()
	}

	if resp.StatusCode != 200 {
		t.Errorf("Status code unexpected: %d", resp.StatusCode)
	}
	loc := Location{}
	err = resp.Unmarshal(&loc)
	if err != nil {
		t.Error("unmarshaling failed")
		t.Error(err)
		t.FailNow()
	}
	expct := Location{
		PostCode:            "90210",
		Country:             "United States",
		CountryAbbreviation: "US",
		Places: []Place{
			{
				PlaceName:         "Beverly Hills",
				Longitude:         "-118.4065",
				Latitude:          "34.0901",
				State:             "California",
				StateAbbreviation: "CA",
			},
		},
	}
	if !cmp.Equal(expct, loc) {
		t.Error("Response wasn't expected")
		t.FailNow()
	}

	// - Testing 404
	resp, err = httpClient.Get(
		map[string]string{},
		"https://api.zippopotam.us/us/{}",
		"9021123",
	)

	if err != nil {
		t.Error("Request failed")
		t.Error(err)
		t.FailNow()
	}

	if resp.StatusCode != 404 {
		t.Errorf("Status code unexpected: %d", resp.StatusCode)
	}
}
