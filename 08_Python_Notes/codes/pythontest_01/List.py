r'''Python3 列表
序列是 Python 中最基本的数据结构。

序列中的每个值都有对应的位置值，称之为索引，第一个索引是 0，第二个索引是 1，依此类推。

Python 有 6 个序列的内置类型，但最常见的是列表和元组。

列表都可以进行的操作包括索引，切片，加，乘，检查成员。

此外，Python 已经内置确定序列的长度以及确定最大和最小的元素的方法。

列表是最常用的 Python 数据类型，它可以作为一个方括号内的逗号分隔值出现。

列表的数据项不需要具有相同的类型

创建一个列表，只要把逗号分隔的不同的数据项使用方括号括起来即可。如下所示：

list1 = ['Google', 'Runoob', 1997, 2000]
list2 = [1, 2, 3, 4, 5 ]
list3 = ["a", "b", "c", "d"]
list4 = ['red', 'green', 'blue', 'yellow', 'white', 'black']'''

r'''访问列表中的值
与字符串的索引一样，列表索引从 0 开始，第二个索引是 1，依此类推。

通过索引列表可以进行截取、组合等操作。
#!/usr/bin/python3

list = ['red', 'green', 'blue', 'yellow', 'white', 'black']
print( list[0] )
print( list[3] )
print( list[5] )
以上实例输出结果：
red
yellow
black
'''
r'''索引也可以从尾部开始，最后一个元素的索引为 -1，往前一位为 -2，以此类推。
实例
#!/usr/bin/python3

list = ['red', 'green', 'blue', 'yellow', 'white', 'black']
print( list[-1] )
print( list[-2] )
print( list[-3] )
以上实例输出结果：

black
white
yellow
'''
r'''使用下标索引来访问列表中的值，同样你也可以使用方括号 [] 的形式截取字符
实例
#!/usr/bin/python3

nums = [10, 20, 30, 40, 50, 60, 70, 80, 90]
print(nums[0:4])
以上实例输出结果：

[10, 20, 30, 40]
使用负数索引值截取：

实例
#!/usr/bin/python3

list = ['Google', 'Runoob', "Zhihu", "Taobao", "Wiki"]

# 读取第二位
print ("list[1]: ", list[1])
# 从第二位开始（包含）截取到倒数第二位（不包含）
print ("list[1:-2]: ", list[1:-2])
以上实例输出结果：

list[1]:  Runoob
list[1:-2]:  ['Runoob', 'Zhihu']'''

r'''更新列表
你可以对列表的数据项进行修改或更新，你也可以使用 append() 方法来添加列表项，如下所示：

实例(Python 3.0+)
#!/usr/bin/python3

list = ['Google', 'Runoob', 1997, 2000]

print ("第三个元素为 : ", list[2])
list[2] = 2001
print ("更新后的第三个元素为 : ", list[2])

list1 = ['Google', 'Runoob', 'Taobao']
list1.append('Baidu')
print ("更新后的列表 : ", list1)
注意：我们会在接下来的章节讨论 append() 方法的使用。

以上实例输出结果：

第三个元素为 :  1997
更新后的第三个元素为 :  2001
更新后的列表 :  ['Google', 'Runoob', 'Taobao', 'Baidu']'''

r'''删除列表元素
可以使用 del 语句来删除列表中的元素，如下实例：

实例(Python 3.0+)
#!/usr/bin/python3
 
list = ['Google', 'Runoob', 1997, 2000]
 
print ("原始列表 : ", list)
del list[2]
print ("删除第三个元素 : ", list)
以上实例输出结果：

原始列表 :  ['Google', 'Runoob', 1997, 2000]
删除第三个元素 :  ['Google', 'Runoob', 2000]
注意：我们会在接下来的章节讨论 remove() 方法的使用'''

r'''Python列表脚本操作符
列表对 + 和 * 的操作符与字符串相似。+ 号用于组合列表，* 号用于重复列表。

如下所示：

Python 表达式	结果	描述
len([1, 2, 3])	3	长度
[1, 2, 3] + [4, 5, 6]	[1, 2, 3, 4, 5, 6]	组合
['Hi!'] * 4	['Hi!', 'Hi!', 'Hi!', 'Hi!']	重复
3 in [1, 2, 3]	True	元素是否存在于列表中
for x in [1, 2, 3]: print(x, end=" ")	1 2 3	迭代
Python 列表截取与拼接
Python 的列表截取与字符串操作类似，如下所示：

L=['Google', 'Runoob', 'Taobao']
操作：

Python 表达式	结果	描述
L[2]	'Taobao'	读取第三个元素
L[-2]	'Runoob'	从右侧开始读取倒数第二个元素: count from the right
L[1:]	['Runoob', 'Taobao']	输出从第二个元素开始后的所有元素
>>> L=['Google', 'Runoob', 'Taobao']
>>> L[2]
'Taobao'
>>> L[-2]
'Runoob'
>>> L[1:]
['Runoob', 'Taobao']
>>>
列表还支持拼接操作：

>>> squares = [1, 4, 9, 16, 25]
>>> squares += [36, 49, 64, 81, 100]
>>> squares
[1, 4, 9, 16, 25, 36, 49, 64, 81, 100]
>>>
嵌套列表
使用嵌套列表即在列表里创建其它列表，例如：

>>> a = ['a', 'b', 'c']
>>> n = [1, 2, 3]
>>> x = [a, n]
>>> x
[['a', 'b', 'c'], [1, 2, 3]]
>>> x[0]
['a', 'b', 'c']
>>> x[0][1]
'b'/'''

r'''列表比较
列表比较需要引入 operator 模块的 eq 方法（详见：Python operator 模块）：

实例
# 导入 operator 模块
import operator

a = [1, 2]
b = [2, 3]
c = [2, 3]
print("operator.eq(a,b): ", operator.eq(a,b))
print("operator.eq(c,b): ", operator.eq(c,b))

以上代码输出结果为：

operator.eq(a,b):  False
operator.eq(c,b):  True'''
r'''Python3 operator 模块
Python2.x 版本中，使用 cmp() 函数来比较两个列表、数字或字符串等的大小关系。

Python 3.X 的版本中已经没有 cmp() 函数，如果你需要实现比较功能，需要引入 operator 模块，适合任何对象，包含的方法有：

operator 模块包含的方法
operator.lt(a, b)
operator.le(a, b)
operator.eq(a, b)
operator.ne(a, b)
operator.ge(a, b)
operator.gt(a, b)
operator.__lt__(a, b)
operator.__le__(a, b)
operator.__eq__(a, b)
operator.__ne__(a, b)
operator.__ge__(a, b)
operator.__gt__(a, b)
operator.lt(a, b) 与 a < b 相同， operator.le(a, b) 与 a <= b 相同，operator.eq(a, b) 与 a == b 相同，operator.ne(a, b) 与 a != b 相同，operator.gt(a, b) 与 a > b 相同，operator.ge(a, b) 与 a >= b 相同。

实例
# 导入 operator 模块
import operator
 
# 数字
x = 10
y = 20

print("x:",x, ", y:",y)
print("operator.lt(x,y): ", operator.lt(x,y))
print("operator.gt(y,x): ", operator.gt(y,x))
print("operator.eq(x,x): ", operator.eq(x,x))
print("operator.ne(y,y): ", operator.ne(y,y))
print("operator.le(x,y): ", operator.le(x,y))
print("operator.ge(y,x): ", operator.ge(y,x))
print()

# 字符串
x = "Google"
y = "Runoob"

print("x:",x, ", y:",y)
print("operator.lt(x,y): ", operator.lt(x,y))
print("operator.gt(y,x): ", operator.gt(y,x))
print("operator.eq(x,x): ", operator.eq(x,x))
print("operator.ne(y,y): ", operator.ne(y,y))
print("operator.le(x,y): ", operator.le(x,y))
print("operator.ge(y,x): ", operator.ge(y,x))
print()

# 查看返回值
print("type((operator.lt(x,y)): ", type(operator.lt(x,y)))
以上代码输出结果为：

x: 10 , y: 20
operator.lt(x,y):  True
operator.gt(y,x):  True
operator.eq(x,x):  True
operator.ne(y,y):  False
operator.le(x,y):  True
operator.ge(y,x):  True

x: Google , y: Runoob
operator.lt(x,y):  True
operator.gt(y,x):  True
operator.eq(x,x):  True
operator.ne(y,y):  False
operator.le(x,y):  True
operator.ge(y,x):  True
比较两个列表：

实例
# 导入 operator 模块
import operator

a = [1, 2]
b = [2, 3]
c = [2, 3]
print("operator.eq(a,b): ", operator.eq(a,b))
print("operator.eq(c,b): ", operator.eq(c,b))
以上代码输出结果为：

operator.eq(a,b):  False
operator.eq(c,b):  True
运算符函数
operator 模块提供了一套与 Python 的内置运算符对应的高效率函数。例如，operator.add(x, y) 与表达式 x+y 相同。

函数包含的种类有：对象的比较运算、逻辑运算、数学运算以及序列运算。

对象比较函数适用于所有的对象，函数名根据它们对应的比较运算符命名。

许多函数名与特殊方法名相同，只是没有双下划线。为了向后兼容性，也保留了许多包含双下划线的函数，为了表述清楚，建议使用没有双下划线的函数。

实例
# Python 实例
# add(), sub(), mul()
 
# 导入  operator 模块
import operator
 
# 初始化变量
a = 4
 
b = 3
 
# 使用 add() 让两个值相加
print ("add() 运算结果 :",end="");
print (operator.add(a, b))
 
# 使用 sub() 让两个值相减
print ("sub() 运算结果 :",end="");
print (operator.sub(a, b))
 
# 使用 mul() 让两个值相乘
print ("mul() 运算结果 :",end="");
print (operator.mul(a, b))
以上代码输出结果为：

add() 运算结果 :7
sub() 运算结果 :1
mul() 运算结果 :12
运算

语法

函数

加法

a + b

add(a, b)

字符串拼接

seq1 + seq2

concat(seq1, seq2)

包含测试

obj in seq

contains(seq, obj)

除法

a / b

truediv(a, b)

除法

a // b

floordiv(a, b)

按位与

a & b

and_(a, b)

按位异或

a ^ b

xor(a, b)

按位取反

~ a

invert(a)

按位或

a | b

or_(a, b)

取幂

a ** b

pow(a, b)

标识

a is b

is_(a, b)

标识

a is not b

is_not(a, b)

索引赋值

obj[k] = v

setitem(obj, k, v)

索引删除

del obj[k]

delitem(obj, k)

索引取值

obj[k]

getitem(obj, k)

左移

a << b

lshift(a, b)

取模

a % b

mod(a, b)

乘法

a * b

mul(a, b)

矩阵乘法

a @ b

matmul(a, b)

取反（算术）

- a

neg(a)

取反（逻辑）

not a

not_(a)

正数

+ a

pos(a)

右移

a >> b

rshift(a, b)

切片赋值

seq[i:j] = values

setitem(seq, slice(i, j), values)

切片删除

del seq[i:j]

delitem(seq, slice(i, j))

切片取值

seq[i:j]

getitem(seq, slice(i, j))

字符串格式化

s % obj

mod(s, obj)

减法

a - b

sub(a, b)

真值测试

obj

truth(obj)

比较

a < b

lt(a, b)

比较

a <= b

le(a, b)

相等

a == b

eq(a, b)

不等

a != b

ne(a, b)

比较

a >= b

ge(a, b)

比较

a > b

gt(a, b)'''

r'''Python列表函数&方法
Python包含以下函数:

序号	函数
1	len(list)
列表元素个数
2	max(list)
返回列表元素最大值
3	min(list)
返回列表元素最小值
4	list(seq)
将元组转换为列表
Python包含以下方法:

序号	方法
1	list.append(obj)
在列表末尾添加新的对象
2	list.count(obj)
统计某个元素在列表中出现的次数
3	list.extend(seq)
在列表末尾一次性追加另一个序列中的多个值（用新列表扩展原来的列表）
4	list.index(obj)
从列表中找出某个值第一个匹配项的索引位置
5	list.insert(index, obj)
将对象插入列表
6	list.pop([index=-1])
移除列表中的一个元素（默认最后一个元素），并且返回该元素的值
7	list.remove(obj)
移除列表中某个值的第一个匹配项
8	list.reverse()
反向列表中元素
9	list.sort( key=None, reverse=False)
对原列表进行排序
10	list.clear()
清空列表
11	list.copy()
复制列表'''