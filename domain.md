#### 关于域名到期删除规则实施的解释：

国际域名：

(1) 到期当天暂停解析，如果在72小时未续费，则修改域名DNS指向广告页面（停放）。域名到期后30-45天为域名保留期（不同注册商政策规定时间不同）

(2) 过了保留期域名将进入赎回期（REDEMPTIONPERIOD，为期30天）

(3) 过了赎回期域名将进入为期5天左右的删除期，删除期过后域名开放，任何人可注。

国内域名：

(1) 到期当天暂停解析，如果在72小时未续费，则修改域名DNS指向广告页面（停放）。35天内，可以自动续费。

(2) 过期后36－48天，将进入13天的高价赎回期，此期间域名无法管理。

(3) 过期后48天后仍未续费的，域名将随时被删除。


#### 关于域名状态的解释

cn域名各个状态说明：

以client开头的状态表示由客户端(注册商)可以增加的状态

以server开头的状态表示服务器端(CNNIC)操作增加的状态

既不以client开头也不以server开头的状态由服务器端管理


域名的状态解释：

ok 正常状态

inactive 非激活状态(注册的时候没有填写域名服务器，不能进行解析)

clientDeleteProhibited 禁止删除

serverDeleteProhibited 禁止删除

clientUpdateProhibited 禁止修改

serverUpdateProhibited 禁止修改

pendingDelete 正在删除过程中

pendingTransfer 正在转移过程中

clientTransferProhibited 禁止转移

serverTransferProhibited 禁止转移

clientRenewProhibited 禁止续费

serverRenewProhibited 禁止续费

clientHold 停止解析

serverHold 停止解析

pendingVerification 注册信息正在确认过程中
