from flask import Flask, render_template,request
import sqlite3
from flask import Response

app = Flask(__name__)
@app.route("/", methods=["GET"])
def database():
    return render_template("index.html")

@app.route("/home",methods=["GET"])
def home():
    return render_template("home.html")  

@app.route("/form",methods=["GET"])
def form():
     return render_template("form.html")

@app.route("/post",methods=["POST"])
def post():
    name = requst.form.get("name")  
    email = request.form.get("email") 
    conn = sqlite3.connect("user.db")
    cursor = conn.cursor()
    cursor.execute("""CREATE TABLE IF NOT EXISTS user(
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL
    )""")

    cursor.execute("INSERT INTO user(name,email)VALUES(?,?)", name,email)
    return Response("save to database successfull",
    status="200",
    mimetype="text/plain"
    )


if __name__== "__main__":
        app.run(debug=True)
