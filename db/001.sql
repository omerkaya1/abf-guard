create table ip_list (
    id serial primary key,
    ip text not null,
    bl boolean not null
);

create index ip_idx on ip_list (ip);
