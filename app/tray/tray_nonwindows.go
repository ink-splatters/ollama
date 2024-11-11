//go:build !windows

package tray

import (
	"errors"

	"github.com/ink-splatters/ollama/app/tray/commontray"
)

func InitPlatformTray(icon, updateIcon []byte) (commontray.OllamaTray, error) {
	return nil, errors.New("not implemented")
}
