# 执行 SQL 并打印新增后的表
from pathlib import Path
import sqlite3

# 当前脚本所在目录
base_dir = Path(__file__).resolve().parent

# 创建内存SQLite数据库
conn = sqlite3.connect(":memory:")

try:
    # 1.创建表并插入原始数据
    setup_sql = (base_dir / "setup.sql").read_text(
        encoding="utf-8"
    )
    conn.executescript(setup_sql)

    print("Before INSERT:")

    rows = conn.execute(
        "SELECT id, name, age FROM students ORDER BY id;"
    ).fetchall()

    for row in rows:
        print(*row, sep = " | ")

    # 2.执行本次新增数据的sql
    insert_sql = (base_dir / "insert.sql").read_text(
        encoding="utf-8"
    )
    conn.executescript(insert_sql)

    print("\nAfter INSERT:")

    # 3.查询新增后的全部学生
    rows = conn.execute(
        "SELECT id, name, age FROM students ORDER BY id;"
    ).fetchall()

    for row in rows:
        print(*row, sep = " | ")

    # 4.更新用户信息
    update_sql = (base_dir / "update.sql").read_text(
    encoding="utf-8"
    )

    conn.executescript(update_sql)

    print("\nAfter UPDATE:")

    rows = conn.execute(
        "SELECT id, name, age FROM students ORDER BY id;"
    ).fetchall()

    for row in rows:
        print(*row, sep=" | ")
    
    # 5.删除用户
    delete_sql = (base_dir / "delete.sql").read_text(
        encoding="utf-8"
    )

    conn.executescript(delete_sql)

    print("\nAfter DELETE:")

    rows = conn.execute(
        "SELECT id, name, age FROM students ORDER BY id;"
    ).fetchall()

    for row in rows:
        print(*row, sep=" | ")

finally:
    #关闭数据库连接
    conn.close()