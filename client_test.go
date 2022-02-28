package main

import (
	"fmt"
	"testing"
)

type SampleResponse struct {
	Response string `json:"response"`
	Success  bool   `json:"success"`
}

var baseurl string = "https://betalixt-testapis.herokuapp.com"

func TestGet(t *testing.T) {
	httpClient := NewClient()

	// - Testing 200 simple
	resp, err := httpClient.Get(
		map[string]string{},
		fmt.Sprintf("%s/get", baseurl),
	)
	if err != nil {
		t.Error("Request failed")
		t.Error(err)
		t.FailNow()
	}

	if resp.StatusCode != 200 {
		t.Errorf("Status code unexpected: %d", resp.StatusCode)
	}
	res := SampleResponse{}
	err = resp.Unmarshal(&res)
	if err != nil {
		t.Error("unmarshaling failed")
		t.Error(err)
		t.FailNow()
	}
	if !res.Success || res.Response != "Successful No body" {
		t.Error("Response wasn't expected")
		t.FailNow()
	}

	// - Testing 200 one path param
	resp, err = httpClient.Get(
		map[string]string{},
		fmt.Sprintf("%s/get/{}", baseurl),
		"valid",
	)
	if err != nil {
		t.Error("Request failed")
		t.Error(err)
		t.FailNow()
	}

	if resp.StatusCode != 200 {
		t.Errorf("Status code unexpected: %d", resp.StatusCode)
	}
	res = SampleResponse{}
	err = resp.Unmarshal(&res)
	if err != nil {
		t.Error("unmarshaling failed")
		t.Error(err)
		t.FailNow()
	}
	if !res.Success || res.Response != "Successful one param" {
		t.Error("Response wasn't expected")
		t.FailNow()
	}

	// - Testing 404 one path param
	resp, err = httpClient.Get(
		map[string]string{},
		fmt.Sprintf("%s/get/{}", baseurl),
		"missing",
	)
	if err != nil {
		t.Error("Request failed")
		t.Error(err)
		t.FailNow()
	}

	if resp.StatusCode != 404 {
		t.Errorf("Status code unexpected: %d", resp.StatusCode)
	}
	res = SampleResponse{}
	err = resp.Unmarshal(&res)
	if err != nil {
		t.Error("unmarshaling failed")
		t.Error(err)
		t.FailNow()
	}
	if res.Success || res.Response != "Unsuccessful one param" {
		t.Error("Response wasn't expected")
		t.FailNow()
	}

	// - Testing 200 one two param
	resp, err = httpClient.Get(
		map[string]string{},
		fmt.Sprintf("%s/get/{}/var2/{}", baseurl),
		"valid",
		"valid",
	)
	if err != nil {
		t.Error("Request failed")
		t.Error(err)
		t.FailNow()
	}

	if resp.StatusCode != 200 {
		t.Errorf("Status code unexpected: %d", resp.StatusCode)
	}
	res = SampleResponse{}
	err = resp.Unmarshal(&res)
	if err != nil {
		t.Error("unmarshaling failed")
		t.Error(err)
		t.FailNow()
	}
	if !res.Success || res.Response != "Successful two param" {
		t.Error("Response wasn't expected")
		t.FailNow()
	}

	// - Testing 404 two path param
	resp, err = httpClient.Get(
		map[string]string{},
		fmt.Sprintf("%s/get/{}/var2/{}", baseurl),
		"missing",
		"valid",
	)
	if err != nil {
		t.Error("Request failed")
		t.Error(err)
		t.FailNow()
	}

	if resp.StatusCode != 404 {
		t.Errorf("Status code unexpected: %d", resp.StatusCode)
	}
	res = SampleResponse{}
	err = resp.Unmarshal(&res)
	if err != nil {
		t.Error("unmarshaling failed")
		t.Error(err)
		t.FailNow()
	}
	if res.Success || res.Response != "Unsuccessful two param" {
		t.Error("Response wasn't expected")
		t.FailNow()
	}

	// - Testing 200 one two param and closing
	resp, err = httpClient.Get(
		map[string]string{},
		fmt.Sprintf("%s/get/{}/var2/{}/closing", baseurl),
		"valid",
		"valid",
	)
	if err != nil {
		t.Error("Request failed")
		t.Error(err)
		t.FailNow()
	}

	if resp.StatusCode != 200 {
		t.Errorf("Status code unexpected: %d", resp.StatusCode)
	}
	res = SampleResponse{}
	err = resp.Unmarshal(&res)
	if err != nil {
		t.Error("unmarshaling failed")
		t.Error(err)
		t.FailNow()
	}
	if !res.Success || res.Response != "Successful two param" {
		t.Error("Response wasn't expected")
		t.FailNow()
	}

	// - Testing 404 two path param and closing
	resp, err = httpClient.Get(
		map[string]string{},
		fmt.Sprintf("%s/get/{}/var2/{}/closing", baseurl),
		"valid",
		"missing",
	)
	if err != nil {
		t.Error("Request failed")
		t.Error(err)
		t.FailNow()
	}

	if resp.StatusCode != 404 {
		t.Errorf("Status code unexpected: %d", resp.StatusCode)
	}
	res = SampleResponse{}
	err = resp.Unmarshal(&res)
	if err != nil {
		t.Error("unmarshaling failed")
		t.Error(err)
		t.FailNow()
	}
	if res.Success || res.Response != "Unsuccessful two param" {
		t.Error("Response wasn't expected")
		t.FailNow()
	}
}
