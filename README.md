# gRPC demo

This repository is a simple demo of gRPC and some related functions. I plan to add more eventually.

## Table of Contents

- [Requirements](#requirements)
- [Tests](#tests)

## Requirements

* [GNU Make] `>= 3.81`
* [Go] `>= 1.16.0`

## Tests

You can run all tests with:

```bash
go test ./...
```

You can run a single Article test with:

```bash
cd internal/article
go test -run Test_ArticleFromProto -v
```

[GNU Make]: https://www.gnu.org/software/make/
[Go]: https://golang.org/
