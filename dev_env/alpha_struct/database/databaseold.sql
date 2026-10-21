--this is for project management
CREATE TABLE atributes (
  curso   TEXT NOT NULL,
  bolsa  TEXT NOT NULL,
  limite  INTEGER NOT NULL,
 CONSTRAINT pk_atrib PRIMARY KEY (curso, bolsa, limite)
);

CREATE TABLE studentdent (
  email   TEXT PRIMARY KEY,
  curso   TEXT NOT NULL,
  bolsa  TEXT NOT NULL,
  limite  INTEGER NOT NULL,
  CONSTRAINT atributes FOREIGN KEY (curso, bolsa, limite)
    FOREIGN KEY (curso, bolsa, limite)
    REFERENCES atrib (curso, bolsa, limite)
);

CREATE TABLE edital (
  nome    TEXT PRIMARY KEY,
  curso   TEXT NOT NULL,
  bolsa  TEXT NOT NULL,
  limite  INTEGER NOT NULL,
  CONSTRAINT fk_edital_atributes
  FOREIGN KEY (curso, bolsa, limite)
    REFERENCES atrib (curso, bolsa, limite)
);

--////--

SELECT studentdent.email, edital.nome
FROM student
JOIN edital
  ON  edital.curso  = student.curso
  AND edital.limite = student.limite;

  CREATE VIEW matches AS
SELECT student.email, edital.nome
FROM student
JOIN edital
  ON  edital.curso  = student.curso
  AND edital.bolsa = student.bolsa
  AND edital.limite = student.limite;


  --verify output from tinylm before commiting to this file

  --{
  "instruction": "You will be given the file_tag and content of a Brazilian Academical notice. From the content, determine: (1) which program this applies to — EngComp, EngAut, EnsMed, or NULL if none apply; (2) whether a period/semester eligibility requirement is stated (e.g. 'estar cursando o 4º período' means term=4), otherwise NULL; (3) the edital's own release date, formatted as Year.Semester (e.g. 2024.2). From the file_tag, copy the link exactly as given — if no link is present in file_tag, use NULL. Output two SQL INSERT statements: one into atributes(degree, term), and one into edital(link, degree, term, release_date).",
  "input": "file_tag: <[certificacao-do-ensino-medio-com-base-nas-notas-encceja-edital-no-008-19-2018-proen-ifms-certificado.txt] [https://selecao.ifms.edu.br/edital/files/certificacao-do-ensino-medio-com-base-nas-notas-encceja-edital-no-008-19-2018-proen-ifms-certificado-de-conclusao-do-ensino-medio-atendimentos-no-ano-de-2021-enem-e-encceja.pdf]:>\ncontent: EDITAL Nº 008.19/2018 – PROEN/IFMS A PRÓ-REITORA DE ENSINO DO INSTITUTO FEDERAL DE EDUCAÇÃO, CIÊNCIA E TECNOLOGIA DE MATO GROSSO DO SUL, no uso das atribuições legais que lhe confere a Portaria IFMS n° 1.485, de 31 de dezembro de 2022... [rest of content, verbatim]",
  "output": "INSERT INTO atributes (degree, term) VALUES ('EnsMed', NULL);\nINSERT INTO edital (link, degree, term, release_date) VALUES ('https://selecao.ifms.edu.br/edital/files/certificacao-do-ensino-medio-com-base-nas-notas-encceja-edital-no-008-19-2018-proen-ifms-certificado-de-conclusao-do-ensino-medio-atendimentos-no-ano-de-2021-enem-e-encceja.pdf', 'EnsMed', NULL, '2024.2');"
