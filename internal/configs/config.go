package configs

import (
	"fmt"
	"gateserver/internal/loggers"
	"strconv"

	"github.com/beevik/etree"
)

/*
<info>
	<gameinfo zoneid="1" groupid="18" bid="18" key="88888888" gamename="萌途2020" groupname="新萌途" type="0" istest="1"/>
	<fakelogin/>
	<ipinfo>
		<svr type="WS" index="1" users="4000" thread="0">
			<listen type="DB" ip="127.0.0.1" port="31111" recvbuf="4194304" sendbuf="4194304"/>
			<listen type="CT" ip="127.0.0.1" port="32222" recvbuf="1024000" sendbuf="1024000"/>
			<listen type="GS" ip="127.0.0.1" port="33333" recvbuf="4194304" sendbuf="4194304"/>
			<listen type="GT" ip="127.0.0.1" port="34444" recvbuf="1024000" sendbuf="1024000"/>
		</svr>
		<svr type="CT" index="1" users="4000" thread="0">
			<listen type="GS" ip="127.0.0.1" port="36666" recvbuf="1024000" sendbuf="1024000"/>
			<listen type="GT" ip="127.0.0.1" port="35555" recvbuf="1024000" sendbuf="1024000"/>
		</svr>
		<svr type="GS" index="1" users="4000" thread="0">
			<listen type="GT" ip="127.0.0.1" port="37777" recvbuf="4194304" sendbuf="4194304"/>
		</svr>
		<!-- svr type="GS" index="2" users="4000" thread="0">
			<listen type="CT" ip="127.0.0.1" port="6667" recvbuf="1024000" sendbuf="1024000" />
			<listen type="GT" ip="127.0.0.1" port="7778" recvbuf="4194304" sendbuf="4194304" />
		</svr -->
		<svr type="GT" index="1" users="4000" thread="0">
			<listen type="CLIENT" ip="0.0.0.0" port="8888" recvbuf="8191" sendbuf="65535"/>
		</svr>
		<svr type="GT" index="2" users="4000" thread="0">
			<listen type="CLIENT" ip="0.0.0.0" port="8889" recvbuf="8191" sendbuf="65535"/>
		</svr>
		<svr type="LS" index="1" users="4000" thread="0">
			<listen type="ALL" ip="127.0.0.1" port="10591" recvbuf="1024000" sendbuf="8192"/>
		</svr>
		<svr type="RG" index="1" users="4000" thread="0">
			<listen type="DB" ip="127.0.0.1" port="10592" recvbuf="65535" sendbuf="65535"/>
			<listen type="RC" ip="127.0.0.1" port="10593" recvbuf="65535" sendbuf="65535"/>
		</svr>
	</ipinfo>
	<dbinfo>
		<mysql type="DB" host="127.0.0.1" port="3306" user="root" password="root" db="dbmhxyzj1365806" charset="utf8"/>
		<mysql type="LS" host="127.0.0.1" port="3306" user="root" password="root" db="dbmhxyzj1239264_log" charset="utf8"/>
	</dbinfo>
	<runlog dir="./" flush="60" cri="open" wrn="open" inf="open" dbg="open"/>
	<ctrl host="127.0.0.1" port="10594" recvbuf="65535" sendbuf="65535"/>
</info>
*/

const (
	ServerIdNil     = 0
	ServerIdGt      = 1
	ServerIdGs      = 2
	ServerIdCs      = 3
	ServerIdDb      = 4
	ServerIdCl      = 5
	ServerIdRg      = 6
	ServerIdLs      = 7
	ServerIdWa      = 8
	ServerIdWeb     = 9
	ServerIdSg      = 10
	ServerIdVs      = 11
	ServerIdRc      = 12
	ServerIdAgent   = 13
	ServerIdCtrl    = 14
	ServerIdMd      = 15
	ServerIdMdAgent = 16
	ServerIdWs      = 17
	ServerIdCt      = 18
	ServerIdAll     = 19
	ServerIdMax     = 20
)

var serverNames = [...]string{
	"N/A",
	"GT",
	"GS",
	"CS",
	"DB",
	"CLIENT",
	"RG",
	"LS",
	"WA",
	"WEB",
	"SG",
	"VS",
	"RC",
	"AGENT",
	"CTRL",
	"MD",
	"MDAGENT",
	"WS",
	"CT",
	"ALL",
}

var serverIds = map[string]int{
	"N/A":     ServerIdNil,
	"GT":      ServerIdGt,
	"GS":      ServerIdGs,
	"CS":      ServerIdCs,
	"DB":      ServerIdDb,
	"CLIENT":  ServerIdCl,
	"RG":      ServerIdRg,
	"LS":      ServerIdLs,
	"WA":      ServerIdWa,
	"WEB":     ServerIdWeb,
	"SG":      ServerIdSg,
	"VS":      ServerIdVs,
	"RC":      ServerIdRc,
	"AGENT":   ServerIdAgent,
	"CTRL":    ServerIdCtrl,
	"MD":      ServerIdMd,
	"MDAGENT": ServerIdMdAgent,
	"WS":      ServerIdWs,
	"CT":      ServerIdCt,
	"ALL":     ServerIdAll,
}

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
	Flag   int
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
	Listens map[int]*ListenAttr
}

type ControllerAttr struct {
	Host    string
	Port    int
	Recvbuf int
	Sendbuf int
}

type Config struct {
	fakeAttr    bool
	zoneAttr    ZoneAttr
	logAttr     LogAttr
	ctrllerAttr ControllerAttr
	svrAttrs    map[int]map[int]*ServerAttr
	dbAttrs     map[int]*DatabaseAttr
}

func GetServerName(flag int) string {
	if flag <= ServerIdNil ||
		flag >= ServerIdMax {
		flag = ServerIdNil
	}

	return serverNames[flag]
}

func LoadConfig(path string) *Config {
	config := &Config{}

	doc := etree.NewDocument()
	if err := doc.ReadFromFile(path); err != nil {
		fmt.Println(err)
		return nil
	}

	root := doc.SelectElement("info")
	if root == nil {
		return nil
	}

	if !config.loadZoneAttr(root) {
		return nil
	}

	if !config.loadFakeAttr(root) {
		return nil
	}

	if !config.loadLogAttr(root) {
		return nil
	}

	if !config.loadCtrllerAttr(root) {
		return nil
	}

	if !config.loadServerAttr(root) {
		return nil
	}

	if !config.loadDbAttr(root) {
		return nil
	}

	return config
}

func (config *Config) loadZoneAttr(root *etree.Element) bool {
	zone := root.SelectElement("gameinfo")
	if zone == nil {
		return false
	}

	var err error
	config.zoneAttr.Zid, err = strconv.Atoi(zone.SelectAttrValue("zoneid", ""))
	if err != nil || config.zoneAttr.Zid == 0 {
		return false
	}

	config.zoneAttr.Gid, err = strconv.Atoi(zone.SelectAttrValue("groupid", ""))
	if err != nil {
		config.zoneAttr.Gid = 0
	}

	config.zoneAttr.Bid, err = strconv.Atoi(zone.SelectAttrValue("bid", ""))
	if err != nil {
		config.zoneAttr.Bid = 0
	}

	config.zoneAttr.Key = zone.SelectAttrValue("key", "")
	config.zoneAttr.Zname = zone.SelectAttrValue("gamename", "")
	config.zoneAttr.Gname = zone.SelectAttrValue("groupname", "")

	config.zoneAttr.Ztype, err = strconv.Atoi(zone.SelectAttrValue("type", ""))
	if err != nil {
		config.zoneAttr.Ztype = 0
	}

	config.zoneAttr.Ztest, err = strconv.Atoi(zone.SelectAttrValue("istest", ""))
	if err != nil {
		config.zoneAttr.Ztest = 0
	}
	return true
}

func (config *Config) loadFakeAttr(root *etree.Element) bool {
	fake := root.SelectElement("fakelogin")
	if fake != nil {
		config.fakeAttr = true
	}

	return true
}

func (config *Config) loadLogAttr(root *etree.Element) bool {
	log := root.SelectElement("runlog")
	if log == nil {
		return false
	}

	var err error
	config.logAttr.Output = log.SelectAttrValue("dir", "./")
	config.logAttr.Flush, err = strconv.Atoi(log.SelectAttrValue("flush", ""))
	if err != nil || config.logAttr.Flush <= 0 {
		config.logAttr.Flush = 30
	}

	config.logAttr.Flag = loggers.LogTypeSys
	if config.logAttr.Dbg = log.SelectAttrValue("dbg", "") == "open"; config.logAttr.Dbg {
		config.logAttr.Flag |= loggers.LogTypeDbg
	}

	if config.logAttr.Inf = log.SelectAttrValue("inf", "") == "open"; config.logAttr.Inf {
		config.logAttr.Flag |= loggers.LogTypeInf
	}

	if config.logAttr.Wrn = log.SelectAttrValue("wrn", "") == "open"; config.logAttr.Wrn {
		config.logAttr.Flag |= loggers.LogTypeWrn
	}

	if config.logAttr.Err = log.SelectAttrValue("cri", "") == "open"; config.logAttr.Err {
		config.logAttr.Flag |= loggers.LogTypeErr
	}

	return true
}

func (config *Config) loadCtrllerAttr(root *etree.Element) bool {
	ctrller := root.SelectElement("ctrl")
	if ctrller == nil {
		return false
	}

	var err error
	config.ctrllerAttr.Host = ctrller.SelectAttrValue("host", "")
	if config.ctrllerAttr.Host == "" {
		return false
	}

	config.ctrllerAttr.Port, err = strconv.Atoi(ctrller.SelectAttrValue("port", ""))
	if err != nil || config.ctrllerAttr.Port <= 0 {
		return false
	}

	config.ctrllerAttr.Recvbuf, err = strconv.Atoi(ctrller.SelectAttrValue("recvbuf", ""))
	if err != nil || config.ctrllerAttr.Recvbuf <= 0 {
		return false
	}

	config.ctrllerAttr.Sendbuf, err = strconv.Atoi(ctrller.SelectAttrValue("sendbuf", ""))
	if err != nil || config.ctrllerAttr.Sendbuf <= 0 {
		return false
	}
	return true
}

func (config *Config) loadServerAttr(root *etree.Element) bool {
	config.svrAttrs = make(map[int]map[int]*ServerAttr)

	servers := root.SelectElement("ipinfo")
	if servers == nil {
		return false
	}

	var err error
	for _, server := range servers.SelectElements("svr") {
		serverattr := &ServerAttr{}
		serverattr.Listens = make(map[int]*ListenAttr)

		serverattr.Index, err = strconv.Atoi(server.SelectAttrValue("index", ""))
		if err != nil || serverattr.Index <= 0 {
			return false
		}

		serverattr.Users, err = strconv.Atoi(server.SelectAttrValue("users", ""))
		if err != nil || serverattr.Users <= 0 {
			return false
		}

		serverattr.Thread, err = strconv.Atoi(server.SelectAttrValue("thread", ""))
		if err != nil {
			return false
		}

		for _, listen := range server.SelectElements("listen") {
			listenattr := &ListenAttr{}
			listenattr.Ip = listen.SelectAttrValue("ip", "")

			listenattr.Port, err = strconv.Atoi(listen.SelectAttrValue("port", ""))
			if err != nil || listenattr.Port <= 0 {
				return false
			}

			listenattr.Recvbuf, err = strconv.Atoi(listen.SelectAttrValue("recvbuf", ""))
			if err != nil || listenattr.Recvbuf <= 0 {
				return false
			}

			listenattr.Sendbuf, err = strconv.Atoi(listen.SelectAttrValue("sendbuf", ""))
			if err != nil || listenattr.Sendbuf <= 0 {
				return false
			}

			listenattr.Name = listen.SelectAttrValue("type", "N/A")
			if id, ok := serverIds[listenattr.Name]; ok && id != ServerIdNil {
				serverattr.Listens[id] = listenattr
			} else {
				return false
			}
		}

		serverattr.Name = server.SelectAttrValue("type", "N/A")
		if id, ok := serverIds[serverattr.Name]; ok && id != ServerIdNil {
			if _id, _ok := config.svrAttrs[id]; _ok {
				_id[serverattr.Index] = serverattr
			} else {
				_id := make(map[int]*ServerAttr)
				_id[serverattr.Index] = serverattr
				config.svrAttrs[id] = _id
			}
		} else {
			return false
		}
	}

	return true
}

func (config *Config) loadDbAttr(root *etree.Element) bool {
	config.dbAttrs = make(map[int]*DatabaseAttr)

	databases := root.SelectElement("dbinfo")
	if databases == nil {
		return false
	}

	var err error
	for _, database := range databases.SelectElements("mysql") {
		attr := &DatabaseAttr{}
		attr.Host = database.SelectAttrValue("host", "")

		attr.Port, err = strconv.Atoi(database.SelectAttrValue("port", ""))
		if err != nil || attr.Port <= 0 {
			return false
		}

		attr.User = database.SelectAttrValue("user", "")
		attr.Password = database.SelectAttrValue("password", "")
		attr.Catalog = database.SelectAttrValue("db", "")
		attr.Charset = database.SelectAttrValue("charset", "")

		attr.Name = database.SelectAttrValue("type", "N/A")
		if v, ok := serverIds[attr.Name]; ok && v != ServerIdNil {
			config.dbAttrs[v] = attr
		} else {
			return false
		}
	}

	return true
}

func (config *Config) GetZoneId() int {
	return config.zoneAttr.Zid
}

func (config *Config) GetLogAttr() *LogAttr {
	return &config.logAttr
}

func (config *Config) GetServerAttr(id, index, subid int) *ListenAttr {
	if v1, ok1 := config.svrAttrs[id]; ok1 {
		if v2, ok2 := v1[index]; ok2 {
			if v3, ok3 := v2.Listens[subid]; ok3 {
				return v3
			}
		}
	}

	return nil
}

func (config *Config) IsVirtual() bool {
	return config.fakeAttr
}
