# Graphql Deduplicator
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
