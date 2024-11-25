# Autoscaler

## Notes

### Node
needs per connection 50 - 100ms
= 75ms on average
1000ms / 75ms = 13.33 = 13 connections per second

### Up Scaling
new node once current capacity is at 90% for at least 5 seconds
maximum: all cores / 2

### Down Scaling
kill nodes until current capacity has reached 80%, if current capacity is at 72% for at least 10 seconds
minimum of 2 nodes
do not kill node immediately, kill it gracefully: no new connections, keep alive until all connections to this
node have disconnected

### Load balancing
next node: the one with the least load
if all have +-2% same load then the first node is next

### Extras
- Error handling
  - retry process max 3 times

