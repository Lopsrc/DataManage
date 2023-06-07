import json
import requests

NAME = "dbWork"
base_url = 'http://localhost:8080'
def createRecordToWorkspace(name_workspace, date, price, time_work, penalty, token):
    url = base_url + '/workspace'
    headers = {'Content-Type': 'application/json'}
    data = {
        'name_workspace': name_workspace,
        'date': date,
        'price': price,
        'time_work': time_work,
        'penalty': penalty,
        'token': token
    }
    response = requests.post(url, headers=headers, data=json.dumps(data))
    return response

def updateRecordToWorkspace( record_id, name_workspace, date, price, time_work, penalty, token):
    url = base_url + '/workspace'
    headers = {'Content-Type': 'application/json'}
    data = {
        'id': record_id,
        'name_workspace': name_workspace,
        'date': date,
        'price': price,
        'time_work': time_work,
        'penalty': penalty,
        'token': token
    }
    response = requests.put(url, headers=headers, data=json.dumps(data))
    return response

def deleteRecordToWorkspace(record_id, token):
    url = base_url + '/workspace'
    headers = {'Content-Type': 'application/json'}
    data = {'id': record_id,'token': token}
    response = requests.delete(url, headers=headers, data=json.dumps(data))
    return response


def createRecordToListPayments(name_workspace,date, price, token):
    url = base_url + '/payments'
    headers = {'Content-Type': 'application/json'}
    data = {
        'name_workspace': name_workspace,
        'date': date,
        'price': price,
        'token': token
    }
    response = requests.post(url, headers=headers, data=json.dumps(data))
    return response

def updateRecordToListPayments(entry_id, name_workspace, date, price, token):
    url = base_url + '/payments'
    headers = {'Content-Type': 'application/json'}
    data = {
        'id': entry_id,
        'name_workspace': name_workspace,
        'date': date,
        'price': price,
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


def getAllRecordsToWorkspace(token):
    url = base_url + '/workspace'
    headers = {'Content-Type': 'application/json'}
    data = { 'token': token}
    response = requests.get(url, headers=headers, data=json.dumps(data))
    return response

def getAllRecordsToListPayments( token):
    url = base_url + '/payments'
    headers = {'Content-Type': 'application/json'}
    data = {'token': token}
    response = requests.get(url, headers=headers, data=json.dumps(data))
    return response