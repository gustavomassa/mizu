package mizu

import (
	"os"
	"path"
)

var (
	SemVer                                    = "0.0.1"
	Branch                                    = "develop"
	GitCommitHash                             = "" // this var is overridden using ldflags in makefile when building
	BuildTimestamp                            = "" // this var is overridden using ldflags in makefile when building
	RBACVersion                               = "v1"
	Platform                                  = ""
	DaemonModePersistentVolumeSizeBufferBytes = int64(500 * 1000 * 1000) //500mb
)

const DEVENVVAR = "MIZU_DISABLE_TELEMTRY"

func GetMizuFolderPath() string {
	home, homeDirErr := os.UserHomeDir()
	if homeDirErr != nil {
		return ""
	}
	return path.Join(home, ".mizu")
}
