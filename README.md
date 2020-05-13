go-slugify
==============

[![Build Status](https://travis-ci.org/shopsmart/go-slugify.svg?branch=master)](https://travis-ci.org/shopsmart/go-slugify)
[![Coverage Status](https://coveralls.io/repos/shopsmart/go-slugify/badge.svg?branch=master)](https://coveralls.io/r/shopsmart/go-slugify?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/shopsmart/go-slugify)](https://goreportcard.com/report/github.com/shopsmart/go-slugify)
[![GoDoc](https://godoc.org/github.com/shopsmart/go-slugify?status.svg)](https://godoc.org/github.com/shopsmart/go-slugify)

Make Pretty Slug.


Installation
------------

```
go get -u github.com/shopsmart/go-slugify
```

Install CLI tool:

```
go get -u github.com/shopsmart/go-slugify/slugify
$ slugify "Brad's Deals & abc"
brads-deals-abc
```


Documentation
--------------

API documentation can be found here:
https://godoc.org/github.com/mozillazg/go-slugify


Usage
------

```go
package main

import (
	"fmt"
	"github.com/shopsmart/go-slugify"
)

func main() {
	s := "Brad's Deals & abc"
	fmt.Println(slugify.Slugify(s))
	// Output: brads-deals-abc
}
```
