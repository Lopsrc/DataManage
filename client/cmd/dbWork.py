import json
import requests

NAME = "dbWork"
base_url = 'http://localhost:8080'
def create_entry(id_table, token, date=None, price=None, time_work=None, penalty=None, price_day=None, payments=None):
    url = base_url + '/entry'
    headers = {'Content-Type': 'application/json'}
    data = {
        'id_table': id_table,
        'date': date,
        'price': price,
        'time_work': time_work,
        'penalty': penalty,
        'price_day': price_day,
        'payments': payments,
        'token': token
    }
    response = requests.post(url, headers=headers, data=json.dumps(data))
    return response

def update_entry(id_table, token, entry_id, date=None, price=None, time_work=None, penalty=None, price_day=None, payments=None):
    url = base_url + '/entry'
    headers = {'Content-Type': 'application/json'}
    data = {
        'id': entry_id,
        'id_table': id_table,
        'date': date,
        'price': price,
        'time_work': time_work,
        'penalty': penalty,
        'price_day': price_day,
        'payments': payments,
        'token': token
    }
    response = requests.put(url, headers=headers, data=json.dumps(data))
    return response

def delete_entry(id_table, entry_id, token):
    url = base_url + '/entry'
    headers = {'Content-Type': 'application/json'}
    data = {'id': entry_id,'id_table': id_table, 'token': token}
    response = requests.delete(url, headers=headers, data=json.dumps(data))
    return response