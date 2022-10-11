CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, email TEXT, password TEXT, created_at DATETIME, updated_at DATETIME);

CREATE TABLE IF NOT EXISTS tokens (id INTEGER PRIMARY KEY, content TEXT, user_id INTEGER, created_at DATETIME, updated_at DATETIME, FOREIGN KEY(user_id) REFERENCES users(id));

CREATE TRIGGER IF NOT EXISTS tokens_updated_at AFTER UPDATE OF id, user_id, content ON tokens FOR EACH ROW BEGIN UPDATE tokens SET updated_at = DATETIME ('NOW') WHERE rowid = new.rowid; END;
CREATE TRIGGER IF NOT EXISTS tokens_created_at AFTER INSERT ON tokens FOR EACH ROW BEGIN UPDATE tokens SET created_at = DATETIME ('NOW') WHERE rowid = new.rowid; END;