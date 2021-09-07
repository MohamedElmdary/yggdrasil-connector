package yggdrasil

import "github.com/MohamedElmdary/yggdrasil-connector/src/helpers"

func ConcatPeers(peers map[string][]string, countries []string) string {
	result := ""
	for country, ips := range peers {
		idx := helpers.FindIndex(countries, country)
		if idx == -1 {
			continue
		}

		for _, ip := range ips {
			result += "    " + ip + "\n"
		}
	}

	return "[\n" + result + "  ]"
	// result := ""
	// for _, peer := range peers {
	// 	result += "    " + peer + "\n"
	// }
	// return "[\n" + result + "  ]"
}
