# ULCL Traffic Influence

## Test Command

```bash
./ci-test-ulcl.sh TestULCLTrafficInfluence
```

## Test File

- [ulclTrafficInfluence_test.go](../test/goTest/ulclTrafficInfluence_test.go)

## Test Cases

1. Before Traffic Influence
   - Ping n6gw: expected ping success
   - Ping mec: expected ping failed
2. After Traffic Influence
   - Ping n6gw: expected ping failed
   - Ping mec: expected ping success
3. Reset Traffic Influence
   - Ping n6gw: expected ping success
   - Ping mec: expected ping failed

## Test Steps

1. Post ue subscription data to db via web console's api
2. Activate PacketRusher
3. Run [test cases](#test-cases)
4. Deactivate PacketRusher
5. Delete ue subscription data from db via web console's api
