print("\n*** WELCOME TO MULTI FILE HANDLER ***")

name = input("Enter your name: ")
age = input("Enter your age: ")
country = input("Enter your country name: ")

user = f"{name}, {age}, {country}\n"

with open("info.txt", "a") as file:
    file.write(user)

search = input("Enter name to search: ")

found = False

with open("info.txt", "r") as file:
    for line in file:
        cleanName = line.strip().split(", ")

        if cleanName[0] == search:
            print(f"Found: {line}")
            found = True

if found == False:
    print("Name not found")

delete_name = input("Enter delete name: ")

students = []
deleted = False

with open("info.txt", "r") as file:
    for line in file:
        file_name = line.strip().split(", ")

        if file_name[0] != delete_name:
            students.append(line)
        else:
            deleted = True

with open("info.txt", "w") as file:
    for student in students:
        file.write(student)

if deleted:
    print("Name deleted")
else:
    print("Name not found")

with open("info.txt", "r") as file:
    for line in file:
        data = line.strip().split(", ")
        print(data)
        