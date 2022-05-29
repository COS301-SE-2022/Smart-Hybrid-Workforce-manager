# syntax=docker/dockerfile:1

FROM python:3.10-slim-buster

WORKDIR /smart-scheduler

COPY requirements.txt requirements.txt

RUN pip3 install -r requirements.txt

COPY . .
