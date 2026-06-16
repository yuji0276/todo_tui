CREATE TABLE if NOT EXISTS tasks(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  title TEXT NOT NULL,
  description TEXT,
  done INTEGER CHECK(done IN(0,
  1)),
  priority INTEGER CHECK(priority IN(0,
  1,
  2,
  3)),
  due_date TEXT,
  created_at TEXT,
  updated_at TEXT
);
