CREATE TABLE students (
    dni varchar(255) NOT NULL,
    nombre varchar(255) NOT NULL,
    direccion varchar(255),
    fecha_nacimiento varchar(255),
    PRIMARY KEY (dni)
);

INSERT INTO
    students (dni, nombre, direccion, fecha_nacimiento)
VALUES
    (
        '12345678',
        'Jose',
        'Calle falsa 123',
        '2020-01-01'
    ),
    (
        '87654321',
        'Juan',
        'Calle verdadera 456',
        '2020-01-01'
    ),
    ('1234124', 'Adrian', 'Valle', '2020-01-01'),
    ('12345154', 'John', 'Valle', '2020-01-01'),
    ('12453124', 'Mary', 'Valle', '2020-01-01');

CREATE TABLE courses (
    id_courses varchar(10) NOT NULL,
    nombre varchar(255) NOT NULL,
    descripcion varchar(255),
    temas varchar(255),
    PRIMARY KEY (id_courses)
);

INSERT INTO
    courses (id_courses, nombre, descripcion, temas)
VALUES
    (
        '1',
        'PHP',
        'Programacion en PHP',
        'PHP, MySQL, HTML, CSS'
    ),
    (
        '2',
        'Java',
        'Programacion en Java',
        'Java, MySQL, HTML, CSS'
    ),
    (
        '3',
        'Python',
        'Programacion en Python',
        'Python, MySQL, HTML, CSS'
    ),
    (
        '4',
        'C#',
        'Programacion en C#',
        'C#, MySQL, HTML, CSS'
    ),
    (
        '5',
        'C++',
        'Programacion en C++',
        'C++, MySQL, HTML, CSS'
    );

CREATE TABLE records (
    id_records varchar(10) NOT NULL,
    student varchar(255) NOT NULL,
    course varchar(255) NOT NULL,
    startdate varchar(255) NOT NULL,
    finishdate varchar(255) NOT NULL,
    PRIMARY KEY (id_records),
    FOREIGN KEY (student) REFERENCES students(dni),
    FOREIGN KEY (course) REFERENCES courses(id_courses)
);

INSERT INTO
    records (
        id_records,
        student,
        course,
        startdate,
        finishdate
    )
VALUES
    ('1', '12345678', '1', '2020-01-01', '2020-01-01'),
    ('2', '12345678', '2', '2020-01-01', '2020-01-01'),
    ('3', '12345678', '3', '2020-01-01', '2020-01-01'),
    ('4', '12345678', '4', '2020-01-01', '2020-01-01'),
    ('5', '12345678', '5', '2020-01-01', '2020-01-01'),
    ('6', '87654321', '1', '2020-01-01', '2020-01-01'),
    ('7', '87654321', '2', '2020-01-01', '2020-01-01'),
    ('8', '87654321', '3', '2020-01-01', '2020-01-01'),
    ('9', '87654321', '4', '2020-01-01', '2020-01-01'),
    (
        '10',
        '87654321',
        '5',
        '2020-01-01',
        '2020-01-01'
    ),
    ('11', '1234124', '1', '2020-01-01', '2020-01-01'),
    ('12', '1234124', '2', '2020-01-01', '2020-01-01'),
    ('13', '1234124', '3', '2020-01-01', '2020-01-01'),
    ('14', '1234124', '4', '2020-01-01', '2020-01-01'),
    ('15', '1234124', '5', '2020-01-01', '2020-01-01'),
    (
        '16',
        '12345154',
        '1',
        '2020-01-01',
        '2020-01-01'
    ),
    (
        '17',
        '12345154',
        '2',
        '2020-01-01',
        '2020-01-01'
    ),
    (
        '18',
        '12345154',
        '3',
        '2020-01-01',
        '2020-01-01'
    ),
    (
        '19',
        '12345154',
        '4',
        '2020-01-01',
        '2020-01-01'
    );