package main

func fillMsg(f JSONfile) (string) {
	var msg string

	msg += "Ip: " + f.Ip + endl
	msg += "Hostname: " + f.Hostname + endl
	msg += "Type: " + f.Type + endl
	msg += "Continent code: " + f.ContinentCode + endl
	msg += "Continent name: " + f.ContinentName + endl
	msg += "Country code: " + f.CountryCode + endl
	msg += "Country name: " + f.CountryName + endl
	msg += "Region code: " + f.RegionCode + endl
	msg += "Region name: " + f.RegionName + endl
	msg += "City: " + f.City + endl
	msg += "Zip code: " + f.Zip + endl
	msg += "Latitude: " + f.Latitude + endl
	msg += "Longitude: " + f.Longitude + endl
	msg += "Geoname ID: " + f.Location.GeonameID + endl
	msg += "Capital: " + f.Location.Capital + endl
	msg += "Country flag: " + f.Location.CountryFlag + endl
	msg += "Flag image: " + f.Location.CountryFlagEmoji + endl
	msg += "Location calling code: " + f.Location.CallingCode + endl
	msg += "Language code: " + f.Location.Languages.Code + endl
	msg += "Language: " + f.Location.Languages.Name + endl
	msg += "Native language: " + f.Location.Languages.Native + endl

	return msg
}

func fillInfo() string {
	return "To use the bot you have to write valid ip address in ipv4 notation"
}