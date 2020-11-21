[![Issues][issues-shield]][issues-url]

<!-- PROJECT LOGO -->
<br />
<p align="center">
  <!-- <a href="https://github.com/roerohan/wait-for-it">
    <img src="https://project-logo.png" alt="Logo" width="80">
  </a> -->

  <h3 align="center">wait-for-it</h3>

  <p align="center">
    A Golang package to wait on the availability of a TCP host and port. 
    <br />
    <a href="https://github.com/roerohan/wait-for-it"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/roerohan/wait-for-it">View Demo</a>
    ·
    <a href="https://github.com/roerohan/wait-for-it/issues">Report Bug</a>
    ·
    <a href="https://github.com/roerohan/wait-for-it/issues">Request Feature</a>
  </p>
</p>



<!-- TABLE OF CONTENTS -->
## Table of Contents

* [About the Project](#about-the-project)
  * [Built With](#built-with)
* [Getting Started](#getting-started)
  * [Prerequisites](#prerequisites)
  * [Installation](#installation)
* [Usage](#usage)
* [Roadmap](#roadmap)
* [Contributing](#contributing)
* [License](#license)
* [Contributors](#contributors-)



<!-- ABOUT THE PROJECT -->
## About The Project

<img src="./assets/wait-for-it.png" alt="wait-for-it" width="800">

This package is adapted from [vishnubob/wait-for-it](https://github.com/vishnubob/wait-for-it), a popular project used to wait for TCP connections until a service is up. This is commonly used in `docker-compose` files to make one service wait for another, for example, to make a web server wait for a `mysql` database.

Since [vishnubob/wait-for-it](https://github.com/vishnubob/wait-for-it) is a bash script, it does not work directly with minimal containers like [scratch](https://hub.docker.com/_/scratch), which are commonly used to run binaries.

With the help of this package, you can generate a binary, which can run inside minimal Docker containers and wait for a TCP connection such as a `mysql` database. You can find an example here: [csivitu/bl0b](https://github.com/csivitu/bl0b/blob/master/docker-compose.yml).


### Built With

* [Go](https://golang.org/)



<!-- GETTING STARTED -->
## Getting Started

A binary for `linux` is available in the [GitHub releases](https://github.com/roerohan/wait-for-it/releases/).

If you want to build a binary for a different Operating System / Architecture, you can follow the procedure below.

### Prerequisites

The only prerequisite is `golang` which you can get [here](https://golang.org/).

* go

### Installation
 
1. Get the package using `go get`.
```bash
go get github.com/roerohan/wait-for-it
```

Alternatively, you can follow these steps:

1. Clone the repository.
```bash
git clone https://github.com/roerohan/wait-for-it
```

2. Build a go binary from source.
```bash
cd wait-for-it
go build -o ./bin/wait-for-it
```

3. Use the binary inside the bin folder.
```bash
./bin/wait-for-it google.com:80 -- echo "It works!"
```


<!-- USAGE EXAMPLES -->
## Usage

The usage is similar to [vishnubob/wait-for-it](https://github.com/vishnubob/wait-for-it).

Use `wait-for-it -h` to display the following list.

```
Usage of wait-for-it:
  -q    Quiet, don't output any status messages
  -s    Only execute subcommand if the test succeeds
  -t int
        Timeout in seconds, zero for no timeout (default 15)
  -w host:port
        Services to be waiting for, in the form host:port
```

### Examples:

1. Waiting for multiple services in parallel.

```sh
wait-for-it -w google.com:80 -w localhost:27017 -t 30 -- echo "Waiting for 30 seconds for google.com:80 and localhost:27017"
```

2. Strict mode will not execute the subcommand only if TCP connection was successful.

```sh
$ wait-for-it . -w abcd:80 -s -t 5 -- echo "Done\!"
wait-for-it: waiting 5 seconds for abcd:80
wait-for-it: timeout occured after waiting for 5 seconds
wait-for-it: strict mode, refusing to execute subprocess
```


<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/roerohan/wait-for-it/issues) for a list of proposed features (and known issues).



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'feat: Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

You are requested to follow the contribution guidelines specified in [CONTRIBUTING.md](./CONTRIBUTING.md) while contributing to the project :smile:.

<!-- LICENSE -->
## License

Distributed under the MIT License. See [`LICENSE`](./LICENSE) for more information.




<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[roerohan-url]: https://roerohan.github.io
[issues-shield]: https://img.shields.io/github/issues/othneildrew/Best-README-Template.svg?style=flat-square
[issues-url]: https://github.com/roerohan/wait-for-it/issues
