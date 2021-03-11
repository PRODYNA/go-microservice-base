package probe

import "net/url"

func NewDnsProbe(url url.URL) func() (string, bool) {
	return func() (string, bool) {
		return "ok", true
	}
}
