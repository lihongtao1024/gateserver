package configs

import (
	"fmt"
	"gateserver/pkg"
	"os"
	"strconv"

	"github.com/beevik/etree"
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

var serverIds = map[string]pkg.ServerType{
	"N/A":     pkg.ServerIdNil,
	"GT":      pkg.ServerIdGt,
	"GS":      pkg.ServerIdGs,
	"CS":      pkg.ServerIdCs,
	"DB":      pkg.ServerIdDb,
	"CLIENT":  pkg.ServerIdCl,
	"RG":      pkg.ServerIdRg,
	"LS":      pkg.ServerIdLs,
	"WA":      pkg.ServerIdWa,
	"WEB":     pkg.ServerIdWeb,
	"SG":      pkg.ServerIdSg,
	"VS":      pkg.ServerIdVs,
	"RC":      pkg.ServerIdRc,
	"AGENT":   pkg.ServerIdAgent,
	"CTRL":    pkg.ServerIdCtrl,
	"MD":      pkg.ServerIdMd,
	"MDAGENT": pkg.ServerIdMdAgent,
	"WS":      pkg.ServerIdWs,
	"CT":      pkg.ServerIdCt,
	"ALL":     pkg.ServerIdAll,
}

type configImpl struct {
	fakeAttr    bool
	zoneAttr    pkg.ZoneAttr
	logAttr     pkg.LogAttr
	ctrllerAttr pkg.ControllerAttr
	svrAttrs    map[pkg.ServerType]map[int]*pkg.ServerAttr
	dbAttrs     map[pkg.ServerType]*pkg.DatabaseAttr
}

func NewConfig() pkg.Config {
	return &configImpl{}
}

func (config *configImpl) LoadConfig(path string) bool {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(path); err != nil {
		fmt.Println(err)
		return false
	}

	root := doc.SelectElement("info")
	if root == nil {
		return false
	}

	if !config.loadZoneAttr(root) {
		return false
	}

	if !config.loadFakeAttr(root) {
		return false
	}

	if !config.loadLogAttr(root) {
		return false
	}

	if !config.loadCtrllerAttr(root) {
		return false
	}

	if !config.loadServerAttr(root) {
		return false
	}

	if !config.loadDbAttr(root) {
		return false
	}

	return true
}

func (config *configImpl) loadZoneAttr(root *etree.Element) bool {
	zone := root.SelectElement("gameinfo")
	if zone == nil {
		return false
	}

	var err error
	config.zoneAttr.Zid, err = strconv.Atoi(zone.SelectAttrValue("zoneid", ""))
	if err != nil || config.zoneAttr.Zid == 0 {
		fmt.Fprintf(
			os.Stderr,
			"load config fail, errmsg: \"unexpected '%s'\".\n",
			"gameinfo->zonzid",
		)
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

func (config *configImpl) loadFakeAttr(root *etree.Element) bool {
	fake := root.SelectElement("fakelogin")
	if fake != nil {
		config.fakeAttr = true
	}

	return true
}

func (config *configImpl) loadLogAttr(root *etree.Element) bool {
	log := root.SelectElement("runlog")
	if log == nil {
		fmt.Fprintf(
			os.Stderr,
			"load config fail, errmsg: \"unexpected '%s'\".\n",
			"gameinfo->runlog",
		)
		return false
	}

	var err error
	config.logAttr.Output = log.SelectAttrValue("dir", "./")
	config.logAttr.Flush, err = strconv.Atoi(log.SelectAttrValue("flush", ""))
	if err != nil || config.logAttr.Flush <= 0 {
		config.logAttr.Flush = 30
	}

	config.logAttr.Dbg = log.SelectAttrValue("dbg", "") == "open"
	config.logAttr.Inf = log.SelectAttrValue("inf", "") == "open"
	config.logAttr.Wrn = log.SelectAttrValue("wrn", "") == "open"
	config.logAttr.Err = log.SelectAttrValue("cri", "") == "open"
	return true
}

func (config *configImpl) loadCtrllerAttr(root *etree.Element) bool {
	ctrller := root.SelectElement("ctrl")
	if ctrller == nil {
		fmt.Fprintf(
			os.Stderr,
			"load config fail, errmsg: \"unexpected '%s'\".\n",
			"gameinfo->ctrl",
		)
		return false
	}

	var err error
	config.ctrllerAttr.Host = ctrller.SelectAttrValue("host", "")
	if config.ctrllerAttr.Host == "" {
		fmt.Fprintf(
			os.Stderr,
			"load config fail, errmsg: \"unexpected '%s'\".\n",
			"gameinfo->ctrl->host",
		)
		return false
	}

	config.ctrllerAttr.Port, err = strconv.Atoi(ctrller.SelectAttrValue("port", ""))
	if err != nil || config.ctrllerAttr.Port <= 0 {
		fmt.Fprintf(
			os.Stderr,
			"load config fail, errmsg: \"unexpected '%s'\".\n",
			"gameinfo->ctrl->port",
		)
		return false
	}

	config.ctrllerAttr.Recvbuf, err = strconv.Atoi(ctrller.SelectAttrValue("recvbuf", ""))
	if err != nil || config.ctrllerAttr.Recvbuf <= 0 {
		fmt.Fprintf(
			os.Stderr,
			"load config fail, errmsg: \"unexpected '%s'\".\n",
			"gameinfo->ctrl->recvbuf",
		)
		return false
	}

	config.ctrllerAttr.Sendbuf, err = strconv.Atoi(ctrller.SelectAttrValue("sendbuf", ""))
	if err != nil || config.ctrllerAttr.Sendbuf <= 0 {
		fmt.Fprintf(
			os.Stderr,
			"load config fail, errmsg: \"unexpected '%s'\".\n",
			"gameinfo->ctrl->sendbuf",
		)
		return false
	}
	return true
}

func (config *configImpl) loadServerAttr(root *etree.Element) bool {
	config.svrAttrs = make(map[pkg.ServerType]map[int]*pkg.ServerAttr)

	servers := root.SelectElement("ipinfo")
	if servers == nil {
		fmt.Fprintf(
			os.Stderr,
			"load config fail, errmsg: \"unexpected '%s'\".\n",
			"gameinfo->ipinfo",
		)
		return false
	}

	var err error
	for _, server := range servers.SelectElements("svr") {
		serverattr := &pkg.ServerAttr{}
		serverattr.Listens = make(map[pkg.ServerType]*pkg.ListenAttr)

		serverattr.Index, err = strconv.Atoi(server.SelectAttrValue("index", ""))
		if err != nil || serverattr.Index <= 0 {
			fmt.Fprintf(
				os.Stderr,
				"load config fail, errmsg: \"unexpected '%s'\".\n",
				"gameinfo->ipinfo->svr->index",
			)
			return false
		}

		serverattr.Users, err = strconv.Atoi(server.SelectAttrValue("users", ""))
		if err != nil || serverattr.Users <= 0 {
			fmt.Fprintf(
				os.Stderr,
				"load config fail, errmsg: \"unexpected '%s'\".\n",
				"gameinfo->ipinfo->svr->users",
			)
			return false
		}

		serverattr.Thread, err = strconv.Atoi(server.SelectAttrValue("thread", ""))
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"load config fail, errmsg: \"unexpected '%s'\".\n",
				"gameinfo->ipinfo->svr->thread",
			)
			return false
		}

		for _, listen := range server.SelectElements("listen") {
			listenattr := &pkg.ListenAttr{}
			listenattr.Ip = listen.SelectAttrValue("ip", "")
			if listenattr.Ip == "" {
				fmt.Fprintf(
					os.Stderr,
					"load config fail, errmsg: \"unexpected '%s'\".\n",
					"gameinfo->ipinfo->svr->listen->ip",
				)
				return false
			}

			listenattr.Port, err = strconv.Atoi(listen.SelectAttrValue("port", ""))
			if err != nil || listenattr.Port <= 0 {
				fmt.Fprintf(
					os.Stderr,
					"load config fail, errmsg: \"unexpected '%s'\".\n",
					"gameinfo->ipinfo->svr->listen->port",
				)
				return false
			}

			listenattr.Recvbuf, err = strconv.Atoi(listen.SelectAttrValue("recvbuf", ""))
			if err != nil || listenattr.Recvbuf <= 0 {
				fmt.Fprintf(
					os.Stderr,
					"load config fail, errmsg: \"unexpected '%s'\".\n",
					"gameinfo->ipinfo->svr->listen->recvbuf",
				)
				return false
			}

			listenattr.Sendbuf, err = strconv.Atoi(listen.SelectAttrValue("sendbuf", ""))
			if err != nil || listenattr.Sendbuf <= 0 {
				fmt.Fprintf(
					os.Stderr,
					"load config fail, errmsg: \"unexpected '%s'\".\n",
					"gameinfo->ipinfo->svr->listen->sendbuf",
				)
				return false
			}

			listenattr.Name = listen.SelectAttrValue("type", "N/A")
			if id, ok := serverIds[listenattr.Name]; ok && id != pkg.ServerIdNil {
				serverattr.Listens[id] = listenattr
			} else {
				fmt.Fprintf(
					os.Stderr,
					"load config fail, errmsg: \"unexpected '%s'\".\n",
					"gameinfo->ipinfo->svr->listen->type",
				)
				return false
			}
		}

		serverattr.Name = server.SelectAttrValue("type", "N/A")
		if id, ok := serverIds[serverattr.Name]; ok && id != pkg.ServerIdNil {
			if _id, _ok := config.svrAttrs[id]; _ok {
				_id[serverattr.Index] = serverattr
			} else {
				_id := make(map[int]*pkg.ServerAttr)
				_id[serverattr.Index] = serverattr
				config.svrAttrs[id] = _id
			}
		} else {
			fmt.Fprintf(
				os.Stderr,
				"load config fail, errmsg: \"unexpected '%s'\".\n",
				"gameinfo->ipinfo->svr->type",
			)
			return false
		}
	}

	return true
}

func (config *configImpl) loadDbAttr(root *etree.Element) bool {
	config.dbAttrs = make(map[pkg.ServerType]*pkg.DatabaseAttr)

	databases := root.SelectElement("dbinfo")
	if databases == nil {
		fmt.Fprintf(
			os.Stderr,
			"load config fail, errmsg: \"unexpected '%s'\".\n",
			"gameinfo->dbinf",
		)
		return false
	}

	var err error
	for _, database := range databases.SelectElements("mysql") {
		attr := &pkg.DatabaseAttr{}
		attr.Host = database.SelectAttrValue("host", "")
		if attr.Host == "" {
			fmt.Fprintf(
				os.Stderr,
				"load config fail, errmsg: \"unexpected '%s'\".\n",
				"gameinfo->dbinf->mysql->host",
			)
			return false
		}

		attr.Port, err = strconv.Atoi(database.SelectAttrValue("port", ""))
		if err != nil || attr.Port <= 0 {
			fmt.Fprintf(
				os.Stderr,
				"load config fail, errmsg: \"unexpected '%s'\".\n",
				"gameinfo->dbinf->mysql->port",
			)
			return false
		}

		attr.User = database.SelectAttrValue("user", "")
		if attr.User == "" {
			fmt.Fprintf(
				os.Stderr,
				"load config fail, errmsg: \"unexpected '%s'\".\n",
				"gameinfo->dbinf->mysql->user",
			)
			return false
		}

		attr.Password = database.SelectAttrValue("password", "")
		if attr.Password == "" {
			fmt.Fprintf(
				os.Stderr,
				"load config fail, errmsg: \"unexpected '%s'\".\n",
				"gameinfo->dbinf->mysql->password",
			)
			return false
		}

		attr.Catalog = database.SelectAttrValue("db", "")
		if attr.Catalog == "" {
			fmt.Fprintf(
				os.Stderr,
				"load config fail, errmsg: \"unexpected '%s'\".\n",
				"gameinfo->dbinf->mysql->db",
			)
			return false
		}

		attr.Charset = database.SelectAttrValue("charset", "")
		if attr.Charset == "" {
			fmt.Fprintf(
				os.Stderr,
				"load config fail, errmsg: \"unexpected '%s'\".\n",
				"gameinfo->dbinf->mysql->charset",
			)
			return false
		}

		attr.Name = database.SelectAttrValue("type", "N/A")
		if v, ok := serverIds[attr.Name]; ok && v != pkg.ServerIdNil {
			config.dbAttrs[v] = attr
		} else {
			fmt.Fprintf(
				os.Stderr,
				"load config fail, errmsg: \"unexpected '%s'\".\n",
				"gameinfo->dbinf->mysql->type",
			)
			return false
		}
	}

	return true
}

func (config *configImpl) GetZoneId() int {
	return config.zoneAttr.Zid
}

func (config *configImpl) GetLogAttr() *pkg.LogAttr {
	return &config.logAttr
}

func (config *configImpl) GetListenAttr(id pkg.ServerType, index int,
	subid pkg.ServerType) *pkg.ListenAttr {
	if v1, ok1 := config.svrAttrs[id]; ok1 {
		if v2, ok2 := v1[index]; ok2 {
			if v3, ok3 := v2.Listens[subid]; ok3 {
				return v3
			}
		}
	}

	return nil
}

func (config *configImpl) GetServerAttr(id pkg.ServerType,
	index int) *pkg.ServerAttr {
	if v1, ok1 := config.svrAttrs[id]; ok1 {
		if v2, ok2 := v1[index]; ok2 {
			return v2
		}
	}

	return nil
}

func (config *configImpl) IsVirtual() bool {
	return config.fakeAttr
}

func (config *configImpl) Close() {

}

func GetServerName(typ1 pkg.ServerType) string {
	if typ1 <= pkg.ServerIdNil ||
		typ1 >= pkg.ServerIdMax {
		typ1 = pkg.ServerIdNil
	}

	return serverNames[typ1]
}
