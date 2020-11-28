package logn

import (
	"archive/zip"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// InfoLog used for log type info
func InfoLog(logMessage interface{}) error {
	err := makeLog(0, "INFO", logMessage, "")
	if err != nil {
		return err
	}
	return nil
}

// WarningLog used for log type warning
func WarningLog(logMessage interface{}) error {
	function, fileName, line, _ := runtime.Caller(1)
	loc := fmt.Sprintf(" - file: *%s*  func: *%s* line: *%d*", filePath(fileName), runtime.FuncForPC(function).Name(), line)

	err := makeLog(1, "WARNING", logMessage, loc)
	if err != nil {
		return err
	}
	return nil
}

// WarningLog used for log type error
func ErrorLog(logMessage interface{}) error {
	function, fileName, line, _ := runtime.Caller(1)
	loc := fmt.Sprintf(" - file: *%s*  func: *%s* line: *%d*", filePath(fileName), runtime.FuncForPC(function).Name(), line)

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
	log.SetFlags(0)
	log.Printf("%v %v %v %s", YMDHis, title, logMessage, strings.Replace(loc, "*", "", -1))

	var appName string
	if len(os.Getenv("logn_app_name")) == 0 {
		appName = "Logn-App LOG"
	} else {
		appName = os.Getenv("logn_app_name")
	}

	logMessageTg := fmt.Sprintf("*"+appName+"* \n\n - *Timestamp:* %v \n - *Type:* %v \n - *Message:* %v \n\n - *Scene:* %s", YMDHis, title, logMessage, loc)

	var isSendTg string
	if len(os.Getenv("tg_send")) == 0 {
		isSendTg = "false"
	} else {
		isSendTg = os.Getenv("tg_send")
	}

	if typeLog != "info" && isSendTg == "true" {
		sendTg(url.QueryEscape(logMessageTg))
	}

	makeZip()

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

// CronZip used for cron zipped log
func CronZip() error {
	if len(os.Args) == 2 {
		if os.Args[1] == "logn_zip_run" {
			err := makeZip()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func makeZip() error {
	var dirZip, isZipped, dir string

	if len(os.Getenv("logn_dir")) != 0 {
		dir = os.Getenv("logn_dir")
	} else {
		dir = "log/"
	}

	if len(os.Getenv("logn_dir_zip")) != 0 {
		dirZip = os.Getenv("logn_dir_zip")
	} else {
		dirZip = "log_zip/"
	}

	if len(os.Getenv("is_zipped")) != 0 {
		isZipped = os.Getenv("is_zipped")
	} else {
		isZipped = "false"
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

	if day == "01" && isZipped == "true" {

		monthInt, err := strconv.Atoi(month)
		if err != nil {
			return err
		}

		prevMonth := monthInt - 1

		if month == "01" {
			prevMonth = 12

			yearInt, err := strconv.Atoi(year)
			if err != nil {
				return err
			}

			prevYear := yearInt - 1

			year = strconv.Itoa(prevYear)
		}

		dirFile := dirZip + year

		_, errYear := os.Stat(dirFile)
		if os.IsNotExist(errYear) {
			errDir := os.MkdirAll(dirFile, 0755)
			if errDir != nil {
				return errDir
			}
		}

		filename := fmt.Sprintf("log_%s_%s.zip", year, strconv.Itoa(prevMonth))

		baseDir := dir + year + "/" + strconv.Itoa(prevMonth) + "/"
		_, errYear = os.Stat(baseDir)
		if os.IsNotExist(errYear) {
			return errors.New("base dir doesn't exists")
		}

		outDir := dirFile + "/" + filename

		err = zipWriter(baseDir, outDir)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	return nil
}

func zipWriter(baseDir, outDir string) error {
	outFile, err := os.Create(outDir)
	if err != nil {
		return err
	}
	defer outFile.Close()

	w := zip.NewWriter(outFile)

	err = addFiles(w, baseDir, "")
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return nil
}

func addFiles(w *zip.Writer, basePath, baseInZip string) error {
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		return err
	}

	for _, file := range files {
		fmt.Println(basePath + file.Name())
		if !file.IsDir() {
			dat, err := ioutil.ReadFile(basePath + file.Name())
			if err != nil {
				return err
			}

			f, err := w.Create(baseInZip + file.Name())
			if err != nil {
				return err
			}
			_, err = f.Write(dat)
			if err != nil {
				return err
			}
		} else if file.IsDir() {
			newBase := basePath + file.Name() + "/"

			addFiles(w, newBase, baseInZip+file.Name()+"/")
		}
	}

	return nil
}
