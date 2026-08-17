# 🧮 Django Calculator

A simple web-based calculator built with **Django**.  
Users can input two numbers, choose an operation (Add, Subtract, Multiply, Divide), and get the result displayed instantly.

---

## 🚀 Features
- Perform basic arithmetic operations  
- Error handling for invalid inputs  
- Clean and simple user interface with Django templates  

---

## 📂 Project Structure
myproject/  
│  
├── calculator/   # App with views, templates, urls  
│   ├── templates/  
│   │   └── calculator/  
│   │       └── home.html  
│   ├── views.py  
│   ├── urls.py  
│   └── models.py  
│  
├── myproject/    # Main project settings  
│   ├── settings.py  
│   ├── urls.py  
│   └── wsgi.py  
│  
├── db.sqlite3    # Database (ignored in .gitignore)  
├── manage.py  
└── README.md  

---

## ⚙️ Installation & Setup

1. Clone the repository  
   `git clone https://github.com/Theo1039/django_calculator.git`  
   `cd django_calculator`  

2. Create & activate a virtual environment  
   - On Windows: `venv\Scripts\activate`  
   - On Mac/Linux: `source venv/bin/activate`  

3. Install dependencies  
   `pip install django`  

4. Run migrations  
   `python manage.py migrate`  

5. Start the development server  
   `python manage.py runserver`  

6. Open in browser  
   `http://127.0.0.1:8000/`  

---

## 📜 License
This project is licensed under the MIT License.  

---

## ✨ Author
**Theophilus Kpurugbara**  
📧 theophiluskpurugbara@gmail.com  
🔗 [GitHub](https://github.com/Theo1039)  
