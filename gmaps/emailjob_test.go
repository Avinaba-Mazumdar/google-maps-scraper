package gmaps

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/require"
)

func Test_DocSocialLinksExtractor(t *testing.T) {
	html := `
	<!DOCTYPE html>
	<html>
	<head><title>Test Page</title></head>
	<body>
		<a href="https://www.facebook.com/businesspage">Facebook</a>
		<a href="https://instagram.com/businesspage/">Instagram</a>
		<a href="https://twitter.com/business">Twitter</a>
		<a href="https://x.com/business">X</a>
		<a href="https://www.linkedin.com/company/business">LinkedIn</a>
		<a href="https://youtube.com/@business">YouTube</a>
		<a href="https://tiktok.com/@business">TikTok</a>
		<a href="https://pinterest.com/business">Pinterest</a>
		<a href="https://fb.com/shortfb">FB Short</a>
		<!-- duplicate links -->
		<a href="https://www.facebook.com/businesspage">Facebook Dup</a>
		<a href="https://instagram.com/businesspage/">Instagram Dup</a>
		<!-- non-social links -->
		<a href="https://example.com/about">About</a>
		<a href="mailto:info@example.com">Email</a>
		<a href="tel:+1234567890">Phone</a>
		<a href="#section">Anchor</a>
		<a href="javascript:void(0)">JS</a>
		<a href="">Empty</a>
	</body>
	</html>
	`

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	require.NoError(t, err)

	links := docSocialLinksExtractor(doc)

	expected := []string{
		"https://www.facebook.com/businesspage",
		"https://instagram.com/businesspage/",
		"https://twitter.com/business",
		"https://x.com/business",
		"https://www.linkedin.com/company/business",
		"https://youtube.com/@business",
		"https://tiktok.com/@business",
		"https://pinterest.com/business",
		"https://fb.com/shortfb",
	}

	require.Equal(t, expected, links)
}
