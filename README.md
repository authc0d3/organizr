# Organizr

A simple cli application written in [Go](https://golang.org/) to automate file organization.
This project uses [equalfile](https://github.com/udhos/equalfile) module to compare files efficiently.

## ✅ Requirements

You only need [Go](https://golang.org/) instaled on your machine to compile the source code.

## 🚀 Get started

### 1. Clone this repo

```
> git clone https://github.com/authc0d3/organizr.git
```

### 2. Install dependencies

```
> cd organizr
> go get -u -v -f all
```

### 3. Build CLI tool

```
> go build
```

### 4. Ejecute in a directory to organize files

Only change "source-dir-path" by the path you want to organize. Flag -dest is optional, use it only if you want to move al organized files into another directory.

```
> organizr -src="source-dir-path" [-dest="output-dir-path"]
```

### 5. Customize configuration

You can configure output folders by creating a **config.json** in same folder where you place the binary. You have an example on **test.config.json**, you can customize and rename it as config.json to make it work.

```
{
  "extendsDefault": true,
  "output": {
    "defaultFolder": "Other Files",
    "folders": [
      { "ext": "svg,eps", "folder": "Vectors" },
      { "ext": "jpg,tiff", "folder": "Photos" }
    ]
  }
}
```

The available options are:

- **folders**: Array in which you can configure the output folders according to the file extension.

- **extendsDefault**: If true, your config will be extended with the default config. At the moment, this property only affects output folders. The "folders" prop will be extended with default folders that Organizr use (Documents, Images, Videos, Audios and Applications). If false, Organizer only will creates the ones you define in "folders" array.

- **defaultFolder**: Folder name in which all files that don't match the extensions for the specified folders will be included. If you don't indicate this, files with others no configured extensions will be stored in a folder with a name equal to its extension.

### 6. Advanced usage

```
> organizr -help
```

## ⏲️ Coming soon...

- Allow reverse operations
- Erase empty folders after organization

## :book: License

This project is under [MIT](https://opensource.org/licenses/MIT) license.

Copyright 2021 Antonio González [authc0d3](https://github.com/authc0d3)

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
