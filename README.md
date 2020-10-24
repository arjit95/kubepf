[![Build Status][build-shield]][build-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]

<p align="center">
  <h3 align="center">kubepf</h3>
  <p align="center">A utility to manage kubernetes port forwarding efficiently
  <br>
    <a href="https://github.com/arjit95/kubepf/issues">Report Bug</a>
  ·
  <a href="https://github.com/arjit95/kubepf/issues">Request Feature</a>
  </p>
</p>

## Table of Contents
* [About the Project](#about-the-project)
* [Install](#install)
  * [Pre-compiled binary](#pre-compiled-binary)
  * [Compile from source](#compile-from-source)
* [Usage](#usage)
* [Shortcuts](#shortcuts)
* [License](#license)
* [Acknowledgements](#acknowledgements)

## About The Project
![kubepf screenshot](_images/screenshot.svg)

Port forwarding in kubernetes using kubectl can be a bit difficult if the user wants to port forward multiple resources, where he needs to switch between multiple terminals or write scripts which allows you to run multiple resources on a single terminal. This could be made easier using an interactive prompt, where the user can run port-forwarding for multiple resources from the same terminal easily.

## Install

### Pre-compiled binary
Get the latest release [here](https://github.com/arjit95/kubepf/releases).

### Compile from source
__Clone__:
```bash
$ git clone github.com/arjit95/kubepf
$ cd kubepf
```
__Get the dependencies__:
```bash
$ go get -t -v ./...
```

__Build:__
```bash
$ go build -ldflags="-s -w" -o kubepf ./cmd/kubepf
```

## Usage

__Interactive__:
```bash
$ kubepf -i
```

__CLI__:
```bash
$ kubepf
kubepf handles all the different resources in a single
session, so you don't have to switch terminals or write
bash scripts to port-forward multiple resources.

Usage:
  kubepf [command]

Available Commands:
  completion  Generate completion script
  help        Help about any command
  start       Starts port forwarding on resource
  stop        Stops port forwarding on resource

Flags:
      --as string                      Username to impersonate for the operation
      --as-group stringArray           Group to impersonate for the operation, this flag can be repeated to specify multiple groups.
      --cache-dir string               Default cache directory (default "/home/arjit/.kube/cache")
      --certificate-authority string   Path to a cert file for the certificate authority
      --client-certificate string      Path to a client certificate file for TLS
      --client-key string              Path to a client key file for TLS
      --cluster string                 The name of the kubeconfig cluster to use
      --context string                 The name of the kubeconfig context to use
  -h, --help                           help for kubepf
      --insecure-skip-tls-verify       If true, the server's certificate will not be checked for validity. This will make your HTTPS connections insecure
  -i, --interactive                    Start kubepf in interactive mode
      --kubeconfig string              Path to the kubeconfig file to use for CLI requests.
      --match-server-version           Require server version to match client version
  -n, --namespace string               If present, the namespace scope for this CLI request
      --no-cache                       Do not use cached state
      --password string                Password for basic authentication to the API server
      --request-timeout string         The length of time to wait before giving up on a single server request. Non-zero values should contain a corresponding time unit (e.g. 1s, 2m, 3h). A value of zero means don't timeout requests. (default "0")
  -s, --server string                  The address and port of the Kubernetes API server
      --tls-server-name string         Server name to use for server certificate validation. If it is not provided, the hostname used to contact the server is used
      --token string                   Bearer token for authentication to the API server
      --user string                    The name of the kubeconfig user to use
      --username string                Username for basic authentication to the API server

Use "kubepf [command] --help" for more information about a command.
```

__Completions__:
Completions can be generated using the below command
```bash
$ kubepf completion [bash|zsh|fish|powershell]
```

## Shortcuts:
List of shortcuts could be found [here](https://github.com/arjit95/cobi#shortcuts) 

## License:
Distributed under the MIT License. See `LICENSE` for more information.

## Acknowledgements
- [cobra](https://github.com/spf13/cobra)
- [tview](https://github.com/rivo/tview)
- [promptui](https://github.com/manifoldco/promptui)
- [kubernetes](https://github.com/kubernetes/kubernetes)


[build-shield]: https://travis-ci.com/arjit95/kubepf.svg?branch=main
[build-url]: https://travis-ci.com/arjit95/kubepf
[issues-shield]: https://img.shields.io/github/issues/arjit95/kubepf.svg
[issues-url]: https://github.com/arjit95/kubepf/issues
[license-shield]: https://img.shields.io/github/license/arjit95/kubepf.svg
[license-url]: https://github.com/arjit95/kubepf/blob/main/LICENSE