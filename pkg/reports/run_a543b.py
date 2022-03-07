import json

import MCC
import requests

headers = {
    # 'Authorization': 'Bearer N1DyzjVoxdjmZTXaO0L1TeX4RxFTVJZ3vzJFhdU0vGWCY6PO1WRScVXrsn3sRFoaJIbyXmclztTWyPXvjoO6l8wJEY0EifAlRkjx4HHdoul4QqddI8KSpBvkSsz4hNA2FQrwQ4eKPaQ8ydSOmCUpWsjSfkQFSbJy17pfapkOLfKwl58piO6tFUiWEiirUMTk2IkxwhjcNC6NWrioKXb3O6KBVwCpKEIXGEhdHyNtMT6xfF0E50BeXP9H0DaMl783Iu3ubUi7AoboCWvEAdtKhLgiabhSr4n9yQBFUAPyPfwSsocLoV5h202609e1KwrLwSjmfYG48Y08C3OYaQh8Url71MP0Fy9uA5Qz0lHzgT1E4mzjVDrnsQJc2tgbkbqh9bO4fi7Dpah7TbShhfHEqIEoUZcyI5aKlZHUAfTKHnhVZZu1diQDFAKTnnU0ySGqbgsyNqt4FqzwjR26GzensTmbyFS6aihyZ4tJbYiJWl3NxdaCej5Y2kdyOItj82a1'
    'Authorization': 'Bearer 1'
}


gs_id = "1e2Pmi5FDwoSmQzBXGCiMrt-CxsBsXrD1zWiHHTWrlfw"
gs_sheet_name = "Sheet1"


# url = "http://159.203.47.150:50001/report/a543_b?split_by="+"2021-10-01"
url = "http://127.0.0.1:50001/report/a543_b?split_by="+"2021-10-01"

response = requests.request("GET", url, headers=headers)

obj = json.loads(response.text)

# print(obj)

gs = MCC.GoogleSheet(gs_id)
gs.jsonToSheet(obj, gs_sheet_name+"!A2")


