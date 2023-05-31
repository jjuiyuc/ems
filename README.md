# DER-EMS

## Backend [![pipeline status](https://gitlab.com/ubiik/ems/der-ems/badges/main/pipeline.svg)](https://gitlab.com/ubiik/ems/der-ems/-/commits/main) [![coverage report](https://gitlab.com/ubiik/ems/der-ems/badges/main/coverage.svg)](https://gitlab.com/ubiik/ems/der-ems/-/commits/main)

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

  $ sqlboiler -c sqlboiler.der-ems.toml --struct-tag-casing camel mysql
  # The model files should be generated in `models/der-ems`
  ```
- Option 2: Prepare `sqlboiler.der-ems.toml` ready, then run `make gen-models`

### Deployment (PM2)
- Make sure the model codes have been generated with the correct databases.
- Create the user `derems`  if user not existed.
  ```shell
  $ adduser derems --disabled-password
  ```
- [Install PM2 and PM2 modules](backend/docs/PM2.md)
- Switch to `backend` directory
  ```shell
  $ cd backend
  ```
- Prepare `config/derems.yaml` and `pm2/derems.config.js`
- Deploy by commands
  ```shell
  # Build code
  $ make go-build

  # Create folders as follows
  /opt/derems
  ├── etc
  ├── sbin

  # Copy binary files
  $ sudo cp dist/* /opt/derems/sbin/
  # Copy specific binary file
  # Take `local-cc-worker` for example
  $ sudo cp dist/local-cc-worker /opt/derems/sbin/

  # Copy config files
  $ sudo cp pm2/derems.config.js /opt/derems/etc
  $ sudo cp config/derems.yaml /opt/derems/etc
  ```
- Run processes
  ```shell
  # Create logs softlink
  sudo ln -s /home/derems/.pm2/logs/ /opt/derems/logs

  cd /opt/derems/etc

  # Start specific process
  # Take `core-local-cc-worker` for example
  pm2 start derems.config.js --only "core-local-cc-worker"

  # Check processes
  pm2 list
  # Stop a process
  pm2 stop "core-local-cc-worker"
  # Restart a process
  pm2 restart "core-local-cc-worker"
  ```

## Frontend

### Deploy with GitLab CI/CD

#### Basic Information
- All the commits can be deployed although there's no any frontend changes in it.
- The version name would be `{Branch Name}-{Commit SHA}`.
- The folder `/var/www/dist` on the target server will be replaced without backup.

#### Steps
1. Go to the [Pipeline](https://gitlab.com/ubiik/ems/der-ems/-/pipelines) page and locate the relevant pipeline.
2. Click stage button to open the job list for that pipeline.
3. Find the job named `deploy-frontend-dev` or `deploy-frontend-production`.
   - Show all sections in TOU page:
      - Click the `play` button to run the job directly.
   - Hide `Mid Peak` sections:
      1. Go to the job page.
      2. Fill out the form with following information:
         - Key: `HIDE_MID_PEAK`
         - Value: `true`
      3. Click `Run job` button below the form to run the job with the specified version number.