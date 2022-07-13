CREATE TABLE `students`
(
    nombre varchar(255) NOT NULL,
    dni varchar(255) NOT NULL,
    direccion varchar(255) NOT NULL,
    fecha_nacimiento varchar(255) NOT NULL,
    PRIMARY KEY (`dni`)
);

CREATE TABLE `courses`
(
    nombre varchar(255) NOT NULL,
    descripcion varchar(255) NOT NULL,
    temas varchar(255) NOT NULL,
    PRIMARY KEY (`nombre`)
);

INSERT INTO `students` (`nombre`, `dni`,`direccion`,`fecha_nacimiento`) VALUES 
       ('Jose', '12345678', 'Calle falsa 123', '2020-01-01'),
       ('Juan', '87654321', 'Calle verdadera 456', '2020-01-01'),
       ('Adrian','1234124', 'Valle', '2020-01-01'),
       ('John', '12345154', 'Valle', '2020-01-01'),
       ('Mary', '12453124', 'Valle', '2020-01-01');

INSERT INTO `courses` (`nombre`, `descripcion`, `temas`) VALUES 
       ('PHP', 'Programacion en PHP', 'PHP, MySQL, HTML, CSS'),
       ('Java', 'Programacion en Java', 'Java, MySQL, HTML, CSS'),
       ('Python', 'Programacion en Python', 'Python, MySQL, HTML, CSS'),
       ('C#', 'Programacion en C#', 'C#, MySQL, HTML, CSS'),
       ('C++', 'Programacion en C++', 'C++, MySQL, HTML, CSS');