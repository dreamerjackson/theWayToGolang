/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */


//当出现defer时，函数和参数表达式将会被执行。但是函数真正的调用将会被延迟，直到包含defer的函数完成。可能会有多个defer，但是他们的执行顺序是相反的。
//上面的例子可以修改为：




package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// Copied from gopl.io/ch5/outline2.
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

//!+
func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	// ...print doc's title element...
	//!-
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" &&
			n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	forEachNode(doc, visitNode, nil)
	//!+

	return nil
}

//!-

func main() {
	for _, arg := range os.Args[1:] {
		if err := title(arg); err != nil {
			fmt.Fprintf(os.Stderr, "title: %v\n", err)
		}
	}
}

/*
defer常常使用在资源的打开或关闭中。
```
package ioutil
func ReadFile(filename string) ([]byte, error) {
  f, err := os.Open(filename)
    if err != nil {
      return nil, err
    }
  defer f.Close()
  return ReadAll(f)
}
```

```
  var mu sync.Mutex
  var m = make(map[string]int)

func lookup(key string) int {
  mu.Lock()
  defer mu.Unlock()
  return m[key]
}
```

*/