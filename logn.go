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
	err := makeLog(0, "INFO :", logMessage, "")
	if err != nil {
		return err
	}

	return nil
}

// WarningLog used for log type warning
func WarningLog(logMessage interface{}) error {
	function, fileName, line, _ := runtime.Caller(1)
	loc := fmt.Sprintf("- file: *%s*  func: *%s* line: *%d*", filePath(fileName), runtime.FuncForPC(function).Name(), line)

	err := makeLog(1, "WARNING :", logMessage, loc)
	if err != nil {
		return err
	}

	return nil
}

// ErrorLog used for log type error
func ErrorLog(logMessage interface{}) error {
	function, fileName, line, _ := runtime.Caller(1)
	loc := fmt.Sprintf("- file: *%s*  func: *%s* line: *%d*", filePath(fileName), runtime.FuncForPC(function).Name(), line)

	err := makeLog(2, "ERROR :", logMessage, loc)
	if err != nil {
		return err
	}

	return nil
}

// CronZip used for cron zipped log
func CronZip() error {
	if len(os.Args) >= 2 {
		if os.Args[1] == "logn_zip_run" {
			err := makeZip()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func makeLog(logTypes int, title string, logMessage interface{}, loc string) error {
	config, err := config()
	if err != nil {
		return err
	}

	if !config.LognIsActive {
		return errors.New("logn is not active")
	}

	var logType string

	switch logTypes {
	case 0:
		logType = "info"
	case 1:
		logType = "warning"
	case 2:
		logType = "error"
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
		"year":  config.Log.LognDir + year,
		"month": config.Log.LognDir + year + "/" + month,
		"day":   config.Log.LognDir + year + "/" + month + "/" + day,
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

	logFile := config.Log.LognDir + year + "/" + month + "/" + day + "/" + logType + ".log"

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	defer file.Close()

	log.SetOutput(file)
	log.SetFlags(0)
	log.Printf("%v %v %v %s", YMDHis, title, logMessage, strings.Replace(loc, "*", "", -1))

	if config.Log.LognPrintConsole {
		fmt.Printf("%v %v %v %s\n", YMDHis, title, logMessage, strings.Replace(loc, "*", "", -1))
	}

	logMessageTg := fmt.Sprintf("*"+config.Tg.LognAppName+"* \n\n - *Timestamp:* %v \n - *Type:* %v \n - *Message:* %v \n\n - *Scene:* %s", YMDHis, title[:len(title)-1], logMessage, loc)

	if logType != "info" && config.Tg.LognTgSend {
		sendTg(url.QueryEscape(logMessageTg))
	}

	if config.Zip.LognIsZipped {
		makeZip()
	}

	return nil
}

func timeNow(args string) (string, error) {
	config, err := config()
	if err != nil {
		return "", err
	}

	if len(args) == 0 {
		args = "YMD"
	}

	loc, err := time.LoadLocation(config.Log.LognDefaultLoc)
	if err != nil {
		return "", err
	}

	timein := time.Now().In(loc).String()
	switch args {
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
	config, err := config()
	if err != nil {
		return err
	}

	logMessageStr := fmt.Sprintf("%v", logMessage)

	url := "https://api.telegram.org/bot" + config.Tg.LognTgToken + "/sendMessage?chat_id=" + config.Tg.LognTgChatId + "&parse_mode=markdown&text=" + logMessageStr

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

func makeZip() error {
	config, err := config()
	if err != nil {
		return err
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

	if day == "01" && config.Zip.LognIsZipped {
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

		dirFile := config.Zip.LognDirZip + year

		_, errYear := os.Stat(dirFile)
		if os.IsNotExist(errYear) {
			errDir := os.MkdirAll(dirFile, 0755)
			if errDir != nil {
				return errDir
			}
		}

		filename := fmt.Sprintf("log_%s_%s.zip", year, strconv.Itoa(prevMonth))

		baseDir := config.Log.LognDir + year + "/" + strconv.Itoa(prevMonth) + "/"
		_, errYear = os.Stat(baseDir)
		if os.IsNotExist(errYear) {
			return errors.New("base dir doesn't exists")
		}

		outDir := dirFile + "/" + filename

		err = zipWriter(baseDir, outDir)
		if err != nil {
			return err
		}

		// auto remove old logn directory after zipped
		if config.Zip.LognDelOldDir {
			if _, err := os.Stat(outDir); err == nil {
				delOldLognDir(baseDir)
			}
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

// delOldLognDir used to delete old logn directory
func delOldLognDir(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}

	return nil
}
