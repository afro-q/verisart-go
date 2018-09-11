package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	coreTypes "github.com/quinlanmorake/verisart-go/types/core"

	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"
	errorMessages "github.com/quinlanmorake/verisart-go/types/core/errorMessages"
)

type Config struct {
	Cors     cors             `json:"cors"`
	Database coreTypes.String `json:"database"`
	Jwt      jwt              `json:"jwt"`
	Logging  logging          `json:"logging"`
	Server   server           `json:"server"`
}

const CONFIG_FILE_NAME = "app.config"

var AppConfig Config

func LoadConfig() coreTypes.Result {
	if pathToConfigFile, getPathToConfigFileResult := getConfigFilePath(); getPathToConfigFileResult.IsNotOk() {
		return getPathToConfigFileResult
	} else {
		return loadConfigFromFile(pathToConfigFile)
	}
}

/*
 NOTE: Perhaps in the future one would like to load config from somewhere else, we as such
 have a method `loadConfigFromFile`

 In my experience, depending on the project, one may want to manage this though an api,
 allowing the modification of config at runtime, for which there are various methods to do,
 bearing in mind which elements cannot be modified / designing for such and so on.
*/

func loadConfigFromFile(filePath coreTypes.String) coreTypes.Result {
	if fileData, readError := ioutil.ReadFile(filePath.ToString()); readError != nil {
		return coreTypes.Result{
			Code:    errorCodes.CONFIG_INIT_COULD_NOT_READ_FILE,
			Message: errorMessages.ErrorMessage(readError.Error()),
		}
	} else {
		if unmarshalError := json.Unmarshal(fileData, &AppConfig); unmarshalError != nil {
			return coreTypes.Result{
				Code:    errorCodes.CONFIG_INIT_COULD_NOT_UNMARSHAL_CONFIG_FILE,
				Message: errorMessages.ErrorMessage(unmarshalError.Error()),
			}
		} else {
			return coreTypes.NewSuccessResult()
		}
	}
}

/*
  Relying on Go library functions here, so not unit testing this function as it does nothing other than call Go library functions and map an error code
  If that changes, do add unit tests

  If one really feels the need, one could test that the right error code is returned on the right failure, but I feel that unwarranted at this stage.
*/

func getConfigFilePath() (filePath coreTypes.String, result coreTypes.Result) {
	currentExePath, getExeError := os.Executable()

	if getExeError != nil {
		result = coreTypes.Result{
			Code:    errorCodes.CONFIG_INIT_COULD_NOT_GET_CURRENT_EXE_PATH,
			Message: errorMessages.ErrorMessage(getExeError.Error()),
		}
		return
	}

	// Parse out symbolic links
	pathWithoutSymbolLinks, getPathError := filepath.EvalSymlinks(currentExePath)
	if getPathError != nil {
		result = coreTypes.Result{
			Code:    errorCodes.CONFIG_INIT_SYMBOL_LINK_ERROR,
			Message: errorMessages.ErrorMessage(getPathError.Error()),
		}
	}

	exeDirectoryPath := filepath.Dir(pathWithoutSymbolLinks)
	configFilePath := filepath.Join(exeDirectoryPath, CONFIG_FILE_NAME)

	if _, fileExistsError := os.Stat(configFilePath); fileExistsError != nil {
		result = coreTypes.Result{
			Code:    errorCodes.CONFIG_INIT_COULD_NOT_FIND_CONFIG_FILE,
			Message: errorMessages.ErrorMessage(fileExistsError.Error()),
		}
	}

	return coreTypes.String(configFilePath), coreTypes.NewSuccessResult()
}
