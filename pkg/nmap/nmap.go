package nmap

type Run struct {
	Id               int64
	Start            float64
	Version          float64
	XmlOutputVersion float64
	Args             string
	StartStr         string
	Scanner          string
	ScanInfo         ScanInfo
	Verbose          Verbose `gorm:"foreignkey:RunId"`
	Debugging        Debugging
	FinishedStat     FinishedStat
	HostStat         HostStat
}

type ScanInfo struct {
	Id          int64
	RunId       int64
	NumServices int64
	Type        string
	Protocol    string
	Services    string
}

type Verbose struct {
	Id    int64
	RunId int64
}

type Debugging struct {
	Id    int64
	RunId int64
}

type Host struct {
	Id        int64
	RunId     int64
	EndTime   int64
	StartTime int64
	Status    HostStatus
	Address   HostAddress
}

type HostAddress struct {
	Id       int64
	RunId    int64
	HostId   int64
	AddrType string
	Addr     string
}

type HostName struct {
	Id     int64
	RunId  int64
	HostId int64
	Name   string
	Type   string
}

type HostStatus struct {
	Id        int64
	RunId     int64
	HostId    int64
	ReasonTTL int64
	State     string
	Reason    string
	ExtraInfo string
}

type HostTime struct {
	Id     int64
	RunId  int64
	HostId int64
	To     int64
	SRTT   int64
	RTTVAR int64
}

type HostPort struct {
	Id       int64
	RunId    int64
	HostId   int64
	PortId   int64
	Protocol string
	Service  PortService `gorm:"foreignkey:PortId"`
	State    PortState   `gorm:"foreignkey:PortId"`
}

type PortService struct {
	Id     int64
	RunId  int64
	PortId int64
	Conf   int64
	Name   string
	Method string
}

type PortState struct {
	Id        int64
	RunId     int64
	PortId    int64
	State     string
	Reason    string
	ReasonTTL string
}

type Task struct {
	Id        int64
	RunId     int64
	Time      int64
	Task      string
	Type      string
	ExtraInfo string
}

type FinishedStat struct {
	Id      int64
	RunId   int64
	Elapsed int64
	Time    int64
	TimeStr string
	Summary string
	Exit    string
}

type HostStat struct {
	Id    int64
	RunId int64
	Up    int64
	Down  int64
	Total int64
}

func (r Run) TableName() string {
	return "nmap_run"
}

func (s ScanInfo) TableName() string {
	return "scan_info"
}

func (v Verbose) TableName() string {
	return "verbose"
}

func (d Debugging) TableName() string {
	return "debugging"
}

func (f FinishedStat) TableName() string {
	return "run_finished_stat"
}

func (h HostStat) TableName() string {
	return "run_host_stat"
}

func (t Task) TableName() string {
	return "task"
}

func (h Host) TableName() string {
	return "host"
}

func (h HostAddress) TableName() string {
	return "host_address"
}

func (h HostName) TableName() string {
	return "host_name"
}

func (h HostPort) TableName() string {
	return "host_port"
}

func (h HostStatus) TableName() string {
	return "host_status"
}

func (h HostTime) TableName() string {
	return "host_time"
}

func (p PortService) TableName() string {
	return "port_service"
}

func (p PortState) TableName() string {
	return "port_state"
}
