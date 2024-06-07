# sideshow
This repository provides a complete reference implementation for simplifying consistent data ingestion using protocol buffers, connectrpc, spark and delta lake 

## Building and Generating the Local Descriptor for Apache Spark
~~~
make clean && make gen && make build-descriptor
~~~

> Get the Protobuf Descriptor for Apache Spark
~~~
curl -s "https://buf.build/sideshow/coffeeco/descriptor/4d2bc9114f1845598d9d3c480b313bf9" \
    --output ~/Desktop/coffeeco.bin
~~~