///>本代码由自动化工具批量生成
package protocols

import "bytes"

type CustomStrData struct { //>自定义键值对
	Key   []uint8 //>key
	Value []uint8 //>value
}

type CustomIntData struct { //>自定义键值对
	Key   []uint8 //>key
	Value int64   //>value
}

type CustomData struct { //>自定义数据
	Strdata []CustomStrData //>字符串数据集
	Intdata []CustomIntData //>数值数据集
}

type SkillData struct { //>Skill数据
	Id              uint16 //>技能id
	Performance     uint32 //>当前熟练度
	Max_performance uint32 //>最大熟练度
	Tick            uint64 //>时间
	Enable          uint8  //>是否激活
	Bind            uint8  //>是否绑定
}

type DynAttrData struct { //>动态属性数据
	Mark  uint16 //>动态标记
	Attr  uint16 //>属性名
	Value int64  //>属性值
}

type BuffData struct { //>Buff数据
	Buff_id     uint16        //>Buff的ID
	Start_times uint32        //>开始时间(仅绝对时间buff有效)
	Duration    uint32        //>持续数据 相对时间buff:为剩余时间 绝时时间buff:为持续时间 次数buff:为剩余次数
	Dyn_attrs   []DynAttrData //>动态属性数据(固定属性值请从buff表读取)
}

type QuestData struct { //>Quest数据
	Quest_id       uint32  //>任务ID
	Quest_state    uint8   //>任务状态
	Accept_npc     uint32  //>接任务目标NPC
	Commit_npc     uint32  //>交任务目标NPC
	Quest_store_id uint32  //>任务库ID
	Accpet_time    uint32  //>任务接取时间
	Limit_time     uint32  //>任务限制时间
	Changed_time   uint32  //>任务状态改变时间
	Ext            []uint8 //>任务扩展数据
}

type QuestRingData struct { //>Quest环数据
	Quest_id   uint32 //>任务ID
	Quest_ring uint32 //>任务环数
}

type QuestStoreData struct { //>QuestStore数据
	Quest_store_id  uint32          //>任务库ID
	Ring_auto       uint32          //>自动接取环数
	Ring_day        uint32          //>当天已完成的环数
	Ring_week       uint32          //>本周已完成的环数
	Ring_display    uint32          //>显示环数
	Quest_rings     []QuestRingData //>子任务环数
	Last_reset_day  uint32          //>上次天数重置
	Last_reset_week uint32          //>上次周数重置
}

type QuestCountData struct { //>Quest计数数据
	Quest_id uint32 //>任务ID
	Count    uint32 //>任务计数
}

type QuestContainerData struct { //>任务容器数据
	Quests          []QuestData      //>任务
	Npc_visible     []uint32         //>可见npc列表
	Npc_invisible   []uint32         //>不可见npc列表
	Stores          []QuestStoreData //>任务库
	Abandoned       []uint32         //>放弃任务列表
	Finished        []uint32         //>完成任务列表
	Activated       []uint32         //>激活任务列表
	Failed_counts   []QuestCountData //>失败任务计数列表
	Finished_counts []QuestCountData //>完成任务计数列表
}

type AttrData struct { //>属性数据
	Attr  uint16 //>属性名
	Value int64  //>属性值
}

type ItemData struct { //>Item数据
	Guid       uint64        //>guid
	Id         uint16        //>模板Id
	Role       uint64        //>拥有者GUID
	Site       uint16        //>当前所在的位置
	Battlesite int16         //>战斗物品site
	Amount     uint16        //>当前的堆叠数量
	Isbound    uint8         //>是否已绑定
	Life       uint32        //>限时物品截止时间，0表示永久
	Dyn_attrs  []DynAttrData //>动态属性数据
	Customs    CustomData    //>自定义
}

type ItemContainerData struct { //>Item容器数据
	Container_type uint16     //>item容器类型
	Capacity       uint16     //>item容器容量上限
	Items          []ItemData //>道具
}

type UserData struct { //>账号数据
	Uid        uint32 //>账号ID
	Sid        uint8  //>子区ID
	Uname      string //>账号名字
	Billinyb   int64  //>累计充值元宝
	Billoutyb  int64  //>累计消耗的充值元宝
	Gameinyb   int64  //>累计游戏内产出元宝
	Gameoutyb  int64  //>累计游戏内消耗元宝
	Createtime uint32 //>创建时间
	Lastlogin  uint32 //>最后一次登陆时间
}

type SceneData struct { //>场景数据
	Map_guid uint64 //>地图guid
	X        uint16 //>地图x坐标
	Y        uint16 //>地图y坐标
}

type TicketData struct { //>门票数据
	Map_guid    uint64 //>地图guid
	Create_time uint32 //>创建时间
	Id          uint16 //>地图ID
	X           uint16 //>地图x坐标
	Y           uint16 //>地图y坐标
}

type PetData struct { //>Pet数据
	Guid        uint64            //>宠物GUID
	Name        string            //>宠物名字
	Equips      ItemContainerData //>宠物装备
	Attrs       []AttrData        //>属性数据
	Extra_attrs []AttrData        //>额外属性数据
	Skills      []SkillData       //>技能
	Buffs       []BuffData        //>buff
	Dyn_attrs   []DynAttrData     //>动态属性数据
	Customs     CustomData        //>自定义
}

type PetContainerData struct { //>Pet容器数据
	Container_type uint16    //>pet容器类型
	Capacity       uint16    //>pet容器容量上限
	Pets           []PetData //>宠物
}

type TitleData struct { //>称号数据
	Title_id uint16     //>称号ID
	End_time uint32     //>限时称号截止时间，0表示永久称号
	Custom   CustomData //>自定义数据
}

type TitleContainerData struct { //>称号容器数据
	Titles []TitleData //>称号
}

type SeatBasic struct { //>阵法基本数据
	Id    uint16 //>阵法id
	Level uint8  //>阵法等级
	Score uint32 //>积分值
}

type SeatData struct { //>阵法个人数据
	Seats []SeatBasic //>所有阵法数据
}

type LineupBasic struct { //>阵容基本数据
	Seat_id uint16   //>阵法id
	Lineup  []uint64 //>阵容顺序
}

type LineupData struct { //>阵容个人数据
	Index   uint8         //>当前阵容索引
	Lineups []LineupBasic //>所有阵容
}

type InstructionBasic struct { //>指令基础数据
	Type    uint8  //>指令类型
	Content string //>指令内容
}

type PetLineupData struct { //>宠物阵容基本数据
	Lineup []uint64 //>宠物阵容
}

type InstructionData struct { //>指令个人数据
	Frienddata []InstructionBasic //>友方指令
	Enemydata  []InstructionBasic //>敌方指令
}

type EmojiData struct { //>表情包数据
	Emojis []uint16 //>表情包记录
}

type PlayerBrief struct { //>玩家简略
	Uid         uint32     //>账号ID
	Sid         uint8      //>子区ID
	Guid        uint64     //>玩家GUID
	Unid        string     //>玩家UNID
	Name        string     //>玩家名字
	Attrs       []AttrData //>属性数据
	Map         uint16     //>地图ID
	X           uint16     //>x坐标
	Y           uint16     //>y坐标
	State       uint8      //>玩家状态 0:彻底销毁 1:使用中 2:已删除
	Createtime  uint32     //>创建时间
	Lastlogin   uint32     //>最后一次登陆时间
	Destroytime uint32     //>销毁时间 state为0=销毁时间 state为2=删除时间
	Customs     CustomData //>自定义
}

type GuardData struct { //>Guard数据
	Guid        uint64            //>侍从GUID
	Equips      ItemContainerData //>侍从装备
	Attrs       []AttrData        //>属性数据
	Extra_attrs []AttrData        //>额外属性数据
	Skills      []SkillData       //>技能
	Buffs       []BuffData        //>buff
	Dyn_attrs   []DynAttrData     //>动态属性数据
	Customs     CustomData        //>自定义
}

type GuardContainerData struct { //>Guard容器数据
	Guards []GuardData //>侍从
}

type PlayerData struct { //>玩家详细
	Uid             uint32              //>账号ID
	Sid             uint8               //>子区ID
	Guid            uint64              //>玩家GUID
	Unid            string              //>玩家UNID
	Name            string              //>玩家名字
	State           uint8               //>玩家状态 0:彻底销毁 1:使用中 2:已删除
	Createtime      uint32              //>创建时间
	Destroytime     uint32              //>销毁时间 state为0=销毁时间 state为2=删除时间
	Lastlogin       uint32              //>最后一次登陆时间
	Lastloginip     uint32              //>最后一次登陆ip
	Lastlogout      uint32              //>最后一次退出时间
	Onlines         uint32              //>总计在线时间
	Cur_map         SceneData           //>当前地图信息
	Last_common_map SceneData           //>最后固定地图信息
	Attrs           []AttrData          //>属性数据
	Dyn_attrs       []DynAttrData       //>动态属性数据
	Extra_attrs     []AttrData          //>额外属性数据
	Items           []ItemContainerData //>物品容器
	Skills          []SkillData         //>技能
	Buffs           []BuffData          //>buff
	Quests          QuestContainerData  //>任务
	Pets            []PetContainerData  //>宠物
	Guards          GuardContainerData  //>侍从
	Tickets         []TicketData        //>门票
	Customs         CustomData          //>自定义
	Gameinyb        int64               //>游戏内给予的元宝(不存档)
	Gameoutyb       int64               //>游戏内消耗的元宝(不存档)
	Billinyb        int64               //>充值给予的元宝(不存档)
	Billoutyb       int64               //>充值消耗的元宝(不存档)
	Titles          TitleContainerData  //>称号
	Username        string              //>账号名
	Lineups         LineupData          //>阵容数据
	Seats           SeatData            //>阵法数据
	Instructions    InstructionData     //>指令数据
	Pet_lineups     PetLineupData       //>宠物阵容
	Emojis          EmojiData           //>表情包阵容
}

type MapRegion struct { //>地图区域
	Min_x uint16 //>最小x坐标
	Max_x uint16 //>最大x坐标
	Min_y uint16 //>最小y坐标
	Max_y uint16 //>最大y坐标
}

type ItemAttrValue64 struct { //>物品属性
	Attr_type  uint32 //>物品属性类型
	Attr_value int64  //>物品属性值
}

type ItemAttrValueS struct { //>物品属性
	Attr_type  uint32  //>物品属性类型
	Attr_value []uint8 //>物品属性值
}

type ItemAttrValueS64 struct { //>物品属性
	Attr_key   string //>key
	Attr_value int64  //>物品属性值
}

type ItemAttrValueSS struct { //>物品属性
	Attr_key   string  //>key
	Attr_value []uint8 //>物品属性值
}

type ItemAttrValueList struct { //>物品属性值列表
	Item_guid       uint64             //>物品guid
	Item_values_64  []ItemAttrValue64  //>物品属性值列表:U64
	Item_values_s   []ItemAttrValueS   //>物品属性值列表:S
	Item_values_s64 []ItemAttrValueS64 //>物品属性值列表:SU64
	Item_values_ss  []ItemAttrValueSS  //>物品属性值列表:SS
	Dyn_attrs       []DynAttrData      //>物品属性值列表:动态属性
}

type ChatMsg struct { //>聊天消息
	Role_name string  //>角色名字
	Role_id   uint32  //>角色模板ID
	Message   []uint8 //>消息内容
	Vip       uint32  //>vip等级
}

type ParamData struct { //>命令参数
	Type  uint8   //>数据类型 1=INT8 2=UINT8 3=INT16 4=UINT16 5=INT32 6=UINT32 7=INT64 8=UINT64 9=字符串
	Param []uint8 //>命令参数
}

type FightPetData struct { //>战斗宠物数据
	Guid  uint64 //>宠物guid
	State uint8  //>0:不可召出 1:可召出
}

type FightBuffData struct { //>战斗BUFF数据
	Id    uint16 //>buff id
	Stack uint8  //>堆叠数
}

type FighterData struct { //>战斗者数据
	Site               uint8           //>战斗者站位
	Fighter_id         uint16          //>战斗者id
	Guid               uint64          //>guid
	Type               uint8           //>对象类型
	Template_id        uint16          //>模板id
	Name               string          //>名称
	Attrs              []AttrData      //>属性数据
	Pets               []FightPetData  //>宠物召出记录
	Buffs              []FightBuffData //>战斗buff数据
	Model              uint32          //>默认模型,玩家填0
	Friend_instruction string          //>显示给友方看的指令
	Enemy_instruction  string          //>显示给敌方看的指令
	Customs            CustomData      //>自定义
}

type FightGroupData struct { //>战斗组数据
	Type      uint8         //>0=观战组 1=组1 2=组2
	Lineup_id uint16        //>阵型id
	Lineup_lv uint16        //>阵型lv
	Fighters  []FighterData //>战斗者数据
}

type FightOperateData struct { //>战斗操作数据
	Actor   uint64 //>发起者guid 玩家自己:填0 宠物:填宠物guid
	Operate uint8  //>指令
	Target  uint64 //>目标guid
	Param   uint64 //>参数(使用技能时的技能ID，使用物品时物品的GUID，召唤宠物时宠物的GUID)
}

type MemberData struct { //>队伍成员数据
	Guid       uint64     //>玩家GUID
	Name       string     //>玩家名字
	Temp_leave uint8      //>暂离标记
	Attrs      []AttrData //>属性数据
	Customs    CustomData //>自定义
}

type KeyValueU32 struct { //>键值对
	Key   uint32 //>键
	Value uint32 //>值
}

type KeyValueU64 struct { //>键值对
	Key   uint32 //>键
	Value uint64 //>值
}

type KeyValueS struct { //>键值对
	Key   uint32 //>键
	Value string //>值
}

type OptionData struct { //>选项数据
	Id   uint16 //>选项id
	Mark uint32 //>选项标识
	Text string //>选项文字
}

type ShopItemData struct { //>商品数据
	Id         uint16 //>根据商店类型决定是道具或者宠物id
	Price_type uint16 //>价格类型 数值参考属性表对应字段
	Price      uint32 //>价格
	Bind       uint8  //>0为不绑定 1为绑定
}

type SellItemData struct { //>出售商品数据
	Guid   uint64 //>出售物品guid
	Amount uint16 //>出售物品数量
}

type Col1Data struct { //>列数据1
	Col1 string //>列数据
}

type Col2Data struct { //>列数据2
	Col1 string //>列数据
	Col2 string //>列数据
}

type Col3Data struct { //>列数据3
	Col1 string //>列数据
	Col2 string //>列数据
	Col3 string //>列数据
}

type Col4Data struct { //>列数据4
	Col1 string //>列数据
	Col2 string //>列数据
	Col3 string //>列数据
	Col4 string //>列数据
}

type Col5Data struct { //>列数据5
	Col1 string //>列数据
	Col2 string //>列数据
	Col3 string //>列数据
	Col4 string //>列数据
	Col5 string //>列数据
}

type Col6Data struct { //>列数据6
	Col1 string //>列数据
	Col2 string //>列数据
	Col3 string //>列数据
	Col4 string //>列数据
	Col5 string //>列数据
	Col6 string //>列数据
}

type Col7Data struct { //>列数据7
	Col1 string //>列数据
	Col2 string //>列数据
	Col3 string //>列数据
	Col4 string //>列数据
	Col5 string //>列数据
	Col6 string //>列数据
	Col7 string //>列数据
}

type Col8Data struct { //>列数据8
	Col1 string //>列数据
	Col2 string //>列数据
	Col3 string //>列数据
	Col4 string //>列数据
	Col5 string //>列数据
	Col6 string //>列数据
	Col7 string //>列数据
	Col8 string //>列数据
}

type Col9Data struct { //>列数据9
	Col1 string //>列数据
	Col2 string //>列数据
	Col3 string //>列数据
	Col4 string //>列数据
	Col5 string //>列数据
	Col6 string //>列数据
	Col7 string //>列数据
	Col8 string //>列数据
	Col9 string //>列数据
}

type Col10Data struct { //>列数据10
	Col1  string //>列数据
	Col2  string //>列数据
	Col3  string //>列数据
	Col4  string //>列数据
	Col5  string //>列数据
	Col6  string //>列数据
	Col7  string //>列数据
	Col8  string //>列数据
	Col9  string //>列数据
	Col10 string //>列数据
}

type TalkNodeData struct { //>对话节点数据
	Node_type    string //>列数据
	Node_id      string //>列数据
	Display_text string //>列数据
	Connect_text string //>列数据
	Npc_id       string //>列数据
	Npc_emotion  string //>列数据
	Npc_talk     string //>列数据
	Npc_template string //>列数据
	Func_type    string //>列数据
	Link_node    string //>列数据
	Belong_to    string //>列数据
	Player_talk  string //>列数据
}

type QuestClientData struct { //>Quest客户端数据
	Quest_id               uint32  //>任务ID
	Quest_type             uint32  //>任务类型
	Quest_type_name        string  //>任务类型名称
	Quest_name             string  //>任务名称
	Quest_color            uint32  //>任务颜色
	Quest_state            uint8   //>任务状态
	Display_location       uint8   //>显示位置，0:显示在可接列表 1:显示在已接列表
	Display_accept_tip     uint8   //>是否发送接受任务提示
	Display_accomplish_tip uint8   //>是否发送完成任务提示
	Can_track              uint8   //>任务追踪可见
	Can_abandon            uint8   //>任务是否可放弃
	Accpet_time            uint32  //>任务接受时间
	Limit_time             uint32  //>任务结束时间
	Changed_time           uint32  //>任务状态变更时间
	Accept_npc             uint64  //>接任务目标NPC
	Accept_npc_info        string  //>任务目标NPC信息
	Commit_npc             uint64  //>交任务目标NPC
	Commit_npc_info        string  //>交任务目标NPC信息
	Quest_goal_type        uint16  //>任务目标信息
	Quest_goal             []uint8 //>任务目标信息
	Quest_award            []uint8 //>任务奖励信息
	Track_desc             string  //>任务目标追踪信息
	Cant_accept_track_desc string  //>不可接任务目标追踪信息
	Cant_accept_desc       string  //>不可接描述信息
	Can_accept_desc        string  //>可接描述信息
	Accepted_desc          string  //>已接描述信息
	Is_display_cycle_num   uint8   //>是否显示轮环数
	Finish_cycle_num       uint32  //>完成环数
	Total_cycle_num        uint32  //>总环数
}

type QuestTalkData struct { //>任务对话数据
	Npc_id    uint32       //>NPC头像ID
	Node_id   uint32       //>对话节点ID
	Talk_text []uint8      //>对话内容
	Options   []OptionData //>对话选项
}

type IDCountData struct { //>ID数量
	Id     uint32 //>ID
	Amount uint16 //>数量
}

type GUIDCountData struct { //>GUID数量
	Guid   uint64 //>GUID
	Amount uint16 //>数量
}

type SkillTipsData struct { //>技能tips数据
	Name  uint8  //>tips名字
	Attr  uint16 //>属性名
	Value int64  //>属性值
}

type GoalVisitNpcData struct { //>拜访NPC目标数据
	Npc_id uint32 //>NPC ID
}

type GoalKillBossData struct { //>杀死BOSS目标数据
	Boss_id    uint32 //>BOSS ID
	Map_id     uint16 //>地图
	Loc_x      uint16 //>x坐标
	Loc_y      uint16 //>y坐标
	Kill_count uint16 //>杀死次数
}

type GoalCommitItemData struct { //>提交物品目标数据
	Item_id    uint32 //>物品ID
	Item_count uint32 //>物品数量
}

type GoalCommitPetData struct { //>提交宠物目标数据
	Pet_id    uint32 //>宠物ID
	Pet_count uint32 //>宠物数量
}

type GoalEscortNpcData struct { //>护送NPC目标数据
	Npc_id uint32 //>NPC ID
}

type GoalArriveAddrData struct { //>到达地址目标数据
	Map_id  uint16 //>地图ID
	Loc_x   uint16 //>x坐标
	Loc_y   uint16 //>y坐标
	Range   uint16 //>范围
	Item_id uint32 //>物品ID
	Npc_id  uint32 //>NPC ID
	Desc    string //>描述
}

type GoalFightMapData struct { //>地图战斗目标数据
	Map_id         uint16 //>地图ID
	Count          uint16 //>次数
	Count_required uint16 //>要求次数
}

type GoalFightMapCountData struct { //>地图战斗目标数据
	Fight_map []GoalFightMapData //>地图战斗
}

type GoalKillMonsterData struct { //>杀怪目标数据
	Monster_id     uint32 //>怪物ID
	Count          uint16 //>次数
	Count_required uint16 //>要求次数
	Map_id         uint16 //>地图ID
	Left           uint16 //>左
	Top            uint16 //>上
	Width          uint16 //>宽
	Height         uint16 //>高
}

type GoalKillMonsterCountData struct { //>杀怪目标数据
	Kill_monster []GoalKillMonsterData //>杀怪
}

type GoalEncounterFightData struct { //>遭遇战斗目标数据
	Cur_num uint8 //>次数
	Max_num uint8 //>最大次数
}

type GoalUseItemData struct { //>使用物品目标数据
	Map_id  uint16 //>地图ID
	Loc_x   uint16 //>x坐标
	Loc_y   uint16 //>y坐标
	Range   uint16 //>范围
	Item_id uint32 //>物品ID
	Npc_id  uint32 //>NPC ID
	Desc    string //>描述
}

type GoalCapturePetData struct { //>捕获目标数据
	Pet_id         uint16 //>宠物ID
	Count          uint16 //>次数
	Count_required uint16 //>要求次数
}

type GoalPVEData struct { //>PVE任务数据
	Win_cur      uint32 //>胜利次数
	Win_required uint32 //>要求胜利次数
}

type IPConfig struct { //>链接配置
	Type     uint8  //>服务器类型
	Port     uint16 //>服务器端口
	Recv_buf uint32 //>接收缓冲大小
	Send_buf uint32 //>发送缓冲大小
	Ip       string //>服务器IP
}

type ActivityData struct { //>活动数据
	Id     uint16  //>活动id
	State  uint8   //>活动状态 0:未开启 1:进行中 2:已过期
	Today  uint8   //>是否今天的活动
	Custom []uint8 //>脚本填充的自定义数据
}

type StringData struct { //>长字符串数据
	Str []uint8 //>字符串
}

type GUIDData struct { //>GUID数组
	Guid uint64 //>guids
}

type AuctionObjectData struct { //>拍卖品数据
	Guid          uint64   //>拍卖品guid
	Id            uint16   //>道具ID或宠物ID
	Type          uint16   //>1代表道具 2代表宠物
	Custom1       int32    //>脚本自定义分类1
	Custom2       int32    //>脚本自定义分类2
	Custom3       int32    //>脚本自定义分类3
	Level         uint16   //>道具:无 宠物:等级
	Fight_value   int32    //>道具:无 宠物:战力
	Coin_type     uint16   //>attr表里的货币属性名
	Coin_value    int64    //>货币值
	Amount        uint16   //>堆叠数量
	State         uint8    //>是否上架 0:公示中(预留) 1:上架中 2:已下架
	Puton_times   uint32   //>上架时间
	Pulloff_times uint32   //>下架时间
	Seller_guid   uint64   //>出售者guid
	Seller_name   string   //>出售者姓名
	Item_data     ItemData //>item数据(拍卖品为道具)
	Pet_data      PetData  //>pet数据(拍卖品为宠物)
}

type AuctionCookieData struct { //>拍卖行个人数据
	Guid      uint64       //>玩家guid
	Coins     []AttrData   //>各种货币
	Favorites []GUIDData   //>收藏的拍卖品
	Records   []StringData //>交易记录
}

type GuildMemberBriefData struct { //>帮派成员简略信息
	Player_guid      uint64     //>玩家guid
	Player_name      string     //>玩家姓名
	Level            uint32     //>等级
	Role             uint8      //>角色
	Guild_job        uint8      //>成员职位
	Total_contrb     uint32     //>历史贡献度
	Join_time        uint32     //>加入帮派时间
	Last_login_time  uint32     //>上次登陆时间
	Last_logout_time uint32     //>上次离线时间
	Status           uint32     //>成员玩家状态(参见logic_def的枚举guild_member_status)
	Vip              uint32     //>vip等级
	Reincarnation    uint32     //>转生等级
	Job              uint8      //>职业
	Customs          CustomData //>自定义
}

type GuildMemberData struct { //>帮派成员信息
	Player_guid        uint64     //>玩家guid
	Player_name        string     //>玩家姓名
	Level              uint32     //>等级
	Role               uint8      //>角色
	Guild_job          uint8      //>成员职位
	Total_contrb       uint32     //>历史贡献度
	Surplus_contrb     uint32     //>剩余贡献度
	Last_week_contrb   uint32     //>上周贡献度
	This_week_contrb   uint32     //>本周贡献度
	Permission         uint32     //>帮派权限(参见logic_def的枚举guild_permission)
	Join_time          uint32     //>加入帮派时间
	Donate_bind_gold   uint32     //>捐赠金币
	Last_donate_time   uint32     //>上次捐赠时间
	Last_dividend_time uint32     //>上次分红时间
	Last_login_time    uint32     //>上次登陆时间
	Last_logout_time   uint32     //>上次离线时间
	Forbid_end_time    uint32     //>禁言结束时间
	Status             uint32     //>成员玩家状态(参见logic_def的枚举guild_member_status)
	Vip                uint32     //>vip等级
	Reincarnation      uint32     //>转生等级
	Customs            CustomData //>自定义
	Job                uint8      //>职业
}

type GuildApplicantData struct { //>申请者信息
	Guid            uint64 //>申请人guid
	Name            string //>申请人姓名
	Level           uint32 //>申请人等级
	Role            uint8  //>申请人角色
	Sn              uint32 //>申请人串号
	Guild_contrb    uint32 //>帮贡
	Last_apply_time uint32 //>申请截止时间
	Vip             uint32 //>vip等级
	Reincarnation   uint32 //>转生等级
	Job             uint8  //>职业
}

type GuildEventData struct { //>帮派事件信息
	Event_type uint32  //>(参见logic_def的枚举guild_event_type)
	Event_time uint32  //>时间时间
	Event_msg  []uint8 //>事件消息
}

type GuildBriefData struct { //>帮派简略数据
	Guid             uint64 //>帮派guid
	Guild_id         uint32 //>帮派id
	Name             string //>帮派名称
	Leader_guid      uint64 //>帮主guid
	Leader_name      string //>帮主名称
	Leader_level     uint32 //>帮主等级
	Leader_role      uint8  //>帮主角色
	Level            uint32 //>帮派等级
	Member_count     uint16 //>帮派成员数量
	Max_member_count uint16 //>最大成员数量(对应当前等级)
	Declaration      string //>帮派宣言
	Applyed          uint8  //>是否申请：1是已申请， 0是未申请
	Leader_school    uint32 //>帮主门派ID
}

type GuildData struct { //>帮派数据
	Guid             uint64           //>帮派guid
	Guild_id         uint32           //>帮派id
	Name             string           //>帮派名称
	Leader_guid      uint64           //>帮主GUID
	Leader_name      string           //>帮主名称
	Level            uint32           //>帮派等级
	Base_level       uint32           //>帮派忠义堂等级
	Wing_room_level  uint32           //>帮派厢房/赏功堂等级
	Vault_level      uint32           //>帮派金库/帮派商店等级
	Academy_level    uint32           //>帮派书院/堂口等级
	Pharmacy_level   uint32           //>帮派宝阁等级
	Upgrade_cooldown uint32           //>帮派建筑升级冷却时间
	Member_count     uint16           //>帮派成员数量
	Max_member_count uint32           //>最大成员数量(对应当前等级)
	Online_count     uint16           //>在线成员数量
	Fund             uint32           //>帮派资金
	Liveness         uint32           //>活跃度
	Build_degree     uint32           //>帮派建设度
	Prestige         uint32           //>威望
	Fight_score      uint32           //>战绩
	Config           uint32           //>帮派设置
	Create_time      uint32           //>创建时间
	Dismissed_time   uint32           //>解散时间(为0表示不会解散，不为0表示解散最后时限)
	Impeach_time     uint32           //>弹劾帮主开始时间
	Impeach_guid     uint64           //>弹劾帮主者的GUID
	Declaration      string           //>帮派宣言
	Events           []GuildEventData //>帮派事件消息
	Customs          CustomData       //>自定义
}

type NoticeData struct { //>通知内容
	Notice_time uint32 //>发送通知时间
	Notice      string //>通知内容
}

type DBGuildData struct { //>帮派数据
	Guild  GuildData  //>帮派数据
	Notice NoticeData //>通知数据
	Items  []ItemData //>物品列表
}

type DBGuildMemberData struct { //>帮派成员数据
	Member            GuildMemberData //>帮派成员数据
	Guild_guid        uint64          //>帮派guid
	Total_fight_value int64           //>战力
}

type DBGuildApplicantData struct { //>帮派申请数据
	Applicant         GuildApplicantData //>帮派申请数据
	Guild_guid        uint64             //>帮派guid
	Total_fight_value int64              //>战力
}

type ContactData struct { //>联系人数据
	Guid              uint64     //>联系人GUID
	Contact_type      uint8      //>联系人类型：对应枚举ContactType的最近联系人、好友、黑名单
	Name              string     //>好友名字
	Role              uint8      //>role
	Job               uint8      //>职业
	Level             uint32     //>等级
	Friendship        uint32     //>好友度
	Last_contact_time uint32     //>最近联系时间
	Sn                int32      //>用户序列号
	Status            uint32     //>是否在线：1在线，0离线
	Vip               uint32     //>vip等级
	Reincarnation     uint32     //>转生等级
	Customs           CustomData //>自定义
}

type DBContactData struct { //>DB联系人数据
	Player_guid       uint64     //>玩家GUID
	Contact_guid      uint64     //>联系人GUID
	Contact_type      uint8      //>联系人类型：对应枚举ContactType的最近联系人、好友、黑名单
	Friendship        uint32     //>好友度
	Last_contact_time uint32     //>最近联系时间
	Customs           CustomData //>自定义
}

type MessageData struct { //>联系人消息数据
	Message    []uint8 //>发送信息内容
	Send_guid  uint64  //>发送者GUID
	Send_time  uint32  //>发送时间
	Recv_guid  uint64  //>接受者GUID
	Auto_reply uint8   //>是否自动回复消息
}

type ContactInfoData struct { //>联系人数据
	Guid              uint64 //>联系人GUID
	Name              string //>好友名字
	Role              uint8  //>role
	Job               uint8  //>职业
	Level             uint32 //>等级
	Sn                int32  //>用户序列号
	Status            uint32 //>是否在线：1在线，0离线
	Vip               uint32 //>vip等级
	Reincarnation     uint32 //>转生等级
	Last_login_time   uint32 //>上次登录时间
	Last_logout_time  uint32 //>上次登出时间
	Total_fight_value uint64 //>总战力
}

type MailHeadData struct { //>邮件标题信息
	Mail_guid uint64  //>邮件GUID
	Mail_type uint8   //>邮件类型 0:玩家邮件；1:系统邮件；2:帮派邮件；3:脚本邮件
	Title     []uint8 //>标题
	Send_guid uint64  //>发件人GUID
	Send_name string  //>发件人名字
	Send_time uint32  //>发件时间
	Status    uint32  //>邮件状态
}

type MailBodyData struct { //>邮件正文信息
	Content []uint8    //>邮件正文
	Attrs   []AttrData //>属性列表
	Items   []ItemData //>物品列表
	Pets    []PetData  //>宠物列表
}

type MailData struct { //>邮件数据
	Head MailHeadData //>邮件标题
	Body MailBodyData //>邮件内容
}

type DBMailData struct { //>邮件数据
	Recv_guid uint64       //>收件人GUID
	Guids     []uint64     //>邮件关联GUID列表
	Head      MailHeadData //>邮件标题
	Body      MailBodyData //>邮件内容
}

type DBVarData struct { //>DB自定义变量
	Key        []uint8 //>key值
	Value      []uint8 //>value
	Value_type uint8   //>值类型
	Merge_type uint8   //>合区类型
}

type RanklistData struct { //>排行榜数据
	Name       string //>上榜名称
	Rank_info  string //>排名信息
	Rank_data1 int64  //>排名数据1
	Rank_data2 int64  //>排名数据2
	Rank_data3 int64  //>排名数据3
}

type DBRanklistData struct { //>排行榜数据
	Name        string //>上榜名称
	Rank_info   string //>排名信息
	Rank_data1  int64  //>排名数据1
	Rank_data2  int64  //>排名数据2
	Rank_data3  int64  //>排名数据3
	Rank_guid   int64  //>排名guid
	Object_guid int64  //>排名对象guid
	Rank_type   int32  //>排行榜类型
}

type TeamBriefData struct { //>队伍简略信息
	Team_guid         uint64     //>队伍GUID
	Target            uint16     //>目标ID
	Min_require_level uint16     //>最低要求等级
	Max_require_level uint16     //>最高要求等级
	Member_count      uint16     //>成员数量
	Leader_data       MemberData //>队长数据
	Has_applicant     uint8      //>是否存在申请
	Min_require_reinc uint16     //>最低要求转生次数
	Max_require_reinc uint16     //>最高要求转生次数
	Customs           CustomData //>自定义
}

type TeamApplicantData struct { //>队伍申请信息
	Guid  uint64     //>玩家GUID
	Name  string     //>玩家名字
	Attrs []AttrData //>属性数据
}

type ExchangeData struct { //>商会数据
	Day     uint64        //>交易日期
	Records []KeyValueU32 //>交易记录
}

type ForbidData struct { //>禁止数据
	Sid     uint8   //>sid
	Keyword []uint8 //>禁止关键字
	Keytype uint8   //>1:角色名 2:帐号名 3:IP 4:硬件码
	Fbdtype uint8   //>1:禁止登陆
	Begin   int64   //>开始时间
	End     int64   //>结束时间
	Desc    []uint8 //>描述
}

type ForbidTalkData struct { //>禁止聊天数据
	Guid  uint64  //>角色guid
	Begin int64   //>开始时间
	End   int64   //>结束时间
	Desc  []uint8 //>描述
}

type PasturePetData struct { //>牧场宠物信息
	Owner_guid           uint64  //>拥有者GUID
	Pet                  PetData //>子区ID
	Map                  uint16  //>地图
	X                    uint16  //>x坐标
	Y                    uint16  //>y坐标
	Placed_time          uint32  //>放置时间
	Placed_status        uint32  //>放置状态
	Replaced_time        uint32  //>替换时间
	Replaced_player_guid uint64  //>替换的玩家GUID
	Replaced_player_name string  //>替换的玩家名字
	Replaced_pet_guid    uint64  //>替换的宠物GUID
	Replaced_pet_name    string  //>替换的宠物名字
}

type LadderData struct { //>天梯数据
	Player_guid uint64     //>玩家GUID
	Rank        uint32     //>排行
	Player      PlayerData //>玩家数据
}

type LadderRecordData struct { //>天梯记录数据
	Action      uint8  //>动作
	Time        uint32 //>时间
	Player_guid uint64 //>玩家GUID
	Player_name string //>玩家名字
	Player_rank uint32 //>排行
	Target_guid uint64 //>玩家GUID
	Target_name string //>玩家名字
	Target_rank uint32 //>排行
}

type PlayerDetail struct { //>玩家详细信息
	Uid         uint32            //>账号ID
	Sid         uint8             //>子区ID
	Guid        uint64            //>玩家GUID
	Unid        string            //>玩家UNID
	Name        string            //>玩家名字
	Attrs       []AttrData        //>属性数据
	Map         uint16            //>地图ID
	X           uint16            //>x坐标
	Y           uint16            //>y坐标
	State       uint8             //>玩家状态 0:彻底销毁 1:使用中 2:已删除
	Createtime  uint32            //>创建时间
	Lastlogin   uint32            //>最后一次登陆时间
	Destroytime uint32            //>销毁时间 state为0=销毁时间 state为2=删除时间
	Customs     CustomData        //>自定义
	Equips      ItemContainerData //>装备
}

type Pt struct { //>
	X uint16 //>x坐标
	Y uint16 //>y坐标
}

type ViolateData struct { //>违禁字数据
	Type    uint8   //>违禁字类型
	Violate []uint8 //>违禁字
}

type ChannelSwitchData struct { //>频道开关数据
	Type uint8 //>频道类型
	Sw   uint8 //>开关
}

type VerifyData struct { //>验证附加数据
	Fbd_time      uint64 //>封禁时间
	Accountstatus int32  //>账号状态 0=正常 1=禁用 2=临时密码使用中 3=异常 4=待销户 5=已销户
}

type IPPort struct { //>tcp连接信息
	Ip   string //>ip
	Port uint16 //>port
}

type CrossIPPort struct { //>跨服IP端口
	Gid        uint32 //>区id
	Gcp        uint16
	Gsp        uint16
	Innerip    string //>内网ip
	Extip      string //>外网ip
	Innerflag  string //>内网标记
	Clientname string //>区显示名称
}

type TestJsonName struct { //>Json测试
	Name []uint8 //>名称
}

type BillQueryData struct { //>充值记录
	Bill_id  uint32 //>订单ID
	User_id  uint32 //>用户ID
	Sid      uint32 //>用户SID
	Added_yb int32  //>用户充值元宝数
	Award_yb int32  //>用户奖励元宝数
	User     string //>用户名
	Desc     string //>充值描述
	Time     uint32 //>充值时间
}

func (proto *CustomStrData) GetMid() uint16 {
	return 100
}

func (proto *CustomStrData) GetPid() uint16 {
	return 1
}

func (proto *CustomStrData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteIntegerArray(b, proto.Key, uint16(65535)) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Value, uint32(4294967295)) {
		return false
	}

	return true
}

func (proto *CustomStrData) Read(b *bytes.Buffer) bool {
	if !ProtoReadIntegerArray(b, &proto.Key, uint16(65535)) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Value, uint32(4294967295)) {
		return false
	}

	return true
}

func (proto *CustomIntData) GetMid() uint16 {
	return 100
}

func (proto *CustomIntData) GetPid() uint16 {
	return 2
}

func (proto *CustomIntData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteIntegerArray(b, proto.Key, uint16(65535)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Value) {
		return false
	}

	return true
}

func (proto *CustomIntData) Read(b *bytes.Buffer) bool {
	if !ProtoReadIntegerArray(b, &proto.Key, uint16(65535)) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Value) {
		return false
	}

	return true
}

func (proto *CustomData) GetMid() uint16 {
	return 100
}

func (proto *CustomData) GetPid() uint16 {
	return 3
}

func (proto *CustomData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteCustomArray(b, proto.Strdata, uint32(4294967295)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Intdata, uint32(4294967295)) {
		return false
	}

	return true
}

func (proto *CustomData) Read(b *bytes.Buffer) bool {
	if !ProtoReadCustomArray(b, &proto.Strdata, uint32(4294967295)) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Intdata, uint32(4294967295)) {
		return false
	}

	return true
}

func (proto *SkillData) GetMid() uint16 {
	return 100
}

func (proto *SkillData) GetPid() uint16 {
	return 4
}

func (proto *SkillData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Performance) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Max_performance) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Tick) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Enable) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Bind) {
		return false
	}

	return true
}

func (proto *SkillData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Performance) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Max_performance) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Tick) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Enable) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Bind) {
		return false
	}

	return true
}

func (proto *DynAttrData) GetMid() uint16 {
	return 100
}

func (proto *DynAttrData) GetPid() uint16 {
	return 5
}

func (proto *DynAttrData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Mark) {
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

func (proto *DynAttrData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mark) {
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

func (proto *BuffData) GetMid() uint16 {
	return 100
}

func (proto *BuffData) GetPid() uint16 {
	return 6
}

func (proto *BuffData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Buff_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Start_times) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Duration) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Dyn_attrs, uint16(65535)) {
		return false
	}

	return true
}

func (proto *BuffData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Buff_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Start_times) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Duration) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Dyn_attrs, uint16(65535)) {
		return false
	}

	return true
}

func (proto *QuestData) GetMid() uint16 {
	return 100
}

func (proto *QuestData) GetPid() uint16 {
	return 7
}

func (proto *QuestData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Quest_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Quest_state) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Accept_npc) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Commit_npc) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Quest_store_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Accpet_time) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Limit_time) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Changed_time) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Ext, uint16(65535)) {
		return false
	}

	return true
}

func (proto *QuestData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Quest_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Quest_state) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Accept_npc) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Commit_npc) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Quest_store_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Accpet_time) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Limit_time) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Changed_time) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Ext, uint16(65535)) {
		return false
	}

	return true
}

func (proto *QuestRingData) GetMid() uint16 {
	return 100
}

func (proto *QuestRingData) GetPid() uint16 {
	return 8
}

func (proto *QuestRingData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Quest_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Quest_ring) {
		return false
	}

	return true
}

func (proto *QuestRingData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Quest_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Quest_ring) {
		return false
	}

	return true
}

func (proto *QuestStoreData) GetMid() uint16 {
	return 100
}

func (proto *QuestStoreData) GetPid() uint16 {
	return 9
}

func (proto *QuestStoreData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Quest_store_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Ring_auto) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Ring_day) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Ring_week) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Ring_display) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Quest_rings, uint16(65535)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Last_reset_day) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Last_reset_week) {
		return false
	}

	return true
}

func (proto *QuestStoreData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Quest_store_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Ring_auto) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Ring_day) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Ring_week) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Ring_display) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Quest_rings, uint16(65535)) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Last_reset_day) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Last_reset_week) {
		return false
	}

	return true
}

func (proto *QuestCountData) GetMid() uint16 {
	return 100
}

func (proto *QuestCountData) GetPid() uint16 {
	return 10
}

func (proto *QuestCountData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Quest_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Count) {
		return false
	}

	return true
}

func (proto *QuestCountData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Quest_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Count) {
		return false
	}

	return true
}

func (proto *QuestContainerData) GetMid() uint16 {
	return 100
}

func (proto *QuestContainerData) GetPid() uint16 {
	return 11
}

func (proto *QuestContainerData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteCustomArray(b, proto.Quests, uint16(65535)) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Npc_visible, uint16(65535)) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Npc_invisible, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Stores, uint16(65535)) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Abandoned, uint16(65535)) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Finished, uint16(65535)) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Activated, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Failed_counts, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Finished_counts, uint16(65535)) {
		return false
	}

	return true
}

func (proto *QuestContainerData) Read(b *bytes.Buffer) bool {
	if !ProtoReadCustomArray(b, &proto.Quests, uint16(65535)) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Npc_visible, uint16(65535)) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Npc_invisible, uint16(65535)) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Stores, uint16(65535)) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Abandoned, uint16(65535)) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Finished, uint16(65535)) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Activated, uint16(65535)) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Failed_counts, uint16(65535)) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Finished_counts, uint16(65535)) {
		return false
	}

	return true
}

func (proto *AttrData) GetMid() uint16 {
	return 100
}

func (proto *AttrData) GetPid() uint16 {
	return 12
}

func (proto *AttrData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Attr) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Value) {
		return false
	}

	return true
}

func (proto *AttrData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Attr) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Value) {
		return false
	}

	return true
}

func (proto *ItemData) GetMid() uint16 {
	return 100
}

func (proto *ItemData) GetPid() uint16 {
	return 13
}

func (proto *ItemData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Role) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Site) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Battlesite) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Amount) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Isbound) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Life) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Dyn_attrs, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Customs) {
		return false
	}

	return true
}

func (proto *ItemData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Role) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Site) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Battlesite) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Amount) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Isbound) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Life) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Dyn_attrs, uint16(65535)) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Customs) {
		return false
	}

	return true
}

func (proto *ItemContainerData) GetMid() uint16 {
	return 100
}

func (proto *ItemContainerData) GetPid() uint16 {
	return 14
}

func (proto *ItemContainerData) Write(b *bytes.Buffer) bool {
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

func (proto *ItemContainerData) Read(b *bytes.Buffer) bool {
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

func (proto *UserData) GetMid() uint16 {
	return 100
}

func (proto *UserData) GetPid() uint16 {
	return 15
}

func (proto *UserData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Uid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Sid) {
		return false
	}

	if !ProtoWriteString(b, proto.Uname, 32) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Billinyb) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Billoutyb) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Gameinyb) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Gameoutyb) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Createtime) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Lastlogin) {
		return false
	}

	return true
}

func (proto *UserData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Uid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Sid) {
		return false
	}

	if !ProtoReadString(b, &proto.Uname, 32) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Billinyb) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Billoutyb) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Gameinyb) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Gameoutyb) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Createtime) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Lastlogin) {
		return false
	}

	return true
}

func (proto *SceneData) GetMid() uint16 {
	return 100
}

func (proto *SceneData) GetPid() uint16 {
	return 16
}

func (proto *SceneData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Map_guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.X) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Y) {
		return false
	}

	return true
}

func (proto *SceneData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Map_guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.X) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Y) {
		return false
	}

	return true
}

func (proto *TicketData) GetMid() uint16 {
	return 100
}

func (proto *TicketData) GetPid() uint16 {
	return 17
}

func (proto *TicketData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Map_guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Create_time) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.X) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Y) {
		return false
	}

	return true
}

func (proto *TicketData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Map_guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Create_time) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.X) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Y) {
		return false
	}

	return true
}

func (proto *PetData) GetMid() uint16 {
	return 100
}

func (proto *PetData) GetPid() uint16 {
	return 18
}

func (proto *PetData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteString(b, proto.Name, 32) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Equips) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Attrs, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Extra_attrs, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Skills, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Buffs, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Dyn_attrs, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Customs) {
		return false
	}

	return true
}

func (proto *PetData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadString(b, &proto.Name, 32) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Equips) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Attrs, uint16(65535)) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Extra_attrs, uint16(65535)) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Skills, uint16(65535)) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Buffs, uint16(65535)) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Dyn_attrs, uint16(65535)) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Customs) {
		return false
	}

	return true
}

func (proto *PetContainerData) GetMid() uint16 {
	return 100
}

func (proto *PetContainerData) GetPid() uint16 {
	return 19
}

func (proto *PetContainerData) Write(b *bytes.Buffer) bool {
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

func (proto *PetContainerData) Read(b *bytes.Buffer) bool {
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

func (proto *TitleData) GetMid() uint16 {
	return 100
}

func (proto *TitleData) GetPid() uint16 {
	return 20
}

func (proto *TitleData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Title_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.End_time) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Custom) {
		return false
	}

	return true
}

func (proto *TitleData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Title_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.End_time) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Custom) {
		return false
	}

	return true
}

func (proto *TitleContainerData) GetMid() uint16 {
	return 100
}

func (proto *TitleContainerData) GetPid() uint16 {
	return 21
}

func (proto *TitleContainerData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteCustomArray(b, proto.Titles, uint16(65535)) {
		return false
	}

	return true
}

func (proto *TitleContainerData) Read(b *bytes.Buffer) bool {
	if !ProtoReadCustomArray(b, &proto.Titles, uint16(65535)) {
		return false
	}

	return true
}

func (proto *SeatBasic) GetMid() uint16 {
	return 100
}

func (proto *SeatBasic) GetPid() uint16 {
	return 22
}

func (proto *SeatBasic) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Level) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Score) {
		return false
	}

	return true
}

func (proto *SeatBasic) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Level) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Score) {
		return false
	}

	return true
}

func (proto *SeatData) GetMid() uint16 {
	return 100
}

func (proto *SeatData) GetPid() uint16 {
	return 23
}

func (proto *SeatData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteCustomArray(b, proto.Seats, uint8(255)) {
		return false
	}

	return true
}

func (proto *SeatData) Read(b *bytes.Buffer) bool {
	if !ProtoReadCustomArray(b, &proto.Seats, uint8(255)) {
		return false
	}

	return true
}

func (proto *LineupBasic) GetMid() uint16 {
	return 100
}

func (proto *LineupBasic) GetPid() uint16 {
	return 24
}

func (proto *LineupBasic) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Seat_id) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Lineup, uint8(255)) {
		return false
	}

	return true
}

func (proto *LineupBasic) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Seat_id) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Lineup, uint8(255)) {
		return false
	}

	return true
}

func (proto *LineupData) GetMid() uint16 {
	return 100
}

func (proto *LineupData) GetPid() uint16 {
	return 25
}

func (proto *LineupData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Index) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Lineups, uint8(255)) {
		return false
	}

	return true
}

func (proto *LineupData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Index) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Lineups, uint8(255)) {
		return false
	}

	return true
}

func (proto *InstructionBasic) GetMid() uint16 {
	return 100
}

func (proto *InstructionBasic) GetPid() uint16 {
	return 26
}

func (proto *InstructionBasic) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Type) {
		return false
	}

	if !ProtoWriteString(b, proto.Content, 32) {
		return false
	}

	return true
}

func (proto *InstructionBasic) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Type) {
		return false
	}

	if !ProtoReadString(b, &proto.Content, 32) {
		return false
	}

	return true
}

func (proto *PetLineupData) GetMid() uint16 {
	return 100
}

func (proto *PetLineupData) GetPid() uint16 {
	return 27
}

func (proto *PetLineupData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteIntegerArray(b, proto.Lineup, uint8(255)) {
		return false
	}

	return true
}

func (proto *PetLineupData) Read(b *bytes.Buffer) bool {
	if !ProtoReadIntegerArray(b, &proto.Lineup, uint8(255)) {
		return false
	}

	return true
}

func (proto *InstructionData) GetMid() uint16 {
	return 100
}

func (proto *InstructionData) GetPid() uint16 {
	return 28
}

func (proto *InstructionData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteCustomArray(b, proto.Frienddata, uint8(255)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Enemydata, uint8(255)) {
		return false
	}

	return true
}

func (proto *InstructionData) Read(b *bytes.Buffer) bool {
	if !ProtoReadCustomArray(b, &proto.Frienddata, uint8(255)) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Enemydata, uint8(255)) {
		return false
	}

	return true
}

func (proto *EmojiData) GetMid() uint16 {
	return 100
}

func (proto *EmojiData) GetPid() uint16 {
	return 29
}

func (proto *EmojiData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteIntegerArray(b, proto.Emojis, uint32(4294967295)) {
		return false
	}

	return true
}

func (proto *EmojiData) Read(b *bytes.Buffer) bool {
	if !ProtoReadIntegerArray(b, &proto.Emojis, uint32(4294967295)) {
		return false
	}

	return true
}

func (proto *PlayerBrief) GetMid() uint16 {
	return 100
}

func (proto *PlayerBrief) GetPid() uint16 {
	return 30
}

func (proto *PlayerBrief) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Uid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Sid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteString(b, proto.Unid, 32) {
		return false
	}

	if !ProtoWriteString(b, proto.Name, 32) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Attrs, uint16(65535)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Map) {
		return false
	}

	if !ProtoWriteInteger(b, proto.X) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Y) {
		return false
	}

	if !ProtoWriteInteger(b, proto.State) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Createtime) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Lastlogin) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Destroytime) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Customs) {
		return false
	}

	return true
}

func (proto *PlayerBrief) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Uid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Sid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadString(b, &proto.Unid, 32) {
		return false
	}

	if !ProtoReadString(b, &proto.Name, 32) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Attrs, uint16(65535)) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Map) {
		return false
	}

	if !ProtoReadInteger(b, &proto.X) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Y) {
		return false
	}

	if !ProtoReadInteger(b, &proto.State) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Createtime) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Lastlogin) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Destroytime) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Customs) {
		return false
	}

	return true
}

func (proto *GuardData) GetMid() uint16 {
	return 100
}

func (proto *GuardData) GetPid() uint16 {
	return 31
}

func (proto *GuardData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Equips) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Attrs, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Extra_attrs, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Skills, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Buffs, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Dyn_attrs, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Customs) {
		return false
	}

	return true
}

func (proto *GuardData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Equips) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Attrs, uint16(65535)) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Extra_attrs, uint16(65535)) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Skills, uint16(65535)) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Buffs, uint16(65535)) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Dyn_attrs, uint16(65535)) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Customs) {
		return false
	}

	return true
}

func (proto *GuardContainerData) GetMid() uint16 {
	return 100
}

func (proto *GuardContainerData) GetPid() uint16 {
	return 32
}

func (proto *GuardContainerData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteCustomArray(b, proto.Guards, uint16(65535)) {
		return false
	}

	return true
}

func (proto *GuardContainerData) Read(b *bytes.Buffer) bool {
	if !ProtoReadCustomArray(b, &proto.Guards, uint16(65535)) {
		return false
	}

	return true
}

func (proto *PlayerData) GetMid() uint16 {
	return 100
}

func (proto *PlayerData) GetPid() uint16 {
	return 33
}

func (proto *PlayerData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Uid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Sid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteString(b, proto.Unid, 32) {
		return false
	}

	if !ProtoWriteString(b, proto.Name, 32) {
		return false
	}

	if !ProtoWriteInteger(b, proto.State) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Createtime) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Destroytime) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Lastlogin) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Lastloginip) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Lastlogout) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Onlines) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Cur_map) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Last_common_map) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Attrs, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Dyn_attrs, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Extra_attrs, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Items, uint8(255)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Skills, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Buffs, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Quests) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Pets, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Guards) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Tickets, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Customs) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Gameinyb) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Gameoutyb) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Billinyb) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Billoutyb) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Titles) {
		return false
	}

	if !ProtoWriteString(b, proto.Username, 255) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Lineups) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Seats) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Instructions) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Pet_lineups) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Emojis) {
		return false
	}

	return true
}

func (proto *PlayerData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Uid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Sid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadString(b, &proto.Unid, 32) {
		return false
	}

	if !ProtoReadString(b, &proto.Name, 32) {
		return false
	}

	if !ProtoReadInteger(b, &proto.State) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Createtime) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Destroytime) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Lastlogin) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Lastloginip) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Lastlogout) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Onlines) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Cur_map) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Last_common_map) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Attrs, uint16(65535)) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Dyn_attrs, uint16(65535)) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Extra_attrs, uint16(65535)) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Items, uint8(255)) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Skills, uint16(65535)) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Buffs, uint16(65535)) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Quests) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Pets, uint16(65535)) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Guards) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Tickets, uint16(65535)) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Customs) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Gameinyb) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Gameoutyb) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Billinyb) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Billoutyb) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Titles) {
		return false
	}

	if !ProtoReadString(b, &proto.Username, 255) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Lineups) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Seats) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Instructions) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Pet_lineups) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Emojis) {
		return false
	}

	return true
}

func (proto *MapRegion) GetMid() uint16 {
	return 100
}

func (proto *MapRegion) GetPid() uint16 {
	return 34
}

func (proto *MapRegion) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Min_x) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Max_x) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Min_y) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Max_y) {
		return false
	}

	return true
}

func (proto *MapRegion) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Min_x) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Max_x) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Min_y) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Max_y) {
		return false
	}

	return true
}

func (proto *ItemAttrValue64) GetMid() uint16 {
	return 100
}

func (proto *ItemAttrValue64) GetPid() uint16 {
	return 35
}

func (proto *ItemAttrValue64) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Attr_type) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Attr_value) {
		return false
	}

	return true
}

func (proto *ItemAttrValue64) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Attr_type) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Attr_value) {
		return false
	}

	return true
}

func (proto *ItemAttrValueS) GetMid() uint16 {
	return 100
}

func (proto *ItemAttrValueS) GetPid() uint16 {
	return 36
}

func (proto *ItemAttrValueS) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Attr_type) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Attr_value, uint16(65535)) {
		return false
	}

	return true
}

func (proto *ItemAttrValueS) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Attr_type) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Attr_value, uint16(65535)) {
		return false
	}

	return true
}

func (proto *ItemAttrValueS64) GetMid() uint16 {
	return 100
}

func (proto *ItemAttrValueS64) GetPid() uint16 {
	return 37
}

func (proto *ItemAttrValueS64) Write(b *bytes.Buffer) bool {
	if !ProtoWriteString(b, proto.Attr_key, 255) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Attr_value) {
		return false
	}

	return true
}

func (proto *ItemAttrValueS64) Read(b *bytes.Buffer) bool {
	if !ProtoReadString(b, &proto.Attr_key, 255) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Attr_value) {
		return false
	}

	return true
}

func (proto *ItemAttrValueSS) GetMid() uint16 {
	return 100
}

func (proto *ItemAttrValueSS) GetPid() uint16 {
	return 38
}

func (proto *ItemAttrValueSS) Write(b *bytes.Buffer) bool {
	if !ProtoWriteString(b, proto.Attr_key, 255) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Attr_value, uint16(65535)) {
		return false
	}

	return true
}

func (proto *ItemAttrValueSS) Read(b *bytes.Buffer) bool {
	if !ProtoReadString(b, &proto.Attr_key, 255) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Attr_value, uint16(65535)) {
		return false
	}

	return true
}

func (proto *ItemAttrValueList) GetMid() uint16 {
	return 100
}

func (proto *ItemAttrValueList) GetPid() uint16 {
	return 39
}

func (proto *ItemAttrValueList) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Item_guid) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Item_values_64, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Item_values_s, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Item_values_s64, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Item_values_ss, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Dyn_attrs, uint16(65535)) {
		return false
	}

	return true
}

func (proto *ItemAttrValueList) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Item_guid) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Item_values_64, uint16(65535)) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Item_values_s, uint16(65535)) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Item_values_s64, uint16(65535)) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Item_values_ss, uint16(65535)) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Dyn_attrs, uint16(65535)) {
		return false
	}

	return true
}

func (proto *ChatMsg) GetMid() uint16 {
	return 100
}

func (proto *ChatMsg) GetPid() uint16 {
	return 40
}

func (proto *ChatMsg) Write(b *bytes.Buffer) bool {
	if !ProtoWriteString(b, proto.Role_name, 255) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Role_id) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Message, uint16(65535)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Vip) {
		return false
	}

	return true
}

func (proto *ChatMsg) Read(b *bytes.Buffer) bool {
	if !ProtoReadString(b, &proto.Role_name, 255) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Role_id) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Message, uint16(65535)) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Vip) {
		return false
	}

	return true
}

func (proto *ParamData) GetMid() uint16 {
	return 100
}

func (proto *ParamData) GetPid() uint16 {
	return 41
}

func (proto *ParamData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Type) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Param, uint16(65535)) {
		return false
	}

	return true
}

func (proto *ParamData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Type) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Param, uint16(65535)) {
		return false
	}

	return true
}

func (proto *FightPetData) GetMid() uint16 {
	return 100
}

func (proto *FightPetData) GetPid() uint16 {
	return 42
}

func (proto *FightPetData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.State) {
		return false
	}

	return true
}

func (proto *FightPetData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.State) {
		return false
	}

	return true
}

func (proto *FightBuffData) GetMid() uint16 {
	return 100
}

func (proto *FightBuffData) GetPid() uint16 {
	return 43
}

func (proto *FightBuffData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Stack) {
		return false
	}

	return true
}

func (proto *FightBuffData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Stack) {
		return false
	}

	return true
}

func (proto *FighterData) GetMid() uint16 {
	return 100
}

func (proto *FighterData) GetPid() uint16 {
	return 44
}

func (proto *FighterData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Site) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Fighter_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Type) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Template_id) {
		return false
	}

	if !ProtoWriteString(b, proto.Name, 32) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Attrs, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Pets, uint8(255)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Buffs, uint8(255)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Model) {
		return false
	}

	if !ProtoWriteString(b, proto.Friend_instruction, 32) {
		return false
	}

	if !ProtoWriteString(b, proto.Enemy_instruction, 32) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Customs) {
		return false
	}

	return true
}

func (proto *FighterData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Site) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Fighter_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Type) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Template_id) {
		return false
	}

	if !ProtoReadString(b, &proto.Name, 32) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Attrs, uint16(65535)) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Pets, uint8(255)) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Buffs, uint8(255)) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Model) {
		return false
	}

	if !ProtoReadString(b, &proto.Friend_instruction, 32) {
		return false
	}

	if !ProtoReadString(b, &proto.Enemy_instruction, 32) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Customs) {
		return false
	}

	return true
}

func (proto *FightGroupData) GetMid() uint16 {
	return 100
}

func (proto *FightGroupData) GetPid() uint16 {
	return 45
}

func (proto *FightGroupData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Type) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Lineup_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Lineup_lv) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Fighters, uint8(255)) {
		return false
	}

	return true
}

func (proto *FightGroupData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Type) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Lineup_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Lineup_lv) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Fighters, uint8(255)) {
		return false
	}

	return true
}

func (proto *FightOperateData) GetMid() uint16 {
	return 100
}

func (proto *FightOperateData) GetPid() uint16 {
	return 46
}

func (proto *FightOperateData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Actor) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Operate) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Target) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Param) {
		return false
	}

	return true
}

func (proto *FightOperateData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Actor) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Operate) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Target) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Param) {
		return false
	}

	return true
}

func (proto *MemberData) GetMid() uint16 {
	return 100
}

func (proto *MemberData) GetPid() uint16 {
	return 47
}

func (proto *MemberData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteString(b, proto.Name, 32) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Temp_leave) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Attrs, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Customs) {
		return false
	}

	return true
}

func (proto *MemberData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadString(b, &proto.Name, 32) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Temp_leave) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Attrs, uint16(65535)) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Customs) {
		return false
	}

	return true
}

func (proto *KeyValueU32) GetMid() uint16 {
	return 100
}

func (proto *KeyValueU32) GetPid() uint16 {
	return 48
}

func (proto *KeyValueU32) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Key) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Value) {
		return false
	}

	return true
}

func (proto *KeyValueU32) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Key) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Value) {
		return false
	}

	return true
}

func (proto *KeyValueU64) GetMid() uint16 {
	return 100
}

func (proto *KeyValueU64) GetPid() uint16 {
	return 49
}

func (proto *KeyValueU64) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Key) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Value) {
		return false
	}

	return true
}

func (proto *KeyValueU64) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Key) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Value) {
		return false
	}

	return true
}

func (proto *KeyValueS) GetMid() uint16 {
	return 100
}

func (proto *KeyValueS) GetPid() uint16 {
	return 50
}

func (proto *KeyValueS) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Key) {
		return false
	}

	if !ProtoWriteString(b, proto.Value, 255) {
		return false
	}

	return true
}

func (proto *KeyValueS) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Key) {
		return false
	}

	if !ProtoReadString(b, &proto.Value, 255) {
		return false
	}

	return true
}

func (proto *OptionData) GetMid() uint16 {
	return 100
}

func (proto *OptionData) GetPid() uint16 {
	return 51
}

func (proto *OptionData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Mark) {
		return false
	}

	if !ProtoWriteString(b, proto.Text, 255) {
		return false
	}

	return true
}

func (proto *OptionData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Mark) {
		return false
	}

	if !ProtoReadString(b, &proto.Text, 255) {
		return false
	}

	return true
}

func (proto *ShopItemData) GetMid() uint16 {
	return 100
}

func (proto *ShopItemData) GetPid() uint16 {
	return 52
}

func (proto *ShopItemData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Price_type) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Price) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Bind) {
		return false
	}

	return true
}

func (proto *ShopItemData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Price_type) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Price) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Bind) {
		return false
	}

	return true
}

func (proto *SellItemData) GetMid() uint16 {
	return 100
}

func (proto *SellItemData) GetPid() uint16 {
	return 53
}

func (proto *SellItemData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Amount) {
		return false
	}

	return true
}

func (proto *SellItemData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Amount) {
		return false
	}

	return true
}

func (proto *Col1Data) GetMid() uint16 {
	return 100
}

func (proto *Col1Data) GetPid() uint16 {
	return 54
}

func (proto *Col1Data) Write(b *bytes.Buffer) bool {
	if !ProtoWriteString(b, proto.Col1, 255) {
		return false
	}

	return true
}

func (proto *Col1Data) Read(b *bytes.Buffer) bool {
	if !ProtoReadString(b, &proto.Col1, 255) {
		return false
	}

	return true
}

func (proto *Col2Data) GetMid() uint16 {
	return 100
}

func (proto *Col2Data) GetPid() uint16 {
	return 55
}

func (proto *Col2Data) Write(b *bytes.Buffer) bool {
	if !ProtoWriteString(b, proto.Col1, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col2, 255) {
		return false
	}

	return true
}

func (proto *Col2Data) Read(b *bytes.Buffer) bool {
	if !ProtoReadString(b, &proto.Col1, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col2, 255) {
		return false
	}

	return true
}

func (proto *Col3Data) GetMid() uint16 {
	return 100
}

func (proto *Col3Data) GetPid() uint16 {
	return 56
}

func (proto *Col3Data) Write(b *bytes.Buffer) bool {
	if !ProtoWriteString(b, proto.Col1, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col2, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col3, 255) {
		return false
	}

	return true
}

func (proto *Col3Data) Read(b *bytes.Buffer) bool {
	if !ProtoReadString(b, &proto.Col1, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col2, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col3, 255) {
		return false
	}

	return true
}

func (proto *Col4Data) GetMid() uint16 {
	return 100
}

func (proto *Col4Data) GetPid() uint16 {
	return 57
}

func (proto *Col4Data) Write(b *bytes.Buffer) bool {
	if !ProtoWriteString(b, proto.Col1, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col2, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col3, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col4, 255) {
		return false
	}

	return true
}

func (proto *Col4Data) Read(b *bytes.Buffer) bool {
	if !ProtoReadString(b, &proto.Col1, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col2, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col3, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col4, 255) {
		return false
	}

	return true
}

func (proto *Col5Data) GetMid() uint16 {
	return 100
}

func (proto *Col5Data) GetPid() uint16 {
	return 58
}

func (proto *Col5Data) Write(b *bytes.Buffer) bool {
	if !ProtoWriteString(b, proto.Col1, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col2, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col3, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col4, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col5, 255) {
		return false
	}

	return true
}

func (proto *Col5Data) Read(b *bytes.Buffer) bool {
	if !ProtoReadString(b, &proto.Col1, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col2, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col3, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col4, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col5, 255) {
		return false
	}

	return true
}

func (proto *Col6Data) GetMid() uint16 {
	return 100
}

func (proto *Col6Data) GetPid() uint16 {
	return 59
}

func (proto *Col6Data) Write(b *bytes.Buffer) bool {
	if !ProtoWriteString(b, proto.Col1, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col2, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col3, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col4, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col5, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col6, 255) {
		return false
	}

	return true
}

func (proto *Col6Data) Read(b *bytes.Buffer) bool {
	if !ProtoReadString(b, &proto.Col1, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col2, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col3, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col4, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col5, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col6, 255) {
		return false
	}

	return true
}

func (proto *Col7Data) GetMid() uint16 {
	return 100
}

func (proto *Col7Data) GetPid() uint16 {
	return 60
}

func (proto *Col7Data) Write(b *bytes.Buffer) bool {
	if !ProtoWriteString(b, proto.Col1, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col2, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col3, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col4, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col5, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col6, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col7, 255) {
		return false
	}

	return true
}

func (proto *Col7Data) Read(b *bytes.Buffer) bool {
	if !ProtoReadString(b, &proto.Col1, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col2, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col3, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col4, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col5, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col6, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col7, 255) {
		return false
	}

	return true
}

func (proto *Col8Data) GetMid() uint16 {
	return 100
}

func (proto *Col8Data) GetPid() uint16 {
	return 61
}

func (proto *Col8Data) Write(b *bytes.Buffer) bool {
	if !ProtoWriteString(b, proto.Col1, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col2, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col3, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col4, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col5, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col6, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col7, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col8, 255) {
		return false
	}

	return true
}

func (proto *Col8Data) Read(b *bytes.Buffer) bool {
	if !ProtoReadString(b, &proto.Col1, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col2, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col3, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col4, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col5, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col6, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col7, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col8, 255) {
		return false
	}

	return true
}

func (proto *Col9Data) GetMid() uint16 {
	return 100
}

func (proto *Col9Data) GetPid() uint16 {
	return 62
}

func (proto *Col9Data) Write(b *bytes.Buffer) bool {
	if !ProtoWriteString(b, proto.Col1, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col2, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col3, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col4, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col5, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col6, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col7, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col8, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col9, 255) {
		return false
	}

	return true
}

func (proto *Col9Data) Read(b *bytes.Buffer) bool {
	if !ProtoReadString(b, &proto.Col1, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col2, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col3, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col4, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col5, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col6, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col7, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col8, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col9, 255) {
		return false
	}

	return true
}

func (proto *Col10Data) GetMid() uint16 {
	return 100
}

func (proto *Col10Data) GetPid() uint16 {
	return 63
}

func (proto *Col10Data) Write(b *bytes.Buffer) bool {
	if !ProtoWriteString(b, proto.Col1, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col2, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col3, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col4, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col5, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col6, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col7, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col8, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col9, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Col10, 255) {
		return false
	}

	return true
}

func (proto *Col10Data) Read(b *bytes.Buffer) bool {
	if !ProtoReadString(b, &proto.Col1, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col2, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col3, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col4, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col5, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col6, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col7, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col8, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col9, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Col10, 255) {
		return false
	}

	return true
}

func (proto *TalkNodeData) GetMid() uint16 {
	return 100
}

func (proto *TalkNodeData) GetPid() uint16 {
	return 64
}

func (proto *TalkNodeData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteString(b, proto.Node_type, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Node_id, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Display_text, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Connect_text, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Npc_id, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Npc_emotion, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Npc_talk, 512) {
		return false
	}

	if !ProtoWriteString(b, proto.Npc_template, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Func_type, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Link_node, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Belong_to, 255) {
		return false
	}

	if !ProtoWriteString(b, proto.Player_talk, 255) {
		return false
	}

	return true
}

func (proto *TalkNodeData) Read(b *bytes.Buffer) bool {
	if !ProtoReadString(b, &proto.Node_type, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Node_id, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Display_text, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Connect_text, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Npc_id, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Npc_emotion, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Npc_talk, 512) {
		return false
	}

	if !ProtoReadString(b, &proto.Npc_template, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Func_type, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Link_node, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Belong_to, 255) {
		return false
	}

	if !ProtoReadString(b, &proto.Player_talk, 255) {
		return false
	}

	return true
}

func (proto *QuestClientData) GetMid() uint16 {
	return 100
}

func (proto *QuestClientData) GetPid() uint16 {
	return 65
}

func (proto *QuestClientData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Quest_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Quest_type) {
		return false
	}

	if !ProtoWriteString(b, proto.Quest_type_name, 128) {
		return false
	}

	if !ProtoWriteString(b, proto.Quest_name, 32) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Quest_color) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Quest_state) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Display_location) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Display_accept_tip) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Display_accomplish_tip) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Can_track) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Can_abandon) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Accpet_time) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Limit_time) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Changed_time) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Accept_npc) {
		return false
	}

	if !ProtoWriteString(b, proto.Accept_npc_info, 256) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Commit_npc) {
		return false
	}

	if !ProtoWriteString(b, proto.Commit_npc_info, 256) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Quest_goal_type) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Quest_goal, uint16(65535)) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Quest_award, uint16(65535)) {
		return false
	}

	if !ProtoWriteString(b, proto.Track_desc, 512) {
		return false
	}

	if !ProtoWriteString(b, proto.Cant_accept_track_desc, 512) {
		return false
	}

	if !ProtoWriteString(b, proto.Cant_accept_desc, 512) {
		return false
	}

	if !ProtoWriteString(b, proto.Can_accept_desc, 512) {
		return false
	}

	if !ProtoWriteString(b, proto.Accepted_desc, 512) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Is_display_cycle_num) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Finish_cycle_num) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Total_cycle_num) {
		return false
	}

	return true
}

func (proto *QuestClientData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Quest_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Quest_type) {
		return false
	}

	if !ProtoReadString(b, &proto.Quest_type_name, 128) {
		return false
	}

	if !ProtoReadString(b, &proto.Quest_name, 32) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Quest_color) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Quest_state) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Display_location) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Display_accept_tip) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Display_accomplish_tip) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Can_track) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Can_abandon) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Accpet_time) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Limit_time) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Changed_time) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Accept_npc) {
		return false
	}

	if !ProtoReadString(b, &proto.Accept_npc_info, 256) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Commit_npc) {
		return false
	}

	if !ProtoReadString(b, &proto.Commit_npc_info, 256) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Quest_goal_type) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Quest_goal, uint16(65535)) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Quest_award, uint16(65535)) {
		return false
	}

	if !ProtoReadString(b, &proto.Track_desc, 512) {
		return false
	}

	if !ProtoReadString(b, &proto.Cant_accept_track_desc, 512) {
		return false
	}

	if !ProtoReadString(b, &proto.Cant_accept_desc, 512) {
		return false
	}

	if !ProtoReadString(b, &proto.Can_accept_desc, 512) {
		return false
	}

	if !ProtoReadString(b, &proto.Accepted_desc, 512) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Is_display_cycle_num) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Finish_cycle_num) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Total_cycle_num) {
		return false
	}

	return true
}

func (proto *QuestTalkData) GetMid() uint16 {
	return 100
}

func (proto *QuestTalkData) GetPid() uint16 {
	return 66
}

func (proto *QuestTalkData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Npc_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Node_id) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Talk_text, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Options, uint8(255)) {
		return false
	}

	return true
}

func (proto *QuestTalkData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Npc_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Node_id) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Talk_text, uint16(65535)) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Options, uint8(255)) {
		return false
	}

	return true
}

func (proto *IDCountData) GetMid() uint16 {
	return 100
}

func (proto *IDCountData) GetPid() uint16 {
	return 67
}

func (proto *IDCountData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Amount) {
		return false
	}

	return true
}

func (proto *IDCountData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Amount) {
		return false
	}

	return true
}

func (proto *GUIDCountData) GetMid() uint16 {
	return 100
}

func (proto *GUIDCountData) GetPid() uint16 {
	return 68
}

func (proto *GUIDCountData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Amount) {
		return false
	}

	return true
}

func (proto *GUIDCountData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Amount) {
		return false
	}

	return true
}

func (proto *SkillTipsData) GetMid() uint16 {
	return 100
}

func (proto *SkillTipsData) GetPid() uint16 {
	return 69
}

func (proto *SkillTipsData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Name) {
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

func (proto *SkillTipsData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Name) {
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

func (proto *GoalVisitNpcData) GetMid() uint16 {
	return 100
}

func (proto *GoalVisitNpcData) GetPid() uint16 {
	return 70
}

func (proto *GoalVisitNpcData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Npc_id) {
		return false
	}

	return true
}

func (proto *GoalVisitNpcData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Npc_id) {
		return false
	}

	return true
}

func (proto *GoalKillBossData) GetMid() uint16 {
	return 100
}

func (proto *GoalKillBossData) GetPid() uint16 {
	return 71
}

func (proto *GoalKillBossData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Boss_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Map_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Loc_x) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Loc_y) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Kill_count) {
		return false
	}

	return true
}

func (proto *GoalKillBossData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Boss_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Map_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Loc_x) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Loc_y) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Kill_count) {
		return false
	}

	return true
}

func (proto *GoalCommitItemData) GetMid() uint16 {
	return 100
}

func (proto *GoalCommitItemData) GetPid() uint16 {
	return 72
}

func (proto *GoalCommitItemData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Item_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Item_count) {
		return false
	}

	return true
}

func (proto *GoalCommitItemData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Item_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Item_count) {
		return false
	}

	return true
}

func (proto *GoalCommitPetData) GetMid() uint16 {
	return 100
}

func (proto *GoalCommitPetData) GetPid() uint16 {
	return 73
}

func (proto *GoalCommitPetData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Pet_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Pet_count) {
		return false
	}

	return true
}

func (proto *GoalCommitPetData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Pet_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pet_count) {
		return false
	}

	return true
}

func (proto *GoalEscortNpcData) GetMid() uint16 {
	return 100
}

func (proto *GoalEscortNpcData) GetPid() uint16 {
	return 74
}

func (proto *GoalEscortNpcData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Npc_id) {
		return false
	}

	return true
}

func (proto *GoalEscortNpcData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Npc_id) {
		return false
	}

	return true
}

func (proto *GoalArriveAddrData) GetMid() uint16 {
	return 100
}

func (proto *GoalArriveAddrData) GetPid() uint16 {
	return 75
}

func (proto *GoalArriveAddrData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Map_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Loc_x) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Loc_y) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Range) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Item_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Npc_id) {
		return false
	}

	if !ProtoWriteString(b, proto.Desc, 255) {
		return false
	}

	return true
}

func (proto *GoalArriveAddrData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Map_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Loc_x) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Loc_y) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Range) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Item_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Npc_id) {
		return false
	}

	if !ProtoReadString(b, &proto.Desc, 255) {
		return false
	}

	return true
}

func (proto *GoalFightMapData) GetMid() uint16 {
	return 100
}

func (proto *GoalFightMapData) GetPid() uint16 {
	return 76
}

func (proto *GoalFightMapData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Map_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Count) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Count_required) {
		return false
	}

	return true
}

func (proto *GoalFightMapData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Map_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Count) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Count_required) {
		return false
	}

	return true
}

func (proto *GoalFightMapCountData) GetMid() uint16 {
	return 100
}

func (proto *GoalFightMapCountData) GetPid() uint16 {
	return 77
}

func (proto *GoalFightMapCountData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteCustomArray(b, proto.Fight_map, uint8(255)) {
		return false
	}

	return true
}

func (proto *GoalFightMapCountData) Read(b *bytes.Buffer) bool {
	if !ProtoReadCustomArray(b, &proto.Fight_map, uint8(255)) {
		return false
	}

	return true
}

func (proto *GoalKillMonsterData) GetMid() uint16 {
	return 100
}

func (proto *GoalKillMonsterData) GetPid() uint16 {
	return 78
}

func (proto *GoalKillMonsterData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Monster_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Count) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Count_required) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Map_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Left) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Top) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Width) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Height) {
		return false
	}

	return true
}

func (proto *GoalKillMonsterData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Monster_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Count) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Count_required) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Map_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Left) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Top) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Width) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Height) {
		return false
	}

	return true
}

func (proto *GoalKillMonsterCountData) GetMid() uint16 {
	return 100
}

func (proto *GoalKillMonsterCountData) GetPid() uint16 {
	return 79
}

func (proto *GoalKillMonsterCountData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteCustomArray(b, proto.Kill_monster, uint8(255)) {
		return false
	}

	return true
}

func (proto *GoalKillMonsterCountData) Read(b *bytes.Buffer) bool {
	if !ProtoReadCustomArray(b, &proto.Kill_monster, uint8(255)) {
		return false
	}

	return true
}

func (proto *GoalEncounterFightData) GetMid() uint16 {
	return 100
}

func (proto *GoalEncounterFightData) GetPid() uint16 {
	return 80
}

func (proto *GoalEncounterFightData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Cur_num) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Max_num) {
		return false
	}

	return true
}

func (proto *GoalEncounterFightData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Cur_num) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Max_num) {
		return false
	}

	return true
}

func (proto *GoalUseItemData) GetMid() uint16 {
	return 100
}

func (proto *GoalUseItemData) GetPid() uint16 {
	return 81
}

func (proto *GoalUseItemData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Map_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Loc_x) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Loc_y) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Range) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Item_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Npc_id) {
		return false
	}

	if !ProtoWriteString(b, proto.Desc, 255) {
		return false
	}

	return true
}

func (proto *GoalUseItemData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Map_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Loc_x) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Loc_y) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Range) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Item_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Npc_id) {
		return false
	}

	if !ProtoReadString(b, &proto.Desc, 255) {
		return false
	}

	return true
}

func (proto *GoalCapturePetData) GetMid() uint16 {
	return 100
}

func (proto *GoalCapturePetData) GetPid() uint16 {
	return 82
}

func (proto *GoalCapturePetData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Pet_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Count) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Count_required) {
		return false
	}

	return true
}

func (proto *GoalCapturePetData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Pet_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Count) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Count_required) {
		return false
	}

	return true
}

func (proto *GoalPVEData) GetMid() uint16 {
	return 100
}

func (proto *GoalPVEData) GetPid() uint16 {
	return 83
}

func (proto *GoalPVEData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Win_cur) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Win_required) {
		return false
	}

	return true
}

func (proto *GoalPVEData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Win_cur) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Win_required) {
		return false
	}

	return true
}

func (proto *IPConfig) GetMid() uint16 {
	return 100
}

func (proto *IPConfig) GetPid() uint16 {
	return 84
}

func (proto *IPConfig) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Type) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Port) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Recv_buf) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Send_buf) {
		return false
	}

	if !ProtoWriteString(b, proto.Ip, 255) {
		return false
	}

	return true
}

func (proto *IPConfig) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Type) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Port) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Recv_buf) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Send_buf) {
		return false
	}

	if !ProtoReadString(b, &proto.Ip, 255) {
		return false
	}

	return true
}

func (proto *ActivityData) GetMid() uint16 {
	return 100
}

func (proto *ActivityData) GetPid() uint16 {
	return 85
}

func (proto *ActivityData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.State) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Today) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Custom, uint16(65535)) {
		return false
	}

	return true
}

func (proto *ActivityData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.State) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Today) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Custom, uint16(65535)) {
		return false
	}

	return true
}

func (proto *StringData) GetMid() uint16 {
	return 100
}

func (proto *StringData) GetPid() uint16 {
	return 86
}

func (proto *StringData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteIntegerArray(b, proto.Str, uint16(65535)) {
		return false
	}

	return true
}

func (proto *StringData) Read(b *bytes.Buffer) bool {
	if !ProtoReadIntegerArray(b, &proto.Str, uint16(65535)) {
		return false
	}

	return true
}

func (proto *GUIDData) GetMid() uint16 {
	return 100
}

func (proto *GUIDData) GetPid() uint16 {
	return 87
}

func (proto *GUIDData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	return true
}

func (proto *GUIDData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	return true
}

func (proto *AuctionObjectData) GetMid() uint16 {
	return 100
}

func (proto *AuctionObjectData) GetPid() uint16 {
	return 88
}

func (proto *AuctionObjectData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Type) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Custom1) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Custom2) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Custom3) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Level) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Fight_value) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Coin_type) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Coin_value) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Amount) {
		return false
	}

	if !ProtoWriteInteger(b, proto.State) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Puton_times) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Pulloff_times) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Seller_guid) {
		return false
	}

	if !ProtoWriteString(b, proto.Seller_name, 32) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Item_data) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Pet_data) {
		return false
	}

	return true
}

func (proto *AuctionObjectData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Type) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Custom1) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Custom2) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Custom3) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Level) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Fight_value) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Coin_type) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Coin_value) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Amount) {
		return false
	}

	if !ProtoReadInteger(b, &proto.State) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Puton_times) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pulloff_times) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Seller_guid) {
		return false
	}

	if !ProtoReadString(b, &proto.Seller_name, 32) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Item_data) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Pet_data) {
		return false
	}

	return true
}

func (proto *AuctionCookieData) GetMid() uint16 {
	return 100
}

func (proto *AuctionCookieData) GetPid() uint16 {
	return 89
}

func (proto *AuctionCookieData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Coins, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Favorites, uint8(255)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Records, uint8(255)) {
		return false
	}

	return true
}

func (proto *AuctionCookieData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Coins, uint16(65535)) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Favorites, uint8(255)) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Records, uint8(255)) {
		return false
	}

	return true
}

func (proto *GuildMemberBriefData) GetMid() uint16 {
	return 100
}

func (proto *GuildMemberBriefData) GetPid() uint16 {
	return 90
}

func (proto *GuildMemberBriefData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Player_guid) {
		return false
	}

	if !ProtoWriteString(b, proto.Player_name, 32) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Level) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Role) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guild_job) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Total_contrb) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Join_time) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Last_login_time) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Last_logout_time) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Status) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Vip) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Reincarnation) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Job) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Customs) {
		return false
	}

	return true
}

func (proto *GuildMemberBriefData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Player_guid) {
		return false
	}

	if !ProtoReadString(b, &proto.Player_name, 32) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Level) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Role) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guild_job) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Total_contrb) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Join_time) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Last_login_time) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Last_logout_time) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Status) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Vip) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Reincarnation) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Job) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Customs) {
		return false
	}

	return true
}

func (proto *GuildMemberData) GetMid() uint16 {
	return 100
}

func (proto *GuildMemberData) GetPid() uint16 {
	return 91
}

func (proto *GuildMemberData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Player_guid) {
		return false
	}

	if !ProtoWriteString(b, proto.Player_name, 32) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Level) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Role) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guild_job) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Total_contrb) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Surplus_contrb) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Last_week_contrb) {
		return false
	}

	if !ProtoWriteInteger(b, proto.This_week_contrb) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Permission) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Join_time) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Donate_bind_gold) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Last_donate_time) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Last_dividend_time) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Last_login_time) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Last_logout_time) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Forbid_end_time) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Status) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Vip) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Reincarnation) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Customs) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Job) {
		return false
	}

	return true
}

func (proto *GuildMemberData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Player_guid) {
		return false
	}

	if !ProtoReadString(b, &proto.Player_name, 32) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Level) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Role) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guild_job) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Total_contrb) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Surplus_contrb) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Last_week_contrb) {
		return false
	}

	if !ProtoReadInteger(b, &proto.This_week_contrb) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Permission) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Join_time) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Donate_bind_gold) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Last_donate_time) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Last_dividend_time) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Last_login_time) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Last_logout_time) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Forbid_end_time) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Status) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Vip) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Reincarnation) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Customs) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Job) {
		return false
	}

	return true
}

func (proto *GuildApplicantData) GetMid() uint16 {
	return 100
}

func (proto *GuildApplicantData) GetPid() uint16 {
	return 92
}

func (proto *GuildApplicantData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteString(b, proto.Name, 32) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Level) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Role) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Sn) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guild_contrb) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Last_apply_time) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Vip) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Reincarnation) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Job) {
		return false
	}

	return true
}

func (proto *GuildApplicantData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadString(b, &proto.Name, 32) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Level) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Role) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Sn) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guild_contrb) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Last_apply_time) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Vip) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Reincarnation) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Job) {
		return false
	}

	return true
}

func (proto *GuildEventData) GetMid() uint16 {
	return 100
}

func (proto *GuildEventData) GetPid() uint16 {
	return 93
}

func (proto *GuildEventData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Event_type) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Event_time) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Event_msg, uint16(65535)) {
		return false
	}

	return true
}

func (proto *GuildEventData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Event_type) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Event_time) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Event_msg, uint16(65535)) {
		return false
	}

	return true
}

func (proto *GuildBriefData) GetMid() uint16 {
	return 100
}

func (proto *GuildBriefData) GetPid() uint16 {
	return 94
}

func (proto *GuildBriefData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guild_id) {
		return false
	}

	if !ProtoWriteString(b, proto.Name, 32) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Leader_guid) {
		return false
	}

	if !ProtoWriteString(b, proto.Leader_name, 32) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Leader_level) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Leader_role) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Level) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Member_count) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Max_member_count) {
		return false
	}

	if !ProtoWriteString(b, proto.Declaration, 255) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Applyed) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Leader_school) {
		return false
	}

	return true
}

func (proto *GuildBriefData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guild_id) {
		return false
	}

	if !ProtoReadString(b, &proto.Name, 32) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Leader_guid) {
		return false
	}

	if !ProtoReadString(b, &proto.Leader_name, 32) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Leader_level) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Leader_role) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Level) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Member_count) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Max_member_count) {
		return false
	}

	if !ProtoReadString(b, &proto.Declaration, 255) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Applyed) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Leader_school) {
		return false
	}

	return true
}

func (proto *GuildData) GetMid() uint16 {
	return 100
}

func (proto *GuildData) GetPid() uint16 {
	return 95
}

func (proto *GuildData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guild_id) {
		return false
	}

	if !ProtoWriteString(b, proto.Name, 32) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Leader_guid) {
		return false
	}

	if !ProtoWriteString(b, proto.Leader_name, 32) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Level) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Base_level) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Wing_room_level) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Vault_level) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Academy_level) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Pharmacy_level) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Upgrade_cooldown) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Member_count) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Max_member_count) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Online_count) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Fund) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Liveness) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Build_degree) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Prestige) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Fight_score) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Config) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Create_time) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Dismissed_time) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Impeach_time) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Impeach_guid) {
		return false
	}

	if !ProtoWriteString(b, proto.Declaration, 255) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Events, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Customs) {
		return false
	}

	return true
}

func (proto *GuildData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guild_id) {
		return false
	}

	if !ProtoReadString(b, &proto.Name, 32) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Leader_guid) {
		return false
	}

	if !ProtoReadString(b, &proto.Leader_name, 32) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Level) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Base_level) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Wing_room_level) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Vault_level) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Academy_level) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Pharmacy_level) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Upgrade_cooldown) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Member_count) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Max_member_count) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Online_count) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Fund) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Liveness) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Build_degree) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Prestige) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Fight_score) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Config) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Create_time) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Dismissed_time) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Impeach_time) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Impeach_guid) {
		return false
	}

	if !ProtoReadString(b, &proto.Declaration, 255) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Events, uint16(65535)) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Customs) {
		return false
	}

	return true
}

func (proto *NoticeData) GetMid() uint16 {
	return 100
}

func (proto *NoticeData) GetPid() uint16 {
	return 96
}

func (proto *NoticeData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Notice_time) {
		return false
	}

	if !ProtoWriteString(b, proto.Notice, 1024) {
		return false
	}

	return true
}

func (proto *NoticeData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Notice_time) {
		return false
	}

	if !ProtoReadString(b, &proto.Notice, 1024) {
		return false
	}

	return true
}

func (proto *DBGuildData) GetMid() uint16 {
	return 100
}

func (proto *DBGuildData) GetPid() uint16 {
	return 97
}

func (proto *DBGuildData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteCustom(b, &proto.Guild) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Notice) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Items, uint8(255)) {
		return false
	}

	return true
}

func (proto *DBGuildData) Read(b *bytes.Buffer) bool {
	if !ProtoReadCustom(b, &proto.Guild) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Notice) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Items, uint8(255)) {
		return false
	}

	return true
}

func (proto *DBGuildMemberData) GetMid() uint16 {
	return 100
}

func (proto *DBGuildMemberData) GetPid() uint16 {
	return 98
}

func (proto *DBGuildMemberData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteCustom(b, &proto.Member) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guild_guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Total_fight_value) {
		return false
	}

	return true
}

func (proto *DBGuildMemberData) Read(b *bytes.Buffer) bool {
	if !ProtoReadCustom(b, &proto.Member) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guild_guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Total_fight_value) {
		return false
	}

	return true
}

func (proto *DBGuildApplicantData) GetMid() uint16 {
	return 100
}

func (proto *DBGuildApplicantData) GetPid() uint16 {
	return 99
}

func (proto *DBGuildApplicantData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteCustom(b, &proto.Applicant) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guild_guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Total_fight_value) {
		return false
	}

	return true
}

func (proto *DBGuildApplicantData) Read(b *bytes.Buffer) bool {
	if !ProtoReadCustom(b, &proto.Applicant) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guild_guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Total_fight_value) {
		return false
	}

	return true
}

func (proto *ContactData) GetMid() uint16 {
	return 100
}

func (proto *ContactData) GetPid() uint16 {
	return 100
}

func (proto *ContactData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Contact_type) {
		return false
	}

	if !ProtoWriteString(b, proto.Name, 32) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Role) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Job) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Level) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Friendship) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Last_contact_time) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Sn) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Status) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Vip) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Reincarnation) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Customs) {
		return false
	}

	return true
}

func (proto *ContactData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Contact_type) {
		return false
	}

	if !ProtoReadString(b, &proto.Name, 32) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Role) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Job) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Level) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Friendship) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Last_contact_time) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Sn) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Status) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Vip) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Reincarnation) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Customs) {
		return false
	}

	return true
}

func (proto *DBContactData) GetMid() uint16 {
	return 100
}

func (proto *DBContactData) GetPid() uint16 {
	return 101
}

func (proto *DBContactData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Player_guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Contact_guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Contact_type) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Friendship) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Last_contact_time) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Customs) {
		return false
	}

	return true
}

func (proto *DBContactData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Player_guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Contact_guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Contact_type) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Friendship) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Last_contact_time) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Customs) {
		return false
	}

	return true
}

func (proto *MessageData) GetMid() uint16 {
	return 100
}

func (proto *MessageData) GetPid() uint16 {
	return 102
}

func (proto *MessageData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteIntegerArray(b, proto.Message, uint16(65535)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Send_guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Send_time) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Recv_guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Auto_reply) {
		return false
	}

	return true
}

func (proto *MessageData) Read(b *bytes.Buffer) bool {
	if !ProtoReadIntegerArray(b, &proto.Message, uint16(65535)) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Send_guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Send_time) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Recv_guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Auto_reply) {
		return false
	}

	return true
}

func (proto *ContactInfoData) GetMid() uint16 {
	return 100
}

func (proto *ContactInfoData) GetPid() uint16 {
	return 103
}

func (proto *ContactInfoData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteString(b, proto.Name, 32) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Role) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Job) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Level) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Sn) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Status) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Vip) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Reincarnation) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Last_login_time) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Last_logout_time) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Total_fight_value) {
		return false
	}

	return true
}

func (proto *ContactInfoData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadString(b, &proto.Name, 32) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Role) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Job) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Level) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Sn) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Status) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Vip) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Reincarnation) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Last_login_time) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Last_logout_time) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Total_fight_value) {
		return false
	}

	return true
}

func (proto *MailHeadData) GetMid() uint16 {
	return 100
}

func (proto *MailHeadData) GetPid() uint16 {
	return 104
}

func (proto *MailHeadData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Mail_guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Mail_type) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Title, uint8(255)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Send_guid) {
		return false
	}

	if !ProtoWriteString(b, proto.Send_name, 32) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Send_time) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Status) {
		return false
	}

	return true
}

func (proto *MailHeadData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Mail_guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Mail_type) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Title, uint8(255)) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Send_guid) {
		return false
	}

	if !ProtoReadString(b, &proto.Send_name, 32) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Send_time) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Status) {
		return false
	}

	return true
}

func (proto *MailBodyData) GetMid() uint16 {
	return 100
}

func (proto *MailBodyData) GetPid() uint16 {
	return 105
}

func (proto *MailBodyData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteIntegerArray(b, proto.Content, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Attrs, uint16(65535)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Items, uint8(255)) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Pets, uint8(255)) {
		return false
	}

	return true
}

func (proto *MailBodyData) Read(b *bytes.Buffer) bool {
	if !ProtoReadIntegerArray(b, &proto.Content, uint16(65535)) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Attrs, uint16(65535)) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Items, uint8(255)) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Pets, uint8(255)) {
		return false
	}

	return true
}

func (proto *MailData) GetMid() uint16 {
	return 100
}

func (proto *MailData) GetPid() uint16 {
	return 106
}

func (proto *MailData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteCustom(b, &proto.Head) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Body) {
		return false
	}

	return true
}

func (proto *MailData) Read(b *bytes.Buffer) bool {
	if !ProtoReadCustom(b, &proto.Head) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Body) {
		return false
	}

	return true
}

func (proto *DBMailData) GetMid() uint16 {
	return 100
}

func (proto *DBMailData) GetPid() uint16 {
	return 107
}

func (proto *DBMailData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Recv_guid) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Guids, uint32(4294967295)) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Head) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Body) {
		return false
	}

	return true
}

func (proto *DBMailData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Recv_guid) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Guids, uint32(4294967295)) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Head) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Body) {
		return false
	}

	return true
}

func (proto *DBVarData) GetMid() uint16 {
	return 100
}

func (proto *DBVarData) GetPid() uint16 {
	return 108
}

func (proto *DBVarData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteIntegerArray(b, proto.Key, uint8(255)) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Value, uint16(65535)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Value_type) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Merge_type) {
		return false
	}

	return true
}

func (proto *DBVarData) Read(b *bytes.Buffer) bool {
	if !ProtoReadIntegerArray(b, &proto.Key, uint8(255)) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Value, uint16(65535)) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Value_type) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Merge_type) {
		return false
	}

	return true
}

func (proto *RanklistData) GetMid() uint16 {
	return 100
}

func (proto *RanklistData) GetPid() uint16 {
	return 109
}

func (proto *RanklistData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteString(b, proto.Name, 32) {
		return false
	}

	if !ProtoWriteString(b, proto.Rank_info, 255) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Rank_data1) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Rank_data2) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Rank_data3) {
		return false
	}

	return true
}

func (proto *RanklistData) Read(b *bytes.Buffer) bool {
	if !ProtoReadString(b, &proto.Name, 32) {
		return false
	}

	if !ProtoReadString(b, &proto.Rank_info, 255) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Rank_data1) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Rank_data2) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Rank_data3) {
		return false
	}

	return true
}

func (proto *DBRanklistData) GetMid() uint16 {
	return 100
}

func (proto *DBRanklistData) GetPid() uint16 {
	return 110
}

func (proto *DBRanklistData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteString(b, proto.Name, 32) {
		return false
	}

	if !ProtoWriteString(b, proto.Rank_info, 255) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Rank_data1) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Rank_data2) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Rank_data3) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Rank_guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Object_guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Rank_type) {
		return false
	}

	return true
}

func (proto *DBRanklistData) Read(b *bytes.Buffer) bool {
	if !ProtoReadString(b, &proto.Name, 32) {
		return false
	}

	if !ProtoReadString(b, &proto.Rank_info, 255) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Rank_data1) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Rank_data2) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Rank_data3) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Rank_guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Object_guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Rank_type) {
		return false
	}

	return true
}

func (proto *TeamBriefData) GetMid() uint16 {
	return 100
}

func (proto *TeamBriefData) GetPid() uint16 {
	return 111
}

func (proto *TeamBriefData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Team_guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Target) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Min_require_level) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Max_require_level) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Member_count) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Leader_data) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Has_applicant) {
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

func (proto *TeamBriefData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Team_guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Target) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Min_require_level) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Max_require_level) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Member_count) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Leader_data) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Has_applicant) {
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

func (proto *TeamApplicantData) GetMid() uint16 {
	return 100
}

func (proto *TeamApplicantData) GetPid() uint16 {
	return 112
}

func (proto *TeamApplicantData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteString(b, proto.Name, 32) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Attrs, uint16(65535)) {
		return false
	}

	return true
}

func (proto *TeamApplicantData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadString(b, &proto.Name, 32) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Attrs, uint16(65535)) {
		return false
	}

	return true
}

func (proto *ExchangeData) GetMid() uint16 {
	return 100
}

func (proto *ExchangeData) GetPid() uint16 {
	return 113
}

func (proto *ExchangeData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Day) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Records, uint16(65535)) {
		return false
	}

	return true
}

func (proto *ExchangeData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Day) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Records, uint16(65535)) {
		return false
	}

	return true
}

func (proto *ForbidData) GetMid() uint16 {
	return 100
}

func (proto *ForbidData) GetPid() uint16 {
	return 114
}

func (proto *ForbidData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Sid) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Keyword, uint8(255)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Keytype) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Fbdtype) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Begin) {
		return false
	}

	if !ProtoWriteInteger(b, proto.End) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Desc, uint8(255)) {
		return false
	}

	return true
}

func (proto *ForbidData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Sid) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Keyword, uint8(255)) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Keytype) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Fbdtype) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Begin) {
		return false
	}

	if !ProtoReadInteger(b, &proto.End) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Desc, uint8(255)) {
		return false
	}

	return true
}

func (proto *ForbidTalkData) GetMid() uint16 {
	return 100
}

func (proto *ForbidTalkData) GetPid() uint16 {
	return 115
}

func (proto *ForbidTalkData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Begin) {
		return false
	}

	if !ProtoWriteInteger(b, proto.End) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Desc, uint8(255)) {
		return false
	}

	return true
}

func (proto *ForbidTalkData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Begin) {
		return false
	}

	if !ProtoReadInteger(b, &proto.End) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Desc, uint8(255)) {
		return false
	}

	return true
}

func (proto *PasturePetData) GetMid() uint16 {
	return 100
}

func (proto *PasturePetData) GetPid() uint16 {
	return 116
}

func (proto *PasturePetData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Owner_guid) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Pet) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Map) {
		return false
	}

	if !ProtoWriteInteger(b, proto.X) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Y) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Placed_time) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Placed_status) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Replaced_time) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Replaced_player_guid) {
		return false
	}

	if !ProtoWriteString(b, proto.Replaced_player_name, 32) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Replaced_pet_guid) {
		return false
	}

	if !ProtoWriteString(b, proto.Replaced_pet_name, 32) {
		return false
	}

	return true
}

func (proto *PasturePetData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Owner_guid) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Pet) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Map) {
		return false
	}

	if !ProtoReadInteger(b, &proto.X) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Y) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Placed_time) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Placed_status) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Replaced_time) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Replaced_player_guid) {
		return false
	}

	if !ProtoReadString(b, &proto.Replaced_player_name, 32) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Replaced_pet_guid) {
		return false
	}

	if !ProtoReadString(b, &proto.Replaced_pet_name, 32) {
		return false
	}

	return true
}

func (proto *LadderData) GetMid() uint16 {
	return 100
}

func (proto *LadderData) GetPid() uint16 {
	return 117
}

func (proto *LadderData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Player_guid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Rank) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Player) {
		return false
	}

	return true
}

func (proto *LadderData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Player_guid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Rank) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Player) {
		return false
	}

	return true
}

func (proto *LadderRecordData) GetMid() uint16 {
	return 100
}

func (proto *LadderRecordData) GetPid() uint16 {
	return 118
}

func (proto *LadderRecordData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Action) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Time) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Player_guid) {
		return false
	}

	if !ProtoWriteString(b, proto.Player_name, 32) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Player_rank) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Target_guid) {
		return false
	}

	if !ProtoWriteString(b, proto.Target_name, 32) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Target_rank) {
		return false
	}

	return true
}

func (proto *LadderRecordData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Action) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Time) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Player_guid) {
		return false
	}

	if !ProtoReadString(b, &proto.Player_name, 32) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Player_rank) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Target_guid) {
		return false
	}

	if !ProtoReadString(b, &proto.Target_name, 32) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Target_rank) {
		return false
	}

	return true
}

func (proto *PlayerDetail) GetMid() uint16 {
	return 100
}

func (proto *PlayerDetail) GetPid() uint16 {
	return 119
}

func (proto *PlayerDetail) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Uid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Sid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Guid) {
		return false
	}

	if !ProtoWriteString(b, proto.Unid, 32) {
		return false
	}

	if !ProtoWriteString(b, proto.Name, 32) {
		return false
	}

	if !ProtoWriteCustomArray(b, proto.Attrs, uint16(65535)) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Map) {
		return false
	}

	if !ProtoWriteInteger(b, proto.X) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Y) {
		return false
	}

	if !ProtoWriteInteger(b, proto.State) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Createtime) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Lastlogin) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Destroytime) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Customs) {
		return false
	}

	if !ProtoWriteCustom(b, &proto.Equips) {
		return false
	}

	return true
}

func (proto *PlayerDetail) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Uid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Sid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Guid) {
		return false
	}

	if !ProtoReadString(b, &proto.Unid, 32) {
		return false
	}

	if !ProtoReadString(b, &proto.Name, 32) {
		return false
	}

	if !ProtoReadCustomArray(b, &proto.Attrs, uint16(65535)) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Map) {
		return false
	}

	if !ProtoReadInteger(b, &proto.X) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Y) {
		return false
	}

	if !ProtoReadInteger(b, &proto.State) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Createtime) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Lastlogin) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Destroytime) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Customs) {
		return false
	}

	if !ProtoReadCustom(b, &proto.Equips) {
		return false
	}

	return true
}

func (proto *Pt) GetMid() uint16 {
	return 100
}

func (proto *Pt) GetPid() uint16 {
	return 120
}

func (proto *Pt) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.X) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Y) {
		return false
	}

	return true
}

func (proto *Pt) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.X) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Y) {
		return false
	}

	return true
}

func (proto *ViolateData) GetMid() uint16 {
	return 100
}

func (proto *ViolateData) GetPid() uint16 {
	return 121
}

func (proto *ViolateData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Type) {
		return false
	}

	if !ProtoWriteIntegerArray(b, proto.Violate, uint8(255)) {
		return false
	}

	return true
}

func (proto *ViolateData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Type) {
		return false
	}

	if !ProtoReadIntegerArray(b, &proto.Violate, uint8(255)) {
		return false
	}

	return true
}

func (proto *ChannelSwitchData) GetMid() uint16 {
	return 100
}

func (proto *ChannelSwitchData) GetPid() uint16 {
	return 122
}

func (proto *ChannelSwitchData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Type) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Sw) {
		return false
	}

	return true
}

func (proto *ChannelSwitchData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Type) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Sw) {
		return false
	}

	return true
}

func (proto *VerifyData) GetMid() uint16 {
	return 100
}

func (proto *VerifyData) GetPid() uint16 {
	return 123
}

func (proto *VerifyData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Fbd_time) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Accountstatus) {
		return false
	}

	return true
}

func (proto *VerifyData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Fbd_time) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Accountstatus) {
		return false
	}

	return true
}

func (proto *IPPort) GetMid() uint16 {
	return 100
}

func (proto *IPPort) GetPid() uint16 {
	return 124
}

func (proto *IPPort) Write(b *bytes.Buffer) bool {
	if !ProtoWriteString(b, proto.Ip, 32) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Port) {
		return false
	}

	return true
}

func (proto *IPPort) Read(b *bytes.Buffer) bool {
	if !ProtoReadString(b, &proto.Ip, 32) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Port) {
		return false
	}

	return true
}

func (proto *CrossIPPort) GetMid() uint16 {
	return 100
}

func (proto *CrossIPPort) GetPid() uint16 {
	return 125
}

func (proto *CrossIPPort) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Gid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Gcp) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Gsp) {
		return false
	}

	if !ProtoWriteString(b, proto.Innerip, 32) {
		return false
	}

	if !ProtoWriteString(b, proto.Extip, 32) {
		return false
	}

	if !ProtoWriteString(b, proto.Innerflag, 64) {
		return false
	}

	if !ProtoWriteString(b, proto.Clientname, 64) {
		return false
	}

	return true
}

func (proto *CrossIPPort) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Gid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Gcp) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Gsp) {
		return false
	}

	if !ProtoReadString(b, &proto.Innerip, 32) {
		return false
	}

	if !ProtoReadString(b, &proto.Extip, 32) {
		return false
	}

	if !ProtoReadString(b, &proto.Innerflag, 64) {
		return false
	}

	if !ProtoReadString(b, &proto.Clientname, 64) {
		return false
	}

	return true
}

func (proto *TestJsonName) GetMid() uint16 {
	return 100
}

func (proto *TestJsonName) GetPid() uint16 {
	return 126
}

func (proto *TestJsonName) Write(b *bytes.Buffer) bool {
	if !ProtoWriteIntegerArray(b, proto.Name, uint16(65535)) {
		return false
	}

	return true
}

func (proto *TestJsonName) Read(b *bytes.Buffer) bool {
	if !ProtoReadIntegerArray(b, &proto.Name, uint16(65535)) {
		return false
	}

	return true
}

func (proto *BillQueryData) GetMid() uint16 {
	return 100
}

func (proto *BillQueryData) GetPid() uint16 {
	return 127
}

func (proto *BillQueryData) Write(b *bytes.Buffer) bool {
	if !ProtoWriteInteger(b, proto.Bill_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.User_id) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Sid) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Added_yb) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Award_yb) {
		return false
	}

	if !ProtoWriteString(b, proto.User, 20) {
		return false
	}

	if !ProtoWriteString(b, proto.Desc, 255) {
		return false
	}

	if !ProtoWriteInteger(b, proto.Time) {
		return false
	}

	return true
}

func (proto *BillQueryData) Read(b *bytes.Buffer) bool {
	if !ProtoReadInteger(b, &proto.Bill_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.User_id) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Sid) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Added_yb) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Award_yb) {
		return false
	}

	if !ProtoReadString(b, &proto.User, 20) {
		return false
	}

	if !ProtoReadString(b, &proto.Desc, 255) {
		return false
	}

	if !ProtoReadInteger(b, &proto.Time) {
		return false
	}

	return true
}

