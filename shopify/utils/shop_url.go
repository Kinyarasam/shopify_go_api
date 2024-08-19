package utils

import (
	"net/url"
	"regexp"
	"strings"
)

const HostnamePattern = `[a-z0-9][a-z0-9-]*[a-z0-9]`

func SanitizeShopDomain(shopDomain, myshopifyDomain string) string {
	name := strings.ToLower(strings.TrimSpace(shopDomain))

	if !strings.Contains(name, myshopifyDomain) && !strings.Contains(name, ".") {
		name += "." + myshopifyDomain
	}

	name = strings.TrimPrefix(name, "http://")
	name = strings.TrimPrefix(name, "https://")

	uri, err := url.Parse("http://" + name)
	if err != nil {
		return ""
	}

	hostnamePattern := regexp.MustCompile(HostnamePattern + `\.` + regexp.QuoteMeta(myshopifyDomain))
	if hostnamePattern.MatchString(uri.Hostname()) {
		return uri.Hostname()
	}

	return ""
}
