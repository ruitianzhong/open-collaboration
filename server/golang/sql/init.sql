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

CREATE TABLE post
(
    id         bigint primary key auto_increment,
    author     bigint,
    user_group bigint,
    time       int,
    title      varchar(30),
    CHECK ( time > 0 ),
    FOREIGN KEY (author) references user (id),
    FOREIGN KEY (user_group) references user_group (id)
);



INSERT INTO user_group (group_id)
VALUES ("1");

INSERT INTO user (username, group_id, email, phone, password, user_id)
VALUES ("test", "1", "test@example.com", "12345678", "123456", "1");

