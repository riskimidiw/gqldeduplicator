# Graphql Deduplicator
[![godoc](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat)](https://pkg.go.dev/github.com/riskimidiw/graphql-deduplicator) 
[![go report](https://goreportcard.com/badge/github.com/riskimidiw/graphql-deduplicator)](https://goreportcard.com/report/github.com/riskimidiw/graphql-deduplicator)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/riskimidiw/graphql-deduplicator/blob/master/LICENSE)

GraphQL response deduplicator.

Javascript version: https://github.com/gajus/graphql-deduplicator 

### Usage

```
package main

import (
	"log"

	deduplicator "github.com/riskimidiw/graphql-deduplicator"
)

func main() {
    data := []byte(`
    {
        "root": [
            {
                "__typename": "foo",
                "id": 1,
                "name": "foo"
            },
            {
                "__typename": "foo",
                "id": 1,
                "name": "foo"
            }
        ]
    }`)

    deflate, err := deduplicator.Deflate(data)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("deflate:", string(deflate))

    inflate, err := deduplicator.Inflate(deflate)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("inflate:", string(inflate))
}
```
