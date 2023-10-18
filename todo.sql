CREATE TABLE todo
(
    id          int           NOT NULL GENERATED BY DEFAULT AS IDENTITY,
    title       varchar(200)  NOT NULL,
    description varchar(1024) NULL,
    reminder    timestamp     NULL,
    CONSTRAINT todo_pk PRIMARY KEY (id)
);