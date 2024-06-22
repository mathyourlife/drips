CREATE TABLE IF NOT EXISTS exercise (
    exercise_id INTEGER PRIMARY KEY AUTOINCREMENT,
    exercise_class_id INTEGER NOT NULL,
    duration_seconds INTEGER,
    rest_seconds INTEGER,
    repeat INTEGER
);

CREATE TABLE IF NOT EXISTS exercise_modifier (
    exercise_modifier_id INTEGER PRIMARY KEY AUTOINCREMENT,
    exercise_id INTEGER NOT NULL,
    modifier_id INTEGER NOT NULL
);
