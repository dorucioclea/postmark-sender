# PaySuper Postmark Sender

[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-brightgreen.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/paysuper/postmark-sender/issues)
[![Build Status](https://travis-ci.org/paysuper/postmark-sender.svg?branch=master)](https://travis-ci.org/paysuper/postmark-sender) 
[![codecov](https://codecov.io/gh/paysuper/postmark-sender/branch/master/graph/badge.svg)](https://codecov.io/gh/paysuper/postmark-sender)
[![Go Report Card](https://goreportcard.com/badge/github.com/paysuper/postmark-sender)](https://goreportcard.com/report/github.com/paysuper/postmark-sender)

PaySuper Postmark Sender is a RabbitMQ consumer to send emails using the [Postmark Service](https://postmarkapp.com).

***

## Table of Contents

- [Usage](#usage)
- [Usage example](#usage-example)
- [Contributing](#contributing-feature-requests-and-support)
- [License](#license)

## Usage

The application handles all configuration from the environment variables.

### Environment variables:

| Name                            | Required | Default                                        | Description                                                                                                                             |
|:--------------------------------|:--------:|:-----------------------------------------------|:----------------------------------------------------------------------------------------------------------------------------------------|
| BROKER_ADDRESS                  | -        | amqp://127.0.0.1:5672                          | The RabbitMQ URL address.                                                                                                                    |
| POSTMARK_API_URL                | -        | https://api.postmarkapp.com/email/withTemplate | The Postmark API URL.                                                                                                                        |
| POSTMARK_API_TOKEN              | true     | -                                              | The Postmark API security token.                                                                                                             |
| POSTMARK_EMAIL_FROM             | true     | -                                              | The sender email to send emails to users.                                                                                                    |
| POSTMARK_EMAIL_CC               | -        | ""                                             | The CC recipient email address. Multiple addresses are comma separated. Max 50.                                                              |
| POSTMARK_EMAIL_BCC              | -        | ""                                             | The BCC recipient email address. Multiple addresses are comma separated. Max 50.                                                             |
| POSTMARK_EMAIL_TRACK_OPENS      | -        | false                                          | Activate the open tracking for all emails.                                                                                                   |
| POSTMARK_EMAIL_TRACK_LINKS      | -        | ""                                             | Activate the link tracking for links in the HTML or Text bodies of this email. Possible options: None, HtmlAndText, HtmlOnly, TextOnly.      |

## Usage example:

```go
package main

import (
    postmarkSdrPkg "github.com/paysuper/postmark-sender/pkg"
    "github.com/streadway/amqp"
    "gopkg.in/ProtocolONE/rabbitmq.v1/pkg"
    "log"
)

func main()  {
    broker, err := rabbitmq.NewBroker("amqp://127.0.0.1:5672")
    
    if err != nil {
        log.Fatalf("Creating RabbitMQ publisher failed with error: %s\n", err)
    }
    
    payload := &postmarkSdrPkg.Payload{
        TemplateAlias: "template_name",
        TemplateModel: map[string]string{"param1": "value1"},
        To:            "emai@test.com",
    }
    
    err = broker.Publish(postmarkSdrPkg.PostmarkSenderTopicName, payload, amqp.Table{})
    
    if err != nil {
        log.Fatalf("Publication message to queue failed with error: %s\n", err)
    }
}
```

## Contributing, Feature Requests and Support

If you like this project then you can put a ⭐ on it. It means a lot to us.

If you have an idea of how to improve PaySuper (or any of the product parts) or have general feedback, you're welcome to submit a [feature request](../../issues/new?assignees=&labels=&template=feature_request.md&title=).

Chances are, you like what we have already but you may require a custom integration, a special license or something else big and specific to your needs. We're generally open to such conversations.

If you have a question and can't find the answer yourself, you can [raise an issue](../../issues/new?assignees=&labels=&template=issue--support-request.md&title=I+have+a+question+about+<this+and+that>+%5BSupport%5D) and describe what exactly you're trying to do. We'll do our best to reply in a meaningful time.

We feel that a welcoming community is important and we ask that you follow PaySuper's [Open Source Code of Conduct](https://github.com/paysuper/code-of-conduct/blob/master/README.md) in all interactions with the community.

PaySuper welcomes contributions from anyone and everyone. Please refer to [our contribution guide to learn more](CONTRIBUTING.md).

## License

The project is available as open source under the terms of the [GPL v3 License](https://www.gnu.org/licenses/gpl-3.0).