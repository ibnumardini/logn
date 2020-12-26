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
- Zipping log data every one month
- Deleting old log directory

## Logn Tree

```bash
go-app
│   main
│
├───log
│    └───2020
│        └───12
│            ├───01
│            │   info.log
│            │   warning.log
│            │   error.log
│            ├───02
│            │   info.log
│            │   warning.log
│            │   error.log
│
└───log_zip
    ├───2019
    │   log_12_2019.zip
    ├───2020
    │   log_01_2020.zip
```

## Installation

Use the package manager [go get](https://golang.org/cmd/go/#hdr-Download_and_install_packages_and_dependencies) to install this package.

```bash
go get -u github.com/ibnumardini/logn
```

## Configuration

- make config file [`logn_config.json`](https://github.com/ibnumardini/logn/blob/master/logn_config_sample.json) in your root project directory

```json
{
  "logn_is_active": true,
  "log": {
    "logn_dir": "log/",
    "logn_default_loc": "Asia/Jakarta",
    "logn_print_console": true
  },
  "tg": {
    "logn_tg_send": false,
    "logn_app_name": "Logn-App LOG",
    "logn_tg_token": "1416xxxxx:AAF3VOBjt7rIeO4tUL_dHxG0qxxxxxxxxx",
    "logn_tg_chat_id": "-4947xxxxx"
  },
  "zip": {
    "logn_is_zipped": false,
    "logn_dir_zip": "log_zip/",
    "logn_del_old_dir": false
  }
}
```

### Config description

- `logn_is_active` used to actived or deactived logn package

#### log

- `logn_dir` used for set default log directory name
- `logn_default_loc` used for set default timezone
- `logn_print_console` used to print log at console

#### telegram

- `logn_tg_send` used to actived or deactived report warning/error log to telegram
- `logn_app_name` used to set title in report telegram message
- `logn_tg_token` used for telegram bot token
- `logn_tg_chat_id` used for set where bot send report message

#### zip

- `logn_is_zipped` used to actived or deactived automate zipped log in one month
- `logn_dir_zip` used for set default log zip directory name
- `logn_del_old_dir` used to set automate deleting old zip directory after zipped

## Usage

```go
logn.InfoLog("hello world")
logn.WarningLog(fmt.Sprintf("%s\n", warningInfo))
logn.ErrorLog(fmt.Sprintln("error message!"))
```

### Usage description

- `logn.InfoLog()` used to make log type info
- `logn.WarningLog()` used to make log type warning
- `logn.ErrorLog()` used to make log type error

#### parameter

- if you have complex string, you can format first using `fmt.Sprintf()` or `fmt.Sprintln()` before input in the parameters

### Zipping log using cron

- set `logn.CronZip(1)`at your init func
- set like this `./app logn_zip_run` when run your app

#### parameter

- fill in the parameter with the position of the argument according to what you set, param type is int

## Contributing

Let's build this Logn library to make it even better.

## License

Logn is under the [MIT License](LICENSE.md)
