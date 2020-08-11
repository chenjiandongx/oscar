# Oscar

> 🐶 Next generation building tool for nothing

[![Build Status](https://travis-ci.org/chenjiandongx/oscar.svg?branch=master)](https://travis-ci.org/chenjiandongx/oscar) [![Build status](https://ci.appveyor.com/api/projects/status/wdh74a2qce47j9tx/branch/master?svg=true)](https://ci.appveyor.com/project/chenjiandongx/oscar/branch/master) [![MIT License](https://img.shields.io/badge/License-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT) [![Contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/chenjiandongx/oscar)



## Motivation

Imitation is the sincerest form of flattery.

[Oscar](https://github.com/chenjiandongx/oscar) is yet another nonsense activity generator for Golang, which is inspired by the rust project [svenstaro/genact](https://github.com/svenstaro/genact). The aim of oscar is to ~~**pretend**~~ prove that you are busy working so that you could take a rest or drink a cup of cappuccino.


## Features

* blazing fast(nothing happend actually)
* variety modules
* easy to use and extend


## Installation

### Go

The classic way to install

```bash
$ go get -u github.com/chenjiandongx/oscar
```

### MacOS, Linux, Windows

The binary is published on the [release page](https://github.com/chenjiandongx/oscar/releases).


## Usage

```bash
$ oscar --help
Next generation building tool for nothing

Usage:
  oscar [command]

Available Commands:
  help        Help about any command
  list        List all available modules
  run         Run with the specific module

Flags:
  -h, --help      help for oscar
  -v, --version   version for oscar

Use "oscar [command] --help" for more information about a command.
```

### Modules list

```bash
$ oscar list
 - bootlog
 - botnet
 - cargo
 - cc
 - composer
 - cryptomining
 - docker
 - download
 - gomod
 - kernel_compile
 - memdump
 - weblog
```

### At a glance

<p align="center">
    <img src="https://user-images.githubusercontent.com/19553554/89953121-11652080-dc61-11ea-987c-9c2b89bf21ad.gif" alt="image" width=760 />
</p>

![image](https://user-images.githubusercontent.com/19553554/89951631-589de200-dc5e-11ea-869d-635022a5382d.png)
![image](https://user-images.githubusercontent.com/19553554/89951646-5fc4f000-dc5e-11ea-8229-b38bc2bb10f8.png)


## License

MIT [©chenjiandongx](https://github.com/chenjiandongx)
