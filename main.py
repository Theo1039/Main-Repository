class Person:
    def __init__(self, name):
        self.name = name

    class Address:
          def __init__(self, address):
             self.address = address

    address = Address("city: Port Harcourt")
person = Person("Theophilus")

print(f"name:{person.name}\n{person.address.address}")