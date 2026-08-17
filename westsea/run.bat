@echo off
REM Activate virtual environment
call .venv\Scripts\activate

REM Run Django's development server
python manage.py runserver
