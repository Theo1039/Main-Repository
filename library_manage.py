import json

# Data Storage: A list to hold our book dictionaries
library = []

def add_book():
    """Function to add a new book record."""
    try:
        # User Input
        title = input("Enter book title: ")
        author = input("Enter author name: ")
        # Using a Dictionary to store the book record
        book = {"title": title, "author": author}
        library.append(book)
        # f-string for formatted display
        print(f"Book '{title}' added successfully!")
    except Exception as e:
        # Error Handling
        print(f"An error occurred: {e}")

def show_books():
    """Function to display all books using f-strings."""
    print("\n--- Current Library ---")
    for book in library:
        print(f"Title: {book['title']} | Author: {book['author']}")

# Control Flow: Using a while loop for the main menu
while True:
    print("\n1. Add Book\n2. Show Books\n3. Exit")
    choice = input("Select an option: ")
    
    # Control Flow: if/elif/else statements
    if choice == '1':
        add_book()
    elif choice == '2':
        show_books()
    elif choice == '3':
        print("Goodbye!")
        break
    else:
        print("Invalid choice, please try again.")
