CREATE TABLE if NOT EXISTS tasks(
  id TEXT NOT NULL,
  title TEXT NOT NULL,
  description TEXT NOT NULL,
  done INTEGER CHECK(done IN(
    0,
    1
  )),
  priority INTEGER CHECK(priority IN(
    0,
    1,
    2,
    3
  )),
  due_date TEXT,
  created_at TEXT,
  updated_at TEXT,
  PRIMARY KEY(id)
);
