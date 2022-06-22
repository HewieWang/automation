package main

import (
	"context"
	"log"
  "fmt"
  "flag"
	// "strings"
  "time"
	"os"
  "io/ioutil"
	"github.com/chromedp/chromedp"
  // "github.com/chromedp/cdproto/network"
	// "github.com/chromedp/cdproto/cdp"
)

func jscode(url string) string {
  file, err := os.Open(url)
  if err != nil {
     panic(err)
  }
  defer file.Close()
  content, err := ioutil.ReadAll(file)
  return string(content)
}
func main() {

  var domain      string
	var js					string

  flag.StringVar(&domain, "u", "", "域名")
	flag.StringVar(&js, "js", "", "自定义JS")

  flag.Parse()
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 100*time.Second)
	defer cancel()
	var html  string
	err := chromedp.Run(ctx,
		chromedp.Navigate(domain),
		chromedp.EvaluateAsDevTools(jscode(js), &html),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(html)
}
