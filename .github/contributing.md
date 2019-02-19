# Contributing

Thank you for considering contributing. You'll find below useful information about how to contribute to the project.

## Contributing code

### Install from sources

1. Install and run the latest version of Docker
2. Verify your Go version (>= 1.11)
3. Fork this repository
4. Clone it outside of your `GOPATH` (we're using Go modules)

### Working with git

`1` Create your feature branch:
 
```bash
git clone # or git fetch 
git checkout -b my-new-feature origin/master
```

`2` Commit your changes (we're using [Conventional Commits](https://www.conventionalcommits.org)):

```bash
git commit -am "type: description"
```

`3` Push to the branch:

```bash
git push origin my-new-feature
```

`4` See `Your recently pushed branches` on page [Pull requests](../../../pulls).

### Testing

Run all tests:

```bash
make test
# or
make test-in-docker
```

## Reporting bugs and feature request

Your issue or feature request may already be reported!
Please search on the [issue tracker](../../../issues) before creating one.

If you do not find any relevant issue or feature request, feel free to
add a new one!

## Additional resources

* [Get started issue](../../../issues/new/choose)
* [Code of conduct](code_of_conduct.md)
