
books = []
def book():
    name = input("input the of the author: ")
    title = input("input book title: ")
    date = input("date: ")
    
    library = {
        "name":name,
        "title":title,
        "date":date
    }
    books.append(library)
    for book in books:
        for key, value in book.items():
            print(f"{key}: {value}")
book()
animals = [
    {
    "name":"john",
    "age":45
},
    {
        "name":"errorty",
        "age":9
    },
    {
        "name":"Grace",
        "age":8
    }
]
for animal in animals:
    for key,value in animal.items():
        print(f"{key} : {value}")
        
        
go = ["we","he","red"]        
for g in go:
    print(g)
dic = {"got":"road","went":"world"}    
print(len(dic))
for d in dic:
    for k,v in dic.items():
        print(f"{k} : {v}")    