# 使用 SQLite 执行 SQL，显示查询结果
from pathlib import Path
import sqlite3

# 创建一个临时的内存数据库，用于SQL练习
with sqlite3.connect(":memory:") as conn:

    # 读取并执行建表和插入数据的SQL
    setup_sql = Path("setup.sql").read_text(encoding="utf-8")
    conn.executescript(setup_sql)

    # 读取查询SQL
    query_sql = Path("query.sql").read_text(encoding="utf-8")

    # 执行查询，并获取结果
    rows = conn.execute(query_sql).fetchall()

    # 显示查询结果
    for row in rows:
        print(*row, sep=" | ")