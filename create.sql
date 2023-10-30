CREATE table IF NOT EXISTS accounts (
                                        id SERIAL primary key,
                                        first_Name varchar,
                                        second_Name varchar,
                                        email varchar unique,
                                        password varchar
);
CREATE table IF NOT EXISTS bills (
                                     id SERIAL primary key,
                                     account_id int references accounts (id),
                                     number varchar,
                                     card   int,
                                     sum_limit int
);
CREATE table IF NOT EXISTS cards (
                                     id SERIAL primary key,
                                     bill_id int references bills (id),
                                     number bigint,
                                     cvv varchar,
                                     expiration_date timestamp,
                                     balance money,
                                     isCardActive bool
);
CREATE table IF NOT EXISTS history (
                                       history_id SERIAL primary key,
                                       destination_card_id int references cards (id),
                                       arrival_card_id int references cards (id),
                                       date timestamp,
                                       operation_type varchar,
                                       sum int

);

