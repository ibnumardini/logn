<p align="center">
    <img src="logo.png" width="150"></img>
</p>

<div align="center">

![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/ibnumardini/logn)
![GitHub](https://img.shields.io/github/license/ibnumardini/logn)

</div>

## About Logn
Logn <em><strong>:/logen/</strong></em> is a simple log management library for go app ツ, with several features provided such as :

- Generate log by day
- Generate log by type
- Report log type Warning / Error to telegram via bot
- Zipping log data every ont month
- Deleting old log directory

## Logn Tree
```
go-app
│   main
|
└───log
|    └───2020
|        └───12
|            └───01
|            |   info.log
|            |   warning.log
|            |   error.log
|            └───02
|            |   info.log
|            |   warning.log
|            |   error.log
|
└───log_zip
    └───2019
    |   log_12_2019.zip
    └───2020
    |   log_01_2020.zip
```

## Configuration
See the [wiki](https://github.com/ibnumardini/logn/wiki).

## Usage
See the [wiki](https://github.com/ibnumardini/logn/wiki).

## Contributing
Let's build this Logn library to make it even better, The contribution guide can be found in [here](/ibnumardini/logn/CONTRIBUTIONS.md).

## License
Logn is under the [MIT License](LICENSE.md)
