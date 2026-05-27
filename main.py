while True:
    print("\n1.  Add Student")
    print("2.  View Student")
    print("3.  Exit")
    
    choice = input("Option: ")
    if choice == "1":
        name = input("enter name: ")
        with open("student.txt", "a") as file:
            file.write(name + "\n")
    elif choice == "2":
        try:
            with open("student.txt", "r") as file:
                print(file.read())
        except FileNotFoundError:
            print("no student record found")
    elif choice == "3":
        break   
   
    else:
        print("invalid input")     
"""     
while True:
    print("\n1. Add Student")
    print("2. View Student")
    print("3. Exit")

    choice = input("Option: ")

    if choice == "1":
        name = input("Enter name: ")

        with open("student.txt", "a") as file:
            file.write(name + "\n")

    elif choice == "2":
        try:
            with open("student.txt", "r") as file:
                print(file.read())

        except FileNotFoundError:
            print("No student record found")

    elif choice == "3":
        break

    else:
        print("Invalid option")         
"""                       