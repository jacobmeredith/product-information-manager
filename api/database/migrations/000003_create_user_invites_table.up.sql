ALTER TABLE account_user ADD COLUMN active boolean NOT NULL DEFAULT true;

CREATE TABLE IF NOT EXISTS account_user_invite(
    id TEXT PRIMARY KEY NOT NULL,
    account_id TEXT NOT NULL, 
    email TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (account_id)
      REFERENCES account(id)
      ON DELETE CASCADE
);
