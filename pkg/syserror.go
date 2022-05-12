///>本代码由测试工具自动生成,请勿手动修改
package pkg

type ErrorCode int

const (
	ErrorOk                 = ErrorCode(iota) ///>成功
	ErrorUnknown                              ///>系统错误
	ErrorLoginIllegal                         ///>非法的登陆请求
	ErrorGTOverload                           ///>网关连接负载过高
	ErrorWSOverload                           ///>服务器人数已满
	ErrorGSOverload                           ///>游戏人数已满
	ErrorGSNotFound                           ///>找不到游戏服务器
	ErrorLoginRepeat                          ///>不要重复提交登陆请求
	ErrorLoginPwd                             ///>登陆账号或密码不能为空
	ErrorLoginInfo                            ///>请完善登陆信息
	ErrorLoginPend                            ///>登陆请求正在处理中请稍候
	ErrorLoginAgain                           ///>账号从另一处登陆
	ErrorTimeoutKick                          ///>操作超时
	ErrorOnlineRelogin                        ///>正常的在线账号不允许重连
	ErrorUserRelogin                          ///>找不到重连的账号数据
	ErrorKickExitPend                         ///>请离或退出请求正在处理中请稍候
	ErrorUserNull                             ///>数据库中无此账号数据
	ErrorLogoutIllegal                        ///>非法的退出请求
	ErrorUserState                            ///>该请求无法在当前玩家状态下执行
	ErrorPlayerList                           ///>请先取得玩家列表
	ErrorPlayerPending                        ///>网络错误20
	ErrorPlayerIsGameing                      ///>该帐号下已有角色正在游戏中
	ErrorPlayerName                           ///>玩家名已存在
	ErrorPlayerNull                           ///>找不到指定的玩家
	ErrorPlayerDeleted                        ///>玩家已被删除
	ErrorPlayerNormal                         ///>玩家未被删除
	ErrorPlayerHasDeleted                     ///>超出当日删除上限
	ErrorPlayerNormalLimit                    ///>达到玩家列表上限
	ErrorPlayerDeleteLimit                    ///>达到删除列表上限
	ErrorPlayerRestoreLimit                   ///>超过恢复时效
	ErrorCacheTimeout                         ///>缓存加载超时
	ErrorPlayerId                             ///>玩家ID非法
	ErrorPlayerRace                           ///>玩家种族非法
	ErrorPlayerGender                         ///>玩家性别非法
	ErrorPlayerNameViolate                    ///>玩家名字包含非法字符
	ErrorPlayerSync                           ///>网络错误35
	ErrorLoginWrongUser                       ///>无此用户
	ErrorLoginWrongPassword                   ///>密码输入有误
	ErrorForbidLoginAccount                   ///>帐号禁止登陆
	ErrorAbnormalUserData                     ///>账号数据异常
	ErrorKeepAliveTimeout                     ///>账号心跳检测超时
	ErrorKicked                               ///>您已被请离服务器
	ErrorForbidLoginPlayer                    ///>玩家禁止登陆
	ErrorClientType                           ///>客户端版本错误
	ErrorReloginPend                          ///>重连请求正在处理中请稍候
	ErrorPlayerListPend                       ///>角色列表正在加载中请稍候
	ErrorPlayerCreatePend                     ///>角色正在创建中请稍候
	ErrorPlayerDestroyPend                    ///>角色正在销毁中请稍候
	ErrorPlayerRestorePend                    ///>角色正在恢复中请稍候
	ErrorEnterGSPend                          ///>正在进入游戏中请稍候
)

const (
	ErrorNoMovablePT     = ErrorCode(100 + iota) ///>无可移动点
	ErrorNoMap                                   ///>无地图信息
	ErrorPath                                    ///>路径非法有不可移动点
	ErrorMoveFast                                ///>移动速度过快
	ErrorNoJumpIfo                               ///>无跳转点信息
	ErrorInvalidJump                             ///>非法的跳转请求
	ErrorRepeatPt                                ///>移动目标点重复
	ErrorRepeatMap                               ///>出现地图重复
	ErrorItemInContainer                         ///>道具在包裹中
	ErrorNoPlace                                 ///>无道具可放置点
	ErrorPathNoPT                                ///>坐标点不在路径上
	ErrorInFight                                 ///>战斗中无法操作
	ErrorInView                                  ///>观战中无法操作
	ErrorInTeam                                  ///>队伍中无法操作
	ErrorInFollow                                ///>跟随中无法操作
	ErrorInvalidPt                               ///>无效的坐标
	ErrorScriptJump                              ///>脚本禁止跳转
	ErrorInConvoy                                ///>押镖中禁止跳转
	ErrorForbidMove                              ///>禁止移动
	ErrorForbidJump                              ///>禁止跳转
)

const (
	ErrorItemSwapSite          = ErrorCode(200 + iota) ///>物品交换位置失败
	ErrorItemNotEnough                                 ///>物品不足
	ErrorItemNotInContainer                            ///>物品不在包裹内
	ErrorItemMoveIn                                    ///>移入物品失败
	ErrorItemMoveOut                                   ///>移出物品失败
	ErrorItemSite                                      ///>物品位置异常
	ErrorItemRearrangeSite                             ///>整理物品失败
	ErrorItemHasMap                                    ///>物品在地图上
	ErrorItemCanMoveIn                                 ///>物品不能移入仓库
	ErrorItemCanUse                                    ///>该物品无法直接使用
	ErrorItemCanEquip                                  ///>无法装备
	ErrorItemNotExist                                  ///>物品不存在
	ErrorItemContainerCapacity                         ///>包裹空间不足
	ErrorItemNotOwner                                  ///>玩家不拥有该装备
	ErrorItemNotEquip                                  ///>物品不是装备
	ErrorItemContainerFull                             ///>包裹已满
	ErrorItemNotUseIsFight                             ///>战斗中不能对目标使用该道具
	ErrorItemApplyTarget                               ///>使用战斗道具时目标非法
	ErrorItemBound                                     ///>物品绑定的
	ErrorItemUserNotMatched                            ///>不能对目标使用物品
	ErrorItemAttrRequire                               ///>不满足穿戴属性要求
	ErrorItemCanUseAll                                 ///>该物品不能批量使用
	ErrorItemNoNeed                                    ///>该物品无需使用
	ErrorItemCountOverflow                             ///>物品数量超出上限，无法购买
	ErrorBindGoldNotEnough                             ///>银币不足，无法购买
	ErrorGoldNotEnough                                 ///>金币不足，无法购买
	ErrorBindIngotNotEnough                            ///>银元宝不足，无法购买
	ErrorIngotNotEnough                                ///>元宝不足，无法购买
	ErrorItemCantEquip                                 ///>无法脱下装备
)

const (
	ErrorBuffTriggerTime      = ErrorCode(300 + iota) ///>BUFF未满足触发时机
	ErrorBuffCooldown                                 ///>BUFF冷却中
	ErrorBuffTriggerTotal                             ///>超出BUFF最大触发次数
	ErrorBuffTriggerPerRound                          ///>超出BUFF每回合最大触发次数
	ErrorBuffDelay                                    ///>BUFF需要延时
	ErrorBuffRateUnreaches                            ///>BUFF触发概率未满足
	ErrorBuffTriggerSearch                            ///>BUFF查找触发目标失败
	ErrorBuffTriggerCondition                         ///>BUFF触发条件未满足
	ErrorBuffActionSearch                             ///>BUFF查找行为目标失败
	ErrorBuffActionTarget                             ///>BUFF无行为目标
	ErrorBuffActionExecute                            ///>BUFF执行行为失败
	ErrorBuffActorDied                                ///>BUFF不能在死亡时触发
)

const (
	ErrorSkillActor           = ErrorCode(400 + iota) ///>当前状态无法施放该技能
	ErrorSkillDelay                                   ///>技能需要延时
	ErrorSkillCooldown                                ///>技能冷却中
	ErrorSkillMajorTarget                             ///>技能主目标非法
	ErrorSkillTargetType                              ///>未知的技能目标类型
	ErrorSkillShotCondition                           ///>技能施放条件不满足
	ErrorSkillShotCost                                ///>技能消耗不足
	ErrorSkillSearchTarget                            ///>查找技能目标失败
	ErrorSkillExtraTargetType                         ///>未知的技能额外目标类型
	ErrorSkillDamageType                              ///>未知的技能伤害类型
	ErrorSkillShotDamage                              ///>技能造成伤害失败
	ErrorSkillDebuffType                              ///>未知的DEBUFF命中类型
	ErrorSkillShotDebuff                              ///>技能添加DEBUFF失败
	ErrorSkillShotBuff                                ///>技能添加BUFF失败
	ErrorSkillActorControl                            ///>技能施放者被控制
	ErrorSkillTargetEmpty                             ///>技能没有施放目标
	ErrorSkillNotFound                                ///>找不到指定的技能
)

const (
	ErrorAITriggerTime      = ErrorCode(500 + iota) ///>AI触发类型不符
	ErrorAIExpire                                   ///>AI已过期
	ErrorAICooldown                                 ///>AI冷却中
	ErrorAITriggerTotal                             ///>超出AI最大触发次数
	ErrorAITriggerPerRound                          ///>超出AI每回合最大触发次数
	ErrorAIDelay                                    ///>AI需要延时
	ErrorAIRateUnreaches                            ///>AI触发概率未满足
	ErrorAITriggerSearch                            ///>AI查找触发目标失败
	ErrorAITriggerCondition                         ///>AI触发条件未满足
	ErrorAIActionSearch                             ///>AI查找行为目标失败
	ErrorAIActionTarget                             ///>AI无行为目标
	ErrorAIActionExecute                            ///>AI执行行为失败
	ErrorAIActorDied                                ///>AI不能在死亡时触发
)

const (
	ErrorPlayerIsFight         = ErrorCode(600 + iota) ///>玩家正在战斗中
	ErrorPlayerNotFight                                ///>玩家不在战斗中
	ErrorPlayerIsView                                  ///>玩家正在观战中
	ErrorPlayerNotView                                 ///>玩家不在观战中
	ErrorFighterActorNull                              ///>找不到战斗动作发起者
	ErrorFighterTargetNull                             ///>找不到战斗动作目标
	ErrorFightSceneState                               ///>当前战斗状态不允许
	ErrorFighterRepeatOperate                          ///>请不要重复发送战斗操作
	ErrorFighterChaos                                  ///>混乱中行动失败
	ErrorFighterIce                                    ///>封印中行动失败
	ErrorFighterSleep                                  ///>昏睡中行动失败
	ErrorFighterHuaWu                                  ///>技能被化无释放失败
	ErrorFighterDied                                   ///>已死亡行动失败
	ErrorFighterCatchTarget                            ///>该目标无法被捕捉
	ErrorFighterCatchLevel                             ///>捕捉等级不足
	ErrorFighterCatchAdd                               ///>宠物栏已满
	ErrorFighterCmd                                    ///>非法的战斗指令
	ErrorFighterEnterView                              ///>进入观战失败
	ErrorFighterLeaveView                              ///>退出观战失败
	ErrorFighterIdle                                   ///>待机中行动失败
	ErrorFighterGeneralAtk                             ///>强制物理攻击
	ErrorFighterIgnoreProtect                          ///>忽略目标保护
	ErrorFighterForget                                 ///>技能被遗忘执行物理普攻
	ErrorFighterDefense                                ///>强制防御
	ErrorOperateForbid                                 ///>该操作被禁止行动失败
	ErrorFighterIgnoreDefense                          ///>忽略目标防御
	ErrorPetLoyaltyForceAttack                         ///>因忠诚度不足无法施放技能
	ErrorPetLoyaltyForceEscape                         ///>因忠诚度不足逃跑了
	ErrorPetSummonOut                                  ///>已出战或参战过的宠物不可再出战
	ErrorPetCatchFail                                  ///>捕捉宠物失败
	ErrorEscapeFail                                    ///>逃跑失败
)

const (
	ErrorTeamExist             = ErrorCode(700 + iota) ///>玩家已经有队伍
	ErrorTeamAddMember                                 ///>队伍添加成员失败
	ErrorTeamDelMember                                 ///>队伍移除成员失败
	ErrorTeamNotMatched                                ///>此队伍已停止招募
	ErrorTeamIsLeader                                  ///>玩家是队长
	ErrorTeamNotIsLeader                               ///>玩家不是队长
	ErrorTeamIsTempLeave                               ///>玩家是暂离的
	ErrorTeamNotSameGuild                              ///>当前队长所在的位置无法归队
	ErrorTeamNotExist                                  ///>玩家没有队伍
	ErrorTeamApplicantExist                            ///>申请已存在，无法操作
	ErrorTeamApplicantNotExist                         ///>申请不存在，无法操作
	ErrorTeamCantApply                                 ///>队长已取消招募
	ErrorTeamConvoy                                    ///>押镖中不能组队
)

const (
	ErrorQuestAcceptCondition    = ErrorCode(800 + iota) ///>玩家不满足任务接取条件
	ErrorQuestNotFound                                   ///>任务不存在
	ErrorQuestAdd                                        ///>添加任务失败
	ErrorQuestCanAdd                                     ///>不能添加任务
	ErrorQuestCanAccept                                  ///>不能接取任务
	ErrorQuestCanCommit                                  ///>不能完成任务
	ErrorQuestCanAbandon                                 ///>不能放弃任务
	ErrorQuestCanFailed                                  ///>不能失败任务
	ErrorQuestHasAccepted                                ///>任务已经接取
	ErrorQuestHasAbandoned                               ///>任务已经放弃
	ErrorQuestHasFailed                                  ///>任务已经失败
	ErrorQuestHasFinished                                ///>任务已经完成
	ErrorQuestHasNotActivated                            ///>任务未激活
	ErrorQuestHasMutex                                   ///>已经接取了互斥任务
	ErrorQuestHasLine                                    ///>已经接取了线性任务的前置任务
	ErrorQuestOperationCondition                         ///>玩家不满足任务操作条件
	ErrorQuestMustAlone                                  ///>任务必须单人
	ErrorQuestMustTeam                                   ///>任务必须组队
	ErrorQuestTeamAverageLevel                           ///>队伍平均等级不满足任务要求
	ErrorQuestTeamNumLower                               ///>队伍成员数量不足
	ErrorQuestTeamNumUpper                               ///>队伍成员数量太多
	ErrorQuestLogout                                     ///>任务已经注销
	ErrorQuestCanFight                                   ///>任务条件不满足
	ErrorQuestCanReady                                   ///>不能准备完成任务
	ErrorQuestTeamOrgGuild                               ///>队伍成员不是同一帮派
	ErrorQuestTeamOrgJob                                 ///>队伍成员不是同一门派
	ErrorQuestMustAllAccept                              ///>队伍所有成员必须满足接取条件
	ErrorQuestCantTeam                                   ///>组队中不能接押镖任务
)

const (
	ErrorPetNotInContainer  = ErrorCode(900 + iota) ///>宠物不在包裹内
	ErrorPetInContainer                             ///>宠物已经在包裹内
	ErrorPetShow                                    ///>宠物展示失败
	ErrorPetHide                                    ///>宠物隐藏失败
	ErrorPetSetLineup                               ///>宠物上阵失败
	ErrorPetClrLineup                               ///>宠物下阵失败
	ErrorPetInWarehouse                             ///>宠物在仓库内
	ErrorPetMoveIn                                  ///>移入宠物失败
	ErrorPetMoveOut                                 ///>移出宠物失败
	ErrorPetIsShowed                                ///>宠物已经展示
	ErrorPetIsLineup                                ///>宠物已经上阵
	ErrorPetCanMoveIn                               ///>宠物仓库已满，不能移入
	ErrorPetShowed                                  ///>展示宠物不能放入仓库
	ErrorPetBind                                    ///>绑定宠物不能放入仓库
	ErrorPetLocked                                  ///>锁定宠物不能放入仓库
	ErrorPetLineup                                  ///>上阵宠物不能放入仓库
	ErrorPetNotFound                                ///>宠物不存在
	ErrorPetNotLineup                               ///>没有上阵宠物
	ErrorPetNotShowed                               ///>没有展示宠物
	ErrorPetNotOwner                                ///>宠物没有主人
	ErrorGuardNotFound                              ///>侍从不存在
	ErrorGuardContainerFull                         ///>上阵侍从已满
	ErrorPetContainerFull                           ///>宠物包裹已满
	ErrorPetInPasture                               ///>宠物放牧中
	ErrorPetNotYour                                 ///>该宠物不属于您
	ErrorPetReplace                                 ///>宠物替换失败
	ErrorPetLevel                                   ///>宠物与角色等级差距过大无法出战
)

const (
	ErrorShopSellAdd      = ErrorCode(1000 + iota) ///>出售物品添加回购列表失败
	ErrorShopCantSell                              ///>物品不可出售
	ErrorShopBindGoldlack                          ///>绑定金币不足
	ErrorShopSplitItem                             ///>出售分离物品失败
	ErrorShopOpen                                  ///>商店打开失败
	ErrorShopNoItem                                ///>商店中无此物品
	ErrorShopSubCost                               ///>扣除费用失败
	ErrorShopBuySuccess                            ///>购买物品成功
	ErrorShopPetMoveIn                             ///>购买宠物移入宠物包裹失败
)

const (
	ErrorAttrIllegal = ErrorCode(1100 + iota) ///>属性非法
)

const (
	ErrorTradeExist          = ErrorCode(1200 + iota) ///>交易不存在
	ErrorTradeRequestor                               ///>您已经存在交易
	ErrorTradeResponsor                               ///>对方已经在交易
	ErrorTradeLocked                                  ///>交易已经锁定
	ErrorTradeConfirmed                               ///>交易已经确认
	ErrorTradeNoItem                                  ///>不存在交易物品
	ErrorTradeNoPet                                   ///>不存在交易宠物
	ErrorTradeOwnerNotLocked                          ///>未锁定，不能交易
	ErrorTradeOtherNotLocked                          ///>对方未锁定，不能交易
	ErrorTradePetStatus                               ///>该宠物无法交易
)

const (
	ErrorGuildExist              = ErrorCode(1300 + iota) ///>帮派已经存在
	ErrorGuildNotExist                                    ///>您没有加入帮派
	ErrorGuildMemberExist                                 ///>帮派成员已经存在
	ErrorGuildMemberNotExist                              ///>帮派成员不存在
	ErrorGuildApplicantExist                              ///>入帮申请已存在
	ErrorGuildApplicantNotExist                           ///>入帮申请不存在
	ErrorGuildLeader                                      ///>帮派中还有帮众时，帮主不可退出帮派
	ErrorGuildPermission                                  ///>没有操作权限
	ErrorGuildJobDuplicate                                ///>该职位不能重复分配
	ErrorGuildImpeachExist                                ///>不能重复弹劾
	ErrorGuildImpeachNotExist                             ///>弹劾不存在
	ErrorGuildImpeachSelf                                 ///>不能弹劾自己
	ErrorGuildLeaderOnline                                ///>帮主在线，不能弹劾
	ErrorGuildLeaderActive                                ///>帮主离线时间不超过七日，不能弹劾
	ErrorGuildImpeachVote                                 ///>不能重复投票
	ErrorGuildImpeaching                                  ///>不能删除弹劾的成员
	ErrorGuildFull                                        ///>帮派人数已满
	ErrorGuildRewardItemFull                              ///>帮派赏功堂物品已满
	ErrorGuildRewardItemNotFound                          ///>帮派赏功堂奖励物品不存在
)

const (
	ErrorContactExist          = ErrorCode(1400 + iota) ///>联系人已存在
	ErrorContactNotExist                                ///>联系人不存在
	ErrorContactGroupExist                              ///>目标联系人存在
	ErrorContactGroupNotExist                           ///>目标联系人不存在
	ErrorContactNotApply                                ///>好友申请不存在
	ErrorContactNotFriend                               ///>好友不存在
	ErrorContactTypeNotMatched                          ///>联系人类型不匹配
	ErrorContactInfoNotExist                            ///>玩家不存在
	ErrorContactBlackList                               ///>玩家在黑名单内
)

const (
	ErrorMailExist        = ErrorCode(1500 + iota) ///>邮件已存在
	ErrorMailNotExist                              ///>邮件不存在
	ErrorMailInvalidAttr                           ///>邮件无效的属性
	ErrorMailInvalidItem                           ///>邮件无效的物品
	ErrorMailInvalidPet                            ///>邮件无效的宠物
	ErrorMailNoAttachment                          ///>邮件没有附件
	ErrorMailOutOfDate                             ///>邮件过期了
)

const (
	ErrorTitleExist    = ErrorCode(1600 + iota) ///>称号已存在
	ErrorTitleNotExist                          ///>称号不存在
)

const (
	ErrorGuardNotExist          = ErrorCode(1700 + iota) ///>侍从不存在
	ErrorGuardNotInContainer                             ///>侍从不在包裹内
	ErrorGuardDelGuard                                   ///>释放侍从失败
	ErrorGuardActivated                                  ///>侍从已经激活
	ErrorGuardActivateFail                               ///>激活侍从失败
	ErrorNotPasture                                      ///>禁止放牧
	ErrorPastureFull                                     ///>牧场已满
	ErrorPasturePetNotFound                              ///>牧场宠物不存在
	ErrorPastureNotMatchedMap                            ///>不在牧场地图，不能操作
	ErrorLadderNotFound                                  ///>找不到玩家天梯信息
	ErrorLadderExist                                     ///>玩家天梯信息已存在
	ErrorInstructionTypeIllegal                          ///>指令类型非法
	ErrorInstructionIdIllegal                            ///>指令ID非法
	ErrorInstructionIllegal                              ///>指令文字非法
	ErrorInstructionLimit                                ///>达到最大指令数上限
	ErrorInstructionModifyFire                           ///>集火指令不可修改
	ErrorInstructionNull                                 ///>该指令不存在
	ErrorInstructionNeedFriend                           ///>该指令仅对友方有效
	ErrorInstructionNeedEnemy                            ///>该指令仅对敌方有效
	ErrorMax
)

type SysError interface {
	Code() int32
	Error() string
}
