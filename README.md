# dynts-bann3r

> A simple application written in go to provide a dynamic updated banner for your teamspeak 3 server with multiple information.

## Installation

The recommended way to install dynts-bann3r is using the provided docker container. You can either build the docker container yourself or use the pre-built one.

### Build the docker container

*Requirements:*

- docker

```bash
make docker_build
```

### Run the docker container

```bash
	docker run --rm -it -p 9000:9000 \
        -v path/to/config.json:/config.json \
        -v path/to/template.png:/template.png \
        -v /etc/localtime:/etc/localtime:ro \
        dynts-bann3r
```

To configure the application you have to mount a `config.json` and the template banner `banner.png` into the docker container. It is also recommended do mount the localtime of your host system to have correct datetime in the container.

### Build the project without docker

*Requirements:*

- golang

You can use the make target `make build` to build a binary of the project. After you run the command you can find the binary in `out`. To run the application place the `config.json` and the banner template picture named as `banner.png` in the same directory as the binary.

## Configuration

The configuration is managed by the `config.json` which needs to be placed in the same path as the binary or be mounted into the docker container.

A example configuration looks like this:

```json
{
  "refreshInterval": 60,
  "connection": {
    "host": "TEAMSPEAK_HOST",
    "port": 10011,
    "serverId": 1,
    "user": "TEAMSPEAK_USER",
    "password": "TEAMSPEAK_PW"
  },
  "labels": [
    {
      "text": "%timeHH%:%timeMM%",
      "x": 460,
      "y": 185,
      "fontSize": 45,
      "font": "Arial.ttf",
      "color": "#000000"
    }
  ]
}
```

The configuration file has three parts. The `refreshInterval` is the intervals in second within the banner is updated with information from your teamspeak 3 server. 

The `connection` part configures the connection to your teamspeak 3 server:

| key        | default       | description                               |
| ---------- | ------------- | ----------------------------------------- |
| `host`     |               | The hostname or ip of your ts3 server     |
| `port`     | `10011`       | Server query port of the server           |
| `serverId` | `1`           | Virtual server id                         |
| `user`     | `serveradmin` | A server query user (See [Permissions](#permissions)) |
| `password` |               | Password for the servery query user       |

The `labels` part configures the shown information on the banner. You can add as many labels you want and configure each one individually. A label is configured in the following way:

| key        | default   | description                                                                                                                    |
| ---------- | --------- | ------------------------------------------------------------------------------------------------------------------------------ |
| `text`     |           | Custom text and placeholder which are replaced with servery query information or time. (See [Placeholder](#placeholder) for further details) |
| `x`        |           | The `x` position of the label in pixel.                                                                                        |
| `y`        |           | The `y` position of the label in pixel.                                                                                        |
| `fontSize` | `16`      | Font size of the label in pixel                                                                                                |
| `font`     | `16`      | Path to a font (See [Fonts](#fonts))                                                                                                   |
| `color`    | `#000000` | hex color of the label                                                                                                         |

## Placeholder

You can dynamically configure the labels by using placeholders in the `text` key of the `label` configuration. The following placeholders are available:

| placeholder       | description                                                 |
| ----------------- | ----------------------------------------------------------- |
| `%clientsonline%` | It shows the count of all clients that are currently online |
| `%maxclients%`    | It shows the maximal available slots on the server          |
| `%timeHH%`        | Local time hours e.g. `11`                                  |
| `%timeMM%`        | Local time minutes e.g. `24`                                |
| `%timeSS%`        | Local time seconds e.g. `55`                                |

There are also special placeholders which need arguments to work. Arguments are given comma-separated within square brackets behind the placeholders name (e.g `%placeholder[1,2,3,4]%`) :

| placeholder      | description                                                                                                                  |
| ---------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `%groupcount[]%` | It counts all online clients in the specified groups. The groups can be specified with their group id in the square brackets |

## Fonts

Currently there is a choice of available fonts. In future releases you should be able to integrate custom fonts.
You can use them by assigning one of the following fonts to the `font` key of your `label` configuration.

- `Arial.ttf`

## Permissions

The server query user needs permissions to fetch information from your teamspeak server. Here is a detailed view which permission is needed for which information.

## Developing setup

*Requirements:*

- golang or docker

You can either use docker as described at the beginning to run the project but it is easier for developing to use go directly.

You can use `make run` to run the application in development mode. It will install all go dependencies and run it.

### Testing

The project is tested by several unit tests. You can run them either directly with go by using

```bash
make test_unit
```

or by using docker

```bash
make docker_test
```