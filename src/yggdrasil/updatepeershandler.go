package yggdrasil

import "github.com/MohamedElmdary/yggdrasil-connector/src/helpers"

func UpdatePeersHandler(peers map[string][]string, initialCountries []string) func(bool, string) {
	selectedCountries := initialCountries

	return func(checked bool, country string) {
		if checked {
			selectedCountries = append(selectedCountries, country)
		} else {
			selectedCountries = helpers.RemoveCountry(selectedCountries, country)
		}

		helpers.UpdateCountries(selectedCountries)
		UpdatePeers(peers, selectedCountries)
	}
}
