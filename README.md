# golang_line

golang test step:https://hackmd.io/AqwKuHViRASxKGRYYs571w

這是一個golang的Line bot專案\
Line bot 將會記錄使用者傳入的訊息 也可以詢問Line Bot使用者ID

另外可以藉由API的方式 讓Line Bot傳送訊息給使用者\
或是查詢使用者曾經輸入的訊息\
API參考網址為：https://golangline.docs.apiary.io/# 

## 部屬前置作業
- 需先申請Line Message API 取得Channel secret及Channel access token
- 需安裝docker及docker compose
- 需安裝ngork

## 部屬程序

1. 下載專案至本機
```
git clone https://github.com/monkeyteacher/golang_line.git
```
  
2. 創建程式設定檔```vim app.env```
```
SERVER_PROT=":8088"
CHANNEL_SECRET="" ##line Channel secret
CHANNEL_TOKEN="" ##Channel access token
DB_IP="mongodb://{資料庫帳號}:{資料庫密碼}@{DBIP}:27017"
DB_NAME="line"
```
3. 創建docker compose 設定檔```vim .env```
```
MONGO_INITDB_ROOT_USERNAME={資料庫帳號}
MONGO_INITDB_ROOT_PASSWORD={資料庫密碼}
```
4. 執行docker compose 
```
docker compose up -d
```
5. 執行ngrok 
```
ngrok http 8088
```
6. 在Line Developers的控制台中 設定Webhook URL(ngrok輸出的https)
