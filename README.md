# Promock

Promock = Proxy + Mock

Promock根据匹配规则进行匹配，然后进行Proxy代理，或者直接返回Mock数据。

**架构设计**

Promock分为规则匹配和规则执行两部分。

## 规则匹配

目前，规则匹配只针对URL中的param参数进行匹配，参数param的数据类型有：

- INT: 整数
- FLOAT: 浮点数
- STRING: 字符串
- VERSION: 版本类型

我们在进行规则定制的时候，首先要定义可以参与匹配的参数，每个参数必须定义数据类型与优先级。

> **优先级**: 每个参数拥有*不同*的优先级，数字越大优先级越高；

比如，如下的参数定义：

- UID: 数据类型INT，优先级5
- Device: 数据类型STRING，优先级4
- Version: 数据类型VERSION，优先级3
- Platform: 数据类型STRING，优先级2
- Model: 数据类型STRING，优先级1

比如，如下的规则匹配：

1. UID = 632025298
2. Device = XFGH653ES2K
3. Version >= 2.7 & Platform = IOS
4. Version >= 2.7 & Model = HUAWEI
5. Version >= 2.7
6. Platfrom = IOS

> 上面的6条规则，按照优先级进行了排序！
>
> 每条规则里面多个条件只支持AND操作，不支持OR操作！

## 规则执行

### 1、两种策略

- Proxy
- Mock

同时加上延时控制，比如控制在500ms的时候返回！

### 2、PATH匹配

- ALL
- 前缀匹配
- 精确匹配

### 3、匹配优先级

`精确匹配` > `前缀匹配` > `ALL`

**高级匹配**:

- 第几次出现
- 前置条件（在某接口之后）
- 重置开始（清空计数、清空历史）

### 4、改写数据

包括：

- 改写Header
- 改写Path
- 改写Param
- 改写Host
