-- 编写今天要学习的 SQL 查询
-- 查询年龄大于等于20岁的学生
-- 按年龄从大到小排列
-- 只返回前2条

SELECT name, age
FROM students
WHERE age >= 20
ORDER BY age DESC
LIMIT 2;