from pathlib import Path
import sqlite3

# 当前脚本所在目录
base_dir = Path(__file__).resolve().parent

# 数据库文件的完整路径
# 它使数据库文件保存在当前实验目录下。
db_path = base_dir / "students.db"

# 如果数据库文件已经存在，就不重复初始化
# 避免覆盖之前保存的数据
if db_path.exists():
    print("Database already exists. Initialization skipped.")
else:
    # 连接到SQLite文件数据库
    # sqlite3.connect() 在指定文件不存在时通常会自动创建数据库文件
    # 但这并不意味着会自动创建 students 表；表仍然要通过 SQL 创建。
    conn = sqlite3.connect(db_path)

    try:
        # 读取建表和插入数据的SQL
        setup_sql = (base_dir / "setup.sql").read_text(
            encoding="utf-8"
        )

        # 执行SQL
        conn.executescript(setup_sql)
        conn.commit()

        print("Database created: students.db")

    finally:
        # 关闭数据库连接
        conn.close()