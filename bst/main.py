class Node:
    def __init__(self, value):
        self.value = value
        self.left = None
        self.right = None

def insert(node, value):
    if node is None:
        return Node(value)

    if value < node.value:
        node.left = insert(node.left, value)
    elif value > node.value:
        node.right = insert(node.right, value)

        return node  # Crucial: This must be un-indented so the function always returns the node

    # Visual helper function to print the tree layout
def print_tree(node, level=0, prefix="Root: "):
    if node is not None:
        print_tree(node.right, level + 1, "R---- ")
        print(" " * 4 * level + prefix + str(node.value))
        print_tree(node.left, level + 1, "L---- ")

        # Initialize and build the tree (Moved outside the insert function)
root = None
root = insert(root, 10)
root = insert(root, 5)
root = insert(root, 15)
root = insert(root, 3)
root = insert(root, 7)

# Run the print function
print_tree(root)
