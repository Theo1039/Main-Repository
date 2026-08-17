from django.contrib import admin
from .models import Order, GalleryImage, Contact, Video

@admin.register(Order)
class OrderAdmin(admin.ModelAdmin):
    list_display = ('id', 'name', 'email', 'product', 'quantity', 'created_at')
    list_filter = ('created_at',)
    search_fields = ('name', 'email', 'product')

@admin.register(GalleryImage)
class GalleryImageAdmin(admin.ModelAdmin):
    list_display = ('id', 'title', 'uploaded_at')
    list_filter = ('uploaded_at',)
    search_fields = ('title',)

@admin.register(Contact)
class ContactAdmin(admin.ModelAdmin):
    list_display = ('id', 'name', 'email', 'created_at')
    list_filter = ('created_at',)
    search_fields = ('name', 'email')

@admin.register(Video)
class VideoAdmin(admin.ModelAdmin):
    list_display = ('id', 'title', 'uploaded_at')
    list_filter = ('uploaded_at',)
    search_fields = ('title',)
