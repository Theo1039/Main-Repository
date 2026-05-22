print("simple calculator")

n1 = int(input("enter num1: "))
op = input("enter operator[+,-,*,/] ")
n2 = int(input("enter num2: "))

if op == '+':
    result = n1 + n2
    print(result)
elif op == '-':
     result = n1 - n2
     print(result)       
elif op == '/':
    if n2 != 0:
        result = n1 / n2
        print(result)
    else: 
        print("cannot be divided by zero")    
        
elif op == '*':
     print(n1 * n2)   

else:
    print("invalid operation")          