CREATE TABLE user_group
(
    id bigint primary key
);

CREATE TABLE user
(
    id         bigint primary key auto_increment,
    username   varchar(30) NOT NULL,
    user_group bigint      NOT NULL,
    email      varchar(30),
    phone      varchar(20),
    password   varchar(30) NOT NULL,
    FOREIGN KEY (user_group) references user_group (id)
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

CREATE TABLE files
(
    user_group   bigint      NOT NULL,
    uploader     bigint      NOT NULL,
    filename     varchar(30) NOT NULL,
    uploadedTime int         NOT NULL,
    primary key (user_group, filename),
    foreign key (user_group) references user_group (id),
    foreign key (uploader) references user (id)
);


INSERT INTO user_group (id)
VALUES (1);

INSERT INTO user (username, user_group, email, phone, password)
VALUES ("test", 1, "test@example.com", "12345678", "123456");

