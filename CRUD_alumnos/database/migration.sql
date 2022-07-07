CREATE TABLE `students`
(
    id   bigint auto_increment,
    nombre varchar(255) NOT NULL,
    dni varchar(255) NOT NULL,
    direccion varchar(255) NOT NULL,
    fecha_nacimiento varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO `students` (`nombre`, `dni`,`direccion`,`fecha_nacimiento`) VALUES 
       ('Jose', '12345678', 'Calle falsa 123', '2020-01-01'),
       ('Juan', '87654321', 'Calle verdadera 456', '2020-01-01'),
       ('Adrian','1234124', 'Valle', '2020-01-01'),
       ('John', '1234124', 'Valle', '2020-01-01'),
       ('Mary', '1234124', 'Valle', '2020-01-01');