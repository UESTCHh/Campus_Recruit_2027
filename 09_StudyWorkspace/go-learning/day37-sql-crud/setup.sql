-- 创建 students 表并添加原始测试数据
-- 创建学生表
CREATE TABLE students(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    age INTEGET NOT NULL
);

-- 添加原始的三条学生记录
INSERT INTO students(id, name, age)
VALUES
    (1001, 'Alice', 20),
    (1002, 'Bob', 21),
    (1003, 'Tom', 22);