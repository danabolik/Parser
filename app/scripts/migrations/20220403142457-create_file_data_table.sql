
-- +migrate Up
CREATE TABLE file_data (
    id int not null,
    file_name varchar(255) NOT NULL,
    extension varchar(4) NOT NULL,
    `data` text NOT NULL,
    parsing_date_time  timestamp NOT NULL,
    PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE file_data;
