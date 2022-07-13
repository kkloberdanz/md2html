# md2html

md2html reads Markdown text from stdin and outputs HTML to stdout

## Build

```Go
go build .
```

## Usage

The following command will convert this README.md file into html and write it
to index.html.
```bash
cat README.md | md2html > index.html
```
