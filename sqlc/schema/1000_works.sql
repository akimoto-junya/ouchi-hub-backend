CREATE TABLE categories (
  id VARCHAR(36) NOT NULL,
  name VARCHAR(512) NOT NULL,
  PRIMARY KEY(id)
);

CREATE TABLE makers (
  id VARCHAR(36) NOT NULL,
  name VARCHAR(512) NOT NULL,
  PRIMARY KEY(id)
);

CREATE TABLE works (
  id VARCHAR(36) NOT NULL,
  title VARCHAR(260) NOT NULL,
  category_id VARCHAR(36) NOT NULL,
  maker_id VARCHAR(36) NOT NULL,
  bucket_id VARCHAR(36) NOT NULL,
  PRIMARY KEY(id),
  FOREIGN KEY (category_id) REFERENCES categories (id),
  FOREIGN KEY (maker_id) REFERENCES makers (id)
);
