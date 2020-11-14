package logn

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"
	"time"
)

// InfoLog used for log type info
func InfoLog(logMessage interface{}) error {
	function, fileName, line, _ := runtime.Caller(1)
	loc := fmt.Sprintf("file: *%s*  func: *%s* line: *%d*", filePath(fileName), runtime.FuncForPC(function).Name(), line)

	err := makeLog(0, "INFO", logMessage, loc)
	if err != nil {
		return err
	}
	return nil
}

// WarningLog used for log type warning
func WarningLog(logMessage interface{}) error {
	function, fileName, line, _ := runtime.Caller(1)
	loc := fmt.Sprintf("file: *%s*  func: *%s* line: *%d*", filePath(fileName), runtime.FuncForPC(function).Name(), line)

	err := makeLog(1, "WARNING", logMessage, loc)
	if err != nil {
		return err
	}
	return nil
}

// WarningLog used for log type error
func ErrorLog(logMessage interface{}) error {
	function, fileName, line, _ := runtime.Caller(1)
	loc := fmt.Sprintf("file: *%s*  func: *%s* line: *%d*", filePath(fileName), runtime.FuncForPC(function).Name(), line)

	err := makeLog(2, "ERROR", logMessage, loc)
	if err != nil {
		return err
	}
	return nil
}

func makeLog(typeLogs int, title string, logMessage interface{}, loc string) error {
	var (
		dir     string
		typeLog string
	)

	if len(os.Getenv("logn_dir")) != 0 {
		dir = os.Getenv("logn_dir")
	} else {
		dir = "log/"
	}

	switch typLog := typeLogs; typLog {
	case 0:
		typeLog = "info"
	case 1:
		typeLog = "warning"
	case 2:
		typeLog = "error"
	default:
		typeLog = "other"
	}

	year, err := timeNow("Y")
	if err != nil {
		return err
	}

	month, err := timeNow("M")
	if err != nil {
		return err
	}

	day, err := timeNow("D")
	if err != nil {
		return err
	}

	YMDHis, err := timeNow("YMDHis")
	if err != nil {
		return err
	}

	var dates = map[string]string{
		"year":  dir + year,
		"month": dir + year + "/" + month,
		"day":   dir + year + "/" + month + "/" + day,
	}

	for _, date := range dates {
		_, errYear := os.Stat(date)
		if os.IsNotExist(errYear) {
			errDir := os.MkdirAll(date, 0755)
			if errDir != nil {
				return err
			}
		}
	}

	logFile := dir + year + "/" + month + "/" + day + "/" + typeLog + ".log"

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	defer file.Close()

	log.SetOutput(file)

	log.Printf("%v %v %v - %s", YMDHis, title, logMessage, strings.Replace(loc, "*", "", -1))

	var appName string
	if len(os.Getenv("logn_app_name")) == 0 {
		appName = "Logn-App LOG"
	} else {
		appName = os.Getenv("logn_app_name")
	}

	logMessageTg := fmt.Sprintf("*"+appName+"* \n\n - *Timestamp:* %v \n - *ErrorType:* %v \n - *Message:* %v \n\n - *Scene:* %s", YMDHis, title, logMessage, loc)

	var isSendTg string
	if len(os.Getenv("tg_send")) == 0 {
		isSendTg = "false"
	} else {
		isSendTg = os.Getenv("tg_send")
	}

	if typeLog != "info" && isSendTg == "true" {
		sendTg(url.QueryEscape(logMessageTg))
	}
	return nil
}

func timeNow(args string) (string, error) {
	if len(args) == 0 {
		args = "YMD"
	}

	var def_loc string
	if len(os.Getenv("logn_default_loc")) != 0 {
		def_loc = os.Getenv("logn_default_loc")
	} else {
		def_loc = "Asia/Jakarta"
	}

	loc, err := time.LoadLocation(def_loc)
	if err != nil {
		return "0", err
	}
	timein := time.Now().In(loc).String()
	switch mode := args; mode {
	case "Y":
		return timein[:4], nil
	case "M":
		return timein[5:7], nil
	case "D":
		return timein[8:10], nil
	case "YMDHis":
		return timein[:19], nil
	default:
		return timein[:10], nil
	}
}

func sendTg(logMessage interface{}) error {
	var (
		token  string
		chatId string
	)

	if len(os.Getenv("tg_token")) != 0 {
		token = os.Getenv("tg_token")
	}

	if len(os.Getenv("tg_chat_id")) != 0 {
		chatId = os.Getenv("tg_chat_id")
	}

	logMessageStr := fmt.Sprintf("%v", logMessage)

	url := "https://api.telegram.org/bot" + token + "/sendMessage?chat_id=" + chatId + "&parse_mode=markdown&text=" + logMessageStr

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return nil
}

func filePath(original string) string {
	i := strings.LastIndex(original, "/")
	if i == -1 {
		return original
	} else {
		return original[i+1:]
	}
}
