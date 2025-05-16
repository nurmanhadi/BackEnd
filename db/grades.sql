CREATE TABLE grades (
    id BIGINT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    student_id VARCHAR(36) NOT NULL,
    subject_id INT NOT NULL,
    attendance_score INT NOT NULL,
    task_score INT NOT NULL,
    midtern_score INT NOT NULL,
    final_score INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_grades_students FOREIGN KEY (student_id) REFERENCES students(id),
    CONSTRAINT fk_grades_subjects FOREIGN KEY (subject_id) REFERENCES subjects(id)
);