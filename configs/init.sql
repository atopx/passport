create table user
(
    id        int unsigned auto_increment primary key,
    account   varchar(50) unique not null comment '账号',
    rule      enum ('RULE_NONE', 'RULE_ROUTINE', 'RULE_VIP', 'RULE_ADMIN', 'RULE_BLACK', 'RULE_OWNER') default 'RULE_NONE' comment '角色',
    login_at  int comment '登陆时间',
    create_at int comment '创建时间',
    update_at int comment '更新时间',
    delete_at int comment '删除时间'
);

create table password
(
    id        int unsigned auto_increment primary key,
    user_id   int          not null comment '用户ID',
    keyword   int          not null comment '关键词',
    value     varchar(128) not null comment '密码值',
    domain    varchar(255) default '' comment '作用域',
    encrypted bool         default false comment '加密的',
    creator   int comment '创建人',
    updater   int comment '更新人',
    deleter   int comment '删除人',
    create_at int comment '创建时间',
    update_at int comment '更新时间',
    delete_at int comment '删除时间'
);

create table token
(
    id        int unsigned auto_increment primary key,
    user_id   int  not null comment '用户ID',
    client    varchar(100) default '' comment '认证客户端',
    invalid   bool not null comment '已失效',
    expire    int          default 0 comment '有效期',
    create_at int  not null comment '创建时间'
);

