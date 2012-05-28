Paging
======

A simple paging script written in Go, for use with template engines like mustache, or the builting template package of Go.

Usage
======
The Paging(current_page, page_count, visible int, url_str string)  function will return a []paging.Pelem.
The visible parameter controls how many
```
package main

import(
	"github.com/opesun/paging"
)

func main() {
	x, _ := Paging(1, 5, 2, "http://www.opesun.com?cat=embedded")
	paging.Print(x)
	fmt.Println("---")
	paging.PrintWithUrl(x)
}
```

The above example will output
```
[1] 2 3 ... 5
---
[1] http://www.opesun.com?cat=embedded&page=1
2	http://www.opesun.com?cat=embedded&page=2
3	http://www.opesun.com?cat=embedded&page=3
...
5	http://www.opesun.com?cat=embedded&page=5
```

Unfortunately the get parameters will be ordered randomly in the url since the net/url package url.Value type does not keep the order of the parameters, since it uses
a map internally.