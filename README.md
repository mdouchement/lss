# Light Storage Service

LSS is a simple web storage API.

## Requirements

- Golang 1.7.x

## Installation

```bash
$ go get github.com/Masterminds/glide
$ glide install
```

## Usage

### API Reference

#### Upload a file (`multipart/form-data`)

```
POST /my_directory/file.txt
```

Attributes:

name | value
:--- | :---
file | binary file


#### List files

```
GET /my_directory?list
```

```json
{
  "/my_directory": {
    "directory": true,
    "size": 102,
    "updated_at": "2017-01-16T09:55:36+01:00"
  },
  "/my_directory/file.txt": {
    "directory": false,
    "size": 764,
    "updated_at": "2017-01-16T09:55:36+01:00"
  }
}
```


#### Check if a file or directory exists

```
HEAD /my_directory/file.txt
```

`200 OK` or `404 Not Found`


#### Metadata of a file or a directory

```
GET /my_directory/file.txt?metadata
```

```json
{
  "directory": false,
  "size": 585,
  "updated_at": "2017-01-16T10:08:06+01:00"
}
```


#### Download a file

```
GET /my_directory/file.txt
```

```
Lorem ipsum dolor sit amet
```


### Configuration

Environment variables:
- `LSS_WORKSPACE` (default: `./workspace`)
- `LSS_UPLOAD_SIZE_LIMIT` (default: `8192` MiB)
- `LSS_ROUTER_NAMESPACE` (default: `""` ; e.g. `/lss`)

### Development

```bash
$ go run lss.go -h

$ go run lss.go server -p 4005

# Before pushing to Github
$ find . -name '*.go' -not -path './vendor*' -exec go fmt {} \;
```

### Production

```bash
$ go build -o lss lss.go
$ ./lss -p 4005 -w /data
```

## License

**MIT**

## Contributing

1. Fork it
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
5. Push to the branch (git push origin my-new-feature)
6. Create new Pull Request
