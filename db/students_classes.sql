CREATE TABLE students_classes (
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    class_id INT NOT NULL,
    student_id VARCHAR(36) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_student_classes_classes FOREIGN KEY (class_id) REFERENCES classes(id),
    CONSTRAINT fk_student_classes_students FOREIGN KEY (student_id) REFERENCES students(id)
);