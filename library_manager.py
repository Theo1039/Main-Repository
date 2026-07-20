import json
import random

Library_file = "library.json" # this acts as storage for library data (library and wishlist)


def load_library()-> tuple[list, list]:
    """Load the library from a JSON file. Return empty list if file doesn't exist."""
    # TODO: Use try/except to open LIBRARY_FILE and return json.load(file)
    # If FileNotFoundError, return []

    try:
        with open(Library_file, "r") as file:
            data = json.load(file)
            #Assuming we saved them as a dictionary containing both lists (library and wishlist)
            return data.get("library", []), data.get("wishlist", [])
            #get used here search for the "library" key and also "wishlist" key, if not found it returns an empty list []
    except FileNotFoundError:
        # if Library_file cannot be found return two empty list
        return [], []


def save_library(library: list, wishlist: list) -> None:
    """Save the library list to a JSON file."""
    # TODO: Open LIBRARY_FILE in 'w' mode and use json.dump(library, file, indent=2)
    save_data = {
        "library": library,
        "wishlist": wishlist
    }
    
    with open(Library_file, "w") as file:
        #json.dump saves the current state of the library and wishlist to library_file
        json.dump(save_data, file, indent=4)


def add_book(library: list):
    """Ask the user for book details and append a dict to the library list."""
    # TODO: Ask for title, author, year (convert to int), read status (y/n)
    # Append {"title": ..., "author": ..., "year": ..., "read": ..., "rating": 0}
    pass


def display_books(library: list):
    """Print all books in a nice format."""
    # TODO: Loop through library with enumerate() and print each book
    # If library is empty, print "Your library is empty!"
    pass


def mark_as_read(library: list): list:
    """Let the user pick a book by number and mark it as read + add a rating."""
    # TODO: Display books, ask for book number, set read=True, ask for rating 1-5
    pass


def search_books(library: list):
    """Search for books by title or author."""
    # TODO: Ask for a search term, loop through books, print matches
    # Hint: use .lower() for case-insensitive search
    pass


def show_statistics(library: list):
    """Show total books, percentage read, and average rating."""
    # TODO: Calculate and print stats
    pass


def random_recommendation(library: list):
    """Pick and display a random unread book."""
    # TODO: Filter unread books, use random.choice() to recommend one
    pass


def delete_book(library: list):
    """Remove a book from the library by its number."""
    # TODO: Display books, ask for number, pop() from list
    pass

def import_books(library: list, filename: str = 'books.csv'):
    pass

def export_books(library: list, filename: str = 'library_exporter.csv'):
    pass

def wishlist_menu(library: list, wishlist: list):
    pass

def print_menu():
    print("\n" + "=" * 40)
    print("    📚 PERSONAL LIBRARY MANAGER")
    print("=" * 40)
    print("1. Add a Book")
    print("2. View All Books")
    print("3. Search Books")
    print("4. Mark a Book as Read")
    print("5. Delete a Book")
    print("6. Show Statistics")
    print("7. Random Recommendation")
    print("8. Import book books from a CSV")
    print("9. Export library to CSV")
    print("10. View reading wishlist")
    print("11. Save & Exit")
    print("=" * 40)


def main():
    library, wishlist = load_library()

    while True:
        print_menu()
        choice = input("Enter your choice (1-11): ").strip()
        # strip() takes away any whitespaces (spaces, tabs, newlines) from the users input

        if choice == "1":
            add_book(library)
        elif choice == "2":
            display_books(library)
        elif choice == "3":
            search_books(library)
        elif choice == "4":
            mark_as_read(library)
        elif choice == "5":
            delete_book(library)
        elif choice == "6":
            show_statistics(library)
        elif choice == "7":
            random_recommendation(library)
        elif choice == "8":
            import_books(library)
        elif choice == "9":
            export_books(library)
        elif choice == "10":
            wishlist_menu(library, wishlist)
        elif choice == "11":
            save_library(library, wishlist)
            print("✅ Library saved! Goodbye!")
            break
        else:
            print("❌ Invalid choice. Please enter a number 1-11.")


if __name__ == "__main__":
    main()
