package resource

import (
	"fmt"
	"os"
	"sort"

	"github.com/joho/godotenv"
	log "github.com/pmpavl/lgen-log"
)

const dotEnvFilePath string = ".env"

func (r *Resource) loadDotEnv() error {
	var dotEnvRead map[string]string

	if _, err := os.Stat(dotEnvFilePath); err == nil || os.IsExist(err) {
		if err := godotenv.Load(dotEnvFilePath); err != nil {
			return fmt.Errorf("godotenv load: %s", err)
		}

		if dotEnvRead, err = godotenv.Read(dotEnvFilePath); err != nil {
			return fmt.Errorf("godotenv read: %s", err)
		}
	}

	dotEnvReadLog(dotEnvRead)

	return nil
}

// Print dotEnv read map with pretty format.
func dotEnvReadLog(dotEnvRead map[string]string) {
	if len(dotEnvRead) == 0 {
		log.Logger.Log().Msgf("no env in %s file",
			string(dotEnvFilePath),
		)

		return
	}

	dotEnvReadSlice := make([]string, 0, len(dotEnvRead))

	for env := range dotEnvRead {
		dotEnvReadSlice = append(dotEnvReadSlice, env)
	}

	sort.Strings(dotEnvReadSlice)

	dotEnvReadLog := fmt.Sprintf("env read from %s", dotEnvFilePath)
	for _, env := range dotEnvReadSlice {
		dotEnvReadLog += fmt.Sprintf("\n\t%s: %s", env, dotEnvRead[env])
	}

	log.Logger.Log().Msg(dotEnvReadLog)

	return
}
