package main

import "strings"

func simplifyPath(path string) string {
	var parts []string
	for _, el := range strings.Split(path, "/") {
		switch el {
		case "", ".":
			continue
		case "..":
			if len(parts) > 0 {
				parts = parts[:len(parts)-1]
			}
		default:
			parts = append(parts, el)
		}
	}
	return "/" + strings.Join(parts, "/")
}
