package scene

import (
	"encoding/json"
	"io"
)

func (scene *Scene) ToJson() ([]byte, error) {
	jsonBytes, err := json.Marshal(scene)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

func LoadStateFromJson(reader io.Reader) (*Scene, error) {
	sceneState := &Scene{}
	err := json.NewDecoder(reader).Decode(sceneState)
	if err != nil {
		return nil, err
	}
	return sceneState, nil
}
