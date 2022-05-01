# Smart Hybrid Workforce Manager

# The team

## Kaylee Posthumus (Team Lead) - u19061359

<img src="https://drive.google.com/uc?export=view&id=1V-Oyk261MFmbf28FgBKZpwe-EE-AiSss">

I am a 3rd year Computer Science Student at the University of Pretoria. I have worked part time(16 hours a week and full time on holidays) for 5DT for one and a half years where I have worked on large web based applications as well as some network based system applications. I have been tutoring for the University for two years in the Computer Science Department.

[LinkedIn](https://www.linkedin.com/in/kaylee-posthumus-1a538b238/)

## Estian Nel - u20427736

<img src="https://drive.google.com/uc?export=view&id=1uUj6kdns3AWIpZN5ZxGj54VrpyB_HGx0">

I am a 3rd year BSc Comp Sci student at the University of Pretoria and I plan on starting to work after my degree and work part time as I am doing my honours degree. I love learning new technologies and I have a large interest in machine learning as well as back-end development.

[LinkedIn](https://www.linkedin.com/in/estian-nel-061296238/)

## Thashil Naidoo - u20491141

<img src="https://drive.google.com/uc?export=view&id=1D6o0HPr1TWjjVaanTh1hqPiEEQ5B9sKQ">

I am a 3rd year BSc Computer Science student. I have a great interest in both computer graphics and AI. After completing my honours next year, I plan on working full time as a software engineer. I enjoy challenging myself to learn new concepts as well as encouraging others to always do their best.

[LinkedIn](https://www.linkedin.com/in/thashilnaidoo/)

## Ryan Healy - u20662302

<img src="https://drive.google.com/uc?export=view&id=1gIRI5IcjOQO77UfTToQc8ofYxbYFnq6A">

I am a 3rd year BSc Computer Science student. I am passionate about artificial intelligence and mathematics. I plan on completing my honours degree next year before working as a software engineer.

[LinkedIn](https://www.linkedin.com/in/ryan-healy-6a4389238/)

## Jonathn Enslin - u19103345

<img src="https://drive.google.com/uc?export=view&id=1Zc33pK4GaZny3IgI6VN6tRFqKzzv6joJ">

I am a 3rd year BSc Information and Knowledge systems student, specialising in data science. I have a great interest in AI and theoretical computer science, and a thorough understanding and intuition in the fields of mathematics, and physics.

[LinkedIn](https://www.linkedin.com/in/jonathan-enslin-947293238/)

# Functional Requirements

# Project Board

# Demo Recordings

# Starting up development environment

Please note that the first time `docker-compose up` is run it will error out and have to be re-run this is because the database has to be created and the api depends on the db creation to finish.

## Building docker containers

    docker-compose build

## Running docker containers

    docker-compose up

## Stopping docker containers

    docker-compose down

# API

## Documentation

Documentation in the form of a postman collection for the api is located in `documentation/api`

## Restarting the API

- Windows
  docker-compose down; cls; docker-compose up --build
- Linux
  docker-compose down; reset; docker-compose up --build

# Postgres

## Restarting Postgres

- Windows
  docker-compose down; cls; docker-compose up --build
- Linux
  docker-compose down; reset; docker-compose up --build

## Removing the volume

    rm -r db/postgres-data/
