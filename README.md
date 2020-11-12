<p align="center">
    <img src="logn-logo.png"></img>
</p>

<p align="center">
    Logn is a simple log management library for go app ツ.
</p>

## Installation

Use the package manager [go get](https://golang.org/cmd/go/#hdr-Download_and_install_packages_and_dependencies) to install this package.

```bash
go get -u github.com/ibnumardini/logn
```

## Configuration

You can put this config in init function.

```go
os.Setenv("logn_app_name", "Logn-App LOG") // your app name
os.Setenv("logn_dir", "log/") // set dir to save your log
os.Setenv("logn_default_loc", "Asia/Jakarta") // timezone in log
os.Setenv("tg_send", "true") // if you need report log type warning & error to telegam
os.Setenv("tg_token", "1416xxxxx:AAF3VOBjt7rIeO4tUL_dHxG0qxxxxxxxxx") // tg bot token
os.Setenv("tg_chat_id", "-4121xxxx") //  tg grup / chat_id
```

## Usage

```go
logn.InfoLog("this is log info")
logn.WarningLog("this is log warning")
logn.ErrorLog("this is log error")
```

## Contributing

Thank you for considering contributing to the Logn!.

## License

[MIT](https://choosealicense.com/licenses/mit/)
