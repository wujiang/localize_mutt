# Loccalize datetime in mutt's header

### Setup development environment
http://golang.org/doc/install

TLTR:

1. `wget http://golang.org/dl/go1.3.linux-amd64.tar.gz`
2. `sudo tar -C /usr/local -xzf go1.3.linux-amd64.tar.gz`
3. `export PATH=$PATH:/usr/local/go/bin`
4. `export GOPATH=$HOME/go`
5. `export PATH=$PATH:$GOPATH/bin`

### Compile

- go build

## Usage

In your mutt configuration file, set
[`display_filter`](http://www.mutt.org/doc/manual/manual-6.html#display_filter)
variable to the location of the excutable, i.e.:

```

set display_filter = "~/opt/bin/localize_mutt"

```
