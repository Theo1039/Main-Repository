import json

with open("pra.json", "r") as file:
    data = json.load(file)

    for key, value in data.items():
        print(f"Section: {key}")
        print(f"Name: {value['name']}")
        print(f"Age: {value['age']}")
        print(f"City: {value['city']}")
        print()
        
der = [1,2,3,5,78,9,89,8]
iteration = iter(der)
print(next(iteration))
print(next(iteration))
print(next(iteration))

def add():
    print("Theo")
def aws():
    print("Hello")
    
df = [add,aws]
df[0]()

def addno(a,b):
    return a + b   
def result(func):
    func
result(print(addno(2,4)))

glo = 56
def va():
    c = "Theo"
    print(c)
    global glo
    glo = 100
    print(glo)
va()    
print(glo)

def seed(name):
    print(f"my name is {name}")
    def greet():
        print("hello")
    greet()
seed("Theophilus")        
    

        