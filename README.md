<p align="center">
    <img src="logo.png"></img>
</p>

<div align="center">
    
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/ibnumardini/logn)
![GitHub](https://img.shields.io/github/license/ibnumardini/logn)

</div>

<p align="center">
    Logn <em><strong>:/logen/</strong></em> is a simple log management library for go app ツ.
</p>

## Installation

Use the package manager [go get](https://golang.org/cmd/go/#hdr-Download_and_install_packages_and_dependencies) to install this package.

```bash
go get -u github.com/ibnumardini/logn
```

## Configuration

Set your env for configuration Logn.

```go
os.Setenv("logn_app_name", "Logn-App LOG") // your app name
os.Setenv("logn_dir", "log/") // set dir to save your log
os.Setenv("logn_default_loc", "Asia/Jakarta") // timezone in log
os.Setenv("tg_send", "true") // set true, if you need report log type warning & error to telegam
os.Setenv("tg_token", "1416xxxxx:AAF3VOBjt7rIeO4tUL_dHxG0qxxxxxxxxx") // tg bot token
os.Setenv("tg_chat_id", "-4121xxxx") //  tg grup / chat_id
os.Setenv("is_zipped", "true") // if you need to zip your log every month
os.Setenv("logn_dir_zip", "log_zip/") // set directory to save your log
```

Set this in main func.

```go
logn.CronZip() // this command used to run zip log with cronjob
```

## Usage

```go
logn.InfoLog("this is log info")
logn.WarningLog("this is log warning")
logn.ErrorLog("this is log error")
```

#### To run zip with cron

```script
./app logn_zip_run // add this args in first args

```

## Contributing

Thank you for considering contributing to the Logn!.

## License

Logn is under the [MIT License](LICENSE.md)
