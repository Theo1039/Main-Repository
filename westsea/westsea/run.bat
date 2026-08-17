@echo off
REM Activate virtual environment
call .venv\Scripts\activate

REM Install dependencies
pip install -r requirements.txt

REM Apply database migrations
python manage.py migrate

REM Collect static files
python manage.py collectstatic --noinput

REM Start Gunicorn equivalent for local testing (using Django dev server)
echo Starting Django development server...
python manage.py runserver 0.0.0.0:8000

