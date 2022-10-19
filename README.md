# Go CLI

Just a little CLI application that gets an URL and returns the IP address.

## Usage

### SDK

If you want to run the application with an installed SDK, just go to the Golang Official website and install the SDK yourself.

After installation, clone this repository with Git or download the latest version.

You can run the application with the following command:

```bash
go run *.go
```

If you want to build the application for your OS and architecture, just run the following command:

```bash
go build -o go-cli
```

After building, run the application as an executable.

### Docker

If you want to run the application with Docker, just run the following command:

```bash
sudo docker build --tag go-cli . && sudo docker run go-cli
```

## Licence

This project is licenced under the MIT Licence, check [HERE](LICENSE)

## Author

[Gabriel Santos Cardoso](https://linkedin.com/in/eng-gabrielscardoso)
