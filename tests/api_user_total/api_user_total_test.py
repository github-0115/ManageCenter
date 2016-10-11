# coding: utf8

import requests
import json
import hashlib
from datetime import date
import subprocess


class ManageCenterTest():
    managername = 'deepir'
    password = 'deepir1234'
    login_token = ''
    header = {}

    username = 'wei111'
    month_total = 300
    concurrency = 10

    base_url = 'http://120.27.157.146:8020'

    def __init__(self, month_total):
        self.month_total = month_total

    def login(self):
        login_url = self.base_url + '/login_token'
        params = {'managername': self.managername, 'password': self.password}
        resp = requests.post(login_url, json=params)
        print resp.status_code
        print resp.content
        content = json.loads(resp.content)
        if content.get('code') == 0:
            print 'login OK'
        else:
            print 'login ERROR'
        token = content.get('token')
        print token
        self.login_token = token
        self.header['LoginToken'] = token
        return token

    def set_user_service(self):

        params = {'username': self.username, 'month_total': self.month_total, 'concurrency': self.concurrency}
        url = self.base_url + '/set_user_service'
        resp = requests.post(url, json=params, headers=self.header)

        print resp.status_code
        print resp.content
        content = json.loads(resp.content)
        if content.get('code') == 0:
            print 'set user service OK'
        else:
            print 'set user service ERROR'


class UserCenter():
    base_url = 'http://120.27.157.146:8015'
    ak = 'e84bc127-2756-40cf-9c32-da958728057f'
    sk = '99fb8d12-8581-4ff5-ae44-a28977007ce6'

    def api_token(self):
        url = self.base_url + '/api_token'
        m = hashlib.md5()
        today = date.today().strftime('%Y-%m-%d')
        data_keys = ['accessKey', 'secretKey', 'date']
        data_values = [self.ak, self.sk, today]
        data = [data_keys[i] +'='+ data_values[i] for i in xrange(len(data_keys))]
        plain = '&'.join(data)
        print 'plain=', plain
        m.update(plain)

        sign = m.hexdigest()
        print 'sign=', sign
        params = {'access_key': self.ak, 'sign': sign}
        resp = requests.post(url, json=params)
        print resp.status_code
        print resp.content
        content = json.loads(resp.content)
        if content.get('code') == 0:
            print 'api token OK'
        else:
            print 'api token ERROR'
        api_token = content.get('api_token')
        return api_token


class APITest():
    # base_url = 'http://classify.deepir.com'
    base_url = 'http://120.55.114.15'
    api_token = ''
    header = {}
    image = None
    image_filename = 'bf0024.jpg'
    fp_image = None

    def __init__(self, api_token):
        self.api_token = api_token
        self.header = {'token': api_token}
        self.fp_image = open(self.image_filename)

    def __del__(self):
        self.fp_image.close()

    def image_classify(self):
        url = self.base_url + '/image_classify'
        # params = {'image': base64.b64encode(self.image)}
        files = {'image': self.fp_image}
        # resp = requests.post(url, files=files, headers=self.header)
        with open(self.image_filename) as fp:
            files = {'image': fp}
            resp = requests.post(url, files=files, headers=self.header)
        # resp = requests.post(url, json=params, headers=self.header)
        print resp.status_code
        print resp.content
        content = json.loads(resp.content)
        if content.get('code') == 0:
            print 'classify OK'
        else:
            print 'classify ERROR'
            print resp.content


def main():
    month_total = 300
    day_total = month_total / 30

    mc = ManageCenterTest(month_total)
    mc.login()
    mc.set_user_service()

    uc = UserCenter()
    api_token = uc.api_token()

    # run cronjob set daily api redis
    ret = subprocess.call(['go', 'run', '../../cronjob/daily_api_user_total/daily_api_user_total.go'])
    if ret != 0:
        print 'cronjob ERROR, now exit'
        return

    api = APITest(api_token)
    for i in xrange(day_total + 2):
        print 'No.{0} classify'.format(i+1)
        api.image_classify()


if __name__ == '__main__':
    main()
