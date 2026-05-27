# write file
file = open("Theo.txt", "w")
file.write("Hello! How are you?")
file.close()

file = open("Theo.txt", "a")
file.write("\nHope you are doing great!")
file.close()

file = open("Theo.txt", "r")
content = file.read()
print(content)
file.close()
print("---------------------")

with open("john.txt", "w") as file:
    file.write("My Name Is John")
    
with open("john.txt", "a") as file:
    file.write("\nadditional comment")    
    
with open("john.txt", "r") as file:
    con = file.read()
    print(con)    

#save users input
name = input("enter your name? ")
with open("users.txt", "a") as file:
    file.write(name + "\n")
    print("saved")
    
with open("users.txt", "r") as file:
    content = file.read()
    print(content)    
    
with open("john.txt", "r") as file:
    for line in file:
        print(line)
#find missing file
try:
    with open("unknown.txt", "r") as file:
        print(file.read()) 
except FileNotFoundError:
    print("file not found")               