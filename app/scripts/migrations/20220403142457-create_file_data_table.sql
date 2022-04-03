
-- +migrate Up
CREATE TABLE file (
    id int not null,
    file_name varchar(255) NOT NULL,
    extension varchar(4) NOT NULL,
    parsing_date_time  timestamp NOT NULL,
    PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE file;
