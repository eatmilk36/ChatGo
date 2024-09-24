create table ChatroomMessageHistory
(
    Id        int auto_increment primary key,
    Message   varchar(1000) charset utf8 not null,
    UserId    int         not null,
    TimeStamp bigint      not null,
    GroupName varchar(30) not null
);

create table users
(
    Id          int auto_increment primary key,
    Account     varchar(30) not null,
    Password    varchar(32) not null,
    CreatedTime datetime    not null,
    constraint Member_Account_uindex
        unique (Account)
);