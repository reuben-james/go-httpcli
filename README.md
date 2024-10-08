# A simple HTTP Client CLI written in Go

## Cobra Initialisation

Project was initialised using cobra:
```
cobra-cli init --author "Reuben James reuben.james.uk@outlook.com" --license mit --viper
```

## Building & Running the CLI tool

### Build
```
go build -o httpcli
```

### Place on your PATH
```
mv httpcli /usr/loca/bin/
```

### Get help
```
httpcli --help
```

### Run a simple GET
```
httpcli get --url http://localhost/example 
```

## TODO

* Add post functionality
* Add config functionality