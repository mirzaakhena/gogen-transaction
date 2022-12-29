# Gogen-transaction

1. First make sure which database you want to run, `gorm sqlite` or `mongodb` by swithing it in `application/app_apptrx.go` line 32
   
   ```
   // use this for gorm sqlite
   datasource := withgorm.NewGateway(log, appData, cfg)
   
   // or use this for mongodb
   datasource := withmongodb.NewGateway(log, appData, cfg)
   ```
   
2. Run the application with command 

   ```
   $ go run main.go apptrx
   ```
   
3. Use this curl command in console to simulate the success insert dan trigger the transaction commit
   ```
   $ curl -X POST 'http://localhost:8000/api/v1/trx?user=1'
   ```
   
   or use this curl command to simulate the fail insert and trigger the transaction rollback
   ```
   $ curl -X POST 'http://localhost:8000/api/v1/trx?user='
   ```

4. Commit transaction will have output
   ```
   INFO  582AH892MS8NVGBS {"user":"1"}        restapi.(*ginController).runTransactionHandler.func1:53
   INFO  582AH892MS8NVGBS Begin trx           database.(*GormWithTransaction).BeginTransaction:41
   INFO  582AH892MS8NVGBS called SaveProduct  withgorm.(*gateway).SaveProduct:45
   INFO  582AH892MS8NVGBS called SaveOrder    withgorm.(*gateway).SaveOrder:56
   INFO  582AH892MS8NVGBS Commit trx          database.(*GormWithTransaction).CommitTransaction:48
   INFO  582AH892MS8NVGBS {}                  restapi.(*ginController).runTransactionHandler.func1:65
   [GIN] 2022/12/29 - 15:53:28 | 200 |    1.518875ms |       127.0.0.1 | POST     "/api/v1/trx?user=1"
   ```
   
   Rollback transction will have output
   ```
   INFO  MH0INZ1CDDVTZZC5 {"user":""}         restapi.(*ginController).runTransactionHandler.func1:53
   INFO  MH0INZ1CDDVTZZC5 Begin trx           database.(*GormWithTransaction).BeginTransaction:41
   INFO  MH0INZ1CDDVTZZC5 called SaveProduct  withgorm.(*gateway).SaveProduct:45
   INFO  MH0INZ1CDDVTZZC5 Rollback trx        database.(*GormWithTransaction).RollbackTransaction:53
   ERROR MH0INZ1CDDVTZZC5 user must not empty restapi.(*ginController).runTransactionHandler.func1:57
   [GIN] 2022/12/29 - 16:09:55 | 400 |    1.430917ms |       127.0.0.1 | POST     "/api/v1/trx?user="   
   ```