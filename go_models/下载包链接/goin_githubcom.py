

import datetime

import pymysql
import requests
from lxml import etree
import json
from queue import Queue
import threading
from requests.exceptions import RequestException
import re





import asyncio
import aiohttp


def run_forever(func):
    def wrapper(obj):
        while True:
            func(obj)
    return wrapper

def find_longest_str(str_list):
    '''
    找到列表中字符串最长的位置索引
    先获取列表中每个字符串的长度，查找长度最大位置的索引值即可
    '''
    num_list = [len(one) for one in str_list]
    index_num = num_list.index(max(num_list))
    return str_list[int(index_num)]


def remove_block(items):
    new_items = []
    for it in items:
        f = "".join(it.split())
        new_items.append(f)
    return new_items



async def get_title(i):


    url = 'https://pkg.go.dev/search?page={0}&q=github.com'.format(i)
    # conn = aiohttp.TCPConnector(verify_ssl=False)  # 防止ssl报错
    async with aiohttp.ClientSession(connector=aiohttp.TCPConnector(limit=10, verify_ssl=False)) as session:
        async with session.get(url) as resp:
            print(resp.status)
            text = await resp.text()
            print('start', i)

    big_list = []
    selector = etree.HTML(text)
    package_name = selector.xpath('/html/body/main/div/div/div[3]/div/h2/a/text()')
    package_url = selector.xpath('/html/body/main/div/div/div[3]/div/h2/a/@href')

    import_num =[]

    patt = re.compile('<b class="InfoLabel-title">Imported by:</b>(.*?)<span class="InfoLabel-divider">|</span>', re.S)
    item = re.findall(patt, text)
    for it in range(2, len(item) + 1, 4):
        try:
            import_num.append(item[it].split()[0])
        except IndexError:
            import_num.append("可能很多")

    for i1,i2,i3 in zip(package_name,package_url,import_num):
        big_list.append((i1,'https://pkg.go.dev/'+i2,i3))


    connection = pymysql.connect(host='127.0.0.1', port=3306, user='root', password='123456', db='goPkg',charset='utf8mb4', cursorclass=pymysql.cursors.DictCursor)
    cursor = connection.cursor()
    try:
        cursor.executemany('insert into goPkg_githubcom (pname,purl,import_num) values (%s,%s,%s)', big_list)
        connection.commit()
        connection.close()
        print('向MySQL中添加数据成功！')
    except TypeError :
        pass






loop = asyncio.get_event_loop()
fun_list = (get_title(i) for i in range(1,10001))
try:
    loop.run_until_complete(asyncio.gather(*fun_list))
except:
    pass


# create table goPkg_githubcom(
# id int not null primary key auto_increment,
# pname text,
# purl text,
# import_num text
# ) engine=InnoDB  charset=utf8;

# drop table goPkg_githubcom;