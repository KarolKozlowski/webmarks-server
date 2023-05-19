
# WebMarks

>**Disclaimer**: This is my `golang` learning project

**WebMarks** is supposed to be a simple web service that allows creating shortcuts accessible directly via URL.

The shortcuts are defined as URL path, for example let's assume the following shortcut definition:
```
example -> http://example.com
```

Entering WebMarks at the following URL: `https://webmarks/example` will trigger a redirect to `http://example.com`.

You can then set up a service within your local DNS search suffix under the hostname `goto` (or similar) which will enable you to enter `goto/example` in the browser's location bar.

