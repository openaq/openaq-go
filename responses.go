package openaq

import "time"

type Meta struct {
	Name    string `json:"name"`
	Website string `json:"website"`
	Page    int64  `json:"page"`
	Limit   int64  `json:"limit"`
	Found   any    `json:"found"`
}

// a coordinate pair of latitude and longitude in WGS84
type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Datetime struct {
	UTC   time.Time `json:"utc"`
	Local time.Time `json:"local"`
}

type ProviderBase struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Units       string `json:"units"`
	DisplayName string `json:"displayName"`
}

type OwnerBase struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type InstrumentBase struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type ParameterBase struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Units       string `json:"units"`
	DisplayName string `json:"displayName"`
}

type SensorBase struct {
	ID        int64         `json:"id"`
	Name      string        `json:"name"`
	Parameter ParameterBase `json:"parameter"`
}

type OwnerEntityBase struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Instrument struct {
	ID             int64            `json:"id"`
	Name           string           `json:"name"`
	Manufacturer   ManufacturerBase `json:"manufacturer"`
	LocationsCount int64            `json:"locationsCount"`
}

type InstrumentsResponse struct {
	Meta    Meta         `json:"meta"`
	Results []Instrument `json:"results"`
}

type Parameter struct {
	ID                int64  `json:"id"`
	Name              string `json:"name"`
	Units             string `json:"units"`
	DisplayName       string `json:"displayName"`
	Description       string `json:"description"`
	LocationsCount    int64  `json:"locationsCount"`
	MeasurementsCount int64  `json:"measurementsCount"`
}

type ParametersResponse struct {
	Meta    Meta        `json:"meta"`
	Results []Parameter `json:"results"`
}

type Location struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Locality string `json:"locality"`
	Timezone string `json:"timezone"`
	Country  struct {
		ID   int64  `json:"id"`
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"country"`
	Owner struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"owner"`
	Provider      ProviderBase     `json:"provider"`
	IsMobile      bool             `json:"isMobile"`
	IsMonitor     bool             `json:"isMonitor"`
	Instruments   []InstrumentBase `json:"instruments"`
	Sensors       []SensorBase     `json:"sensors"`
	Coordinates   Coordinates      `json:"coordinates"`
	Bounds        []float64        `json:"bounds"`
	Distance      float64          `json:"distance"`
	DatetimeFirst Datetime         `json:"datetimeFirst"`
	DatetimeLast  Datetime         `json:"datetimeLast"`
}

type LocationsResponse struct {
	Meta    Meta       `json:"meta"`
	Results []Location `json:"results"`
}

// bounding box to define the geographic bounds of the data
// coordinates in the form of []
type Bbox struct {
	Type        string        `json:"type"`
	Coordinates []interface{} `json:"coordinates"`
}

type Country struct {
	ID                int64           `json:"id"`
	Code              string          `json:"code"`
	Name              string          `json:"name"`
	DatetimeFirst     time.Time       `json:"datetimeFirst"`
	DatetimeLast      time.Time       `json:"datetimeLast"`
	Parameters        []ParameterBase `json:"parameters"`
	LocationsCount    int64           `json:"locationsCount"`
	MeasurementsCount int64           `json:"measurementsCount"`
	ProvidersCount    int64           `json:"providersCount"`
}

type CountriesResponse struct {
	Meta    Meta      `json:"meta"`
	Results []Country `json:"results"`
}

type Provider struct {
	ID                int64           `json:"id"`
	Name              string          `json:"name"`
	SourceName        string          `json:"sourceName"`
	ExportPrefix      string          `json:"exportPrefix"`
	License           string          `json:"license"`
	DatetimeAdded     time.Time       `json:"datetimeAdded"`
	DatetimeFirst     time.Time       `json:"datetimeFirst"`
	DatetimeLast      time.Time       `json:"datetimeLast"`
	OwnerEntity       OwnerEntityBase `json:"ownerEntity"`
	LocationsCount    int64           `json:"locationsCount"`
	MeasurementsCount int64           `json:"measurementsCount"`
	CountriesCount    int64           `json:"countriesCount"`
	Parameters        []ParameterBase `json:"parameters"`
	Bbox              Bbox            `json:"bbox"`
}

type ProvidersResponse struct {
	Meta    Meta       `json:"meta"`
	Results []Provider `json:"results"`
}

type ManufacturerBase struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Manufacturer struct {
	ID             int64            `json:"id"`
	Name           string           `json:"name"`
	Instruments    []InstrumentBase `json:"instruments"`
	LocationsCount int64            `json:"locationsCount"`
}

type ManufacturersResponse struct {
	Meta    Meta           `json:"meta"`
	Results []Manufacturer `json:"results"`
}

type period struct {
	Label        string   `json:"label"`
	Interval     string   `json:"interval"`
	DatetimeFrom Datetime `json:"datetimeFrom"`
	DatetimeTo   Datetime `json:"datetimeTo"`
}

type summary struct {
	Min    float64 `json:"min"`
	Q02    float64 `json:"q02"`
	Q25    float64 `json:"q25"`
	Median float64 `json:"median"`
	Q75    float64 `json:"q75"`
	Q98    float64 `json:"q98"`
	Max    float64 `json:"max"`
	StdDev float64 `json:"sd"`
}

type coverage struct {
	ExpectedCount    int64    `json:"expectedCount"`
	ExpectedInterval string   `json:"expectedInterval"`
	ObservedCount    int64    `json:"observedCount"`
	ObservedInterval string   `json:"observedInterval"`
	PercentComplete  float64  `json:"percentComplete"`
	PercentCoverage  float64  `json:"percentCoverage"`
	DatetimeFrom     Datetime `json:"datetimeFrom"`
	DatetimeTo       Datetime `json:"datetimeTo"`
}

type Measurement struct {
	Period      period        `json:"period"`
	Value       float64       `json:"value"`
	Parameter   ParameterBase `json:"parameter"`
	Coordinates Coordinates   `json:"coordinates"`
	Summary     summary       `json:"summary"`
	Coverage    coverage      `json:"coverage"`
}

type MeasurementsResponse struct {
	Meta    Meta          `json:"meta"`
	Results []Measurement `json:"results"`
}

type Owner struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	LocationsCount int64  `json:"locationsCount"`
}

type OwnersResponse struct {
	Meta    Meta    `json:"meta"`
	Results []Owner `json:"results"`
}

type SensorsResponse struct{}

type TrendsResponse struct{}
