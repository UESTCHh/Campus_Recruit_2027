r'''Python3 字符串
字符串是 Python 中最常用的数据类型。我们可以使用引号( ' 或 " )来创建字符串。

创建字符串很简单，只要为变量分配一个值即可。例如：

var1 = 'Hello World!'
var2 = "Runoob"
Python 访问字符串中的值
Python 不支持单字符类型，单字符在 Python 中也是作为一个字符串使用。

Python 访问子字符串，可以使用方括号 [] 来截取字符串，字符串的截取的语法格式如下：

变量[头下标:尾下标]
索引值以 0 为开始值，-1 为从末尾的开始位置。
如下实例：'''

r'''var1 = 'Hello World!'
var2 = "Runoob"

print("var1[0]: ", var1[0])
print("var2[1:5]: ", var2[1:5])
以上实例执行结果：

var1[0]:  H
var2[1:5]:  
'''

r'''Python 字符串更新
你可以截取字符串的一部分并与其他字段拼接，如下实例：

实例(Python 3.0+)
#!/usr/bin/python3
 
var1 = 'Hello World!'
 
print ("已更新字符串 : ", var1[:6] + 'Runoob!')
以上实例执行结果

已更新字符串 :  Hello Runoob!
'''

r'''Python 转义字符
在需要在字符中使用特殊字符时，python 用反斜杠 \ 转义字符。如下表：

转义字符	描述	实例
\(在行尾时)	续行符	
>>> print("line1 \
... line2 \
... line3")
line1 line2 line3
>>> 
\\	反斜杠符号	
>>> print("\\")
\
\'	单引号	
>>> print('\'')
'
\"	双引号	
>>> print("\"")
"
\a	响铃	
>>> print("\a")
执行后电脑有响声。
\b	退格(Backspace)	
>>> print("Hello \b World!")
Hello World!
\000	空	
>>> print("\000")

>>> 
\n	换行	
>>> print("\n")


>>>
\v	纵向制表符	
>>> print("Hello \v World!")
Hello 
       World!
>>>
\t	横向制表符	
>>> print("Hello \t World!")
Hello      World!
>>>
\r	回车，将 \r 后面的内容移到字符串开头，并逐一替换开头部分的字符，直至将 \r 后面的内容完全替换完成。	
>>> print("Hello\rWorld!")
World!
>>> print('google runoob taobao\r123456')
123456 runoob taobao
\f	换页	
>>> print("Hello \f World!")
Hello 
       World!
>>> 
\yyy	八进制数，y 代表 0~7 的字符，例如：\012 代表换行。	
>>> print("\110\145\154\154\157\40\127\157\162\154\144\41")
Hello World!
\xyy	十六进制数，以 \x 开头，y 代表的字符，例如：\x0a 代表换行	
>>> print("\x48\x65\x6c\x6c\x6f\x20\x57\x6f\x72\x6c\x64\x21")
Hello World!
\other	其它的字符以普通格式输出	 
使用 \r 实现百分比进度：'''
r'''import time

for i in range(101): # 添加进度条图形和百分比
    bar = '[' + ('=' * (i // 2)) + (' ' * (50 - i // 2)) + ']'
    print(f"\r{bar} {i:3}%", end='', flush=True)
    time.sleep(0.05)
print()'''

r'''代码拆解
1. import time
导入 Python 标准库 time，主要用其中的 sleep 函数来制造延迟，让进度条动画可见。

2. for i in range(101):
range(101) 生成从 0 到 100 的整数序列，共 101 个数。
i 在这里直接充当百分比数值：i=0 表示 0%，i=100 表示 100%。
循环 101 次的原因就是要包含 0% 和 100% 两个端点，让进度条能从空跑到满。

3. 构建进度条字符串
python
bar = '[' + '=' * (i // 2) + ' ' * (50 - i // 2) + ']'
这一行动态生成进度条的外观。

i // 2
整除运算。因为 i 最大是 100，i // 2 的最大值是 50，所以整个进度条内部的可视长度被固定为 50 个字符。
当 i=0 时，i//2 = 0；当 i=100 时，i//2 = 50。

'=' * (i // 2)
字符串乘法，生成由 = 组成的字符串，长度为 i//2。
这一部分代表已完成进度。

' ' * (50 - i // 2)
生成由空格组成的字符串，长度为 50 - i//2。
这一部分代表剩余未完成进度，保证整条进度条内部始终占满 50 个字符。

'[' + ... + ']'
在两端加上方括号作为边框。
最终 bar 的总长度是 52 个字符（左括号 + 50个内部字符 + 右括号）。

举例：

i=0：[ + 50个空格 + ]

i=50：[ + 25个 = + 25个空格 + ]

i=100：[ + 50个 = + ]

4. 打印并覆盖当前行
python
print(f"\r{bar} {i:3}%", end='', flush=True)
\r (回车符)
将光标移动到当前行的最开头，但不会换到下一行。这样后面打印的内容就会覆盖掉这一行之前显示的内容。这是实现单行动态刷新的核心。

{bar} {i:3}%
f-string 格式化输出。
{bar} 放置前面构造的进度条字符串。
{i:3} 表示将整数 i 格式化为宽度至少为 3 个字符（右对齐）。
例如 i=0 显示 0%，i=100 显示 100%。这样百分比始终占 3 格，整个输出长度不会随数字位数变化而抖动。

end=''
print 默认会在结尾添加换行符 \n。这里将它改为空字符串，禁止自动换行，确保 \r 之后光标仍留在同一行。

flush=True
强制立即刷新输出缓冲区。
很多终端出于性能考虑会缓冲输出内容，等积累到一定量或遇到换行才真正显示。设置 flush=True 可以保证每次 print 的内容立刻出现在屏幕上，动画才不会卡顿或延迟。

5. time.sleep(0.05)
让程序暂停 0.05 秒（即 50 毫秒）。
这控制了进度条的更新速度：大约每秒更新 20 次（1 ÷ 0.05 = 20）。
去掉这行代码，整个循环会瞬间跑完，根本看不到动画过程。

6. 最后的 print()
循环结束后调用一个空的 print()，作用是输出一个换行符。
因为之前所有 print 都用了 end=''，循环结束时光标还停留在进度条那一行。
加这一行把光标移到下一行，防止后续程序的输出与进度条混在同一行。'''

r'''以下实例，我们使用了不同的转义字符来演示单引号、换行符、制表符、退格符、换页符、ASCII、二进制、八进制数和十六进制数的效果：

实例
print('\'Hello, world!\'')  # 输出：'Hello, world!'

print("Hello, world!\nHow are you?")  # 输出：Hello, world!
                                        #       How are you?

print("Hello, world!\tHow are you?")  # 输出：Hello, world!    How are you?

print("Hello,\b world!")  # 输出：Hello world!

print("Hello,\f world!")  # 输出：
                           # Hello,
                           #  world!

print("A 对应的 ASCII 值为：", ord('A'))  # 输出：A 对应的 ASCII 值为： 65

print("\x41 为 A 的 ASCII 代码")  # 输出：A 为 A 的 ASCII 代码

decimal_number = 42
binary_number = bin(decimal_number)  # 十进制转换为二进制
print('转换为二进制:', binary_number)  # 转换为二进制: 0b101010

octal_number = oct(decimal_number)  # 十进制转换为八进制
print('转换为八进制:', octal_number)  # 转换为八进制: 0o52

hexadecimal_number = hex(decimal_number)  # 十进制转换为十六进制
print('转换为十六进制:', hexadecimal_number) # 转换为十六进制: 0x2a'''

r'''Python 字符串运算符
下表实例变量 a 值为字符串 "Hello"，b 变量值为 "Python"：

操作符	描述	实例
+	字符串连接	a + b 输出结果： HelloPython
*	重复输出字符串	a*2 输出结果：HelloHello
[]	通过索引获取字符串中字符	a[1] 输出结果 e
[ : ]	截取字符串中的一部分，遵循左闭右开原则，str[0:2] 是不包含第 3 个字符的。	a[1:4] 输出结果 ell
in	成员运算符 - 如果字符串中包含给定的字符返回 True	'H' in a 输出结果 True
not in	成员运算符 - 如果字符串中不包含给定的字符返回 True	'M' not in a 输出结果 True
r/R	原始字符串 - 原始字符串：所有的字符串都是直接按照字面的意思来使用，没有转义特殊或不能打印的字符。 原始字符串除在字符串的第一个引号前加上字母 r（可以大小写）以外，与普通字符串有着几乎完全相同的语法。	
print( r'\n' )
print( R'\n' )
%	格式字符串	请看下一节内容。'''
r'''实例(Python 3.0+)
#!/usr/bin/python3
 
a = "Hello"
b = "Python"
 
print("a + b 输出结果：", a + b)
print("a * 2 输出结果：", a * 2)
print("a[1] 输出结果：", a[1])
print("a[1:4] 输出结果：", a[1:4])
 
if( "H" in a) :
    print("H 在变量 a 中")
else :
    print("H 不在变量 a 中")
 
if( "M" not in a) :
    print("M 不在变量 a 中")
else :
    print("M 在变量 a 中")
 
print (r'\n')
print (R'\n')
以上实例输出结果为：

a + b 输出结果： HelloPython
a * 2 输出结果： HelloHello
a[1] 输出结果： e
a[1:4] 输出结果： ell
H 在变量 a 中
M 不在变量 a 中
\n
\n'''

r'''Python 字符串格式化
Python 支持格式化字符串的输出 。尽管这样可能会用到非常复杂的表达式，但最基本的用法是将一个值插入到一个有字符串格式符 %s 的字符串中。

在 Python 中，字符串格式化使用与 C 中 sprintf 函数一样的语法。

实例(Python 3.0+)
#!/usr/bin/python3
 
print ("我叫 %s 今年 %d 岁!" % ('小明', 10))
以上实例输出结果：

我叫 小明 今年 10 岁!
python字符串格式化符号:

    符   号	描述
      %c	 格式化字符及其ASCII码
      %s	 格式化字符串
      %d	 格式化整数
      %u	 格式化无符号整型
      %o	 格式化无符号八进制数
      %x	 格式化无符号十六进制数
      %X	 格式化无符号十六进制数（大写）
      %f	 格式化浮点数字，可指定小数点后的精度
      %e	 用科学计数法格式化浮点数
      %E	 作用同%e，用科学计数法格式化浮点数
      %g	 %f和%e的简写
      %G	 %f 和 %E 的简写
      %p	 用十六进制数格式化变量的地址
格式化操作符辅助指令:

符号	功能
*	定义宽度或者小数点精度
-	用做左对齐
+	在正数前面显示加号( + )
<sp>	在正数前面显示空格
#	在八进制数前面显示零('0')，在十六进制前面显示'0x'或者'0X'(取决于用的是'x'还是'X')
0	显示的数字前面填充'0'而不是默认的空格
%	'%%'输出一个单一的'%'
(var)	映射变量(字典参数)
m.n.	m 是显示的最小总宽度,n 是小数点后的位数(如果可用的话)
Python2.6 开始，新增了一种格式化字符串的函数 str.format()，它增强了字符串格式化的功能。'''

r'''Python format 格式化函数
Python 字符串 Python 字符串

Python2.6 开始，新增了一种格式化字符串的函数 str.format()，它增强了字符串格式化的功能。

基本语法是通过 {} 和 : 来代替以前的 % 。

format 函数可以接受不限个参数，位置可以不按顺序。

实例
>>>"{} {}".format("hello", "world")    # 不设置指定位置，按默认顺序
'hello world'
 
>>> "{0} {1}".format("hello", "world")  # 设置指定位置
'hello world'
 
>>> "{1} {0} {1}".format("hello", "world")  # 设置指定位置
'world hello world'
也可以设置参数：

实例
#!/usr/bin/python
# -*- coding: UTF-8 -*-
 
print("网站名：{name}, 地址 {url}".format(name="菜鸟教程", url="www.runoob.com"))
 
# 通过字典设置参数
site = {"name": "菜鸟教程", "url": "www.runoob.com"}
print("网站名：{name}, 地址 {url}".format(**site))
 
# 通过列表索引设置参数
my_list = ['菜鸟教程', 'www.runoob.com']
print("网站名：{0[0]}, 地址 {0[1]}".format(my_list))  # "0" 是必须的
输出结果为：

网站名：菜鸟教程, 地址 www.runoob.com
网站名：菜鸟教程, 地址 www.runoob.com
网站名：菜鸟教程, 地址 www.runoob.com
也可以向 str.format() 传入对象：

实例
#!/usr/bin/python
# -*- coding: UTF-8 -*-
 
class AssignValue(object):
    def __init__(self, value):
        self.value = value
my_value = AssignValue(6)
print('value 为: {0.value}'.format(my_value))  # "0" 是可选的
输出结果为：

value 为: 6
数字格式化
下表展示了 str.format() 格式化数字的多种方法：

>>> print("{:.2f}".format(3.1415926))
3.14
数字	格式	输出	描述
3.1415926	{:.2f}	3.14	保留小数点后两位
3.1415926	{:+.2f}	+3.14	带符号保留小数点后两位
-1	{:-.2f}	-1.00	带符号保留小数点后两位
2.71828	{:.0f}	3	不带小数
5	{:0>2d}	05	数字补零 (填充左边, 宽度为2)
5	{:x<4d}	5xxx	数字补x (填充右边, 宽度为4)
10	{:x<4d}	10xx	数字补x (填充右边, 宽度为4)
1000000	{:,}	1,000,000	以逗号分隔的数字格式
0.25	{:.2%}	25.00%	百分比格式
1000000000	{:.2e}	1.00e+09	指数记法
13	{:>10d}	        13	右对齐 (默认, 宽度为10)
13	{:<10d}	13	左对齐 (宽度为10)
13	{:^10d}	    13	中间对齐 (宽度为10)
11	
'{:b}'.format(11)
'{:d}'.format(11)
'{:o}'.format(11)
'{:x}'.format(11)
'{:#x}'.format(11)
'{:#X}'.format(11)
1011
11
13
b
0xb
0XB
进制
^, <, > 分别是居中、左对齐、右对齐，后面带宽度， : 号后面带填充的字符，只能是一个字符，不指定则默认是用空格填充。

+ 表示在正数前显示 +，负数前显示 -；  （空格）表示在正数前加空格

b、d、o、x 分别是二进制、十进制、八进制、十六进制。

此外我们可以使用大括号 {} 来转义大括号，如下实例：

实例
#!/usr/bin/python
# -*- coding: UTF-8 -*-
 
print ("{} 对应的位置是 {{0}}".format("runoob"))
输出结果为：

runoob 对应的位置是 {0}'''

r'''Python 三引号
Python 中三引号可以将复杂的字符串进行赋值。

Python 三引号允许一个字符串跨多行，字符串中可以包含换行符、制表符以及其他特殊字符。

三引号的语法是一对连续的单引号或者双引号（通常都是成对的用）。

 >>> hi = \'''hi
there\'''
>>> hi   # repr()
'hi\nthere'
>>> print hi  # str()
hi 
there  
三引号让程序员从引号和特殊字符串的泥潭里面解脱出来，自始至终保持一小块字符串的格式是所谓的WYSIWYG（所见即所得）格式的。

一个典型的用例是，当你需要一块HTML或者SQL时，这时当用三引号标记，使用传统的转义字符体系将十分费神。

 errHTML = \'''
<HTML><HEAD><TITLE>
Friends CGI Demo</TITLE></HEAD>
<BODY><H3>ERROR</H3>
<B>%s</B><P>
<FORM><INPUT TYPE=button VALUE=Back
ONCLICK="window.history.back()"></FORM>
</BODY></HTML>
\'''
cursor.execute(\'''
CREATE TABLE users (  
login VARCHAR(8), 
uid INTEGER,
prid INTEGER)
\''')
'''

r'''Unicode 字符串
Python 中定义一个 Unicode 字符串和定义一个普通字符串一样简单：

>>> u'Hello World !'
u'Hello World !'
引号前小写的"u"表示这里创建的是一个 Unicode 字符串。如果你想加入一个特殊字符，可以使用 Python 的 Unicode-Escape 编码。如下例所示：

>>> u'Hello\u0020World !'
u'Hello World !'
被替换的 \u0020 标识表示在给定位置插入编码值为 0x0020 的 Unicode 字符（空格符）。'''

r'''Python 的字符串内建函数
字符串方法是从 Python1.6 到 2.0 慢慢加进来的 —— 它们也被加到了Jython 中。

这些方法实现了 string 模块的大部分方法，如下表所示列出了目前字符串内建支持的方法，所有的方法都包含了对 Unicode 的支持，有一些甚至是专门用于 Unicode 的。

方法	描述
string.capitalize()

把字符串的第一个字符大写

string.center(width)

返回一个原字符串居中,并使用空格填充至长度 width 的新字符串

string.count(str, beg=0, end=len(string))

返回 str 在 string 里面出现的次数，如果 beg 或者 end 指定则返回指定范围内 str 出现的次数

string.decode(encoding='UTF-8', errors='strict')

以 encoding 指定的编码格式解码 string，如果出错默认报一个 ValueError 的 异 常 ， 除非 errors 指 定 的 是 'ignore' 或 者'replace'

string.encode(encoding='UTF-8', errors='strict')

以 encoding 指定的编码格式编码 string，如果出错默认报一个ValueError 的异常，除非 errors 指定的是'ignore'或者'replace'

string.endswith(obj, beg=0, end=len(string))

检查字符串是否以 obj 结束，如果beg 或者 end 指定则检查指定的范围内是否以 obj 结束，如果是，返回 True,否则返回 False.

string.expandtabs(tabsize=8)

把字符串 string 中的 tab 符号转为空格，tab 符号默认的空格数是 8。

string.find(str, beg=0, end=len(string))

检测 str 是否包含在 string 中，如果 beg 和 end 指定范围，则检查是否包含在指定范围内，如果是返回开始的索引值，否则返回-1

string.format()

格式化字符串

string.index(str, beg=0, end=len(string))

跟find()方法一样，只不过如果str不在 string中会报一个异常.

string.isalnum()

如果 string 至少有一个字符并且所有字符都是字母或数字则返

回 True,否则返回 False

string.isalpha()

如果 string 至少有一个字符并且所有字符都是字母则返回 True,

否则返回 False

string.isdecimal()

如果 string 只包含十进制数字则返回 True 否则返回 False.

string.isdigit()

如果 string 只包含数字则返回 True 否则返回 False.

string.islower()

如果 string 中包含至少一个区分大小写的字符，并且所有这些(区分大小写的)字符都是小写，则返回 True，否则返回 False

string.isnumeric()

如果 string 中只包含数字字符，则返回 True，否则返回 False

string.isspace()

如果 string 中只包含空格，则返回 True，否则返回 False.

string.istitle()

如果 string 是标题化的(见 title())则返回 True，否则返回 False

string.isupper()

如果 string 中包含至少一个区分大小写的字符，并且所有这些(区分大小写的)字符都是大写，则返回 True，否则返回 False

string.join(seq)

以 string 作为分隔符，将 seq 中所有的元素(的字符串表示)合并为一个新的字符串

string.ljust(width)

返回一个原字符串左对齐,并使用空格填充至长度 width 的新字符串

string.lower()

转换 string 中所有大写字符为小写.

string.lstrip()

截掉 string 左边的空格

string.maketrans(intab, outtab)

maketrans() 方法用于创建字符映射的转换表，对于接受两个参数的最简单的调用方式，第一个参数是字符串，表示需要转换的字符，第二个参数也是字符串表示转换的目标。

max(str)

返回字符串 str 中最大的字母。

min(str)

返回字符串 str 中最小的字母。

string.partition(str)

有点像 find()和 split()的结合体,从 str 出现的第一个位置起,把 字 符 串 string 分 成 一 个 3 元 素 的 元 组 (string_pre_str,str,string_post_str),如果 string 中不包含str 则 string_pre_str == string.

string.replace(str1, str2,  num=string.count(str1))

把 string 中的 str1 替换成 str2,如果 num 指定，则替换不超过 num 次.

string.rfind(str, beg=0,end=len(string) )

类似于 find() 函数，返回字符串最后一次出现的位置，如果没有匹配项则返回 -1。

string.rindex( str, beg=0,end=len(string))

类似于 index()，不过是返回最后一个匹配到的子字符串的索引号。

string.rjust(width)

返回一个原字符串右对齐,并使用空格填充至长度 width 的新字符串

string.rpartition(str)

类似于 partition()函数,不过是从右边开始查找

string.rstrip()

删除 string 字符串末尾的空格.

string.split(str="", num=string.count(str))

以 str 为分隔符切片 string，如果 num 有指定值，则仅分隔 num+1 个子字符串

string.splitlines([keepends])

按照行('\r', '\r\n', '\n')分隔，返回一个包含各行作为元素的列表，如果参数 keepends 为 False，不包含换行符，如果为 True，则保留换行符。

string.startswith(obj, beg=0,end=len(string))

检查字符串是否是以 obj 开头，是则返回 True，否则返回 False。如果beg 和 end 指定值，则在指定范围内检查.

string.strip([obj])

在 string 上执行 lstrip()和 rstrip()

string.swapcase()

翻转 string 中的大小写

string.title()

返回"标题化"的 string,就是说所有单词都是以大写开始，其余字母均为小写(见 istitle())

string.translate(str, del="")

根据 str 给出的表(包含 256 个字符)转换 string 的字符,

要过滤掉的字符放到 del 参数中

string.upper()

转换 string 中的小写字母为大写

string.zfill(width)

返回长度为 width 的字符串，原字符串 string 右对齐，前面填充0'''
