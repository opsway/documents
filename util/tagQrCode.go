package util

import "tagQrCode"

import (
	"fmt"
	"log"
	"strings"

	"github.com/opsway/documents/cmd/template"
	"golang.org/x/net/html"
)

type tags struct {
	tagVal string
}

func processQrCodeTags(content string, data template.Context) (string, error) {
	tagVals := parseTagVals(content)
	fmt.Println(tagVals)
}

func parseTagVals(text string) (data []string) {

	tkn := html.NewTokenizer(strings.NewReader(text))

	var vals []string
	var isQrCode bool
	replacer := strings.NewReplacer("{", "", "}", "", " ", "")

	for {

		tt := tkn.Next()

		switch {

		case tt == html.ErrorToken:
			return vals

		case tt == html.StartTagToken:

			t := tkn.Token()
			isQrCode = t.Data == "qrCode"

		case tt == html.TextToken:

			t := tkn.Token()

			if isQrCode {
				vals = append(vals, replacer.Replace(t.Data))
			}

			isLi = false
		}
	}
}
