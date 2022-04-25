///>本代码由自动化工具批量生成
package protocols

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unsafe"
)

type KeepAliveReq struct { //>心跳请求服务器主动发给客户端
	Mid uint16
	Pid uint16
	Seq uint32 //>自增长序号
}

type KeepAliveAck struct { //>心跳回应客户端回应服务器
	Mid  uint16
	Pid  uint16
	Seq  uint32 //>自增长序号
	Tick uint32 //>回应tick值
}

type AttrNtf struct { //>属性通知
	Mid   uint16    
	Pid   uint16    
	Guid  uint64     //>角色guid
	Attrs []AttrData //>属性数据
}

type PlayerAppearNtf struct { //>玩家出现通知
	Mid    uint16    
	Pid    uint16    
	Guid   uint64     //>玩家guid
	Sn     int32      //>用户序列号
	X      uint16     //>x坐标
	Y      uint16     //>y坐标
	Attrs  []AttrData //>属性数据
	Buffs  []BuffData //>buff数据
	Custom CustomData //>自定义数据
	Name   string     //>玩家名字
}

type NPCAppearNtf struct { //>NPC出现通知
	Mid         uint16    
	Pid         uint16    
	Guid        uint64     //>NPC guid
	Template_id uint16     //>模板id
	X           uint16     //>x坐标
	Y           uint16     //>y坐标
	Dir         uint8      //>方向
	Attrs       []AttrData //>属性数据
	Custom      CustomData //>自定义数据
	Name        string     //>玩家名字
}

type ItemAppearNtf struct { //>道具出现通知
	Mid         uint16       
	Pid         uint16       
	Guid        uint64        //>道具 guid
	Template_id uint16        //>模板id
	X           uint16        //>x坐标
	Y           uint16        //>y坐标
	Custom      CustomData    //>自定义数据
	Name        string        //>道具名字
	Dyn_attrs   []DynAttrData //>动态属性数据
}

type ObjDisAppearNtf struct { //>对象消失通知
	Mid  uint16
	Pid  uint16
	Guid uint64 //>对象 guid
}

type ObjMoveNtf struct { //>对象移动通知
	Mid   uint16
	Pid   uint16
	Guid  uint64 //>对象 guid
	Dst_x uint16 //>目标点x坐标
	Dst_y uint16 //>目标点y坐标
	Type  uint8  //>移动方式 1=正常 2=跳跃
}

type EnterMapNtf struct { //>角色进入地图通知
	Mid         uint16
	Pid         uint16
	Template_id uint16 //>模板id
	Dst_x       uint16 //>目标点x坐标
	Dst_y       uint16 //>目标点y坐标
}

type MoveReq struct { //>移动请求
	Mid   uint16
	Pid   uint16
	Dst_x uint16 //>目标点x
	Dst_y uint16 //>目标点y
	Cur_x uint16 //>当前点x
	Cur_y uint16 //>当前点y
	Idx   uint32 //>客户端序列号
	Tick  uint32 //>gate收到此包的tick
}

type MoveAck struct { //>移动回应
	Mid     uint16
	Pid     uint16
	Errcode int32  //>0=成功, 其他表示错误码
	Dst_x   uint16 //>当前的坐标点x
	Dst_y   uint16 //>当前的坐标点y
	Idx     uint32 //>客户端序列号
}

type JumpMapReq struct { //>地图跳转请求
	Mid uint16
	Pid uint16
	Idx uint32 //>跳转区索引 静态跳转区idx 小于 65535 动态跳转区 idx 大于 65535 
}

type JumpMapAck struct { //>地图跳转应答
	Mid     uint16
	Pid     uint16
	Errcode int32  //>0=成功, 其他表示错误码
}

type AddJumpMapRegionNtf struct { //>添加地图跳转区域通知
	Mid         uint16   
	Pid         uint16   
	Idx         uint32    //>跳转区索引 静态跳转区idx 小于 65535 动态跳转区 idx 大于 65535 
	Jump_region MapRegion //>起跳区域
}

type DelJumpMapRegionNtf struct { //>删除地图跳转区域通知
	Mid uint16
	Pid uint16
	Idx uint32 //>跳转区索引
}

type ItemAddNtf struct { //>新增物品通知
	Mid            uint16    
	Pid            uint16    
	Guid           uint64     //>容器属于谁(玩家、宠物)
	Container_type uint16     //>item容器类型
	Items          []ItemData //>物品列表
}

type ItemUpdateNtf struct { //>物品更新通知
	Mid            uint16             
	Pid            uint16             
	Guid           uint64              //>容器属于谁(玩家、宠物)
	Container_type uint16              //>item容器类型
	Attrs          []ItemAttrValueList //>变更属性列表
}

type ItemDestroyNtf struct { //>物品销毁通知
	Mid            uint16  
	Pid            uint16  
	Guid           uint64   //>容器属于谁(玩家、宠物)
	Container_type uint16   //>item容器类型
	Item_guids     []uint64 //>物品GUID列表
}

type TipsMsgNtf struct { //>客户端左上角消息通知
	Mid uint16 
	Pid uint16 
	Msg []uint8 //>消息内容
}

type TopMsgNtf struct { //>客户端顶部消息通知（全服公告）
	Mid        uint16 
	Pid        uint16 
	Foreground uint8   //>消息前景颜色ID
	Background uint8   //>消息背景颜色ID
	Count      uint8   //>消息滚动次数
	Msg        []uint8 //>消息内容
}

type SysMsgNtf struct { //>客户端底部消息通知（普通聊天窗口）
	Mid        uint16 
	Pid        uint16 
	Foreground uint8   //>消息前景颜色ID
	Background uint8   //>消息背景颜色ID
	Msg        []uint8 //>消息内容
}

type PopupMsgNtf struct { //>客户端弹框消息
	Mid  uint16 
	Pid  uint16 
	Type uint8   //>消息类型 0:通用 1:任务
	Msg  []uint8 //>消息内容
}

type ItemContainerNtf struct { //>上线物品容器通知
	Mid            uint16    
	Pid            uint16    
	Guid           uint64     //>容器属于谁(玩家、宠物)
	Container_type uint16     //>item容器类型
	Capacity       uint16     //>item容器容量
	Items          []ItemData //>物品列表
}

type ItemContainerUpdateNtf struct { //>物品容器更新通知
	Mid            uint16
	Pid            uint16
	Guid           uint64 //>容器属于谁(玩家、宠物)
	Container_type uint16 //>item容器类型
	Capacity       uint16 //>item容器容量
}

type SubmitForm struct { //>请求执行表单
	Mid  uint16     
	Pid  uint16     
	Form string      //>表单table
	Func string      //>表单函数
	Args []ParamData //>脚本命令参数
}

type ShowFormNtf struct { //>表单通知
	Mid        uint16 
	Pid        uint16 
	Form       string  //>表单名称
	Compressed uint8   //>是否压缩：1-压缩，0-未压缩
	Context    []uint8 //>表单内容
}

type ExecuteGMReq struct { //>请求执行GM命令
	Mid uint16
	Pid uint16
	Cmd string //>GM命令
	Arg string //>GM命令参数
}

type FightBeginNtf struct { //>战斗开始通知
	Mid        uint16          
	Pid        uint16          
	Groups     []FightGroupData //>战斗组信息
	Self_group uint8            //>自己所属的组
	Is_pvp     uint8            //>是否pvp战斗
}

type TurnRoundNtf struct { //>新回合通知
	Mid   uint16
	Pid   uint16
	Round uint16 //>回合数
}

type FightOperateListNtf struct { //>战斗指令通知
	Mid             uint16 
	Pid             uint16 
	Player_operates []uint8 //>玩家允许操作的指令列表
	Pet_operates    []uint8 //>宠物允许操作的指令列表
	Count_down      uint16  //>倒计时 单位:毫秒
}

type FightOperateReq struct { //>战斗指令请求
	Mid     uint16          
	Pid     uint16          
	Operate FightOperateData //>操作数据
}

type FightOperateAck struct { //>战斗指令返回结果
	Mid     uint16
	Pid     uint16
	Errcode int32  //>错误码
	Errmsg  string //>错误描述
}

type FightOperateNtf struct { //>战斗指令完成通知
	Mid        uint16
	Pid        uint16
	Fighter_id uint16 //>战斗者id
}

type FightDisplayNtf struct { //>战斗显示用数据
	Mid    uint16 
	Pid    uint16 
	Data   []uint8 //>显示数据
	Crypto []uint8 //>加密串,显示完成时原样返回
}

type FightDisplayCompleteNtf struct { //>战斗显示完成知通
	Mid    uint16 
	Pid    uint16 
	Crypto []uint8 //>加密串
}

type FightAutoReq struct { //>自动战斗请求
	Mid     uint16
	Pid     uint16
	Is_auto uint8  //>0:关闭 1:开启
}

type FightAutoAck struct { //>自动战斗回应
	Mid     uint16
	Pid     uint16
	Is_auto uint8  //>0:关闭 1:开启
	Errcode int32  //>错误码
	Errmsg  string //>错误描述
}

type FightAutoNtf struct { //>自动战斗通知
	Mid        uint16
	Pid        uint16
	Fighter_id uint16 //>战斗者id
	Is_auto    uint8  //>0:关闭 1:开启
}

type FightAutoSkillReq struct { //>自动战斗技能请求
	Mid     uint16
	Pid     uint16
	Actor   uint64 //>发起者guid 玩家自己:填0 宠物:填宠物guid
	Skillid uint16 //>技能id
}

type FightAutoSkillAck struct { //>自动战斗技能回应
	Mid     uint16
	Pid     uint16
	Actor   uint64 //>发起者guid 玩家自己:填0 宠物:填宠物guid
	Skillid uint16 //>技能id
	Errcode int32  //>错误码
	Errmsg  string //>错误描述
}

type FightEndNtf struct { //>战斗结束通知
	Mid    uint16
	Pid    uint16
	Result uint8  //>战斗结果
}

type AddFighterNtf struct { //>战斗中添加战斗者
	Mid   uint16     
	Pid   uint16     
	Data  FighterData //>战斗者信息
	Group uint8       //>自己所属的组
}

type DelFighterNtf struct { //>战斗中实时删除战斗者(退出观战时用)
	Mid        uint16
	Pid        uint16
	Fighter_id uint16 //>战斗者id
}

type AddFightPetData struct { //>添加战斗宠物数据
	Mid uint16      
	Pid uint16      
	Pet FightPetData //>战斗宠物数据
}

type UpdateFightPetData struct { //>修改战斗宠物数据
	Mid uint16      
	Pid uint16      
	Pet FightPetData //>战斗宠物数据
}

type PlayerKillReq struct { //>攻击玩家请求
	Mid    uint16
	Pid    uint16
	Target uint64 //>玩家guid
}

type PlayerKillAck struct { //>攻击玩家回应
	Mid     uint16
	Pid     uint16
	Target  uint64 //>玩家guid
	Errcode int32  //>错误码
	Errmsg  string //>错误描述
}

type EnterFightViewReq struct { //>进入观战请求
	Mid    uint16
	Pid    uint16
	Target uint64 //>观战目标guid
}

type EnterFightViewAck struct { //>进入观战回应
	Mid     uint16
	Pid     uint16
	Target  uint64 //>观战目标guid
	Errcode int32  //>错误码
	Errmsg  string //>错误描述
}

type LeaveFightViewReq struct { //>退出观战请求
	Mid uint16
	Pid uint16
}

type LeaveFightViewAck struct { //>退出观战回应
	Mid     uint16
	Pid     uint16
	Errcode int32  //>错误码
	Errmsg  string //>错误描述
}

type TeamLeaderOprReq struct { //>队长操作请求
	Mid uint16
	Pid uint16
	Opr uint8  //>队长操作
}

type TeamNtf struct { //>队伍通知
	Mid               uint16      
	Pid               uint16      
	Team_guid         uint64       //>队伍GUID
	Target            uint16       //>目标ID
	Permission        uint16       //>1=开启队友招人权限 0=关闭
	Min_require_level uint16       //>最低要求等级
	Max_require_level uint16       //>最高要求等级
	Leader_guid       uint64       //>队长GUID
	Leader_guards     []uint64     //>队长侍从GUID列表
	Members           []MemberData //>队伍成员
	Min_require_reinc uint16       //>最低要求转生次数
	Max_require_reinc uint16       //>最高要求转生次数
	Customs           CustomData   //>自定义
}

type TeamLeaderNtf struct { //>队伍队长信息通知
	Mid           uint16  
	Pid           uint16  
	Leader_guid   uint64   //>队长GUID
	Leader_guards []uint64 //>队长侍从GUID列表
}

type TeamDestroyNtf struct { //>队伍解散通知
	Mid       uint16
	Pid       uint16
	Team_guid uint64 //>队伍GUID
}

type TeamMemberNtf struct { //>成员信息刷新通知
	Mid    uint16    
	Pid    uint16    
	Member MemberData //>队伍成员信息
}

type TeamMemberLeaveNtf struct { //>成员离开通知
	Mid         uint16
	Pid         uint16
	Action      uint8  //>1=被踢，2=主动退出, 3=暂离, 4=归队
	Player_guid uint64 //>成员guid
}

type NpcSelectReq struct { //>NPC功能选择请求
	Mid       uint16
	Pid       uint16
	Npc_guid  uint64 //>NPC guid
	Select_id uint16 //>选择项id 0表示获取初始选项
}

type NpcSelectAck struct { //>NPC功能选择应答
	Mid      uint16      
	Pid      uint16      
	Npc_guid uint64       //>NPC guid
	Options  []OptionData //>NPC选项
	Msg      string       //>回应消息,没有就从npc表中读取
}

type NpcTalkReq struct { //>NPC对话请求
	Mid        uint16 
	Pid        uint16 
	Npc_guid   uint64  //>NPC guid
	Compressed uint8   //>是否压缩：1-压缩，0-未压缩
	Talk       []uint8 //>对话内容
}

type NpcTalkAck struct { //>NPC对话应答
	Mid        uint16 
	Pid        uint16 
	Npc_guid   uint64  //>NPC guid
	Compressed uint8   //>是否压缩：1-压缩，0-未压缩
	Talk       []uint8 //>对话内容
}

type InviteMsgNtf struct { //>邀请消息
	Mid          uint16 
	Pid          uint16 
	Inviter_guid uint64  //>邀请人guid
	Type         uint64  //>邀请类型
	Msg          []uint8 //>邀请内容
	Interval     uint32  //>邀请有效时长：秒
}

type ReplyInvite struct { //>回复邀请
	Mid          uint16
	Pid          uint16
	Inviter_guid uint64 //>邀请人guid
	Type         uint32 //>邀请类型
	Agreed       uint8  //>1:表示同意邀请 0:表示不同意
}

type MoveItem struct { //>移动物品
	Mid            uint16
	Pid            uint16
	Item_guid      uint64 //>物品guid
	Container_type uint16 //>item容器类型
}

type UseItem struct { //>回复邀请
	Mid       uint16
	Pid       uint16
	Item_guid uint64 //>物品guid
}

type RearrangeItem struct { //>回复邀请
	Mid            uint16
	Pid            uint16
	Container_type uint16 //>item容器类型
}

type SkillContainerNtf struct { //>上线技能容器通知
	Mid    uint16     
	Pid    uint16     
	Guid   uint64      //>技能容器属于谁(玩家、宠物)
	Skills []SkillData //>技能列表
}

type AddSkillNtf struct { //>添加技能通知
	Mid   uint16   
	Pid   uint16   
	Guid  uint64    //>技能属于谁(玩家、宠物)
	Skill SkillData //>技能数据
}

type UpdateSkillNtf struct { //>添加技能通知
	Mid   uint16
	Pid   uint16
	Guid  uint64 //>技能属于谁(玩家、宠物)
	Id    uint16 //>技能id
	Attr  uint16 //>技能属性名 1:当前熟练度 2:最大熟练度 3:是否激活
	Value uint32 //>技能属性值
}

type DelSkillNtf struct { //>删除技能通知
	Mid  uint16
	Pid  uint16
	Guid uint64 //>技能属于谁(玩家、宠物)
	Id   uint16 //>技能id
}

type PetAppearNtf struct { //>宠物出现通知
	Mid    uint16    
	Pid    uint16    
	Guid   uint64     //>宠物guid
	Name   string     //>宠物名字
	X      uint16     //>x坐标
	Y      uint16     //>y坐标
	Attrs  []AttrData //>属性数据
	Custom CustomData //>自定义数据
}

type PetContainerNtf struct { //>上线宠物容器通知
	Mid            uint16   
	Pid            uint16   
	Container_type uint16    //>pet容器类型
	Capacity       uint16    //>pet容器容量
	Pets           []PetData //>宠物列表
}

type PetContainerUpdateNtf struct { //>宠物容器更新通知
	Mid            uint16
	Pid            uint16
	Container_type uint16 //>pet容器类型
	Capacity       uint16 //>pet容器容量
}

type PetAddNtf struct { //>新增宠物通知
	Mid            uint16 
	Pid            uint16 
	Container_type uint16  //>pet容器类型
	Pet            PetData //>物品列表
}

type PetDestroyNtf struct { //>宠物销毁通知
	Mid            uint16
	Pid            uint16
	Container_type uint16 //>pet容器类型
	Guid           uint64 //>宠物guid
}

type SetPetLineup struct { //>设置宠物上阵请求
	Mid    uint16
	Pid    uint16
	Guid   uint64 //>宠物guid
	Lineup uint8  //>上阵：0 下阵，1 上阵
}

type ShowPet struct { //>设置宠物显示请求
	Mid  uint16
	Pid  uint16
	Guid uint64 //>宠物guid
	Show uint8  //>显示：0 隐藏，1 显示
}

type ReleasePet struct { //>放生宠物请求
	Mid  uint16
	Pid  uint16
	Guid uint64 //>宠物guid
}

type MovePet struct { //>移动宠物
	Mid            uint16
	Pid            uint16
	Guid           uint64 //>宠物guid
	Container_type uint16 //>pet容器类型
}

type ShopOpenNtf struct { //>商店打开通知
	Mid            uint16        
	Pid            uint16        
	Shop_id        uint16         //>商店id
	Def_item_id    uint16         //>默认选中商品id
	Type           uint8          //>商店类型0=宠物商店 1=道具商店 2=随身商店
	Shop_item_list []ShopItemData //>商品列表
	Buy_back_list  []ItemData     //>回购列表
}

type ShopBuyNtf struct { //>商店购买通知
	Mid           uint16
	Pid           uint16
	Shop_id       uint16 //>商店id
	Shop_item_id  uint16 //>商品id
	Shop_item_num uint16 //>商品数量
}

type SellNtf struct { //>商店出售通知
	Mid        uint16        
	Pid        uint16        
	Sell_items []SellItemData //>出售物品列表
}

type BuyBackNtf struct { //>商店回购通知
	Mid      uint16
	Pid      uint16
	Buy_guid uint64 //>出售物品guid
}

type BuyBackListNtf struct { //>商店回购通知
	Mid           uint16    
	Pid           uint16    
	Buy_back_list []ItemData //>回购列表
}

type TeamAttrNtf struct { //>队伍成员属性通知
	Mid   uint16    
	Pid   uint16    
	Guid  uint64     //>角色guid
	Attrs []AttrData //>属性数据
}

type TipsMsgExNtf struct { //>客户端飘字提示+左下角通知
	Mid uint16 
	Pid uint16 
	Msg []uint8 //>消息内容
}

type ItemNewAddNtf struct { //>物品新增通知
	Mid       uint16
	Pid       uint16
	Item_guid uint64 //>物品GUID
}

type QuestContainerNtf struct { //>任务容器通知
	Mid    uint16 
	Pid    uint16 
	Buffer []uint8 //>任务列表: QuestClientData[]
}

type QuestAddNtf struct { //>新增任务通知
	Mid    uint16 
	Pid    uint16 
	Buffer []uint8 //>任务列表: QuestClientData
}

type QuestStateNtf struct { //>任务状态变更通知
	Mid          uint16
	Pid          uint16
	Quest_id     uint32 //>任务ID
	Quest_state  uint8  //>任务状态
	Changed_time uint32 //>状态改变时间
}

type QuestNpcStateNtf struct { //>任务Npc状态通知
	Mid    uint16
	Pid    uint16
	Npc_id uint32 //>npc id
	State  uint8  //>任务状态
	Logo   uint64 //>任务logo
}

type QuestTrackCountNtf struct { //>任务追踪目标数量通知
	Mid         uint16
	Pid         uint16
	Quest_id    uint32 //>任务ID
	Track_id    uint32 //>追踪目标ID
	Track_count uint32 //>追踪数量
}

type QuestTalkSelectReq struct { //>NPC功能选择应答
	Mid       uint16
	Pid       uint16
	Npc_guid  uint64 //>NPC guid
	Quest_id  uint32 //>任务ID
	Option_id uint32 //>对话选项ID
}

type QuestTalkSelectAck struct { //>NPC功能选择应答
	Mid      uint16 
	Pid      uint16 
	Npc_guid uint64  //>NPC guid
	Quest_id uint32  //>任务ID
	Talk     []uint8 //>npc对话压缩: QuestTalkData[]
}

type ObjStrCustomNtf struct { //>字符串自定义变量通知
	Mid    uint16       
	Pid    uint16       
	Guid   uint64        //>角色guid
	Custom CustomStrData //>自定义变量
}

type ObjIntCustomNtf struct { //>整形自定义变量通知
	Mid    uint16       
	Pid    uint16       
	Guid   uint64        //>角色guid
	Custom CustomIntData //>自定义变量
}

type ObjDynAttrNtf struct { //>动态属性通知
	Mid      uint16     
	Pid      uint16     
	Guid     uint64      //>角色guid
	Dyn_attr DynAttrData //>动态属性
}

type CommitQuestItemRsp struct { //>上交任务物品应答
	Mid      uint16         
	Pid      uint16         
	Quest_id uint32          //>任务ID
	Items    []GUIDCountData //>上交物品
}

type CommitQuestItemNtf struct { //>通知上交任务物品
	Mid      uint16       
	Pid      uint16       
	Quest_id uint32        //>任务ID
	Items    []IDCountData //>上交物品
}

type CommitQuestPetRsp struct { //>上交任务宠物应答
	Mid      uint16         
	Pid      uint16         
	Quest_id uint32          //>任务ID
	Pets     []GUIDCountData //>上交宠物
}

type CommitQuestPetNtf struct { //>通知上交任务宠物
	Mid      uint16       
	Pid      uint16       
	Quest_id uint32        //>任务ID
	Pets     []IDCountData //>上交宠物
}

type AbandonQuest struct { //>放弃任务
	Mid      uint16
	Pid      uint16
	Quest_id uint32 //>任务ID
}

type SkillTipsReq struct { //>技能tips请求
	Mid      uint16
	Pid      uint16
	Actor    uint64 //>宠物guid,自己填0
	Skill_id uint16 //>技能ID
}

type SkillTipsAck struct { //>技能tips回应
	Mid      uint16         
	Pid      uint16         
	Skill_id uint16          //>技能ID
	Tips     []SkillTipsData //>tip属性
	Errcode  int32           //>错误码
	Errmsg   string          //>错误描述
}

type TeamLeaderOprNtf struct { //>队长操作通知
	Mid uint16
	Pid uint16
	Opr uint8  //>队长操作
}

type KeepAliveCtrlNtf struct { //>心跳控制通知
	Mid uint16
	Pid uint16
	Opt uint8  //>0=关闭心跳 1=开启心跳
}

type NetDelayReq struct { //>网络延时探测请求
	Mid     uint16
	Pid     uint16
	Gs_tick uint32 //>GS发送此请求tick
	Gt_tick uint32 //>GT发送此请求tick
	Sg_tick uint32 //>SG发送此请求tick
	Ping    uint8  //>是否ping 1=开始 0=结束
}

type NetDelayAck struct { //>网络延时探测请求
	Mid          uint16
	Pid          uint16
	Gs_elpase    int32  //>GS从发出NetDelayReq到收到NetDelayAck的耗时
	Gt_elpase    int32  //>GT从发出NetDelayReq到收到NetDelayAck的耗时
	Sg_elpase    int32  //>SG从发出NetDelayReq到收到NetDelayAck的耗时
	Gt_cpu       int32  //>GT收到NetDelayAck时的cpu利用率
	Baidu_delay  int32  //>百度的延时
	Taobao_delay int32  //>淘宝的延时
	Sg_id        string //>云网关id
}

type SystemSetupNtf struct { //>系统设置通知
	Mid   uint16    
	Pid   uint16    
	Attrs []AttrData //>属性数据
}

type BuffListNtf struct { //>buff列表通知
	Mid   uint16    
	Pid   uint16    
	Guid  uint64     //>玩家或宠物guid
	Buffs []BuffData //>BUFF列表
}

type BuffAddNtf struct { //>新增buff通知
	Mid  uint16  
	Pid  uint16  
	Guid uint64   //>玩家或宠物guid
	Buff BuffData //>BUFF列表
}

type BuffDurationNtf struct { //>Buff持续数值改变通知
	Mid      uint16
	Pid      uint16
	Guid     uint64 //>玩家或宠物guid
	Id       uint16 //>buff的id
	Duration uint32 //>持续数值
}

type BuffDynAttrNtf struct { //>Buff动态属性改变通知
	Mid       uint16       
	Pid       uint16       
	Guid      uint64        //>玩家或宠物guid
	Id        uint16        //>buff的id
	Dyn_attrs []DynAttrData //>动态属性数据
}

type StopBuff struct { //>客户通知服务器停用buff
	Mid  uint16
	Pid  uint16
	Guid uint64 //>玩家或宠物guid,玩家的填0
	Id   uint16 //>buff的id
}

type BuffDelNtf struct { //>删除buff通知
	Mid  uint16
	Pid  uint16
	Guid uint64 //>玩家或宠物guid
	Id   uint16 //>buff的id
}

type ActivityListReq struct { //>活动列表请求
	Mid uint16
	Pid uint16
}

type ActivityListAck struct { //>活动列表回应
	Mid     uint16        
	Pid     uint16        
	Datas   []ActivityData //>活动数据
	Errcode int32          //>错误码
	Errmsg  string         //>错误描述
}

type ActivityJoinNtf struct { //>加入活动通知
	Mid uint16
	Pid uint16
	Id  uint16 //>活动的id
}

type GuildListNtf struct { //>帮派列表通知
	Mid    uint16          
	Pid    uint16          
	Pos    int32            //>帮派索引
	Count  uint16           //>帮派总数
	Guilds []GuildBriefData //>帮派简略信息
}

type GuildNtf struct { //>帮派信息通知
	Mid   uint16         
	Pid   uint16         
	Guild GuildData       //>帮派数据
	Self  GuildMemberData //>帮派成员信息
}

type GuildBaseNtf struct { //>帮派基本信息，上线通知
	Mid  uint16
	Pid  uint16
	Guid uint64 //>帮派guid
	Name string //>帮派名称
}

type GuildNoticeNtf struct { //>帮派通知
	Mid    uint16
	Pid    uint16
	Guid   uint64 //>帮派guid
	Notice string //>通知内容
}

type GuildBriefNtf struct { //>帮派成员信息
	Mid   uint16        
	Pid   uint16        
	Guild GuildBriefData //>帮派简略信息
}

type GuildMemberListNtf struct { //>帮派列表通知
	Mid     uint16                
	Pid     uint16                
	Guid    uint64                 //>帮派GUID
	Pos     int32                  //>成员索引
	Count   uint16                 //>帮派成员总数
	Members []GuildMemberBriefData //>帮派成员数据
}

type GuildMemberNtf struct { //>帮派成员信息
	Mid     uint16           
	Pid     uint16           
	Guid    uint64            //>帮派GUID
	Members []GuildMemberData //>帮派成员数据
}

type GuildApplicantListNtf struct { //>帮派列表通知
	Mid        uint16              
	Pid        uint16              
	Guid       uint64               //>帮派GUID
	Applicants []GuildApplicantData //>帮派申请列表
}

type GuildOperationNtf struct { //>帮派操作通知
	Mid     uint16
	Pid     uint16
	Type    uint16 //>操作通知
	Param1  string //>参数1
	Param2  string //>参数2
	Errcode int32  //>0=成功, 其他表示错误码
	Errmsg  string //>错误码不为0时表示 错误消息
}

type GetTimestampReq struct { //>服务器时间戳请求
	Mid uint16
	Pid uint16
}

type GetTimestampAck struct { //>服务器时间戳应答
	Mid uint16
	Pid uint16
	Now uint32 //>当前时间
}

type ContactListNtf struct { //>联系人列表
	Mid          uint16       
	Pid          uint16       
	Contact_type uint8         //>联系人类型：对应枚举contact_type的陌生人、好友、申请者
	Contacts     []ContactData //>联系人列表
}

type AddContactNtf struct { //>新增联系人通知
	Mid     uint16     
	Pid     uint16     
	Contact ContactData //>联系人
}

type UpdateContactNtf struct { //>联系人刷新通知
	Mid     uint16     
	Pid     uint16     
	Contact ContactData //>联系人
}

type DelContactNtf struct { //>删除联系人通知
	Mid         uint16
	Pid         uint16
	Target_guid uint64 //>玩家GUID
}

type AddContactMessageNtf struct { //>联系人消息通知
	Mid      uint16     
	Pid      uint16     
	Messages MessageData //>消息内容
}

type ItemQueryNtf struct { //>物品查询通知
	Mid  uint16  
	Pid  uint16  
	Item ItemData //>物品数据
}

type PetQueryNtf struct { //>宠物查询通知
	Mid uint16 
	Pid uint16 
	Pet PetData //>宠物数据
}

type ContactInfoNtf struct { //>联系人信息通知
	Mid  uint16         
	Pid  uint16         
	Data ContactInfoData //>联系人信息
}

type MailListNtf struct { //>邮件列表通知
	Mid   uint16        
	Pid   uint16        
	Mails []MailHeadData //>邮件标题列表
}

type AddMailNtf struct { //>邮件通知
	Mid  uint16      
	Pid  uint16      
	Mail MailHeadData //>邮件标题
}

type DelMailNtf struct { //>邮件通知
	Mid       uint16
	Pid       uint16
	Mail_guid uint64 //>邮件GUID
}

type MailBodyNtf struct { //>邮件内容通知
	Mid       uint16      
	Pid       uint16      
	Mail_guid uint64       //>邮件GUID
	Mail      MailBodyData //>邮件内容
}

type UpdateMailBodyNtf struct { //>邮件内容变更通知
	Mid       uint16      
	Pid       uint16      
	Mail_guid uint64       //>邮件GUID
	Mail      MailBodyData //>邮件内容
}

type UpdateMailHeadNtf struct { //>邮件标题变更通知
	Mid       uint16      
	Pid       uint16      
	Mail_guid uint64       //>邮件GUID
	Mail      MailHeadData //>邮件标题
}

type RanklistReq struct { //>获取排行榜请求
	Mid   uint16
	Pid   uint16
	Type  int32  //>榜单类型 
	Begin uint16 //>获取排行开始
	End   uint16 //>获取排行结束
}

type RanklistAck struct { //>获取排行榜应答
	Mid   uint16        
	Pid   uint16        
	Type  int32          //>榜单类型 1=等级总榜 2=等级人榜 3=等级魔榜 4=等级仙榜 5=等级鬼榜 6=等级龙榜 7=装备总榜 8=装备人榜 9=装备魔榜 10=装备仙榜 11=装备鬼榜 12=装备龙榜 13=帮派榜 14=宠物榜 15=水陆大会榜 16=竞技场榜 16以后是其他自定义榜
	Total uint16         //>排行榜总数
	Begin uint16         //>获取排行开始
	End   uint16         //>获取排行结束
	Data  []RanklistData //>排行榜内容
}

type GetRankReq struct { //>获取排行排位
	Mid  uint16
	Pid  uint16
	Type int32  //>榜单类型 
	Guid uint64 //>玩家宠物或者其他榜单上对象的GUID
}

type GetRankAck struct { //>获取排行排位
	Mid  uint16      
	Pid  uint16      
	Type int32        //>榜单类型 
	Rank uint16       //>玩家宠物或者其他榜单上排名数, 0=未上榜
	Data RanklistData //>榜单数据
}

type TitleContainerNtf struct { //>称号容器通知
	Mid    uint16     
	Pid    uint16     
	Titles []TitleData //>称号列表
}

type TitleAddNtf struct { //>新增称号通知
	Mid   uint16   
	Pid   uint16   
	Title TitleData //>称号数据
}

type TitleDelNtf struct { //>宠物销毁通知
	Mid      uint16
	Pid      uint16
	Title_id uint16 //>称号ID
}

type AgentKeyReq struct { //>代理Key值请求
	Mid uint16
	Pid uint16
}

type AgentKeyAck struct { //>代理Key值返回
	Mid uint16 
	Pid uint16 
	Key []uint8 //>代理Key值
}

type HeadMsgNtf struct { //>头顶消息
	Mid  uint16 
	Pid  uint16 
	Guid uint64  //>对象GUID
	Msg  []uint8 //>消息内容
}

type AutoContainerNtf struct { //>自动战斗信息通知
	Mid        uint16
	Pid        uint16
	Guid       uint64 //>自己或自己的宠物
	Is_auto    uint8  //>是否开启了自动(宠物跟随主人自动)
	Auto_skill uint16 //>自动战斗技能
}

type PlayerQueryNtf struct { //>玩家查询通知
	Mid   uint16     
	Pid   uint16     
	Brief PlayerBrief //>玩家简略数据
}

type UseAllItem struct { //>使用全部物品
	Mid         uint16
	Pid         uint16
	Item_guid   uint64 //>物品guid
	Item_amount uint16 //>物品数量
}

type GuardContainerNtf struct { //>上线侍从容器通知
	Mid    uint16     
	Pid    uint16     
	Guards []GuardData //>侍从
}

type GuardAddNtf struct { //>新增侍从通知
	Mid   uint16   
	Pid   uint16   
	Guard GuardData //>侍从列表
}

type SetGuardLineup struct { //>设置侍从上阵请求
	Mid    uint16
	Pid    uint16
	Guid   uint64 //>侍从guid
	Lineup uint8  //>上阵：0 下阵，1 上阵
}

type PetNewAddNtf struct { //>宠物新增通知
	Mid      uint16
	Pid      uint16
	Pet_guid uint64 //>宠物GUID
}

type TeamPlatformNtf struct { //>组队平台通知
	Mid   uint16         
	Pid   uint16         
	Teams []TeamBriefData //>队伍列表
}

type TeamApplicantsNtf struct { //>队伍申请列表
	Mid        uint16             
	Pid        uint16             
	Team_guid  uint64              //>队伍GUID
	Applicants []TeamApplicantData //>申请人列表
}

type TeamOperationNtf struct { //>队伍操作通知
	Mid     uint16
	Pid     uint16
	Type    uint16 //>操作通知
	Errcode int32  //>0=成功, 其他表示错误码
	Errmsg  string //>错误码不为0时表示 错误消息
}

type TeamTargetNtf struct { //>队伍目标信息变更通知
	Mid               uint16
	Pid               uint16
	Team_guid         uint64 //>队伍GUID
	Target            uint16 //>目标ID
	Permission        uint16 //>1=开启队友招人权限 0=关闭
	Min_require_level uint16 //>最低要求等级
	Max_require_level uint16 //>最高要求等级
	Leader_guid       uint64 //>队长GUID
	Min_require_reinc uint16 //>最低要求转生次数
	Max_require_reinc uint16 //>最高要求转生次数
}

type ChangedNameNtf struct { //>改名通知
	Mid  uint16
	Pid  uint16
	Guid uint64 //>对象 guid
	Name string //>玩家名字
}

type CustomDataNtf struct { //>自定义变量通知
	Mid    uint16    
	Pid    uint16    
	Guid   uint64     //>玩家guid
	Custom CustomData //>自定义数据
}

type SpeedCheckNtf struct { //>游戏速度检查通知
	Mid  uint16
	Pid  uint16
	Type uint8  //>1=加速 2=超时 3=异常
	Pct  uint8  //>加速百分比
}

type ConsoleMsgNtf struct { //>控制台消息通知
	Mid uint16 
	Pid uint16 
	Msg []uint8 //>控制台消息
}

type PetSwapNtf struct { //>宠物交换位置通知
	Mid            uint16
	Pid            uint16
	Container_type uint16 //>pet容器类型
	Pet_guid1      uint64 //>宠物1guid
	Pet_guid2      uint64 //>宠物2guid
}

type GuardDestroyNtf struct { //>侍从遣散通知
	Mid  uint16
	Pid  uint16
	Guid uint64 //>侍从GUID
}

type ActivateGuard struct { //>激活侍从
	Mid      uint16
	Pid      uint16
	Guard_id uint32 //>侍从ID
}

type ReleaseGuard struct { //>遣散侍从
	Mid        uint16
	Pid        uint16
	Guard_guid uint64 //>侍从GUID
}

type TeamMemberSwapNtf struct { //>队伍成员交换位置通知
	Mid          uint16
	Pid          uint16
	Member_guid1 uint64 //>队伍成员GUID1
	Member_guid2 uint64 //>队伍成员GUID2
}

type GuardSwapNtf struct { //>侍从交换位置通知
	Mid         uint16
	Pid         uint16
	Guard_guid1 uint64 //>侍从1guid
	Guard_guid2 uint64 //>侍从2guid
}

type PetReplaceNtf struct { //>宠物替换通知
	Mid            uint16 
	Pid            uint16 
	Container_type uint16  //>pet容器类型
	Dst_pet_guid   uint64  //>目标宠物GUID
	Src_pet        PetData //>替换宠物
}

type GuardAppearNtf struct { //>侍从出现通知
	Mid    uint16    
	Pid    uint16    
	Guid   uint64     //>侍从guid
	Name   string     //>侍从名字
	X      uint16     //>x坐标
	Y      uint16     //>y坐标
	Attrs  []AttrData //>属性数据
	Buffs  []BuffData //>buff数据
	Custom CustomData //>自定义数据
}

type InstructionContainerNtf struct { //>指令容器通知
	Mid   uint16         
	Pid   uint16         
	Datas InstructionData //>指令数据
}

type InstructionAddReq struct { //>指令添加请求
	Mid     uint16
	Pid     uint16
	Type    uint8  //>指令类型
	Content string //>指令内容
}

type InstructionAddAck struct { //>指令添加回应
	Mid     uint16
	Pid     uint16
	Type    uint8  //>指令类型
	Content string //>指令内容
	Errcode int32  //>0=成功, 其他表示错误码
	Errmsg  string //>错误码不为0时表示 错误消息
}

type InstructionDeleteReq struct { //>指令删除请求
	Mid   uint16
	Pid   uint16
	Type  uint8  //>指令类型
	Index uint8  //>指令索引
}

type InstructionDeleteAck struct { //>指令删除回应
	Mid     uint16
	Pid     uint16
	Type    uint8  //>指令类型
	Index   uint8  //>指令索引
	Errcode int32  //>0=成功, 其他表示错误码
	Errmsg  string //>错误码不为0时表示 错误消息
}

type InstructionModfityReq struct { //>指令修改请求
	Mid     uint16
	Pid     uint16
	Type    uint8  //>指令类型
	Index   uint8  //>指令索引
	Content string //>指令内容
}

type InstructionModifyAck struct { //>指令修改回应
	Mid     uint16
	Pid     uint16
	Type    uint8  //>指令类型
	Index   uint8  //>指令索引
	Content string //>指令内容
	Errcode int32  //>0=成功, 其他表示错误码
	Errmsg  string //>错误码不为0时表示 错误消息
}

type InstructionDefaultReq struct { //>指令恢复默认请求
	Mid  uint16
	Pid  uint16
	Type uint8  //>指令类型
}

type InstructionDefaultAck struct { //>指令恢复默认回应
	Mid     uint16            
	Pid     uint16            
	Type    uint8              //>指令类型
	Datas   []InstructionBasic //>恢复后的指令
	Errcode int32              //>0=成功, 其他表示错误码
	Errmsg  string             //>错误码不为0时表示 错误消息
}

type InstructionAttachReq struct { //>指令设置请求
	Mid    uint16
	Pid    uint16
	Type   uint8  //>指令类型
	Index  uint8  //>指令索引
	Target uint64 //>目标GUID
}

type InstructionAttachAck struct { //>指令设置回应
	Mid     uint16
	Pid     uint16
	Type    uint8  //>指令类型
	Index   uint8  //>指令索引
	Target  uint64 //>目标GUID
	Errcode int32  //>0=成功, 其他表示错误码
	Errmsg  string //>错误码不为0时表示 错误消息
}

type InstructionAttachNtf struct { //>指令设置通知
	Mid        uint16
	Pid        uint16
	Fighter_id uint32 //>战斗者id
	Content    string //>指令内容
}

type InstructionDetachReq struct { //>指令重置请求
	Mid    uint16
	Pid    uint16
	Target uint64 //>目标GUID
}

type InstructionDetachAck struct { //>指令重置回应
	Mid     uint16
	Pid     uint16
	Target  uint64 //>目标GUID
	Errcode int32  //>0=成功, 其他表示错误码
	Errmsg  string //>错误码不为0时表示 错误消息
}

type InstructionDetachNtf struct { //>指令重置通知
	Mid        uint16
	Pid        uint16
	Fighter_id uint32 //>战斗者id
}

type PlayerDetailNtf struct { //>玩家信息通知
	Mid  uint16      
	Pid  uint16      
	Data PlayerDetail //>玩家详细信息
}

type MapDynBlockPtNtf struct { //>地图动态阻挡点通知
	Mid           uint16
	Pid           uint16
	Map_id        uint16 //>地图id
	Block_pt_list []Pt   //>玩家详细信息
}

type GuardQueryNtf struct { //>侍从查询通知
	Mid   uint16   
	Pid   uint16   
	Guard GuardData //>侍从数据
}

type BuyBackNtfEx struct { //>商店回购通知
	Mid        uint16
	Pid        uint16
	Buy_guid   uint64 //>出售物品guid
	Buy_amount uint16 //>出售物品数量
}

type GuildCustomNtf struct { //>帮派自定义信息通知
	Mid         uint16    
	Pid         uint16    
	Player_guid uint64     //>角色guid
	Is_guild    uint8      //>是否是帮派：0-帮派成员，1-帮派
	Customs     CustomData //>自定义变量
}

type PreTurnRoundNtf struct { //>新回合前通知
	Mid   uint16
	Pid   uint16
	Round uint16 //>回合数
}

type FighterSpecialPetNtf struct { //>战斗主宠通知
	Mid         uint16
	Pid         uint16
	Special_pet uint64 //>指定主宠
}

func (proto *KeepAliveReq) GetMid() uint16 {
	return 102
}

func (proto *KeepAliveReq) GetPid() uint16 {
	return 1
}

func (proto *KeepAliveReq) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(1)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Seq) {
		return false
	}

	return true
}

func (proto *KeepAliveReq) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Seq) {
		return false
	}

	return true
}

func (proto *KeepAliveAck) GetMid() uint16 {
	return 102
}

func (proto *KeepAliveAck) GetPid() uint16 {
	return 2
}

func (proto *KeepAliveAck) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(2)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Seq) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Tick) {
		return false
	}

	return true
}

func (proto *KeepAliveAck) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Seq) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Tick) {
		return false
	}

	return true
}

func (proto *AttrNtf) GetMid() uint16 {
	return 102
}

func (proto *AttrNtf) GetPid() uint16 {
	return 3
}

func (proto *AttrNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(3)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Attrs, uint16(65535)) {
		return false
	}

	return true
}

func (proto *AttrNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Attrs, uint16(65535)) {
		return false
	}

	return true
}

func (proto *PlayerAppearNtf) GetMid() uint16 {
	return 102
}

func (proto *PlayerAppearNtf) GetPid() uint16 {
	return 4
}

func (proto *PlayerAppearNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(4)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Sn) {
		return false
	}

	if !ProtoWriteInteger(b, proto.X) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Y) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Attrs, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Buffs, uint8(255)) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Custom) {
		return false
	}

	if !ProtoWriteString(b, proto.Name, 255) {
		return false
	}

	return true
}

func (proto *PlayerAppearNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Sn) {
		return false
	}

	if !ProtoReadInteger(b, &proto.X) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Y) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Attrs, uint16(65535)) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Buffs, uint8(255)) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Custom) {
		return false
	}

	if !ProtoReadString(b, &proto.Name, 255) {
		return false
	}

	return true
}

func (proto *NPCAppearNtf) GetMid() uint16 {
	return 102
}

func (proto *NPCAppearNtf) GetPid() uint16 {
	return 5
}

func (proto *NPCAppearNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(5)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Template_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.X) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Y) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Dir) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Attrs, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Custom) {
		return false
	}

	if !ProtoWriteString(b, proto.Name, 255) {
		return false
	}

	return true
}

func (proto *NPCAppearNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Template_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.X) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Y) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Dir) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Attrs, uint16(65535)) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Custom) {
		return false
	}

	if !ProtoReadString(b, &proto.Name, 255) {
		return false
	}

	return true
}

func (proto *ItemAppearNtf) GetMid() uint16 {
	return 102
}

func (proto *ItemAppearNtf) GetPid() uint16 {
	return 6
}

func (proto *ItemAppearNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(6)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Template_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.X) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Y) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Custom) {
		return false
	}

	if !ProtoWriteString(b, proto.Name, 255) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Dyn_attrs, uint16(65535)) {
		return false
	}

	return true
}

func (proto *ItemAppearNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Template_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.X) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Y) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Custom) {
		return false
	}

	if !ProtoReadString(b, &proto.Name, 255) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Dyn_attrs, uint16(65535)) {
		return false
	}

	return true
}

func (proto *ObjDisAppearNtf) GetMid() uint16 {
	return 102
}

func (proto *ObjDisAppearNtf) GetPid() uint16 {
	return 7
}

func (proto *ObjDisAppearNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(7)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	return true
}

func (proto *ObjDisAppearNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	return true
}

func (proto *ObjMoveNtf) GetMid() uint16 {
	return 102
}

func (proto *ObjMoveNtf) GetPid() uint16 {
	return 8
}

func (proto *ObjMoveNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(8)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Dst_x) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Dst_y) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Type) {
		return false
	}

	return true
}

func (proto *ObjMoveNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Dst_x) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Dst_y) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Type) {
		return false
	}

	return true
}

func (proto *EnterMapNtf) GetMid() uint16 {
	return 102
}

func (proto *EnterMapNtf) GetPid() uint16 {
	return 9
}

func (proto *EnterMapNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(9)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Template_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Dst_x) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Dst_y) {
		return false
	}

	return true
}

func (proto *EnterMapNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Template_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Dst_x) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Dst_y) {
		return false
	}

	return true
}

func (proto *MoveReq) GetMid() uint16 {
	return 102
}

func (proto *MoveReq) GetPid() uint16 {
	return 10
}

func (proto *MoveReq) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(10)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Dst_x) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Dst_y) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Cur_x) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Cur_y) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Idx) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Tick) {
		return false
	}

	return true
}

func (proto *MoveReq) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Dst_x) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Dst_y) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Cur_x) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Cur_y) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Idx) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Tick) {
		return false
	}

	return true
}

func (proto *MoveAck) GetMid() uint16 {
	return 102
}

func (proto *MoveAck) GetPid() uint16 {
	return 11
}

func (proto *MoveAck) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(11)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Errcode) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Dst_x) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Dst_y) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Idx) {
		return false
	}

	return true
}

func (proto *MoveAck) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Errcode) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Dst_x) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Dst_y) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Idx) {
		return false
	}

	return true
}

func (proto *JumpMapReq) GetMid() uint16 {
	return 102
}

func (proto *JumpMapReq) GetPid() uint16 {
	return 12
}

func (proto *JumpMapReq) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(12)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Idx) {
		return false
	}

	return true
}

func (proto *JumpMapReq) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Idx) {
		return false
	}

	return true
}

func (proto *JumpMapAck) GetMid() uint16 {
	return 102
}

func (proto *JumpMapAck) GetPid() uint16 {
	return 13
}

func (proto *JumpMapAck) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(13)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Errcode) {
		return false
	}

	return true
}

func (proto *JumpMapAck) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Errcode) {
		return false
	}

	return true
}

func (proto *AddJumpMapRegionNtf) GetMid() uint16 {
	return 102
}

func (proto *AddJumpMapRegionNtf) GetPid() uint16 {
	return 14
}

func (proto *AddJumpMapRegionNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(14)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Idx) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Jump_region) {
		return false
	}

	return true
}

func (proto *AddJumpMapRegionNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Idx) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Jump_region) {
		return false
	}

	return true
}

func (proto *DelJumpMapRegionNtf) GetMid() uint16 {
	return 102
}

func (proto *DelJumpMapRegionNtf) GetPid() uint16 {
	return 15
}

func (proto *DelJumpMapRegionNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(15)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Idx) {
		return false
	}

	return true
}

func (proto *DelJumpMapRegionNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Idx) {
		return false
	}

	return true
}

func (proto *ItemAddNtf) GetMid() uint16 {
	return 102
}

func (proto *ItemAddNtf) GetPid() uint16 {
	return 16
}

func (proto *ItemAddNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(16)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Container_type) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Items, uint16(65535)) {
		return false
	}

	return true
}

func (proto *ItemAddNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Container_type) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Items, uint16(65535)) {
		return false
	}

	return true
}

func (proto *ItemUpdateNtf) GetMid() uint16 {
	return 102
}

func (proto *ItemUpdateNtf) GetPid() uint16 {
	return 17
}

func (proto *ItemUpdateNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(17)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Container_type) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Attrs, uint16(65535)) {
		return false
	}

	return true
}

func (proto *ItemUpdateNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Container_type) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Attrs, uint16(65535)) {
		return false
	}

	return true
}

func (proto *ItemDestroyNtf) GetMid() uint16 {
	return 102
}

func (proto *ItemDestroyNtf) GetPid() uint16 {
	return 18
}

func (proto *ItemDestroyNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(18)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Container_type) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Item_guids, uint16(65535)) {
		return false
	}

	return true
}

func (proto *ItemDestroyNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Container_type) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Item_guids, uint16(65535)) {
		return false
	}

	return true
}

func (proto *TipsMsgNtf) GetMid() uint16 {
	return 102
}

func (proto *TipsMsgNtf) GetPid() uint16 {
	return 19
}

func (proto *TipsMsgNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(19)) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Msg, uint16(65535)) {
		return false
	}

	return true
}

func (proto *TipsMsgNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Msg, uint16(65535)) {
		return false
	}

	return true
}

func (proto *TopMsgNtf) GetMid() uint16 {
	return 102
}

func (proto *TopMsgNtf) GetPid() uint16 {
	return 20
}

func (proto *TopMsgNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(20)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Foreground) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Background) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Count) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Msg, uint16(65535)) {
		return false
	}

	return true
}

func (proto *TopMsgNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Foreground) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Background) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Count) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Msg, uint16(65535)) {
		return false
	}

	return true
}

func (proto *SysMsgNtf) GetMid() uint16 {
	return 102
}

func (proto *SysMsgNtf) GetPid() uint16 {
	return 21
}

func (proto *SysMsgNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(21)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Foreground) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Background) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Msg, uint16(65535)) {
		return false
	}

	return true
}

func (proto *SysMsgNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Foreground) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Background) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Msg, uint16(65535)) {
		return false
	}

	return true
}

func (proto *PopupMsgNtf) GetMid() uint16 {
	return 102
}

func (proto *PopupMsgNtf) GetPid() uint16 {
	return 22
}

func (proto *PopupMsgNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(22)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Type) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Msg, uint16(65535)) {
		return false
	}

	return true
}

func (proto *PopupMsgNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Type) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Msg, uint16(65535)) {
		return false
	}

	return true
}

func (proto *ItemContainerNtf) GetMid() uint16 {
	return 102
}

func (proto *ItemContainerNtf) GetPid() uint16 {
	return 23
}

func (proto *ItemContainerNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(23)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Container_type) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Capacity) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Items, uint16(65535)) {
		return false
	}

	return true
}

func (proto *ItemContainerNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Container_type) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Capacity) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Items, uint16(65535)) {
		return false
	}

	return true
}

func (proto *ItemContainerUpdateNtf) GetMid() uint16 {
	return 102
}

func (proto *ItemContainerUpdateNtf) GetPid() uint16 {
	return 24
}

func (proto *ItemContainerUpdateNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(24)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Container_type) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Capacity) {
		return false
	}

	return true
}

func (proto *ItemContainerUpdateNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Container_type) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Capacity) {
		return false
	}

	return true
}

func (proto *SubmitForm) GetMid() uint16 {
	return 102
}

func (proto *SubmitForm) GetPid() uint16 {
	return 25
}

func (proto *SubmitForm) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(25)) {
		return false
	}

	if !ProtoWriteString(b, proto.Form, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Func, 255) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Args, uint8(255)) {
		return false
	}

	return true
}

func (proto *SubmitForm) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadString(b, &proto.Form, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Func, 255) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Args, uint8(255)) {
		return false
	}

	return true
}

func (proto *ShowFormNtf) GetMid() uint16 {
	return 102
}

func (proto *ShowFormNtf) GetPid() uint16 {
	return 26
}

func (proto *ShowFormNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(26)) {
		return false
	}

	if !ProtoWriteString(b, proto.Form, 255) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Compressed) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Context, uint16(65535)) {
		return false
	}

	return true
}

func (proto *ShowFormNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadString(b, &proto.Form, 255) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Compressed) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Context, uint16(65535)) {
		return false
	}

	return true
}

func (proto *ExecuteGMReq) GetMid() uint16 {
	return 102
}

func (proto *ExecuteGMReq) GetPid() uint16 {
	return 27
}

func (proto *ExecuteGMReq) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(27)) {
		return false
	}

	if !ProtoWriteString(b, proto.Cmd, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Arg, 255) {
		return false
	}

	return true
}

func (proto *ExecuteGMReq) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadString(b, &proto.Cmd, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Arg, 255) {
		return false
	}

	return true
}

func (proto *FightBeginNtf) GetMid() uint16 {
	return 102
}

func (proto *FightBeginNtf) GetPid() uint16 {
	return 28
}

func (proto *FightBeginNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(28)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Groups, uint8(255)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Self_group) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Is_pvp) {
		return false
	}

	return true
}

func (proto *FightBeginNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Groups, uint8(255)) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Self_group) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Is_pvp) {
		return false
	}

	return true
}

func (proto *TurnRoundNtf) GetMid() uint16 {
	return 102
}

func (proto *TurnRoundNtf) GetPid() uint16 {
	return 29
}

func (proto *TurnRoundNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(29)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Round) {
		return false
	}

	return true
}

func (proto *TurnRoundNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Round) {
		return false
	}

	return true
}

func (proto *FightOperateListNtf) GetMid() uint16 {
	return 102
}

func (proto *FightOperateListNtf) GetPid() uint16 {
	return 30
}

func (proto *FightOperateListNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(30)) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Player_operates, uint8(255)) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Pet_operates, uint8(255)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Count_down) {
		return false
	}

	return true
}

func (proto *FightOperateListNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Player_operates, uint8(255)) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Pet_operates, uint8(255)) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Count_down) {
		return false
	}

	return true
}

func (proto *FightOperateReq) GetMid() uint16 {
	return 102
}

func (proto *FightOperateReq) GetPid() uint16 {
	return 31
}

func (proto *FightOperateReq) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(31)) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Operate) {
		return false
	}

	return true
}

func (proto *FightOperateReq) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Operate) {
		return false
	}

	return true
}

func (proto *FightOperateAck) GetMid() uint16 {
	return 102
}

func (proto *FightOperateAck) GetPid() uint16 {
	return 32
}

func (proto *FightOperateAck) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(32)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Errcode) {
		return false
	}

	if !ProtoWriteString(b, proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *FightOperateAck) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Errcode) {
		return false
	}

	if !ProtoReadString(b, &proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *FightOperateNtf) GetMid() uint16 {
	return 102
}

func (proto *FightOperateNtf) GetPid() uint16 {
	return 33
}

func (proto *FightOperateNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(33)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Fighter_id) {
		return false
	}

	return true
}

func (proto *FightOperateNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Fighter_id) {
		return false
	}

	return true
}

func (proto *FightDisplayNtf) GetMid() uint16 {
	return 102
}

func (proto *FightDisplayNtf) GetPid() uint16 {
	return 34
}

func (proto *FightDisplayNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(34)) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Data, uint32(4294967295)) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Crypto, uint8(255)) {
		return false
	}

	return true
}

func (proto *FightDisplayNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Data, uint32(4294967295)) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Crypto, uint8(255)) {
		return false
	}

	return true
}

func (proto *FightDisplayCompleteNtf) GetMid() uint16 {
	return 102
}

func (proto *FightDisplayCompleteNtf) GetPid() uint16 {
	return 35
}

func (proto *FightDisplayCompleteNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(35)) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Crypto, uint8(255)) {
		return false
	}

	return true
}

func (proto *FightDisplayCompleteNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Crypto, uint8(255)) {
		return false
	}

	return true
}

func (proto *FightAutoReq) GetMid() uint16 {
	return 102
}

func (proto *FightAutoReq) GetPid() uint16 {
	return 36
}

func (proto *FightAutoReq) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(36)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Is_auto) {
		return false
	}

	return true
}

func (proto *FightAutoReq) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Is_auto) {
		return false
	}

	return true
}

func (proto *FightAutoAck) GetMid() uint16 {
	return 102
}

func (proto *FightAutoAck) GetPid() uint16 {
	return 37
}

func (proto *FightAutoAck) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(37)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Is_auto) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Errcode) {
		return false
	}

	if !ProtoWriteString(b, proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *FightAutoAck) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Is_auto) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Errcode) {
		return false
	}

	if !ProtoReadString(b, &proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *FightAutoNtf) GetMid() uint16 {
	return 102
}

func (proto *FightAutoNtf) GetPid() uint16 {
	return 38
}

func (proto *FightAutoNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(38)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Fighter_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Is_auto) {
		return false
	}

	return true
}

func (proto *FightAutoNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Fighter_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Is_auto) {
		return false
	}

	return true
}

func (proto *FightAutoSkillReq) GetMid() uint16 {
	return 102
}

func (proto *FightAutoSkillReq) GetPid() uint16 {
	return 39
}

func (proto *FightAutoSkillReq) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(39)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Actor) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Skillid) {
		return false
	}

	return true
}

func (proto *FightAutoSkillReq) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Actor) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Skillid) {
		return false
	}

	return true
}

func (proto *FightAutoSkillAck) GetMid() uint16 {
	return 102
}

func (proto *FightAutoSkillAck) GetPid() uint16 {
	return 40
}

func (proto *FightAutoSkillAck) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(40)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Actor) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Skillid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Errcode) {
		return false
	}

	if !ProtoWriteString(b, proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *FightAutoSkillAck) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Actor) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Skillid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Errcode) {
		return false
	}

	if !ProtoReadString(b, &proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *FightEndNtf) GetMid() uint16 {
	return 102
}

func (proto *FightEndNtf) GetPid() uint16 {
	return 41
}

func (proto *FightEndNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(41)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Result) {
		return false
	}

	return true
}

func (proto *FightEndNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Result) {
		return false
	}

	return true
}

func (proto *AddFighterNtf) GetMid() uint16 {
	return 102
}

func (proto *AddFighterNtf) GetPid() uint16 {
	return 42
}

func (proto *AddFighterNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(42)) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Data) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Group) {
		return false
	}

	return true
}

func (proto *AddFighterNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Data) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Group) {
		return false
	}

	return true
}

func (proto *DelFighterNtf) GetMid() uint16 {
	return 102
}

func (proto *DelFighterNtf) GetPid() uint16 {
	return 43
}

func (proto *DelFighterNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(43)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Fighter_id) {
		return false
	}

	return true
}

func (proto *DelFighterNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Fighter_id) {
		return false
	}

	return true
}

func (proto *AddFightPetData) GetMid() uint16 {
	return 102
}

func (proto *AddFightPetData) GetPid() uint16 {
	return 44
}

func (proto *AddFightPetData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(44)) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Pet) {
		return false
	}

	return true
}

func (proto *AddFightPetData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Pet) {
		return false
	}

	return true
}

func (proto *UpdateFightPetData) GetMid() uint16 {
	return 102
}

func (proto *UpdateFightPetData) GetPid() uint16 {
	return 45
}

func (proto *UpdateFightPetData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(45)) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Pet) {
		return false
	}

	return true
}

func (proto *UpdateFightPetData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Pet) {
		return false
	}

	return true
}

func (proto *PlayerKillReq) GetMid() uint16 {
	return 102
}

func (proto *PlayerKillReq) GetPid() uint16 {
	return 46
}

func (proto *PlayerKillReq) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(46)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Target) {
		return false
	}

	return true
}

func (proto *PlayerKillReq) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Target) {
		return false
	}

	return true
}

func (proto *PlayerKillAck) GetMid() uint16 {
	return 102
}

func (proto *PlayerKillAck) GetPid() uint16 {
	return 47
}

func (proto *PlayerKillAck) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(47)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Target) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Errcode) {
		return false
	}

	if !ProtoWriteString(b, proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *PlayerKillAck) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Target) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Errcode) {
		return false
	}

	if !ProtoReadString(b, &proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *EnterFightViewReq) GetMid() uint16 {
	return 102
}

func (proto *EnterFightViewReq) GetPid() uint16 {
	return 48
}

func (proto *EnterFightViewReq) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(48)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Target) {
		return false
	}

	return true
}

func (proto *EnterFightViewReq) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Target) {
		return false
	}

	return true
}

func (proto *EnterFightViewAck) GetMid() uint16 {
	return 102
}

func (proto *EnterFightViewAck) GetPid() uint16 {
	return 49
}

func (proto *EnterFightViewAck) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(49)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Target) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Errcode) {
		return false
	}

	if !ProtoWriteString(b, proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *EnterFightViewAck) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Target) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Errcode) {
		return false
	}

	if !ProtoReadString(b, &proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *LeaveFightViewReq) GetMid() uint16 {
	return 102
}

func (proto *LeaveFightViewReq) GetPid() uint16 {
	return 50
}

func (proto *LeaveFightViewReq) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(50)) {
		return false
	}

	return true
}

func (proto *LeaveFightViewReq) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	return true
}

func (proto *LeaveFightViewAck) GetMid() uint16 {
	return 102
}

func (proto *LeaveFightViewAck) GetPid() uint16 {
	return 51
}

func (proto *LeaveFightViewAck) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(51)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Errcode) {
		return false
	}

	if !ProtoWriteString(b, proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *LeaveFightViewAck) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Errcode) {
		return false
	}

	if !ProtoReadString(b, &proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *TeamLeaderOprReq) GetMid() uint16 {
	return 102
}

func (proto *TeamLeaderOprReq) GetPid() uint16 {
	return 52
}

func (proto *TeamLeaderOprReq) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(52)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Opr) {
		return false
	}

	return true
}

func (proto *TeamLeaderOprReq) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Opr) {
		return false
	}

	return true
}

func (proto *TeamNtf) GetMid() uint16 {
	return 102
}

func (proto *TeamNtf) GetPid() uint16 {
	return 53
}

func (proto *TeamNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(53)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Team_guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Target) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Permission) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Min_require_level) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Max_require_level) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Leader_guid) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Leader_guards, uint8(255)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Members, uint8(255)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Min_require_reinc) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Max_require_reinc) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Customs) {
		return false
	}

	return true
}

func (proto *TeamNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Team_guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Target) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Permission) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Min_require_level) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Max_require_level) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Leader_guid) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Leader_guards, uint8(255)) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Members, uint8(255)) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Min_require_reinc) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Max_require_reinc) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Customs) {
		return false
	}

	return true
}

func (proto *TeamLeaderNtf) GetMid() uint16 {
	return 102
}

func (proto *TeamLeaderNtf) GetPid() uint16 {
	return 54
}

func (proto *TeamLeaderNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(54)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Leader_guid) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Leader_guards, uint8(255)) {
		return false
	}

	return true
}

func (proto *TeamLeaderNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Leader_guid) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Leader_guards, uint8(255)) {
		return false
	}

	return true
}

func (proto *TeamDestroyNtf) GetMid() uint16 {
	return 102
}

func (proto *TeamDestroyNtf) GetPid() uint16 {
	return 55
}

func (proto *TeamDestroyNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(55)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Team_guid) {
		return false
	}

	return true
}

func (proto *TeamDestroyNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Team_guid) {
		return false
	}

	return true
}

func (proto *TeamMemberNtf) GetMid() uint16 {
	return 102
}

func (proto *TeamMemberNtf) GetPid() uint16 {
	return 56
}

func (proto *TeamMemberNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(56)) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Member) {
		return false
	}

	return true
}

func (proto *TeamMemberNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Member) {
		return false
	}

	return true
}

func (proto *TeamMemberLeaveNtf) GetMid() uint16 {
	return 102
}

func (proto *TeamMemberLeaveNtf) GetPid() uint16 {
	return 57
}

func (proto *TeamMemberLeaveNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(57)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Action) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Player_guid) {
		return false
	}

	return true
}

func (proto *TeamMemberLeaveNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Action) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Player_guid) {
		return false
	}

	return true
}

func (proto *NpcSelectReq) GetMid() uint16 {
	return 102
}

func (proto *NpcSelectReq) GetPid() uint16 {
	return 58
}

func (proto *NpcSelectReq) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(58)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Npc_guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Select_id) {
		return false
	}

	return true
}

func (proto *NpcSelectReq) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Npc_guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Select_id) {
		return false
	}

	return true
}

func (proto *NpcSelectAck) GetMid() uint16 {
	return 102
}

func (proto *NpcSelectAck) GetPid() uint16 {
	return 59
}

func (proto *NpcSelectAck) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(59)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Npc_guid) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Options, uint16(65535)) {
		return false
	}

	if !ProtoWriteString(b, proto.Msg, 4096) {
		return false
	}

	return true
}

func (proto *NpcSelectAck) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Npc_guid) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Options, uint16(65535)) {
		return false
	}

	if !ProtoReadString(b, &proto.Msg, 4096) {
		return false
	}

	return true
}

func (proto *NpcTalkReq) GetMid() uint16 {
	return 102
}

func (proto *NpcTalkReq) GetPid() uint16 {
	return 60
}

func (proto *NpcTalkReq) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(60)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Npc_guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Compressed) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Talk, uint16(65535)) {
		return false
	}

	return true
}

func (proto *NpcTalkReq) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Npc_guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Compressed) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Talk, uint16(65535)) {
		return false
	}

	return true
}

func (proto *NpcTalkAck) GetMid() uint16 {
	return 102
}

func (proto *NpcTalkAck) GetPid() uint16 {
	return 61
}

func (proto *NpcTalkAck) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(61)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Npc_guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Compressed) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Talk, uint16(65535)) {
		return false
	}

	return true
}

func (proto *NpcTalkAck) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Npc_guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Compressed) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Talk, uint16(65535)) {
		return false
	}

	return true
}

func (proto *InviteMsgNtf) GetMid() uint16 {
	return 102
}

func (proto *InviteMsgNtf) GetPid() uint16 {
	return 62
}

func (proto *InviteMsgNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(62)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Inviter_guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Type) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Msg, uint16(65535)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Interval) {
		return false
	}

	return true
}

func (proto *InviteMsgNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Inviter_guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Type) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Msg, uint16(65535)) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Interval) {
		return false
	}

	return true
}

func (proto *ReplyInvite) GetMid() uint16 {
	return 102
}

func (proto *ReplyInvite) GetPid() uint16 {
	return 63
}

func (proto *ReplyInvite) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(63)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Inviter_guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Type) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Agreed) {
		return false
	}

	return true
}

func (proto *ReplyInvite) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Inviter_guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Type) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Agreed) {
		return false
	}

	return true
}

func (proto *MoveItem) GetMid() uint16 {
	return 102
}

func (proto *MoveItem) GetPid() uint16 {
	return 64
}

func (proto *MoveItem) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(64)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Item_guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Container_type) {
		return false
	}

	return true
}

func (proto *MoveItem) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Item_guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Container_type) {
		return false
	}

	return true
}

func (proto *UseItem) GetMid() uint16 {
	return 102
}

func (proto *UseItem) GetPid() uint16 {
	return 65
}

func (proto *UseItem) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(65)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Item_guid) {
		return false
	}

	return true
}

func (proto *UseItem) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Item_guid) {
		return false
	}

	return true
}

func (proto *RearrangeItem) GetMid() uint16 {
	return 102
}

func (proto *RearrangeItem) GetPid() uint16 {
	return 66
}

func (proto *RearrangeItem) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(66)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Container_type) {
		return false
	}

	return true
}

func (proto *RearrangeItem) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Container_type) {
		return false
	}

	return true
}

func (proto *SkillContainerNtf) GetMid() uint16 {
	return 102
}

func (proto *SkillContainerNtf) GetPid() uint16 {
	return 67
}

func (proto *SkillContainerNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(67)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Skills, uint16(65535)) {
		return false
	}

	return true
}

func (proto *SkillContainerNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Skills, uint16(65535)) {
		return false
	}

	return true
}

func (proto *AddSkillNtf) GetMid() uint16 {
	return 102
}

func (proto *AddSkillNtf) GetPid() uint16 {
	return 68
}

func (proto *AddSkillNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(68)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Skill) {
		return false
	}

	return true
}

func (proto *AddSkillNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Skill) {
		return false
	}

	return true
}

func (proto *UpdateSkillNtf) GetMid() uint16 {
	return 102
}

func (proto *UpdateSkillNtf) GetPid() uint16 {
	return 69
}

func (proto *UpdateSkillNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(69)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Attr) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Value) {
		return false
	}

	return true
}

func (proto *UpdateSkillNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Attr) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Value) {
		return false
	}

	return true
}

func (proto *DelSkillNtf) GetMid() uint16 {
	return 102
}

func (proto *DelSkillNtf) GetPid() uint16 {
	return 70
}

func (proto *DelSkillNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(70)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Id) {
		return false
	}

	return true
}

func (proto *DelSkillNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Id) {
		return false
	}

	return true
}

func (proto *PetAppearNtf) GetMid() uint16 {
	return 102
}

func (proto *PetAppearNtf) GetPid() uint16 {
	return 71
}

func (proto *PetAppearNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(71)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteString(b, proto.Name, 255) {
		return false
	}

	if !ProtoWriteInteger(b, proto.X) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Y) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Attrs, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Custom) {
		return false
	}

	return true
}

func (proto *PetAppearNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadString(b, &proto.Name, 255) {
		return false
	}

	if !ProtoReadInteger(b, &proto.X) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Y) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Attrs, uint16(65535)) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Custom) {
		return false
	}

	return true
}

func (proto *PetContainerNtf) GetMid() uint16 {
	return 102
}

func (proto *PetContainerNtf) GetPid() uint16 {
	return 72
}

func (proto *PetContainerNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(72)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Container_type) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Capacity) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Pets, uint16(65535)) {
		return false
	}

	return true
}

func (proto *PetContainerNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Container_type) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Capacity) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Pets, uint16(65535)) {
		return false
	}

	return true
}

func (proto *PetContainerUpdateNtf) GetMid() uint16 {
	return 102
}

func (proto *PetContainerUpdateNtf) GetPid() uint16 {
	return 73
}

func (proto *PetContainerUpdateNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(73)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Container_type) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Capacity) {
		return false
	}

	return true
}

func (proto *PetContainerUpdateNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Container_type) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Capacity) {
		return false
	}

	return true
}

func (proto *PetAddNtf) GetMid() uint16 {
	return 102
}

func (proto *PetAddNtf) GetPid() uint16 {
	return 74
}

func (proto *PetAddNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(74)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Container_type) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Pet) {
		return false
	}

	return true
}

func (proto *PetAddNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Container_type) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Pet) {
		return false
	}

	return true
}

func (proto *PetDestroyNtf) GetMid() uint16 {
	return 102
}

func (proto *PetDestroyNtf) GetPid() uint16 {
	return 75
}

func (proto *PetDestroyNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(75)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Container_type) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	return true
}

func (proto *PetDestroyNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Container_type) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	return true
}

func (proto *SetPetLineup) GetMid() uint16 {
	return 102
}

func (proto *SetPetLineup) GetPid() uint16 {
	return 76
}

func (proto *SetPetLineup) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(76)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Lineup) {
		return false
	}

	return true
}

func (proto *SetPetLineup) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Lineup) {
		return false
	}

	return true
}

func (proto *ShowPet) GetMid() uint16 {
	return 102
}

func (proto *ShowPet) GetPid() uint16 {
	return 77
}

func (proto *ShowPet) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(77)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Show) {
		return false
	}

	return true
}

func (proto *ShowPet) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Show) {
		return false
	}

	return true
}

func (proto *ReleasePet) GetMid() uint16 {
	return 102
}

func (proto *ReleasePet) GetPid() uint16 {
	return 78
}

func (proto *ReleasePet) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(78)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	return true
}

func (proto *ReleasePet) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	return true
}

func (proto *MovePet) GetMid() uint16 {
	return 102
}

func (proto *MovePet) GetPid() uint16 {
	return 79
}

func (proto *MovePet) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(79)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Container_type) {
		return false
	}

	return true
}

func (proto *MovePet) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Container_type) {
		return false
	}

	return true
}

func (proto *ShopOpenNtf) GetMid() uint16 {
	return 102
}

func (proto *ShopOpenNtf) GetPid() uint16 {
	return 80
}

func (proto *ShopOpenNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(80)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Shop_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Def_item_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Type) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Shop_item_list, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Buy_back_list, uint16(65535)) {
		return false
	}

	return true
}

func (proto *ShopOpenNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Shop_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Def_item_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Type) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Shop_item_list, uint16(65535)) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Buy_back_list, uint16(65535)) {
		return false
	}

	return true
}

func (proto *ShopBuyNtf) GetMid() uint16 {
	return 102
}

func (proto *ShopBuyNtf) GetPid() uint16 {
	return 81
}

func (proto *ShopBuyNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(81)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Shop_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Shop_item_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Shop_item_num) {
		return false
	}

	return true
}

func (proto *ShopBuyNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Shop_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Shop_item_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Shop_item_num) {
		return false
	}

	return true
}

func (proto *SellNtf) GetMid() uint16 {
	return 102
}

func (proto *SellNtf) GetPid() uint16 {
	return 82
}

func (proto *SellNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(82)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Sell_items, uint16(65535)) {
		return false
	}

	return true
}

func (proto *SellNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Sell_items, uint16(65535)) {
		return false
	}

	return true
}

func (proto *BuyBackNtf) GetMid() uint16 {
	return 102
}

func (proto *BuyBackNtf) GetPid() uint16 {
	return 83
}

func (proto *BuyBackNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(83)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Buy_guid) {
		return false
	}

	return true
}

func (proto *BuyBackNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Buy_guid) {
		return false
	}

	return true
}

func (proto *BuyBackListNtf) GetMid() uint16 {
	return 102
}

func (proto *BuyBackListNtf) GetPid() uint16 {
	return 84
}

func (proto *BuyBackListNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(84)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Buy_back_list, uint16(65535)) {
		return false
	}

	return true
}

func (proto *BuyBackListNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Buy_back_list, uint16(65535)) {
		return false
	}

	return true
}

func (proto *TeamAttrNtf) GetMid() uint16 {
	return 102
}

func (proto *TeamAttrNtf) GetPid() uint16 {
	return 85
}

func (proto *TeamAttrNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(85)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Attrs, uint16(65535)) {
		return false
	}

	return true
}

func (proto *TeamAttrNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Attrs, uint16(65535)) {
		return false
	}

	return true
}

func (proto *TipsMsgExNtf) GetMid() uint16 {
	return 102
}

func (proto *TipsMsgExNtf) GetPid() uint16 {
	return 86
}

func (proto *TipsMsgExNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(86)) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Msg, uint16(65535)) {
		return false
	}

	return true
}

func (proto *TipsMsgExNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Msg, uint16(65535)) {
		return false
	}

	return true
}

func (proto *ItemNewAddNtf) GetMid() uint16 {
	return 102
}

func (proto *ItemNewAddNtf) GetPid() uint16 {
	return 87
}

func (proto *ItemNewAddNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(87)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Item_guid) {
		return false
	}

	return true
}

func (proto *ItemNewAddNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Item_guid) {
		return false
	}

	return true
}

func (proto *QuestContainerNtf) GetMid() uint16 {
	return 102
}

func (proto *QuestContainerNtf) GetPid() uint16 {
	return 88
}

func (proto *QuestContainerNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(88)) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Buffer, uint16(65535)) {
		return false
	}

	return true
}

func (proto *QuestContainerNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Buffer, uint16(65535)) {
		return false
	}

	return true
}

func (proto *QuestAddNtf) GetMid() uint16 {
	return 102
}

func (proto *QuestAddNtf) GetPid() uint16 {
	return 89
}

func (proto *QuestAddNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(89)) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Buffer, uint16(65535)) {
		return false
	}

	return true
}

func (proto *QuestAddNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Buffer, uint16(65535)) {
		return false
	}

	return true
}

func (proto *QuestStateNtf) GetMid() uint16 {
	return 102
}

func (proto *QuestStateNtf) GetPid() uint16 {
	return 90
}

func (proto *QuestStateNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(90)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Quest_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Quest_state) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Changed_time) {
		return false
	}

	return true
}

func (proto *QuestStateNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Quest_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Quest_state) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Changed_time) {
		return false
	}

	return true
}

func (proto *QuestNpcStateNtf) GetMid() uint16 {
	return 102
}

func (proto *QuestNpcStateNtf) GetPid() uint16 {
	return 91
}

func (proto *QuestNpcStateNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(91)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Npc_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.State) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Logo) {
		return false
	}

	return true
}

func (proto *QuestNpcStateNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Npc_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.State) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Logo) {
		return false
	}

	return true
}

func (proto *QuestTrackCountNtf) GetMid() uint16 {
	return 102
}

func (proto *QuestTrackCountNtf) GetPid() uint16 {
	return 92
}

func (proto *QuestTrackCountNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(92)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Quest_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Track_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Track_count) {
		return false
	}

	return true
}

func (proto *QuestTrackCountNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Quest_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Track_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Track_count) {
		return false
	}

	return true
}

func (proto *QuestTalkSelectReq) GetMid() uint16 {
	return 102
}

func (proto *QuestTalkSelectReq) GetPid() uint16 {
	return 93
}

func (proto *QuestTalkSelectReq) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(93)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Npc_guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Quest_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Option_id) {
		return false
	}

	return true
}

func (proto *QuestTalkSelectReq) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Npc_guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Quest_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Option_id) {
		return false
	}

	return true
}

func (proto *QuestTalkSelectAck) GetMid() uint16 {
	return 102
}

func (proto *QuestTalkSelectAck) GetPid() uint16 {
	return 94
}

func (proto *QuestTalkSelectAck) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(94)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Npc_guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Quest_id) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Talk, uint16(65535)) {
		return false
	}

	return true
}

func (proto *QuestTalkSelectAck) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Npc_guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Quest_id) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Talk, uint16(65535)) {
		return false
	}

	return true
}

func (proto *ObjStrCustomNtf) GetMid() uint16 {
	return 102
}

func (proto *ObjStrCustomNtf) GetPid() uint16 {
	return 95
}

func (proto *ObjStrCustomNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(95)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Custom) {
		return false
	}

	return true
}

func (proto *ObjStrCustomNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Custom) {
		return false
	}

	return true
}

func (proto *ObjIntCustomNtf) GetMid() uint16 {
	return 102
}

func (proto *ObjIntCustomNtf) GetPid() uint16 {
	return 96
}

func (proto *ObjIntCustomNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(96)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Custom) {
		return false
	}

	return true
}

func (proto *ObjIntCustomNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Custom) {
		return false
	}

	return true
}

func (proto *ObjDynAttrNtf) GetMid() uint16 {
	return 102
}

func (proto *ObjDynAttrNtf) GetPid() uint16 {
	return 97
}

func (proto *ObjDynAttrNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(97)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Dyn_attr) {
		return false
	}

	return true
}

func (proto *ObjDynAttrNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Dyn_attr) {
		return false
	}

	return true
}

func (proto *CommitQuestItemRsp) GetMid() uint16 {
	return 102
}

func (proto *CommitQuestItemRsp) GetPid() uint16 {
	return 98
}

func (proto *CommitQuestItemRsp) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(98)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Quest_id) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Items, uint8(255)) {
		return false
	}

	return true
}

func (proto *CommitQuestItemRsp) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Quest_id) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Items, uint8(255)) {
		return false
	}

	return true
}

func (proto *CommitQuestItemNtf) GetMid() uint16 {
	return 102
}

func (proto *CommitQuestItemNtf) GetPid() uint16 {
	return 99
}

func (proto *CommitQuestItemNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(99)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Quest_id) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Items, uint8(255)) {
		return false
	}

	return true
}

func (proto *CommitQuestItemNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Quest_id) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Items, uint8(255)) {
		return false
	}

	return true
}

func (proto *CommitQuestPetRsp) GetMid() uint16 {
	return 102
}

func (proto *CommitQuestPetRsp) GetPid() uint16 {
	return 100
}

func (proto *CommitQuestPetRsp) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(100)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Quest_id) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Pets, uint8(255)) {
		return false
	}

	return true
}

func (proto *CommitQuestPetRsp) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Quest_id) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Pets, uint8(255)) {
		return false
	}

	return true
}

func (proto *CommitQuestPetNtf) GetMid() uint16 {
	return 102
}

func (proto *CommitQuestPetNtf) GetPid() uint16 {
	return 101
}

func (proto *CommitQuestPetNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(101)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Quest_id) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Pets, uint8(255)) {
		return false
	}

	return true
}

func (proto *CommitQuestPetNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Quest_id) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Pets, uint8(255)) {
		return false
	}

	return true
}

func (proto *AbandonQuest) GetMid() uint16 {
	return 102
}

func (proto *AbandonQuest) GetPid() uint16 {
	return 102
}

func (proto *AbandonQuest) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Quest_id) {
		return false
	}

	return true
}

func (proto *AbandonQuest) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Quest_id) {
		return false
	}

	return true
}

func (proto *SkillTipsReq) GetMid() uint16 {
	return 102
}

func (proto *SkillTipsReq) GetPid() uint16 {
	return 103
}

func (proto *SkillTipsReq) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(103)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Actor) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Skill_id) {
		return false
	}

	return true
}

func (proto *SkillTipsReq) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Actor) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Skill_id) {
		return false
	}

	return true
}

func (proto *SkillTipsAck) GetMid() uint16 {
	return 102
}

func (proto *SkillTipsAck) GetPid() uint16 {
	return 104
}

func (proto *SkillTipsAck) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(104)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Skill_id) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Tips, uint8(255)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Errcode) {
		return false
	}

	if !ProtoWriteString(b, proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *SkillTipsAck) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Skill_id) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Tips, uint8(255)) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Errcode) {
		return false
	}

	if !ProtoReadString(b, &proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *TeamLeaderOprNtf) GetMid() uint16 {
	return 102
}

func (proto *TeamLeaderOprNtf) GetPid() uint16 {
	return 105
}

func (proto *TeamLeaderOprNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(105)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Opr) {
		return false
	}

	return true
}

func (proto *TeamLeaderOprNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Opr) {
		return false
	}

	return true
}

func (proto *KeepAliveCtrlNtf) GetMid() uint16 {
	return 102
}

func (proto *KeepAliveCtrlNtf) GetPid() uint16 {
	return 106
}

func (proto *KeepAliveCtrlNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(106)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Opt) {
		return false
	}

	return true
}

func (proto *KeepAliveCtrlNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Opt) {
		return false
	}

	return true
}

func (proto *NetDelayReq) GetMid() uint16 {
	return 102
}

func (proto *NetDelayReq) GetPid() uint16 {
	return 107
}

func (proto *NetDelayReq) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(107)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Gs_tick) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Gt_tick) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Sg_tick) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Ping) {
		return false
	}

	return true
}

func (proto *NetDelayReq) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Gs_tick) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Gt_tick) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Sg_tick) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Ping) {
		return false
	}

	return true
}

func (proto *NetDelayAck) GetMid() uint16 {
	return 102
}

func (proto *NetDelayAck) GetPid() uint16 {
	return 108
}

func (proto *NetDelayAck) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(108)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Gs_elpase) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Gt_elpase) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Sg_elpase) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Gt_cpu) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Baidu_delay) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Taobao_delay) {
		return false
	}

	if !ProtoWriteString(b, proto.Sg_id, 255) {
		return false
	}

	return true
}

func (proto *NetDelayAck) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Gs_elpase) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Gt_elpase) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Sg_elpase) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Gt_cpu) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Baidu_delay) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Taobao_delay) {
		return false
	}

	if !ProtoReadString(b, &proto.Sg_id, 255) {
		return false
	}

	return true
}

func (proto *SystemSetupNtf) GetMid() uint16 {
	return 102
}

func (proto *SystemSetupNtf) GetPid() uint16 {
	return 109
}

func (proto *SystemSetupNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(109)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Attrs, uint16(65535)) {
		return false
	}

	return true
}

func (proto *SystemSetupNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Attrs, uint16(65535)) {
		return false
	}

	return true
}

func (proto *BuffListNtf) GetMid() uint16 {
	return 102
}

func (proto *BuffListNtf) GetPid() uint16 {
	return 110
}

func (proto *BuffListNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(110)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Buffs, uint16(65535)) {
		return false
	}

	return true
}

func (proto *BuffListNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Buffs, uint16(65535)) {
		return false
	}

	return true
}

func (proto *BuffAddNtf) GetMid() uint16 {
	return 102
}

func (proto *BuffAddNtf) GetPid() uint16 {
	return 111
}

func (proto *BuffAddNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(111)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Buff) {
		return false
	}

	return true
}

func (proto *BuffAddNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Buff) {
		return false
	}

	return true
}

func (proto *BuffDurationNtf) GetMid() uint16 {
	return 102
}

func (proto *BuffDurationNtf) GetPid() uint16 {
	return 112
}

func (proto *BuffDurationNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(112)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Duration) {
		return false
	}

	return true
}

func (proto *BuffDurationNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Duration) {
		return false
	}

	return true
}

func (proto *BuffDynAttrNtf) GetMid() uint16 {
	return 102
}

func (proto *BuffDynAttrNtf) GetPid() uint16 {
	return 113
}

func (proto *BuffDynAttrNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(113)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Id) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Dyn_attrs, uint16(65535)) {
		return false
	}

	return true
}

func (proto *BuffDynAttrNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Id) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Dyn_attrs, uint16(65535)) {
		return false
	}

	return true
}

func (proto *StopBuff) GetMid() uint16 {
	return 102
}

func (proto *StopBuff) GetPid() uint16 {
	return 114
}

func (proto *StopBuff) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(114)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Id) {
		return false
	}

	return true
}

func (proto *StopBuff) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Id) {
		return false
	}

	return true
}

func (proto *BuffDelNtf) GetMid() uint16 {
	return 102
}

func (proto *BuffDelNtf) GetPid() uint16 {
	return 115
}

func (proto *BuffDelNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(115)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Id) {
		return false
	}

	return true
}

func (proto *BuffDelNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Id) {
		return false
	}

	return true
}

func (proto *ActivityListReq) GetMid() uint16 {
	return 102
}

func (proto *ActivityListReq) GetPid() uint16 {
	return 116
}

func (proto *ActivityListReq) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(116)) {
		return false
	}

	return true
}

func (proto *ActivityListReq) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	return true
}

func (proto *ActivityListAck) GetMid() uint16 {
	return 102
}

func (proto *ActivityListAck) GetPid() uint16 {
	return 117
}

func (proto *ActivityListAck) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(117)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Datas, uint16(65535)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Errcode) {
		return false
	}

	if !ProtoWriteString(b, proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *ActivityListAck) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Datas, uint16(65535)) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Errcode) {
		return false
	}

	if !ProtoReadString(b, &proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *ActivityJoinNtf) GetMid() uint16 {
	return 102
}

func (proto *ActivityJoinNtf) GetPid() uint16 {
	return 118
}

func (proto *ActivityJoinNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(118)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Id) {
		return false
	}

	return true
}

func (proto *ActivityJoinNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Id) {
		return false
	}

	return true
}

func (proto *GuildListNtf) GetMid() uint16 {
	return 102
}

func (proto *GuildListNtf) GetPid() uint16 {
	return 119
}

func (proto *GuildListNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(119)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Pos) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Count) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Guilds, uint16(65535)) {
		return false
	}

	return true
}

func (proto *GuildListNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pos) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Count) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Guilds, uint16(65535)) {
		return false
	}

	return true
}

func (proto *GuildNtf) GetMid() uint16 {
	return 102
}

func (proto *GuildNtf) GetPid() uint16 {
	return 120
}

func (proto *GuildNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(120)) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Guild) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Self) {
		return false
	}

	return true
}

func (proto *GuildNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Guild) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Self) {
		return false
	}

	return true
}

func (proto *GuildBaseNtf) GetMid() uint16 {
	return 102
}

func (proto *GuildBaseNtf) GetPid() uint16 {
	return 121
}

func (proto *GuildBaseNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(121)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteString(b, proto.Name, 32) {
		return false
	}

	return true
}

func (proto *GuildBaseNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadString(b, &proto.Name, 32) {
		return false
	}

	return true
}

func (proto *GuildNoticeNtf) GetMid() uint16 {
	return 102
}

func (proto *GuildNoticeNtf) GetPid() uint16 {
	return 122
}

func (proto *GuildNoticeNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(122)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteString(b, proto.Notice, 1024) {
		return false
	}

	return true
}

func (proto *GuildNoticeNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadString(b, &proto.Notice, 1024) {
		return false
	}

	return true
}

func (proto *GuildBriefNtf) GetMid() uint16 {
	return 102
}

func (proto *GuildBriefNtf) GetPid() uint16 {
	return 123
}

func (proto *GuildBriefNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(123)) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Guild) {
		return false
	}

	return true
}

func (proto *GuildBriefNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Guild) {
		return false
	}

	return true
}

func (proto *GuildMemberListNtf) GetMid() uint16 {
	return 102
}

func (proto *GuildMemberListNtf) GetPid() uint16 {
	return 124
}

func (proto *GuildMemberListNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(124)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Pos) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Count) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Members, uint16(65535)) {
		return false
	}

	return true
}

func (proto *GuildMemberListNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pos) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Count) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Members, uint16(65535)) {
		return false
	}

	return true
}

func (proto *GuildMemberNtf) GetMid() uint16 {
	return 102
}

func (proto *GuildMemberNtf) GetPid() uint16 {
	return 125
}

func (proto *GuildMemberNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(125)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Members, uint16(65535)) {
		return false
	}

	return true
}

func (proto *GuildMemberNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Members, uint16(65535)) {
		return false
	}

	return true
}

func (proto *GuildApplicantListNtf) GetMid() uint16 {
	return 102
}

func (proto *GuildApplicantListNtf) GetPid() uint16 {
	return 126
}

func (proto *GuildApplicantListNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(126)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Applicants, uint16(65535)) {
		return false
	}

	return true
}

func (proto *GuildApplicantListNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Applicants, uint16(65535)) {
		return false
	}

	return true
}

func (proto *GuildOperationNtf) GetMid() uint16 {
	return 102
}

func (proto *GuildOperationNtf) GetPid() uint16 {
	return 127
}

func (proto *GuildOperationNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(127)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Type) {
		return false
	}

	if !ProtoWriteString(b, proto.Param1, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Param2, 255) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Errcode) {
		return false
	}

	if !ProtoWriteString(b, proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *GuildOperationNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Type) {
		return false
	}

	if !ProtoReadString(b, &proto.Param1, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Param2, 255) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Errcode) {
		return false
	}

	if !ProtoReadString(b, &proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *GetTimestampReq) GetMid() uint16 {
	return 102
}

func (proto *GetTimestampReq) GetPid() uint16 {
	return 128
}

func (proto *GetTimestampReq) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(128)) {
		return false
	}

	return true
}

func (proto *GetTimestampReq) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	return true
}

func (proto *GetTimestampAck) GetMid() uint16 {
	return 102
}

func (proto *GetTimestampAck) GetPid() uint16 {
	return 129
}

func (proto *GetTimestampAck) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(129)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Now) {
		return false
	}

	return true
}

func (proto *GetTimestampAck) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Now) {
		return false
	}

	return true
}

func (proto *ContactListNtf) GetMid() uint16 {
	return 102
}

func (proto *ContactListNtf) GetPid() uint16 {
	return 130
}

func (proto *ContactListNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(130)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Contact_type) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Contacts, uint16(65535)) {
		return false
	}

	return true
}

func (proto *ContactListNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Contact_type) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Contacts, uint16(65535)) {
		return false
	}

	return true
}

func (proto *AddContactNtf) GetMid() uint16 {
	return 102
}

func (proto *AddContactNtf) GetPid() uint16 {
	return 131
}

func (proto *AddContactNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(131)) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Contact) {
		return false
	}

	return true
}

func (proto *AddContactNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Contact) {
		return false
	}

	return true
}

func (proto *UpdateContactNtf) GetMid() uint16 {
	return 102
}

func (proto *UpdateContactNtf) GetPid() uint16 {
	return 132
}

func (proto *UpdateContactNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(132)) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Contact) {
		return false
	}

	return true
}

func (proto *UpdateContactNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Contact) {
		return false
	}

	return true
}

func (proto *DelContactNtf) GetMid() uint16 {
	return 102
}

func (proto *DelContactNtf) GetPid() uint16 {
	return 133
}

func (proto *DelContactNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(133)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Target_guid) {
		return false
	}

	return true
}

func (proto *DelContactNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Target_guid) {
		return false
	}

	return true
}

func (proto *AddContactMessageNtf) GetMid() uint16 {
	return 102
}

func (proto *AddContactMessageNtf) GetPid() uint16 {
	return 134
}

func (proto *AddContactMessageNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(134)) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Messages) {
		return false
	}

	return true
}

func (proto *AddContactMessageNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Messages) {
		return false
	}

	return true
}

func (proto *ItemQueryNtf) GetMid() uint16 {
	return 102
}

func (proto *ItemQueryNtf) GetPid() uint16 {
	return 135
}

func (proto *ItemQueryNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(135)) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Item) {
		return false
	}

	return true
}

func (proto *ItemQueryNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Item) {
		return false
	}

	return true
}

func (proto *PetQueryNtf) GetMid() uint16 {
	return 102
}

func (proto *PetQueryNtf) GetPid() uint16 {
	return 136
}

func (proto *PetQueryNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(136)) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Pet) {
		return false
	}

	return true
}

func (proto *PetQueryNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Pet) {
		return false
	}

	return true
}

func (proto *ContactInfoNtf) GetMid() uint16 {
	return 102
}

func (proto *ContactInfoNtf) GetPid() uint16 {
	return 137
}

func (proto *ContactInfoNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(137)) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *ContactInfoNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *MailListNtf) GetMid() uint16 {
	return 102
}

func (proto *MailListNtf) GetPid() uint16 {
	return 138
}

func (proto *MailListNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(138)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Mails, uint8(255)) {
		return false
	}

	return true
}

func (proto *MailListNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Mails, uint8(255)) {
		return false
	}

	return true
}

func (proto *AddMailNtf) GetMid() uint16 {
	return 102
}

func (proto *AddMailNtf) GetPid() uint16 {
	return 139
}

func (proto *AddMailNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(139)) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Mail) {
		return false
	}

	return true
}

func (proto *AddMailNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Mail) {
		return false
	}

	return true
}

func (proto *DelMailNtf) GetMid() uint16 {
	return 102
}

func (proto *DelMailNtf) GetPid() uint16 {
	return 140
}

func (proto *DelMailNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(140)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Mail_guid) {
		return false
	}

	return true
}

func (proto *DelMailNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Mail_guid) {
		return false
	}

	return true
}

func (proto *MailBodyNtf) GetMid() uint16 {
	return 102
}

func (proto *MailBodyNtf) GetPid() uint16 {
	return 141
}

func (proto *MailBodyNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(141)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Mail_guid) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Mail) {
		return false
	}

	return true
}

func (proto *MailBodyNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Mail_guid) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Mail) {
		return false
	}

	return true
}

func (proto *UpdateMailBodyNtf) GetMid() uint16 {
	return 102
}

func (proto *UpdateMailBodyNtf) GetPid() uint16 {
	return 142
}

func (proto *UpdateMailBodyNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(142)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Mail_guid) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Mail) {
		return false
	}

	return true
}

func (proto *UpdateMailBodyNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Mail_guid) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Mail) {
		return false
	}

	return true
}

func (proto *UpdateMailHeadNtf) GetMid() uint16 {
	return 102
}

func (proto *UpdateMailHeadNtf) GetPid() uint16 {
	return 143
}

func (proto *UpdateMailHeadNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(143)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Mail_guid) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Mail) {
		return false
	}

	return true
}

func (proto *UpdateMailHeadNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Mail_guid) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Mail) {
		return false
	}

	return true
}

func (proto *RanklistReq) GetMid() uint16 {
	return 102
}

func (proto *RanklistReq) GetPid() uint16 {
	return 144
}

func (proto *RanklistReq) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(144)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Type) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Begin) {
		return false
	}

	if !ProtoWriteInteger(b, proto.End) {
		return false
	}

	return true
}

func (proto *RanklistReq) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Type) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Begin) {
		return false
	}

	if !ProtoReadInteger(b, &proto.End) {
		return false
	}

	return true
}

func (proto *RanklistAck) GetMid() uint16 {
	return 102
}

func (proto *RanklistAck) GetPid() uint16 {
	return 145
}

func (proto *RanklistAck) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(145)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Type) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Total) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Begin) {
		return false
	}

	if !ProtoWriteInteger(b, proto.End) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Data, uint16(65535)) {
		return false
	}

	return true
}

func (proto *RanklistAck) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Type) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Total) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Begin) {
		return false
	}

	if !ProtoReadInteger(b, &proto.End) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Data, uint16(65535)) {
		return false
	}

	return true
}

func (proto *GetRankReq) GetMid() uint16 {
	return 102
}

func (proto *GetRankReq) GetPid() uint16 {
	return 146
}

func (proto *GetRankReq) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(146)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Type) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	return true
}

func (proto *GetRankReq) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Type) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	return true
}

func (proto *GetRankAck) GetMid() uint16 {
	return 102
}

func (proto *GetRankAck) GetPid() uint16 {
	return 147
}

func (proto *GetRankAck) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(147)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Type) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Rank) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *GetRankAck) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Type) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Rank) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *TitleContainerNtf) GetMid() uint16 {
	return 102
}

func (proto *TitleContainerNtf) GetPid() uint16 {
	return 148
}

func (proto *TitleContainerNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(148)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Titles, uint16(65535)) {
		return false
	}

	return true
}

func (proto *TitleContainerNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Titles, uint16(65535)) {
		return false
	}

	return true
}

func (proto *TitleAddNtf) GetMid() uint16 {
	return 102
}

func (proto *TitleAddNtf) GetPid() uint16 {
	return 149
}

func (proto *TitleAddNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(149)) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Title) {
		return false
	}

	return true
}

func (proto *TitleAddNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Title) {
		return false
	}

	return true
}

func (proto *TitleDelNtf) GetMid() uint16 {
	return 102
}

func (proto *TitleDelNtf) GetPid() uint16 {
	return 150
}

func (proto *TitleDelNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(150)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Title_id) {
		return false
	}

	return true
}

func (proto *TitleDelNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Title_id) {
		return false
	}

	return true
}

func (proto *AgentKeyReq) GetMid() uint16 {
	return 102
}

func (proto *AgentKeyReq) GetPid() uint16 {
	return 151
}

func (proto *AgentKeyReq) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(151)) {
		return false
	}

	return true
}

func (proto *AgentKeyReq) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	return true
}

func (proto *AgentKeyAck) GetMid() uint16 {
	return 102
}

func (proto *AgentKeyAck) GetPid() uint16 {
	return 152
}

func (proto *AgentKeyAck) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(152)) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Key, uint8(255)) {
		return false
	}

	return true
}

func (proto *AgentKeyAck) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Key, uint8(255)) {
		return false
	}

	return true
}

func (proto *HeadMsgNtf) GetMid() uint16 {
	return 102
}

func (proto *HeadMsgNtf) GetPid() uint16 {
	return 153
}

func (proto *HeadMsgNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(153)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Msg, uint16(65535)) {
		return false
	}

	return true
}

func (proto *HeadMsgNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Msg, uint16(65535)) {
		return false
	}

	return true
}

func (proto *AutoContainerNtf) GetMid() uint16 {
	return 102
}

func (proto *AutoContainerNtf) GetPid() uint16 {
	return 154
}

func (proto *AutoContainerNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(154)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Is_auto) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Auto_skill) {
		return false
	}

	return true
}

func (proto *AutoContainerNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Is_auto) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Auto_skill) {
		return false
	}

	return true
}

func (proto *PlayerQueryNtf) GetMid() uint16 {
	return 102
}

func (proto *PlayerQueryNtf) GetPid() uint16 {
	return 155
}

func (proto *PlayerQueryNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(155)) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Brief) {
		return false
	}

	return true
}

func (proto *PlayerQueryNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Brief) {
		return false
	}

	return true
}

func (proto *UseAllItem) GetMid() uint16 {
	return 102
}

func (proto *UseAllItem) GetPid() uint16 {
	return 156
}

func (proto *UseAllItem) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(156)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Item_guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Item_amount) {
		return false
	}

	return true
}

func (proto *UseAllItem) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Item_guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Item_amount) {
		return false
	}

	return true
}

func (proto *GuardContainerNtf) GetMid() uint16 {
	return 102
}

func (proto *GuardContainerNtf) GetPid() uint16 {
	return 157
}

func (proto *GuardContainerNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(157)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Guards, uint16(65535)) {
		return false
	}

	return true
}

func (proto *GuardContainerNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Guards, uint16(65535)) {
		return false
	}

	return true
}

func (proto *GuardAddNtf) GetMid() uint16 {
	return 102
}

func (proto *GuardAddNtf) GetPid() uint16 {
	return 158
}

func (proto *GuardAddNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(158)) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Guard) {
		return false
	}

	return true
}

func (proto *GuardAddNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Guard) {
		return false
	}

	return true
}

func (proto *SetGuardLineup) GetMid() uint16 {
	return 102
}

func (proto *SetGuardLineup) GetPid() uint16 {
	return 159
}

func (proto *SetGuardLineup) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(159)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Lineup) {
		return false
	}

	return true
}

func (proto *SetGuardLineup) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Lineup) {
		return false
	}

	return true
}

func (proto *PetNewAddNtf) GetMid() uint16 {
	return 102
}

func (proto *PetNewAddNtf) GetPid() uint16 {
	return 160
}

func (proto *PetNewAddNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(160)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Pet_guid) {
		return false
	}

	return true
}

func (proto *PetNewAddNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pet_guid) {
		return false
	}

	return true
}

func (proto *TeamPlatformNtf) GetMid() uint16 {
	return 102
}

func (proto *TeamPlatformNtf) GetPid() uint16 {
	return 161
}

func (proto *TeamPlatformNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(161)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Teams, uint16(65535)) {
		return false
	}

	return true
}

func (proto *TeamPlatformNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Teams, uint16(65535)) {
		return false
	}

	return true
}

func (proto *TeamApplicantsNtf) GetMid() uint16 {
	return 102
}

func (proto *TeamApplicantsNtf) GetPid() uint16 {
	return 162
}

func (proto *TeamApplicantsNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(162)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Team_guid) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Applicants, uint16(65535)) {
		return false
	}

	return true
}

func (proto *TeamApplicantsNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Team_guid) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Applicants, uint16(65535)) {
		return false
	}

	return true
}

func (proto *TeamOperationNtf) GetMid() uint16 {
	return 102
}

func (proto *TeamOperationNtf) GetPid() uint16 {
	return 163
}

func (proto *TeamOperationNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(163)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Type) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Errcode) {
		return false
	}

	if !ProtoWriteString(b, proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *TeamOperationNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Type) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Errcode) {
		return false
	}

	if !ProtoReadString(b, &proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *TeamTargetNtf) GetMid() uint16 {
	return 102
}

func (proto *TeamTargetNtf) GetPid() uint16 {
	return 164
}

func (proto *TeamTargetNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(164)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Team_guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Target) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Permission) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Min_require_level) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Max_require_level) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Leader_guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Min_require_reinc) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Max_require_reinc) {
		return false
	}

	return true
}

func (proto *TeamTargetNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Team_guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Target) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Permission) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Min_require_level) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Max_require_level) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Leader_guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Min_require_reinc) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Max_require_reinc) {
		return false
	}

	return true
}

func (proto *ChangedNameNtf) GetMid() uint16 {
	return 102
}

func (proto *ChangedNameNtf) GetPid() uint16 {
	return 165
}

func (proto *ChangedNameNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(165)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteString(b, proto.Name, 255) {
		return false
	}

	return true
}

func (proto *ChangedNameNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadString(b, &proto.Name, 255) {
		return false
	}

	return true
}

func (proto *CustomDataNtf) GetMid() uint16 {
	return 102
}

func (proto *CustomDataNtf) GetPid() uint16 {
	return 166
}

func (proto *CustomDataNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(166)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Custom) {
		return false
	}

	return true
}

func (proto *CustomDataNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Custom) {
		return false
	}

	return true
}

func (proto *SpeedCheckNtf) GetMid() uint16 {
	return 102
}

func (proto *SpeedCheckNtf) GetPid() uint16 {
	return 167
}

func (proto *SpeedCheckNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(167)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Type) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Pct) {
		return false
	}

	return true
}

func (proto *SpeedCheckNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Type) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pct) {
		return false
	}

	return true
}

func (proto *ConsoleMsgNtf) GetMid() uint16 {
	return 102
}

func (proto *ConsoleMsgNtf) GetPid() uint16 {
	return 168
}

func (proto *ConsoleMsgNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(168)) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Msg, uint16(65535)) {
		return false
	}

	return true
}

func (proto *ConsoleMsgNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Msg, uint16(65535)) {
		return false
	}

	return true
}

func (proto *PetSwapNtf) GetMid() uint16 {
	return 102
}

func (proto *PetSwapNtf) GetPid() uint16 {
	return 169
}

func (proto *PetSwapNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(169)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Container_type) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Pet_guid1) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Pet_guid2) {
		return false
	}

	return true
}

func (proto *PetSwapNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Container_type) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pet_guid1) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pet_guid2) {
		return false
	}

	return true
}

func (proto *GuardDestroyNtf) GetMid() uint16 {
	return 102
}

func (proto *GuardDestroyNtf) GetPid() uint16 {
	return 170
}

func (proto *GuardDestroyNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(170)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	return true
}

func (proto *GuardDestroyNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	return true
}

func (proto *ActivateGuard) GetMid() uint16 {
	return 102
}

func (proto *ActivateGuard) GetPid() uint16 {
	return 171
}

func (proto *ActivateGuard) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(171)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guard_id) {
		return false
	}

	return true
}

func (proto *ActivateGuard) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guard_id) {
		return false
	}

	return true
}

func (proto *ReleaseGuard) GetMid() uint16 {
	return 102
}

func (proto *ReleaseGuard) GetPid() uint16 {
	return 172
}

func (proto *ReleaseGuard) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(172)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guard_guid) {
		return false
	}

	return true
}

func (proto *ReleaseGuard) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guard_guid) {
		return false
	}

	return true
}

func (proto *TeamMemberSwapNtf) GetMid() uint16 {
	return 102
}

func (proto *TeamMemberSwapNtf) GetPid() uint16 {
	return 173
}

func (proto *TeamMemberSwapNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(173)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Member_guid1) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Member_guid2) {
		return false
	}

	return true
}

func (proto *TeamMemberSwapNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Member_guid1) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Member_guid2) {
		return false
	}

	return true
}

func (proto *GuardSwapNtf) GetMid() uint16 {
	return 102
}

func (proto *GuardSwapNtf) GetPid() uint16 {
	return 174
}

func (proto *GuardSwapNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(174)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guard_guid1) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guard_guid2) {
		return false
	}

	return true
}

func (proto *GuardSwapNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guard_guid1) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guard_guid2) {
		return false
	}

	return true
}

func (proto *PetReplaceNtf) GetMid() uint16 {
	return 102
}

func (proto *PetReplaceNtf) GetPid() uint16 {
	return 175
}

func (proto *PetReplaceNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(175)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Container_type) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Dst_pet_guid) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Src_pet) {
		return false
	}

	return true
}

func (proto *PetReplaceNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Container_type) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Dst_pet_guid) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Src_pet) {
		return false
	}

	return true
}

func (proto *GuardAppearNtf) GetMid() uint16 {
	return 102
}

func (proto *GuardAppearNtf) GetPid() uint16 {
	return 176
}

func (proto *GuardAppearNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(176)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteString(b, proto.Name, 255) {
		return false
	}

	if !ProtoWriteInteger(b, proto.X) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Y) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Attrs, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Buffs, uint8(255)) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Custom) {
		return false
	}

	return true
}

func (proto *GuardAppearNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadString(b, &proto.Name, 255) {
		return false
	}

	if !ProtoReadInteger(b, &proto.X) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Y) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Attrs, uint16(65535)) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Buffs, uint8(255)) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Custom) {
		return false
	}

	return true
}

func (proto *InstructionContainerNtf) GetMid() uint16 {
	return 102
}

func (proto *InstructionContainerNtf) GetPid() uint16 {
	return 177
}

func (proto *InstructionContainerNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(177)) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Datas) {
		return false
	}

	return true
}

func (proto *InstructionContainerNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Datas) {
		return false
	}

	return true
}

func (proto *InstructionAddReq) GetMid() uint16 {
	return 102
}

func (proto *InstructionAddReq) GetPid() uint16 {
	return 178
}

func (proto *InstructionAddReq) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(178)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Type) {
		return false
	}

	if !ProtoWriteString(b, proto.Content, 32) {
		return false
	}

	return true
}

func (proto *InstructionAddReq) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Type) {
		return false
	}

	if !ProtoReadString(b, &proto.Content, 32) {
		return false
	}

	return true
}

func (proto *InstructionAddAck) GetMid() uint16 {
	return 102
}

func (proto *InstructionAddAck) GetPid() uint16 {
	return 179
}

func (proto *InstructionAddAck) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(179)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Type) {
		return false
	}

	if !ProtoWriteString(b, proto.Content, 32) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Errcode) {
		return false
	}

	if !ProtoWriteString(b, proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *InstructionAddAck) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Type) {
		return false
	}

	if !ProtoReadString(b, &proto.Content, 32) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Errcode) {
		return false
	}

	if !ProtoReadString(b, &proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *InstructionDeleteReq) GetMid() uint16 {
	return 102
}

func (proto *InstructionDeleteReq) GetPid() uint16 {
	return 180
}

func (proto *InstructionDeleteReq) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(180)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Type) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Index) {
		return false
	}

	return true
}

func (proto *InstructionDeleteReq) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Type) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Index) {
		return false
	}

	return true
}

func (proto *InstructionDeleteAck) GetMid() uint16 {
	return 102
}

func (proto *InstructionDeleteAck) GetPid() uint16 {
	return 181
}

func (proto *InstructionDeleteAck) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(181)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Type) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Index) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Errcode) {
		return false
	}

	if !ProtoWriteString(b, proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *InstructionDeleteAck) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Type) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Index) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Errcode) {
		return false
	}

	if !ProtoReadString(b, &proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *InstructionModfityReq) GetMid() uint16 {
	return 102
}

func (proto *InstructionModfityReq) GetPid() uint16 {
	return 182
}

func (proto *InstructionModfityReq) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(182)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Type) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Index) {
		return false
	}

	if !ProtoWriteString(b, proto.Content, 32) {
		return false
	}

	return true
}

func (proto *InstructionModfityReq) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Type) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Index) {
		return false
	}

	if !ProtoReadString(b, &proto.Content, 32) {
		return false
	}

	return true
}

func (proto *InstructionModifyAck) GetMid() uint16 {
	return 102
}

func (proto *InstructionModifyAck) GetPid() uint16 {
	return 183
}

func (proto *InstructionModifyAck) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(183)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Type) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Index) {
		return false
	}

	if !ProtoWriteString(b, proto.Content, 32) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Errcode) {
		return false
	}

	if !ProtoWriteString(b, proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *InstructionModifyAck) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Type) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Index) {
		return false
	}

	if !ProtoReadString(b, &proto.Content, 32) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Errcode) {
		return false
	}

	if !ProtoReadString(b, &proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *InstructionDefaultReq) GetMid() uint16 {
	return 102
}

func (proto *InstructionDefaultReq) GetPid() uint16 {
	return 184
}

func (proto *InstructionDefaultReq) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(184)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Type) {
		return false
	}

	return true
}

func (proto *InstructionDefaultReq) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Type) {
		return false
	}

	return true
}

func (proto *InstructionDefaultAck) GetMid() uint16 {
	return 102
}

func (proto *InstructionDefaultAck) GetPid() uint16 {
	return 185
}

func (proto *InstructionDefaultAck) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(185)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Type) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Datas, uint8(255)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Errcode) {
		return false
	}

	if !ProtoWriteString(b, proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *InstructionDefaultAck) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Type) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Datas, uint8(255)) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Errcode) {
		return false
	}

	if !ProtoReadString(b, &proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *InstructionAttachReq) GetMid() uint16 {
	return 102
}

func (proto *InstructionAttachReq) GetPid() uint16 {
	return 186
}

func (proto *InstructionAttachReq) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(186)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Type) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Index) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Target) {
		return false
	}

	return true
}

func (proto *InstructionAttachReq) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Type) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Index) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Target) {
		return false
	}

	return true
}

func (proto *InstructionAttachAck) GetMid() uint16 {
	return 102
}

func (proto *InstructionAttachAck) GetPid() uint16 {
	return 187
}

func (proto *InstructionAttachAck) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(187)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Type) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Index) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Target) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Errcode) {
		return false
	}

	if !ProtoWriteString(b, proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *InstructionAttachAck) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Type) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Index) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Target) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Errcode) {
		return false
	}

	if !ProtoReadString(b, &proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *InstructionAttachNtf) GetMid() uint16 {
	return 102
}

func (proto *InstructionAttachNtf) GetPid() uint16 {
	return 188
}

func (proto *InstructionAttachNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(188)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Fighter_id) {
		return false
	}

	if !ProtoWriteString(b, proto.Content, 32) {
		return false
	}

	return true
}

func (proto *InstructionAttachNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Fighter_id) {
		return false
	}

	if !ProtoReadString(b, &proto.Content, 32) {
		return false
	}

	return true
}

func (proto *InstructionDetachReq) GetMid() uint16 {
	return 102
}

func (proto *InstructionDetachReq) GetPid() uint16 {
	return 189
}

func (proto *InstructionDetachReq) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(189)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Target) {
		return false
	}

	return true
}

func (proto *InstructionDetachReq) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Target) {
		return false
	}

	return true
}

func (proto *InstructionDetachAck) GetMid() uint16 {
	return 102
}

func (proto *InstructionDetachAck) GetPid() uint16 {
	return 190
}

func (proto *InstructionDetachAck) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(190)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Target) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Errcode) {
		return false
	}

	if !ProtoWriteString(b, proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *InstructionDetachAck) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Target) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Errcode) {
		return false
	}

	if !ProtoReadString(b, &proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *InstructionDetachNtf) GetMid() uint16 {
	return 102
}

func (proto *InstructionDetachNtf) GetPid() uint16 {
	return 191
}

func (proto *InstructionDetachNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(191)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Fighter_id) {
		return false
	}

	return true
}

func (proto *InstructionDetachNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Fighter_id) {
		return false
	}

	return true
}

func (proto *PlayerDetailNtf) GetMid() uint16 {
	return 102
}

func (proto *PlayerDetailNtf) GetPid() uint16 {
	return 192
}

func (proto *PlayerDetailNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(192)) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *PlayerDetailNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *MapDynBlockPtNtf) GetMid() uint16 {
	return 102
}

func (proto *MapDynBlockPtNtf) GetPid() uint16 {
	return 193
}

func (proto *MapDynBlockPtNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(193)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Map_id) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Block_pt_list, uint16(65535)) {
		return false
	}

	return true
}

func (proto *MapDynBlockPtNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Map_id) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Block_pt_list, uint16(65535)) {
		return false
	}

	return true
}

func (proto *GuardQueryNtf) GetMid() uint16 {
	return 102
}

func (proto *GuardQueryNtf) GetPid() uint16 {
	return 194
}

func (proto *GuardQueryNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(194)) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Guard) {
		return false
	}

	return true
}

func (proto *GuardQueryNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Guard) {
		return false
	}

	return true
}

func (proto *BuyBackNtfEx) GetMid() uint16 {
	return 102
}

func (proto *BuyBackNtfEx) GetPid() uint16 {
	return 195
}

func (proto *BuyBackNtfEx) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(195)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Buy_guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Buy_amount) {
		return false
	}

	return true
}

func (proto *BuyBackNtfEx) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Buy_guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Buy_amount) {
		return false
	}

	return true
}

func (proto *GuildCustomNtf) GetMid() uint16 {
	return 102
}

func (proto *GuildCustomNtf) GetPid() uint16 {
	return 196
}

func (proto *GuildCustomNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(196)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Player_guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Is_guild) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Customs) {
		return false
	}

	return true
}

func (proto *GuildCustomNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Player_guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Is_guild) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Customs) {
		return false
	}

	return true
}

func (proto *PreTurnRoundNtf) GetMid() uint16 {
	return 102
}

func (proto *PreTurnRoundNtf) GetPid() uint16 {
	return 197
}

func (proto *PreTurnRoundNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(197)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Round) {
		return false
	}

	return true
}

func (proto *PreTurnRoundNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Round) {
		return false
	}

	return true
}

func (proto *FighterSpecialPetNtf) GetMid() uint16 {
	return 102
}

func (proto *FighterSpecialPetNtf) GetPid() uint16 {
	return 198
}

func (proto *FighterSpecialPetNtf) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, uint16(102)) {
		return false
	}

	if !ProtoWriteInteger(b, uint16(198)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Special_pet) {
		return false
	}

	return true
}

func (proto *FighterSpecialPetNtf) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Special_pet) {
		return false
	}

	return true
}

type IKeepAliveReq interface {
	OnKeepAliveReq(proto *KeepAliveReq)
}

type IKeepAliveAck interface {
	OnKeepAliveAck(proto *KeepAliveAck)
}

type IAttrNtf interface {
	OnAttrNtf(proto *AttrNtf)
}

type IPlayerAppearNtf interface {
	OnPlayerAppearNtf(proto *PlayerAppearNtf)
}

type INPCAppearNtf interface {
	OnNPCAppearNtf(proto *NPCAppearNtf)
}

type IItemAppearNtf interface {
	OnItemAppearNtf(proto *ItemAppearNtf)
}

type IObjDisAppearNtf interface {
	OnObjDisAppearNtf(proto *ObjDisAppearNtf)
}

type IObjMoveNtf interface {
	OnObjMoveNtf(proto *ObjMoveNtf)
}

type IEnterMapNtf interface {
	OnEnterMapNtf(proto *EnterMapNtf)
}

type IMoveReq interface {
	OnMoveReq(proto *MoveReq)
}

type IMoveAck interface {
	OnMoveAck(proto *MoveAck)
}

type IJumpMapReq interface {
	OnJumpMapReq(proto *JumpMapReq)
}

type IJumpMapAck interface {
	OnJumpMapAck(proto *JumpMapAck)
}

type IAddJumpMapRegionNtf interface {
	OnAddJumpMapRegionNtf(proto *AddJumpMapRegionNtf)
}

type IDelJumpMapRegionNtf interface {
	OnDelJumpMapRegionNtf(proto *DelJumpMapRegionNtf)
}

type IItemAddNtf interface {
	OnItemAddNtf(proto *ItemAddNtf)
}

type IItemUpdateNtf interface {
	OnItemUpdateNtf(proto *ItemUpdateNtf)
}

type IItemDestroyNtf interface {
	OnItemDestroyNtf(proto *ItemDestroyNtf)
}

type ITipsMsgNtf interface {
	OnTipsMsgNtf(proto *TipsMsgNtf)
}

type ITopMsgNtf interface {
	OnTopMsgNtf(proto *TopMsgNtf)
}

type ISysMsgNtf interface {
	OnSysMsgNtf(proto *SysMsgNtf)
}

type IPopupMsgNtf interface {
	OnPopupMsgNtf(proto *PopupMsgNtf)
}

type IItemContainerNtf interface {
	OnItemContainerNtf(proto *ItemContainerNtf)
}

type IItemContainerUpdateNtf interface {
	OnItemContainerUpdateNtf(proto *ItemContainerUpdateNtf)
}

type ISubmitForm interface {
	OnSubmitForm(proto *SubmitForm)
}

type IShowFormNtf interface {
	OnShowFormNtf(proto *ShowFormNtf)
}

type IExecuteGMReq interface {
	OnExecuteGMReq(proto *ExecuteGMReq)
}

type IFightBeginNtf interface {
	OnFightBeginNtf(proto *FightBeginNtf)
}

type ITurnRoundNtf interface {
	OnTurnRoundNtf(proto *TurnRoundNtf)
}

type IFightOperateListNtf interface {
	OnFightOperateListNtf(proto *FightOperateListNtf)
}

type IFightOperateReq interface {
	OnFightOperateReq(proto *FightOperateReq)
}

type IFightOperateAck interface {
	OnFightOperateAck(proto *FightOperateAck)
}

type IFightOperateNtf interface {
	OnFightOperateNtf(proto *FightOperateNtf)
}

type IFightDisplayNtf interface {
	OnFightDisplayNtf(proto *FightDisplayNtf)
}

type IFightDisplayCompleteNtf interface {
	OnFightDisplayCompleteNtf(proto *FightDisplayCompleteNtf)
}

type IFightAutoReq interface {
	OnFightAutoReq(proto *FightAutoReq)
}

type IFightAutoAck interface {
	OnFightAutoAck(proto *FightAutoAck)
}

type IFightAutoNtf interface {
	OnFightAutoNtf(proto *FightAutoNtf)
}

type IFightAutoSkillReq interface {
	OnFightAutoSkillReq(proto *FightAutoSkillReq)
}

type IFightAutoSkillAck interface {
	OnFightAutoSkillAck(proto *FightAutoSkillAck)
}

type IFightEndNtf interface {
	OnFightEndNtf(proto *FightEndNtf)
}

type IAddFighterNtf interface {
	OnAddFighterNtf(proto *AddFighterNtf)
}

type IDelFighterNtf interface {
	OnDelFighterNtf(proto *DelFighterNtf)
}

type IAddFightPetData interface {
	OnAddFightPetData(proto *AddFightPetData)
}

type IUpdateFightPetData interface {
	OnUpdateFightPetData(proto *UpdateFightPetData)
}

type IPlayerKillReq interface {
	OnPlayerKillReq(proto *PlayerKillReq)
}

type IPlayerKillAck interface {
	OnPlayerKillAck(proto *PlayerKillAck)
}

type IEnterFightViewReq interface {
	OnEnterFightViewReq(proto *EnterFightViewReq)
}

type IEnterFightViewAck interface {
	OnEnterFightViewAck(proto *EnterFightViewAck)
}

type ILeaveFightViewReq interface {
	OnLeaveFightViewReq(proto *LeaveFightViewReq)
}

type ILeaveFightViewAck interface {
	OnLeaveFightViewAck(proto *LeaveFightViewAck)
}

type ITeamLeaderOprReq interface {
	OnTeamLeaderOprReq(proto *TeamLeaderOprReq)
}

type ITeamNtf interface {
	OnTeamNtf(proto *TeamNtf)
}

type ITeamLeaderNtf interface {
	OnTeamLeaderNtf(proto *TeamLeaderNtf)
}

type ITeamDestroyNtf interface {
	OnTeamDestroyNtf(proto *TeamDestroyNtf)
}

type ITeamMemberNtf interface {
	OnTeamMemberNtf(proto *TeamMemberNtf)
}

type ITeamMemberLeaveNtf interface {
	OnTeamMemberLeaveNtf(proto *TeamMemberLeaveNtf)
}

type INpcSelectReq interface {
	OnNpcSelectReq(proto *NpcSelectReq)
}

type INpcSelectAck interface {
	OnNpcSelectAck(proto *NpcSelectAck)
}

type INpcTalkReq interface {
	OnNpcTalkReq(proto *NpcTalkReq)
}

type INpcTalkAck interface {
	OnNpcTalkAck(proto *NpcTalkAck)
}

type IInviteMsgNtf interface {
	OnInviteMsgNtf(proto *InviteMsgNtf)
}

type IReplyInvite interface {
	OnReplyInvite(proto *ReplyInvite)
}

type IMoveItem interface {
	OnMoveItem(proto *MoveItem)
}

type IUseItem interface {
	OnUseItem(proto *UseItem)
}

type IRearrangeItem interface {
	OnRearrangeItem(proto *RearrangeItem)
}

type ISkillContainerNtf interface {
	OnSkillContainerNtf(proto *SkillContainerNtf)
}

type IAddSkillNtf interface {
	OnAddSkillNtf(proto *AddSkillNtf)
}

type IUpdateSkillNtf interface {
	OnUpdateSkillNtf(proto *UpdateSkillNtf)
}

type IDelSkillNtf interface {
	OnDelSkillNtf(proto *DelSkillNtf)
}

type IPetAppearNtf interface {
	OnPetAppearNtf(proto *PetAppearNtf)
}

type IPetContainerNtf interface {
	OnPetContainerNtf(proto *PetContainerNtf)
}

type IPetContainerUpdateNtf interface {
	OnPetContainerUpdateNtf(proto *PetContainerUpdateNtf)
}

type IPetAddNtf interface {
	OnPetAddNtf(proto *PetAddNtf)
}

type IPetDestroyNtf interface {
	OnPetDestroyNtf(proto *PetDestroyNtf)
}

type ISetPetLineup interface {
	OnSetPetLineup(proto *SetPetLineup)
}

type IShowPet interface {
	OnShowPet(proto *ShowPet)
}

type IReleasePet interface {
	OnReleasePet(proto *ReleasePet)
}

type IMovePet interface {
	OnMovePet(proto *MovePet)
}

type IShopOpenNtf interface {
	OnShopOpenNtf(proto *ShopOpenNtf)
}

type IShopBuyNtf interface {
	OnShopBuyNtf(proto *ShopBuyNtf)
}

type ISellNtf interface {
	OnSellNtf(proto *SellNtf)
}

type IBuyBackNtf interface {
	OnBuyBackNtf(proto *BuyBackNtf)
}

type IBuyBackListNtf interface {
	OnBuyBackListNtf(proto *BuyBackListNtf)
}

type ITeamAttrNtf interface {
	OnTeamAttrNtf(proto *TeamAttrNtf)
}

type ITipsMsgExNtf interface {
	OnTipsMsgExNtf(proto *TipsMsgExNtf)
}

type IItemNewAddNtf interface {
	OnItemNewAddNtf(proto *ItemNewAddNtf)
}

type IQuestContainerNtf interface {
	OnQuestContainerNtf(proto *QuestContainerNtf)
}

type IQuestAddNtf interface {
	OnQuestAddNtf(proto *QuestAddNtf)
}

type IQuestStateNtf interface {
	OnQuestStateNtf(proto *QuestStateNtf)
}

type IQuestNpcStateNtf interface {
	OnQuestNpcStateNtf(proto *QuestNpcStateNtf)
}

type IQuestTrackCountNtf interface {
	OnQuestTrackCountNtf(proto *QuestTrackCountNtf)
}

type IQuestTalkSelectReq interface {
	OnQuestTalkSelectReq(proto *QuestTalkSelectReq)
}

type IQuestTalkSelectAck interface {
	OnQuestTalkSelectAck(proto *QuestTalkSelectAck)
}

type IObjStrCustomNtf interface {
	OnObjStrCustomNtf(proto *ObjStrCustomNtf)
}

type IObjIntCustomNtf interface {
	OnObjIntCustomNtf(proto *ObjIntCustomNtf)
}

type IObjDynAttrNtf interface {
	OnObjDynAttrNtf(proto *ObjDynAttrNtf)
}

type ICommitQuestItemRsp interface {
	OnCommitQuestItemRsp(proto *CommitQuestItemRsp)
}

type ICommitQuestItemNtf interface {
	OnCommitQuestItemNtf(proto *CommitQuestItemNtf)
}

type ICommitQuestPetRsp interface {
	OnCommitQuestPetRsp(proto *CommitQuestPetRsp)
}

type ICommitQuestPetNtf interface {
	OnCommitQuestPetNtf(proto *CommitQuestPetNtf)
}

type IAbandonQuest interface {
	OnAbandonQuest(proto *AbandonQuest)
}

type ISkillTipsReq interface {
	OnSkillTipsReq(proto *SkillTipsReq)
}

type ISkillTipsAck interface {
	OnSkillTipsAck(proto *SkillTipsAck)
}

type ITeamLeaderOprNtf interface {
	OnTeamLeaderOprNtf(proto *TeamLeaderOprNtf)
}

type IKeepAliveCtrlNtf interface {
	OnKeepAliveCtrlNtf(proto *KeepAliveCtrlNtf)
}

type INetDelayReq interface {
	OnNetDelayReq(proto *NetDelayReq)
}

type INetDelayAck interface {
	OnNetDelayAck(proto *NetDelayAck)
}

type ISystemSetupNtf interface {
	OnSystemSetupNtf(proto *SystemSetupNtf)
}

type IBuffListNtf interface {
	OnBuffListNtf(proto *BuffListNtf)
}

type IBuffAddNtf interface {
	OnBuffAddNtf(proto *BuffAddNtf)
}

type IBuffDurationNtf interface {
	OnBuffDurationNtf(proto *BuffDurationNtf)
}

type IBuffDynAttrNtf interface {
	OnBuffDynAttrNtf(proto *BuffDynAttrNtf)
}

type IStopBuff interface {
	OnStopBuff(proto *StopBuff)
}

type IBuffDelNtf interface {
	OnBuffDelNtf(proto *BuffDelNtf)
}

type IActivityListReq interface {
	OnActivityListReq(proto *ActivityListReq)
}

type IActivityListAck interface {
	OnActivityListAck(proto *ActivityListAck)
}

type IActivityJoinNtf interface {
	OnActivityJoinNtf(proto *ActivityJoinNtf)
}

type IGuildListNtf interface {
	OnGuildListNtf(proto *GuildListNtf)
}

type IGuildNtf interface {
	OnGuildNtf(proto *GuildNtf)
}

type IGuildBaseNtf interface {
	OnGuildBaseNtf(proto *GuildBaseNtf)
}

type IGuildNoticeNtf interface {
	OnGuildNoticeNtf(proto *GuildNoticeNtf)
}

type IGuildBriefNtf interface {
	OnGuildBriefNtf(proto *GuildBriefNtf)
}

type IGuildMemberListNtf interface {
	OnGuildMemberListNtf(proto *GuildMemberListNtf)
}

type IGuildMemberNtf interface {
	OnGuildMemberNtf(proto *GuildMemberNtf)
}

type IGuildApplicantListNtf interface {
	OnGuildApplicantListNtf(proto *GuildApplicantListNtf)
}

type IGuildOperationNtf interface {
	OnGuildOperationNtf(proto *GuildOperationNtf)
}

type IGetTimestampReq interface {
	OnGetTimestampReq(proto *GetTimestampReq)
}

type IGetTimestampAck interface {
	OnGetTimestampAck(proto *GetTimestampAck)
}

type IContactListNtf interface {
	OnContactListNtf(proto *ContactListNtf)
}

type IAddContactNtf interface {
	OnAddContactNtf(proto *AddContactNtf)
}

type IUpdateContactNtf interface {
	OnUpdateContactNtf(proto *UpdateContactNtf)
}

type IDelContactNtf interface {
	OnDelContactNtf(proto *DelContactNtf)
}

type IAddContactMessageNtf interface {
	OnAddContactMessageNtf(proto *AddContactMessageNtf)
}

type IItemQueryNtf interface {
	OnItemQueryNtf(proto *ItemQueryNtf)
}

type IPetQueryNtf interface {
	OnPetQueryNtf(proto *PetQueryNtf)
}

type IContactInfoNtf interface {
	OnContactInfoNtf(proto *ContactInfoNtf)
}

type IMailListNtf interface {
	OnMailListNtf(proto *MailListNtf)
}

type IAddMailNtf interface {
	OnAddMailNtf(proto *AddMailNtf)
}

type IDelMailNtf interface {
	OnDelMailNtf(proto *DelMailNtf)
}

type IMailBodyNtf interface {
	OnMailBodyNtf(proto *MailBodyNtf)
}

type IUpdateMailBodyNtf interface {
	OnUpdateMailBodyNtf(proto *UpdateMailBodyNtf)
}

type IUpdateMailHeadNtf interface {
	OnUpdateMailHeadNtf(proto *UpdateMailHeadNtf)
}

type IRanklistReq interface {
	OnRanklistReq(proto *RanklistReq)
}

type IRanklistAck interface {
	OnRanklistAck(proto *RanklistAck)
}

type IGetRankReq interface {
	OnGetRankReq(proto *GetRankReq)
}

type IGetRankAck interface {
	OnGetRankAck(proto *GetRankAck)
}

type ITitleContainerNtf interface {
	OnTitleContainerNtf(proto *TitleContainerNtf)
}

type ITitleAddNtf interface {
	OnTitleAddNtf(proto *TitleAddNtf)
}

type ITitleDelNtf interface {
	OnTitleDelNtf(proto *TitleDelNtf)
}

type IAgentKeyReq interface {
	OnAgentKeyReq(proto *AgentKeyReq)
}

type IAgentKeyAck interface {
	OnAgentKeyAck(proto *AgentKeyAck)
}

type IHeadMsgNtf interface {
	OnHeadMsgNtf(proto *HeadMsgNtf)
}

type IAutoContainerNtf interface {
	OnAutoContainerNtf(proto *AutoContainerNtf)
}

type IPlayerQueryNtf interface {
	OnPlayerQueryNtf(proto *PlayerQueryNtf)
}

type IUseAllItem interface {
	OnUseAllItem(proto *UseAllItem)
}

type IGuardContainerNtf interface {
	OnGuardContainerNtf(proto *GuardContainerNtf)
}

type IGuardAddNtf interface {
	OnGuardAddNtf(proto *GuardAddNtf)
}

type ISetGuardLineup interface {
	OnSetGuardLineup(proto *SetGuardLineup)
}

type IPetNewAddNtf interface {
	OnPetNewAddNtf(proto *PetNewAddNtf)
}

type ITeamPlatformNtf interface {
	OnTeamPlatformNtf(proto *TeamPlatformNtf)
}

type ITeamApplicantsNtf interface {
	OnTeamApplicantsNtf(proto *TeamApplicantsNtf)
}

type ITeamOperationNtf interface {
	OnTeamOperationNtf(proto *TeamOperationNtf)
}

type ITeamTargetNtf interface {
	OnTeamTargetNtf(proto *TeamTargetNtf)
}

type IChangedNameNtf interface {
	OnChangedNameNtf(proto *ChangedNameNtf)
}

type ICustomDataNtf interface {
	OnCustomDataNtf(proto *CustomDataNtf)
}

type ISpeedCheckNtf interface {
	OnSpeedCheckNtf(proto *SpeedCheckNtf)
}

type IConsoleMsgNtf interface {
	OnConsoleMsgNtf(proto *ConsoleMsgNtf)
}

type IPetSwapNtf interface {
	OnPetSwapNtf(proto *PetSwapNtf)
}

type IGuardDestroyNtf interface {
	OnGuardDestroyNtf(proto *GuardDestroyNtf)
}

type IActivateGuard interface {
	OnActivateGuard(proto *ActivateGuard)
}

type IReleaseGuard interface {
	OnReleaseGuard(proto *ReleaseGuard)
}

type ITeamMemberSwapNtf interface {
	OnTeamMemberSwapNtf(proto *TeamMemberSwapNtf)
}

type IGuardSwapNtf interface {
	OnGuardSwapNtf(proto *GuardSwapNtf)
}

type IPetReplaceNtf interface {
	OnPetReplaceNtf(proto *PetReplaceNtf)
}

type IGuardAppearNtf interface {
	OnGuardAppearNtf(proto *GuardAppearNtf)
}

type IInstructionContainerNtf interface {
	OnInstructionContainerNtf(proto *InstructionContainerNtf)
}

type IInstructionAddReq interface {
	OnInstructionAddReq(proto *InstructionAddReq)
}

type IInstructionAddAck interface {
	OnInstructionAddAck(proto *InstructionAddAck)
}

type IInstructionDeleteReq interface {
	OnInstructionDeleteReq(proto *InstructionDeleteReq)
}

type IInstructionDeleteAck interface {
	OnInstructionDeleteAck(proto *InstructionDeleteAck)
}

type IInstructionModfityReq interface {
	OnInstructionModfityReq(proto *InstructionModfityReq)
}

type IInstructionModifyAck interface {
	OnInstructionModifyAck(proto *InstructionModifyAck)
}

type IInstructionDefaultReq interface {
	OnInstructionDefaultReq(proto *InstructionDefaultReq)
}

type IInstructionDefaultAck interface {
	OnInstructionDefaultAck(proto *InstructionDefaultAck)
}

type IInstructionAttachReq interface {
	OnInstructionAttachReq(proto *InstructionAttachReq)
}

type IInstructionAttachAck interface {
	OnInstructionAttachAck(proto *InstructionAttachAck)
}

type IInstructionAttachNtf interface {
	OnInstructionAttachNtf(proto *InstructionAttachNtf)
}

type IInstructionDetachReq interface {
	OnInstructionDetachReq(proto *InstructionDetachReq)
}

type IInstructionDetachAck interface {
	OnInstructionDetachAck(proto *InstructionDetachAck)
}

type IInstructionDetachNtf interface {
	OnInstructionDetachNtf(proto *InstructionDetachNtf)
}

type IPlayerDetailNtf interface {
	OnPlayerDetailNtf(proto *PlayerDetailNtf)
}

type IMapDynBlockPtNtf interface {
	OnMapDynBlockPtNtf(proto *MapDynBlockPtNtf)
}

type IGuardQueryNtf interface {
	OnGuardQueryNtf(proto *GuardQueryNtf)
}

type IBuyBackNtfEx interface {
	OnBuyBackNtfEx(proto *BuyBackNtfEx)
}

type IGuildCustomNtf interface {
	OnGuildCustomNtf(proto *GuildCustomNtf)
}

type IPreTurnRoundNtf interface {
	OnPreTurnRoundNtf(proto *PreTurnRoundNtf)
}

type IFighterSpecialPetNtf interface {
	OnFighterSpecialPetNtf(proto *FighterSpecialPetNtf)
}

type ClientGS struct {
	protoDispatch interface{}
}

func NewClientGS[T any](dispatch *T) *ClientGS {
	return &ClientGS{dispatch}
}

func (protos *ClientGS) GetMid() uint16 {
	return 102
}

func (protos *ClientGS) DispatchProto(data []byte) bool {
	b := bytes.NewBuffer(data)

	mid := binary.LittleEndian.Uint16(data)
	if mid != protos.GetMid() {
		return false
	}

	pid := binary.LittleEndian.Uint16(data[unsafe.Sizeof(uint16(0)):])
	switch pid {
	case 1:
		{
			t, ok := protos.protoDispatch.(IKeepAliveReq)
			if !ok {
				return false
			}

			proto := &KeepAliveReq{}
			if !proto.Read(b) {
				fmt.Println("read KeepAliveReq fail, system error.")
				return false
			}

			t.OnKeepAliveReq(proto)
		}
	case 2:
		{
			t, ok := protos.protoDispatch.(IKeepAliveAck)
			if !ok {
				return false
			}

			proto := &KeepAliveAck{}
			if !proto.Read(b) {
				fmt.Println("read KeepAliveAck fail, system error.")
				return false
			}

			t.OnKeepAliveAck(proto)
		}
	case 3:
		{
			t, ok := protos.protoDispatch.(IAttrNtf)
			if !ok {
				return false
			}

			proto := &AttrNtf{}
			if !proto.Read(b) {
				fmt.Println("read AttrNtf fail, system error.")
				return false
			}

			t.OnAttrNtf(proto)
		}
	case 4:
		{
			t, ok := protos.protoDispatch.(IPlayerAppearNtf)
			if !ok {
				return false
			}

			proto := &PlayerAppearNtf{}
			if !proto.Read(b) {
				fmt.Println("read PlayerAppearNtf fail, system error.")
				return false
			}

			t.OnPlayerAppearNtf(proto)
		}
	case 5:
		{
			t, ok := protos.protoDispatch.(INPCAppearNtf)
			if !ok {
				return false
			}

			proto := &NPCAppearNtf{}
			if !proto.Read(b) {
				fmt.Println("read NPCAppearNtf fail, system error.")
				return false
			}

			t.OnNPCAppearNtf(proto)
		}
	case 6:
		{
			t, ok := protos.protoDispatch.(IItemAppearNtf)
			if !ok {
				return false
			}

			proto := &ItemAppearNtf{}
			if !proto.Read(b) {
				fmt.Println("read ItemAppearNtf fail, system error.")
				return false
			}

			t.OnItemAppearNtf(proto)
		}
	case 7:
		{
			t, ok := protos.protoDispatch.(IObjDisAppearNtf)
			if !ok {
				return false
			}

			proto := &ObjDisAppearNtf{}
			if !proto.Read(b) {
				fmt.Println("read ObjDisAppearNtf fail, system error.")
				return false
			}

			t.OnObjDisAppearNtf(proto)
		}
	case 8:
		{
			t, ok := protos.protoDispatch.(IObjMoveNtf)
			if !ok {
				return false
			}

			proto := &ObjMoveNtf{}
			if !proto.Read(b) {
				fmt.Println("read ObjMoveNtf fail, system error.")
				return false
			}

			t.OnObjMoveNtf(proto)
		}
	case 9:
		{
			t, ok := protos.protoDispatch.(IEnterMapNtf)
			if !ok {
				return false
			}

			proto := &EnterMapNtf{}
			if !proto.Read(b) {
				fmt.Println("read EnterMapNtf fail, system error.")
				return false
			}

			t.OnEnterMapNtf(proto)
		}
	case 10:
		{
			t, ok := protos.protoDispatch.(IMoveReq)
			if !ok {
				return false
			}

			proto := &MoveReq{}
			if !proto.Read(b) {
				fmt.Println("read MoveReq fail, system error.")
				return false
			}

			t.OnMoveReq(proto)
		}
	case 11:
		{
			t, ok := protos.protoDispatch.(IMoveAck)
			if !ok {
				return false
			}

			proto := &MoveAck{}
			if !proto.Read(b) {
				fmt.Println("read MoveAck fail, system error.")
				return false
			}

			t.OnMoveAck(proto)
		}
	case 12:
		{
			t, ok := protos.protoDispatch.(IJumpMapReq)
			if !ok {
				return false
			}

			proto := &JumpMapReq{}
			if !proto.Read(b) {
				fmt.Println("read JumpMapReq fail, system error.")
				return false
			}

			t.OnJumpMapReq(proto)
		}
	case 13:
		{
			t, ok := protos.protoDispatch.(IJumpMapAck)
			if !ok {
				return false
			}

			proto := &JumpMapAck{}
			if !proto.Read(b) {
				fmt.Println("read JumpMapAck fail, system error.")
				return false
			}

			t.OnJumpMapAck(proto)
		}
	case 14:
		{
			t, ok := protos.protoDispatch.(IAddJumpMapRegionNtf)
			if !ok {
				return false
			}

			proto := &AddJumpMapRegionNtf{}
			if !proto.Read(b) {
				fmt.Println("read AddJumpMapRegionNtf fail, system error.")
				return false
			}

			t.OnAddJumpMapRegionNtf(proto)
		}
	case 15:
		{
			t, ok := protos.protoDispatch.(IDelJumpMapRegionNtf)
			if !ok {
				return false
			}

			proto := &DelJumpMapRegionNtf{}
			if !proto.Read(b) {
				fmt.Println("read DelJumpMapRegionNtf fail, system error.")
				return false
			}

			t.OnDelJumpMapRegionNtf(proto)
		}
	case 16:
		{
			t, ok := protos.protoDispatch.(IItemAddNtf)
			if !ok {
				return false
			}

			proto := &ItemAddNtf{}
			if !proto.Read(b) {
				fmt.Println("read ItemAddNtf fail, system error.")
				return false
			}

			t.OnItemAddNtf(proto)
		}
	case 17:
		{
			t, ok := protos.protoDispatch.(IItemUpdateNtf)
			if !ok {
				return false
			}

			proto := &ItemUpdateNtf{}
			if !proto.Read(b) {
				fmt.Println("read ItemUpdateNtf fail, system error.")
				return false
			}

			t.OnItemUpdateNtf(proto)
		}
	case 18:
		{
			t, ok := protos.protoDispatch.(IItemDestroyNtf)
			if !ok {
				return false
			}

			proto := &ItemDestroyNtf{}
			if !proto.Read(b) {
				fmt.Println("read ItemDestroyNtf fail, system error.")
				return false
			}

			t.OnItemDestroyNtf(proto)
		}
	case 19:
		{
			t, ok := protos.protoDispatch.(ITipsMsgNtf)
			if !ok {
				return false
			}

			proto := &TipsMsgNtf{}
			if !proto.Read(b) {
				fmt.Println("read TipsMsgNtf fail, system error.")
				return false
			}

			t.OnTipsMsgNtf(proto)
		}
	case 20:
		{
			t, ok := protos.protoDispatch.(ITopMsgNtf)
			if !ok {
				return false
			}

			proto := &TopMsgNtf{}
			if !proto.Read(b) {
				fmt.Println("read TopMsgNtf fail, system error.")
				return false
			}

			t.OnTopMsgNtf(proto)
		}
	case 21:
		{
			t, ok := protos.protoDispatch.(ISysMsgNtf)
			if !ok {
				return false
			}

			proto := &SysMsgNtf{}
			if !proto.Read(b) {
				fmt.Println("read SysMsgNtf fail, system error.")
				return false
			}

			t.OnSysMsgNtf(proto)
		}
	case 22:
		{
			t, ok := protos.protoDispatch.(IPopupMsgNtf)
			if !ok {
				return false
			}

			proto := &PopupMsgNtf{}
			if !proto.Read(b) {
				fmt.Println("read PopupMsgNtf fail, system error.")
				return false
			}

			t.OnPopupMsgNtf(proto)
		}
	case 23:
		{
			t, ok := protos.protoDispatch.(IItemContainerNtf)
			if !ok {
				return false
			}

			proto := &ItemContainerNtf{}
			if !proto.Read(b) {
				fmt.Println("read ItemContainerNtf fail, system error.")
				return false
			}

			t.OnItemContainerNtf(proto)
		}
	case 24:
		{
			t, ok := protos.protoDispatch.(IItemContainerUpdateNtf)
			if !ok {
				return false
			}

			proto := &ItemContainerUpdateNtf{}
			if !proto.Read(b) {
				fmt.Println("read ItemContainerUpdateNtf fail, system error.")
				return false
			}

			t.OnItemContainerUpdateNtf(proto)
		}
	case 25:
		{
			t, ok := protos.protoDispatch.(ISubmitForm)
			if !ok {
				return false
			}

			proto := &SubmitForm{}
			if !proto.Read(b) {
				fmt.Println("read SubmitForm fail, system error.")
				return false
			}

			t.OnSubmitForm(proto)
		}
	case 26:
		{
			t, ok := protos.protoDispatch.(IShowFormNtf)
			if !ok {
				return false
			}

			proto := &ShowFormNtf{}
			if !proto.Read(b) {
				fmt.Println("read ShowFormNtf fail, system error.")
				return false
			}

			t.OnShowFormNtf(proto)
		}
	case 27:
		{
			t, ok := protos.protoDispatch.(IExecuteGMReq)
			if !ok {
				return false
			}

			proto := &ExecuteGMReq{}
			if !proto.Read(b) {
				fmt.Println("read ExecuteGMReq fail, system error.")
				return false
			}

			t.OnExecuteGMReq(proto)
		}
	case 28:
		{
			t, ok := protos.protoDispatch.(IFightBeginNtf)
			if !ok {
				return false
			}

			proto := &FightBeginNtf{}
			if !proto.Read(b) {
				fmt.Println("read FightBeginNtf fail, system error.")
				return false
			}

			t.OnFightBeginNtf(proto)
		}
	case 29:
		{
			t, ok := protos.protoDispatch.(ITurnRoundNtf)
			if !ok {
				return false
			}

			proto := &TurnRoundNtf{}
			if !proto.Read(b) {
				fmt.Println("read TurnRoundNtf fail, system error.")
				return false
			}

			t.OnTurnRoundNtf(proto)
		}
	case 30:
		{
			t, ok := protos.protoDispatch.(IFightOperateListNtf)
			if !ok {
				return false
			}

			proto := &FightOperateListNtf{}
			if !proto.Read(b) {
				fmt.Println("read FightOperateListNtf fail, system error.")
				return false
			}

			t.OnFightOperateListNtf(proto)
		}
	case 31:
		{
			t, ok := protos.protoDispatch.(IFightOperateReq)
			if !ok {
				return false
			}

			proto := &FightOperateReq{}
			if !proto.Read(b) {
				fmt.Println("read FightOperateReq fail, system error.")
				return false
			}

			t.OnFightOperateReq(proto)
		}
	case 32:
		{
			t, ok := protos.protoDispatch.(IFightOperateAck)
			if !ok {
				return false
			}

			proto := &FightOperateAck{}
			if !proto.Read(b) {
				fmt.Println("read FightOperateAck fail, system error.")
				return false
			}

			t.OnFightOperateAck(proto)
		}
	case 33:
		{
			t, ok := protos.protoDispatch.(IFightOperateNtf)
			if !ok {
				return false
			}

			proto := &FightOperateNtf{}
			if !proto.Read(b) {
				fmt.Println("read FightOperateNtf fail, system error.")
				return false
			}

			t.OnFightOperateNtf(proto)
		}
	case 34:
		{
			t, ok := protos.protoDispatch.(IFightDisplayNtf)
			if !ok {
				return false
			}

			proto := &FightDisplayNtf{}
			if !proto.Read(b) {
				fmt.Println("read FightDisplayNtf fail, system error.")
				return false
			}

			t.OnFightDisplayNtf(proto)
		}
	case 35:
		{
			t, ok := protos.protoDispatch.(IFightDisplayCompleteNtf)
			if !ok {
				return false
			}

			proto := &FightDisplayCompleteNtf{}
			if !proto.Read(b) {
				fmt.Println("read FightDisplayCompleteNtf fail, system error.")
				return false
			}

			t.OnFightDisplayCompleteNtf(proto)
		}
	case 36:
		{
			t, ok := protos.protoDispatch.(IFightAutoReq)
			if !ok {
				return false
			}

			proto := &FightAutoReq{}
			if !proto.Read(b) {
				fmt.Println("read FightAutoReq fail, system error.")
				return false
			}

			t.OnFightAutoReq(proto)
		}
	case 37:
		{
			t, ok := protos.protoDispatch.(IFightAutoAck)
			if !ok {
				return false
			}

			proto := &FightAutoAck{}
			if !proto.Read(b) {
				fmt.Println("read FightAutoAck fail, system error.")
				return false
			}

			t.OnFightAutoAck(proto)
		}
	case 38:
		{
			t, ok := protos.protoDispatch.(IFightAutoNtf)
			if !ok {
				return false
			}

			proto := &FightAutoNtf{}
			if !proto.Read(b) {
				fmt.Println("read FightAutoNtf fail, system error.")
				return false
			}

			t.OnFightAutoNtf(proto)
		}
	case 39:
		{
			t, ok := protos.protoDispatch.(IFightAutoSkillReq)
			if !ok {
				return false
			}

			proto := &FightAutoSkillReq{}
			if !proto.Read(b) {
				fmt.Println("read FightAutoSkillReq fail, system error.")
				return false
			}

			t.OnFightAutoSkillReq(proto)
		}
	case 40:
		{
			t, ok := protos.protoDispatch.(IFightAutoSkillAck)
			if !ok {
				return false
			}

			proto := &FightAutoSkillAck{}
			if !proto.Read(b) {
				fmt.Println("read FightAutoSkillAck fail, system error.")
				return false
			}

			t.OnFightAutoSkillAck(proto)
		}
	case 41:
		{
			t, ok := protos.protoDispatch.(IFightEndNtf)
			if !ok {
				return false
			}

			proto := &FightEndNtf{}
			if !proto.Read(b) {
				fmt.Println("read FightEndNtf fail, system error.")
				return false
			}

			t.OnFightEndNtf(proto)
		}
	case 42:
		{
			t, ok := protos.protoDispatch.(IAddFighterNtf)
			if !ok {
				return false
			}

			proto := &AddFighterNtf{}
			if !proto.Read(b) {
				fmt.Println("read AddFighterNtf fail, system error.")
				return false
			}

			t.OnAddFighterNtf(proto)
		}
	case 43:
		{
			t, ok := protos.protoDispatch.(IDelFighterNtf)
			if !ok {
				return false
			}

			proto := &DelFighterNtf{}
			if !proto.Read(b) {
				fmt.Println("read DelFighterNtf fail, system error.")
				return false
			}

			t.OnDelFighterNtf(proto)
		}
	case 44:
		{
			t, ok := protos.protoDispatch.(IAddFightPetData)
			if !ok {
				return false
			}

			proto := &AddFightPetData{}
			if !proto.Read(b) {
				fmt.Println("read AddFightPetData fail, system error.")
				return false
			}

			t.OnAddFightPetData(proto)
		}
	case 45:
		{
			t, ok := protos.protoDispatch.(IUpdateFightPetData)
			if !ok {
				return false
			}

			proto := &UpdateFightPetData{}
			if !proto.Read(b) {
				fmt.Println("read UpdateFightPetData fail, system error.")
				return false
			}

			t.OnUpdateFightPetData(proto)
		}
	case 46:
		{
			t, ok := protos.protoDispatch.(IPlayerKillReq)
			if !ok {
				return false
			}

			proto := &PlayerKillReq{}
			if !proto.Read(b) {
				fmt.Println("read PlayerKillReq fail, system error.")
				return false
			}

			t.OnPlayerKillReq(proto)
		}
	case 47:
		{
			t, ok := protos.protoDispatch.(IPlayerKillAck)
			if !ok {
				return false
			}

			proto := &PlayerKillAck{}
			if !proto.Read(b) {
				fmt.Println("read PlayerKillAck fail, system error.")
				return false
			}

			t.OnPlayerKillAck(proto)
		}
	case 48:
		{
			t, ok := protos.protoDispatch.(IEnterFightViewReq)
			if !ok {
				return false
			}

			proto := &EnterFightViewReq{}
			if !proto.Read(b) {
				fmt.Println("read EnterFightViewReq fail, system error.")
				return false
			}

			t.OnEnterFightViewReq(proto)
		}
	case 49:
		{
			t, ok := protos.protoDispatch.(IEnterFightViewAck)
			if !ok {
				return false
			}

			proto := &EnterFightViewAck{}
			if !proto.Read(b) {
				fmt.Println("read EnterFightViewAck fail, system error.")
				return false
			}

			t.OnEnterFightViewAck(proto)
		}
	case 50:
		{
			t, ok := protos.protoDispatch.(ILeaveFightViewReq)
			if !ok {
				return false
			}

			proto := &LeaveFightViewReq{}
			if !proto.Read(b) {
				fmt.Println("read LeaveFightViewReq fail, system error.")
				return false
			}

			t.OnLeaveFightViewReq(proto)
		}
	case 51:
		{
			t, ok := protos.protoDispatch.(ILeaveFightViewAck)
			if !ok {
				return false
			}

			proto := &LeaveFightViewAck{}
			if !proto.Read(b) {
				fmt.Println("read LeaveFightViewAck fail, system error.")
				return false
			}

			t.OnLeaveFightViewAck(proto)
		}
	case 52:
		{
			t, ok := protos.protoDispatch.(ITeamLeaderOprReq)
			if !ok {
				return false
			}

			proto := &TeamLeaderOprReq{}
			if !proto.Read(b) {
				fmt.Println("read TeamLeaderOprReq fail, system error.")
				return false
			}

			t.OnTeamLeaderOprReq(proto)
		}
	case 53:
		{
			t, ok := protos.protoDispatch.(ITeamNtf)
			if !ok {
				return false
			}

			proto := &TeamNtf{}
			if !proto.Read(b) {
				fmt.Println("read TeamNtf fail, system error.")
				return false
			}

			t.OnTeamNtf(proto)
		}
	case 54:
		{
			t, ok := protos.protoDispatch.(ITeamLeaderNtf)
			if !ok {
				return false
			}

			proto := &TeamLeaderNtf{}
			if !proto.Read(b) {
				fmt.Println("read TeamLeaderNtf fail, system error.")
				return false
			}

			t.OnTeamLeaderNtf(proto)
		}
	case 55:
		{
			t, ok := protos.protoDispatch.(ITeamDestroyNtf)
			if !ok {
				return false
			}

			proto := &TeamDestroyNtf{}
			if !proto.Read(b) {
				fmt.Println("read TeamDestroyNtf fail, system error.")
				return false
			}

			t.OnTeamDestroyNtf(proto)
		}
	case 56:
		{
			t, ok := protos.protoDispatch.(ITeamMemberNtf)
			if !ok {
				return false
			}

			proto := &TeamMemberNtf{}
			if !proto.Read(b) {
				fmt.Println("read TeamMemberNtf fail, system error.")
				return false
			}

			t.OnTeamMemberNtf(proto)
		}
	case 57:
		{
			t, ok := protos.protoDispatch.(ITeamMemberLeaveNtf)
			if !ok {
				return false
			}

			proto := &TeamMemberLeaveNtf{}
			if !proto.Read(b) {
				fmt.Println("read TeamMemberLeaveNtf fail, system error.")
				return false
			}

			t.OnTeamMemberLeaveNtf(proto)
		}
	case 58:
		{
			t, ok := protos.protoDispatch.(INpcSelectReq)
			if !ok {
				return false
			}

			proto := &NpcSelectReq{}
			if !proto.Read(b) {
				fmt.Println("read NpcSelectReq fail, system error.")
				return false
			}

			t.OnNpcSelectReq(proto)
		}
	case 59:
		{
			t, ok := protos.protoDispatch.(INpcSelectAck)
			if !ok {
				return false
			}

			proto := &NpcSelectAck{}
			if !proto.Read(b) {
				fmt.Println("read NpcSelectAck fail, system error.")
				return false
			}

			t.OnNpcSelectAck(proto)
		}
	case 60:
		{
			t, ok := protos.protoDispatch.(INpcTalkReq)
			if !ok {
				return false
			}

			proto := &NpcTalkReq{}
			if !proto.Read(b) {
				fmt.Println("read NpcTalkReq fail, system error.")
				return false
			}

			t.OnNpcTalkReq(proto)
		}
	case 61:
		{
			t, ok := protos.protoDispatch.(INpcTalkAck)
			if !ok {
				return false
			}

			proto := &NpcTalkAck{}
			if !proto.Read(b) {
				fmt.Println("read NpcTalkAck fail, system error.")
				return false
			}

			t.OnNpcTalkAck(proto)
		}
	case 62:
		{
			t, ok := protos.protoDispatch.(IInviteMsgNtf)
			if !ok {
				return false
			}

			proto := &InviteMsgNtf{}
			if !proto.Read(b) {
				fmt.Println("read InviteMsgNtf fail, system error.")
				return false
			}

			t.OnInviteMsgNtf(proto)
		}
	case 63:
		{
			t, ok := protos.protoDispatch.(IReplyInvite)
			if !ok {
				return false
			}

			proto := &ReplyInvite{}
			if !proto.Read(b) {
				fmt.Println("read ReplyInvite fail, system error.")
				return false
			}

			t.OnReplyInvite(proto)
		}
	case 64:
		{
			t, ok := protos.protoDispatch.(IMoveItem)
			if !ok {
				return false
			}

			proto := &MoveItem{}
			if !proto.Read(b) {
				fmt.Println("read MoveItem fail, system error.")
				return false
			}

			t.OnMoveItem(proto)
		}
	case 65:
		{
			t, ok := protos.protoDispatch.(IUseItem)
			if !ok {
				return false
			}

			proto := &UseItem{}
			if !proto.Read(b) {
				fmt.Println("read UseItem fail, system error.")
				return false
			}

			t.OnUseItem(proto)
		}
	case 66:
		{
			t, ok := protos.protoDispatch.(IRearrangeItem)
			if !ok {
				return false
			}

			proto := &RearrangeItem{}
			if !proto.Read(b) {
				fmt.Println("read RearrangeItem fail, system error.")
				return false
			}

			t.OnRearrangeItem(proto)
		}
	case 67:
		{
			t, ok := protos.protoDispatch.(ISkillContainerNtf)
			if !ok {
				return false
			}

			proto := &SkillContainerNtf{}
			if !proto.Read(b) {
				fmt.Println("read SkillContainerNtf fail, system error.")
				return false
			}

			t.OnSkillContainerNtf(proto)
		}
	case 68:
		{
			t, ok := protos.protoDispatch.(IAddSkillNtf)
			if !ok {
				return false
			}

			proto := &AddSkillNtf{}
			if !proto.Read(b) {
				fmt.Println("read AddSkillNtf fail, system error.")
				return false
			}

			t.OnAddSkillNtf(proto)
		}
	case 69:
		{
			t, ok := protos.protoDispatch.(IUpdateSkillNtf)
			if !ok {
				return false
			}

			proto := &UpdateSkillNtf{}
			if !proto.Read(b) {
				fmt.Println("read UpdateSkillNtf fail, system error.")
				return false
			}

			t.OnUpdateSkillNtf(proto)
		}
	case 70:
		{
			t, ok := protos.protoDispatch.(IDelSkillNtf)
			if !ok {
				return false
			}

			proto := &DelSkillNtf{}
			if !proto.Read(b) {
				fmt.Println("read DelSkillNtf fail, system error.")
				return false
			}

			t.OnDelSkillNtf(proto)
		}
	case 71:
		{
			t, ok := protos.protoDispatch.(IPetAppearNtf)
			if !ok {
				return false
			}

			proto := &PetAppearNtf{}
			if !proto.Read(b) {
				fmt.Println("read PetAppearNtf fail, system error.")
				return false
			}

			t.OnPetAppearNtf(proto)
		}
	case 72:
		{
			t, ok := protos.protoDispatch.(IPetContainerNtf)
			if !ok {
				return false
			}

			proto := &PetContainerNtf{}
			if !proto.Read(b) {
				fmt.Println("read PetContainerNtf fail, system error.")
				return false
			}

			t.OnPetContainerNtf(proto)
		}
	case 73:
		{
			t, ok := protos.protoDispatch.(IPetContainerUpdateNtf)
			if !ok {
				return false
			}

			proto := &PetContainerUpdateNtf{}
			if !proto.Read(b) {
				fmt.Println("read PetContainerUpdateNtf fail, system error.")
				return false
			}

			t.OnPetContainerUpdateNtf(proto)
		}
	case 74:
		{
			t, ok := protos.protoDispatch.(IPetAddNtf)
			if !ok {
				return false
			}

			proto := &PetAddNtf{}
			if !proto.Read(b) {
				fmt.Println("read PetAddNtf fail, system error.")
				return false
			}

			t.OnPetAddNtf(proto)
		}
	case 75:
		{
			t, ok := protos.protoDispatch.(IPetDestroyNtf)
			if !ok {
				return false
			}

			proto := &PetDestroyNtf{}
			if !proto.Read(b) {
				fmt.Println("read PetDestroyNtf fail, system error.")
				return false
			}

			t.OnPetDestroyNtf(proto)
		}
	case 76:
		{
			t, ok := protos.protoDispatch.(ISetPetLineup)
			if !ok {
				return false
			}

			proto := &SetPetLineup{}
			if !proto.Read(b) {
				fmt.Println("read SetPetLineup fail, system error.")
				return false
			}

			t.OnSetPetLineup(proto)
		}
	case 77:
		{
			t, ok := protos.protoDispatch.(IShowPet)
			if !ok {
				return false
			}

			proto := &ShowPet{}
			if !proto.Read(b) {
				fmt.Println("read ShowPet fail, system error.")
				return false
			}

			t.OnShowPet(proto)
		}
	case 78:
		{
			t, ok := protos.protoDispatch.(IReleasePet)
			if !ok {
				return false
			}

			proto := &ReleasePet{}
			if !proto.Read(b) {
				fmt.Println("read ReleasePet fail, system error.")
				return false
			}

			t.OnReleasePet(proto)
		}
	case 79:
		{
			t, ok := protos.protoDispatch.(IMovePet)
			if !ok {
				return false
			}

			proto := &MovePet{}
			if !proto.Read(b) {
				fmt.Println("read MovePet fail, system error.")
				return false
			}

			t.OnMovePet(proto)
		}
	case 80:
		{
			t, ok := protos.protoDispatch.(IShopOpenNtf)
			if !ok {
				return false
			}

			proto := &ShopOpenNtf{}
			if !proto.Read(b) {
				fmt.Println("read ShopOpenNtf fail, system error.")
				return false
			}

			t.OnShopOpenNtf(proto)
		}
	case 81:
		{
			t, ok := protos.protoDispatch.(IShopBuyNtf)
			if !ok {
				return false
			}

			proto := &ShopBuyNtf{}
			if !proto.Read(b) {
				fmt.Println("read ShopBuyNtf fail, system error.")
				return false
			}

			t.OnShopBuyNtf(proto)
		}
	case 82:
		{
			t, ok := protos.protoDispatch.(ISellNtf)
			if !ok {
				return false
			}

			proto := &SellNtf{}
			if !proto.Read(b) {
				fmt.Println("read SellNtf fail, system error.")
				return false
			}

			t.OnSellNtf(proto)
		}
	case 83:
		{
			t, ok := protos.protoDispatch.(IBuyBackNtf)
			if !ok {
				return false
			}

			proto := &BuyBackNtf{}
			if !proto.Read(b) {
				fmt.Println("read BuyBackNtf fail, system error.")
				return false
			}

			t.OnBuyBackNtf(proto)
		}
	case 84:
		{
			t, ok := protos.protoDispatch.(IBuyBackListNtf)
			if !ok {
				return false
			}

			proto := &BuyBackListNtf{}
			if !proto.Read(b) {
				fmt.Println("read BuyBackListNtf fail, system error.")
				return false
			}

			t.OnBuyBackListNtf(proto)
		}
	case 85:
		{
			t, ok := protos.protoDispatch.(ITeamAttrNtf)
			if !ok {
				return false
			}

			proto := &TeamAttrNtf{}
			if !proto.Read(b) {
				fmt.Println("read TeamAttrNtf fail, system error.")
				return false
			}

			t.OnTeamAttrNtf(proto)
		}
	case 86:
		{
			t, ok := protos.protoDispatch.(ITipsMsgExNtf)
			if !ok {
				return false
			}

			proto := &TipsMsgExNtf{}
			if !proto.Read(b) {
				fmt.Println("read TipsMsgExNtf fail, system error.")
				return false
			}

			t.OnTipsMsgExNtf(proto)
		}
	case 87:
		{
			t, ok := protos.protoDispatch.(IItemNewAddNtf)
			if !ok {
				return false
			}

			proto := &ItemNewAddNtf{}
			if !proto.Read(b) {
				fmt.Println("read ItemNewAddNtf fail, system error.")
				return false
			}

			t.OnItemNewAddNtf(proto)
		}
	case 88:
		{
			t, ok := protos.protoDispatch.(IQuestContainerNtf)
			if !ok {
				return false
			}

			proto := &QuestContainerNtf{}
			if !proto.Read(b) {
				fmt.Println("read QuestContainerNtf fail, system error.")
				return false
			}

			t.OnQuestContainerNtf(proto)
		}
	case 89:
		{
			t, ok := protos.protoDispatch.(IQuestAddNtf)
			if !ok {
				return false
			}

			proto := &QuestAddNtf{}
			if !proto.Read(b) {
				fmt.Println("read QuestAddNtf fail, system error.")
				return false
			}

			t.OnQuestAddNtf(proto)
		}
	case 90:
		{
			t, ok := protos.protoDispatch.(IQuestStateNtf)
			if !ok {
				return false
			}

			proto := &QuestStateNtf{}
			if !proto.Read(b) {
				fmt.Println("read QuestStateNtf fail, system error.")
				return false
			}

			t.OnQuestStateNtf(proto)
		}
	case 91:
		{
			t, ok := protos.protoDispatch.(IQuestNpcStateNtf)
			if !ok {
				return false
			}

			proto := &QuestNpcStateNtf{}
			if !proto.Read(b) {
				fmt.Println("read QuestNpcStateNtf fail, system error.")
				return false
			}

			t.OnQuestNpcStateNtf(proto)
		}
	case 92:
		{
			t, ok := protos.protoDispatch.(IQuestTrackCountNtf)
			if !ok {
				return false
			}

			proto := &QuestTrackCountNtf{}
			if !proto.Read(b) {
				fmt.Println("read QuestTrackCountNtf fail, system error.")
				return false
			}

			t.OnQuestTrackCountNtf(proto)
		}
	case 93:
		{
			t, ok := protos.protoDispatch.(IQuestTalkSelectReq)
			if !ok {
				return false
			}

			proto := &QuestTalkSelectReq{}
			if !proto.Read(b) {
				fmt.Println("read QuestTalkSelectReq fail, system error.")
				return false
			}

			t.OnQuestTalkSelectReq(proto)
		}
	case 94:
		{
			t, ok := protos.protoDispatch.(IQuestTalkSelectAck)
			if !ok {
				return false
			}

			proto := &QuestTalkSelectAck{}
			if !proto.Read(b) {
				fmt.Println("read QuestTalkSelectAck fail, system error.")
				return false
			}

			t.OnQuestTalkSelectAck(proto)
		}
	case 95:
		{
			t, ok := protos.protoDispatch.(IObjStrCustomNtf)
			if !ok {
				return false
			}

			proto := &ObjStrCustomNtf{}
			if !proto.Read(b) {
				fmt.Println("read ObjStrCustomNtf fail, system error.")
				return false
			}

			t.OnObjStrCustomNtf(proto)
		}
	case 96:
		{
			t, ok := protos.protoDispatch.(IObjIntCustomNtf)
			if !ok {
				return false
			}

			proto := &ObjIntCustomNtf{}
			if !proto.Read(b) {
				fmt.Println("read ObjIntCustomNtf fail, system error.")
				return false
			}

			t.OnObjIntCustomNtf(proto)
		}
	case 97:
		{
			t, ok := protos.protoDispatch.(IObjDynAttrNtf)
			if !ok {
				return false
			}

			proto := &ObjDynAttrNtf{}
			if !proto.Read(b) {
				fmt.Println("read ObjDynAttrNtf fail, system error.")
				return false
			}

			t.OnObjDynAttrNtf(proto)
		}
	case 98:
		{
			t, ok := protos.protoDispatch.(ICommitQuestItemRsp)
			if !ok {
				return false
			}

			proto := &CommitQuestItemRsp{}
			if !proto.Read(b) {
				fmt.Println("read CommitQuestItemRsp fail, system error.")
				return false
			}

			t.OnCommitQuestItemRsp(proto)
		}
	case 99:
		{
			t, ok := protos.protoDispatch.(ICommitQuestItemNtf)
			if !ok {
				return false
			}

			proto := &CommitQuestItemNtf{}
			if !proto.Read(b) {
				fmt.Println("read CommitQuestItemNtf fail, system error.")
				return false
			}

			t.OnCommitQuestItemNtf(proto)
		}
	case 100:
		{
			t, ok := protos.protoDispatch.(ICommitQuestPetRsp)
			if !ok {
				return false
			}

			proto := &CommitQuestPetRsp{}
			if !proto.Read(b) {
				fmt.Println("read CommitQuestPetRsp fail, system error.")
				return false
			}

			t.OnCommitQuestPetRsp(proto)
		}
	case 101:
		{
			t, ok := protos.protoDispatch.(ICommitQuestPetNtf)
			if !ok {
				return false
			}

			proto := &CommitQuestPetNtf{}
			if !proto.Read(b) {
				fmt.Println("read CommitQuestPetNtf fail, system error.")
				return false
			}

			t.OnCommitQuestPetNtf(proto)
		}
	case 102:
		{
			t, ok := protos.protoDispatch.(IAbandonQuest)
			if !ok {
				return false
			}

			proto := &AbandonQuest{}
			if !proto.Read(b) {
				fmt.Println("read AbandonQuest fail, system error.")
				return false
			}

			t.OnAbandonQuest(proto)
		}
	case 103:
		{
			t, ok := protos.protoDispatch.(ISkillTipsReq)
			if !ok {
				return false
			}

			proto := &SkillTipsReq{}
			if !proto.Read(b) {
				fmt.Println("read SkillTipsReq fail, system error.")
				return false
			}

			t.OnSkillTipsReq(proto)
		}
	case 104:
		{
			t, ok := protos.protoDispatch.(ISkillTipsAck)
			if !ok {
				return false
			}

			proto := &SkillTipsAck{}
			if !proto.Read(b) {
				fmt.Println("read SkillTipsAck fail, system error.")
				return false
			}

			t.OnSkillTipsAck(proto)
		}
	case 105:
		{
			t, ok := protos.protoDispatch.(ITeamLeaderOprNtf)
			if !ok {
				return false
			}

			proto := &TeamLeaderOprNtf{}
			if !proto.Read(b) {
				fmt.Println("read TeamLeaderOprNtf fail, system error.")
				return false
			}

			t.OnTeamLeaderOprNtf(proto)
		}
	case 106:
		{
			t, ok := protos.protoDispatch.(IKeepAliveCtrlNtf)
			if !ok {
				return false
			}

			proto := &KeepAliveCtrlNtf{}
			if !proto.Read(b) {
				fmt.Println("read KeepAliveCtrlNtf fail, system error.")
				return false
			}

			t.OnKeepAliveCtrlNtf(proto)
		}
	case 107:
		{
			t, ok := protos.protoDispatch.(INetDelayReq)
			if !ok {
				return false
			}

			proto := &NetDelayReq{}
			if !proto.Read(b) {
				fmt.Println("read NetDelayReq fail, system error.")
				return false
			}

			t.OnNetDelayReq(proto)
		}
	case 108:
		{
			t, ok := protos.protoDispatch.(INetDelayAck)
			if !ok {
				return false
			}

			proto := &NetDelayAck{}
			if !proto.Read(b) {
				fmt.Println("read NetDelayAck fail, system error.")
				return false
			}

			t.OnNetDelayAck(proto)
		}
	case 109:
		{
			t, ok := protos.protoDispatch.(ISystemSetupNtf)
			if !ok {
				return false
			}

			proto := &SystemSetupNtf{}
			if !proto.Read(b) {
				fmt.Println("read SystemSetupNtf fail, system error.")
				return false
			}

			t.OnSystemSetupNtf(proto)
		}
	case 110:
		{
			t, ok := protos.protoDispatch.(IBuffListNtf)
			if !ok {
				return false
			}

			proto := &BuffListNtf{}
			if !proto.Read(b) {
				fmt.Println("read BuffListNtf fail, system error.")
				return false
			}

			t.OnBuffListNtf(proto)
		}
	case 111:
		{
			t, ok := protos.protoDispatch.(IBuffAddNtf)
			if !ok {
				return false
			}

			proto := &BuffAddNtf{}
			if !proto.Read(b) {
				fmt.Println("read BuffAddNtf fail, system error.")
				return false
			}

			t.OnBuffAddNtf(proto)
		}
	case 112:
		{
			t, ok := protos.protoDispatch.(IBuffDurationNtf)
			if !ok {
				return false
			}

			proto := &BuffDurationNtf{}
			if !proto.Read(b) {
				fmt.Println("read BuffDurationNtf fail, system error.")
				return false
			}

			t.OnBuffDurationNtf(proto)
		}
	case 113:
		{
			t, ok := protos.protoDispatch.(IBuffDynAttrNtf)
			if !ok {
				return false
			}

			proto := &BuffDynAttrNtf{}
			if !proto.Read(b) {
				fmt.Println("read BuffDynAttrNtf fail, system error.")
				return false
			}

			t.OnBuffDynAttrNtf(proto)
		}
	case 114:
		{
			t, ok := protos.protoDispatch.(IStopBuff)
			if !ok {
				return false
			}

			proto := &StopBuff{}
			if !proto.Read(b) {
				fmt.Println("read StopBuff fail, system error.")
				return false
			}

			t.OnStopBuff(proto)
		}
	case 115:
		{
			t, ok := protos.protoDispatch.(IBuffDelNtf)
			if !ok {
				return false
			}

			proto := &BuffDelNtf{}
			if !proto.Read(b) {
				fmt.Println("read BuffDelNtf fail, system error.")
				return false
			}

			t.OnBuffDelNtf(proto)
		}
	case 116:
		{
			t, ok := protos.protoDispatch.(IActivityListReq)
			if !ok {
				return false
			}

			proto := &ActivityListReq{}
			if !proto.Read(b) {
				fmt.Println("read ActivityListReq fail, system error.")
				return false
			}

			t.OnActivityListReq(proto)
		}
	case 117:
		{
			t, ok := protos.protoDispatch.(IActivityListAck)
			if !ok {
				return false
			}

			proto := &ActivityListAck{}
			if !proto.Read(b) {
				fmt.Println("read ActivityListAck fail, system error.")
				return false
			}

			t.OnActivityListAck(proto)
		}
	case 118:
		{
			t, ok := protos.protoDispatch.(IActivityJoinNtf)
			if !ok {
				return false
			}

			proto := &ActivityJoinNtf{}
			if !proto.Read(b) {
				fmt.Println("read ActivityJoinNtf fail, system error.")
				return false
			}

			t.OnActivityJoinNtf(proto)
		}
	case 119:
		{
			t, ok := protos.protoDispatch.(IGuildListNtf)
			if !ok {
				return false
			}

			proto := &GuildListNtf{}
			if !proto.Read(b) {
				fmt.Println("read GuildListNtf fail, system error.")
				return false
			}

			t.OnGuildListNtf(proto)
		}
	case 120:
		{
			t, ok := protos.protoDispatch.(IGuildNtf)
			if !ok {
				return false
			}

			proto := &GuildNtf{}
			if !proto.Read(b) {
				fmt.Println("read GuildNtf fail, system error.")
				return false
			}

			t.OnGuildNtf(proto)
		}
	case 121:
		{
			t, ok := protos.protoDispatch.(IGuildBaseNtf)
			if !ok {
				return false
			}

			proto := &GuildBaseNtf{}
			if !proto.Read(b) {
				fmt.Println("read GuildBaseNtf fail, system error.")
				return false
			}

			t.OnGuildBaseNtf(proto)
		}
	case 122:
		{
			t, ok := protos.protoDispatch.(IGuildNoticeNtf)
			if !ok {
				return false
			}

			proto := &GuildNoticeNtf{}
			if !proto.Read(b) {
				fmt.Println("read GuildNoticeNtf fail, system error.")
				return false
			}

			t.OnGuildNoticeNtf(proto)
		}
	case 123:
		{
			t, ok := protos.protoDispatch.(IGuildBriefNtf)
			if !ok {
				return false
			}

			proto := &GuildBriefNtf{}
			if !proto.Read(b) {
				fmt.Println("read GuildBriefNtf fail, system error.")
				return false
			}

			t.OnGuildBriefNtf(proto)
		}
	case 124:
		{
			t, ok := protos.protoDispatch.(IGuildMemberListNtf)
			if !ok {
				return false
			}

			proto := &GuildMemberListNtf{}
			if !proto.Read(b) {
				fmt.Println("read GuildMemberListNtf fail, system error.")
				return false
			}

			t.OnGuildMemberListNtf(proto)
		}
	case 125:
		{
			t, ok := protos.protoDispatch.(IGuildMemberNtf)
			if !ok {
				return false
			}

			proto := &GuildMemberNtf{}
			if !proto.Read(b) {
				fmt.Println("read GuildMemberNtf fail, system error.")
				return false
			}

			t.OnGuildMemberNtf(proto)
		}
	case 126:
		{
			t, ok := protos.protoDispatch.(IGuildApplicantListNtf)
			if !ok {
				return false
			}

			proto := &GuildApplicantListNtf{}
			if !proto.Read(b) {
				fmt.Println("read GuildApplicantListNtf fail, system error.")
				return false
			}

			t.OnGuildApplicantListNtf(proto)
		}
	case 127:
		{
			t, ok := protos.protoDispatch.(IGuildOperationNtf)
			if !ok {
				return false
			}

			proto := &GuildOperationNtf{}
			if !proto.Read(b) {
				fmt.Println("read GuildOperationNtf fail, system error.")
				return false
			}

			t.OnGuildOperationNtf(proto)
		}
	case 128:
		{
			t, ok := protos.protoDispatch.(IGetTimestampReq)
			if !ok {
				return false
			}

			proto := &GetTimestampReq{}
			if !proto.Read(b) {
				fmt.Println("read GetTimestampReq fail, system error.")
				return false
			}

			t.OnGetTimestampReq(proto)
		}
	case 129:
		{
			t, ok := protos.protoDispatch.(IGetTimestampAck)
			if !ok {
				return false
			}

			proto := &GetTimestampAck{}
			if !proto.Read(b) {
				fmt.Println("read GetTimestampAck fail, system error.")
				return false
			}

			t.OnGetTimestampAck(proto)
		}
	case 130:
		{
			t, ok := protos.protoDispatch.(IContactListNtf)
			if !ok {
				return false
			}

			proto := &ContactListNtf{}
			if !proto.Read(b) {
				fmt.Println("read ContactListNtf fail, system error.")
				return false
			}

			t.OnContactListNtf(proto)
		}
	case 131:
		{
			t, ok := protos.protoDispatch.(IAddContactNtf)
			if !ok {
				return false
			}

			proto := &AddContactNtf{}
			if !proto.Read(b) {
				fmt.Println("read AddContactNtf fail, system error.")
				return false
			}

			t.OnAddContactNtf(proto)
		}
	case 132:
		{
			t, ok := protos.protoDispatch.(IUpdateContactNtf)
			if !ok {
				return false
			}

			proto := &UpdateContactNtf{}
			if !proto.Read(b) {
				fmt.Println("read UpdateContactNtf fail, system error.")
				return false
			}

			t.OnUpdateContactNtf(proto)
		}
	case 133:
		{
			t, ok := protos.protoDispatch.(IDelContactNtf)
			if !ok {
				return false
			}

			proto := &DelContactNtf{}
			if !proto.Read(b) {
				fmt.Println("read DelContactNtf fail, system error.")
				return false
			}

			t.OnDelContactNtf(proto)
		}
	case 134:
		{
			t, ok := protos.protoDispatch.(IAddContactMessageNtf)
			if !ok {
				return false
			}

			proto := &AddContactMessageNtf{}
			if !proto.Read(b) {
				fmt.Println("read AddContactMessageNtf fail, system error.")
				return false
			}

			t.OnAddContactMessageNtf(proto)
		}
	case 135:
		{
			t, ok := protos.protoDispatch.(IItemQueryNtf)
			if !ok {
				return false
			}

			proto := &ItemQueryNtf{}
			if !proto.Read(b) {
				fmt.Println("read ItemQueryNtf fail, system error.")
				return false
			}

			t.OnItemQueryNtf(proto)
		}
	case 136:
		{
			t, ok := protos.protoDispatch.(IPetQueryNtf)
			if !ok {
				return false
			}

			proto := &PetQueryNtf{}
			if !proto.Read(b) {
				fmt.Println("read PetQueryNtf fail, system error.")
				return false
			}

			t.OnPetQueryNtf(proto)
		}
	case 137:
		{
			t, ok := protos.protoDispatch.(IContactInfoNtf)
			if !ok {
				return false
			}

			proto := &ContactInfoNtf{}
			if !proto.Read(b) {
				fmt.Println("read ContactInfoNtf fail, system error.")
				return false
			}

			t.OnContactInfoNtf(proto)
		}
	case 138:
		{
			t, ok := protos.protoDispatch.(IMailListNtf)
			if !ok {
				return false
			}

			proto := &MailListNtf{}
			if !proto.Read(b) {
				fmt.Println("read MailListNtf fail, system error.")
				return false
			}

			t.OnMailListNtf(proto)
		}
	case 139:
		{
			t, ok := protos.protoDispatch.(IAddMailNtf)
			if !ok {
				return false
			}

			proto := &AddMailNtf{}
			if !proto.Read(b) {
				fmt.Println("read AddMailNtf fail, system error.")
				return false
			}

			t.OnAddMailNtf(proto)
		}
	case 140:
		{
			t, ok := protos.protoDispatch.(IDelMailNtf)
			if !ok {
				return false
			}

			proto := &DelMailNtf{}
			if !proto.Read(b) {
				fmt.Println("read DelMailNtf fail, system error.")
				return false
			}

			t.OnDelMailNtf(proto)
		}
	case 141:
		{
			t, ok := protos.protoDispatch.(IMailBodyNtf)
			if !ok {
				return false
			}

			proto := &MailBodyNtf{}
			if !proto.Read(b) {
				fmt.Println("read MailBodyNtf fail, system error.")
				return false
			}

			t.OnMailBodyNtf(proto)
		}
	case 142:
		{
			t, ok := protos.protoDispatch.(IUpdateMailBodyNtf)
			if !ok {
				return false
			}

			proto := &UpdateMailBodyNtf{}
			if !proto.Read(b) {
				fmt.Println("read UpdateMailBodyNtf fail, system error.")
				return false
			}

			t.OnUpdateMailBodyNtf(proto)
		}
	case 143:
		{
			t, ok := protos.protoDispatch.(IUpdateMailHeadNtf)
			if !ok {
				return false
			}

			proto := &UpdateMailHeadNtf{}
			if !proto.Read(b) {
				fmt.Println("read UpdateMailHeadNtf fail, system error.")
				return false
			}

			t.OnUpdateMailHeadNtf(proto)
		}
	case 144:
		{
			t, ok := protos.protoDispatch.(IRanklistReq)
			if !ok {
				return false
			}

			proto := &RanklistReq{}
			if !proto.Read(b) {
				fmt.Println("read RanklistReq fail, system error.")
				return false
			}

			t.OnRanklistReq(proto)
		}
	case 145:
		{
			t, ok := protos.protoDispatch.(IRanklistAck)
			if !ok {
				return false
			}

			proto := &RanklistAck{}
			if !proto.Read(b) {
				fmt.Println("read RanklistAck fail, system error.")
				return false
			}

			t.OnRanklistAck(proto)
		}
	case 146:
		{
			t, ok := protos.protoDispatch.(IGetRankReq)
			if !ok {
				return false
			}

			proto := &GetRankReq{}
			if !proto.Read(b) {
				fmt.Println("read GetRankReq fail, system error.")
				return false
			}

			t.OnGetRankReq(proto)
		}
	case 147:
		{
			t, ok := protos.protoDispatch.(IGetRankAck)
			if !ok {
				return false
			}

			proto := &GetRankAck{}
			if !proto.Read(b) {
				fmt.Println("read GetRankAck fail, system error.")
				return false
			}

			t.OnGetRankAck(proto)
		}
	case 148:
		{
			t, ok := protos.protoDispatch.(ITitleContainerNtf)
			if !ok {
				return false
			}

			proto := &TitleContainerNtf{}
			if !proto.Read(b) {
				fmt.Println("read TitleContainerNtf fail, system error.")
				return false
			}

			t.OnTitleContainerNtf(proto)
		}
	case 149:
		{
			t, ok := protos.protoDispatch.(ITitleAddNtf)
			if !ok {
				return false
			}

			proto := &TitleAddNtf{}
			if !proto.Read(b) {
				fmt.Println("read TitleAddNtf fail, system error.")
				return false
			}

			t.OnTitleAddNtf(proto)
		}
	case 150:
		{
			t, ok := protos.protoDispatch.(ITitleDelNtf)
			if !ok {
				return false
			}

			proto := &TitleDelNtf{}
			if !proto.Read(b) {
				fmt.Println("read TitleDelNtf fail, system error.")
				return false
			}

			t.OnTitleDelNtf(proto)
		}
	case 151:
		{
			t, ok := protos.protoDispatch.(IAgentKeyReq)
			if !ok {
				return false
			}

			proto := &AgentKeyReq{}
			if !proto.Read(b) {
				fmt.Println("read AgentKeyReq fail, system error.")
				return false
			}

			t.OnAgentKeyReq(proto)
		}
	case 152:
		{
			t, ok := protos.protoDispatch.(IAgentKeyAck)
			if !ok {
				return false
			}

			proto := &AgentKeyAck{}
			if !proto.Read(b) {
				fmt.Println("read AgentKeyAck fail, system error.")
				return false
			}

			t.OnAgentKeyAck(proto)
		}
	case 153:
		{
			t, ok := protos.protoDispatch.(IHeadMsgNtf)
			if !ok {
				return false
			}

			proto := &HeadMsgNtf{}
			if !proto.Read(b) {
				fmt.Println("read HeadMsgNtf fail, system error.")
				return false
			}

			t.OnHeadMsgNtf(proto)
		}
	case 154:
		{
			t, ok := protos.protoDispatch.(IAutoContainerNtf)
			if !ok {
				return false
			}

			proto := &AutoContainerNtf{}
			if !proto.Read(b) {
				fmt.Println("read AutoContainerNtf fail, system error.")
				return false
			}

			t.OnAutoContainerNtf(proto)
		}
	case 155:
		{
			t, ok := protos.protoDispatch.(IPlayerQueryNtf)
			if !ok {
				return false
			}

			proto := &PlayerQueryNtf{}
			if !proto.Read(b) {
				fmt.Println("read PlayerQueryNtf fail, system error.")
				return false
			}

			t.OnPlayerQueryNtf(proto)
		}
	case 156:
		{
			t, ok := protos.protoDispatch.(IUseAllItem)
			if !ok {
				return false
			}

			proto := &UseAllItem{}
			if !proto.Read(b) {
				fmt.Println("read UseAllItem fail, system error.")
				return false
			}

			t.OnUseAllItem(proto)
		}
	case 157:
		{
			t, ok := protos.protoDispatch.(IGuardContainerNtf)
			if !ok {
				return false
			}

			proto := &GuardContainerNtf{}
			if !proto.Read(b) {
				fmt.Println("read GuardContainerNtf fail, system error.")
				return false
			}

			t.OnGuardContainerNtf(proto)
		}
	case 158:
		{
			t, ok := protos.protoDispatch.(IGuardAddNtf)
			if !ok {
				return false
			}

			proto := &GuardAddNtf{}
			if !proto.Read(b) {
				fmt.Println("read GuardAddNtf fail, system error.")
				return false
			}

			t.OnGuardAddNtf(proto)
		}
	case 159:
		{
			t, ok := protos.protoDispatch.(ISetGuardLineup)
			if !ok {
				return false
			}

			proto := &SetGuardLineup{}
			if !proto.Read(b) {
				fmt.Println("read SetGuardLineup fail, system error.")
				return false
			}

			t.OnSetGuardLineup(proto)
		}
	case 160:
		{
			t, ok := protos.protoDispatch.(IPetNewAddNtf)
			if !ok {
				return false
			}

			proto := &PetNewAddNtf{}
			if !proto.Read(b) {
				fmt.Println("read PetNewAddNtf fail, system error.")
				return false
			}

			t.OnPetNewAddNtf(proto)
		}
	case 161:
		{
			t, ok := protos.protoDispatch.(ITeamPlatformNtf)
			if !ok {
				return false
			}

			proto := &TeamPlatformNtf{}
			if !proto.Read(b) {
				fmt.Println("read TeamPlatformNtf fail, system error.")
				return false
			}

			t.OnTeamPlatformNtf(proto)
		}
	case 162:
		{
			t, ok := protos.protoDispatch.(ITeamApplicantsNtf)
			if !ok {
				return false
			}

			proto := &TeamApplicantsNtf{}
			if !proto.Read(b) {
				fmt.Println("read TeamApplicantsNtf fail, system error.")
				return false
			}

			t.OnTeamApplicantsNtf(proto)
		}
	case 163:
		{
			t, ok := protos.protoDispatch.(ITeamOperationNtf)
			if !ok {
				return false
			}

			proto := &TeamOperationNtf{}
			if !proto.Read(b) {
				fmt.Println("read TeamOperationNtf fail, system error.")
				return false
			}

			t.OnTeamOperationNtf(proto)
		}
	case 164:
		{
			t, ok := protos.protoDispatch.(ITeamTargetNtf)
			if !ok {
				return false
			}

			proto := &TeamTargetNtf{}
			if !proto.Read(b) {
				fmt.Println("read TeamTargetNtf fail, system error.")
				return false
			}

			t.OnTeamTargetNtf(proto)
		}
	case 165:
		{
			t, ok := protos.protoDispatch.(IChangedNameNtf)
			if !ok {
				return false
			}

			proto := &ChangedNameNtf{}
			if !proto.Read(b) {
				fmt.Println("read ChangedNameNtf fail, system error.")
				return false
			}

			t.OnChangedNameNtf(proto)
		}
	case 166:
		{
			t, ok := protos.protoDispatch.(ICustomDataNtf)
			if !ok {
				return false
			}

			proto := &CustomDataNtf{}
			if !proto.Read(b) {
				fmt.Println("read CustomDataNtf fail, system error.")
				return false
			}

			t.OnCustomDataNtf(proto)
		}
	case 167:
		{
			t, ok := protos.protoDispatch.(ISpeedCheckNtf)
			if !ok {
				return false
			}

			proto := &SpeedCheckNtf{}
			if !proto.Read(b) {
				fmt.Println("read SpeedCheckNtf fail, system error.")
				return false
			}

			t.OnSpeedCheckNtf(proto)
		}
	case 168:
		{
			t, ok := protos.protoDispatch.(IConsoleMsgNtf)
			if !ok {
				return false
			}

			proto := &ConsoleMsgNtf{}
			if !proto.Read(b) {
				fmt.Println("read ConsoleMsgNtf fail, system error.")
				return false
			}

			t.OnConsoleMsgNtf(proto)
		}
	case 169:
		{
			t, ok := protos.protoDispatch.(IPetSwapNtf)
			if !ok {
				return false
			}

			proto := &PetSwapNtf{}
			if !proto.Read(b) {
				fmt.Println("read PetSwapNtf fail, system error.")
				return false
			}

			t.OnPetSwapNtf(proto)
		}
	case 170:
		{
			t, ok := protos.protoDispatch.(IGuardDestroyNtf)
			if !ok {
				return false
			}

			proto := &GuardDestroyNtf{}
			if !proto.Read(b) {
				fmt.Println("read GuardDestroyNtf fail, system error.")
				return false
			}

			t.OnGuardDestroyNtf(proto)
		}
	case 171:
		{
			t, ok := protos.protoDispatch.(IActivateGuard)
			if !ok {
				return false
			}

			proto := &ActivateGuard{}
			if !proto.Read(b) {
				fmt.Println("read ActivateGuard fail, system error.")
				return false
			}

			t.OnActivateGuard(proto)
		}
	case 172:
		{
			t, ok := protos.protoDispatch.(IReleaseGuard)
			if !ok {
				return false
			}

			proto := &ReleaseGuard{}
			if !proto.Read(b) {
				fmt.Println("read ReleaseGuard fail, system error.")
				return false
			}

			t.OnReleaseGuard(proto)
		}
	case 173:
		{
			t, ok := protos.protoDispatch.(ITeamMemberSwapNtf)
			if !ok {
				return false
			}

			proto := &TeamMemberSwapNtf{}
			if !proto.Read(b) {
				fmt.Println("read TeamMemberSwapNtf fail, system error.")
				return false
			}

			t.OnTeamMemberSwapNtf(proto)
		}
	case 174:
		{
			t, ok := protos.protoDispatch.(IGuardSwapNtf)
			if !ok {
				return false
			}

			proto := &GuardSwapNtf{}
			if !proto.Read(b) {
				fmt.Println("read GuardSwapNtf fail, system error.")
				return false
			}

			t.OnGuardSwapNtf(proto)
		}
	case 175:
		{
			t, ok := protos.protoDispatch.(IPetReplaceNtf)
			if !ok {
				return false
			}

			proto := &PetReplaceNtf{}
			if !proto.Read(b) {
				fmt.Println("read PetReplaceNtf fail, system error.")
				return false
			}

			t.OnPetReplaceNtf(proto)
		}
	case 176:
		{
			t, ok := protos.protoDispatch.(IGuardAppearNtf)
			if !ok {
				return false
			}

			proto := &GuardAppearNtf{}
			if !proto.Read(b) {
				fmt.Println("read GuardAppearNtf fail, system error.")
				return false
			}

			t.OnGuardAppearNtf(proto)
		}
	case 177:
		{
			t, ok := protos.protoDispatch.(IInstructionContainerNtf)
			if !ok {
				return false
			}

			proto := &InstructionContainerNtf{}
			if !proto.Read(b) {
				fmt.Println("read InstructionContainerNtf fail, system error.")
				return false
			}

			t.OnInstructionContainerNtf(proto)
		}
	case 178:
		{
			t, ok := protos.protoDispatch.(IInstructionAddReq)
			if !ok {
				return false
			}

			proto := &InstructionAddReq{}
			if !proto.Read(b) {
				fmt.Println("read InstructionAddReq fail, system error.")
				return false
			}

			t.OnInstructionAddReq(proto)
		}
	case 179:
		{
			t, ok := protos.protoDispatch.(IInstructionAddAck)
			if !ok {
				return false
			}

			proto := &InstructionAddAck{}
			if !proto.Read(b) {
				fmt.Println("read InstructionAddAck fail, system error.")
				return false
			}

			t.OnInstructionAddAck(proto)
		}
	case 180:
		{
			t, ok := protos.protoDispatch.(IInstructionDeleteReq)
			if !ok {
				return false
			}

			proto := &InstructionDeleteReq{}
			if !proto.Read(b) {
				fmt.Println("read InstructionDeleteReq fail, system error.")
				return false
			}

			t.OnInstructionDeleteReq(proto)
		}
	case 181:
		{
			t, ok := protos.protoDispatch.(IInstructionDeleteAck)
			if !ok {
				return false
			}

			proto := &InstructionDeleteAck{}
			if !proto.Read(b) {
				fmt.Println("read InstructionDeleteAck fail, system error.")
				return false
			}

			t.OnInstructionDeleteAck(proto)
		}
	case 182:
		{
			t, ok := protos.protoDispatch.(IInstructionModfityReq)
			if !ok {
				return false
			}

			proto := &InstructionModfityReq{}
			if !proto.Read(b) {
				fmt.Println("read InstructionModfityReq fail, system error.")
				return false
			}

			t.OnInstructionModfityReq(proto)
		}
	case 183:
		{
			t, ok := protos.protoDispatch.(IInstructionModifyAck)
			if !ok {
				return false
			}

			proto := &InstructionModifyAck{}
			if !proto.Read(b) {
				fmt.Println("read InstructionModifyAck fail, system error.")
				return false
			}

			t.OnInstructionModifyAck(proto)
		}
	case 184:
		{
			t, ok := protos.protoDispatch.(IInstructionDefaultReq)
			if !ok {
				return false
			}

			proto := &InstructionDefaultReq{}
			if !proto.Read(b) {
				fmt.Println("read InstructionDefaultReq fail, system error.")
				return false
			}

			t.OnInstructionDefaultReq(proto)
		}
	case 185:
		{
			t, ok := protos.protoDispatch.(IInstructionDefaultAck)
			if !ok {
				return false
			}

			proto := &InstructionDefaultAck{}
			if !proto.Read(b) {
				fmt.Println("read InstructionDefaultAck fail, system error.")
				return false
			}

			t.OnInstructionDefaultAck(proto)
		}
	case 186:
		{
			t, ok := protos.protoDispatch.(IInstructionAttachReq)
			if !ok {
				return false
			}

			proto := &InstructionAttachReq{}
			if !proto.Read(b) {
				fmt.Println("read InstructionAttachReq fail, system error.")
				return false
			}

			t.OnInstructionAttachReq(proto)
		}
	case 187:
		{
			t, ok := protos.protoDispatch.(IInstructionAttachAck)
			if !ok {
				return false
			}

			proto := &InstructionAttachAck{}
			if !proto.Read(b) {
				fmt.Println("read InstructionAttachAck fail, system error.")
				return false
			}

			t.OnInstructionAttachAck(proto)
		}
	case 188:
		{
			t, ok := protos.protoDispatch.(IInstructionAttachNtf)
			if !ok {
				return false
			}

			proto := &InstructionAttachNtf{}
			if !proto.Read(b) {
				fmt.Println("read InstructionAttachNtf fail, system error.")
				return false
			}

			t.OnInstructionAttachNtf(proto)
		}
	case 189:
		{
			t, ok := protos.protoDispatch.(IInstructionDetachReq)
			if !ok {
				return false
			}

			proto := &InstructionDetachReq{}
			if !proto.Read(b) {
				fmt.Println("read InstructionDetachReq fail, system error.")
				return false
			}

			t.OnInstructionDetachReq(proto)
		}
	case 190:
		{
			t, ok := protos.protoDispatch.(IInstructionDetachAck)
			if !ok {
				return false
			}

			proto := &InstructionDetachAck{}
			if !proto.Read(b) {
				fmt.Println("read InstructionDetachAck fail, system error.")
				return false
			}

			t.OnInstructionDetachAck(proto)
		}
	case 191:
		{
			t, ok := protos.protoDispatch.(IInstructionDetachNtf)
			if !ok {
				return false
			}

			proto := &InstructionDetachNtf{}
			if !proto.Read(b) {
				fmt.Println("read InstructionDetachNtf fail, system error.")
				return false
			}

			t.OnInstructionDetachNtf(proto)
		}
	case 192:
		{
			t, ok := protos.protoDispatch.(IPlayerDetailNtf)
			if !ok {
				return false
			}

			proto := &PlayerDetailNtf{}
			if !proto.Read(b) {
				fmt.Println("read PlayerDetailNtf fail, system error.")
				return false
			}

			t.OnPlayerDetailNtf(proto)
		}
	case 193:
		{
			t, ok := protos.protoDispatch.(IMapDynBlockPtNtf)
			if !ok {
				return false
			}

			proto := &MapDynBlockPtNtf{}
			if !proto.Read(b) {
				fmt.Println("read MapDynBlockPtNtf fail, system error.")
				return false
			}

			t.OnMapDynBlockPtNtf(proto)
		}
	case 194:
		{
			t, ok := protos.protoDispatch.(IGuardQueryNtf)
			if !ok {
				return false
			}

			proto := &GuardQueryNtf{}
			if !proto.Read(b) {
				fmt.Println("read GuardQueryNtf fail, system error.")
				return false
			}

			t.OnGuardQueryNtf(proto)
		}
	case 195:
		{
			t, ok := protos.protoDispatch.(IBuyBackNtfEx)
			if !ok {
				return false
			}

			proto := &BuyBackNtfEx{}
			if !proto.Read(b) {
				fmt.Println("read BuyBackNtfEx fail, system error.")
				return false
			}

			t.OnBuyBackNtfEx(proto)
		}
	case 196:
		{
			t, ok := protos.protoDispatch.(IGuildCustomNtf)
			if !ok {
				return false
			}

			proto := &GuildCustomNtf{}
			if !proto.Read(b) {
				fmt.Println("read GuildCustomNtf fail, system error.")
				return false
			}

			t.OnGuildCustomNtf(proto)
		}
	case 197:
		{
			t, ok := protos.protoDispatch.(IPreTurnRoundNtf)
			if !ok {
				return false
			}

			proto := &PreTurnRoundNtf{}
			if !proto.Read(b) {
				fmt.Println("read PreTurnRoundNtf fail, system error.")
				return false
			}

			t.OnPreTurnRoundNtf(proto)
		}
	case 198:
		{
			t, ok := protos.protoDispatch.(IFighterSpecialPetNtf)
			if !ok {
				return false
			}

			proto := &FighterSpecialPetNtf{}
			if !proto.Read(b) {
				fmt.Println("read FighterSpecialPetNtf fail, system error.")
				return false
			}

			t.OnFighterSpecialPetNtf(proto)
		}
	default:
		{
			fmt.Println("illegal protocol, Mid =", mid, "Pid =", pid)
		}
	}

	return true
}