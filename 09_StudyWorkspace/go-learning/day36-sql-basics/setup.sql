-- 创建学生表并添加测试数据
-- 创建学生表
CREATE TABLE students (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    age INTEGER NOT NULL
);

-- 添加三条学生记录
INSERT INTO students (id, name, age)
VALUES
    (1001, 'Alice', 20),
    (1002, 'Bob', 21),
    (1003, 'Tom', 22);