CREATE TABLE pets.owns (
    pid INTEGER REFERENCES pets.person(id),
    did INTEGER REFERENCES pets.dog(id),
    PRIMARY KEY (pid, did)
);
