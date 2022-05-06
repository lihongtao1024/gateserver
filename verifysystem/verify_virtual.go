package verifysystem

import (
	"bufio"
	"encoding/json"
	"gateserver/internal/errors"
	"gateserver/logsystem"
	"gateserver/netsystem/clients"
	"io/ioutil"
	"os"
)

const jsonFile = "./user.json"
const jsonRead = os.O_RDONLY | os.O_CREATE
const jsonWrite = os.O_WRONLY | os.O_CREATE

type virtualUser struct {
	User string `json:"user"`
	Uid  uint32 `json:"uid"`
}

type virtualUsers struct {
	Users []virtualUser `json:"users"`
}

type virtualImpl struct {
	isDirty bool
	maxUid  uint32
	userMap map[string]uint32
}

func (impl *virtualImpl) loadJson() bool {
	impl.userMap = make(map[string]uint32)
	impl.maxUid = 0

	file, err := os.OpenFile(jsonFile, jsonRead, os.ModePerm)
	if err != nil {
		logsystem.Instance.Err("load \"%s\" fail, errmsg: \"%s\"", jsonFile, err)
		return false
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		logsystem.Instance.Err("load \"%s\" fail, errmsg: \"%s\"", jsonFile, err)
		return false
	}

	if len(data) == 0 {
		return true
	}

	users := &virtualUsers{}
	err = json.Unmarshal(data, users)
	if err != nil {
		logsystem.Instance.Err("load \"%s\" fail, errmsg: \"%s\"", jsonFile, err)
		return false
	}

	for i := 0; i < len(users.Users); i++ {
		account := &users.Users[i]
		impl.userMap[account.User] = account.Uid

		if account.Uid > impl.maxUid {
			impl.maxUid = account.Uid
		}
	}

	return true
}

func (impl *virtualImpl) PostRequest(client *clients.Client) {
	uname := client.GetUName()
	uid, ok := impl.userMap[uname]
	if !ok {
		impl.isDirty = true
		impl.maxUid++
		uid = impl.maxUid
		impl.userMap[uname] = uid
	}

	client.SetUid(uid)
	Instance.ReceiveResponse(client, nil)
}

func (impl *virtualImpl) ReceiveResponse(client *clients.Client, err error) {
	client.SendLoginAck(errors.NewError(errors.ErrorOk))
}

func (impl *virtualImpl) Close() {
	if !impl.isDirty {
		return
	}

	users := &virtualUsers{Users: make([]virtualUser, 0)}
	for user, uid := range impl.userMap {
		users.Users = append(users.Users, virtualUser{user, uid})
	}

	data, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		logsystem.Instance.Err("save \"%s\" fail, errmsg: \"%s\"", jsonFile, err)
		return
	}

	file, err := os.OpenFile(jsonFile, jsonWrite, os.ModePerm)
	if err != nil {
		logsystem.Instance.Err("save \"%s\" fail, errmsg: \"%s\"", jsonFile, err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.Write(data)
	writer.Flush()
}
