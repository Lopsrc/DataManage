

import dbWork

fix_price_of_the_day = 3200
TOKEN = 'valid_token'

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
        return fix_price_of_the_day/2-penalty
    return fix_price_of_the_day-penalty

def getRecordId():
    pass


def table_01(id_table, token_):
    action = int(input("Select action (1. Create new record, 2. Update record, 3. Delete record): "))
    if action==1:
        date = str(input("Enter date(yyyy-mm-dd): "))
        time_work = int(input("Enter time of work(8 or 4): "))
        penalty = int(input("Enter penalty(0 or other num): "))
        price = count_price_of_the_day(time_work, penalty)

        response = dbWork.createRecordToWorkspace(id_table,token_,date, price, time_work, penalty)
        if response.status_code == 200:
            print('Запись успешно создана')
        else:
            print('Ошибка при создании записи:', response.text, response.status_code)
    elif action==2:
        date = str(input("Enter date(yyyy-mm-dd): "))
        time_work = int(input("Enter time of work(8 or 4): "))
        penalty = int(input("Enter penalty(0 or other num): "))
        price = count_price_of_the_day(time_work, penalty)
        record_id = 1#getRecordId()

        response = dbWork.updateRecordToWorkspace(id_table, token_,record_id, date, price, time_work, penalty)
        if response.status_code == 200:
            print('Запись успешно создана')
        else:
            print('Ошибка при создании записи:', response.text, response.status_code)
    elif action==3:
        record_id = 1  # Идентификатор записи, которую нужно удалить

        response = dbWork.deleteRecordToWorkspace(id_table, record_id, token_)
        if response.status_code == 200:
            print('Запись успешно удалена')
        else:
            print('Ошибка при удалении записи:', response.text)



def table_02(id_table, token_):
    action = int(input("Select action (1. Create new record, 2. Update record, 3. Delete record): "))
    if action==1:
        date = str(input("Enter date(yyyy-mm-dd): "))
        price = int(input("Enter price of the day(3200): "))

        response = dbWork.createRecordToListPrices(id_table,token_,date, price)
        if response.status_code == 200:
            print('Запись успешно создана')
        else:
            print('Ошибка при создании записи:', response.text, response.status_code)
    elif action==2:
        date = str(input("Enter date(yyyy-mm-dd): "))
        
        price = int(input("Enter price of the day(3200): "))
        record_id = 1#getRecordId()

        response = dbWork.updateRecordToListPrices(id_table, token_,record_id, date, price)
        if response.status_code == 200:
            print('Запись успешно создана')
        else:
            print('Ошибка при создании записи:', response.text, response.status_code)
    elif action==3:
        record_id = 1  # Идентификатор записи, которую нужно удалить

        response = dbWork.deleteRecordToListPrices(id_table, record_id, token_)
        if response.status_code == 200:
            print('Запись успешно удалена')
        else:
            print('Ошибка при удалении записи:', response.text)



def table_03(id_table, token_):
    action = int(input("Select action (1. Create new record, 2. Update record, 3. Delete record): "))
    if action==1:
        date = str(input("Enter date(yyyy-mm-dd): "))
        price = int(input("Enter price of the day(3200): "))

        response = dbWork.createRecordToListPayments(id_table,token_,date, price)
        if response.status_code == 200:
            print('Запись успешно создана')
        else:
            print('Ошибка при создании записи:', response.text, response.status_code)
    elif action==2:
        date = str(input("Enter date(yyyy-mm-dd): "))
        
        price = int(input("Enter price of the day(3200): "))
        record_id = 1#getRecordId()

        response = dbWork.updateRecordToListPrices(id_table, token_,record_id, date, price)
        if response.status_code == 200:
            print('Запись успешно создана')
        else:
            print('Ошибка при создании записи:', response.text, response.status_code)
    elif action==3:
        record_id = 1  # Идентификатор записи, которую нужно удалить

        response = dbWork.deleteRecordToListPrices(id_table, record_id, token_)
        if response.status_code == 200:
            print('Запись успешно удалена')
        else:
            print('Ошибка при удалении записи:', response.text)



if __name__=="__main__":
    token_ = TOKEN #str(input("Enter an integer as input: "))
    if TOKEN!=token_:
        print("Иди на хуй")
        exit()
    while(True):
        
        id_table = int(input("Enter the table number: "))
        if id_table==1:
            table_01(id_table, token_)
        elif id_table==2:
            table_02(id_table, token_)
        elif id_table==3:
            table_03(id_table, token_)

        if str(input("Do you want to repeat?(y/n): " ))=="n":
            exit()
