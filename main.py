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

        return node


def search(node, target):

    if node is None:
        return None

    if node.value == target:
        return node

    if target < node.value:
        return search(node.left, target)

    return search(node.right, target)

root = Node(4)

insert(root, 3)
insert(root, 5)
insert(root, 2)
insert(root, 7)

found = search(root, 7)

print(found.value)