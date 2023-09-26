# Messages

## What's Messages?

Messages is a simple message-passing repo for gamedev. It's made to make it possible to pass messages in an ambiguous, but still safely abstracted and Go-idiomatic way.

## How do I get it?

```go get github.com/solarlune/messages```

## How do I use it? 

Check the example. It's pretty straightforward - you simply create a Dispatcher, register listener objects to the dispatcher, and send pre-defined messages through the dispatcher. Registered listeners will receive the messages as necessary.