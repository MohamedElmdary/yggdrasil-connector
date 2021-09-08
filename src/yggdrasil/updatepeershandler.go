package yggdrasil

import "github.com/MohamedElmdary/yggdrasil-connector/src/helpers"

func UpdatePeersHandler(peers map[string][]string) func(bool, string) {
	selectedCountries := []string{}

	return func(checked bool, country string) {
		if checked {
			selectedCountries = append(selectedCountries, country)
		} else {
			selectedCountries = helpers.RemoveCountry(selectedCountries, country)
		}

		UpdatePeers(peers, selectedCountries)
	}
}
