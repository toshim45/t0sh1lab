import requests
import threading
import logging
import time
import datetime
import os

url2 = "http://localhost:9602/product/update"
url3 = "http://localhost:9603/product/update"
url4 = "http://localhost:9604/product/update"
maxLoop = int(os.getenv('MAX_LOOP',5))
maxParalel = int(os.getenv('MAX_PARALEL',100))

payload = {
    "action": {"name": "edit_product"},
    "input": {
        "id": "0018e493-aa8c-4412-8b34-80e954ac9438",
        "company_id": "3993d009-0032-4e18-8e9f-3fb9ae0e0c3a",
        "tenant_id": "8bc82521-f6a0-4143-a7ce-77a9e15b40b2",
        "sku": "7113-5343",
        "name": "product-7113-5343",
        "description": "test product 7113-5343-1",
        "type": "inventory",
        "shelf_life_management": "disabled",
        "lot_management": "disabled",
        "serial_number_management": "disabled",
        "uom_conversion_management": "disabled",
        "image_url": "path-to-image-7113-5343-1",
        "status": "active"
    },
    "session_variables": {
        "x-hasura-user-id": "06daa23c-5848-45f8-b714-013e98eb502f",
        "x-hasura-company-id": "3993d009-0032-4e18-8e9f-3fb9ae0e0c3a",
        "x-hasura-warehouse-id": "2e73984a-9fae-44e4-be69-e645bac0274c"
    }
}
headers = {"Content-Type": "application/json"}

def hit(idx):
    _url = url2
    if (idx % 3) == 0:
        _url = url4
    elif (idx % 2) == 0:
        _url = url3
        
    response = requests.request("POST", _url, json=payload, headers=headers)
    logging.info("[HIT] %s %s result: %s", idx, _url, response.status_code)


if __name__ == "__main__":
    format = "%(asctime)s: %(message)s"
    logging.basicConfig(format=format, level=logging.INFO,
                        datefmt="%H:%M:%S")

    threads = list()
    startTime = time.time()
    for l in range(maxLoop):
        for index in range(maxParalel):
            # logging.info("Main    : create and start thread %d.", index)
            x = threading.Thread(target=hit, args=(l*maxParalel+index,))
            threads.append(x)
            x.start()

        for index, thread in enumerate(threads):
            thread.join()
    endTime = time.time()

    durationTime = datetime.timedelta(seconds=endTime-startTime)
    logging.info("[FINISHED] in %s", durationTime)