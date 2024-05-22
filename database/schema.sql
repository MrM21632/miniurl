create table url (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    shortened_url TEXT NOT NULL,
    original_url TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
);