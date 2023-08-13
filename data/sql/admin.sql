create database if not exists go_zero_admin;

use go_zero_admin;

create table if not exists user(
    id bigint not null auto_increment,
    username varchar(50) not null default'' comment '用户名',
    password char(32) not null default '' comment '密码',
    avatar varchar(255) not null default '' comment '头像',
    introduction varchar(244) not null default '' comment '个人介绍',
    role_id int  not null default 0 comment '角色id',
    status int not null default 0 comment'0-待激活1-已入职-2-离职中3-已离职',
    created_at datetime not null default CURRENT_TIMESTAMP,
    updated_at datetime not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    deleted_at datetime null,
    primary key(`id`),
    index(`username`),
    index(`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

create table if not exists role(
    id int not null primary key auto_increment,
    name varchar(50) not null default '' comment'角色名称',
    title varchar(50) not null default '' comment'角色中文名称',
    status int not null default 0 comment'0-未启用1-启用-2-禁用-3-删除',
    created_at datetime default CURRENT_TIMESTAMP comment '创建时间',
    updated_at datetime default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP comment '更新时间',
    deleted_at datetime default null comment'删除时间',
    primary key(`id`),
    index(`name`),
    index(`status`)
)engine=innodb default charset=utf8mb4 COLLATE=utf8mb4_unicode_ci;


create table if not exists rule(
    id int not null primary key auto_increment,
    name varchar(50) not null default '' comment'权限名称',
    title varchar(50) not null default '' comment'权限的中文名称',
    type int not null default 0 comment'0菜单-1按钮',
    created_at datetime default CURRENT_TIMESTAMP,
    updated_at datetime default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    deleted_at datetime default null,
    primary key(`id`),
    index(`username`),
    index(`status`)
) engine=innodb default charset=utf8mb4 COLLATE=utf8mb4_unicode_ci;


create table if not exists role_rule(
    id int not null auto_increment,
    role_id int not null default 0,
    rule_id int not null default 0,
    created_at datetime default CURRENT_TIMESTAMP,
    updated_at datetime default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    primary key(`id`),
    index(role_id, rule_id)
)engine=innodb default charset=utf8mb4 COLLATE=utf8mb4_unicode_ci;