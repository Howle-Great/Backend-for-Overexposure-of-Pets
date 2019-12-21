create table users_rights
(
    u_id          integer    not null
        constraint users_rights_pk
            primary key,
    type_of_right varchar(5) not null
);

alter table users_rights
    owner to postgres;

create table users
(
    u_id     serial       not null
        constraint users_pk
            primary key
        constraint users_users_rights_u_id_fk
            references users_rights
            on delete cascade,
    nickname varchar(30)  not null,
    password varchar(100) not null,
    mail     varchar(100) not null
);

alter table users
    owner to postgres;

create unique index users_nickname_uindex
    on users (nickname);

create unique index users_u_id_uindex
    on users (u_id);

create unique index users_rights_u_id_uindex
    on users_rights (u_id);

create table places_in_hotels
(
    h_id           integer not null
        constraint places_in_hotels_pk
            primary key,
    type_of_places varchar not null,
    amount         integer not null
);

alter table places_in_hotels
    owner to postgres;

create table hotels
(
    h_id serial       not null
        constraint hotels_pk
            primary key
        constraint hotels_places_in_hotels_h_id_fk
            references places_in_hotels
            on update cascade on delete cascade,
    name varchar(100) not null,
    url  varchar
);

alter table hotels
    owner to postgres;

create unique index hotels_h_id_uindex
    on hotels (h_id);

create table order_list
(
    ord_id integer not null
        constraint order_list_pk
            primary key,
    h_id   integer not null
        constraint order_list_hotels_h_id_fk
            references hotels
);

alter table order_list
    owner to postgres;

create table orders
(
    ord_id         serial    not null
        constraint orders_pk
            primary key
        constraint orders_order_list_ord_id_fk
            references order_list
            on update cascade on delete cascade,
    u_id           integer   not null
        constraint orders_users_rights_u_id_fk
            references users_rights
            on update cascade on delete cascade,
    type_of_places varchar   not null,
    pick_from      varchar   not null,
    pick_date      timestamp not null,
    bring_in       varchar   not null,
    bring_date     timestamp not null
);

alter table orders
    owner to postgres;

create unique index orders_ord_id_uindex
    on orders (ord_id);

create unique index order_list_ord_id_uindex
    on order_list (ord_id);

create table companies_users
(
    com_id integer not null,
    u_id   integer not null
        constraint companies_users_pk
            primary key
        constraint companies_users_users_u_id_fk
            references users
            on update cascade on delete cascade
);

alter table companies_users
    owner to postgres;

create table companies
(
    com_id serial       not null
        constraint companies_pk
            primary key
        constraint companies_companies_users_com_id_fk
            references companies_users (com_id)
            on update cascade on delete cascade,
    name   varchar(100) not null,
    url    varchar
);

alter table companies
    owner to postgres;

create unique index companies_com_id_uindex
    on companies (com_id);

create unique index companies_users_com_id_uindex
    on companies_users (com_id);

create unique index companies_users_u_id_uindex
    on companies_users (u_id);

