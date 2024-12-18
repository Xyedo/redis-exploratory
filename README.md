# Redis is not just a Cache
## Redis Exploratory using Golang Example


### Message Queue with Push Based using List
It using Producer and consumer with Blocking Consumer when empty list
it also can have multiple producer, multiple consumer

### PubSub with Pull Based using Streams
It using Publisher and Consumer with Pull Based based on lastId


### Optimistic-Locking like auction app with thousand client trying to change to current bid
It use WATCH CLI and whenever the watched value key is changed the other client get aborted and return error redis.Nil

## Still Learnin...
