CREATE TABLE user_group
(
    group_id varchar(20) primary key
);

CREATE TABLE user
(
    user_id  varchar(20) primary key,
    username varchar(30) NOT NULL,
    group_id varchar(20) NOT NULL,
    email    varchar(30),
    phone    varchar(20),
    password varchar(30) NOT NULL,
    FOREIGN KEY (group_id) references user_group (group_id)
);

CREATE TABLE id_allocation
(
    topic varchar(20) primary key,
    count bigint NOT NULL
);

INSERT INTO id_allocation (topic, count)
values ('docs', 100);

INSERT INTO user_group (group_id)
VALUES ('1');

INSERT INTO user (username, group_id, email, phone, password, user_id)
VALUES ('test', '1', 'test@example.com', '12345678"', '123456', '1');

