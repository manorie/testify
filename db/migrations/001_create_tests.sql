CREATE TABLE IF NOT EXISTS Tests (
  Title             varchar(255),
  Author            varchar(255),
  Path              varchar(255) PRIMARY KEY NOT NULL,
  SecretKey         varchar(255) NOT NULL,
  AuthorEmail       varchar(255) NOT NULL,
  IsPublished       boolean,
  TimeLimit         integer NOT NULL,
  AnswerSize        integer NOT NULL,
  CreationDate      datetime NOT NULL,
  ExpirationDate    datetime NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS TestPathsIndex ON Tests(Path);
