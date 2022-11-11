import json
import os.path
from pprint import pprint
import xlsxwriter

import MCC
import requests

headers = {
    'Authorization': 'Bearer N1DyzjVoxdjmZTXaO0L1TeX4RxFTVJZ3vzJFhdU0vGWCY6PO1WRScVXrsn3sRFoaJIbyXmclztTWyPXvjoO6l8wJEY0EifAlRkjx4HHdoul4QqddI8KSpBvkSsz4hNA2FQrwQ4eKPaQ8ydSOmCUpWsjSfkQFSbJy17pfapkOLfKwl58piO6tFUiWEiirUMTk2IkxwhjcNC6NWrioKXb3O6KBVwCpKEIXGEhdHyNtMT6xfF0E50BeXP9H0DaMl783Iu3ubUi7AoboCWvEAdtKhLgiabhSr4n9yQBFUAPyPfwSsocLoV5h202609e1KwrLwSjmfYG48Y08C3OYaQh8Url71MP0Fy9uA5Qz0lHzgT1E4mzjVDrnsQJc2tgbkbqh9bO4fi7Dpah7TbShhfHEqIEoUZcyI5aKlZHUAfTKHnhVZZu1diQDFAKTnnU0ySGqbgsyNqt4FqzwjR26GzensTmbyFS6aihyZ4tJbYiJWl3NxdaCej5Y2kdyOItj82a1'
}


gs_id = "1xLlkUETuM5_PTdzlXDxg1GN0Hbo9JqCfguSR4Jwqe08"
gs_sheet_name = "Sheet1"
split_data = "2022-10-31"

# url = "http://127.0.0.1:50001/report/a543_a?split_by=2021-10-01"
# url = "http://159.203.47.150:50001/report/a543_a?split_by=2021-11-30"
# url = "http://159.203.47.150:50001/report/a543_a?split_by=2021-12-31"
# url = "http://159.203.47.150:50001/report/a543_a?split_by="+"2022-01-31"
# url = "http://159.203.47.150:50001/report/a543_a?split_by="+"2022-02-28"
# url = "http://159.203.47.150:50001/report/a543_a?split_by="+"2022-03-31"
# url = "http://159.203.47.150:50001/report/a543_a?split_by="+"2022-04-30"
# url = "http://159.203.47.150:50001/report/a543_a?split_by="+"2022-05-31"
# url = "http://159.203.47.150:50001/report/a543_a?split_by="+"2022-06-30"
# url = "http://159.203.47.150:50001/report/a543_a?split_by="+"2022-07-31"
url = "http://159.203.47.150:50001/report/a543_a?split_by="+split_data

obj = {}
if os.path.isfile(split_data+".json"):
    with open(split_data+".json", "r") as f:
        obj = json.load(f)
else:
    response = requests.request("GET", url, headers=headers)
    obj = json.loads(response.text)
    with open(split_data+".json", "w") as f:
        json.dump(obj, f)

pprint(obj)

# gs = MCC.GoogleSheet(gs_id)
# gs.jsonToSheet(obj, gs_sheet_name+"!A2")

workbook = xlsxwriter.Workbook(split_data + '.xlsx')
worksheet = workbook.add_worksheet()

col = 0
for key in obj["header"]:
    worksheet.write(0, col, obj["header"][key])
    col +=1

r = 1
for row in obj["data"]:
    c = 0
    for key in obj["header"]:
        print(row[key])
        worksheet.write(r, c, row[key])
        c +=1
    r +=1

workbook.close()

