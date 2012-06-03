// Package paging is a simple paging scripts to be used with template engines, like mustache, or the
// tpl package of Go.
// Example:
//
// x, _ := Paging(1, 5, 2, "http://www.opesun.com?cat=embedded")
// Print(x)
// PrintWithUrl(x)
package paging

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Pelem struct {
	IsDot     bool
	Url       string
	Page      int
	IsCurrent bool
	IsFirst   bool
	IsLast    bool
}

func handleNumbers(current, all, visible int) ([]Pelem, bool) {
	ret := []Pelem{}
	if current > all || current == 0 {
		return ret, false
	}
	if current-visible > 1 {
		ret = append(ret, Pelem{Page: 1, IsFirst: true})
		if current-visible > 2 {
			ret = append(ret, Pelem{IsDot: true})
		}
	}
	for i := current - visible; i < current; i++ {
		if i > 0 {
			ret = append(ret, Pelem{Page: i})
		}
	}
	ret = append(ret, Pelem{IsCurrent: true, Page: current})
	for i := current + 1; i <= all && (i-current) <= visible; i++ {
		ret = append(ret, Pelem{Page: i})
	}
	if current+visible < all {
		if current+visible < all-1 {
			ret = append(ret, Pelem{IsDot: true})
		}
		ret = append(ret, Pelem{Page: all})
	}
	return ret, true
}

func handleUrls(p []Pelem, url_str string) ([]Pelem, bool) {
	urlp := strings.Split(url_str, "?")
	getparams := urlp[len(urlp)-1]
	v, err := url.ParseQuery(getparams)
	if err == nil {
		for i, _ := range p {
			if p[i].IsDot {
				continue
			}
			v.Set("page", strconv.Itoa(p[i].Page))
			if len(urlp) > 1 {
				fmt.Println(v.Encode())
				p[i].Url = urlp[0] + "?" + v.Encode()
			} else {
				p[i].Url = v.Encode()
			}
		}
		return p, true
	}
	return p, false
}

func P(current, all, visible int, url_str string) ([]Pelem, bool) {
	ret, ok := handleNumbers(current, all, visible)
	if ok {
		return handleUrls(ret, url_str)
	}
	return ret, false
}

func Print(pages []Pelem) {
	for _, k := range pages {
		if k.IsDot {
			fmt.Print("... ")
		} else {
			if k.IsCurrent {
				fmt.Print("[", k.Page, "] ")
			} else {
				fmt.Print(k.Page, " ")
			}
		}
	}
}

func PrintWithUrl(pages []Pelem) {
	for _, k := range pages {
		if k.IsDot {
			fmt.Println("... ")
		} else {
			if k.IsCurrent {
				fmt.Println("[", k.Page, "] ", k.Url)
			} else {
				fmt.Println(k.Page, " ", k.Url)
			}
		}
	}
}
