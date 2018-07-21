# go-web-scraper
Map a website and learn some go

This takes a single url as input aka https://www.monozo.com. Will attempt to scrape all urls within the domain. It will
write the results out into a /tmp/results.json file. It should also write any errors encountered.

This also assumes that any non 200 response code is an error (this may not be technically true but its truthy)

The results are in the format

```json
[
    {
        "Hostname" : "https://monzo.com",
        "Paths" [
            //Child pages go here
        ],
        "Err" : null // Error goes here

    }
]
```


To test run go test.

To run use go run go-scrape-some-things.go  _url_to_visit_

ex: go run go-scrape-some-things.go http://www.monzo.com

OR

use go build and run the executable as ./go-scrape-some-things _url_to_visit_
### External Libraries

You will probably need these to run this.
* golang.org/x/net/html
