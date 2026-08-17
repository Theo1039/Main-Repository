from django.shortcuts import render

def home(request):
    result = None
    if request.method == "POST":
        num1 = request.POST.get('num1')
        num2 = request.POST.get('num2')
        operation = request.POST.get('operation')

        try:
            num1 = float(num1)
            num2 = float(num2)
            if operation == "add":
                result = num1 + num2
            elif operation == "subtract":
                result = num1 - num2
            elif operation == "multiply":
                result = num1 * num2
            elif operation == "divide":
                result = num1 / num2
        except:
            result = "Error"

    return render(request, "calculator/home.html", {"result": result})