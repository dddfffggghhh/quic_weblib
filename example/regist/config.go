package regist

import (
	"fmt"
)

//eureka client config
type Config struct {
	//eureka server address
	DefaultZone string
	//heart beat interval, default 30s
	RenewalIntervalInSecs int
	//get serve list interval, default 15s
	RegistryFetchIntervalSeconds int
	//duration interval, default 90s
	DurationInSecs int
	//application name
	App string
	//port number
	Port int
	Metadata map[string]interface{}
	instance *Instance
}

// Applications eureka server register apps
type Applications struct {
	VersionsDelta string        `xml:"versions__delta,omitempty" json:"versions__delta,omitempty"`
	AppsHashcode  string        `xml:"apps__hashcode,omitempty" json:"apps__hashcode,omitempty"`
	Applications  []Application `xml:"application,omitempty" json:"application,omitempty"`
}

// Application
type Application struct {
	Name      string     `xml:"name" json:"name"`
	Instances []Instance `xml:"instance" json:"instance"`
}

// Instance
type Instance struct {
	HostName                      string                 `xml:"hostName" json:"hostName"`
	HomePageURL                   string                 `xml:"homePageUrl,omitempty" json:"homePageUrl,omitempty"`
	StatusPageURL                 string                 `xml:"statusPageUrl" json:"statusPageUrl"`
	HealthCheckURL                string                 `xml:"healthCheckUrl,omitempty" json:"healthCheckUrl,omitempty"`
	App                           string                 `xml:"app" json:"app"`
	IPAddr                        string                 `xml:"ipAddr" json:"ipAddr"`
	VipAddress                    string                 `xml:"vipAddress" json:"vipAddress"`
	SecureVipAddress              string                 `xml:"secureVipAddress,omitempty" json:"secureVipAddress,omitempty"`
	Status                        string                 `xml:"status" json:"status"`
	Port                          *Port                  `xml:"port,omitempty" json:"port,omitempty"`
	SecurePort                    *Port                  `xml:"securePort,omitempty" json:"securePort,omitempty"`
	DataCenterInfo                *DataCenterInfo        `xml:"dataCenterInfo" json:"dataCenterInfo"`
	LeaseInfo                     *LeaseInfo             `xml:"leaseInfo,omitempty" json:"leaseInfo,omitempty"`
	Metadata                      map[string]interface{} `xml:"metadata,omitempty" json:"metadata,omitempty"`
	IsCoordinatingDiscoveryServer string                 `xml:"isCoordinatingDiscoveryServer,omitempty" json:"isCoordinatingDiscoveryServer,omitempty"`
	LastUpdatedTimestamp          string                 `xml:"lastUpdatedTimestamp,omitempty" json:"lastUpdatedTimestamp,omitempty"`
	LastDirtyTimestamp            string                 `xml:"lastDirtyTimestamp,omitempty" json:"lastDirtyTimestamp,omitempty"`
	ActionType                    string                 `xml:"actionType,omitempty" json:"actionType,omitempty"`
	OverriddenStatus              string                 `xml:"overriddenstatus,omitempty" json:"overriddenstatus,omitempty"`
	CountryID                     int                    `xml:"countryId,omitempty" json:"countryId,omitempty"`
	InstanceID                    string                 `xml:"instanceId,omitempty" json:"instanceId,omitempty"`
}

// Port
type Port struct {
	Port    int    `xml:",chardata" json:"$"`
	Enabled string `xml:"enabled,attr" json:"@enabled"`
}

// DataCenterInfo
type DataCenterInfo struct {
	Name     string              `xml:"name" json:"name"`
	Class    string              `xml:"class,attr" json:"@class"`
	Metadata *DataCenterMetadata `xml:"metadata,omitempty" json:"metadata,omitempty"`
}

// DataCenterMetadata
type DataCenterMetadata struct {
	AmiLaunchIndex   string `xml:"ami-launch-index,omitempty" json:"ami-launch-index,omitempty"`
	LocalHostname    string `xml:"local-hostname,omitempty" json:"local-hostname,omitempty"`
	AvailabilityZone string `xml:"availability-zone,omitempty" json:"availability-zone,omitempty"`
	InstanceID       string `xml:"instance-id,omitempty" json:"instance-id,omitempty"`
	PublicIpv4       string `xml:"public-ipv4,omitempty" json:"public-ipv4,omitempty"`
	PublicHostname   string `xml:"public-hostname,omitempty" json:"public-hostname,omitempty"`
	AmiManifestPath  string `xml:"ami-manifest-path,omitempty" json:"ami-manifest-path,omitempty"`
	LocalIpv4        string `xml:"local-ipv4,omitempty" json:"local-ipv4,omitempty"`
	Hostname         string `xml:"hostname,omitempty" json:"hostname,omitempty"`
	AmiID            string `xml:"ami-id,omitempty" json:"ami-id,omitempty"`
	InstanceType     string `xml:"instance-type,omitempty" json:"instance-type,omitempty"`
}

// LeaseInfo
type LeaseInfo struct {
	RenewalIntervalInSecs int `xml:"renewalIntervalInSecs,omitempty" json:"renewalIntervalInSecs,omitempty"`
	DurationInSecs        int `xml:"durationInSecs,omitempty" json:"durationInSecs,omitempty"`
}

// NewInstance
func NewInstance(ip string, config *Config) *Instance {
	instance := &Instance{
		InstanceID: fmt.Sprintf("%s:%s:%d", ip, config.App, config.Port),
		HostName:   ip,
		App:        config.App,
		IPAddr:     ip,
		Port: &Port{
			Port:    config.Port,
			Enabled: "true",
		},
		VipAddress:       config.App,
		SecureVipAddress: config.App,
		// lease information
		LeaseInfo: &LeaseInfo{
			RenewalIntervalInSecs: config.RenewalIntervalInSecs,
			DurationInSecs:        config.DurationInSecs,
		},
		Status:           "UP",
		OverriddenStatus: "UNKNOWN",
		// data center
		DataCenterInfo: &DataCenterInfo{
			Name:  "MyOwn",
			Class: "com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo",
		},
		// ?????????
		Metadata: config.Metadata,
	}
	instance.HomePageURL = fmt.Sprintf("http://%s:%d", ip, config.Port)
	instance.StatusPageURL = fmt.Sprintf("http://%s:%d/info", ip, config.Port)
	return instance
}

