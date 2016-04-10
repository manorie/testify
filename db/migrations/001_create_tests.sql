CREATE TABLE IF NOT EXISTS Tests (
  Id                integer PRIMARY KEY AUTOINCREMENT,
  Title             varchar(255),
  Author            varchar(255),
  Path              varchar(255),
  SecretKey         varchar(255) NOT NULL,
  AuthorEmail       varchar(255) NOT NULL,
  IsPublished       boolean,
  TimeLimit         integer NOT NULL,
  AnswerSize        integer NOT NULL,
  CreationDate      integer NOT NULL,
  ExpirationDate    integer NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS TestPathsIndex ON Tests(Path);
