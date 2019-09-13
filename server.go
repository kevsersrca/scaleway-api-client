package scalewayapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/scaleway/scaleway-sdk-go/namegenerator"
	"io/ioutil"
	"net/http"
)

func New(token string, organization string, region string, userAgent string) (scw *ScalewayAPI, err error){
	if region != "fr-par-1" && region != "nl-ams-1" {
		return nil, errors.New("Region Not Found!")
	}
	if region == "fr-par-1" {
		ScalewayApiUrl = EndpointPar1
	}
	if region == "nl-ams-1" {
		ScalewayApiUrl = EndpointAms1
	}
	scw = &ScalewayAPI{
		BaseUrl:      ScalewayApiUrl,
		Organization: organization,
		Token:        token,
		client:       &http.Client{},
		userAgent:    userAgent,
		Region:       region,
	}
	return scw, nil
}
func (scw *ScalewayAPI) FilterServer(name string) (server ServerList, err error ) {
	//TODO data split
	resource := fmt.Sprintf("servers?name=%s", name)

	resp, err := scw.GetResponse(scw.BaseUrl, resource, nil)
	if err != nil {
		return server, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return server, err
	}
	if err = json.Unmarshal(body, &server); err != nil {
		return server, err
	}
	return server, nil
}

func (scw *ScalewayAPI) ListServer() (servers ServerList, err error){
	//TODO data split
	resource := "servers"

	resp, err := scw.GetResponse(scw.BaseUrl, resource, nil)
	if err != nil {
		return servers, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return servers, err
	}
	if err = json.Unmarshal(body, &servers); err != nil {
		return servers, err
	}
	return servers, nil
}

func (scw *ScalewayAPI) CreateServer(definition ServerCreateRequest) (server Server, err error){
	//TODO data split

	if definition.Organization == "" {
		definition.Organization = scw.Organization
	}

	if definition.Name == "" {
		namegenerator.GetRandomName("server")
	}

	resp, err := scw.PostResponse(scw.BaseUrl, "servers", definition)
	if err != nil {
		return server, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return server, err
	}
	if err = json.Unmarshal(body, &server); err != nil {
		return server, err
	}
	return server, nil
}

func(scw *ScalewayAPI) ListServerActions(server Server) (action string, err error) {
	resource := fmt.Sprintf("servers/%s/action", server.ID)
	resp, err := scw.GetResponse(scw.BaseUrl, resource, nil)
	if err != nil {
		return action, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return action, err
	}
	action = string(body)
	return action, nil
}

func(scw *ScalewayAPI) PowerOnServer(server Server) (task string, err error) {
	resource := fmt.Sprintf("servers/%s/action", server.ID)
	data := ServerAction{ Action:"poweron" }
	resp, err := scw.PostResponse(scw.BaseUrl, resource, data)
	if err != nil {
		return "",err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "",err
	}
	task = string(body)
	return task,nil
}
func(scw *ScalewayAPI) PowerOffServer(server Server) (task string, err error) {
	resource := fmt.Sprintf("servers/%s/action", server.ID)
	data := ServerAction{ Action:"poweron" }
	resp, err := scw.PostResponse(scw.BaseUrl, resource, data)
	if err != nil {
		return "",err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "",err
	}
	task = string(body)
	return task,nil
}


func (scw *ScalewayAPI) DeleteServer(server Server) (err error) {
	resource := fmt.Sprintf("servers/%s", server.ID)
	resp, err := scw.DeleteResponse(scw.BaseUrl, resource)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

