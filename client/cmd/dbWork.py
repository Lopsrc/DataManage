import json
import requests

NAME = "dbWork"
base_url = 'http://localhost:8080'
def createRecordToWorkspace(id_table, token, date=None, price=None, time_work=None, penalty=None, payments=None):
    url = base_url + '/workspace'
    headers = {'Content-Type': 'application/json'}
    data = {
        'id_table': id_table,
        'date': date,
        'price': price,
        'time_work': time_work,
        'penalty': penalty,
        'payments': payments,
        'token': token
    }
    response = requests.post(url, headers=headers, data=json.dumps(data))
    return response

def updateRecordToWorkspace(id_table, token, record_id, date=None, price=None, time_work=None, penalty=None, payments=None):
    url = base_url + '/workspace'
    headers = {'Content-Type': 'application/json'}
    data = {
        'id': record_id,
        'id_table': id_table,
        'date': date,
        'price': price,
        'time_work': time_work,
        'penalty': penalty,
        'payments': payments,
        'token': token
    }
    response = requests.put(url, headers=headers, data=json.dumps(data))
    return response

def deleteRecordToWorkspace(id_table, record_id, token):
    url = base_url + '/workspace'
    headers = {'Content-Type': 'application/json'}
    data = {'id': record_id,'id_table': id_table, 'token': token}
    response = requests.delete(url, headers=headers, data=json.dumps(data))
    return response


def createRecordToListPrices(id_table, token, date=None, price=None, time_work=None, penalty=None, payments=None):
    url = base_url + '/price'
    headers = {'Content-Type': 'application/json'}
    data = {
        'id_table': id_table,
        'date': date,
        'price': price,
        'time_work': time_work,
        'penalty': penalty,
        'payments': payments,
        'token': token
    }
    response = requests.post(url, headers=headers, data=json.dumps(data))
    return response

def updateRecordToListPrices(id_table, token, entry_id, date=None, price=None, time_work=None, penalty=None, payments=None):
    url = base_url + '/price'
    headers = {'Content-Type': 'application/json'}
    data = {
        'id': entry_id,
        'id_table': id_table,
        'date': date,
        'price': price,
        'time_work': time_work,
        'penalty': penalty,
        'payments': payments,
        'token': token
    }
    response = requests.put(url, headers=headers, data=json.dumps(data))
    return response

def deleteRecordToListPrices(id_table, entry_id, token):
    url = base_url + '/price'
    headers = {'Content-Type': 'application/json'}
    data = {'id': entry_id,'id_table': id_table, 'token': token}
    response = requests.delete(url, headers=headers, data=json.dumps(data))
    return response

def createRecordToListPayments(id_table, token, date=None, price=None, time_work=None, penalty=None, payments=None):
    url = base_url + '/payments'
    headers = {'Content-Type': 'application/json'}
    data = {
        'id_table': id_table,
        'date': date,
        'price': price,
        'time_work': time_work,
        'penalty': penalty,
        'payments': payments,
        'token': token
    }
    response = requests.post(url, headers=headers, data=json.dumps(data))
    return response

def updateRecordToListPayments(id_table, token, entry_id, date=None, price=None, time_work=None, penalty=None, payments=None):
    url = base_url + '/payments'
    headers = {'Content-Type': 'application/json'}
    data = {
        'id': entry_id,
        'id_table': id_table,
        'date': date,
        'price': price,
        'time_work': time_work,
        'penalty': penalty,
        'payments': payments,
        'token': token
    }
    response = requests.put(url, headers=headers, data=json.dumps(data))
    return response

def deleteRecordToListPayments(entry_id, token):
    url = base_url + '/payments'
    headers = {'Content-Type': 'application/json'}
    data = {'id': entry_id, 'token': token}
    response = requests.delete(url, headers=headers, data=json.dumps(data))
    return response


def getAllRecordsToWorkspace(entry_id, token):
    url = base_url + '/workspace'
    headers = {'Content-Type': 'application/json'}
    data = {'id': entry_id, 'token': token}
    response = requests.get(url, headers=headers, data=json.dumps(data))
    return response
def getAllRecordsToListPrices(entry_id, token):
    url = base_url + '/price'
    headers = {'Content-Type': 'application/json'}
    data = {'id': entry_id, 'token': token}
    response = requests.get(url, headers=headers, data=json.dumps(data))
    return response
def getAllRecordsToListPayments(entry_id, token):
    url = base_url + '/payments'
    headers = {'Content-Type': 'application/json'}
    data = {'id': entry_id, 'token': token}
    response = requests.get(url, headers=headers, data=json.dumps(data))
    return response