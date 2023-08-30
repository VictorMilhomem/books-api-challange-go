CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS authors (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(50)
);

CREATE TABLE IF NOT EXISTS books (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(50),
    edition INT NOT NULL,
    publication_year VARCHAR(4) NOT NULL
);

CREATE TABLE IF NOT EXISTS book_authors (
    book_id UUID REFERENCES books(id),
    author_id UUID REFERENCES authors(id),
    PRIMARY KEY (book_id, author_id)
);