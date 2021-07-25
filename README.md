# Organizr

A simple cli application written in [Go](https://golang.org/) to automate file organization.
This project uses [equalfile](https://github.com/udhos/equalfile) module to compare files efficiently.

## ✅ Requirements

You only need [Go](https://golang.org/) instaled on your machine to compile the source code.

## 🚀 Get started

1. Clone this repo

```
> git clone https://github.com/authc0d3/organizr.git
```

2. Install dependencies

```
> cd organizr
> go get -u -v -f all
```

3. Build CLI tool

```
> go build
```

4. Ejecute in a directory to organize files

Only change "source-dir-path" by the path you want to organize. Flag -dest is optional.

```
> organizr -src="source-dir-path" [-dest="output-dir-path"]
```

5. Advanced usage

```
> organizr -help
```

## ⏲️ Coming soon...

- Allow reverse operations
- Can config output subfolder names by extension
- Erase empty folders after organization

## :book: License

[MIT](https://opensource.org/licenses/MIT)
