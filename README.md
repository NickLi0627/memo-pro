# Memo Pro

## Development

### Environment

* Install the golang
* Download the binary file [here](https://go.dev/dl/) with the version: **1.17.5** and install it.
* Setup the environment path for golang's bin, e.g.

    ```bash
    echo 'export PATH="$PATH:/usr/local/go/bin"' >> ~/.zshrc
    source ~/.zshrc

    # Should be go1.17.5
    go version
    ```

### Build

Running the command: `go build` will build a executable file

### Test

Running the unit test: `go test -v`
