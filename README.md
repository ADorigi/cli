# checkctl

CLI for opengovernance

![Root Command GIF](./tapes/gif/root.gif)

## Installation

```
brew tap kaytu-io/cli-tap && brew install checkctl
```

## List of Commands

- [configure](#configure)
- [get]()
    - [benchmarks](#benchmarks)
    - compliance-summary-for-benchmark 
    - compliance-summary-for-integration
    - controls 
    - jobs
    - job-details
    - findings
- run compliance
- run discovery

## Configure

Configuration for checkctl

![Configure Command GIF](./tapes/gif/configure.gif)

Interactive mode:
```
checkctl configure 
```
Non-interactive mode:
```
checktl configure --api-key <<api-key>> --app-endpoint https://path.to.app.endpoint --output json
```

---
**_NOTE:_**  Pre-defined integartions can be manually added to the configuration file, by updating the `integrations` field inside `$HOME/.checkctl/config.json`. 

config.json:
```
{
    ...
    "integrations": {"acc3":"id_name=account3"}
}
```

---
## get 

### benchmarks

Get a list of benchmarks

![Get Benchmarks Command GIF](./tapes/gif/getbenchmarks.gif)

Available flags:

| Flag | Description | Default
|----- | ----------- | -------
|`--page-size` | Defines page size of response | 25
|`--page-number` | Defines page number of response| 1
|`--show-only-root` | Show only root benchmarks | true
|`--include-findings-summary` | Include findings summary in response | false

### controls

Get a list of controls 

![Get Controls Command GIF](./tapes/gif/getcontrols.gif)

Available flags:

| Flag | Description | Default
|----- | ----------- | -------
|`--page-size` | Defines page size of response | 25
|`--page-number` | Defines page number of response| 1
