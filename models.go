package scalewayapi

import (
	"net/http"
)

const (
	perPage = 50

	EndpointPar1 = "https://api.scaleway.com/instance/v1/zones/fr-par-1"

	EndpointAms1 = "https://api.scaleway.com/instance/v1/zones/nl-ams-1"
)

type VolumeType string

const (
	VolumeTypeLSSD = VolumeType("l_ssd")
	VolumeTypeBSSD = VolumeType("b_ssd")
)

var (
	//Package Type
	ClientPackageType = "DEV1-XL"
	//serp-bulk-requests-client image id
	ClientImageID = "75bf569a-4028-481b-b4e4-0182001b280a"
	//Scaleway api url
	ScalewayApiUrl = "https://api.scaleway.com/instance/v1/zones/"

	//datacenter location code
	ParisZoneCode = "fr-par-1"

	AmsterdamZoneCode = "nl-ams-1"
)

type VolumeTemplate struct {
	// ID display the volumes unique ID
	ID string `json:"id,omitempty"`
	// Name display the volumes name
	Name string `json:"name,omitempty"`
	// Size display the volumes disk size
	Size uint64 `json:"size,omitempty"`
	// VolumeType display the volumes type
	//
	// Default value: l_ssd
	VolumeType VolumeType `json:"volume_type,omitempty"`
	// Organization the organization ID
	Organization string `json:"organization,omitempty"`
}

type ServerBootType string

const (
	ServerBootTypeLocal = ServerBootType("local")
)

type ServerCreateRequest struct {
	Name              string `json:"name"`
	DynamicIPRequired bool `json:"dynamic_ip_required"`
	CommercialType    string `json:"commercial_type"`
	Image             string `json:"image"`
	Volumes map[string]*VolumeTemplate `json:"volumes,omitempty"`
	EnableIpv6        bool   `json:"enable_ipv6"`
	BootType          ServerBootType   `json:"boot_type"`
	Organization      string   `json:"organization"`
	Tags              []string `json:"tags"`
}

type ServerAction struct {
	Action string `json:"action"`
}

type ScalewayAPI struct {
	// Api URL
	BaseUrl string

	// Organization is the identifier of the Scaleway organization
	Organization string

	// Token is the authentication token for the Scaleway organization
	Token string

	// Http Client for requests
	client     *http.Client

	// user agent for requests
	userAgent string

	//server location
	Region string
}

// ScalewayAPIError represents a Scaleway API Error
type ScalewayAPIError struct {
	// Message is a human-friendly error message
	APIMessage string `json:"message,omitempty"`

	// Type is a string code that defines the kind of error
	Type string `json:"type,omitempty"`

	// Fields contains detail about validation error
	Fields map[string][]string `json:"fields,omitempty"`

	// StatusCode is the HTTP status code received
	StatusCode int `json:"-"`

	// Message
	Message string `json:"-"`
}
type ServerList struct {
	Servers []Server `json:"servers"`
}

type Server struct {
	ID                string   `json:"id"`
	Name              string   `json:"name"`
	Organization      string   `json:"organization"`
	AllowedActions    []string `json:"allowed_actions"`
	Tags              []string `json:"tags"`
	CommercialType    string   `json:"commercial_type"`
	CreationDate      string   `json:"creation_date"`
	DynamicIPRequired string   `json:"dynamic_ip_required"`
	EnableIpv6        string   `json:"enable_ipv6"`
	ExtraNetworks     []string `json:"extra_networks"`
	Hostname          string   `json:"hostname"`
	Image             struct {
		ID                string `json:"id"`
		Name              string `json:"name"`
		Arch              string `json:"arch"`
		CreationDate      string `json:"creation_date"`
		ModificationDate  string `json:"modification_date"`
		DefaultBootscript struct {
			Bootcmdargs  string `json:"bootcmdargs"`
			Default      string `json:"default"`
			Dtb          string `json:"dtb"`
			ID           string `json:"id"`
			Initrd       string `json:"initrd"`
			Kernel       string `json:"kernel"`
			Organization string `json:"organization"`
			Public       string `json:"public"`
			Title        string `json:"title"`
			Arch         string `json:"arch"`
		} `json:"default_bootscript"`
		ExtraVolumes struct {
			ExtraVolumeKey struct {
				ID               string `json:"id"`
				Name             string `json:"name"`
				ExportURI        string `json:"export_uri"`
				Size             int    `json:"size"`
				VolumeType       string `json:"volume_type"`
				CreationDate     string `json:"creation_date"`
				ModificationDate string `json:"modification_date"`
				Organization     string `json:"organization"`
				Server           struct {
					ID   string `json:"id"`
					Name string `json:"name"`
				} `json:"server"`
				State string `json:"state"`
			} `json:"<extra_volumeKey>"`
		} `json:"extra_volumes"`
		FromServer struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"from_server"`
		Organization string `json:"organization"`
		Public       string `json:"public"`
		RootVolume   struct {
			ID         string `json:"id"`
			Name       string `json:"name"`
			Size       int    `json:"size"`
			VolumeType string `json:"volume_type"`
		} `json:"root_volume"`
		State string `json:"state"`
	} `json:"image"`
	Protected string `json:"protected"`
	PrivateIP string `json:"private_ip"`
	PublicIP  struct {
		ID      string `json:"id"`
		Address string `json:"address"`
		Dynamic string `json:"dynamic"`
	} `json:"public_ip"`
	ModificationDate string `json:"modification_date"`
	State            string `json:"state"`
	Location         struct {
		ClusterID    string `json:"cluster_id"`
		HypervisorID string `json:"hypervisor_id"`
		NodeID       string `json:"node_id"`
		PlatformID   string `json:"platform_id"`
		ZoneID       string `json:"zone_id"`
	} `json:"location"`
	Ipv6 struct {
		Address string `json:"address"`
		Gateway string `json:"gateway"`
		Netmask string `json:"netmask"`
	} `json:"ipv6"`
	Bootscript struct {
		Bootcmdargs  string `json:"bootcmdargs"`
		Default      string `json:"default"`
		Dtb          string `json:"dtb"`
		ID           string `json:"id"`
		Initrd       string `json:"initrd"`
		Kernel       string `json:"kernel"`
		Organization string `json:"organization"`
		Public       string `json:"public"`
		Title        string `json:"title"`
		Arch         string `json:"arch"`
	} `json:"bootscript"`
	BootType string `json:"boot_type"`
	Volumes  struct {
		VolumeKey struct {
			ID               string `json:"id"`
			Name             string `json:"name"`
			ExportURI        string `json:"export_uri"`
			Size             int    `json:"size"`
			VolumeType       string `json:"volume_type"`
			CreationDate     string `json:"creation_date"`
			ModificationDate string `json:"modification_date"`
			Organization     string `json:"organization"`
			Server           struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"server"`
			State string `json:"state"`
		} `json:"<volumeKey>"`
	} `json:"volumes"`
	SecurityGroup struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"security_group"`
	Maintenances []struct {
	} `json:"maintenances"`
	StateDetail    string `json:"state_detail"`
	Arch           string `json:"arch"`
	ComputeCluster struct {
		ID              string `json:"id"`
		Name            string `json:"name"`
		Organization    string `json:"organization"`
		PolicyMode      string `json:"policy_mode"`
		PolicyType      string `json:"policy_type"`
		PolicyRespected string `json:"policy_respected"`
	} `json:"compute_cluster"`
}

type ImageList struct {
	Images []Image `json:"images"`
}

type Image struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Arch              string `json:"arch"`
	CreationDate      string `json:"creation_date"`
	ModificationDate  string `json:"modification_date"`
	DefaultBootscript struct {
		Bootcmdargs  string `json:"bootcmdargs"`
		Default      string `json:"default"`
		Dtb          string `json:"dtb"`
		ID           string `json:"id"`
		Initrd       string `json:"initrd"`
		Kernel       string `json:"kernel"`
		Organization string `json:"organization"`
		Public       string `json:"public"`
		Title        string `json:"title"`
		Arch         string `json:"arch"`
	} `json:"default_bootscript"`
	ExtraVolumes struct {
		ExtraVolumeKey struct {
			ID               string `json:"id"`
			Name             string `json:"name"`
			ExportURI        string `json:"export_uri"`
			Size             int    `json:"size"`
			VolumeType       string `json:"volume_type"`
			CreationDate     string `json:"creation_date"`
			ModificationDate string `json:"modification_date"`
			Organization     string `json:"organization"`
			Server           struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"server"`
			State string `json:"state"`
		} `json:"<extra_volumeKey>"`
	} `json:"extra_volumes"`
	FromServer struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"from_server"`
	Organization string `json:"organization"`
	Public       string `json:"public"`
	RootVolume   struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		Size       int    `json:"size"`
		VolumeType string `json:"volume_type"`
	} `json:"root_volume"`
	State string `json:"state"`
}

type SecurityGroups struct {
	ID                    string `json:"id"`
	Name                  string `json:"name"`
	Description           string `json:"description"`
	EnableDefaultSecurity string `json:"enable_default_security"`
	InboundDefaultPolicy  string `json:"inbound_default_policy"`
	OutboundDefaultPolicy string `json:"outbound_default_policy"`
	Organization          string `json:"organization"`
	OrganizationDefault   string `json:"organization_default"`
	CreationDate          string `json:"creation_date"`
	ModificationDate      string `json:"modification_date"`
	Servers               []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"servers"`
	Stateful string `json:"stateful"`
}

type SecurityGroupsRules struct {
		ID           string `json:"id"`
		Protocol     string `json:"protocol"`
		Direction    string `json:"direction"`
		Action       string `json:"action"`
		IPRange      string `json:"ip_range"`
		DestPortFrom int    `json:"dest_port_from"`
		DestPortTo   int    `json:"dest_port_to"`
		Position     int    `json:"position"`
		Editable     string `json:"editable"`
}

type ServerPackages struct {
		ServerKey struct {
			MonthlyPrice        int      `json:"monthly_price"`
			HourlyPrice         int      `json:"hourly_price"`
			AltNames            []string `json:"alt_names"`
			PerVolumeConstraint struct {
				LSsd struct {
					MinSize int `json:"min_size"`
					MaxSize int `json:"max_size"`
				} `json:"l_ssd"`
			} `json:"per_volume_constraint"`
			VolumesConstraint struct {
				MinSize int `json:"min_size"`
				MaxSize int `json:"max_size"`
			} `json:"volumes_constraint"`
			Ncpus     int    `json:"ncpus"`
			Gpu       string `json:"gpu"`
			RAM       int    `json:"ram"`
			Arch      string `json:"arch"`
			Baremetal string `json:"baremetal"`
			Network   struct {
				Interfaces []struct {
					InternalBandwidth string `json:"internal_bandwidth"`
					InternetBandwidth string `json:"internet_bandwidth"`
				} `json:"interfaces"`
				SumInternalBandwidth string `json:"sum_internal_bandwidth"`
				SumInternetBandwidth string `json:"sum_internet_bandwidth"`
				Ipv6Support          string `json:"ipv6_support"`
			} `json:"network"`
		} `json:"<serverKey>"`
}

type IpList struct {
	IP []Ip `json:"ips"`
	TotalCount uint64 `json:"total_count"`
}

type ReseveIPResponse struct {
	IP Ip `json:"ip"`
	Location string `json:"location"`
}

type ReseveIP struct {
	Organization string `json:"organization"`
}

type IPServer struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Ip struct {
	ID      string `json:"id"`
	Address string `json:"address"`
	Reverse string `json:"reverse"`
	Server  IPServer `json:"server"`
	Organization string `json:"organization"`
}

type VolumeList struct {
	Volumes []struct {
		ID               string `json:"id"`
		Name             string `json:"name"`
		ExportURI        string `json:"export_uri"`
		Size             int    `json:"size"`
		VolumeType       string `json:"volume_type"`
		CreationDate     string `json:"creation_date"`
		ModificationDate string `json:"modification_date"`
		Organization     string `json:"organization"`
		Server           struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"server"`
		State string `json:"state"`
	} `json:"volumes"`
	TotalCount int `json:"total_count"`
}