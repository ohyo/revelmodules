package wikihelper

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/shurcooL/github_flavored_markdown"
)

func UrlEncode(str string) (encoded string) {
	encoded = strings.Replace(str, "/", "-", -1)
	encoded = url.QueryEscape(encoded)
	encoded = strings.Replace(encoded, "+", "%20", -1)
	return encoded
}

func Render(markdown string) string {
	// Replace bracket link
	body := markdown
	re := regexp.MustCompile("\\[\\[([^\\]\\[\\|]+)(\\|([^\\]\\[]+))?\\]\\]")
	for _, match := range re.FindAllStringSubmatch(body, -1) {
		bracketLink := match[0]
		title := match[1]
		alias := match[3]
		if alias == "" {
			alias = title
		}
		body = strings.Replace(body, bracketLink, "<a href=\"/page/"+UrlEncode(alias)+"\">"+title+"</a>", -1)
	}

	// Markdown Convert to
	html := string(github_flavored_markdown.Markdown([]byte(body)))
	html = strings.Replace(html, "<table>", "<table class=\"table table-bordered table-striped\">", -1)
	return html
}
