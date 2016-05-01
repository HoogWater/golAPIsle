# golAPIsle

## Development

Install the gb build tool and update the vendored packages

    go get github.com/constabulary/gb/...
    gb vendor update --all

You can now build and run the progam 

    gb build all
    ./bin/golapisle

You can now run the test suite

    gb test [-v]
