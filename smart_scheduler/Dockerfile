FROM python:3.10


WORKDIR /smart-scheduler

COPY requirements.txt requirements.txt

RUN pip3 install -r requirements.txt

COPY . .

CMD ["python", "smart_scheduler.py"]
