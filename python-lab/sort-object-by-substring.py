class Item:
    sku = ""
    quantity = 0
    name = ""

    def __init__(self, sku, quantity, name):
        self.sku = sku
        self.quantity = quantity
        self.name = name


items = [Item("sku-101-XL-001", 10, "apel"),
         Item("sku-102-MM-002", 3, "box"),
         Item("sku-103-MM-003", 5, "box"),
         Item("sku-104-LL-004", 8, "box"),
         Item("sku-105-LL-005", 15, "can")]

itemWeight = dict(MM=1, LL=2, XL=3)

# for item in items:
#     print item.sku

items.sort(key=lambda i: itemWeight[i.sku.split("-")[2]], reverse=True)

for item in items:
    print item.sku, " and ", item.quantity, " and ", item.name
