package scraper

import (
	"errors"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ExtractFromHTML(body string, selector string, attr string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		return "", err
	}

	sel := doc.Find(selector).First()
	if sel.Length() == 0 {
		return "", errors.New("selector not found: " + selector)
	}

	if attr != "" {
		val, exists := sel.Attr(attr)
		if !exists {
			return "", errors.New("attribute not found: " + attr)
		}
		return strings.TrimSpace(val), nil
	}

	return strings.TrimSpace(sel.Text()), nil
}
