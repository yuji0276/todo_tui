create table if not exists  tasks
(id text not null,
  title text not null,
  description text not null,
  done integer check(done in (0 , 1)),
  priority integer check(priority in (0 , 1 , 2 , 3)),
  due_date text,
  created_at text,
  updated_at text
  primary key  (id)
);
