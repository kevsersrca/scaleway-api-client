package scalewayapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

func (scw *ScalewayAPI) FilterIP(address string) (ip Ip , err error){
	var ipList IpList
	resp, err := scw.GetResponse(ScalewayApiUrl, "ips", nil)
	if err != nil {
		return ip, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ip, err
	}
	if err = json.Unmarshal(body, &ipList); err != nil {
		return ip, err
	}

	for _,v := range ipList.IP {
		if v.Address == address {
			return v,nil
		}
	}
	return ip, errors.New("IP Not Found!")
}

func (scw *ScalewayAPI) ListIP() (ipList IpList , err error){
	resp, err := scw.GetResponse(ScalewayApiUrl, "ips", nil)
	if err != nil {
		return ipList, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ipList, err
	}
	if err = json.Unmarshal(body, &ipList); err != nil {
		return ipList, err
	}
	return ipList, nil
}

func (scw *ScalewayAPI) ReseveIP() (response ReseveIPResponse , err error){
	data := ReseveIP{Organization:scw.Organization}
	resp, err := scw.PostResponse(ScalewayApiUrl, "ips", data)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	if err = json.Unmarshal(body, &response); err != nil {
		return response, err
	}
	return response, nil
}

func (scw *ScalewayAPI) DeleteIP(ip Ip) (err error){
	resource := fmt.Sprintf("ips/%s", ip.ID)
	resp, err := scw.DeleteResponse(ScalewayApiUrl, resource)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (scw *ScalewayAPI) AttachIP(ip Ip, server Server) (err error){
	data := map[string]interface{}{"server":server.ID}
	resource := fmt.Sprintf("ips/%s", ip.ID)
	resp, err := scw.PatchResponse(ScalewayApiUrl, resource, data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (scw *ScalewayAPI) DetachIP(ip Ip) (err error){
	data := map[string]interface{}{"server":nil}
	resource := fmt.Sprintf("ips/%s", ip.ID)
	resp, err := scw.PatchResponse(ScalewayApiUrl, resource, data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}