import json

with open("data.text", "w") as f:
    f.write("Theophilus is my name")
    
with open("data.text", "r") as f:
    print(f.read())
    
    user = {
       "Name":"theo",
       "Age":78
    }
    
with open("d.text", "w")as f:
        json.dump(user, f, indent=4)
        
with open("d.text", "r")as f:
    d = json.load(f)
    print(d)        