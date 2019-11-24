create table whitelist (
    id serial primary key,
    ip text not null
);

create table blacklist (
    id serial primary key,
    ip text not null
);
