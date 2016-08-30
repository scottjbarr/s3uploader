# S3 Uploader

This repository provides a tiny wrapper library for uploading files to
S3.

Also available, is an application which receives Sidekiq jobs from
Redis, uploads to S3.


## Installation

    go get github.com/scottjbarr/s3uploader/cmd/s3worker


## Makefile

There is a [Makefile](Makefile) which documents how to build and run
etc.


## Tests

Well...

There aren't any tests because there is very little to test at this
stage.  The library and worker fill a need I have for a feature right
now. Most of the functionality is in support libraries.


## S3 Worker Application

Install the worker that uploads files to S3.

The worker receives Sidekiq jobs from Redis and simply uploads the
file.

[Gokiq](https://github.com/cupcake/gokiq) is used to handle the
interactions with Redis as it implements a reliable queue as
recommended by Redis.

Example usage showing AWS environment variables (need by aws-sdk-go),
as well as optional environment variables for the worker app.

    AWS_ACCESS_KEY_ID=xx \
    AWS_SECRET_ACCESS_KEY=xx \
    AWS_REGION=ap-southeast-2 \
    go run cmd/s3worker/main.go

You can also provide environment variables to the `s3worker` app to
specify the Redis URL, namespace and the maximum number of workers to
use.

    AWS_ACCESS_KEY_ID=key \
    AWS_SECRET_ACCESS_KEY=secret \
    AWS_REGION=ap-southeast-2 \
    NAMESPACE=yourapp-test \
    REDIS_URL=:6379 \
    WORKER_COUNT=50 \
    go run cmd/s3worker/main.go


## References

- https://github.com/cupcake/gokiq
- https://github.com/mperham/sidekiq
- http://redis.io/commands/rpoplpush
- https://github.com/aws/aws-sdk-go


## License

The MIT License (MIT)

Copyright (c) 2016 Scott Barr

See [LICENSE.md](LICENSE.md)
