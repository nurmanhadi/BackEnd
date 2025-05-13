DROP TABLE attendances;
CREATE TABLE attendances (
    id BIGINT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    student_id VARCHAR(36) NOT NULL,
    schedule_id BIGINT NOT NULL,
    status VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_attendaces_students FOREIGN KEY (student_id) REFERENCES students(id),
    CONSTRAINT fk_attendaces_schedules FOREIGN KEY (schedule_id) REFERENCES schedules(id)
);