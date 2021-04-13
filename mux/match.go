package mux

import "strings"

// MatchTemplate matches templates in rest url paths.
func MatchTemplate(pathElements, templateElements []string) (bool, map[string]string) {

	pathSize := len(pathElements)
	templateSize := len(templateElements)

	if pathSize == 0 && templateSize == 0 {
		return true, nil
	}

	var max int

	if pathSize >= templateSize {
		max = templateSize
	} else {
		max = pathSize
	}

	if pathSize > max || templateSize > max {
		return false, nil
	}

	for i := 0; i < max; i++ {
		t := templateElements[i]
		p := pathElements[i]

		if !strings.Contains(t, "{") {
			if t != p {
				return false, nil
			}
		}
	}

	return true, make(map[string]string)
}

// Match templates
func Match(path, template string) (bool, map[string]string) {
	return MatchTemplate(splitPath(path), splitPath(template))
}

func splitPath(s string) []string {
	return strings.Split(strings.Trim(s, "/"), "/")
}
