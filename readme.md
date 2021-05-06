package utils // import "github.com/innv8/utils"


FUNCTIONS

func ChannelConnect(conn *amqp.Connection) (channel *amqp.Channel, err error)
    ChannelConnect Connects to a channel Takes conn Returns channel, error

func ConnectToCache(ctx context.Context, host, password string, db int) (client *redis.Client, err error)
    ConnectToCache connects to single redis instance Takes context, host,
    password and db Returns client, error

func DBConnect(driver, uri string) (client *sql.DB, err error)
    DBConnect Connects to DB Takes driver, uri Returns client, error

func DeleteFromCache(ctx context.Context, key string, client *redis.Client) (err error)
    DeleteFromCache Deletes a key from redis Takes context, key, client Returns
    error

func InitLogger(logFolder, env string)
    InitLogger initializes the logger. It requires logFolder and env. logFolder
    is the path of the log directory ending with a slash If env is set to debug,
    the log output will be the stderr. If env is set to anything else (e.g
    staging or prod or nothing), the output will be a files, rotated every 24
    hours

func LogDBStats(client *sql.DB)
    LogDBStats Log db Stats Takes db client Retusn nothing

func LogError(msg string, params ...interface{})
    LogError logs an ERROR message

func LogINFO(msg string, params ...interface{})
    LogINFO logs an INFO message.

func QConnect(qURI string) (conn *amqp.Connection, err error)
    QConnect Connects to rabbitmq Takes qURI Returns conn, error

func QConsumer(prefetchCount int, q string, ack bool, channel *amqp.Channel) (msgChan <-chan amqp.Delivery, err error)
    QConsumer Start Consumer Takes prefetchCount, q, ack, channel Returns <-
    chan Delivery, error

func QPublish(channel *amqp.Channel, exchange, routingKey string, data interface{}) (err error)
    QPublish Publishes data to exchange Takes channel, exchange, routingKey,
    data, Returns error

func ReadFromCache(ctx context.Context, key string, client *redis.Client) (data interface{}, err error)
    ReadFromCache Reads data from redis Takes context, key, client Returns data,
    error

func SaveToCache(ctx context.Context, key string, data interface{}, expiry time.Duration, client *redis.Client) (err error)
    SaveToCache Saves data to redis Takes context, key, data, expiry, client
    Returns error

