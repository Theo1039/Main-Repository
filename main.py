def add(a, b):
    return a + b

def subtract(a, b):
    return a - b

def multiplication(a, b):
    return a * b

def division(a, b):
    if b != 0:
        return a / b
    else:
        return "Cannot divide by zero"


while True:

    print("\n*** OPTIONS ***")
    print("1. Addition")
    print("2. Subtraction")
    print("3. Multiplication")
    print("4. Division")
    print("5. Quit")

    choice = input("Option: ")

    if choice == "5":
        print("Goodbye")
        break

    elif choice in ["1", "2", "3", "4"]:

        print("Valid option")

        num1 = int(input("Input first number: "))
        num2 = int(input("Input second number: "))

        if choice == "1":
            print("Answer =", add(num1, num2))

        elif choice == "2":
            print("Answer =", subtract(num1, num2))

        elif choice == "3":
            print("Answer =", multiplication(num1, num2))

        elif choice == "4":
            print("Answer =", division(num1, num2))

    else:
        print("Invalid input")