package helpers

func RemoveCountry(countries []string, country string) []string {
	idx := FindIndex(countries, country)
	return append(countries[:idx], countries[idx+1:]...)
}
