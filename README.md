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

This project is under [MIT](https://opensource.org/licenses/MIT) license.

Copyright 2021 Antonio González (@authc0d3)

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
