package cfi

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

func TestCFI(t *testing.T) {
	file, _ := os.Open("E:\\code\\workspace\\project\\gin-scaffold\\pkg\\crawler\\stock\\cfi\\cfi.txt")
	defer file.Close()
	// 解析 HTML
	doc, err := goquery.NewDocumentFromReader(io.Reader(file))
	if err != nil {
		log.Fatal(err)
	}

	price := ""
	change := ""
	doc.Find("#last").Each(func(i int, s *goquery.Selection) {
		price = strings.TrimSuffix(s.Text(), "↑") // 去除箭头
	})
	doc.Find("#chg").Each(func(i int, s *goquery.Selection) {
		change = s.Text()
	})
	fmt.Println(price)
	fmt.Println(change[0:4])
	fmt.Println(change[4:])
}
