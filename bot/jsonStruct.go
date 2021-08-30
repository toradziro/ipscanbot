package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type JSONfile struct {
	Ip string				`json:"ip"`
	Hostname string			`json:"hostname"`
	Type string				`json:"type"`
	ContinentCode string	`json:"continent_code"`
	ContinentName string	`json:"continent_name"`
	CountryCode string		`json:"country_code"`
	CountryName string		`json:"country_name"`
	RegionCode string		`json:"region_code"`
	RegionName string		`json:"region_name"`
	City string				`json:"city"`
	Zip string				`json:"zip"`
	Latitude string			`json:"latitude"`
	Longitude string		`json:"longitude"`
	Location LocationJSON	`json:"location"`
}

type LocationJSON struct {
	GeonameID string		`json:"geoname_id"`
	Capital string			`json:"capital"`
	Languages LanguagesJSON	`json:"languages"`
	CountryFlag string		`json:"country_flag"`
	CountryFlagEmoji string	`json:"country_flag_emoji"`
	CallingCode string		`json:"calling_code"`
}

type LanguagesJSON struct {
	Code string				`json:"code"`
	Name string				`json:"name"`
	Native string			`json:"native"`
}

func getJSON (ip string) (string, error) {
	resp, err := http.Get("http://api.ipstack.com/" + ip + "?access_key=" + accessKey)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API service is not available now")
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", fmt.Errorf("error occured during reading info from the API site")
	}

	err = resp.Body.Close()
	if err != nil {
		return "", fmt.Errorf("error occured during closing resp body")
	}

	return string(data), nil
}
