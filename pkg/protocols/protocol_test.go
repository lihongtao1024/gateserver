package protocols

import (
	"bufio"
	"errors"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/beevik/etree"
)

type MoudleAttr struct {
	name    string
	moudle  string
	version string
	comment string
}

type PackageDeclare struct {
	name string
}

type StructDeclare struct {
	name    string
	comment string
}

type FieldAttr struct {
	def     string
	name    string
	comment string
	array   string
	alias   string
	length  string
}

type DefStruct struct {
	isproto bool
	declare StructDeclare
	fields  []*FieldAttr
}

type DefProtoMoudle struct {
	isproto   bool
	moudle    MoudleAttr
	protocols []*DefStruct
	pack      *PackageDeclare
}

func firstToUpper(s string) (r string) {
	wchars := []rune(s)
	for k, v := range wchars {
		if k == 0 {
			if v >= 'a' && v <= 'z' {
				v -= 'a' - 'A'
				r += string(v)
				continue
			}
		}

		r += string(v)
	}
	return
}

func parseProtocol(path string) (*DefProtoMoudle, error) {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(path); err != nil {
		return nil, err
	}

	root := doc.SelectElement("protocol_define")
	if root == nil {
		return nil, errors.New("cannot find 'protocol_define'")
	}

	var protomoudle DefProtoMoudle
	protomoudle.protocols = make([]*DefStruct, 0)

	pack := root.SelectElement("package")
	if pack != nil {
		packdeclare := &PackageDeclare{}

		packdeclare.name = pack.SelectAttrValue("name", "")
		if packdeclare.name == "" {
			return nil, errors.New("invalid 'package->name'")
		}

		protomoudle.pack = packdeclare
	}

	protomoudle.moudle.name = root.SelectAttrValue("name", "")
	if protomoudle.moudle.name == "" {
		return nil, errors.New("invalid 'protocol_define->name'")
	}

	protomoudle.moudle.moudle = root.SelectAttrValue("moudleid", "")
	if protomoudle.moudle.moudle == "" {
		return nil, errors.New("invalid 'protocol_define->moudleid'")
	}

	protomoudle.moudle.version = root.SelectAttrValue("version", "")
	protomoudle.moudle.comment = root.SelectAttrValue("comment", "")

	datas := root.SelectElements("data")
	for _, data := range datas {
		var defstruct DefStruct
		defstruct.isproto = false

		t := data.SelectAttrValue("type", "")
		switch t {
		case "protocol":
			{
				defstruct.isproto = true
			}
			fallthrough
		case "struct":
			{
				defstruct.declare.name = data.SelectAttrValue("name", "")

				if defstruct.declare.name == "" {
					return nil, errors.New("invalid 'data->name'")
				}

				defstruct.declare.comment = data.SelectAttrValue("comment", "")
			}
		default:
			{
				return nil, errors.New("invalid 'data->type'")
			}
		}

		defstruct.fields = make([]*FieldAttr, 0)
		if defstruct.isproto {
			defstruct.fields = append(defstruct.fields, &FieldAttr{def: "uint16", name: "Mid"})
			defstruct.fields = append(defstruct.fields, &FieldAttr{def: "uint16", name: "Pid"})
		}

		items := data.SelectElements("item")
		for _, item := range items {
			field := &FieldAttr{}

			field.def = item.SelectAttrValue("type", "")
			if field.def == "" {
				return nil, errors.New("invalid 'item->type'")
			}

			field.array = item.SelectAttrValue("array", "")
			if field.array != "" {
				field.def = "[]" + field.def
			}

			field.name = item.SelectAttrValue("name", "")
			if field.name == "" {
				return nil, errors.New("invalid 'item->name'")
			}

			field.comment = item.SelectAttrValue("comment", "")
			field.alias = item.SelectAttrValue("alias", "")
			field.length = item.SelectAttrValue("length", "")

			defstruct.fields = append(defstruct.fields, field)
		}

		protomoudle.protocols = append(protomoudle.protocols, &defstruct)
	}

	protomoudle.isproto = isAllValid(true, protomoudle.protocols)
	if !protomoudle.isproto && !isAllValid(false, protomoudle.protocols) {
		return nil, errors.New("invalid 'struct' + 'protocol'")
	}

	return &protomoudle, nil
}

func isAllValid(isproto bool, datas []*DefStruct) bool {
	for _, data := range datas {
		if data.isproto != isproto {
			return false
		}
	}

	return true
}

func writeProtocol(defproto *DefProtoMoudle) error {
	file, err := os.OpenFile(
		"./"+strings.ToLower(defproto.moudle.name)+".go",
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
		os.ModePerm,
	)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)
	defer func() {
		writer.Flush()
		file.Close()
	}()

	writer.WriteString("///>本代码由测试工具自动生成,请勿手动修改\n")
	writer.WriteString("package protocols\n\n")
	if defproto.isproto {
		writer.WriteString("import (\n\t\"bytes\"\n\t\"encoding/binary\"\n\t\"fmt\"\n\t\"unsafe\"\n)\n\n")
	} else {
		writer.WriteString("import \"bytes\"\n\n")
	}

	for _, proto := range defproto.protocols {
		writer.WriteString("type ")
		writer.WriteString(firstToUpper(proto.declare.name))
		writer.WriteString(" struct { //>")
		writer.WriteString(proto.declare.comment)
		writer.WriteString("\n")

		fl := 0
		ml := 0
		for _, field := range proto.fields {
			if l := len(field.name); l > fl {
				fl = l
			}
			if l := len(field.def); l > ml {
				ml = l
			}
		}

		for _, field := range proto.fields {
			writer.WriteString("\t")
			s := firstToUpper(field.name)
			l := len(s)
			writer.WriteString(s)
			for ; l < fl; l++ {
				writer.WriteString(" ")
			}
			writer.WriteString(" ")

			writer.WriteString(field.def)
			for i := len(field.def); i < ml; i++ {
				writer.WriteString(" ")
			}

			if field.comment != "" {
				writer.WriteString(" //>")
				writer.WriteString(field.comment)
			}

			writer.WriteString("\n")
		}

		writer.WriteString("}\n\n")
	}

	for k, proto := range defproto.protocols {
		writer.WriteString("func (proto *")
		writer.WriteString(firstToUpper(proto.declare.name))
		writer.WriteString(") GetMid() uint16 {\n\treturn ")
		writer.WriteString(defproto.moudle.moudle)
		writer.WriteString("\n}\n\n")

		writer.WriteString("func (proto *")
		writer.WriteString(firstToUpper(proto.declare.name))
		writer.WriteString(") GetPid() uint16 {\n\treturn ")
		writer.WriteString(strconv.Itoa(k + 1))
		writer.WriteString("\n}\n\n")

		writer.WriteString("func (proto *")
		writer.WriteString(firstToUpper(proto.declare.name))
		writer.WriteString(") Write(b *bytes.Buffer) bool {\n")

		for _, field := range proto.fields {
			if isBuiltinInteger(field.def) {
				if field.name == "Mid" {
					writer.WriteString("\tif !ProtoWriteInteger(b, uint16(")
					writer.WriteString(defproto.moudle.moudle)
					writer.WriteString(")")
				} else if field.name == "Pid" {
					writer.WriteString("\tif !ProtoWriteInteger(b, uint16(")
					writer.WriteString(strconv.Itoa(k + 1))
					writer.WriteString(")")
				} else {
					writer.WriteString("\tif !ProtoWriteInteger(b, proto.")
					writer.WriteString(firstToUpper(field.name))
				}
			} else if isBuiltinIntegerArray(field.def) {
				writer.WriteString("\tif !ProtoWriteIntegerArray(b, proto.")
				writer.WriteString(firstToUpper(field.name))
				writer.WriteString(", ")
				if field.length != "" {
					writer.WriteString(field.length)
				} else if field.array != "" {
					s := getBuiltinIntegerMax(field.array)
					if s == "" {
						return errors.New("illegal array:" + field.array)
					}
					writer.WriteString(field.array)
					writer.WriteString("(")
					writer.WriteString(s)
					writer.WriteString(")")
				} else {
					return errors.New("system error 3")
				}
			} else if field.def[0:2] == "[]" {
				writer.WriteString("\tif !ProtoWriteCustomArray(b, proto.")
				writer.WriteString(firstToUpper(field.name))
				writer.WriteString(", ")
				if field.array != "" {
					s := getBuiltinIntegerMax(field.array)
					if s == "" {
						return errors.New("illegal array:" + field.array)
					}
					writer.WriteString(field.array)
					writer.WriteString("(")
					writer.WriteString(s)
					writer.WriteString(")")
				} else {
					return errors.New("system error 4")
				}
			} else if field.def == "string" {
				writer.WriteString("\tif !ProtoWriteString(b, proto.")
				writer.WriteString(firstToUpper(field.name))
				writer.WriteString(", ")
				if field.length != "" {
					writer.WriteString(field.length)
				} else {
					writer.WriteString("255")
				}
			} else {
				writer.WriteString("\tif !ProtoWriteCustom(b, &proto.")
				writer.WriteString(firstToUpper(field.name))
			}
			writer.WriteString(") {\n\t\treturn false\n\t}\n\n")
		}
		writer.WriteString("\treturn true\n}\n\n")

		writer.WriteString("func (proto *")
		writer.WriteString(firstToUpper(proto.declare.name))
		writer.WriteString(") Read(b *bytes.Buffer) bool {\n")

		for _, field := range proto.fields {
			if isBuiltinInteger(field.def) {
				writer.WriteString("\tif !ProtoReadInteger(b, &proto.")
				writer.WriteString(firstToUpper(field.name))
			} else if isBuiltinIntegerArray(field.def) {
				writer.WriteString("\tif !ProtoReadIntegerArray(b, &proto.")
				writer.WriteString(firstToUpper(field.name))
				writer.WriteString(", ")
				if field.length != "" {
					writer.WriteString(field.length)
				} else if field.array != "" {
					s := getBuiltinIntegerMax(field.array)
					if s == "" {
						return errors.New("illegal array:" + field.array)
					}
					writer.WriteString(field.array)
					writer.WriteString("(")
					writer.WriteString(s)
					writer.WriteString(")")
				} else {
					return errors.New("system error 3")
				}
			} else if field.def[0:2] == "[]" {
				writer.WriteString("\tif !ProtoReadCustomArray(b, &proto.")
				writer.WriteString(firstToUpper(field.name))
				writer.WriteString(", ")
				if field.array != "" {
					s := getBuiltinIntegerMax(field.array)
					if s == "" {
						return errors.New("illegal array:" + field.array)
					}
					writer.WriteString(field.array)
					writer.WriteString("(")
					writer.WriteString(s)
					writer.WriteString(")")
				} else {
					return errors.New("system error 4")
				}
			} else if field.def == "string" {
				writer.WriteString("\tif !ProtoReadString(b, &proto.")
				writer.WriteString(firstToUpper(field.name))
				writer.WriteString(", ")
				if field.length != "" {
					writer.WriteString(field.length)
				} else {
					writer.WriteString("255")
				}
			} else {
				writer.WriteString("\tif !ProtoReadCustom(b, &proto.")
				writer.WriteString(firstToUpper(field.name))
			}
			writer.WriteString(") {\n\t\treturn false\n\t}\n\n")
		}
		writer.WriteString("\treturn true\n}\n\n")
	}

	if !defproto.isproto {
		return nil
	}

	for _, proto := range defproto.protocols {
		writer.WriteString("type I")
		writer.WriteString(firstToUpper(proto.declare.name))
		writer.WriteString(" interface {\n")
		writer.WriteString("\tOn")
		writer.WriteString(firstToUpper(proto.declare.name))
		writer.WriteString("(proto *")
		writer.WriteString(firstToUpper(proto.declare.name))
		writer.WriteString(")\n}\n\n")
	}

	writer.WriteString("type ")
	writer.WriteString(firstToUpper(defproto.moudle.name))
	writer.WriteString(" struct {\n\tprotoDispatch interface{}\n}\n\n")

	writer.WriteString("func New")
	writer.WriteString(firstToUpper(defproto.moudle.name))
	writer.WriteString("[T any](dispatch *T) *")
	writer.WriteString(firstToUpper(defproto.moudle.name))
	writer.WriteString(" {\n\treturn &")
	writer.WriteString(firstToUpper(defproto.moudle.name))
	writer.WriteString("{dispatch}\n}\n\n")

	writer.WriteString("func (protos *")
	writer.WriteString(firstToUpper(defproto.moudle.name))
	writer.WriteString(") GetMid() uint16 {\n\treturn ")
	writer.WriteString(defproto.moudle.moudle)
	writer.WriteString("\n}\n\n")

	writer.WriteString("func (protos *")
	writer.WriteString(firstToUpper(defproto.moudle.name))
	writer.WriteString(") DispatchProto(data []byte) bool {\n")
	writer.WriteString("\tb := bytes.NewBuffer(data)\n\n")
	writer.WriteString("\tmid := binary.LittleEndian.Uint16(data)\n")
	writer.WriteString("\tif mid != protos.GetMid() {\n\t\treturn false\n\t}\n\n")
	writer.WriteString("\tpid := binary.LittleEndian.Uint16(data[unsafe.Sizeof(uint16(0)):])\n")
	writer.WriteString("\tswitch pid {\n")

	for k, proto := range defproto.protocols {
		writer.WriteString("\tcase ")
		writer.WriteString(strconv.Itoa(k + 1))
		writer.WriteString(":\n\t\t{\n")
		writer.WriteString("\t\t\tt, ok := protos.protoDispatch.(I")
		writer.WriteString(firstToUpper(proto.declare.name))
		writer.WriteString(")\n\t\t\tif !ok {\n\t\t\t\treturn false\n\t\t\t}\n\n")
		writer.WriteString("\t\t\tproto := &")
		writer.WriteString(firstToUpper(proto.declare.name))
		writer.WriteString("{}\n\t\t\tif !proto.Read(b) {\n")
		writer.WriteString("\t\t\t\tfmt.Println(\"read ")
		writer.WriteString(firstToUpper(proto.declare.name))
		writer.WriteString(" fail, system error.\")\n\t\t\t\treturn false\n\t\t\t}\n\n\t\t\tt.On")
		writer.WriteString(firstToUpper(proto.declare.name))
		writer.WriteString("(proto)\n")
		writer.WriteString("\t\t}\n")
	}

	writer.WriteString("\tdefault:\n\t\t{\n")
	writer.WriteString("\t\t\tfmt.Println(\"illegal protocol, Mid =\", mid, \"Pid =\", pid)\n")
	writer.WriteString("\t\t}\n\t}\n\n\treturn true\n}")

	return nil
}

func TestProtocol(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	paths := []string{
		"E:/mt2020/mh_trunk/common/protocol/StructDef.xml",
		"E:/mt2020/mh_trunk/common/protocol/Global.xml",
		"E:/mt2020/mh_trunk/common/protocol/ClientWS.xml",
		"E:/mt2020/mh_trunk/common/protocol/ClientGS.xml",
		"E:/mt2020/mh_trunk/common/protocol/ClientCS.xml",
	}

	waitgroup := sync.WaitGroup{}

	for _, path := range paths {
		waitgroup.Add(1)

		go func(path string) {
			protodef, err := parseProtocol(path)
			if err != nil {
				t.Error(err)
				return
			}

			err = writeProtocol(protodef)
			if err != nil {
				t.Error(err)
				return
			}

			waitgroup.Done()
		}(path)
	}

	waitgroup.Wait()
}
