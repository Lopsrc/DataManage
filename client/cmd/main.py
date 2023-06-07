

import dbWork

fix_price_of_the_day = 3200
TOKEN = 'J/79d4=Yutb7J!YYgMBvsH12PyxAa197v1CTkWkv4QBy?T4UHDUEiexn0t1HVBNAQ-9429md/8!hmkFZOV!9oeyGbwo0q0mDUYEa7cPloIFu8DjDLj=eKQoOQONPKywwOv?MQtv!rkNWVoNUEv2sTwY3HOxeUUBHeOtXD-voZ12vD3pOZQm6VcspJa7jhuCloAx-unzh?0gXXMVVGsjMZc=eKH2LG!5SOEQ3Xy8BxlccLACoHRB2Df-njeMbJ79a'

class Workspace:
    def __init__(self, id: int,
                nameworkspace: str,
                date: str,
                price: int,
                timework: int,
                penalty: int,) -> None:
        self.id = id
        self.nameworkspace = nameworkspace
        self.date = date
        self.price = price
        self.timework = timework
        self.penalty = penalty

class ListPayments:
    def __init__(self, id: int,
                nameworkspace: str,
                date: str,
                price: int,) -> None:
        self.id = id
        self.nameworkspace = nameworkspace
        self.date = date
        self.price = price

def count_price_of_the_day(time_work, penalty):
    if time_work <= 4:
        return int(fix_price_of_the_day/2-penalty)
    return int(fix_price_of_the_day-penalty)

def getRecordId():
    pass


def workspaceTable(token_):
    action = int(input("Select action (1. Create new record, 2. Update record, 3. Delete record, 4. Get all records): "))
    if action==1:
        nameworkspace = str(input("Enter the name of the workplace: "))
        date = str(input("Enter the date(yyyy-mm-dd): "))
        time_work = int(input("Enter the working time(8 or 4): "))
        penalty = int(input("Enter the penalty(0 or other num): "))
        price = count_price_of_the_day(time_work, penalty)

        response = dbWork.createRecordToWorkspace(nameworkspace, date, price, time_work, penalty, token_)
        if response.status_code == 200:
            print('Запись успешно создана')
        else:
            print('Ошибка при создании записи:', response.text, response.status_code)
    elif action==2:
        record_id = int(input("Enter id of the record: "))
        nameworkspace = str(input("Enter the name of the workplace: "))

        date = str(input("Enter date(yyyy-mm-dd): "))
        time_work = int(input("Enter time of work(8 or 4): "))
        penalty = int(input("Enter penalty(0 or other num): "))
        price = count_price_of_the_day(time_work, penalty)
        
        response = dbWork.updateRecordToWorkspace(record_id,nameworkspace, date, price, time_work, penalty, token_)
        if response.status_code == 200:
            print('Запись успешно создана')
        else:
            print('Ошибка при создании записи:', response.text, response.status_code)
    elif action==3:
        record_id = int(input("Enter id of the record: "))  # Идентификатор записи, которую нужно удалить

        response = dbWork.deleteRecordToWorkspace(record_id, token_)
        if response.status_code == 200:
            print('Запись успешно удалена')
        else:
            print('Ошибка при удалении записи:', response.text)
    elif action==4:
        response = dbWork.getAllRecordsToWorkspace(token_)
        if response.status_code == 200:
            print('Запись успешно получена')
        else:
            print('Ошибка при получении записи:', response.text)
            return
        print(response.json())


def listPaymentsTable(token_):
    action = int(input("Select action (1. Create new record, 2. Update record, 3. Delete record, 4. Get all records): "))
    if action==1:
        date = str(input("Enter date(yyyy-mm-dd): "))
        nameworkspace = str(input("Enter the name of the workplace: "))
        price = int(input("Enter price of the day(3200): "))

        response = dbWork.createRecordToListPayments(nameworkspace, date, price, token_)
        if response.status_code == 200:
            print('Запись успешно создана')
        else:
            print('Ошибка при создании записи:', response.text, response.status_code)
    elif action==2:
        record_id = int(input("Enter id of the record: "))
        nameworkspace = str(input("Enter the name of the workplace: "))
        date = str(input("Enter date(yyyy-mm-dd): "))
        price = int(input("Enter price of the day(3200): "))

        response = dbWork.updateRecordToListPayments( record_id, nameworkspace, date, price, token_)
        if response.status_code == 200:
            print('Запись успешно создана')
        else:
            print('Ошибка при создании записи:', response.text, response.status_code)
    elif action==3:
        record_id = int(input("Enter id of the record: ")) # Идентификатор записи, которую нужно удалить

        response = dbWork.deleteRecordToListPayments( record_id, token_)
        if response.status_code == 200:
            print('Запись успешно удалена')
        else:
            print('Ошибка при удалении записи:', response.text)
    elif action==4:
        response = dbWork.getAllRecordsToListPayments(token_)
        if response.status_code == 200:
            print('Запись успешно получена')
        else:
            print('Ошибка при получении записи:', response.text)
            return
        print(response.json())


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
