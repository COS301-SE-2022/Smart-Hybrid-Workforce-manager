import docker
import requests

# from flask import Flask
# app = Flask(__name__) # Flask instance named app

# client = docker.from_env()

# print("Containers")
# for container in client.containers.list():
#   print(container.attrs['NetworkSettings']['IPAddress'])

response = requests.post("http://arche-api:8080/api/user/information", data="{}")

print(response.content)
