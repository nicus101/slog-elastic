# Slog-Elastic - implementation of slog.Handler for elasticsearch

Golang have structural logging for some time,
named [slog](https://pkg.go.dev/log/slog).
It even has rich and
[Awesome Slog](https://github.com/go-slog/awesome-slog)
ecosystem.

## Why?

Why use *ElasticSearch* directly?
There is [Logstash](https://www.elastic.co/logstash) after all.
But sometimes You...

 - Develop locally, and still want persistency.
 - Have IoT solution, and don't want big guns.
 - Didn't get access to logstash from Your operations department.

This package was made to solve the above.
Having somewhat decent implementation would be better than
cobbling ad-hoc solutions in every project.

## Features

 - Extracts values from *Context* using ContextAttrFunc
 - Works with [zinc](https://zincsearch.com/)
 - Works with [elasticsearch](https://www.elastic.co/elasticsearch/)
 
## Known issues

This package is very young.
It only has [Handler.Handle](https://pkg.go.dev/log/slog#Handler) implemented.
And helper methods to load configuration from `.env`.

## Planned features and to-do

 - Implement other methods of Handler.
 - Ability for client code to filter/rename attributes.
 - Safeguard to propagate all logs before application shutdown.
 - Bulk inserts asynchronously.
