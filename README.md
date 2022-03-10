# dots
dns over tls proxy

this project has been replaced by dnslite https://github.com/dilfish/dnslite


[![Build Status](https://travis-ci.org/dilfish/dots.svg?branch=master)](https://travis-ci.org/dilfish/dots)
[![codecov](https://codecov.io/gh/dilfish/dots/branch/master/graph/badge.svg)](https://codecov.io/gh/dilfish/dots)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg)](github.com/dilfish/dots)
[![GoDoc](https://godoc.org/github.com/dilfish/dots?status.svg)](https://godoc.org/github.com/dilfish/dots)

- a proxy between client and 1.1.1.1
- openssl req -new -nodes -x509 -out certs/client.pem -keyout certs/client.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=www.random.com/emailAddress=xxx@yyy" to generate client keys
