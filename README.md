# SuperMarket
SuperMarket API to handle produce inventory
<p><a class="no-attachment-icon" href="https://travis-ci.org/littlebrownham/supermarket" target="_blank"><img src="https://travis-ci.org/littlebrownham/supermarket.svg?branch=master" alt=""></a></p>


## Building
To build binary, you can run `bin/build` and outputs binary to `build/supermarket`.

## Test
To run unit test, you can run `bin/test`. Add `--race` for race detection
build status: https://travis-ci.org/littlebrownham/supermarket/branches

## Docker
To build docker image, `docker run -t supermarket .`
docker repo: https://hub.docker.com/r/dnguy078/supermarket/

## Running
``` bash
# Runs against the binary
`./build/supermarket`

# Runs against docker image
`docker run -p 50200:50200 supermarket`
```


## SuperMarket API
#### [POST] /createproduce
Creates inventory, produce_code must be unique
+ Request
    + Body

            {
                "produce_code": "abcd-1234-1234-asbc",
                "name": "apples",
                "unit_price": 12.12
            }


+ Response 201 (application/json)

        {
          "produce_code": "abcd-1234-1234-asbc"
        }

+ Response 400 (application/json)
        - invalid produce_code, name, or price error message

#### [GET] /getproduce
Returns all produce inventory
+ Response 200 (application/json)

        [
            {
                "produce_code": "abcd-1234-1234-asbc",
                "name": "apples",
                "unit_price": 12.12
            },
            {
                "produce_code": "abcd-1234-1234-asbd",
                "name": "potatoes",
                "unit_price": 12.11
            }
        ]
#### [DELETE] /deleteproduce?produce_code=[produce_code]
Deletes a produce inventory
+ Response 200 (application/json)
+ Response 400 (application/json)
    - Does not exist
