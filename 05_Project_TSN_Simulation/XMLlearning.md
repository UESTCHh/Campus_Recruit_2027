# 📚 学习笔记：XML 基础与 Qt 解析实战（TSN 仿真项目专项）

## 一、 XML 核心基础（我们只需要懂这些）

XML 的本质是一种**用来存储和传输数据的纯文本格式**，它通过自定义的标签将数据组织成具有层级关系的**树状结构**。针对本项目，掌握以下 5 点即可：

1.  **声明 (Declaration)：** 必须在文件的第一行、第一列。
    `<?xml version="1.0" encoding="UTF-8"?>`
2.  **根节点 (Root Element)：** 每个 XML 文件必须且只能有一个最外层的标签（例如 `<Topology>`）。
3.  **元素 (Element)：** 由成对的尖括号组成，必须正确闭合。例如 `<Switch>...</Switch>`。
4.  **属性 (Attribute)：** 写在起始标签内部的附加信息，常用于简短的唯一标识。例如 `<Switch id="SW_1">`。
5.  **注释 (Comment)：** 留给开发者看的备注。``。

## 二、 Qt 解析 XML 的核心武器（DOM 模型）

在 `.pro` 文件中添加 `QT += xml` 后，主要使用以下三个类在内存中操作这棵“树”：

* **`QDomDocument`**：代表整棵树（土壤和树根），负责把本地文件加载进内存。
* **`QDomElement`**：代表一个具体的标签元素（如 `<Switch>`）。
* **`QDomNodeList`**：代表一组搜刮出来的节点集合。

### 🚨 核心防坑指南：节点 (Node) vs 元素 (Element)
* **万物皆节点：** 标签、属性、内部的纯文本、注释，统统都是节点。
* **元素只是节点的一种：** 只有带 `< >` 标签的节点才是元素（`QDomElement` 继承自 `QDomNode`）。
* **提取文本的陷阱：** 如果 `<type>TSN_Switch</type>`，想要提取 `TSN_Switch`，必须找到 `<type>` 这个元素节点，再调用它的 `.text()` 方法。**绝对不能对包含多个子节点的父节点直接调用 `.text()`**，否则所有子节点的文本会粘连在一起。

### 🛠️ 黄金解析代码模板

```cpp
// 1. 加载文件
QDomDocument doc;
QFile file("path/to/your/file.xml");
file.open(QIODevice::ReadOnly | QIODevice::Text);
doc.setContent(&file);
file.close();

// 2. 获取根节点
QDomElement root = doc.documentElement();

// 3. 广度搜索：找出所有的 <Switch> (支持深度递归查找)
QDomNodeList switchList = root.elementsByTagName("Switch");

// 4. 遍历与数据提取
for (int i = 0; i < switchList.count(); ++i) {
    QDomNode node = switchList.at(i);
    if (node.isElement()) {
        QDomElement swElem = node.toElement();
        
        // 提取属性
        QString id = swElem.attribute("id"); 
        
        // 提取子元素内部的纯文本
        QString type = swElem.firstChildElement("type").text(); 
    }
}
```

## 三、 TSN 仿真项目 XML 架构蓝图

根据《项目建议书》的要求，我们的 C++ 软件不再需要硬编码数据，而是通过读取以下 4 个 XML 配置文件来驱动：

1.  **`XX_topo.xml`（拓扑文件）：**
    * **作用：** 定义网络中的硬件实体与连接。
    * **核心标签设计：** `<Topology>` 下包含 `<Nodes>`（分 `<Switch>` 和 `<EndSystem>`）与 `<Links>`。
2.  **`XX_msg.xml`（消息流文件）：**
    * **作用：** 定义网络中传输的流量（最高 2048 条）。
    * **核心标签设计：** `<TrafficFlows>` 下包含多个 `<Flow type="ST" />` 或 `<Flow type="AVB" />`，记录周期、源、目的和优先级。
3.  **`XX_para.xml`（配置参数文件）：**
    * **作用：** 调度算法的全局规则与环境。
    * **核心标签设计：** `<ConfigurationParameters>` 下记录时间同步精度、基础周期，以及流量机制映射表（如 ST -> CQF，AVB -> CBS）。
4.  **`XX_config.xml`（配置结果文件）：**
    * **作用：** C++ 调度算法运行完毕后，**写入保存**的结果。
    * **核心标签设计：** `<ScheduleResult>` 下记录总延迟，以及具体下发给设备的 `<GateControlList>`（GCL 门控列表状态）。

## 四、 下阶段开发路线（告别 XML，回归 C++）

XML 的学习到此为止。接下来的开发重心转移至：

1.  **构建数据载体：** 在 C++ 中编写与 XML 结构对应的 `struct` 或 `class`（如 `TsnNode`、`TsnLink`、`TsnFlow`）。
2.  **联调 Qt 界面：** 实现“点击按钮 -> 弹出文件选择框 -> 解析 XML -> 将数据存入 C++ 列表 (QList)” 的闭环。
3.  **图形化拓扑：** 将 C++ 列表中的节点与链路数据，传递给 Qt 的绘图模块，在界面上画出真实的拓扑图。
4.  **搁置底层仿真：** 在 Qt 前端的这 4 个 XML 文件能完美生成与读取之前，**暂时不要去碰 OMNeT++**。

--- 
*“XML 只是负责把快递送到家门口，如何拆开并组装里面的零件，才是 C++ 程序员的真正战场。”*