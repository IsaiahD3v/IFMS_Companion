

  CREATE TABLE atributes (
  degree TEXT NOT NULL,
  term   INTEGER,
  CONSTRAINT uq_atributes UNIQUE NULLS NOT DISTINCT (degree, term)
); --uq_atributes ensures that there is only one row for each (degree, term) pair, even if term is NULL.
   -- maybe add if this is paid also 
CREATE TABLE edital (
  link         TEXT PRIMARY KEY,
  degree       TEXT NOT NULL,
  term         INTEGER,
  release_date TEXT NOT NULL,  -- 'Year.Semester', e.g. '2024.2'
  CONSTRAINT fk_edital_atributes
    FOREIGN KEY (degree, term) REFERENCES atributes (degree, term)
);

CREATE TABLE student (
  email  TEXT PRIMARY KEY,
  degree TEXT NOT NULL,
  term   INTEGER,
  CONSTRAINT fk_student_atributes
    FOREIGN KEY (degree, term) REFERENCES atributes (degree, term)
);

CREATE VIEW matches AS
SELECT s.email, e.link, e.degree, e.term, e.release_date
FROM student s
JOIN edital e
  ON  e.degree = s.degree
  AND e.term IS NOT DISTINCT FROM s.term;

CREATE TABLE notified (
  email   TEXT NOT NULL REFERENCES student (email),
  link    TEXT NOT NULL REFERENCES edital (link),
  sent_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  PRIMARY KEY (email, link)
);

-- why created a table here