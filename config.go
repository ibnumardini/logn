package logn

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	LognIsActive bool `json:"logn_is_active"`
	Log          Log  `json:"log"`
	Tg           Tg   `json:"tg"`
	Zip          Zip  `json:"zip"`
}

type Log struct {
	LognDir          string `json:"logn_dir"`
	LognDefaultLoc   string `json:"logn_default_loc"`
	LognPrintConsole bool   `json:"logn_print_console"`
}

type Tg struct {
	LognAppName  string `json:"logn_app_name"`
	LognTgSend   bool   `json:"logn_tg_send"`
	LognTgToken  string `json:"logn_tg_token"`
	LognTgChatId string `json:"logn_tg_chat_id"`
}

type Zip struct {
	LognIsZipped  bool   `json:"logn_is_zipped"`
	LognDirZip    string `json:"logn_dir_zip"`
	LognDelOldDir bool   `json:"logn_del_old_dir"`
}

// config used to load config file
func config() (Config, error) {
	var config Config

	file, err := ioutil.ReadFile("logn_config.json")
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return config, err
	}

	if len(config.Tg.LognAppName) == 0 {
		config.Tg.LognAppName = "Logn-App LOG"
	}

	if len(config.Log.LognDir) == 0 {
		config.Log.LognDir = "log/"
	}

	if len(config.Log.LognDefaultLoc) == 0 {
		config.Log.LognDefaultLoc = "Asia/Jakarta"
	}

	if len(config.Zip.LognDirZip) == 0 {
		config.Zip.LognDirZip = "log_zip/"
	}

	if len(config.Tg.LognTgToken) == 0 {
		config.Tg.LognTgToken = "false"
	}

	if len(config.Tg.LognTgChatId) == 0 {
		config.Tg.LognTgToken = "false"
	}

	return config, nil
}
