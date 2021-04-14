create table sectors
(
    id   serial      not null
        constraint sector_pk
            primary key,
    name varchar(10) not null
);

alter table sectors
    owner to postgres;

create unique index sector_id_uindex
    on sectors (id);

create table accounts
(
    id           serial       not null
        constraint accounts_pk
            primary key,
    first_name   varchar(100) not null,
    last_name    varchar(100) not null,
    phone_number varchar(100) not null,
    email        varchar(100) not null
);

alter table accounts
    owner to postgres;

create unique index accounts_id_uindex
    on accounts (id);

create table locations
(
    id           serial       not null
        constraint locations_pk
            primary key,
    account_id   integer      not null
        constraint locations_accounts_id_fk
            references accounts,
    address      varchar(255) not null,
    phone_number varchar(100) not null
);

alter table locations
    owner to postgres;

create unique index locations_id_uindex
    on locations (id);

create table roles
(
    id   serial       not null
        constraint roles_pk
            primary key,
    name varchar(100) not null
);

alter table roles
    owner to postgres;

create unique index roles_id_uindex
    on roles (id);

create table genre
(
    id   serial       not null
        constraint genre_pk
            primary key,
    name varchar(100) not null
);

alter table genre
    owner to postgres;

create unique index genre_id_uindex
    on genre (id);

create table halls
(
    id          serial       not null
        constraint halls_pk
            primary key,
    account_id  integer      not null
        constraint halls_accounts_id_fk
            references accounts,
    name        varchar(100) not null,
    capacity    integer      not null,
    location_id integer      not null
        constraint halls_locations_id_fk
            references locations
);

alter table halls
    owner to postgres;

create unique index halls_id_uindex
    on halls (id);

create table users
(
    id           serial       not null
        constraint users_pk
            primary key,
    account_id   integer      not null
        constraint users_accounts_id_fk
            references accounts,
    first_name   varchar(100) not null,
    last_name    varchar(100) not null,
    role_id      integer      not null
        constraint users_roles_id_fk
            references roles,
    location_id  integer      not null
        constraint users_locations_id_fk
            references locations,
    phone_number varchar(100) not null
);

alter table users
    owner to postgres;

create unique index users_id_uindex
    on users (id);

create table performance
(
    id         serial       not null
        constraint performance_pk
            primary key,
    account_id integer      not null
        constraint performance_accounts_id_fk
            references accounts,
    name       varchar(100) not null,
    genre_id   integer      not null
        constraint performance_genre_id_fk
            references genre,
    duration   time         not null
);

alter table performance
    owner to postgres;

create unique index performance_id_uindex
    on performance (id);

create table schedule
(
    id             serial    not null
        constraint schedule_pk
            primary key,
    account_id     integer   not null
        constraint schedule_accounts_id_fk
            references accounts,
    performance_id integer   not null
        constraint schedule_performance_id_fk
            references performance,
    date           timestamp not null,
    hall_id        integer   not null
        constraint schedule_halls_id_fk
            references halls
);

alter table schedule
    owner to postgres;

create unique index schedule_id_uindex
    on schedule (id);

create table poster
(
    id          serial       not null
        constraint poster_pk
            primary key,
    account_id  integer      not null
        constraint poster_accounts_id_fk
            references accounts,
    schedule_id integer      not null
        constraint poster_schedule_id_fk
            references schedule,
    comment     varchar(255) not null
);

alter table poster
    owner to postgres;

create unique index poster_id_uindex
    on poster (id);

create table places
(
    id        serial  not null
        constraint places_pk
            primary key,
    sector_id integer not null
        constraint places_sectors_id_fk
            references sectors,
    name      integer not null
);

alter table places
    owner to postgres;

create unique index places_id_uindex
    on places (id);

create table price
(
    id             serial  not null
        constraint price_pk
            primary key,
    account_id     integer not null
        constraint price_accounts_id_fk
            references accounts,
    sector_id      integer not null
        constraint price_sectors_id_fk
            references sectors,
    performance_id integer not null
        constraint price_performance_id_fk
            references performance,
    price          integer not null
);

alter table price
    owner to postgres;

create unique index price_id_uindex
    on price (id);

create table tickets
(
    id            serial                not null
        constraint tickets_pk
            primary key,
    account_id    integer               not null
        constraint tickets_accounts_id_fk
            references accounts,
    schedule_id   integer               not null
        constraint tickets_schedule_id_fk
            references schedule,
    place_id      integer               not null
        constraint tickets_places_id_fk
            references places,
    date_of_issue timestamp             not null,
    paid          boolean default false not null,
    reservation   boolean default false not null,
    destroyed     boolean default false not null
);

alter table tickets
    owner to postgres;

create unique index tickets_id_uindex
    on tickets (id);



INSERT INTO sectors (name)
VALUES ('A');
INSERT INTO sectors (name)
VALUES ('B');
INSERT INTO sectors (name)
VALUES ('C');
INSERT INTO sectors (name)
VALUES ('D');
INSERT INTO sectors (name)
VALUES ('E');
INSERT INTO sectors (name)
VALUES ('F');
INSERT INTO sectors (name)
VALUES ('G');
INSERT INTO sectors (name)
VALUES ('I');
INSERT INTO sectors (name)
VALUES ('J');
INSERT INTO sectors (name)
VALUES ('K');

INSERT INTO places (sector_id, name)
VALUES (9, 1);
INSERT INTO places (sector_id, name)
VALUES (10, 1);
INSERT INTO places (sector_id, name)
VALUES (11, 1);
INSERT INTO places (sector_id, name)
VALUES (12, 1);
INSERT INTO places (sector_id, name)
VALUES (13, 1);
INSERT INTO places (sector_id, name)
VALUES (14, 1);
INSERT INTO places (sector_id, name)
VALUES (15, 1);
INSERT INTO places (sector_id, name)
VALUES (16, 1);
INSERT INTO places (sector_id, name)
VALUES (17, 1);
INSERT INTO places (sector_id, name)
VALUES (18, 1);

INSERT INTO accounts(first_name, last_name, phone_number, email)
VALUES ('Alexander', 'Antoshkov', '+37544*******', 'aaaa@gmail.com');
INSERT INTO accounts(first_name, last_name, phone_number, email)
VALUES ('Valeria', 'Abramtsova', '+37544*******', 'vvv@gmail.com');
INSERT INTO accounts(first_name, last_name, phone_number, email)
VALUES ('Alexander', 'Palchik', '+37544*******', 'ppppp@gmail.com');
INSERT INTO accounts(first_name, last_name, phone_number, email)
VALUES ('Kiril', 'Bunich', '+37544*******', 'kkkkkk@gmail.com');
INSERT INTO accounts(first_name, last_name, phone_number, email)
VALUES ('Daniel', 'Martunov', '+37544*******', 'aaaa@gmail.com');

INSERT INTO locations (account_id, address, phone_number)
VALUES (1, 'Gaidara_6', '+375443564987');
INSERT INTO locations (account_id, address, phone_number)
VALUES (2, 'Woll_street', '+375443974651');
INSERT INTO locations (account_id, address, phone_number)
VALUES (3, 'GreenWay_street', '+375442874593');
INSERT INTO locations (account_id, address, phone_number)
VALUES (4, 'Red_street', '+375441893500');
INSERT INTO locations (account_id, address, phone_number)
VALUES (5, 'High_street', '+375441438532');

INSERT INTO roles(name)
VALUES ('Actor');
INSERT INTO roles(name)
VALUES ('Producer');
INSERT INTO roles(name)
VALUES ('Prompter');
INSERT INTO roles(name)
VALUES ('Technical worker');
INSERT INTO roles(name)
VALUES ('Artist and designer');
INSERT INTO roles(name)
VALUES ('Manager-organizer');
INSERT INTO roles(name)
VALUES ('Viewer Assistant');

INSERT INTO genre (name)
VALUES ('a musical');
INSERT INTO genre (name)
VALUES ('a melodrama');
INSERT INTO genre (name)
VALUES ('a comedy');
INSERT INTO genre (name)
VALUES ('a tragedy');
INSERT INTO genre (name)
VALUES ('a history play');
INSERT INTO genre (name)
VALUES ('a farce');
INSERT INTO genre (name)
VALUES ('an epic');
INSERT INTO genre (name)
VALUES ('an opera');
INSERT INTO genre (name)
VALUES ('a vaudeville');
INSERT INTO genre (name)
VALUES ('a pantomime');
INSERT INTO genre (name)
VALUES ('an operetta');

INSERT INTO halls (account_id, name, capacity)
VALUES (1, 'Small', 100);
INSERT INTO halls (account_id, name, capacity)
VALUES (1, 'Big', 3000);
INSERT INTO halls (account_id, name, capacity)
VALUES (1, 'Middle', 1500);
INSERT INTO halls (account_id, name, capacity)
VALUES (1, 'Dollhouse', 1000);
INSERT INTO halls (account_id, name, capacity)
VALUES (1, 'Happy', 500);

INSERT INTO users(account_id, first_name, last_name, role_id, location_id, phone_number)
VALUES (1, 'Charles', 'Dean', 1, 1, '+375445239375');
INSERT INTO users(account_id, first_name, last_name, role_id, location_id, phone_number)
VALUES (1, 'Brian', 'Cobb', 2, 1, '+375445234353');
INSERT INTO users(account_id, first_name, last_name, role_id, location_id, phone_number)
VALUES (1, 'Jordan', 'Moore', 3, 1, '+375445234323');
INSERT INTO users(account_id, first_name, last_name, role_id, location_id, phone_number)
VALUES (1, 'Ethan', 'Snow', 4, 1, '+375445232398');
INSERT INTO users(account_id, first_name, last_name, role_id, location_id, phone_number)
VALUES (1, 'David', 'Leonard', 5, 1, '+375445239455');
INSERT INTO users(account_id, first_name, last_name, role_id, location_id, phone_number)
VALUES (1, 'Dustin', 'Mason', 6, 1, '+375445232125');
INSERT INTO users(account_id, first_name, last_name, role_id, location_id, phone_number)
VALUES (1, 'David', 'Bradley', 7, 1, '+375445209121');

INSERT INTO performance (account_id, name, genre_id, duration)
VALUES (1, 'The Dragon', 1, '4:00');
INSERT INTO performance (account_id, name, genre_id, duration)
VALUES (1, 'Chasing two hares', 5, '2:00');
INSERT INTO performance (account_id, name, genre_id, duration)
VALUES (1, 'Life and destiny', 3, '3:00');
INSERT INTO performance (account_id, name, genre_id, duration)
VALUES (1, 'And the day lasts longer than a century', 10, '5:00');
INSERT INTO performance (account_id, name, genre_id, duration)
VALUES (1, 'Master and Margarita', 8, '7:00');

INSERT INTO schedule (account_id, performance_id, date)
VALUES (1, 1, '2021-04-13 16:00');
INSERT INTO schedule (account_id, performance_id, date)
VALUES (1, 2, '2021-04-25 13:00');
INSERT INTO schedule (account_id, performance_id, date)
VALUES (1, 3, '2021-04-19 19:00');
INSERT INTO schedule (account_id, performance_id, date)
VALUES (1, 4, '2021-05-10 14:00');
INSERT INTO schedule (account_id, performance_id, date)
VALUES (1, 5, '2021-04-15 21:00');

INSERT INTO schedule (account_id, performance_id, date, hall_id)
VALUES (1, 1, '2021-04-13 16:00', 3);
INSERT INTO schedule (account_id, performance_id, date, hall_id)
VALUES (1, 2, '2021-04-25 13:00', 2);
INSERT INTO schedule (account_id, performance_id, date, hall_id)
VALUES (1, 3, '2021-04-19 19:00', 1);
INSERT INTO schedule (account_id, performance_id, date, hall_id)
VALUES (1, 4, '2021-05-10 14:00', 4);
INSERT INTO schedule (account_id, performance_id, date, hall_id)
VALUES (1, 5, '2021-04-15 21:00', 5);

INSERT INTO poster (account_id, schedule_id, comment)
VALUES (1, 6, 'We invite you! It will be cool!!!');
INSERT INTO poster (account_id, schedule_id, comment)
VALUES (1, 7, 'We invite you! It will be cool!!!');
INSERT INTO poster (account_id, schedule_id, comment)
VALUES (1, 8, 'We invite you! It will be cool!!!');
INSERT INTO poster (account_id, schedule_id, comment)
VALUES (1, 9, 'We invite you! It will be cool!!!');
INSERT INTO poster (account_id, schedule_id, comment)
VALUES (1, 10, 'We invite you! It will be cool!!!');

INSERT INTO price (account_id, sector_id, performance_id, price)
VALUES (1, 9, 1, 40);
INSERT INTO price (account_id, sector_id, performance_id, price)
VALUES (1, 9, 2, 34);
INSERT INTO price (account_id, sector_id, performance_id, price)
VALUES (1, 9, 3, 97);
INSERT INTO price (account_id, sector_id, performance_id, price)
VALUES (1, 9, 4, 76);
INSERT INTO price (account_id, sector_id, performance_id, price)
VALUES (1, 9, 5, 88);
INSERT INTO price (account_id, sector_id, performance_id, price)
VALUES (1, 10, 1, 39);
INSERT INTO price (account_id, sector_id, performance_id, price)
VALUES (1, 10, 2, 33);
INSERT INTO price (account_id, sector_id, performance_id, price)
VALUES (1, 10, 3, 78);
INSERT INTO price (account_id, sector_id, performance_id, price)
VALUES (1, 11, 1, 38);
INSERT INTO price (account_id, sector_id, performance_id, price)
VALUES (1, 12, 1, 37);
INSERT INTO price (account_id, sector_id, performance_id, price)
VALUES (1, 13, 1, 36);
INSERT INTO price (account_id, sector_id, performance_id, price)
VALUES (1, 14, 1, 35);
INSERT INTO price (account_id, sector_id, performance_id, price)
VALUES (1, 15, 1, 34);
INSERT INTO price (account_id, sector_id, performance_id, price)
VALUES (1, 16, 1, 33);
INSERT INTO price (account_id, sector_id, performance_id, price)
VALUES (1, 17, 1, 32);
INSERT INTO price (account_id, sector_id, performance_id, price)
VALUES (1, 18, 1, 31);

INSERT INTO tickets (account_id, schedule_id, place_id, date_of_issue, paid)
values (1, 6, 1, now(), true);
INSERT INTO tickets (account_id, schedule_id, place_id, date_of_issue, paid)
values (1, 7, 2, now(), true);
INSERT INTO tickets (account_id, schedule_id, place_id, date_of_issue, paid)
values (1, 8, 3, now(), true);
INSERT INTO tickets (account_id, schedule_id, place_id, date_of_issue)
values (1, 9, 4, now());
INSERT INTO tickets (account_id, schedule_id, place_id, date_of_issue, paid)
values (1, 10, 5, now(), true);
INSERT INTO tickets (account_id, schedule_id, place_id, date_of_issue, paid)
values (1, 6, 6, now(), true);
INSERT INTO tickets (account_id, schedule_id, place_id, date_of_issue, paid)
values (1, 7, 7, now(), true);
INSERT INTO tickets (account_id, schedule_id, place_id, date_of_issue, paid)
values (1, 8, 8, now(), true);
INSERT INTO tickets (account_id, schedule_id, place_id, date_of_issue, paid, reservation)
values (1, 9, 9, now(), true, true);
INSERT INTO tickets (account_id, schedule_id, place_id, date_of_issue, paid)
values (1, 10, 10, now(), true);
INSERT INTO tickets (account_id, schedule_id, place_id, date_of_issue, paid, reservation)
values (1, 6, 1, now(), true, true);
INSERT INTO tickets (account_id, schedule_id, place_id, date_of_issue, paid)
values (1, 7, 2, now(), true);
INSERT INTO tickets (account_id, schedule_id, place_id, date_of_issue, paid, reservation, destroyed)
values (1, 8, 3, now(), true, true, true);
INSERT INTO tickets (account_id, schedule_id, place_id, date_of_issue, paid)
values (1, 9, 4, now(), true);
INSERT INTO tickets (account_id, schedule_id, place_id, date_of_issue, paid)
values (1, 10, 5, now(), true);
INSERT INTO tickets (account_id, schedule_id, place_id, date_of_issue)
values (1, 6, 6, now());
INSERT INTO tickets (account_id, schedule_id, place_id, date_of_issue)
values (1, 7, 7, now());
INSERT INTO tickets (account_id, schedule_id, place_id, date_of_issue, paid, reservation, destroyed)
values (1, 8, 8, now(), true, true, true);
INSERT INTO tickets (account_id, schedule_id, place_id, date_of_issue, paid)
values (1, 9, 9, now(), true);
INSERT INTO tickets (account_id, schedule_id, place_id, date_of_issue, paid, reservation)
values (1, 10, 10, now(), true, true);
INSERT INTO tickets (account_id, schedule_id, place_id, date_of_issue)
values (1, 6, 1, now());
INSERT INTO tickets (account_id, schedule_id, place_id, date_of_issue, paid)
values (1, 7, 2, now(), true);
INSERT INTO tickets (account_id, schedule_id, place_id, date_of_issue, paid, reservation, destroyed)
values (1, 8, 3, now(), true, true, true);
INSERT INTO tickets (account_id, schedule_id, place_id, date_of_issue)
values (1, 9, 4, now());
INSERT INTO tickets (account_id, schedule_id, place_id, date_of_issue, paid, reservation)
values (1, 10, 5, now(), true, true);