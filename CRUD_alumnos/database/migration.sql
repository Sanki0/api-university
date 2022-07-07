CREATE TABLE `students`
(
    id   bigint auto_increment,
    name varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO `students` (`name`)
VALUES ('Jose'),
       ('Adrian'),
       ('John'),
       ('Mary');