def count_down(num):
    if num == 0:
        return
    print(num)
    count_down(num - 1)
count_down(5)    
print("'''''''''''''''''''")
def ascending_count(n):
    if n == 5:
        return print(f"number reach the end: {n - 1}")
    print(n)
    return ascending_count(n + 1)
ascending_count(1)

box = []
def fill_box(n):
    global box
    if len(box) == 5:
        return f"box filled"
    #print(box)
    box.append(n)
    fill_box(n + 1)
fill_box(1)  
print(box)  

print("+++++++++++++++++++++++++++")

even_num = []
odd_num = []
prime_num = []

def sep_max_min(num):
    global even_num
    global odd_num
    global prime_num
    if num < 1:
        return
    if num % 2 == 0:
       even_num.append(num)
        
    elif num % 3 == 0:
        odd_num.append(num)  
    elif num % 2 and num % 3 != 0:
        prime_num.append(num)  
    sep_max_min(num - 1)
print("To check Prime, Odd, or Even numbers")    
number = int(input("Enter a number: "))       
sep_max_min(number)  
print(f"Even numbers: {even_num}")
print(f"Odd numbers: {odd_num}") 
print(f"Prime numbers: {prime_num}")  

        
    

    