package sys_base

import "strings"

func StringIsNullOrEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
