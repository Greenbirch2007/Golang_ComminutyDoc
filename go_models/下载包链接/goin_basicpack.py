# -*- coding: utf-8 -*-

# 读取页面文本
# 按照标题，保存整个文本


import csv
import datetime

import os
import re
import time
import sys

type = sys.getfilesystemencoding()
import pymysql
import xlrd
import requests
from requests.exceptions import RequestException
from lxml import etree


def call_page(url):
    try:
        response = requests.get(url)
        response.encoding = 'utf-8'  #
        if response.status_code == 200:
            return response.text
        return None
    except RequestException:
        return None


def removeDot(item):
    f_l = []
    for it in item:
        f_str = "".join(it.split(","))
        f_l.append(f_str)

    return f_l


def remove_block(items):
    new_items = []
    for it in items:
        f = "".join(it.split())
        new_items.append(f)
    return new_items


def writerDt_csv(headers, rowsdata):
    # rowsdata列表中的数据元组,也可以是字典数据
    with open('tokyoTSN.csv', 'w', newline='') as f:
        f_csv = csv.writer(f)
        f_csv.writerow(headers)
        f_csv.writerows(rowsdata)


def read_xlrd(excelFile):
    data = xlrd.open_workbook(excelFile)
    table = data.sheet_by_index(0)
    dataFile = []
    for rowNum in range(table.nrows):
        dataFile.append(table.row_values(rowNum))

    # # if 去掉表头
    # if rowNum > 0:

    return dataFile


# xlsx---list_url----单页url
def get_allURL():
    lpath = os.getcwd()
    excelFile = '{0}\\urls.xlsx'.format(lpath)
    full_items = read_xlrd(excelFile=excelFile)
    for item in full_items:
        for num in range(1, 4):
            f_url = item[0] + 'list_{0}.html'.format(num)
            BIG_URL.append(f_url)


'<h2>日语常用语180句——23</h2>'


def getOneText(url):
    title = []
    html = call_page(url)
    patt = re.compile('<title>(.*?)_.*?</title>', re.S)
    title1 = re.findall(patt, html)

    selector = etree.HTML(html)
    title2 = selector.xpath('/html/body/div[4]/div[3]/div[1]/div[2]/div[1]/h2/text()')
    if len(title1) == 0:
        title = title2
    else:
        title = title1

    content1 = selector.xpath('//*[@id="article"]/div/text()')
    content2 = selector.xpath('//*[@id="article"]/p/text()')
    f_con = []
    if len(content1) != 0 and len(content2) != 0 and len(content1) > len(content2):
        f_con = content1
    elif len(content1) != 0 and len(content2) != 0 and len(content1) < len(content2):
        f_con = content2
    elif len(content1) == 0:
        f_con = content2
    elif len(content2) == 0:
        f_con = content1
    else:
        pass
    for sname in f_con:

        try:
            with open('{0}.txt'.format(title[0]), 'a') as file_handle:
                # .txt可以不自己新建,代码会自动新建

                file_handle.write(sname + ",")  # 写入
                file_handle.write('\n')  # 有时放在循环里面需要自动转行，不然会覆盖上一条数据
                print("{0} 整理完毕".format(title[0]))
        except:
            pass

url ='https://pkg.go.dev/search?page=100&q=Package'
html = call_page(url)
# patt = re.compile('<b class="InfoLabel-title">Version:</b>(.*?)<span class="InfoLabel-divider">|</span>.*?<b class="InfoLabel-title">Published:</b>(.*?)<span class="InfoLabel-divider">|</span>.*?<b class="InfoLabel-title">Imported by:</b>(.*?)<span class="InfoLabel-divider">|</span>.*?<b class="InfoLabel-title">License:</b>(.*?)</div>',re.S)


