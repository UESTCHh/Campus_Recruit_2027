package main

import "database/sql"

// Student表示一名学生
type Student struct {
	ID   int
	Name string
	Age  int
}

// StudentRepository负责访问学生数据
type StudentRepository struct {
	db *sql.DB
}

// 创建Repository，接收外部的数据库管理对象
func NewStudentRepository(db *sql.DB) *StudentRepository {
	return &StudentRepository{
		db: db,
	}
}

// GetByID根据学生ID查询单条记录
func (r *StudentRepository) GetByID(studentID int) (Student, error) {
	var student Student

	err := r.db.QueryRow(
		"SELECT id, name, age FROM students WHERE id = ?",
		studentID,
	).Scan(
		&student.ID,
		&student.Name,
		&student.Age,
	)

	if err != nil {
		return Student{}, err
	}

	return student, nil
}

// List查询全部学生记录
func (r *StudentRepository) List() ([]Student, error) {
	// 1. 执行多行查询
	rows, err := r.db.Query(
		"SELECT id, name, age FROM students ORDER BY id",
	)

	if err != nil {
		return nil, err
	}

	// 2. 查询结果集由当前函数负责关闭
	// 注意这里关闭的是rows，不是r.db
	defer rows.Close()

	// 3. 准备保存查询结果的切片
	// 使用make可以确保空结果时返回一个非 nil 的空切片。
	students := make([]Student, 0)

	// 4. 遍历查询结果
	for rows.Next() {
		var student Student

		// 将当前行的数据写入student
		err := rows.Scan(
			&student.ID,
			&student.Name,
			&student.Age,
		)

		if err != nil {
			return nil, err
		}

		// 保存当前学生
		students = append(students, student)
	}

	// 5. 检查遍历过程中是否发生错误
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// 6. 返回全部学生记录
	return students, nil
}

func (r *StudentRepository) Create(student Student) error {
	_, err := r.db.Exec("INSERT INTO students(id, name, age) VALUES(?, ?, ?)",
		student.ID,
		student.Name,
		student.Age,
	)
	return err
}
