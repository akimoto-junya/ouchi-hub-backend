-- storages
CREATE TABLE storages (
  id VARCHAR(36) NOT NULL,
  owner_id VARCHAR(36) NOT NULL UNIQUE,

  PRIMARY KEY (id)
);
--CREATE INDEX storage_owner_id_index ON storages (owner_id);

-- buckets
CREATE TABLE buckets (
  id VARCHAR(36) NOT NULL,
  name VARCHAR(256) NOT NULL,
  storage_id VARCHAR(36) NOT NULL,

  PRIMARY KEY (id),
  FOREIGN KEY (storage_id) REFERENCES storages (id) ON DELETE CASCADE
);
CREATE INDEX bucket_storage_id_index ON buckets (storage_id);

-- items
CREATE TABLE items (
  id VARCHAR(36) NOT NULL,
  name VARCHAR(256) NOT NULL,
  bucket_id VARCHAR(36) NOT NULL,
  is_file BOOL NOT NULL,
  lft BIGINT NOT NULL,
  rgt BIGINT NOT NULL,

  PRIMARY KEY (id),
  FOREIGN KEY (bucket_id) REFERENCES buckets (id) ON DELETE CASCADE,
  CHECK (lft < rgt)
);
CREATE INDEX item_bucket_id_lft_rgt_index ON items (bucket_id, lft, rgt);

-- contents
CREATE TABLE contents (
  id VARCHAR(36) NOT NULL,
  item_id VARCHAR(36) NOT NULL,
  path VARCHAR(2100) NOT NULL,
  file_type VARCHAR(64) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

  PRIMARY KEY(id),
  FOREIGN KEY (item_id) REFERENCES items (id) ON DELETE CASCADE
);
CREATE INDEX content_item_id_type_index ON contents (item_id, file_type);

