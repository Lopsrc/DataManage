

import dbWork

fix_price_of_the_day = 3200
TOKEN = 'valid_token'

class Workspace:
    def __init__(self, id: int,
                nameworkspace: str,
                date: str,
                price: int,
                timework: int,
                penalty: int) -> None:
        self.id = id
        self.nameworkspace = nameworkspace
        self.date = date
        self.price = price
        self.timework = timework
        self.penalty = penalty

class ListPrice:
    def __init__(self, id: int,
                date: str,
                price: int,) -> None:
        self.id = id
        self.date = date
        self.price = price

class ListPayments:
    def __init__(self, id: int,
                date: str,
                price: int,) -> None:
        self.id = id
        self.date = date
        self.price = price
# class CreateData:
#     id_table, token, date, price, time_work, penalty, price_day, payments = None
#     def __init__(self, token, date=None, price=None, time_work=None, penalty=None, price_day=None, payments=None):
#         pass
    

#нужно реализовать функции для получения данных из таблиц и выбора конкретного id 



# Пример использования
# date = '2023-07-04'
# price = 100
# time_work = 8
# penalty = 10
# price_day = 90
# payments = 1000


# Создание записи
# response = create_entry(date, price, time_work, penalty, price_day, payments, token)
# if response.status_code == 200:
#     print('Запись успешно создана')
# else:
#     print('Ошибка при создании записи:', response.text, response.status_code)

# Обновление записи
# entry_id = 7  # Идентификатор записи, которую нужно обновить
# new_date = '2023-06-02'
# new_price = 150
# new_time_work = 7
# new_penalty = 5
# new_price_day = 120
# new_payments = 80000
# response = update_entry(entry_id, new_date, new_price, new_time_work, new_penalty, new_price_day, new_payments, token)
# if response.status_code == 200:
#     print('Запись успешно обновлена')
# else:
#     print('Ошибка при обновлении записи:', response.text)

# Удаление записи
# entry_id = 3  # Идентификатор записи, которую нужно удалить
# response = delete_entry(entry_id, token)
# if response.status_code == 200:
#     print('Запись успешно удалена')
# else:
#     print('Ошибка при удалении записи:', response.text)


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
            print('Запись успешно удалена')
        else:
            print('Ошибка при удалении записи:', response.text)
            return
        print(response.json())



def listPriceTable(token_):
    action = int(input("Select action (1. Create new record, 2. Update record, 3. Delete record, 4. Get all records): "))
    if action==1:
        date = str(input("Enter date(yyyy-mm-dd): "))
        price = int(input("Enter price of the day(3200): "))

        response = dbWork.createRecordToListPrices(date, price, token_)
        if response.status_code == 200:
            print('Запись успешно создана')
        else:
            print('Ошибка при создании записи:', response.text, response.status_code)
    elif action==2:
        record_id = int(input("Enter id of the record: "))
        date = str(input("Enter date(yyyy-mm-dd): "))
        price = int(input("Enter price of the day(3200): "))

        response = dbWork.updateRecordToListPrices( record_id, date, price, token_)
        if response.status_code == 200:
            print('Запись успешно создана')
        else:
            print('Ошибка при создании записи:', response.text, response.status_code)
    elif action==3:
        record_id = int(input("Enter id of the record: "))  # Идентификатор записи, которую нужно удалить

        response = dbWork.deleteRecordToListPrices(record_id, token_)
        if response.status_code == 200:
            print('Запись успешно удалена')
        else:
            print('Ошибка при удалении записи:', response.text)
    elif action==4:
        response = dbWork.getAllRecordsToListPrice(token_)
        if response.status_code == 200:
            print('Запись успешно удалена')
        else:
            print('Ошибка при удалении записи:', response.text)
            return
        print(response.json())


def listPaymentsTable(token_):
    action = int(input("Select action (1. Create new record, 2. Update record, 3. Delete record, 4. Get all records): "))
    if action==1:
        date = str(input("Enter date(yyyy-mm-dd): "))
        price = int(input("Enter price of the day(3200): "))

        response = dbWork.createRecordToListPayments(date, price, token_)
        if response.status_code == 200:
            print('Запись успешно создана')
        else:
            print('Ошибка при создании записи:', response.text, response.status_code)
    elif action==2:
        record_id = int(input("Enter id of the record: "))
        date = str(input("Enter date(yyyy-mm-dd): "))
        price = int(input("Enter price of the day(3200): "))

        response = dbWork.updateRecordToListPayments( record_id, date, price, token_)
        if response.status_code == 200:
            print('Запись успешно создана')
        else:
            print('Ошибка при создании записи:', response.text, response.status_code)
    elif action==3:
        record_id = int(input("Enter id of the record: ")) # Идентификатор записи, которую нужно удалить

        response = dbWork.deleteRecordToListPrices( record_id, token_)
        if response.status_code == 200:
            print('Запись успешно удалена')
        else:
            print('Ошибка при удалении записи:', response.text)
    elif action==4:
        response = dbWork.getAllRecordsToListPayments(token_)
        if response.status_code == 200:
            print('Запись успешно удалена')
        else:
            print('Ошибка при удалении записи:', response.text)
            return
        print(response.json())


if __name__=="__main__":
    token = TOKEN #str(input("Enter token: "))
    if TOKEN!=token:
        print("Иди на хуй")
        exit()
    while(True):
        id_table = int(input("Enter the table number: "))
        if id_table==1:
            workspaceTable( token)
        elif id_table==2:
            listPaymentsTable( token)
        elif id_table==3:
            listPriceTable( token)
        
        if str(input("Do you want to repeat?(y/n): " ))=="n":
            exit()
