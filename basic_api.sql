create table articles
(
    id         int auto_increment
        primary key,
    username   varchar(255) null,
    title      varchar(255) null,
    content    text         null,
    created_at datetime     null,
    constraint fk_articles_comments_1
        foreign key (id) references comments (id)
);

create table comments
(
    id         int          not null
        primary key,
    username   varchar(255) null,
    content    text         null,
    created_at datetime     null,
    article_id int          null
);

create table users
(
    id         int auto_increment
        primary key,
    first_name varchar(255) not null,
    last_name  varchar(255) not null,
    dob        date         not null,
    city       varchar(255) not null,
    username   varchar(255) not null,
    password   varchar(255) not null
);
