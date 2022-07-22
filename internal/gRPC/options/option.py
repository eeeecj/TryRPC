

import configparser

# Config 设置
class Config():
    def __init__(self,path) -> None:
        self.path=path
        self.config=self.readConfig()

    def readConfig(self):
        config=configparser.ConfigParser()
        config.read(self.path,encoding='utf-8')
        return config

    def Get(self,section,name):
        return self.config.get(section,name)

    def Set(self,section,key,value):
        self.config.set(section,key,value)

    def Save(self,path):
        with open(path) as f:
            self.config.write(f)

