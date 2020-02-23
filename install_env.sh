#!/bin/bash
set -o xtrace

# Install go packages which are not required to switch version
go get -u github.com/aead/cmac
go get -u github.com/bronze1man/radius
go get -u github.com/cydev/zero
go get -u github.com/ishidawataru/sctp
go get -u github.com/jinzhu/copier
go get -u github.com/mohae/deepcopy
go get -u github.com/xdg/scram
go get -u golang.org/x/crypto
go get -u golang.org/x/net
go get -u golang.org/x/oauth2
go get -u golang.org/x/sys
go get -u gopkg.in/check.v1
go get -u gopkg.in/yaml.v3

