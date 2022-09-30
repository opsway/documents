package util

import (
	//"fmt"
	"encoding/base64"
	"strings"

	"github.com/opsway/documents/cmd/template"
	qrcode "github.com/skip2/go-qrcode"
	"golang.org/x/net/html"
)

type tags struct {
	tagVal string
}

func ProcessQrCodeTags(content string, data template.Context) string {
	tagVals := parseTagVals(content)

	replacer := strings.NewReplacer("{", "", "}", "", " ", "")
	for _, element := range tagVals {
		dataProperty := replacer.Replace(element)
		qrCodeImg := ""

		if data[dataProperty] != nil && data[dataProperty] != "" {
			qrCodeImg = createQrCodeImage(data[dataProperty].(string))
		}

		content = strings.ReplaceAll(content, "<qrcode>"+element+"</qrcode>", qrCodeImg)
	}

	return content
}

func parseTagVals(text string) (data []string) {

	tkn := html.NewTokenizer(strings.NewReader(text))

	var vals []string
	var isQrCode bool

	for {
		tt := tkn.Next()

		switch {

		case tt == html.ErrorToken:
			return vals

		case tt == html.StartTagToken:

			t := tkn.Token()
			isQrCode = t.Data == "qrcode"

		case tt == html.TextToken:

			t := tkn.Token()

			if isQrCode {
				vals = append(vals, t.Data)
			}

			isQrCode = false
		}
	}
}

func createQrCodeImage(text string) string {
	png, err := qrcode.Encode(text, qrcode.Medium, 256)
	if err != nil {
		return ""
	}

	return "<img class=\"qrCode\" alt=\"QR Code\" src=\"data:image/png;base64," + base64.StdEncoding.EncodeToString(png) + "\"/>"
}
