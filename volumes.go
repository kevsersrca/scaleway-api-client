package scalewayapi

import (
	"encoding/json"
	"io/ioutil"
)

func (scw *ScalewayAPI) ListVolumes()(volumes VolumeList, err error) {
	resp, err := scw.GetResponse(ScalewayApiUrl, "volumes", nil)
	if err != nil {
		return volumes, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return volumes, err
	}
	if err = json.Unmarshal(body, &volumes); err != nil {
		return volumes, err
	}
	return volumes, nil
}