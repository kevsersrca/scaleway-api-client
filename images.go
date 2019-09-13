package scalewayapi

import (
	"encoding/json"
	"io/ioutil"
)

func (scw *ScalewayAPI) ListImages()(images ImageList, err error) {
	resp, err := scw.GetResponse(ScalewayApiUrl, "images", nil)
	if err != nil {
		return images, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return images, err
	}

	if err = json.Unmarshal(body, &images); err != nil {
		return images, err
	}
	return images, nil
}