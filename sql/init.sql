CREATE TABLE IF NOT EXISTS users (
  id varchar PRIMARY KEY,
  username varchar UNIQUE, 
  password varchar
);

CREATE TABLE IF NOT EXISTS boxes (
  id varchar PRIMARY KEY,
  title varchar
);

CREATE TABLE IF NOT EXISTS items (
  id varchar PRIMARY KEY,
  title varchar,
  amount int
);

CREATE TABLE IF NOT EXISTS boxes_items (
  boxid varchar,
  itemid varchar,
  
  CONSTRAINT pk_boxes_items PRIMARY KEY(boxid, itemid),
  FOREIGN KEY (boxid) REFERENCES boxes(id),
  FOREIGN KEY (itemid) REFERENCES items(id)
);

CREATE TABLE IF NOT EXISTS users_boxes (
  userid varchar,
  boxid varchar,

  CONSTRAINT pk_users_boxes PRIMARY KEY(userid, boxid),
  FOREIGN KEY (userid) REFERENCES users(id),
  FOREIGN KEY (boxid) REFERENCES boxes(id)
);
