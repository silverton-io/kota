// Copyright (c) 2024 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the Apache-2.0 license, a copy of
// which may be found at https://github.com/silverton-io/kota/blob/main/LICENSE

package consumer

import (
	"context"
	"time"

	awsconf "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/rs/zerolog/log"
	"github.com/silverton.io/kota/pkg/buffer"
	"github.com/silverton.io/kota/pkg/config"
	"github.com/silverton.io/kota/pkg/envelope"
)

type SqsConsumer struct {
	client   *sqs.Client
	queue    string
	buffer   *buffer.Buffer
	ctx      context.Context
	shutdown context.CancelFunc
}

func (c *SqsConsumer) Initialize(config *config.Input, buffer *buffer.Buffer) error {
	ctx, shutdown := context.WithCancel(context.Background())
	log.Debug().Msg("initializing sqs consumer")
	c.ctx, c.shutdown = ctx, shutdown
	c.buffer = buffer
	c.queue = config.Sqs.Queue
	// AWS stuff
	aws_config, err := awsconf.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("could not load aws config")
		return err
	}
	svc := sqs.NewFromConfig(aws_config)
	c.client = svc
	return nil
}

func (c *SqsConsumer) Consume() {
	log.Debug().Msg("starting sqs consumer from queue: " + c.queue)
	go func() {
		for {
			select {
			case <-c.ctx.Done():
				log.Debug().Msg("shutting down sqs consumer")
				return
			default:
				log.Trace().Msg("polling sqs for new messages from queue: " + c.queue)
				results, err := c.client.ReceiveMessage(c.ctx, &sqs.ReceiveMessageInput{
					QueueUrl: &c.queue,
				})
				if err != nil {
					log.Error().Err(err).Msg("could not receive messages from sqs")
				}
				for _, message := range results.Messages {
					// TODO -> Wrap the record in an envelope and pass to buffer
					fake_envelope := envelope.BuildFakeEnvelope()
					c.buffer.Append([]envelope.KotaEnvelope{fake_envelope})
					log.Trace().Msg("deleting message " + *message.MessageId + " from sqs queue: " + c.queue)
					_, err := c.client.DeleteMessage(c.ctx, &sqs.DeleteMessageInput{
						QueueUrl:      &c.queue,
						ReceiptHandle: message.ReceiptHandle,
					})
					if err != nil {
						log.Error().Err(err).Msg("could not delete message from queue")
					}
				}
			}
			time.Sleep(time.Duration(200) * time.Millisecond)
		}
	}()
}

func (c *SqsConsumer) Shutdown() error {
	log.Debug().Msg("shutting down sqs consumer")
	c.shutdown()
	return nil
}

func NewSqsConsumer(config *config.Input, buffer *buffer.Buffer) *SqsConsumer {
	consumer := SqsConsumer{}
	consumer.Initialize(config, buffer)
	consumer.Consume()
	return &consumer
}
