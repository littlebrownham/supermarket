# SuperMarket
SuperMarket API to handle produce inventory
<p><a class="no-attachment-icon" href="" target="_blank"><img src="https://travis-ci.org/littlebrownham/supermarket.svg?branch=master" alt=""></a></p>


## Building
To build binary, you can run `bin/build` and outputs binary to `build/supermarket`.

## Test
To run unit test, you can run `bin/test`. Add `--race` for race detection

## Docker
To build docker image, `docker build -t supermarket .`

## Running
``` bash
# Runs against the binary
`./build/supermarket`

# Runs against docker image
`docker run --rm supermarket`
```

## API Call (Todo)
POST /createProduce
Body:
{
    "produce_code": [produce_code:string],
    "price": [price:float with 2 decimal precision],
    "name" [name:string name of produce]
}
Response:
200 - Entry added
400 - Duplicate entry

