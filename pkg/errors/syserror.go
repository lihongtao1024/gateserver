///>本代码由测试工具自动生成,请勿手动修改
package errors

import "gateserver/pkg"

var errorArray [pkg.ErrorMax]string

func init() {
	errorArray[pkg.ErrorUnknown] = "系统错误"
	errorArray[pkg.ErrorLoginIllegal] = "非法的登陆请求"
	errorArray[pkg.ErrorGTOverload] = "网关连接负载过高"
	errorArray[pkg.ErrorWSOverload] = "服务器人数已满"
	errorArray[pkg.ErrorGSOverload] = "游戏人数已满"
	errorArray[pkg.ErrorGSNotFound] = "找不到游戏服务器"
	errorArray[pkg.ErrorLoginRepeat] = "不要重复提交登陆请求"
	errorArray[pkg.ErrorLoginPwd] = "登陆账号或密码不能为空"
	errorArray[pkg.ErrorLoginInfo] = "请完善登陆信息"
	errorArray[pkg.ErrorLoginPend] = "登陆请求正在处理中请稍候"
	errorArray[pkg.ErrorLoginAgain] = "账号从另一处登陆"
	errorArray[pkg.ErrorTimeoutKick] = "操作超时"
	errorArray[pkg.ErrorOnlineRelogin] = "正常的在线账号不允许重连"
	errorArray[pkg.ErrorUserRelogin] = "找不到重连的账号数据"
	errorArray[pkg.ErrorKickExitPend] = "请离或退出请求正在处理中请稍候"
	errorArray[pkg.ErrorUserNull] = "数据库中无此账号数据"
	errorArray[pkg.ErrorLogoutIllegal] = "非法的退出请求"
	errorArray[pkg.ErrorUserState] = "该请求无法在当前玩家状态下执行"
	errorArray[pkg.ErrorPlayerList] = "请先取得玩家列表"
	errorArray[pkg.ErrorPlayerPending] = "网络错误20"
	errorArray[pkg.ErrorPlayerIsGameing] = "该帐号下已有角色正在游戏中"
	errorArray[pkg.ErrorPlayerName] = "玩家名已存在"
	errorArray[pkg.ErrorPlayerNull] = "找不到指定的玩家"
	errorArray[pkg.ErrorPlayerDeleted] = "玩家已被删除"
	errorArray[pkg.ErrorPlayerNormal] = "玩家未被删除"
	errorArray[pkg.ErrorPlayerHasDeleted] = "超出当日删除上限"
	errorArray[pkg.ErrorPlayerNormalLimit] = "达到玩家列表上限"
	errorArray[pkg.ErrorPlayerDeleteLimit] = "达到删除列表上限"
	errorArray[pkg.ErrorPlayerRestoreLimit] = "超过恢复时效"
	errorArray[pkg.ErrorCacheTimeout] = "缓存加载超时"
	errorArray[pkg.ErrorPlayerId] = "玩家ID非法"
	errorArray[pkg.ErrorPlayerRace] = "玩家种族非法"
	errorArray[pkg.ErrorPlayerGender] = "玩家性别非法"
	errorArray[pkg.ErrorPlayerNameViolate] = "玩家名字包含非法字符"
	errorArray[pkg.ErrorPlayerSync] = "网络错误35"
	errorArray[pkg.ErrorLoginWrongUser] = "无此用户"
	errorArray[pkg.ErrorLoginWrongPassword] = "密码输入有误"
	errorArray[pkg.ErrorForbidLoginAccount] = "帐号禁止登陆"
	errorArray[pkg.ErrorAbnormalUserData] = "账号数据异常"
	errorArray[pkg.ErrorKeepAliveTimeout] = "账号心跳检测超时"
	errorArray[pkg.ErrorKicked] = "您已被请离服务器"
	errorArray[pkg.ErrorForbidLoginPlayer] = "玩家禁止登陆"
	errorArray[pkg.ErrorClientType] = "客户端版本错误"
	errorArray[pkg.ErrorReloginPend] = "重连请求正在处理中请稍候"
	errorArray[pkg.ErrorPlayerListPend] = "角色列表正在加载中请稍候"
	errorArray[pkg.ErrorPlayerCreatePend] = "角色正在创建中请稍候"
	errorArray[pkg.ErrorPlayerDestroyPend] = "角色正在销毁中请稍候"
	errorArray[pkg.ErrorPlayerRestorePend] = "角色正在恢复中请稍候"
	errorArray[pkg.ErrorEnterGSPend] = "正在进入游戏中请稍候"

	errorArray[pkg.ErrorNoMovablePT] = "无可移动点"
	errorArray[pkg.ErrorNoMap] = "无地图信息"
	errorArray[pkg.ErrorPath] = "路径非法有不可移动点"
	errorArray[pkg.ErrorMoveFast] = "移动速度过快"
	errorArray[pkg.ErrorNoJumpIfo] = "无跳转点信息"
	errorArray[pkg.ErrorInvalidJump] = "非法的跳转请求"
	errorArray[pkg.ErrorRepeatPt] = "移动目标点重复"
	errorArray[pkg.ErrorRepeatMap] = "出现地图重复"
	errorArray[pkg.ErrorItemInContainer] = "道具在包裹中"
	errorArray[pkg.ErrorNoPlace] = "无道具可放置点"
	errorArray[pkg.ErrorPathNoPT] = "坐标点不在路径上"
	errorArray[pkg.ErrorInFight] = "战斗中无法操作"
	errorArray[pkg.ErrorInView] = "观战中无法操作"
	errorArray[pkg.ErrorInTeam] = "队伍中无法操作"
	errorArray[pkg.ErrorInFollow] = "跟随中无法操作"
	errorArray[pkg.ErrorInvalidPt] = "无效的坐标"
	errorArray[pkg.ErrorScriptJump] = "脚本禁止跳转"
	errorArray[pkg.ErrorInConvoy] = "押镖中禁止跳转"
	errorArray[pkg.ErrorForbidMove] = "禁止移动"
	errorArray[pkg.ErrorForbidJump] = "禁止跳转"

	errorArray[pkg.ErrorItemSwapSite] = "物品交换位置失败"
	errorArray[pkg.ErrorItemNotEnough] = "物品不足"
	errorArray[pkg.ErrorItemNotInContainer] = "物品不在包裹内"
	errorArray[pkg.ErrorItemMoveIn] = "移入物品失败"
	errorArray[pkg.ErrorItemMoveOut] = "移出物品失败"
	errorArray[pkg.ErrorItemSite] = "物品位置异常"
	errorArray[pkg.ErrorItemRearrangeSite] = "整理物品失败"
	errorArray[pkg.ErrorItemHasMap] = "物品在地图上"
	errorArray[pkg.ErrorItemCanMoveIn] = "物品不能移入仓库"
	errorArray[pkg.ErrorItemCanUse] = "该物品无法直接使用"
	errorArray[pkg.ErrorItemCanEquip] = "无法装备"
	errorArray[pkg.ErrorItemNotExist] = "物品不存在"
	errorArray[pkg.ErrorItemContainerCapacity] = "包裹空间不足"
	errorArray[pkg.ErrorItemNotOwner] = "玩家不拥有该装备"
	errorArray[pkg.ErrorItemNotEquip] = "物品不是装备"
	errorArray[pkg.ErrorItemContainerFull] = "包裹已满"
	errorArray[pkg.ErrorItemNotUseIsFight] = "战斗中不能对目标使用该道具"
	errorArray[pkg.ErrorItemApplyTarget] = "使用战斗道具时目标非法"
	errorArray[pkg.ErrorItemBound] = "物品绑定的"
	errorArray[pkg.ErrorItemUserNotMatched] = "不能对目标使用物品"
	errorArray[pkg.ErrorItemAttrRequire] = "不满足穿戴属性要求"
	errorArray[pkg.ErrorItemCanUseAll] = "该物品不能批量使用"
	errorArray[pkg.ErrorItemNoNeed] = "该物品无需使用"
	errorArray[pkg.ErrorItemCountOverflow] = "物品数量超出上限，无法购买"
	errorArray[pkg.ErrorBindGoldNotEnough] = "银币不足，无法购买"
	errorArray[pkg.ErrorGoldNotEnough] = "金币不足，无法购买"
	errorArray[pkg.ErrorBindIngotNotEnough] = "银元宝不足，无法购买"
	errorArray[pkg.ErrorIngotNotEnough] = "元宝不足，无法购买"
	errorArray[pkg.ErrorItemCantEquip] = "无法脱下装备"

	errorArray[pkg.ErrorBuffTriggerTime] = "BUFF未满足触发时机"
	errorArray[pkg.ErrorBuffCooldown] = "BUFF冷却中"
	errorArray[pkg.ErrorBuffTriggerTotal] = "超出BUFF最大触发次数"
	errorArray[pkg.ErrorBuffTriggerPerRound] = "超出BUFF每回合最大触发次数"
	errorArray[pkg.ErrorBuffDelay] = "BUFF需要延时"
	errorArray[pkg.ErrorBuffRateUnreaches] = "BUFF触发概率未满足"
	errorArray[pkg.ErrorBuffTriggerSearch] = "BUFF查找触发目标失败"
	errorArray[pkg.ErrorBuffTriggerCondition] = "BUFF触发条件未满足"
	errorArray[pkg.ErrorBuffActionSearch] = "BUFF查找行为目标失败"
	errorArray[pkg.ErrorBuffActionTarget] = "BUFF无行为目标"
	errorArray[pkg.ErrorBuffActionExecute] = "BUFF执行行为失败"
	errorArray[pkg.ErrorBuffActorDied] = "BUFF不能在死亡时触发"

	errorArray[pkg.ErrorSkillActor] = "当前状态无法施放该技能"
	errorArray[pkg.ErrorSkillDelay] = "技能需要延时"
	errorArray[pkg.ErrorSkillCooldown] = "技能冷却中"
	errorArray[pkg.ErrorSkillMajorTarget] = "技能主目标非法"
	errorArray[pkg.ErrorSkillTargetType] = "未知的技能目标类型"
	errorArray[pkg.ErrorSkillShotCondition] = "技能施放条件不满足"
	errorArray[pkg.ErrorSkillShotCost] = "技能消耗不足"
	errorArray[pkg.ErrorSkillSearchTarget] = "查找技能目标失败"
	errorArray[pkg.ErrorSkillExtraTargetType] = "未知的技能额外目标类型"
	errorArray[pkg.ErrorSkillDamageType] = "未知的技能伤害类型"
	errorArray[pkg.ErrorSkillShotDamage] = "技能造成伤害失败"
	errorArray[pkg.ErrorSkillDebuffType] = "未知的DEBUFF命中类型"
	errorArray[pkg.ErrorSkillShotDebuff] = "技能添加DEBUFF失败"
	errorArray[pkg.ErrorSkillShotBuff] = "技能添加BUFF失败"
	errorArray[pkg.ErrorSkillActorControl] = "技能施放者被控制"
	errorArray[pkg.ErrorSkillTargetEmpty] = "技能没有施放目标"
	errorArray[pkg.ErrorSkillNotFound] = "找不到指定的技能"

	errorArray[pkg.ErrorAITriggerTime] = "AI触发类型不符"
	errorArray[pkg.ErrorAIExpire] = "AI已过期"
	errorArray[pkg.ErrorAICooldown] = "AI冷却中"
	errorArray[pkg.ErrorAITriggerTotal] = "超出AI最大触发次数"
	errorArray[pkg.ErrorAITriggerPerRound] = "超出AI每回合最大触发次数"
	errorArray[pkg.ErrorAIDelay] = "AI需要延时"
	errorArray[pkg.ErrorAIRateUnreaches] = "AI触发概率未满足"
	errorArray[pkg.ErrorAITriggerSearch] = "AI查找触发目标失败"
	errorArray[pkg.ErrorAITriggerCondition] = "AI触发条件未满足"
	errorArray[pkg.ErrorAIActionSearch] = "AI查找行为目标失败"
	errorArray[pkg.ErrorAIActionTarget] = "AI无行为目标"
	errorArray[pkg.ErrorAIActionExecute] = "AI执行行为失败"
	errorArray[pkg.ErrorAIActorDied] = "AI不能在死亡时触发"

	errorArray[pkg.ErrorPlayerIsFight] = "玩家正在战斗中"
	errorArray[pkg.ErrorPlayerNotFight] = "玩家不在战斗中"
	errorArray[pkg.ErrorPlayerIsView] = "玩家正在观战中"
	errorArray[pkg.ErrorPlayerNotView] = "玩家不在观战中"
	errorArray[pkg.ErrorFighterActorNull] = "找不到战斗动作发起者"
	errorArray[pkg.ErrorFighterTargetNull] = "找不到战斗动作目标"
	errorArray[pkg.ErrorFightSceneState] = "当前战斗状态不允许"
	errorArray[pkg.ErrorFighterRepeatOperate] = "请不要重复发送战斗操作"
	errorArray[pkg.ErrorFighterChaos] = "混乱中行动失败"
	errorArray[pkg.ErrorFighterIce] = "封印中行动失败"
	errorArray[pkg.ErrorFighterSleep] = "昏睡中行动失败"
	errorArray[pkg.ErrorFighterHuaWu] = "技能被化无释放失败"
	errorArray[pkg.ErrorFighterDied] = "已死亡行动失败"
	errorArray[pkg.ErrorFighterCatchTarget] = "该目标无法被捕捉"
	errorArray[pkg.ErrorFighterCatchLevel] = "捕捉等级不足"
	errorArray[pkg.ErrorFighterCatchAdd] = "宠物栏已满"
	errorArray[pkg.ErrorFighterCmd] = "非法的战斗指令"
	errorArray[pkg.ErrorFighterEnterView] = "进入观战失败"
	errorArray[pkg.ErrorFighterLeaveView] = "退出观战失败"
	errorArray[pkg.ErrorFighterIdle] = "待机中行动失败"
	errorArray[pkg.ErrorFighterGeneralAtk] = "强制物理攻击"
	errorArray[pkg.ErrorFighterIgnoreProtect] = "忽略目标保护"
	errorArray[pkg.ErrorFighterForget] = "技能被遗忘执行物理普攻"
	errorArray[pkg.ErrorFighterDefense] = "强制防御"
	errorArray[pkg.ErrorOperateForbid] = "该操作被禁止行动失败"
	errorArray[pkg.ErrorFighterIgnoreDefense] = "忽略目标防御"
	errorArray[pkg.ErrorPetLoyaltyForceAttack] = "因忠诚度不足无法施放技能"
	errorArray[pkg.ErrorPetLoyaltyForceEscape] = "因忠诚度不足逃跑了"
	errorArray[pkg.ErrorPetSummonOut] = "已出战或参战过的宠物不可再出战"
	errorArray[pkg.ErrorPetCatchFail] = "捕捉宠物失败"
	errorArray[pkg.ErrorEscapeFail] = "逃跑失败"

	errorArray[pkg.ErrorTeamExist] = "玩家已经有队伍"
	errorArray[pkg.ErrorTeamAddMember] = "队伍添加成员失败"
	errorArray[pkg.ErrorTeamDelMember] = "队伍移除成员失败"
	errorArray[pkg.ErrorTeamNotMatched] = "此队伍已停止招募"
	errorArray[pkg.ErrorTeamIsLeader] = "玩家是队长"
	errorArray[pkg.ErrorTeamNotIsLeader] = "玩家不是队长"
	errorArray[pkg.ErrorTeamIsTempLeave] = "玩家是暂离的"
	errorArray[pkg.ErrorTeamNotSameGuild] = "当前队长所在的位置无法归队"
	errorArray[pkg.ErrorTeamNotExist] = "玩家没有队伍"
	errorArray[pkg.ErrorTeamApplicantExist] = "申请已存在，无法操作"
	errorArray[pkg.ErrorTeamApplicantNotExist] = "申请不存在，无法操作"
	errorArray[pkg.ErrorTeamCantApply] = "队长已取消招募"
	errorArray[pkg.ErrorTeamConvoy] = "押镖中不能组队"

	errorArray[pkg.ErrorQuestAcceptCondition] = "玩家不满足任务接取条件"
	errorArray[pkg.ErrorQuestNotFound] = "任务不存在"
	errorArray[pkg.ErrorQuestAdd] = "添加任务失败"
	errorArray[pkg.ErrorQuestCanAdd] = "不能添加任务"
	errorArray[pkg.ErrorQuestCanAccept] = "不能接取任务"
	errorArray[pkg.ErrorQuestCanCommit] = "不能完成任务"
	errorArray[pkg.ErrorQuestCanAbandon] = "不能放弃任务"
	errorArray[pkg.ErrorQuestCanFailed] = "不能失败任务"
	errorArray[pkg.ErrorQuestHasAccepted] = "任务已经接取"
	errorArray[pkg.ErrorQuestHasAbandoned] = "任务已经放弃"
	errorArray[pkg.ErrorQuestHasFailed] = "任务已经失败"
	errorArray[pkg.ErrorQuestHasFinished] = "任务已经完成"
	errorArray[pkg.ErrorQuestHasNotActivated] = "任务未激活"
	errorArray[pkg.ErrorQuestHasMutex] = "已经接取了互斥任务"
	errorArray[pkg.ErrorQuestHasLine] = "已经接取了线性任务的前置任务"
	errorArray[pkg.ErrorQuestOperationCondition] = "玩家不满足任务操作条件"
	errorArray[pkg.ErrorQuestMustAlone] = "任务必须单人"
	errorArray[pkg.ErrorQuestMustTeam] = "任务必须组队"
	errorArray[pkg.ErrorQuestTeamAverageLevel] = "队伍平均等级不满足任务要求"
	errorArray[pkg.ErrorQuestTeamNumLower] = "队伍成员数量不足"
	errorArray[pkg.ErrorQuestTeamNumUpper] = "队伍成员数量太多"
	errorArray[pkg.ErrorQuestLogout] = "任务已经注销"
	errorArray[pkg.ErrorQuestCanFight] = "任务条件不满足"
	errorArray[pkg.ErrorQuestCanReady] = "不能准备完成任务"
	errorArray[pkg.ErrorQuestTeamOrgGuild] = "队伍成员不是同一帮派"
	errorArray[pkg.ErrorQuestTeamOrgJob] = "队伍成员不是同一门派"
	errorArray[pkg.ErrorQuestMustAllAccept] = "队伍所有成员必须满足接取条件"
	errorArray[pkg.ErrorQuestCantTeam] = "组队中不能接押镖任务"

	errorArray[pkg.ErrorPetNotInContainer] = "宠物不在包裹内"
	errorArray[pkg.ErrorPetInContainer] = "宠物已经在包裹内"
	errorArray[pkg.ErrorPetShow] = "宠物展示失败"
	errorArray[pkg.ErrorPetHide] = "宠物隐藏失败"
	errorArray[pkg.ErrorPetSetLineup] = "宠物上阵失败"
	errorArray[pkg.ErrorPetClrLineup] = "宠物下阵失败"
	errorArray[pkg.ErrorPetInWarehouse] = "宠物在仓库内"
	errorArray[pkg.ErrorPetMoveIn] = "移入宠物失败"
	errorArray[pkg.ErrorPetMoveOut] = "移出宠物失败"
	errorArray[pkg.ErrorPetIsShowed] = "宠物已经展示"
	errorArray[pkg.ErrorPetIsLineup] = "宠物已经上阵"
	errorArray[pkg.ErrorPetCanMoveIn] = "宠物仓库已满，不能移入"
	errorArray[pkg.ErrorPetShowed] = "展示宠物不能放入仓库"
	errorArray[pkg.ErrorPetBind] = "绑定宠物不能放入仓库"
	errorArray[pkg.ErrorPetLocked] = "锁定宠物不能放入仓库"
	errorArray[pkg.ErrorPetLineup] = "上阵宠物不能放入仓库"
	errorArray[pkg.ErrorPetNotFound] = "宠物不存在"
	errorArray[pkg.ErrorPetNotLineup] = "没有上阵宠物"
	errorArray[pkg.ErrorPetNotShowed] = "没有展示宠物"
	errorArray[pkg.ErrorPetNotOwner] = "宠物没有主人"
	errorArray[pkg.ErrorGuardNotFound] = "侍从不存在"
	errorArray[pkg.ErrorGuardContainerFull] = "上阵侍从已满"
	errorArray[pkg.ErrorPetContainerFull] = "宠物包裹已满"
	errorArray[pkg.ErrorPetInPasture] = "宠物放牧中"
	errorArray[pkg.ErrorPetNotYour] = "该宠物不属于您"
	errorArray[pkg.ErrorPetReplace] = "宠物替换失败"
	errorArray[pkg.ErrorPetLevel] = "宠物与角色等级差距过大无法出战"

	errorArray[pkg.ErrorShopSellAdd] = "出售物品添加回购列表失败"
	errorArray[pkg.ErrorShopCantSell] = "物品不可出售"
	errorArray[pkg.ErrorShopBindGoldlack] = "绑定金币不足"
	errorArray[pkg.ErrorShopSplitItem] = "出售分离物品失败"
	errorArray[pkg.ErrorShopOpen] = "商店打开失败"
	errorArray[pkg.ErrorShopNoItem] = "商店中无此物品"
	errorArray[pkg.ErrorShopSubCost] = "扣除费用失败"
	errorArray[pkg.ErrorShopBuySuccess] = "购买物品成功"
	errorArray[pkg.ErrorShopPetMoveIn] = "购买宠物移入宠物包裹失败"

	errorArray[pkg.ErrorAttrIllegal] = "属性非法"

	errorArray[pkg.ErrorTradeExist] = "交易不存在"
	errorArray[pkg.ErrorTradeRequestor] = "您已经存在交易"
	errorArray[pkg.ErrorTradeResponsor] = "对方已经在交易"
	errorArray[pkg.ErrorTradeLocked] = "交易已经锁定"
	errorArray[pkg.ErrorTradeConfirmed] = "交易已经确认"
	errorArray[pkg.ErrorTradeNoItem] = "不存在交易物品"
	errorArray[pkg.ErrorTradeNoPet] = "不存在交易宠物"
	errorArray[pkg.ErrorTradeOwnerNotLocked] = "未锁定，不能交易"
	errorArray[pkg.ErrorTradeOtherNotLocked] = "对方未锁定，不能交易"
	errorArray[pkg.ErrorTradePetStatus] = "该宠物无法交易"

	errorArray[pkg.ErrorGuildExist] = "帮派已经存在"
	errorArray[pkg.ErrorGuildNotExist] = "您没有加入帮派"
	errorArray[pkg.ErrorGuildMemberExist] = "帮派成员已经存在"
	errorArray[pkg.ErrorGuildMemberNotExist] = "帮派成员不存在"
	errorArray[pkg.ErrorGuildApplicantExist] = "入帮申请已存在"
	errorArray[pkg.ErrorGuildApplicantNotExist] = "入帮申请不存在"
	errorArray[pkg.ErrorGuildLeader] = "帮派中还有帮众时，帮主不可退出帮派"
	errorArray[pkg.ErrorGuildPermission] = "没有操作权限"
	errorArray[pkg.ErrorGuildJobDuplicate] = "该职位不能重复分配"
	errorArray[pkg.ErrorGuildImpeachExist] = "不能重复弹劾"
	errorArray[pkg.ErrorGuildImpeachNotExist] = "弹劾不存在"
	errorArray[pkg.ErrorGuildImpeachSelf] = "不能弹劾自己"
	errorArray[pkg.ErrorGuildLeaderOnline] = "帮主在线，不能弹劾"
	errorArray[pkg.ErrorGuildLeaderActive] = "帮主离线时间不超过七日，不能弹劾"
	errorArray[pkg.ErrorGuildImpeachVote] = "不能重复投票"
	errorArray[pkg.ErrorGuildImpeaching] = "不能删除弹劾的成员"
	errorArray[pkg.ErrorGuildFull] = "帮派人数已满"
	errorArray[pkg.ErrorGuildRewardItemFull] = "帮派赏功堂物品已满"
	errorArray[pkg.ErrorGuildRewardItemNotFound] = "帮派赏功堂奖励物品不存在"

	errorArray[pkg.ErrorContactExist] = "联系人已存在"
	errorArray[pkg.ErrorContactNotExist] = "联系人不存在"
	errorArray[pkg.ErrorContactGroupExist] = "目标联系人存在"
	errorArray[pkg.ErrorContactGroupNotExist] = "目标联系人不存在"
	errorArray[pkg.ErrorContactNotApply] = "好友申请不存在"
	errorArray[pkg.ErrorContactNotFriend] = "好友不存在"
	errorArray[pkg.ErrorContactTypeNotMatched] = "联系人类型不匹配"
	errorArray[pkg.ErrorContactInfoNotExist] = "玩家不存在"
	errorArray[pkg.ErrorContactBlackList] = "玩家在黑名单内"

	errorArray[pkg.ErrorMailExist] = "邮件已存在"
	errorArray[pkg.ErrorMailNotExist] = "邮件不存在"
	errorArray[pkg.ErrorMailInvalidAttr] = "邮件无效的属性"
	errorArray[pkg.ErrorMailInvalidItem] = "邮件无效的物品"
	errorArray[pkg.ErrorMailInvalidPet] = "邮件无效的宠物"
	errorArray[pkg.ErrorMailNoAttachment] = "邮件没有附件"
	errorArray[pkg.ErrorMailOutOfDate] = "邮件过期了"

	errorArray[pkg.ErrorTitleExist] = "称号已存在"
	errorArray[pkg.ErrorTitleNotExist] = "称号不存在"

	errorArray[pkg.ErrorGuardNotExist] = "侍从不存在"
	errorArray[pkg.ErrorGuardNotInContainer] = "侍从不在包裹内"
	errorArray[pkg.ErrorGuardDelGuard] = "释放侍从失败"
	errorArray[pkg.ErrorGuardActivated] = "侍从已经激活"
	errorArray[pkg.ErrorGuardActivateFail] = "激活侍从失败"
	errorArray[pkg.ErrorNotPasture] = "禁止放牧"
	errorArray[pkg.ErrorPastureFull] = "牧场已满"
	errorArray[pkg.ErrorPasturePetNotFound] = "牧场宠物不存在"
	errorArray[pkg.ErrorPastureNotMatchedMap] = "不在牧场地图，不能操作"
	errorArray[pkg.ErrorLadderNotFound] = "找不到玩家天梯信息"
	errorArray[pkg.ErrorLadderExist] = "玩家天梯信息已存在"
	errorArray[pkg.ErrorInstructionTypeIllegal] = "指令类型非法"
	errorArray[pkg.ErrorInstructionIdIllegal] = "指令ID非法"
	errorArray[pkg.ErrorInstructionIllegal] = "指令文字非法"
	errorArray[pkg.ErrorInstructionLimit] = "达到最大指令数上限"
	errorArray[pkg.ErrorInstructionModifyFire] = "集火指令不可修改"
	errorArray[pkg.ErrorInstructionNull] = "该指令不存在"
	errorArray[pkg.ErrorInstructionNeedFriend] = "该指令仅对友方有效"
	errorArray[pkg.ErrorInstructionNeedEnemy] = "该指令仅对敌方有效"
}

type sysError struct {
	errCode int32
	errDesc string
}

func NewError(err pkg.ErrorCode) pkg.SysError {
	return &sysError{
		errCode: int32(err),
		errDesc: errorArray[err],
	}
}

func (e *sysError) Code() int32 {
	return e.errCode
}

func (e *sysError) Error() string {
	return e.errDesc
}