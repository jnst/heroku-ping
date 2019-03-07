# heroku-ping

## Heroku

[Free Dyno Hours](https://devcenter.heroku.com/articles/free-dyno-hours)

>If an app has a free web dyno, and that dyno receives no web traffic in a 30-minute period, it will sleep.

heroku-ping is a package that sends a ping continually so that app does not sleep.

## How to use

Check your heroku web url.

```sh
$ heroku info -s  | grep web_url | cut -d= -f2
https://abc-xyz-12345.herokuapp.com/
```

Set `HEROKU_URL` to config vars.

```sh
$ heroku config:set HEROKU_URL=$(heroku info -s  | grep web_url | cut -d= -f2)
```
