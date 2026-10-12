from pathlib import Path
import sqlite3

base_dir = Path(__file__).resolve().parent
db_path = base_dir / "students.db"

# 确认数据库文件已经初始化
if not db_path.exists():
    raise SystemExit(
        "Database not found. Run init_db.py first."
    )

# 打开已经存在的数据库文件
conn = sqlite3.connect(db_path)

try:
    # 查询所有学生，按id升序排列
    rows = conn.execute(
        "SELECT id, name, age FROM students ORDER BY id;"
    ).fetchall()

    print("Students from database file:")

    for row in rows:
        print(*row, sep=" | ")

finally:
    conn.close()