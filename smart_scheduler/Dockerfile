FROM python:3.10


WORKDIR /smart-scheduler

COPY requirements.txt requirements.txt

RUN pip3 install -r requirements.txt

COPY . .

EXPOSE 8081

# CMD ["python", "smart_scheduler.py"]
CMD ["python", "scheduler_endpoint.py"]
