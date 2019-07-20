<p align="center">
    <img src="https://github.com/opsway/documents/raw/master/docs/logo.png" alt="Documents logo" />
</p>
<h3 align="center">Documents</h3>
<p align="center">Service stateless API for converting HTML documents to PDF.</p>
<p align="center">
  <a href="/.github/contributing.md">Contributing</a>
</p>

---
# 

## Quick start

Open a terminal and run the following command:
 
    docker run --rm -p 8515:8515 -v template:/app/templates quay.io/opsway/documents
 
The API is now available on your host at `http://localhost:8515`.

<p align="center">
    <img src="https://github.com/opsway/documents/raw/master/docs/swagger.png" alt="demo swagger docs" />
</p>

## Build

    make image-release
  
## Badges

[![Travis CI](https://travis-ci.org/opsway/documents.svg?branch=master)](https://travis-ci.org/opsway/documents)
[![codecov](https://codecov.io/gh/opsway/documents/branch/master/graph/badge.svg)](https://codecov.io/gh/opsway/documents)
[![Go Report Card](https://goreportcard.com/badge/github.com/opsway/documents)](https://goreportcard.com/report/opsway/documents)
