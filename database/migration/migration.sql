CREATE TABLE `students`
(
    nombre varchar(255) NOT NULL,
    dni varchar(255) NOT NULL,
    direccion varchar(255) NOT NULL,
    fecha_nacimiento varchar(255) NOT NULL,
    PRIMARY KEY (`dni`)
);
INSERT INTO `students` (`nombre`, `dni`,`direccion`,`fecha_nacimiento`) VALUES 
       ('Jose', '12345678', 'Calle falsa 123', '2020-01-01'),
       ('Juan', '87654321', 'Calle verdadera 456', '2020-01-01'),
       ('Adrian','1234124', 'Valle', '2020-01-01'),
       ('John', '12345154', 'Valle', '2020-01-01'),
       ('Mary', '12453124', 'Valle', '2020-01-01');


CREATE TABLE `courses`
(
    nombre varchar(255) NOT NULL,
    descripcion varchar(255) NOT NULL,
    temas varchar(255) NOT NULL,
    PRIMARY KEY (`nombre`)
);
INSERT INTO `courses` (`nombre`, `descripcion`, `temas`) VALUES 
       ('PHP', 'Programacion en PHP', 'PHP, MySQL, HTML, CSS'),
       ('Java', 'Programacion en Java', 'Java, MySQL, HTML, CSS'),
       ('Python', 'Programacion en Python', 'Python, MySQL, HTML, CSS'),
       ('C#', 'Programacion en C#', 'C#, MySQL, HTML, CSS'),
       ('C++', 'Programacion en C++', 'C++, MySQL, HTML, CSS');


CREATE TABLE records
(
    student varchar(255), 
    course varchar(255),
    startdate varchar(255) NOT NULL,
    finishdate varchar(255) NOT NULL,
    PRIMARY KEY (student,course ),
    FOREIGN KEY (student) REFERENCES students(dni) ,
    FOREIGN KEY (course) REFERENCES courses(nombre)    
);
INSERT INTO `records` (`student`, `course`, `startdate`, `finishdate`) VALUES 
       ('12345678', 'PHP', '2020-01-01', '2020-01-01'),
       ('12345678', 'Java', '2020-01-01', '2020-01-01'),
       ('12345678', 'Python', '2020-01-01', '2020-01-01'),
       ('12345678', 'C#', '2020-01-01', '2020-01-01'),
       ('12345678', 'C++', '2020-01-01', '2020-01-01'),
       ('87654321', 'PHP', '2020-01-01', '2020-01-01'),
       ('87654321', 'Java', '2020-01-01', '2020-01-01'),
       ('87654321', 'Python', '2020-01-01', '2020-01-01'),
       ('87654321', 'C#', '2020-01-01', '2020-01-01'),
       ('87654321', 'C++', '2020-01-01', '2020-01-01'),
       ('1234124', 'PHP', '2020-01-01', '2020-01-01'),
       ('1234124', 'Java', '2020-01-01', '2020-01-01'),
       ('1234124', 'Python', '2020-01-01', '2020-01-01'),
       ('1234124', 'C#', '2020-01-01', '2020-01-01'),
       ('1234124', 'C++', '2020-01-01', '2020-01-01'),
       ('12345154', 'PHP', '2020-01-01', '2020-01-01'),
       ('12345154', 'Java', '2020-01-01', '2020-01-01'),
       ('12345154', 'Python', '2020-01-01', '2020-01-01'),
       ('12345154', 'C#', '2020-01-01', '2020-01-01');