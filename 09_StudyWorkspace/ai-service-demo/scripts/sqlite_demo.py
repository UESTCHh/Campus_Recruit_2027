# 文件作用
# 使用 Python 原生 sqlite3 模块完成最小数据库实验，包括：
# 创建数据库
# 创建表
# 写入消息
# 读取消息
# 清空实验数据
# 为什么新增这个文件
# 文件作用：
# 使用 Python 标准库 sqlite3 演示创建表、写入、读取和清空持久化数据
# 先直接观察 SQLite 的基础行为，再学习 SQLAlchemy。
# 这样后续看到 ORM、Engine 和 Session 时，能够理解它们在替我们封装什么。
import argparse
import sqlite3
from pathlib import Path

DATABASE_PATH = Path("data/sqlite_demo.db")


def connect_database() -> sqlite3.Connection:
    DATABASE_PATH.parent.mkdir(
        parents=True,
        exist_ok=True,
    )

    connection = sqlite3.connect(DATABASE_PATH)
    connection.row_factory = sqlite3.Row

    return connection


def create_table(connection: sqlite3.Connection) -> None:
    connection.execute(
        """
        CREATE TABLE IF NOT EXISTS messages (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            conversation_id TEXT NOT NULL,
            role TEXT NOT NULL,
            content TEXT NOT NULL,
            created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
        )
        """
    )
    connection.commit()


def write_message(
    connection: sqlite3.Connection,
    conversation_id: str,
    role: str,
    content: str,
) -> None:
    cursor = connection.execute(
        """
        INSERT INTO messages (
            conversation_id,
            role,
            content
        )
        VALUES (?, ?, ?)
        """,
        (
            conversation_id,
            role,
            content,
        ),
    )
    connection.commit()

    print(f"saved message id: {cursor.lastrowid}")


def read_messages(connection: sqlite3.Connection) -> None:
    cursor = connection.execute(
        """
        SELECT
            id,
            conversation_id,
            role,
            content,
            created_at
        FROM messages
        ORDER BY id
        """
    )

    rows = cursor.fetchall()

    if not rows:
        print("no messages found")
        return

    for row in rows:
        print(
            " | ".join(
                [
                    f"id={row['id']}",
                    f"conversation_id={row['conversation_id']}",
                    f"role={row['role']}",
                    f"content={row['content']}",
                    f"created_at={row['created_at']}",
                ]
            )
        )


def reset_database(connection: sqlite3.Connection) -> None:
    connection.execute("DELETE FROM messages")
    connection.commit()

    print("all demo messages deleted")


def parse_arguments() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Minimal SQLite persistence experiment.",
    )

    subparsers = parser.add_subparsers(
        dest="command",
        required=True,
    )

    write_parser = subparsers.add_parser(
        "write",
        help="Write one message.",
    )
    write_parser.add_argument("conversation_id")
    write_parser.add_argument("role")
    write_parser.add_argument("content")

    subparsers.add_parser(
        "read",
        help="Read all messages.",
    )

    subparsers.add_parser(
        "reset",
        help="Delete all messages.",
    )

    return parser.parse_args()


def main() -> None:
    arguments = parse_arguments()

    with connect_database() as connection:
        create_table(connection)

        if arguments.command == "write":
            write_message(
                connection=connection,
                conversation_id=arguments.conversation_id,
                role=arguments.role,
                content=arguments.content,
            )
            return

        if arguments.command == "read":
            read_messages(connection)
            return

        if arguments.command == "reset":
            reset_database(connection)


if __name__ == "__main__":
    main()
