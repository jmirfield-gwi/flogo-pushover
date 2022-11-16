# Pushover Activity

The `flogo-pushover` activity sends a Pushover `message` using `App Token` and `Group Token` provided by Pushover and returns a `status` in the `outputs`.

The available service `settings` are as follows:

| Name   |  Type   | Description   |
|:-----------|:--------|:--------------|
| App Token | string | Application token provided by Pushover |
| Group Token | string | Group/User token provided by Pushover |
| Active | boolean | Controls whether to actually request Pushover API |

The available `input` for the request are as follows:

| Name   |  Type   | Description   |
|:-----------|:--------|:--------------|
| message | string | Message to be sent to Pushover (Cannot be empty)  |

The available response `outputs` are as follows:

| Name   |  Type   | Description   |
|:-----------|:--------|:--------------|
| status | number | Returns 200 if message was sent successfully. \n204 is if message is empty or if active is set to false. 400 is if request is bad. |

## Testing

To run tests, first create a .env file in root directory and add the following values:
```
GROUP=GROUP_TOKEN_PROVIDED_BY_PUSHOVER
APP=APP_OR_USER_TOKEN_PROVIDED_BY_PUSHOVER
```
After .env variables are setup, run:
```
go test
```