-- +migrate Up
CREATE TABLE file_rows
(
    id      INT  NOT NULL,
    file_id INT  NOT NULL,
    `data`  text NOT NULL,

    PRIMARY KEY (id),
    CONSTRAINT `FK-file-rows_file_id`
        FOREIGN KEY (file_id)
            REFERENCES file (id)
);

-- +migrate Down

DROP TABLE file_rows;
