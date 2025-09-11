SELECT students.id, students.name
FROM students
LEFT JOIN student_courses
  ON students.id = student_courses.student_id AND student_courses.course_id = 7
WHERE student_courses.student_id IS NULL;