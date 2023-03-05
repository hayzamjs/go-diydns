# go-diyDNS

go-diyDNS is a CLI application that allows you to update your DNS records without having to use a third party service. 

## Installation

### From source
#### Clone the repository

```bash
git clone https://github.com/hayzamjs/go-diydns
```

#### Build the application
```bash
make build
```

#### Run the binary
The binary should be located in the `bin` directory. Make sure to have a config.json file in the same directory as the binary or specify a config file using the `--config` flag.

```bash
./bin/go-diydns --config /path/to/config.json
```

### From release

#### Download the binary

Download the binary from the [releases](https://github.com/hayzamjs/go-diydns) page.

#### Run the binary

```bash
./go-diydns --config /path/to/config.json
```

### Configuration

The configuration file is a JSON file that contains the following fields:

```js
[
    {
        "name": "cloudflare-hayzam-diydns", // a descriptive name for the record (for logs)
        "provider": "cloudflare", // provider (only cloudflare supported for now)
        "token": "<token>", // API token for the provider
        "domain": "diy1.hayzam.com", // domain to update
        "interval": 5 // interval in seconds to update the record
    },
    {
        "name": "cloudflare-hayzam-2",
        "provider": "cloudflare",
        "token": "<token>",
        "domain": "diy2.hayzam.com",
        "interval": 10
    }
]
```

### Todo 

- [ ] Add support for other providers
    - [x] Cloudflare
    - [ ] DigitalOcean
    - [ ] Godaddy
    - [ ] Google Cloud DNS
    - [ ] Route53

- [ ] Add Github Actions to build and release binaries


### Disclaimer

I wrote this application to learn Go and I am not a Go expert. so please use the project with caution. I am sure there are many things that can be improved. If you have any suggestions, please open an issue or a PR. 