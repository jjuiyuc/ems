# DER-EMS

## Backend

### Prerequisite

Make Kafka, MySQL, and [model files](#generate-model-files-by-sqlboiler) ready.

### Run Backend Worker with Go 1.16 and above
- Take `weather-worker` for example
```shell
$ cd backend/daemon/weather-worker

$ go run weather-worker.go
# Start worker with the default config file `backend/config/template.yaml`.
$ go run weather-worker.go -d <config_path> -e <yaml_filename>
# Run worker with your own configuration with YAML with `-d` and `-e` flag.
```

### Generate model files by sqlboiler
- Install [sqlboiler](https://github.com/volatiletech/sqlboiler) by
  ```shell
  # Go 1.16 and above:
  $ go install github.com/volatiletech/sqlboiler/v4@latest
  $ go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest
  ```
- Option 1: Copy `sqlboiler.toml.der-ems` to `sqlboiler.der-ems.toml` to generate corresponding model files.
  ```shell
  $ cp sqlboiler.toml.der-ems sqlboiler.der-ems.toml

  # Modify credential, whitelist in sqlboiler.der-ems.toml if necessary.

  $ sqlboiler -c sqlboiler.der-ems.toml mysql
  # The model files should be generated in `models/der-ems`
  ```
- Option 2: Prepare `sqlboiler.der-ems.toml` ready, then run `make gen-models`

### Deployment (systemd)
- Make sure the model codes have been generated with the correct databases.
- Create the user `derems`  if user not existed.
  ```shell
  $ adduser derems --disabled-password
  ```
- Switch to `backend` directory
  ```shell
  $ cd backend
  ```
- Prepare `config/derems.yaml`
- Deploy by running
  ```shell
  $ make systemd
  ```
- Check log by
  ```shell
  # Take `weather-worker` for example
  $ sudo journalctl -f -u weather-worker
  ```
