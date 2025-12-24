package common

import (
	"bytes"

	"github.com/tyler-sommer/stick"
)

func Render(template string, params map[string]any) (string, error) {
	var tmpBuffer bytes.Buffer
	stickParam := make(map[string]stick.Value)
	for k, v := range params {
		stickParam[k] = v
	}

	errRender := stick.New(nil).Execute(template, &tmpBuffer, stickParam)
	if errRender != nil {
		return "", errRender
	}

	return tmpBuffer.String(), nil
}
