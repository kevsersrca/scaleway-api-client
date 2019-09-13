package scalewayapi

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var ServerName = "Test"

var TestIP = "51.158.109.119"

func Init() (api *ScalewayAPI){
	api, err := New("token", "organization", "region", "")
	if err != nil {
		return api
	}
	return
}

func TestScalewayAPI_NewApi(t *testing.T) {
	Convey("Testing New()", t, func() {
		api := Init()
		So(api, ShouldNotBeNil)
		So(api.Token, ShouldEqual, "token")
		So(api.Organization, ShouldEqual, "organization")
		So(api.client, ShouldNotBeNil)
	})
}

func TestScalewayAPI_ListServer(t *testing.T) {
	api := Init()
	srv, _:= api.ListServer()
	fmt.Println(srv.Servers)
	if srv.Servers != nil {
		t.Skip()
	}
}

func TestScalewayAPI_ImageList(t *testing.T) {
	api := Init()
	images, _ := api.ListImages()
	if images.Images != nil {
		t.Skip()
	}
}

func TestScalewayAPI_VolumeList(t *testing.T) {
	api := Init()
	volumes, _ := api.ListVolumes()
	if volumes.Volumes != nil {
		t.Skip()
	}
}

func TestScalewayAPI_CreateServer(t *testing.T) {
	Convey("Testing New()", t, func() {
		api := Init()
		createServerResponse, err := api.CreateServer(ServerCreateRequest{
			Name:              ServerName,
			Organization:      api.Organization,
			Image:             ClientImageID,
			Volumes:           map[string]*VolumeTemplate{},
			EnableIpv6:        false,
			CommercialType:    ClientPackageType,
			Tags:              []string{"client"},
			DynamicIPRequired: true,
			BootType: ServerBootTypeLocal,
		})
		So(err, ShouldBeNil)
		So(createServerResponse, ShouldNotBeNil)
	})
}

func TestScalewayAPI_DeleteServer(t *testing.T) {
	api := Init()
	server , err := api.FilterServer(ServerName)
	if err != nil {
		t.Error(err)
	}
	if len(server.Servers) == 0  {
		t.Error("Server Not Found")
	}
	_, err = api.PowerOffServer(server.Servers[0])
	if err != nil {
		t.Error(err)
	}
	err = api.DeleteServer(server.Servers[0])
	if err == nil {
		t.Skip()
	}
}

func TestScalewayAPI_PowerOnServer(t *testing.T) {
	api := Init()
	server , _ := api.FilterServer(ServerName)
	if len(server.Servers) == 0  {
		t.Error("server not found")
	}
	_, err := api.PowerOnServer(server.Servers[0])
	if err != nil {
		t.Fatal(err)
	}
	t.Skip()
}

func TestScalewayAPI_IPList(t *testing.T) {
	api := Init()
	ips, err := api.ListIP()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(len(ips.IP))
	if  ips.IP != nil {
		t.Skip()
	}
}

func TestScalewayAPI_IPReseve(t *testing.T) {
	api := Init()
	res, err := api.ReseveIP()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res)
}

func TestScalewayAPI_FilterIP(t *testing.T) {
	api := Init()
	_, err := api.FilterIP(TestIP)
	if err != nil {
		t.Error(err)
	}
}

func TestScalewayAPI_DeleteIP(t *testing.T) {
	api := Init()
	ip, err := api.FilterIP(TestIP)
	if err != nil {
		t.Error(err)
	}
	err = api.DeleteIP(ip)
	if err != nil {
		t.Error(err)
	}
	t.Skip()
}

func TestScalewayAPI_AttachIP(t *testing.T) {
	api := Init()
	ip, err := api.FilterIP(TestIP)
	if err != nil {
		t.Error(err)
	}
	server , _ := api.FilterServer(ServerName)
	if len(server.Servers) == 0  {
		t.Error("server not found")
	}
	err = api.AttachIP(ip, server.Servers[0])
	if err != nil {
		t.Error(err)
	}
	t.Skip()
}

func TestScalewayAPI_DetachIP(t *testing.T) {
	api := Init()
	ip, err := api.FilterIP(TestIP)
	if err != nil {
		t.Error(err)
	}
	err = api.DetachIP(ip)
	if err != nil {
		t.Error(err)
	}
	t.Skip()
}