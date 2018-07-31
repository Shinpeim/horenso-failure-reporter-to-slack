# horenso-failure-reporter-to-slack

## what's this

horenso-failure-reporter-to-slack is a reporter for [horenso](https://github.com/Songmu/horenso) that post the failuer notifications to Slack.

It dose nothing when the command exited with exitcode 0 but notify failure report when the command exited with exitcode non-zero.

## install

go get github.com/Shinpeim/horenso-failure-reporter-to-slack

## usage

```
$ horenso --reporter horenso-failure-reporter-to-slack -- /path/to/yourjob
```
