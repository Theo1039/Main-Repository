from django.shortcuts import render, redirect
from .forms import OrderForm, ContactForm, VideoForm
from .models import GalleryImage, Video

def home(request):
    return render(request, 'home.html')

def order(request):
    if request.method == 'POST':
        form = OrderForm(request.POST)
        if form.is_valid():
            form.save()
            return redirect('order_success')
    else:
        form = OrderForm()
    return render(request, 'order.html', {'form': form})

def order_success(request):
    return render(request, 'order_success.html')

def contact(request):
    if request.method == 'POST':
        form = ContactForm(request.POST)
        if form.is_valid():
            form.save()
            return redirect('contact_success')
    else:
        form = ContactForm()
    return render(request, 'contact.html', {'form': form})

def contact_success(request):
    return render(request, 'contact_success.html')

def gallery(request):
    images = GalleryImage.objects.all()
    return render(request, 'gallery.html', {'images': images})

def videos(request):
    videos = Video.objects.all()
    return render(request, 'videos.html', {'videos': videos})

def upload_video(request):
    if request.method == 'POST':
        form = VideoForm(request.POST, request.FILES)
        if form.is_valid():
            form.save()
            return redirect('videos')
    else:
        form = VideoForm()
    return render(request, 'upload_video.html', {'form': form})
