# SimpleStore

SimpleStore is a demo JSON API webapp which stores and retrieves arbitrary data.

There are two endpoints:

### POST /messages/

  If successful, the reponse will be:

    {"digest":<sha sum of post body>}

### GET /messages/<sha sum>

If successful, the response will be the contents of the message.

## Running the demo

This demo uses the Google Compute Platform's DataStore product to store arbitrary data.
To run the demo I used Google AppEngine because it's extremely easy to deploy live webservices,
and the sandbox they give you is fantastic.

Dependency management is handled with [dep](https://github.com/golang/dep), though you could always use `go get` as well.


### Installation

#### Requirements

* A recent version of [Go](https://golang.org/dl/).
* The google cloud sdk: https://cloud.google.com/appengine/docs/standard/go/download

```
    #set up your GOPATH, or make a temporary one
    export GOPATH=$HOME/skaufmandemo
    mkdir -p $GOPATH/src/github.com/samuelkaufman
    mkdir -p $GOPATH/bin
    export PATH=$GOPATH/bin:$PATH
    cd $GOPATH/src/github.com/samuelkaufman
    git clone https://github.com/samuelkaufman/simplestore
    #You can skip this step if you're managing your own dependencies with GOPATH, or want
    #to use a different method to install dep
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
    cd simplestore
    dep ensure
    cd cmd/simplestore
    #this will use the default port of 8080, you can change the port with --port
    dev_appserver.py .
```

You can now the api at localhost:8080, or you can run the test with:

    go test -v  github.com/samuelkaufman/simplestore/cmd/simplestore -args -port 8080

Google AppEngine's SDK also comes with a nice mock DataStore browser which you can reach at http://localhost:8000/datastore.
