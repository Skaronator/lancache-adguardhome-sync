# Lancache AdGuard Home Sync

## Introduction

Lancache AdGuard Home Sync is a synchronization tool designed to update DNS entries for specific services for the lancache server as rewrites to the AdGuardHome DNS server.

It serves users who already have a running local DNS server (AdGuard Home) in their LAN and wish to avoid running an additional lancache-dns container.
This project simplifies the integration of lancache server benefits without the necessity for additional DNS services.

## Installation and Configuration

### Requirements

- Docker installed on your system
- An existing [AdGuard Home](https://github.com/AdguardTeam/AdGuardHome) setup within your LAN
- An existing [LanCache](https://lancache.net) setup within your lan

### Setup

To start using Lancache AdGuard Home Sync, utilize the following Docker command:

```bash
docker run --name=lancache-adguardhome-sync -d \
  -e ADGUARD_USERNAME='admin' \
  -e ADGUARD_PASSWORD='password' \
  -e LANCACHE_SERVER='192.168.1.100' \
  -e ADGUARD_API='http://adguard.local:3000' \
  -e SYNC_INTERVAL_MINUTES="1440" \
  -e SERVICE_NAMES='steam,epicgames' \
  ghcr.io/skaronator/lancache-adguardhome-sync:latest
```

#### Run Modes

The application supports different run modes:

- **Daemon mode (default)**: Runs continuously with scheduled syncs
- **Once mode**: Run once and exit, useful for testing or manual runs

```bash
# Run once and exit
docker run --rm \
  -e ADGUARD_USERNAME='admin' \
  -e ADGUARD_PASSWORD='password' \
  -e LANCACHE_SERVER='192.168.1.100' \
  -e ADGUARD_API='http://adguard.local:3000' \
  -e SERVICE_NAMES='steam' \
  ghcr.io/skaronator/lancache-adguardhome-sync:latest --once

# Check version
docker run --rm ghcr.io/skaronator/lancache-adguardhome-sync:latest --version
```

#### Environment Variables

| Variable         | Description                        | Required  | Default         | Example                         |
|------------------|------------------------------------|-----------|-----------------|---------------------------------|
| ADGUARD_USERNAME | Username for AdGuard Home          | Yes       |                 | ADGUARD_USERNAME=admin          |
| ADGUARD_PASSWORD | Password for AdGuard Home          | Yes       |                 | ADGUARD_PASSWORD=admin          |
| LANCACHE_SERVER  | IP address of your lancache server | Yes       |                 | LANCACHE_SERVER=192.168.1.1     |
| ADGUARD_API      | API URL for AdGuard Home           | Yes       |                 | ADGUARD_API=http://fw.home:8080 |
| SYNC_INTERVAL_MINUTES | Sync interval in minutes        | No        | "1440"          | SYNC_INTERVAL_MINUTES="60" or "1440" |
| SERVICE_NAMES    | Services to sync DNS entries for   | Yes       |                 | SERVICE_NAMES='*' or SERVICE_NAMES='wsus,epicgames,steam' |

*Note: Use `SERVICE_NAMES='*'` to sync all services, or specify comma-separated service names.

### Building from Source

```bash
# Clone the repository
git clone https://github.com/Skaronator/lancache-adguardhome-sync.git
cd lancache-adguardhome-sync

# Build the Go binary
go build -o lancache-adguardhome-sync main.go

# Build the Docker image
docker build -t lancache-adguardhome-sync .
```

### How It Works

Upon configuring and initiating the container with the required environment variables, Lancache AdGuard Home Sync will automatically sync DNS entries for the designated services to your AdGuard Home installation. This process is governed by the `SYNC_INTERVAL_MINUTES` setting, allowing for periodic updates without manual intervention.

The application downloads domain files concurrently for better performance and uses minimal system resources. It runs as a native daemon with built-in minute-based scheduling.

## Development

### Running Tests

```bash
go test -v ./...
```

### Manual Execution

```bash
export ADGUARD_USERNAME="admin"
export ADGUARD_PASSWORD="password"
export LANCACHE_SERVER="192.168.1.100"
export ADGUARD_API="http://adguard.local:3000"
export SERVICE_NAMES="steam,epicgames"

go run main.go
```

## Contributing

We welcome contributions! For enhancements or fixes, please submit an issue or pull request on GitHub. Your contributions help improve Lancache AdGuard Home Sync for everyone.

## License

This project is available under the [MIT License](LICENSE). You are free to fork, modify, and use it in any way you see fit.
