count = 0
def hello():
    global count
    if count == 4:
        return
    print("Hello")
    count +=1
    hello()
hello()    

def add(n):
   if n == 0: 
       return
   print(n)
   add(n-1)
add(5)    
print(".....................................")
def num(n):
    
    if n== 0:
        return
    num(n-1)
    print(n)
    
nu = int(input("input number: "))
    
num(nu)         

def sum_num(n):
    if n == 0:
        return 0
    print(n)
    return n + sum_num(n - 1)
num = int(input("input a number: "))
sum_num(num)