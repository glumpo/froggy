# froggy TG App

## Basic setup
Build and run in docker:
```
docker build -t froggy .
docker run froggy
```

Use `make` to build localy:
```
make froggy
```

## Provide config
All configuration should be provided with toml config file.

### Docker
Default config path is `./etc/config.toml`
Path could not be configured for now.

### Server binary
Use `-c` to provide path:
```
./froggy -c /custom/path
```

