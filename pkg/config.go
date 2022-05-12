package pkg

type ServerType int

const (
	ServerIdNil = ServerType(iota)
	ServerIdGt
	ServerIdGs
	ServerIdCs
	ServerIdDb
	ServerIdCl
	ServerIdRg
	ServerIdLs
	ServerIdWa
	ServerIdWeb
	ServerIdSg
	ServerIdVs
	ServerIdRc
	ServerIdAgent
	ServerIdCtrl
	ServerIdMd
	ServerIdMdAgent
	ServerIdWs
	ServerIdCt
	ServerIdAll
	ServerIdMax
)

type ZoneAttr struct {
	Ztest int
	Ztype int
	Zid   int
	Gid   int
	Bid   int
	Key   string
	Zname string
	Gname string
}

type LogAttr struct {
	Err    bool
	Wrn    bool
	Inf    bool
	Dbg    bool
	Output string
	Flush  int
}

type ListenAttr struct {
	Name    string
	Ip      string
	Port    int
	Recvbuf int
	Sendbuf int
}

type DatabaseAttr struct {
	Name     string
	Host     string
	Port     int
	User     string
	Password string
	Catalog  string
	Charset  string
}

type ServerAttr struct {
	Name    string
	Index   int
	Users   int
	Thread  int
	Listens map[ServerType]*ListenAttr
}

type ControllerAttr struct {
	Host    string
	Port    int
	Recvbuf int
	Sendbuf int
}

type Config interface {
	LoadConfig(path string) bool
	GetZoneId() int
	GetLogAttr() *LogAttr
	GetListenAttr(id ServerType, index int, subid ServerType) *ListenAttr
	GetServerAttr(id ServerType, index int) *ServerAttr
	Close()
	IsVirtual() bool
}
