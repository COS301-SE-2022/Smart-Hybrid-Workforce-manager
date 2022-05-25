CREATE TABLE team (
    team_id int,
    team_name varchar(255),
    user_member int,
    PRIMARY KEY(team_id, user_member),
    CONSTRAINT user_member
        FOREIGN KEY(user_member)
            REFERENCES "users"(user_id)
);

INSERT INTO team (team_id, team_name, user_member) VALUES
(1, 'CSA', 123);