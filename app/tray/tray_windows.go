package tray

import (
	"github.com/ink-splatters/ollama/app/tray/commontray"
	"github.com/ink-splatters/ollama/app/tray/wintray"
)

func InitPlatformTray(icon, updateIcon []byte) (commontray.OllamaTray, error) {
	return wintray.InitTray(icon, updateIcon)
}
