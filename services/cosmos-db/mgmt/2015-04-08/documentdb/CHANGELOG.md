Generated from https://github.com/Azure/azure-rest-api-specs/tree/3c764635e7d442b3e74caf593029fcd440b3ef82/specification/cosmos-db/resource-manager/readme.md tag: `package-2015-04`

Code generator @microsoft.azure/autorest.go@2.1.168

## Breaking Changes

### Removed Funcs

1. *DatabaseAccountsCreateOrUpdateFuture.Result(DatabaseAccountsClient) (DatabaseAccount, error)
1. *DatabaseAccountsCreateUpdateCassandraKeyspaceFuture.Result(DatabaseAccountsClient) (CassandraKeyspace, error)
1. *DatabaseAccountsCreateUpdateCassandraTableFuture.Result(DatabaseAccountsClient) (CassandraTable, error)
1. *DatabaseAccountsCreateUpdateGremlinDatabaseFuture.Result(DatabaseAccountsClient) (GremlinDatabase, error)
1. *DatabaseAccountsCreateUpdateGremlinGraphFuture.Result(DatabaseAccountsClient) (GremlinGraph, error)
1. *DatabaseAccountsCreateUpdateMongoDBCollectionFuture.Result(DatabaseAccountsClient) (MongoDBCollection, error)
1. *DatabaseAccountsCreateUpdateMongoDBDatabaseFuture.Result(DatabaseAccountsClient) (MongoDBDatabase, error)
1. *DatabaseAccountsCreateUpdateSQLContainerFuture.Result(DatabaseAccountsClient) (SQLContainer, error)
1. *DatabaseAccountsCreateUpdateSQLDatabaseFuture.Result(DatabaseAccountsClient) (SQLDatabase, error)
1. *DatabaseAccountsCreateUpdateTableFuture.Result(DatabaseAccountsClient) (Table, error)
1. *DatabaseAccountsDeleteCassandraKeyspaceFuture.Result(DatabaseAccountsClient) (autorest.Response, error)
1. *DatabaseAccountsDeleteCassandraTableFuture.Result(DatabaseAccountsClient) (autorest.Response, error)
1. *DatabaseAccountsDeleteFuture.Result(DatabaseAccountsClient) (autorest.Response, error)
1. *DatabaseAccountsDeleteGremlinDatabaseFuture.Result(DatabaseAccountsClient) (autorest.Response, error)
1. *DatabaseAccountsDeleteGremlinGraphFuture.Result(DatabaseAccountsClient) (autorest.Response, error)
1. *DatabaseAccountsDeleteMongoDBCollectionFuture.Result(DatabaseAccountsClient) (autorest.Response, error)
1. *DatabaseAccountsDeleteMongoDBDatabaseFuture.Result(DatabaseAccountsClient) (autorest.Response, error)
1. *DatabaseAccountsDeleteSQLContainerFuture.Result(DatabaseAccountsClient) (autorest.Response, error)
1. *DatabaseAccountsDeleteSQLDatabaseFuture.Result(DatabaseAccountsClient) (autorest.Response, error)
1. *DatabaseAccountsDeleteTableFuture.Result(DatabaseAccountsClient) (autorest.Response, error)
1. *DatabaseAccountsFailoverPriorityChangeFuture.Result(DatabaseAccountsClient) (autorest.Response, error)
1. *DatabaseAccountsOfflineRegionFuture.Result(DatabaseAccountsClient) (autorest.Response, error)
1. *DatabaseAccountsOnlineRegionFuture.Result(DatabaseAccountsClient) (autorest.Response, error)
1. *DatabaseAccountsPatchFuture.Result(DatabaseAccountsClient) (DatabaseAccount, error)
1. *DatabaseAccountsRegenerateKeyFuture.Result(DatabaseAccountsClient) (autorest.Response, error)
1. *DatabaseAccountsUpdateCassandraKeyspaceThroughputFuture.Result(DatabaseAccountsClient) (Throughput, error)
1. *DatabaseAccountsUpdateCassandraTableThroughputFuture.Result(DatabaseAccountsClient) (Throughput, error)
1. *DatabaseAccountsUpdateGremlinDatabaseThroughputFuture.Result(DatabaseAccountsClient) (Throughput, error)
1. *DatabaseAccountsUpdateGremlinGraphThroughputFuture.Result(DatabaseAccountsClient) (Throughput, error)
1. *DatabaseAccountsUpdateMongoDBCollectionThroughputFuture.Result(DatabaseAccountsClient) (Throughput, error)
1. *DatabaseAccountsUpdateMongoDBDatabaseThroughputFuture.Result(DatabaseAccountsClient) (Throughput, error)
1. *DatabaseAccountsUpdateSQLContainerThroughputFuture.Result(DatabaseAccountsClient) (Throughput, error)
1. *DatabaseAccountsUpdateSQLDatabaseThroughputFuture.Result(DatabaseAccountsClient) (Throughput, error)
1. *DatabaseAccountsUpdateTableThroughputFuture.Result(DatabaseAccountsClient) (Throughput, error)

## Struct Changes

### Removed Struct Fields

1. DatabaseAccountsCreateOrUpdateFuture.azure.Future
1. DatabaseAccountsCreateUpdateCassandraKeyspaceFuture.azure.Future
1. DatabaseAccountsCreateUpdateCassandraTableFuture.azure.Future
1. DatabaseAccountsCreateUpdateGremlinDatabaseFuture.azure.Future
1. DatabaseAccountsCreateUpdateGremlinGraphFuture.azure.Future
1. DatabaseAccountsCreateUpdateMongoDBCollectionFuture.azure.Future
1. DatabaseAccountsCreateUpdateMongoDBDatabaseFuture.azure.Future
1. DatabaseAccountsCreateUpdateSQLContainerFuture.azure.Future
1. DatabaseAccountsCreateUpdateSQLDatabaseFuture.azure.Future
1. DatabaseAccountsCreateUpdateTableFuture.azure.Future
1. DatabaseAccountsDeleteCassandraKeyspaceFuture.azure.Future
1. DatabaseAccountsDeleteCassandraTableFuture.azure.Future
1. DatabaseAccountsDeleteFuture.azure.Future
1. DatabaseAccountsDeleteGremlinDatabaseFuture.azure.Future
1. DatabaseAccountsDeleteGremlinGraphFuture.azure.Future
1. DatabaseAccountsDeleteMongoDBCollectionFuture.azure.Future
1. DatabaseAccountsDeleteMongoDBDatabaseFuture.azure.Future
1. DatabaseAccountsDeleteSQLContainerFuture.azure.Future
1. DatabaseAccountsDeleteSQLDatabaseFuture.azure.Future
1. DatabaseAccountsDeleteTableFuture.azure.Future
1. DatabaseAccountsFailoverPriorityChangeFuture.azure.Future
1. DatabaseAccountsOfflineRegionFuture.azure.Future
1. DatabaseAccountsOnlineRegionFuture.azure.Future
1. DatabaseAccountsPatchFuture.azure.Future
1. DatabaseAccountsRegenerateKeyFuture.azure.Future
1. DatabaseAccountsUpdateCassandraKeyspaceThroughputFuture.azure.Future
1. DatabaseAccountsUpdateCassandraTableThroughputFuture.azure.Future
1. DatabaseAccountsUpdateGremlinDatabaseThroughputFuture.azure.Future
1. DatabaseAccountsUpdateGremlinGraphThroughputFuture.azure.Future
1. DatabaseAccountsUpdateMongoDBCollectionThroughputFuture.azure.Future
1. DatabaseAccountsUpdateMongoDBDatabaseThroughputFuture.azure.Future
1. DatabaseAccountsUpdateSQLContainerThroughputFuture.azure.Future
1. DatabaseAccountsUpdateSQLDatabaseThroughputFuture.azure.Future
1. DatabaseAccountsUpdateTableThroughputFuture.azure.Future

## Struct Changes

### New Struct Fields

1. DatabaseAccountsCreateOrUpdateFuture.Result
1. DatabaseAccountsCreateOrUpdateFuture.azure.FutureAPI
1. DatabaseAccountsCreateUpdateCassandraKeyspaceFuture.Result
1. DatabaseAccountsCreateUpdateCassandraKeyspaceFuture.azure.FutureAPI
1. DatabaseAccountsCreateUpdateCassandraTableFuture.Result
1. DatabaseAccountsCreateUpdateCassandraTableFuture.azure.FutureAPI
1. DatabaseAccountsCreateUpdateGremlinDatabaseFuture.Result
1. DatabaseAccountsCreateUpdateGremlinDatabaseFuture.azure.FutureAPI
1. DatabaseAccountsCreateUpdateGremlinGraphFuture.Result
1. DatabaseAccountsCreateUpdateGremlinGraphFuture.azure.FutureAPI
1. DatabaseAccountsCreateUpdateMongoDBCollectionFuture.Result
1. DatabaseAccountsCreateUpdateMongoDBCollectionFuture.azure.FutureAPI
1. DatabaseAccountsCreateUpdateMongoDBDatabaseFuture.Result
1. DatabaseAccountsCreateUpdateMongoDBDatabaseFuture.azure.FutureAPI
1. DatabaseAccountsCreateUpdateSQLContainerFuture.Result
1. DatabaseAccountsCreateUpdateSQLContainerFuture.azure.FutureAPI
1. DatabaseAccountsCreateUpdateSQLDatabaseFuture.Result
1. DatabaseAccountsCreateUpdateSQLDatabaseFuture.azure.FutureAPI
1. DatabaseAccountsCreateUpdateTableFuture.Result
1. DatabaseAccountsCreateUpdateTableFuture.azure.FutureAPI
1. DatabaseAccountsDeleteCassandraKeyspaceFuture.Result
1. DatabaseAccountsDeleteCassandraKeyspaceFuture.azure.FutureAPI
1. DatabaseAccountsDeleteCassandraTableFuture.Result
1. DatabaseAccountsDeleteCassandraTableFuture.azure.FutureAPI
1. DatabaseAccountsDeleteFuture.Result
1. DatabaseAccountsDeleteFuture.azure.FutureAPI
1. DatabaseAccountsDeleteGremlinDatabaseFuture.Result
1. DatabaseAccountsDeleteGremlinDatabaseFuture.azure.FutureAPI
1. DatabaseAccountsDeleteGremlinGraphFuture.Result
1. DatabaseAccountsDeleteGremlinGraphFuture.azure.FutureAPI
1. DatabaseAccountsDeleteMongoDBCollectionFuture.Result
1. DatabaseAccountsDeleteMongoDBCollectionFuture.azure.FutureAPI
1. DatabaseAccountsDeleteMongoDBDatabaseFuture.Result
1. DatabaseAccountsDeleteMongoDBDatabaseFuture.azure.FutureAPI
1. DatabaseAccountsDeleteSQLContainerFuture.Result
1. DatabaseAccountsDeleteSQLContainerFuture.azure.FutureAPI
1. DatabaseAccountsDeleteSQLDatabaseFuture.Result
1. DatabaseAccountsDeleteSQLDatabaseFuture.azure.FutureAPI
1. DatabaseAccountsDeleteTableFuture.Result
1. DatabaseAccountsDeleteTableFuture.azure.FutureAPI
1. DatabaseAccountsFailoverPriorityChangeFuture.Result
1. DatabaseAccountsFailoverPriorityChangeFuture.azure.FutureAPI
1. DatabaseAccountsOfflineRegionFuture.Result
1. DatabaseAccountsOfflineRegionFuture.azure.FutureAPI
1. DatabaseAccountsOnlineRegionFuture.Result
1. DatabaseAccountsOnlineRegionFuture.azure.FutureAPI
1. DatabaseAccountsPatchFuture.Result
1. DatabaseAccountsPatchFuture.azure.FutureAPI
1. DatabaseAccountsRegenerateKeyFuture.Result
1. DatabaseAccountsRegenerateKeyFuture.azure.FutureAPI
1. DatabaseAccountsUpdateCassandraKeyspaceThroughputFuture.Result
1. DatabaseAccountsUpdateCassandraKeyspaceThroughputFuture.azure.FutureAPI
1. DatabaseAccountsUpdateCassandraTableThroughputFuture.Result
1. DatabaseAccountsUpdateCassandraTableThroughputFuture.azure.FutureAPI
1. DatabaseAccountsUpdateGremlinDatabaseThroughputFuture.Result
1. DatabaseAccountsUpdateGremlinDatabaseThroughputFuture.azure.FutureAPI
1. DatabaseAccountsUpdateGremlinGraphThroughputFuture.Result
1. DatabaseAccountsUpdateGremlinGraphThroughputFuture.azure.FutureAPI
1. DatabaseAccountsUpdateMongoDBCollectionThroughputFuture.Result
1. DatabaseAccountsUpdateMongoDBCollectionThroughputFuture.azure.FutureAPI
1. DatabaseAccountsUpdateMongoDBDatabaseThroughputFuture.Result
1. DatabaseAccountsUpdateMongoDBDatabaseThroughputFuture.azure.FutureAPI
1. DatabaseAccountsUpdateSQLContainerThroughputFuture.Result
1. DatabaseAccountsUpdateSQLContainerThroughputFuture.azure.FutureAPI
1. DatabaseAccountsUpdateSQLDatabaseThroughputFuture.Result
1. DatabaseAccountsUpdateSQLDatabaseThroughputFuture.azure.FutureAPI
1. DatabaseAccountsUpdateTableThroughputFuture.Result
1. DatabaseAccountsUpdateTableThroughputFuture.azure.FutureAPI
