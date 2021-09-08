package helpers

import (
	"fmt"
	"os"
	"strings"
)

const (
	path = "/tmp/ygg-countries.conf"
)

func isCountriesExist() bool {
	_, err := os.Stat(path)
	return err == nil
}

func LoadCountries() []string {
	if !isCountriesExist() {
		return []string{}
	}

	countries, err := os.ReadFile(path)
	if err != nil {
		return []string{}
	}

	return strings.Split(string(countries), "\n")
}

func UpdateCountries(countries []string) {
	if isCountriesExist() {
		err := os.Remove(path)
		if err != nil {
			fmt.Println(err)
		}
	}

	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	block := ""
	for i, c := range countries {
		if i == 0 {
			block += c
			continue
		}
		block += "\n" + c
	}

	fmt.Fprintln(f, block)
}
