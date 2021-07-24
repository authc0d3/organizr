# Organizr

A simple cli application written in [Go](https://golang.org/) to automate file organization.

## ✅ Requirements

You only need [Go](https://golang.org/) instaled on your machine to compile the source code.

## 🚀 Get started

1. Clone this repo

```
> git clone https://github.com/authc0d3/organizr.git
```

2. Build CLI tool

```
> cd organizr
> go build organizr.go
```

3. Ejecute in a directory to organize files

Only change "source-dir-path" by the path you want to organize. Flag -dest is optional.

```
> organizr -src="source-dir-path" [-dest="output-dir-path"]
```

4. Get help

```
> organizr -h
```

5. Future features

- Allow reverse operations
- Can config output subfolder names by extension

## :book: License

[MIT](https://opensource.org/licenses/MIT)
