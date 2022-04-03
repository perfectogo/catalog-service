create unique index idx_book_category on book_category(book_id, category_id);

alter table authors add column created_at timestamp default current_timestamp;
alter table authors add column updated_at timestamp;
alter table authors add column deleted_at timestamp;

alter table categories add column created_at timestamp default current_timestamp;
alter table categories add column updated_at timestamp;
alter table categories add column deleted_at timestamp;

alter table books add column created_at timestamp default current_timestamp;
alter table books add column updated_at timestamp;
alter table books add column deleted_at timestamp;
