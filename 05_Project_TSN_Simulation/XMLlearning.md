# 📚 XML 基础与解析实战（TSN 仿真项目专项）学习笔记

## 一、 XML 核心概念与目标

### 1.1 XML 的起源与背景

XML（eXtensible Markup Language，可扩展标记语言）起源于 **SGML（标准通用标记语言）** 和 **HTML（超文本标记语言）** 的发展历程，由 **W3C（万维网联盟）** 于 1998 年正式发布。

**核心目标**：XML 是一种**用来存储和传输数据的纯文本格式**，通过自定义的标签将数据组织成具有层级关系的**树状结构**。

**通俗理解**：XML 就像一个**带详细索引的档案袋**。每个档案袋（元素）都有明确的标签（自定义标签名），袋子里可以装具体物品（文本内容）或其他小档案袋（子元素），形成严格的层级关系。

### 1.2 XML 与 HTML 的本质区别

| 特性 | XML | HTML |
|------|-----|------|
| 设计目的 | 数据存储与传输 | 数据展示 |
| 标签 | 可自定义（可扩展） | 预定义（固定标签） |
| 语法严格性 | 极其严格 | 相对宽松 |
| 是否区分大小写 | 区分 | 不区分 |
| 是否必须闭合标签 | 必须 | 可省略 |
| 是否支持自定义属性 | 是 | 是 |
| 主要用途 | 配置文件、数据交换、文档格式 | 网页展示 |

### 1.3 XML 的核心价值

**为什么 TSN 仿真项目需要 XML？**

在 TSN 网络仿真项目中，C++ 软件不再需要硬编码数据，而是通过读取 XML 配置文件来驱动：

1. **数据与代码分离**：修改配置无需重新编译代码
2. **跨平台通用**：纯文本格式，任何语言都可以解析
3. **层级结构清晰**：天然适合表达网络拓扑、配置参数等树状数据
4. **工具链完善**：大量 XML 编辑器、验证工具可用

---

## 二、 XML 语法规则详解

### 2.1 XML 文档结构

```
┌─────────────────────────────────────────────────────────────┐
│                    XML 文档结构                              │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  1. XML 声明 (必须)                                         │
│     ↓                                                       │
│  2. 根元素 (必须，有且仅有一个)                              │
│     ↓                                                       │
│  3. 子元素 (可嵌套任意层级)                                  │
│     ↓                                                       │
│  4. 文本内容 (叶节点的数据)                                  │
│                                                             │
│  ┌─────────────────────────────────────────────────────┐   │
│  │ <?xml version="1.0" encoding="UTF-8"?>             │   │
│  │ <根元素>                                            │   │
│  │   <子元素 属性="值">                               │   │
│  │     <孙元素>文本内容</孙元素>                       │   │
│  │   </子元素>                                         │   │
│  │ </根元素>                                           │   │
│  └─────────────────────────────────────────────────────┘   │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### 2.2 XML 声明 (Declaration)

**位置**：必须是文件的第一行、第一列

**标准格式**：
```xml
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
```

**参数说明**：
| 参数 | 必选 | 说明 | 常见值 |
|------|------|------|--------|
| version | 是 | XML 标准版本 | "1.0", "1.1" |
| encoding | 否 | 字符编码 | "UTF-8", "GBK", "ISO-8859-1" |
| standalone | 否 | 是否独立文档 | "yes"（无需外部DTD）, "no" |

### 2.3 元素 (Element)

**定义**：元素是 XML 的基本构建单元，由成对的尖括号组成。

**语法规则**：
1. **必须正确闭合**：每个起始标签必须有对应的结束标签
2. **区分大小写**：`<Switch>` 和 `<switch>` 是不同的元素
3. **不能交叉嵌套**：`<a><b></a></b>` 是非法的
4. **根元素唯一**：整个文档只能有一个根元素

**元素类型**：

```xml
<!-- 1. 空元素（自闭合） -->
<Switch id="SW_1" />

<!-- 2. 包含文本内容 -->
<Name>TSN_Switch_001</Name>

<!-- 3. 包含子元素 -->
<Switch id="SW_1">
    <Type>CQF_Enabled</Type>
    <PortCount>8</PortCount>
</Switch>

<!-- 4. 混合内容（元素+文本） -->
<Description>交换机型号：<Model>TSN-9000</Model>，支持CQF</Description>
```

### 2.4 属性 (Attribute)

**定义**：属性是写在起始标签内部的键值对，用于提供元素的额外信息。

**语法规则**：
1. 属性必须用引号包裹（单引号或双引号）
2. 同一个元素不能有重复的属性名
3. 属性顺序不重要
4. 属性只能在起始标签中，不能在结束标签中

```xml
<!-- 属性示例 -->
<Switch id="SW_001" type="core" location="zone_A">
    <Port interface="10G" status="up"/>
</Switch>
```

**元素 vs 属性**：
```xml
<!-- 使用属性存储简短标识信息 -->
<Switch id="SW_001" model="TSN-9000" />

<!-- 使用子元素存储可能包含层级结构的数据 -->
<Switch id="SW_001">
    <Model>TSN-9000</Model>
    <Manufacturer>
        <Name>Huawei</Name>
        <Country>China</Country>
    </Manufacturer>
</Switch>
```

### 2.5 文本内容与转义

**特殊字符必须转义**：

| 原字符 | 转义序列 | 说明 |
|--------|----------|------|
| `<` | `&lt;` | 小于号 |
| `>` | `&gt;` | 大于号 |
| `&` | `&amp;` | 和号 |
| `"` | `&quot;` | 双引号 |
| `'` | `&apos;` | 单引号 |

**示例**：
```xml
<!-- 错误：未转义 -->
<Formula>x < y && x > 0</Formula>

<!-- 正确：使用转义序列 -->
<Formula>x &lt; y &amp;&amp; x &gt; 0</Formula>
```

### 2.6 注释 (Comment)

**语法**：`<!-- 注释内容 -->`

**规则**：
1. 注释不能出现在 XML 声明之前
2. 注释不能嵌套
3. 注释内容中不能包含 `--`
4. 注释可以跨越多行

```xml
<?xml version="1.0" encoding="UTF-8"?>
<!-- 这是单行注释 -->

<!--
这是多行注释
可以跨越多行
TSN 仿真项目配置文件
-->
<Configuration>
    <ParamA>100</ParamA>
</Configuration>
```

### 2.7 CDATA 区域

**用途**：当文本内容包含大量特殊字符时，使用 CDATA 可以避免转义。

**语法**：`<![CDATA[ 内容 ]]>`

```xml
<!-- 不使用 CDATA：需要大量转义 -->
<Script>if (x &lt; y &amp;&amp; y &gt; z) { alert(&quot;hello&quot;); }</Script>

<!-- 使用 CDATA：无需转义，原样保留 -->
<Script><![CDATA[if (x < y && y > z) { alert("hello"); }]]></Script>
```

### 2.8 处理指令 (Processing Instruction)

**用途**：向应用程序传递信息。

```xml
<?xml version="1.0" encoding="UTF-8"?>
<?xml-stylesheet type="text/xsl" href="style.xsl"?>
<Configuration>
    <Param>value</Param>
</Configuration>
```

---

## 三、 XML 文档验证

### 3.1 格式良好的 XML (Well-Formed)

一个"格式良好"的 XML 文档必须满足：

| 规则 | 说明 |
|------|------|
| 单一根元素 | 有且仅有一个根元素 |
| 正确嵌套 | 标签不能交叉 |
| 大小写敏感 | 开始标签和结束标签必须完全匹配 |
| 属性引号 | 属性值必须用引号包裹 |
| 唯一属性 | 同一元素不能有重复属性 |
| 转义特殊字符 | `<`, `>`, `&` 必须转义 |

### 3.2 有效的 XML (Valid)

"有效"的 XML 文档除了格式良好外，还必须符合**文档类型定义（DTD）**或**XML Schema**定义的规则。

### 3.3 DTD (Document Type Definition)

**定义**：DTD 是用于定义 XML 文档结构的早期验证标准。

```xml
<!-- 内部 DTD 示例 -->
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE Topology [
    <!ELEMENT Topology (Nodes, Links)>
    <!ELEMENT Nodes (Switch+, EndSystem*)>
    <!ELEMENT Switch (Type, PortCount)>
    <!ELEMENT EndSystem (Type)>
    <!ELEMENT Type (#PCDATA)>
    <!ELEMENT PortCount (#PCDATA)>
    <!ELEMENT Links (Link+)>
    <!ELEMENT Link EMPTY>
    <!ATTLIST Link id ID #REQUIRED
                  source CDATA #REQUIRED
                  target CDATA #REQUIRED>
]>
<Topology>
    <Nodes>
        <Switch id="SW_1">
            <Type>TSN_Switch</Type>
            <PortCount>8</PortCount>
        </Switch>
    </Nodes>
    <Links>
        <Link id="L1" source="SW_1" target="ES_1"/>
    </Links>
</Topology>
```

**DTD 关键字说明**：
| 符号 | 说明 |
|------|------|
| `+` | 出现一次或多次 |
| `*` | 出现零次或多次 |
| `?` | 出现零次或一次 |
| `|` | 或（选择） |
| `,` | 与（顺序） |
| `#PCDATA` | 纯文本内容 |
| `#REQUIRED` | 属性必须存在 |
| `#IMPLIED` | 属性可选 |
| `#FIXED` | 属性固定值 |
| `EMPTY` | 空元素 |

### 3.4 XML Schema (XSD)

**定义**：XML Schema 是 DTD 的现代替代品，使用 XML 语法编写，功能更强大。

```xml
<?xml version="1.0" encoding="UTF-8"?>
<xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema">

    <!-- 定义根元素 -->
    <xs:element name="Topology">
        <xs:complexType>
            <xs:sequence>
                <xs:element name="Nodes" type="NodesType"/>
                <xs:element name="Links" type="LinksType"/>
            </xs:sequence>
        </xs:complexType>
    </xs:element>

    <!-- 定义节点类型 -->
    <xs:complexType name="NodesType">
        <xs:choice maxOccurs="unbounded">
            <xs:element name="Switch" type="SwitchType"/>
            <xs:element name="EndSystem" type="EndSystemType"/>
        </xs:choice>
    </xs:complexType>

    <!-- 定义交换机类型 -->
    <xs:complexType name="SwitchType">
        <xs:sequence>
            <xs:element name="Type" type="xs:string"/>
            <xs:element name="PortCount" type="xs:integer"/>
        </xs:sequence>
        <xs:attribute name="id" type="xs:string" use="required"/>
    </xs:complexType>

</xs:schema>
```

### 3.5 DTD vs XML Schema 对比

| 特性 | DTD | XML Schema |
|------|-----|------------|
| 语法 | 专属语法 | XML 语法 |
| 数据类型 | 弱（仅有 PCDATA） | 强（丰富的内置类型） |
| 命名空间 | 不直接支持 | 完全支持 |
| 继承/扩展 | 不支持 | 支持复杂类型继承 |
| 可读性 | 较差 | 较好 |
| 工具支持 | 广泛 | 良好 |

---

## 四、 命名空间 (Namespace)

### 4.1 为什么需要命名空间

当不同来源的 XML 文档混合使用时，元素名冲突是常见问题。

```xml
<!-- 冲突示例：两个不同定义的 Switch -->
<Switch>
    <Type>Network Switch</Type>  <!-- 网络设备类型 -->
</Switch>

<Switch>
    <Type>Electric Switch</Type>  <!-- 电气开关类型 -->
</Switch>
```

### 4.2 命名空间语法

```xml
<!-- 定义命名空间 -->
<配置 xmlns:net="http://example.com/network"
       xmlns:elec="http://example.com/electric">

    <!-- 使用 net 前缀 -->
    <net:Switch>
        <net:Type>Cisco_CBS</net:Type>
    </net:Switch>

    <!-- 使用 elec 前缀 -->
    <elec:Switch>
        <elec:Type>Power_Switch</elec:Type>
    </elec:Switch>

</配置>
```

### 4.3 默认命名空间

```xml
<!-- 默认命名空间：无需前缀 -->
<Configuration xmlns="http://example.com/tsn">
    <Switch>
        <Type>TSN_Switch</Type>
    </Switch>
</Configuration>
```

---

## 五、 TSN 仿真项目 XML 架构

### 5.1 四大核心配置文件

根据《项目建议书》的要求，C++ 软件通过读取以下 4 个 XML 配置文件来驱动：

```
┌─────────────────────────────────────────────────────────────┐
│           TSN 仿真项目 XML 配置体系                          │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  ┌─────────────────┐    ┌─────────────────┐              │
│  │  XX_topo.xml   │    │  XX_msg.xml     │              │
│  │  拓扑文件       │    │  消息流文件     │              │
│  │  - 节点定义    │    │  - 流量定义     │              │
│  │  - 链路连接    │    │  - 周期/优先级  │              │
│  └─────────────────┘    └─────────────────┘              │
│                                                             │
│  ┌─────────────────┐    ┌─────────────────┐              │
│  │  XX_para.xml    │    │  XX_config.xml  │              │
│  │  参数配置       │    │  配置结果       │              │
│  │  - 同步精度    │    │  - 运行结果     │              │
│  │  - 机制映射    │    │  - GCL 列表    │              │
│  └─────────────────┘    └─────────────────┘              │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### 5.2 拓扑文件 (XX_topo.xml)

**作用**：定义网络中的硬件实体与连接。

```xml
<?xml version="1.0" encoding="UTF-8"?>
<Topology>
    <Meta>
        <Version>1.0</Version>
        <Author>TSN Planner</Author>
        <Created>2026-01-15</Created>
    </Meta>

    <Nodes>
        <!-- 交换机节点 -->
        <Switch id="SW_CORE_1" type="tsn_switch" location="zone_a">
            <Property name="model">TSN-9000</Property>
            <Property name="ports">8</Property>
            <Property name="speed">10Gbps</Property>
        </Switch>

        <Switch id="SW_EDGE_1" type="tsn_switch" location="zone_b">
            <Property name="model">TSN-6000</Property>
            <Property name="ports">4</Property>
            <Property name="speed">1Gbps</Property>
        </Switch>

        <!-- 终端系统节点 -->
        <EndSystem id="ES_SENSOR_1" type="sensor" location="zone_b">
            <Property name="dataRate">100Mbps</Property>
            <Property name="criticality">high</Property>
        </EndSystem>

        <EndSystem id="ES_CAMERA_1" type="camera" location="zone_b">
            <Property name="dataRate">1Gbps</Property>
            <Property name="criticality">medium</Property>
        </EndSystem>

        <EndSystem id="ES_PC_1" type="computer" location="zone_c">
            <Property name="dataRate">1Gbps</Property>
            <Property name="criticality">low</Property>
        </EndSystem>
    </Nodes>

    <Links>
        <Link id="L1" source="SW_CORE_1" target="SW_EDGE_1" delay="1us">
            <Bandwidth>10Gbps</Bandwidth>
        </Link>

        <Link id="L2" source="SW_EDGE_1" target="ES_SENSOR_1" delay="0.5us">
            <Bandwidth>1Gbps</Bandwidth>
        </Link>

        <Link id="L3" source="SW_EDGE_1" target="ES_CAMERA_1" delay="0.5us">
            <Bandwidth>1Gbps</Bandwidth>
        </Link>

        <Link id="L4" source="SW_EDGE_1" target="ES_PC_1" delay="0.5us">
            <Bandwidth>1Gbps</Bandwidth>
        </Link>
    </Links>
</Topology>
```

### 5.3 消息流文件 (XX_msg.xml)

**作用**：定义网络中传输的流量（最高 2048 条）。

```xml
<?xml version="1.0" encoding="UTF-8"?>
<TrafficFlows>
    <Meta>
        <TotalFlows>5</TotalFlows>
        <Created>2026-01-15</Created>
    </Meta>

    <!-- ST 流：时间敏感周期性控制流 -->
    <Flow id="F001" type="ST">
        <Description>传感器控制指令流</Description>
        <Source>ES_SENSOR_1</Source>
        <Destination>ES_ACTUATOR_1</Destination>
        <Period unit="us">500</Period>
        <PayloadSize unit="byte">64</PayloadSize>
        <Priority>6</Priority>
        <Deadline unit="us">400</Deadline>
        <Mechanism>CQF</Mechanism>
    </Flow>

    <!-- AVB 流：音视频桥接流 -->
    <Flow id="F002" type="AVB">
        <Description>摄像头视频流</Description>
        <Source>ES_CAMERA_1</Source>
        <Destination>ES_DISPLAY_1</Destination>
        <Period unit="ms">33</Period>
        <PayloadSize unit="byte">1500</PayloadSize>
        <Priority>4</Priority>
        <Class>Class_A</Class>
        <Mechanism>CBS</Mechanism>
    </Flow>

    <!-- AVB Class B 流 -->
    <Flow id="F003" type="AVB">
        <Description>音频流</Description>
        <Source>ES_MIC_1</Source>
        <Destination>ES_SPEAKER_1</Destination>
        <Period unit="ms">100</Period>
        <PayloadSize unit="byte">300</PayloadSize>
        <Priority>3</Priority>
        <Class>Class_B</Class>
        <Mechanism>CBS</Mechanism>
    </Flow>

    <!-- BE 流：尽力而为流 -->
    <Flow id="F004" type="BE">
        <Description>普通数据流</Description>
        <Source>ES_PC_1</Source>
        <Destination>ES_SERVER_1</Destination>
        <Priority>0</Priority>
        <Mechanism>SP</Mechanism>
    </Flow>
</TrafficFlows>
```

### 5.4 参数配置文件 (XX_para.xml)

**作用**：调度算法的全局规则与环境。

```xml
<?xml version="1.0" encoding="UTF-8"?>
<ConfigurationParameters>
    <TimeSynchronization>
        <Protocol>IEEE_802.1AS</Protocol>
        <Precision unit="ns">100</Precision>
        <SyncInterval unit="ms">100</SyncInterval>
    </TimeSynchronization>

    <GlobalTiming>
        <BaseCycle unit="ms">1</BaseCycle>
        <HyperCycle unit="ms">10</HyperCycle>
    </GlobalTiming>

    <MechanismMapping>
        <!-- 优先级到流控机制的映射 -->
        <Mapping priority="7" mechanism="TAS" queue="TC7"/>
        <Mapping priority="6" mechanism="CQF" queue="TC6"/>
        <Mapping priority="5" mechanism="CQF" queue="TC5"/>
        <Mapping priority="4" mechanism="CBS" class="A" queue="TC4"/>
        <Mapping priority="3" mechanism="CBS" class="B" queue="TC3"/>
        <Mapping priority="0-2" mechanism="SP" queue="TC0-2"/>
    </MechanismMapping>

    <AlgorithmParameters>
        <!-- CBS 参数 -->
        <CBS bandwidthPercent="50" idleSlope="500000000" sendSlope="-500000000"/>

        <!-- CQF 参数 -->
        <CQFSlotTime unit="us">250</CQFSlotTime>

        <!-- TAS 保护带 -->
        <TASGuardBand unit="us">10</TASGuardBand>
    </AlgorithmParameters>

    <NetworkConstraints>
        <MaxHopCount>10</MaxHopCount>
        <MaxLatency unit="us">2000</MaxLatency>
        <MaxJitter unit="us">100</MaxJitter>
    </NetworkConstraints>
</ConfigurationParameters>
```

### 5.5 配置结果文件 (XX_config.xml)

**作用**：C++ 调度算法运行完毕后，**写入保存**的结果。

```xml
<?xml version="1.0" encoding="UTF-8"?>
<ScheduleResult>
    <Meta>
        <Version>1.0</Version>
        <GeneratedBy>TSN_Simulator</GeneratedBy>
        <GeneratedAt>2026-01-15T10:30:00</GeneratedAt>
        <Status>Success</Status>
    </Meta>

    <GlobalMetrics>
        <TotalDelay>
            <Min unit="us">50</Min>
            <Max unit="us">850</Max>
            <Average unit="us">320</Average>
        </TotalDelay>
        <Jitter unit="us">45</Jitter>
        <BandwidthUtilization percent="78"/>
        <ScheduleFeasibility>true</ScheduleFeasibility>
    </GlobalMetrics>

    <GateControlLists>
        <!-- SW_CORE_1 的 GCL -->
        <GCL node="SW_CORE_1">
            <Entry time="0" mask="11110000" action="open"/>
            <Entry time="250" mask="11110000" action="close"/>
            <Entry time="500" mask="11110000" action="open"/>
            <Entry time="750" mask="11110000" action="close"/>
            <CycleTime unit="us">1000</CycleTime>
        </GCL>

        <!-- SW_EDGE_1 的 GCL -->
        <GCL node="SW_EDGE_1">
            <Entry time="0" mask="11000000" action="open"/>
            <Entry time="125" mask="11000000" action="close"/>
            <Entry time="250" mask="00110000" action="open"/>
            <Entry time="375" mask="00110000" action="close"/>
            <Entry time="500" mask="11000000" action="open"/>
            <CycleTime unit="us">1000</CycleTime>
        </GCL>
    </GateControlLists>

    <PerFlowResults>
        <FlowResult flowId="F001">
            <Path>ES_SENSOR_1 -> SW_EDGE_1 -> SW_CORE_1 -> ES_ACTUATOR_1</Path>
            <EndToEndDelay unit="us">380</EndToEndDelay>
            <Jitter unit="us">30</Jitter>
            <MechanismAssigned>CQF</MechanismAssigned>
        </FlowResult>

        <FlowResult flowId="F002">
            <Path>ES_CAMERA_1 -> SW_EDGE_1 -> SW_CORE_1 -> ES_DISPLAY_1</Path>
            <EndToEndDelay unit="us">650</EndToEndDelay>
            <Jitter unit="us">50</Jitter>
            <MechanismAssigned>CBS</MechanismAssigned>
        </FlowResult>
    </PerFlowResults>
</ScheduleResult>
```

---

## 六、 XML 解析方法

### 6.1 DOM 模型详解

**原理**：将整个 XML 文档加载到内存中，构建一棵完整的树结构。

**优点**：
- 随机访问任意节点
- 易于导航（父节点、子节点、兄弟节点）
- 易于修改和写入

**缺点**：
- 需要将整个文档加载到内存
- 对于大文档内存消耗大
- 解析速度相对较慢

### 6.2 Qt DOM 解析实战

#### 6.2.1 核心类说明

| 类名 | 说明 |
|------|------|
| `QDomDocument` | 代表整棵树，负责加载和保存 XML |
| `QDomElement` | 代表一个具体的标签元素 |
| `QDomNodeList` | 代表一组搜索出来的节点集合 |
| `QDomNode` | 基类，代表任何类型的节点 |
| `QDomAttr` | 代表元素的属性 |
| `QDomText` | 代表文本节点 |

#### 6.2.2 节点 vs 元素的区别

**重要概念**：
- **万物皆节点**：标签、属性、内部的纯文本、注释，统统都是节点
- **元素只是节点的一种**：只有带 `< >` 标签的节点才是元素（`QDomElement` 继承自 `QDomNode`）

```
节点类型继承关系：
QDomNode
├── QDomElement (元素节点)
├── QDomText (文本节点)
├── QDomAttr (属性节点)
├── QDomComment (注释节点)
├── QDomDocument (文档节点)
└── ... 其他类型
```

#### 6.2.3 黄金解析代码模板

```cpp
#include <QFile>
#include <QDomDocument>
#include <QDomElement>
#include <QDomNodeList>
#include <QDebug>

class XmlConfigParser {
public:
    bool loadFile(const QString& filePath) {
        QFile file(filePath);
        if (!file.open(QIODevice::ReadOnly | QIODevice::Text)) {
            qWarning() << "Cannot open file:" << filePath;
            return false;
        }

        QString errorMsg;
        int errorLine, errorColumn;
        if (!doc_.setContent(&file, &errorMsg, &errorLine, &errorColumn)) {
            qWarning() << "XML parse error at line" << errorLine
                       << "column" << errorColumn << ":" << errorMsg;
            file.close();
            return false;
        }

        file.close();
        return true;
    }

    bool saveFile(const QString& filePath) {
        QFile file(filePath);
        if (!file.open(QIODevice::WriteOnly | QIODevice::Text)) {
            qWarning() << "Cannot open file for writing:" << filePath;
            return false;
        }

        QTextStream stream(&file);
        stream.setCodec("UTF-8");
        doc_.save(stream, 4);
        file.close();
        return true;
    }

protected:
    QDomDocument doc_;
};
```

#### 6.2.4 遍历节点的方法

```cpp
// 方法1：按标签名搜索（深度递归查找）
QDomNodeList switchList = root.elementsByTagName("Switch");

for (int i = 0; i < switchList.count(); ++i) {
    QDomNode node = switchList.at(i);
    if (node.isElement()) {
        QDomElement elem = node.toElement();

        // 提取属性
        QString id = elem.attribute("id");
        QString type = elem.attribute("type");

        // 提取子元素的文本内容
        QString portCount = elem.firstChildElement("Property").text();

        qDebug() << "Switch:" << id << "Type:" << type;
    }
}

// 方法2：遍历直接子节点
QDomNode child = root.firstChild();
while (!child.isNull()) {
    if (child.isElement()) {
        QDomElement childElem = child.toElement();
        qDebug() << "Child tag:" << childElem.tagName();

        // 遍历当前元素的属性
        QDomNamedNodeMap attrs = childElem.attributes();
        for (int i = 0; i < attrs.count(); ++i) {
            QDomAttr attr = attrs.item(i).toAttr();
            qDebug() << "  Attr:" << attr.name() << "=" << attr.value();
        }
    }
    child = child.nextSibling();
}

// 方法3：使用 firstChildElement / nextSiblingElement
QDomElement metaElem = root.firstChildElement("Meta");
if (!metaElem.isNull()) {
    QString version = metaElem.firstChildElement("Version").text();
    QString author = metaElem.firstChildElement("Author").text();
}
```

#### 6.2.5 提取文本的陷阱

**错误做法**：
```cpp
// 错误：对包含多个子节点的父节点直接调用 .text()
QDomElement switchElem = ...;
QString wrongText = switchElem.text();
// 结果：所有子节点的文本会粘连在一起
```

**正确做法**：
```cpp
// 正确：逐个找到子元素，再调用 .text()
QDomElement switchElem = ...;
QString id = switchElem.attribute("id");
QString type = switchElem.firstChildElement("Type").text();
QString ports = switchElem.firstChildElement("Property").text();
```

#### 6.2.6 创建 XML 文档

```cpp
void createTopologyXml() {
    QDomDocument doc;

    // 创建 XML 声明
    QDomProcessingInstruction instr = doc.createProcessingInstruction(
        "xml", "version=\"1.0\" encoding=\"UTF-8\"");
    doc.appendChild(instr);

    // 创建根元素
    QDomElement root = doc.createElement("Topology");
    doc.appendChild(root);

    // 创建子元素
    QDomElement nodesElem = doc.createElement("Nodes");
    root.appendChild(nodesElem);

    // 创建 Switch 元素
    QDomElement switchElem = doc.createElement("Switch");
    switchElem.setAttribute("id", "SW_001");
    switchElem.setAttribute("type", "tsn_switch");

    // 添加子元素
    QDomElement typeElem = doc.createElement("Type");
    typeElem.appendChild(doc.createTextNode("TSN-9000"));
    switchElem.appendChild(typeElem);

    nodesElem.appendChild(switchElem);

    // 保存到文件
    QFile file("output.xml");
    if (file.open(QIODevice::WriteOnly)) {
        QTextStream stream(&file);
        stream.setCodec("UTF-8");
        doc.save(stream, 4);
        file.close();
    }
}
```

### 6.3 SAX 模型详解

**原理**：基于事件驱动的流式解析，边读边处理，不构建完整树。

**优点**：
- 无需将整个文档加载到内存
- 解析速度快
- 适合处理大文档

**缺点**：
- 只能顺序读取，不能随机访问
- 修改文档困难
- 编程相对复杂

```cpp
#include <QXmlSimpleReader>
#include <QXmlInputSource>

class SaxHandler : public QXmlContentHandler {
public:
    bool startElement(const QString& namespaceUri,
                      const QString& localName,
                      const QString& qName,
                      const QXmlAttributes& atts) override {
        if (localName == "Switch") {
            currentSwitch_.id = atts.value("id");
            currentSwitch_.type = atts.value("type");
        } else if (localName == "Type") {
            currentTag_ = "Type";
        }
        return true;
    }

    bool endElement(const QString& namespaceUri,
                   const QString& localName,
                   const QString& qName) override {
        if (localName == "Switch") {
            switches_.append(currentSwitch_);
            currentSwitch_ = SwitchInfo();
        }
        currentTag_.clear();
        return true;
    }

    bool characters(const QString& ch) override {
        if (currentTag_ == "Type") {
            currentSwitch_.model = ch.trimmed();
        }
        return true;
    }

private:
    struct SwitchInfo {
        QString id;
        QString type;
        QString model;
    };
    SwitchInfo currentSwitch_;
    QString currentTag_;
    QList<SwitchInfo> switches_;
};
```

### 6.4 DOM vs SAX 对比

| 特性 | DOM | SAX |
|------|-----|-----|
| 内存占用 | 高（全文档加载） | 低（流式处理） |
| 速度 | 慢（需构建树） | 快（边读边处理） |
| 随机访问 | 支持 | 不支持 |
| 修改文档 | 容易 | 困难 |
| 适用场景 | 小文档、需频繁访问 | 大文档、只需顺序读取 |
| 编程难度 | 简单 | 较复杂 |

---

## 七、 常见错误与注意事项

### 7.1 编码问题

**错误**：`QTextCodec::setCodecForLocale` 在某些 Qt 版本中已被移除

**正确做法**：
```cpp
// 设置流和文档的编码
QTextStream stream(&file);
stream.setCodec("UTF-8");
doc.save(stream, 4);
```

### 7.2 节点遍历的边界情况

```cpp
// 检查元素是否为空
QDomElement elem = root.firstChildElement("NonExistent");
if (elem.isNull()) {
    qDebug() << "Element not found";
}

// 检查节点是否存在
QDomNode node = switchList.at(i);
if (node.isNull()) {
    continue;
}
```

### 7.3 属性访问的安全写法

```cpp
// 安全的属性访问
QString id = swElem.attribute("id");
if (id.isEmpty()) {
    qWarning() << "Missing required attribute 'id'";
}

// 设置默认值
QString type = swElem.attribute("type", "default_type");
```

### 7.4 中文乱码问题

```cpp
// 确保文件以 UTF-8 编码保存
// 在 .pro 文件中添加
QT += xml

// 使用 QTextStream 显式设置编码
QTextStream stream(&file);
stream.setCodec("UTF-8");
```

---

## 八、 XPath 快速入门

### 8.1 XPath 基础

XPath 是一门用于在 XML 文档中查找信息的语言。

```cpp
// Qt 中使用 XPath（通过 QXmlQuery）
#include <QXmlQuery>

QXmlQuery query;
query.setDocument(&doc);
query.setQuery("descendant::Switch[@type='tsn_switch']/Type");

QString result;
query.evaluateTo(&result);
qDebug() << "Result:" << result;
```

### 8.2 常用路径表达式

| 表达式 | 说明 |
|--------|------|
| `nodeName` | 选取此节点的所有子节点 |
| `/` | 从根节点选取 |
| `//` | 从当前节点选取，不考虑位置 |
| `.` | 当前节点 |
| `..` | 父节点 |
| `@` | 选取属性 |
| `[@attr]` | 选取具有指定属性的元素 |
| `[tag='value']` | 选取具有指定值的元素 |

---

## 九、 面试高频问题与解答

### 9.1 基础问题

#### Q1：XML 的核心特点是什么？

**答案**：
- **纯文本格式**：平台无关，易于传输和存储
- **自描述性**：标签可自定义，语义清晰
- **树状结构**：具有严格的层级关系
- **严格语法**：必须格式良好，否则无法解析

#### Q2：XML 和 HTML 的区别是什么？

**答案**：
- **目的不同**：XML 用于数据存储传输，HTML 用于网页展示
- **标签可扩展性**：XML 标签可自定义，HTML 标签预定义
- **语法严格性**：XML 极其严格，HTML 相对宽松
- **大小写敏感**：XML 区分，HTML 不区分

#### Q3：什么是"格式良好"的 XML？

**答案**：
- 有且仅有一个根元素
- 所有标签必须正确闭合
- 标签不能交叉嵌套
- 属性值必须用引号包裹
- 特殊字符必须转义

### 9.2 解析相关问题

#### Q4：DOM 和 SAX 解析的区别是什么？

**答案**：
| 特性 | DOM | SAX |
|------|-----|-----|
| 内存占用 | 高 | 低 |
| 速度 | 慢 | 快 |
| 随机访问 | 支持 | 不支持 |
| 修改文档 | 容易 | 困难 |
| 适用场景 | 小文档 | 大文档 |

#### Q5：QDomNode 和 QDomElement 的区别是什么？

**答案**：
- `QDomNode` 是基类，代表任何类型的节点
- `QDomElement` 继承自 `QDomNode`，特指元素节点
- 万物皆节点：标签、属性、文本、注释都是节点
- 只有带 `< >` 标签的才是元素

#### Q6：如何正确提取元素的文本内容？

**答案**：
- 如果直接对包含多个子节点的元素调用 `.text()`，所有子节点文本会粘连
- 正确做法是逐个找到子元素，再调用 `.text()`
- 例如：`elem.firstChildElement("Type").text()`

### 9.3 应用相关问题

#### Q7：TSN 项目中 XML 的作用是什么？

**答案**：
- **数据与代码分离**：修改配置无需重新编译
- **配置驱动仿真**：C++ 软件读取 XML 配置来实例化网络
- **四大核心文件**：
  - `XX_topo.xml`：定义网络拓扑
  - `XX_msg.xml`：定义流量
  - `XX_para.xml`：定义参数
  - `XX_config.xml`：保存结果

#### Q8：什么是命名空间？为什么要用它？

**答案**：
- 命名空间用于避免不同来源 XML 文档的元素名冲突
- 使用 `xmlns:` 前缀来声明和区分
- 当混合使用多个 XML 词汇表时必需

---

## 十、 参考文献

1. **W3C XML 1.0 Specification** - https://www.w3.org/TR/xml/
2. **W3C XML Schema** - https://www.w3.org/XML/Schema
3. **Qt XML Patterns** - https://doc.qt.io/qt-5/qtxml-index.html
4. **XML Tutorial** - https://www.w3schools.com/xml/
5. **《项目建议书》** - TSN 仿真系统设计文档

---

*复习打卡：2026-05-13 (知识重建期 Day 7)*
*核心标签：#XML #Qt解析 #DOM #SAX #TSN配置 #XPath #文档验证*

希望这份学习笔记能帮助你更好地理解 XML 的原理与 Qt 解析实战！
