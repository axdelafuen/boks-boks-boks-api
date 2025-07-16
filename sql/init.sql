CREATE TABLE IF NOT EXISTS box (
  id varchar PRIMARY KEY,
  title varchar
);

CREATE TABLE IF NOT EXISTS item (
  id varchar PRIMARY KEY,
  title varchar,
  amount int
);

CREATE TABLE IF NOT EXISTS boxes_items (
  boxid varchar,
  itemid varchar,
  
  CONSTRAINT pk_boxes_items PRIMARY KEY(boxid, itemid),
  FOREIGN KEY (boxid) REFERENCES box(id),
  FOREIGN KEY (itemid) REFERENCES item(id)
);
