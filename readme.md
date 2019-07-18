Postmark sender
=====

[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-brightgreen.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![Build Status](https://travis-ci.org/paysuper/postmark-sender.svg?branch=master)](https://travis-ci.org/paysuper/postmark-sender) 
[![codecov](https://codecov.io/gh/paysuper/postmark-sender/branch/master/graph/badge.svg)](https://codecov.io/gh/paysuper/postmark-sender)
[![Go Report Card](https://goreportcard.com/badge/github.com/paysuper/postmark-sender)](https://goreportcard.com/report/github.com/paysuper/postmark-sender)

RabbitMQ consumer to send emails through [postmark service](https://postmarkapp.com)

## Environment variables:

| Name                            | Required | Default                                        | Description                                                                                                                         |
|:--------------------------------|:--------:|:-----------------------------------------------|:------------------------------------------------------------------|
| BROKER_ADDRESS                  | -        | amqp://127.0.0.1:5672                          | RabbitMQ url address                                              |
| POSTMARK_API_URL                | -        | https://api.postmarkapp.com/email/withTemplate | Postmark API url                                                  |
| POSTMARK_API_TOKEN              | true     | -                                              | Postmark API security token                                       |
| POSTMARK_EMAIL_SENDER           | true     | -                                              | Sender email to send emails to users                              |
| POSTMARK_TEMPLATE_CONFIRM_EMAIL | true     | -                                              | Template alias in postmark for emails for user email confirmation |


## Contributing
We feel that a welcoming community is important and we ask that you follow PaySuper's [Open Source Code of Conduct](https://github.com/paysuper/code-of-conduct/blob/master/README.md) in all interactions with the community.

PaySuper welcomes contributions from anyone and everyone. Please refer to each project's style and contribution guidelines for submitting patches and additions. In general, we follow the "fork-and-pull" Git workflow.

The master branch of this repository contains the latest stable release of this component.