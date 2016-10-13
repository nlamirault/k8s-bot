# k8s-bot

[![License Apache 2][badge-license]](LICENSE)

This IRC Bot could responde you:

* [x] *!k8s :namespaces* : available namespaces
* [x] *!k8s :services* : available services
* [x] *!k8s :pods* : available pods
* [x] *!k8s :nodes* : available nodes

Type */help* to show the Bot commands.


## Installation

You can download the binaries :

* Architecture i386 [ [linux](https://bintray.com/artifact/download/nlamirault/oss/k8s_bot-0.1.0_linux_386) / [darwin](https://bintray.com/artifact/download/nlamirault/oss/k8s_bot-0.1.0_darwin_386) / [freebsd](https://bintray.com/artifact/download/nlamirault/oss/k8s_bot-0.1.0_freebsd_386) / [netbsd](https://bintray.com/artifact/download/nlamirault/oss/k8s_bot-0.1.0_netbsd_386) / [openbsd](https://bintray.com/artifact/download/nlamirault/oss/k8s_bot-0.1.0_openbsd_386) / [windows](https://bintray.com/artifact/download/nlamirault/oss/k8s_bot-0.1.0_windows_386.exe) ]
* Architecture amd64 [ [linux](https://bintray.com/artifact/download/nlamirault/oss/k8s_bot-0.1.0_linux_amd64) / [darwin](https://bintray.com/artifact/download/nlamirault/oss/k8s_bot-0.1.0_darwin_amd64) / [freebsd](https://bintray.com/artifact/download/nlamirault/oss/k8s_bot-0.1.0_freebsd_amd64) / [netbsd](https://bintray.com/artifact/download/nlamirault/oss/k8s_bot-0.1.0_netbsd_amd64) / [openbsd](https://bintray.com/artifact/download/nlamirault/oss/k8s_bot-0.1.0_openbsd_amd64) / [windows](https://bintray.com/artifact/download/nlamirault/oss/k8s_bot-0.1.0_windows_amd64.exe) ]
* Architecture arm [ [linux](https://bintray.com/artifact/download/nlamirault/oss/k8s_bot-0.1.0_linux_arm) / [freebsd](https://bintray.com/artifact/download/nlamirault/oss/k8s_bot-0.1.0_freebsd_arm) / [netbsd](https://bintray.com/artifact/download/nlamirault/oss/k8s_bot-0.1.0_netbsd_arm) ]


## Usage

Launch the bot :

    ./k8s-bot -kubeconfig ~/.kube/config -irc-server irc.freenode.net -irc-nick k8s-bot -irc-channels "#k8s-bot"


## Contributing

See [CONTRIBUTING](CONTRIBUTING.md).


## License

See [LICENSE](LICENSE) for the complete license.


## Changelog

A [changelog](ChangeLog.md) is available


## Contact

Nicolas Lamirault <nicolas.lamirault@gmail.com>

[badge-license]: https://img.shields.io/badge/license-Apache2-green.svg?style=flat
