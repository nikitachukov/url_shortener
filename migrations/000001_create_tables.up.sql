create table public.shortener
(
    uuid         serial,
    key          varchar(10)   not null
        constraint shortener_pk
            primary key,
    original_url varchar(4096) not null
        constraint shortener_unique
            unique
);