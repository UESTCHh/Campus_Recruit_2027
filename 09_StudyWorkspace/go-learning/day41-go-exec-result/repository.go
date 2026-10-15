package main

import "database/sql"

type Student struct {
	ID   int
	Name string
	Age  int
}

type StudentRepository struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) *StudentRepository {
	return &StudentRepository{db: db}
}

// 查询单个学生
func (r *StudentRepository) GetByID(id int) (Student, error) {
	var student Student

	err := r.db.QueryRow(
		"SELECT id, name, age FROM students WHERE id = ?",
		id,
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

// 查询所有学生
func (r *StudentRepository) List() ([]Student, error) {
	rows, err := r.db.Query(
		"SELECT id, name, age FROM students ORDER BY id",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	students := make([]Student, 0)

	for rows.Next() {
		var student Student

		if err := rows.Scan(
			&student.ID,
			&student.Name,
			&student.Age,
		); err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

// 新增学生
func (r *StudentRepository) Create(student Student) error {
	_, err := r.db.Exec(
		"INSERT INTO students (id, name, age) VALUES (?, ?, ?)",
		student.ID,
		student.Name,
		student.Age,
	)
	return err
}

// 根据ID修改学生信息
func (r *StudentRepository) Update(student Student) (int64, error) {
	result, err := r.db.Exec(
		"UPDATE students SET name = ?, age = ? WHERE id = ?",
		student.Name,
		student.Age,
		student.ID,
	)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// 根据ID删除学生
func (r *StudentRepository) Delete(studentID int) (int64, error) {
	result, err := r.db.Exec(
		"DELETE FROM students WHERE id = ?",
		studentID,
	)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
