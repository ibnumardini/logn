# Logn

Logn is a ordinary log library management for golang -app.

## Installation

Use the package manager [go get](https://golang.org/cmd/go/#hdr-Download_and_install_packages_and_dependencies) to install this package.

```bash
go get -u github.com/ibnumardini/logn
```

## Configuration

You can put this config in init function

```go
os.Setenv("logn_app_name", "Logn-App LOG") // your app name
os.Setenv("logn_dir", "log/") // set dir to save your log
os.Setenv("logn_default_loc", "Asia/Jakarta") // timezone in log
os.Setenv("send_tg", "true") // is you need report log type warning & error to telegam
os.Setenv("tg_token", "1416xxxxx:AAF3VOBjt7rIeO4tUL_dHxG0qxxxxxxxxx") // tg bot token
os.Setenv("tg_chat_id", "-4121xxxx") //  tg grup / chat_id
```

## Usage

```go
logn.InfoLog("this if info")
logn.WarningLog("this is warning")
logn.ErrorLog("this is error")
```

## Contributing

I am very happy if you want :)

## License

[MIT](https://choosealicense.com/licenses/mit/)
