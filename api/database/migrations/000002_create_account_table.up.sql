CREATE TABLE IF NOT EXISTS account(
  id TEXT PRIMARY KEY NOT NULL,
  name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS account_user(
  role TEXT NOT NULL,
  account_id TEXT NOT NULL,
  user_id TEXT NOT NULL,
  PRIMARY KEY (account_id, user_id),
  FOREIGN KEY (account_id)
    REFERENCES account(id)
    ON DELETE CASCADE,
  FOREIGN KEY (user_id)
    REFERENCES user(id)
    ON DELETE CASCADE
);
