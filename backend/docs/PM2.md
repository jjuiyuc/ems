# Install PM2 and PM2 modules

### [PM2](https://pm2.keymetrics.io/docs/usage/quick-start/)
- Make sure NVM and Node.js are installed

NVM (version: 0.39.1)
  ```shell
  # Install & update script
  $ curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh | bash
  # Check version
  $ nvm -v
  ```
Node.js  (version: 14.19.3)
  ```shell
  # List available versions
  $ nvm ls-remote
  # Install LTS version
  $ nvm install 14.19.3
  # Check version
  $ node -v
  ```
PM2 (version: 5.2.0)
  ```shell
  # Install the latest version
  $ npm install pm2@latest -g
  # Check version
  $ pm2 -v
  ```

### [pm2-logrotate](https://github.com/keymetrics/pm2-logrotate)
  ```shell
  # Install
  $ pm2 install pm2-logrotate
  # Update config
  $ pm2 set pm2-logrotate:retain 10
  $ pm2 set pm2-logrotate:max_size 256M
  $ pm2 set pm2-logrotate:compress true
  ```

### [pm2-health](https://github.com/pankleks/pm2-health)
  ```shell
  # Install
  $ pm2 install pm2-health
  # Update config: host, port, from, user, password, mailTo, events
  $ vim ~/.pm2/module_conf.json
  # Restart
  $ pm2 restart pm2-health

  # Test email
  $ pm2 trigger pm2-health mail
  ```

  ```json
	"pm2-health": {
		"smtp": {
			"host": "smtp.gmail.com",
			"port": 587,
			"from": "DER-EMS-Monitor <gitlab@ubiik.com>",
			"user": "gitlab@ubiik.com",
			"password": "",
			"secure": false,
			"disabled": false
		},
		"mailTo": "xxx@ubiik.com",
		"replyTo": "",
		"batchPeriodM": 0,
		"batchMaxMessages": 0,
		"events": [
			"exit",
			"restart"
		],
		"exceptions": true,
		"messages": true,
		"messageExcludeExps": [],
		"metric": {},
		"metricIntervalS": 60,
		"aliveTimeoutS": 300,
		"addLogs": true,
		"appsExcluded": [],
		"snapshot": {
			"url": "",
			"token": "",
			"auth": {
				"user": "",
				"password": ""
			},
			"disabled": false
		}
	}
  ```