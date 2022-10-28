create
database if not exists yanxue DEFAULT CHARACTER SET utf8 DEFAULT COLLATE utf8_general_ci;

create
user 'admin'@'%' IDENTIFIED BY 'admin';

GRANT ALL PRIVILEGES on yanxue.* to
'admin'@'%';

Flush
PRIVILEGES;

CREATE TABLE IF NOT EXISTS phone_codes
(
    id
    INT
    UNSIGNED
    AUTO_INCREMENT,
    created_at
    TIMESTAMP,
    updated_at
    TIMESTAMP,
    types
    varchar
(
    128
) NOT NULL,
    country varchar
(
    128
) NOT NULL,
    code varchar
(
    128
) NOT NULL,
    PRIMARY KEY
(
    id
)
    );

CREATE TABLE IF NOT EXISTS users
(
    id
    INT
    UNSIGNED
    AUTO_INCREMENT,
    created_at
    TIMESTAMP,
    updated_at
    TIMESTAMP,
    account
    varchar
(
    128
) NOT NULL,
    PASSWORD varchar
(
    128
) NOT NULL,
    PRIMARY KEY
(
    id
),
    UNIQUE KEY
(
    account
)
    );

CREATE TABLE IF NOT EXISTS roles
(
    id
    INT
    UNSIGNED
    AUTO_INCREMENT,
    created_at
    TIMESTAMP,
    updated_at
    TIMESTAMP,
    role_name
    varchar
(
    128
) NOT NULL,
    alias varchar
(
    128
) NOT NULL,
    PRIMARY KEY
(
    id
),
    UNIQUE KEY
(
    role_name
)
    );

CREATE TABLE IF NOT EXISTS user_roles
(
    id
    INT
    UNSIGNED
    AUTO_INCREMENT,
    created_at
    TIMESTAMP,
    updated_at
    TIMESTAMP,
    role_id
    int
    NOT
    NULL,
    user_id
    int
    NOT
    NULL,
    PRIMARY
    KEY
(
    id
)
    );