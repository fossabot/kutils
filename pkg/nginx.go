package pkg

import (
	"fmt"
	"strings"
)

func NewNginxWWWRewriteRule(host string) string {
	return fmt.Sprintf(
		"if ($host = 'www.%s') {\nrewrite ^ https://%s$request_uri permanent;\n}",
		host, host,
	)
}

func NewNginxWWWRewriteRules(hosts ...string) string {
	var rules []string
	for _, host := range hosts {
		rules = append(rules, NewNginxWWWRewriteRule(host))
	}
	return strings.Join(rules, "\n")
}
