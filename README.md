# media-server
[![pipeline status](https://gitlab.com/kukoon/mediathek/media-server/badges/main/pipeline.svg)](https://gitlab.com/kukoon/mediathek/media-server/-/commits/main)
[![coverage report](https://gitlab.com/kukoon/mediathek/media-server/badges/main/coverage.svg)](https://gitlab.com/kukoon/mediathek/media-server/-/commits/main) 
[![Go Report Card](https://goreportcard.com/badge/github.com/Kukoon/media-server)](https://goreportcard.com/report/github.com/Kukoon/media-server)

This should become the Server of the OpenSource Mediathek and Streaming-Server.

For the Frontend, take a look here: [media-ui](https://github.com/Kukoon/media-ui)

## Compile
```sh
go get -v -u github.com/Kukoon/media-server
```

## Configuration
Copy [`config_example.toml`](./config_example.toml) to `config.toml` and take a look there.

## Pre-requierements
- PostgreSQL - Database (or cockroachdb)

For Testing we use a database called `media_server` with user `root`, take a look in [CI-Scripts](./.github/workflows/ci.yml#L26-L51)

## Startup
Everything is controlled by the `config.toml`.
A database migration happens on every startup if necessary.

### For Developers
```sh
go run ./main.go
```

### In Production
Create [systemd.service](https://www.freedesktop.org/software/systemd/man/systemd.service.html) file under: `/etc/systemd/system/media-server.service`
```ini
[Unit]
Description = media-server

[Service]
Type=simple
User=mediaserver
ExecStart=/usr/local/bin/media-server -c /etc/media-server.conf
Restart=always
RestartSec=5s
Environment=PATH=/usr/bin:/usr/local/bin

[Install]
WantedBy=multi-user.target
```

Store in `media-server.service` given path:
- config file
- binary

Enable service on Startup of service and start now:
```sh
systemctl enable --now media-server.service
```

