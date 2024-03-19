CREATE TABLE items (
  id VARCHAR(64) NOT NULL,
  name VARCHAR(256) NOT NULL,
  type_id VARCHAR(64) NOT NULL,
  work_id VARCHAR(64) NOT NULL,
  lft BIGINT NOT NULL,
  rgt BIGINT NOT NULL,
  CHECK (lft < rgt),
  PRIMARY KEY (id),
  FOREIGN KEY (type_id) REFERENCES item_types(id)
);

CREATE TABLE item_types (
  id VARCHAR(64) NOT NULL,
  name VARCHAR(64) NOT NULL,
  PRIMARY KEY (id)
);
