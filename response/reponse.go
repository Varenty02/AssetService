package response

type AutonomousSystem struct {
	ASN         int    `json:"asn,omitempty"`
	Prefix      string `json:"bgp_prefix,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
}

// DNS
type ReverseDNS struct {
	Names []string `json:"names,omitempty"`
}
type DNS struct {
	ReverseDNS ReverseDNS `json:"reverse_dns,omitempty"`
}

// location
type Coordinates struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}
type Location struct {
	City string `json:"city,omitempty"`
	//Continent   string      `json:"continent"`
	//Coordinates Coordinates `json:"coordinates"`
	Country string `json:"country,omitempty"`
	//CountryCode string `json:"country_code"`
	//Province    string `json:"province"`
	Timezone string `json:"timezone,omitempty"`
}

// Response
type Service struct {
	ExtendedServiceName string `json:"extended_service_name,omitempty"`
	Port                int    `json:"port,omitempty"`
	ServiceName         string `json:"service_name,omitempty"`
	TransportProtocol   string `json:"transport_protocol,omitempty"`
	Certificate         string `json:"certificate,omitempty,omitempty"`
}
type Data struct {
	ASN           AutonomousSystem `json:"autonomous_system,omitempty"`
	DNS           DNS              `json:"dns,omitempty"`
	IP            string           `json:"ip,omitempty"`
	LastUpdatedAt string           `json:"last_updated_at,omitempty"`
	Location      Location         `json:"location,omitempty"`
	Services      []Service        `json:"services,omitempty"`
}
type Response struct {
	Ip         string       `json:"ip,omitempty"`
	Data       []Data       `json:"data,omitempty"`
	SameSubnet []SameSubnet `json:"same_subnet,omitempty"`
}
type SameSubnet struct {
	Prefix string   `json:"prefix,omitempty"`
	Ips    []string `json:"ips,omitempty,omitempty"`
}
