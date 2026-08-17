from django import forms
from .models import Order, Contact, Video

class OrderForm(forms.ModelForm):
    class Meta:
        model = Order
        fields = ['name', 'email', 'product', 'quantity']

class ContactForm(forms.ModelForm):
    class Meta:
        model = Contact
        fields = ['name', 'email', 'message']

class VideoForm(forms.ModelForm):
    class Meta:
        model = Video
        fields = ['title', 'video_file']
