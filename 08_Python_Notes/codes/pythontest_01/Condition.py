# Python3 条件控制
# Python 条件语句是通过一条或多条语句的执行结果（True 或者 False）来决定执行的代码块。
#
# 可以通过下图来简单了解条件语句的执行过程:
#
#
#
# 代码执行过程：
#
#
#
# 条件判断关键字
# 关键字 / 函数	说明	示例
# if	条件判断语句，当条件为 True 时执行代码块	if x > 0:
# elif	多条件判断分支（else if）	elif x == 0:
# else	所有条件不满足时执行	else:
# pass	空语句，占位用，保证语法完整	if x > 0: pass
# match	结构化模式匹配（Python 3.10+，类似 switch）	match x: case 1: ...
# if 语句
# Python中if语句的一般形式如下所示：
#
# if condition_1:
#     statement_block_1
# elif condition_2:
#     statement_block_2
# else:
#     statement_block_3
# 如果 "condition_1" 为 True 将执行 "statement_block_1" 块语句
# 如果 "condition_1" 为False，将判断 "condition_2"
# 如果"condition_2" 为 True 将执行 "statement_block_2" 块语句
# 如果 "condition_2" 为False，将执行"statement_block_3"块语句
# Python 中用 elif 代替了 else if，所以if语句的关键字为：if – elif – else。
#
# 注意：
#
# 1、每个条件后面要使用冒号 :，表示接下来是满足条件后要执行的语句块。
# 2、使用缩进来划分语句块，相同缩进数的语句在一起组成一个语句块。
# 3、在 Python 中没有 switch...case 语句，但在 Python3.10 版本添加了 match...case，功能也类似，详见下文。

# 实例
# 以下是一个简单的 if 实例：
#
# 实例
# # !/usr/bin/python3
#
# var1 = 100
# if var1:
#     print("1 - if 表达式条件为 true")
#     print(var1)
#
# var2 = 0
# if var2:
#     print("2 - if 表达式条件为 true")
#     print(var2)
# print("Good bye!")
# 执行以上代码，输出结果为：
#
# 1 - if 表达式条件为
# true
# 100
# Good
# bye!
# 从结果可以看到由于变量
# var2
# 为
# 0，所以对应的 if 内的语句没有执行。
#
# 以下实例演示了狗的年龄计算判断：
#
# 实例
# # !/usr/bin/python3
#
# age = int(input("请输入你家狗狗的年龄: "))
# print("")
# if age <= 0:
#     print("你是在逗我吧!")
# elif age == 1:
#     print("相当于 14 岁的人。")
# elif age == 2:
#     print("相当于 22 岁的人。")
# elif age > 2:
#     human = 22 + (age - 2) * 5
#     print("对应人类年龄: ", human)
#
# ### 退出提示
# input("点击 enter 键退出")
# 将以上脚本保存在dog.py文件中，并执行该脚本：
#
# $ python3
# dog.py
# 请输入你家狗狗的年龄: 1
#
# 相当于
# 14
# 岁的人。
# 点击
# enter
# 键退出
# 以下为if中常用的操作运算符:
#
# 操作符
# 描述
# < 小于
# <= 小于或等于
# > 大于
# >= 大于或等于
# == 等于，比较两个值是否相等
# != 不等于
# 实例
# # !/usr/bin/python3
#
# # 程序演示了 == 操作符
# # 使用数字
# print(5 == 6)
# # 使用变量
# x = 5
# y = 8
# print(x == y)
# 以上实例输出结果：
#
# False
# False
# high_low.py文件演示了数字的比较运算：
#
# 实例
# # !/usr/bin/python3
#
# # 该实例演示了数字猜谜游戏
# number = 7
# guess = -1
# print("数字猜谜游戏!")
# while guess != number:
#     guess = int(input("请输入你猜的数字："))
#
#     if guess == number:
#         print("恭喜，你猜对了！")
#     elif guess < number:
#         print("猜的数字小了...")
#     elif guess > number:
#         print("猜的数字大了...")
# 执行以上脚本，实例输出结果如下：
#
# $ python3
# high_low.py
# 数字猜谜游戏!
# 请输入你猜的数字：1
# 猜的数字小了...
# 请输入你猜的数字：9
# 猜的数字大了...
# 请输入你猜的数字：7
# 恭喜，你猜对了！

# if 嵌套
#     在嵌套 if 语句中，可以把 if ... elif ... else 结构放在另外一个 if ... elif ... else 结构中。
#
#     if 表达式1:
#         语句
#         if 表达式2:
#             语句
#         elif 表达式3:
#             语句
#         else:
#             语句
#     elif 表达式4:
#         语句
#     else:
#         语句
#     实例
#     # !/usr/bin/python3
#
#     num = int(input("输入一个数字："))
#     if num % 2 == 0:
#         if num % 3 == 0:
#             print("你输入的数字可以整除 2 和 3")
#         else:
#             print("你输入的数字可以整除 2，但不能整除 3")
#     else:
#         if num % 3 == 0:
#             print("你输入的数字可以整除 3，但不能整除 2")
#         else:
#             print("你输入的数字不能整除 2 和 3")
#     将以上程序保存到
#     test_if.py
#     文件中，执行后输出结果为：
#
#     $ python3
#     test.py
#     输入一个数字：6
#     你输入的数字可以整除
#     2
#     和
#

# match...case
# Python
# 3.10
# 增加了
# match...case
# 的条件判断，不需要再使用一连串的 if - else 来判断了。
#
# match
# 后的对象会依次与
# case
# 后的内容进行匹配，如果匹配成功，则执行匹配到的表达式，否则直接跳过，_
# 可以匹配一切。
#
# match
# subject
# case < pattern_1 > ?
# 匹配
# action_1
# 不匹配
# case < pattern_2 > ?
# 匹配
# action_2
# 不匹配
# case < pattern_3 > ?
# 匹配
# action_3
# 不匹配
# case
# _（通配符）
# 必匹配
# wildcard
# 语法格式如下：
#
# match subject:
#     case < pattern_1 >:
#         < action_1 >
#     case < pattern_2 >:
#         < action_2 >
#     case < pattern_3 >:
#         < action_3 >
#     case _:
#         < action_wildcard >
# case
# _: 类似于
# C
# 和
# Java
# 中的
# default:，当其他
# case
# 都无法匹配时，匹配这条，保证永远会匹配成功。
#
# 实例
#
#
# def http_error(status):
#     match status:
#         case 400:
#             return "Bad request"
#         case 404:
#             return "Not found"
#         case 418:
#             return "I'm a teapot"
#         case _:
#             return "Something's wrong with the internet"
#
#
# print(http_error(400))
# print(http_error(404))
# print(http_error(418))
# print(http_error(500))
# 以上是一个输出
# HTTP
# 状态码的实例，多个状态码的输出结果为：
#
# Bad request
# Not found
# I'm a teapot
# Something's wrong with the internet
# 一个case也可以设置多个匹配条件，条件使用 | 隔开，例如：
#
# ...
# case
# 401 | 403 | 404:
# return "Not allowed"
# status
# 401
# 403
# 404
# |
# |
# "Not allowed"
# 实例
#
#
# def check_permission(status):
#     match status:
#         case 200:
#             return "OK - 请求成功"
#         case 301 | 302:
#             return "Redirect - 重定向"
#         case 401 | 403 | 404:
#             return "Not allowed - 无权限或未找到"
#         case 500 | 502 | 503:
#             return "Server Error - 服务器错误"
#         case _:
#             return "Unknown status - 未知状态码"
#
#
# for code in [200, 301, 403, 500, 418]:
#     print(f"状态码 {code}: {check_permission(code)}")
#
# 状态码 200: OK - 请求成功
# 状态码 301: Redirect - 重定向
# 状态码 403: Not allowed - 无权限或未找到
# 状态码 500: Server Error - 服务器错误
# 状态码 418: Unknown status - 未知状态码

# Python match...case 语句
# Python3 条件控制 Python3 条件控制
#
# match...case 提供了一种更强大的模式匹配方法。
#
# 模式匹配是一种在编程中处理数据结构的方式，可以使代码更简洁、易读。
#
# match...case 是 Python 3.10 版本引入的新语法。
#
# match...case 语法结构如下：
#
# match expression:
#     case pattern1:
#         # 处理pattern1的逻辑
#     case pattern2 if condition:
#         # 处理pattern2并且满足condition的逻辑
#     case _:
#         # 处理其他情况的逻辑
# 参数说明：
#
# match语句后跟一个表达式，然后使用case语句来定义不同的模式。
# case后跟一个模式，可以是具体值、变量、通配符等。
# 可以使用if关键字在case中添加条件。
# _通常用作通配符，匹配任何值。

# 实例
# 1. 简单的值匹配
#
# 实例
# def match_example(value):
#     match value:
#         case 1:
#             print("匹配到值为1")
#         case 2:
#             print("匹配到值为2")
#         case _:
#             print("匹配到其他值")
#
# match_example(1)  # 输出: 匹配到值为1
# match_example(2)  # 输出: 匹配到值为2
# match_example(3)  # 输出: 匹配到其他值
# 以上代码中，match 语句用于匹配 value 的不同情况，每个 case 语句表示一种可能的匹配情况，_ 通配符表示其他情况。
#
# 输出结果为：
#
# 匹配到值为1
# 匹配到值为2
# 匹配到其他值
# 2. 使用变量
#
# 实例
# def match_example(item):
#     match item:
#         case (x, y) if x == y:
#             print(f"匹配到相等的元组: {item}")
#         case (x, y):
#             print(f"匹配到元组: {item}")
#         case _:
#             print("匹配到其他情况")
#
# match_example((1, 1))  # 输出: 匹配到相等的元组: (1, 1)
# match_example((1, 2))  # 输出: 匹配到元组: (1, 2)
# match_example("other") # 输出: 匹配到其他情况
# 输出结果为：
#
# 匹配到相等的元组: (1, 1)
# 匹配到元组: (1, 2)
# 匹配到其他情况
# 3. 类型匹配
#
# 实例
# class Circle:
#     def __init__(self, radius):
#         self.radius = radius
#
# class Rectangle:
#     def __init__(self, width, height):
#         self.width = width
#         self.height = height
#
# def match_shape(shape):
#     match shape:
#         case Circle(radius=1):
#             print("匹配到半径为1的圆")
#         case Rectangle(width=1, height=2):
#             print("匹配到宽度为1，高度为2的矩形")
#         case _:
#             print("匹配到其他形状")
#
# match_shape(Circle(radius=1))          # 输出: 匹配到半径为1的圆
# match_shape(Rectangle(width=1, height=2)) # 输出: 匹配到宽度为1，高度为2的矩形
# match_shape(Circle(radius=2))          # 输出: 匹配到其他形状
# match_shape(Rectangle(width=2, height=2)) # 输出: 匹配到其他形状
# match_shape("other")                    # 输出: 匹配到其他形状
# 输出结果为：
#
# 匹配到半径为1的圆
# 匹配到宽度为1，高度为2的矩形
# 匹配到其他形状
# 匹配到其他形状
# 匹配到其他形状