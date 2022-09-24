from os.path import exists
from typing import Set

class Setup:
    def __init__(self):
        self.env_file = None
    def run(self):
        self.setEnviorment()
        
    def setEnviorment(self):
        #POSTGRES
        user = input("Please enter user for postgress db (defualt: admin)")
        password = input("Please enter password for postgress db (defualt: admin)")
        db = input("Please enter db for postgress db (defualt: arche)")
        #API Notifications
        #Redis Database
        
if __name__ == '__main__':
    setup = Setup()
    setup.run()