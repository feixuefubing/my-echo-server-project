package util

import "encoding/hex"

func TransferToHexStr(str string) string {
	dst := hex.EncodeToString([]byte(str))
	return "0x" + dst;
}

func ContainsDuplicate(strArr []string) bool {
	m := make(map[string]struct{})
	for _, v := range strArr {
		_, ok := m[v]
		if ok {
			return true
		} else {
			m[v] = struct{}{}
		}
	}
	return false
}