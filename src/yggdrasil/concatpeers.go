package yggdrasil

func ConcatPeers(peers []string) string {
	result := ""
	for _, peer := range peers {
		result += "    " + peer + "\n"
	}
	return "[" + result + "]"
}
