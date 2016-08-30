package main

import (
	"encoding/base64"

	"github.com/cupcake/gokiq"
	"github.com/garyburd/redigo/redis"
	"github.com/namsral/flag"
	"github.com/scottjbarr/s3uploader"
)

// Args receives the Sidekiq job
type Args struct {
	Bucket  string
	Key     string
	ID      string
	Encoded string
}

// S3UploadWorker facilitates handling the Sidekiq job from Redis.
type S3UploadWorker []Args

// Perform handles the job.
func (w *S3UploadWorker) Perform() error {
	args := (*w)[0]

	decoded, err := base64.StdEncoding.DecodeString(args.Encoded)

	if err != nil {
		return err
	}

	uploader := s3uploader.NewUploader(args.Bucket, args.Key, decoded)
	return uploader.Upload()
}

func main() {
	redisNamespace := flag.String("namespace", "", "Redis namespace")
	redisURL := flag.String("redis_url", ":6379", "Redis URL")
	workerCount := flag.Int("worker_count", 50, "Worker count")
	flag.Parse()

	worker := gokiq.NewWorkerConfig()
	worker.WorkerCount = *workerCount
	worker.RedisNamespace = *redisNamespace

	// reconfigure the default Redis pool
	worker.RedisPool = redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial("tcp", *redisURL)
	}, worker.WorkerCount+1)

	worker.Register(&S3UploadWorker{})

	worker.Queues = gokiq.QueueConfig{
		"upload": 1,
	}

	worker.Run()
}
