from django.urls import path
from . import views

urlpatterns = [
    path('', views.home, name='home'),
    path('order/', views.order, name='order'),
    path('order/success/', views.order_success, name='order_success'),
    path('contact/', views.contact, name='contact'),
    path('contact/success/', views.contact_success, name='contact_success'),
    path('gallery/', views.gallery, name='gallery'),
    path('videos/', views.videos, name='videos'),
    path('videos/upload/', views.upload_video, name='upload_video'),
]
