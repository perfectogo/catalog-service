create table authors (
    author_id uuid not null primary key,
    name varchar(48) not null
);

create table categories (
    category_id uuid not null primary key,
    name varchar(128) not null,
    parent_uuid uuid
);

create table books (
    book_id uuid not null primary key,
    name varchar(128) not null,
    author_id uuid not null references authors(author_id)
);

create table book_category (
    book_id uuid not null references books(book_id),
    category_id uuid not null references categories(category_id)
);
