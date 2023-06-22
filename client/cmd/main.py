import dbWork
from tabulate import tabulate 

TOKEN = 'vYq?VQ?Gvg==9cow9neCUCwc3YpoA5kmOBrdQL!/9k60iNjr3i6?RS?fKr3UDEkbNUKKdt16xhio3yJxj8-1Ql84BUmDZ4fQq5CpeyNaF?!5qiJvjfd5a?OEP0KMYo2vt/yTxpv14hfpDs768?8Et5cZzxJgd2TlkBtKzvj0pXsRzW-95iwWREFwp7wCqzOD1XRg553okTUJKpAQr7wulPLln?fMIrBid9b5x!9oCENq6Z=FtaaiNECuc5?Q6cRi'

def count_price_of_the_day(penalty, price):
    return int(price-penalty)

def printTable(data):
    table = tabulate(data, headers="keys", tablefmt="psql")
    print(table)

def createRecordFromWorkspace(token):
    nameworkspace = str(input("Enter the name of the workplace: "))
    date = str(input("Enter the date(yyyy-mm-dd): "))
    time_work = int(input("Enter the working time(8 or 4): "))
    price = int(input("Enter the price(3200 or other price): "))
    penalty = int(input("Enter the penalty(0 or other num): "))

    price = count_price_of_the_day(penalty, price)

    response = dbWork.createRecordToWorkspace(nameworkspace, date, price, time_work, penalty, token)
    
    if response.status_code == 200:
        print('Запись успешно создана')
    else:
        print('Ошибка при создании записи:', response.text, response.status_code)
def updateRecordFromWorkspace(token):
    record_id = int(input("Enter id of the record: "))
    nameworkspace = str(input("Enter the name of the workplace: "))
    date = str(input("Enter the date(yyyy-mm-dd): "))
    time_work = int(input("Enter the working time(8 or 4): "))
    price = int(input("Enter the price(3200 or other price): "))
    penalty = int(input("Enter the penalty(0 or other num): "))
    price = count_price_of_the_day(penalty, price)
    
    response = dbWork.updateRecordToWorkspace(record_id,nameworkspace, date, price, time_work, penalty, token)
    if response.status_code == 200:
        print('Запись успешно создана')
    else:
        print('Ошибка при создании записи:', response.text, response.status_code)
def deleteRecordFromWorkspace(token):
    record_id = int(input("Enter id of the record: "))  # Идентификатор записи, которую нужно удалить
    response = dbWork.deleteRecordToWorkspace(record_id, token)
    if response.status_code == 200:
        print('Запись успешно удалена')
    else:
        print('Ошибка при удалении записи:', response.text)
def getAllRecordsFromWorkspace(token):
    response = dbWork.getAllRecordsToWorkspace(token)
    if response.status_code == 200:
        print('Запись успешно получена')
    else:
        print('Ошибка при получении записи:', response.text)
        return
    printTable(response.json())

def createRecordFromListPayments(token):
    date          = str(input("Enter date(yyyy-mm-dd): "))
    nameworkspace = str(input("Enter the name of the workplace: "))
    price         = int(input("Enter price: "))

    response = dbWork.createRecordToListPayments(nameworkspace, date, price, token)
    if response.status_code == 200:
        print('Запись успешно создана')
    else:
        print('Ошибка при создании записи:', response.text, response.status_code)
def updateRecordFromListPayments(token):
    record_id     = int(input("Enter id of the record: "))
    nameworkspace = str(input("Enter the name of the workplace: "))
    date          = str(input("Enter date(yyyy-mm-dd): "))
    price         = int(input("Enter price: "))

    response = dbWork.updateRecordToListPayments( record_id, nameworkspace, date, price, token)
    if response.status_code == 200:
        print('Запись успешно создана')
    else:
        print('Ошибка при создании записи:', response.text, response.status_code)
def deleteRecordFromListPayments(token):
    record_id = int(input("Enter id of the record: ")) # Идентификатор записи, которую нужно удалить

    response = dbWork.deleteRecordToListPayments( record_id, token)
    if response.status_code == 200:
        print('Запись успешно удалена')
    else:
        print('Ошибка при удалении записи:', response.text)
def getAllRecordsFromListPayments(token):
    response = dbWork.getAllRecordsToListPayments(token)
    if response.status_code == 200:
        print('Запись успешно получена')
    else:
        print('Ошибка при получении записи:', response.text)
        return
    printTable(response.json())
    
def workspaceTable(token_):
    action = int(input("Select action (1. Create new record, 2. Update record, 3. Delete record, 4. Get all records): "))
    if action==1:
        createRecordFromWorkspace(token_)
    elif action==2:
        updateRecordFromWorkspace(token_)
    elif action==3:
        deleteRecordFromWorkspace(token_)
    elif action==4:
        getAllRecordsFromWorkspace(token_)

def listPaymentsTable(token_):
    action = int(input("Select action (1. Create new record, 2. Update record, 3. Delete record, 4. Get all records): "))
    if action==1:
        createRecordFromListPayments(token_)
    elif action==2:
        updateRecordFromListPayments(token_)
    elif action==3:
        deleteRecordFromListPayments(token_)
    elif action==4:
        getAllRecordsFromListPayments(token_)


if __name__=="__main__":
    token = str(input("Enter token: "))
    if TOKEN!=token:
        print("Иди на хуй")
        exit()
    while(True):
        id_table = int(input("Enter the table number: "))
        if id_table==1:
            workspaceTable(token)
        elif id_table==2:
            listPaymentsTable(token)
        
        if str(input("Do you want to repeat?(y/n): " ))=="n":
            exit()
