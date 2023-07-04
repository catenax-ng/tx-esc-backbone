# Chain upgrades
This guide targets for tangible understanding, how to implement chain upgrades.

We are currently using Cosmos SDK v0.47.2 .
The guide about chain upgrades can be found [here](https://docs.cosmos.network/v0.47/building-modules/upgrade.html).

We currently use the ignite cli v0.27.1 to scaffold chains. Install with `curl https://get.ignite.com/cli@v0.27.1! | bash`.

## When to use the x/upgrade module's migration mechanism.
There are three cases, when to use the migration mechanism.
1. The set of modules changes (addition and removal of modules)
2. The structure of the data stored in the KV stores changes
3. The consensus version of a module changes.

### The set of modules changes
The set of modules is stored in the `UpgradeKeeper`'s module version map.
#### Adding a module
https://docs.cosmos.network/v0.47/core/upgrade.html#adding-new-modules-during-upgrades
An added module is initially not reflected in the `UpgradeKeeper`'s module version map (called `fromVM` in upgrade handlers (from-version-map)).

The `x/upgrade` module does call InitGenesis
#### Removing a module
Due to the composition of Cosmos SDK apps, it's just almost only removing the the modules code from the app.
But you should also consider of cleaning the `UpgradeKeeper`'s module version map (called `fromVM` in upgrade handlers (from-version-map)) and remove
the corresponding KV stores.

### The structure of the data stored in the KV stores changes
Restructuring a module's data (keys, types etc.) can be done with in place migrations

### The consensus version of a module changes.
See the consensus version of a module as the major version of an API. You need to increase it, when breaking changes are done.
Incompatible API or storage changes (see also previous section). 
You can take the messages of a module as the API. 
If the change of the messages are incompatible, you need to change the consensus version.  

## Module migration order
Modules are migrated in alphabetical order. `x/auth` is an exception and migrated last.
Use `func (m *Manager) SetOrderMigrations(moduleNames ...string)` to specify the order.

## Detailed x/upgrade module's migration mechanisms
Let's create a new chain, which should support energy trading between neighbors.
```shell
ignite scaffold chain energytrade --no-module --address-prefix nrg && cd energytrade
```
For simplicity, we assume each household has an account at the chain.
This chain has no special functionality beyond the vanilla Cosmos SDK.
The binary is build with `ignite chain build` at the newly generated folder.
You can start the chain with `ignite chain serve`.

### Preparations to save the state for migrations
Because ignite is a tool for development and mostly recreates all on the fly, we need to take some actions to experience the migration steps.
#### Add mnemonics to accounts.
For convenience add `mnemonic` fields to the `accounts` list entries the `./config.yml`.
Mnemonics can be generated with `energytraded keys mnemonic`.
For example:
```yaml
accounts:
- name: alice
  mnemonic: exchange spatial furnace pear sauce gap fresh normal lobster gold claim beef print royal bulk thought interest grass expire limit control indicate certain bind
  coins:
  - 20000token
  - 200000000stake
- name: bob
  mnemonic: twice eyebrow inherit bring daughter afford forget motor grunt voice prepare dilemma broccoli iron modify usual one grass sunset truth matter duty away cash
  coins: 
  - 10000token
  - 100000000stake
validators:
- name: alice
  bonded: "100000000stake"
...
```
By setting this, ignite will use these instead of randomly generating them at `ignite chain serve`
```shell
git add -A
git commit -m "Add mnemonic"
```
#### Set the proposal voting period
The default voting period is 2 weeks. We reduce this for testing purposes to a minute.
Add the following lines to the `./config.yml`:
```shell
cat << EOF >> config.yml
genesis:
  app_state:
    gov:
      params:
        voting_period: "60s"
EOF
git add -A
git commit -m "Set proposal voting period to one minute"
```
#### Start with ignite at a non-default chain home
When starting the chain with ignite, the chain home is located at `~/.energytrade`.

Because we mainly use ignite to set up the chain and start it later directly with `energytraded start`,
we start once 
```shell
ignite chain serve --home ~/.energytrade-mig
```
As soon as you see the endpoint urls you can abort the process with `ctrl+c`.
```
🌍 Tendermint node: http://0.0.0.0:26657
🌍 Blockchain API: http://0.0.0.0:1317
🌍 Token faucet: http://0.0.0.0:4500
```
The folder `~/.energytrade-mig` is now precious to us. 
It contains the chain home, at which we from now on can operate at with `energytraded --home ~/.energytraded-mig`.

#### Use the validator operator's key for later voting
At the standard implementation migrations take place, after a software upgrade proposal was accepted.
To vote for a proposal, the account has to have a stake. 
In the above snippet of the `config.yaml` you can see, that alice has staked enough to be a validator.

We want to use alice for later interaction with the chain. Ignite imported the key for us already.

### Prepare Cosmovisor
To handle different versions of the chain, we will use git tags.
We will check out these tags into separate working trees and build the version with this `Makefile`: 
```shell
cat <<'EOF' >> Makefile
#!/usr/bin/make -f

BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
COMMIT := $(shell git log -1 --format='%H')

# don't override user values
ifeq (,$(VERSION))
  VERSION := $(shell git describe --exact-match 2>/dev/null)
  # if VERSION is empty, then populate it with branch's name and raw commit hash
  ifeq (,$(VERSION))
    VERSION := $(BRANCH)-$(COMMIT)
  endif
endif

SDK_PACK := $(shell go list -m github.com/cosmos/cosmos-sdk | sed  's/ /\@/g')
CMT_VERSION := $(shell go list -m github.com/cometbft/cometbft | sed 's:.* ::') # grab everything after the space in "github.com/cometbft/cometbft v0.34.7"
DOCKER := $(shell which docker)
BUILDDIR ?= $(CURDIR)/build

export GO111MODULE = on

# process build tags

build_tags = netgo

ifeq (cleveldb,$(findstring cleveldb,$(NRG_BUILD_OPTIONS)))
  build_tags += gcc cleveldb
endif
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=energytrade \
<tab>          -X github.com/cosmos/cosmos-sdk/version.AppName=energytraded \
<tab>          -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
<tab>          -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
<tab>          -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)" \
<tab>          -X github.com/cometbft/cometbft/version.TMCoreSemVer=$(CMT_VERSION)

ifeq (cleveldb,$(findstring cleveldb,$(NRG_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif
ifeq (,$(findstring nostrip,$(NRG_BUILD_OPTIONS)))
  ldflags += -w -s
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'
# check for nostrip option
ifeq (,$(findstring nostrip,$(NRG_BUILD_OPTIONS)))
  BUILD_FLAGS += -trimpath
endif

$(info $$BUILD_FLAGS is [$(BUILD_FLAGS)])

all: install

BUILD_TARGETS := build install

build: BUILD_ARGS=-o $(BUILDDIR)/

$(BUILD_TARGETS): go.sum $(BUILDDIR)/ 
<tab>go $@ -mod=readonly $(BUILD_FLAGS) $(BUILD_ARGS) ./...

$(BUILDDIR)/: 
<tab>mkdir -p $(BUILDDIR)/
EOF
 ( TAB=$'\t' ; sed -i "s/<tab>/$TAB/" Makefile )
```

Add supplementary scripts for easier version handling
```shell
mkdir scripts
cat << EOF > scripts/prepare-version.sh
#!/bin/bash

SCRIPT_LOCATION=\$( dirname -- "\${BASH_SOURCE[0]}")
COMMITISH=\${1:?"Provide a tag, commit or branch"}
COPY_TARGET=\${2:?"Proved copy target location for the binary"}
WORKTREE=\${SCRIPT_LOCATION}/worktrees/\$COMMITISH

\$SCRIPT_LOCATION/get-commitish.sh \$COMMITISH


cd \$WORKTREE
make build
cd - > /dev/null
# search the build folder
# then skip the build folder name in the search results
# then select the first result as binary
BINARY=\`find \$WORKTREE/build -print |  tail -n +2 | head -1 \`

mkdir -p \$COPY_TARGET
cp \$BINARY \$COPY_TARGET/\$(basename \$BINARY)
EOF
cat << EOF > scripts/get-commitish.sh
#!/bin/bash

SCRIPT_LOCATION=\$( dirname -- "\${BASH_SOURCE[0]}")
COMMITISH=\${1:?"Provide a tag, commit or branch"}
WORKTREE=\${SCRIPT_LOCATION}/worktrees/\$COMMITISH

if [ -d \$WORKTREE ] ; then
  echo "already created \$WORKTREE"
else
  # go-git up to v5.4.2 cannot handle git worktree correctly
  # git worktree add \$WORKTREE \$COMMITISH
  mkdir -p \$WORKTREE
  # TODO put the content of .gitignore into GLOBIGNORE
  ROOTPATH=\`realpath \$SCRIPT_LOCATION/..\`
  export GLOBIGNORE=".:..:\${SCRIPT_LOCATION##*/}"
  shopt -s dotglob
  cd \$ROOTPATH
  cp -r * \$WORKTREE/
  cd \$WORKTREE
  git checkout \$COMMITISH
fi

EOF
cat << EOF > scripts/del-commitish.sh
#!/bin/bash

SCRIPT_LOCATION=\$( dirname -- "\${BASH_SOURCE[0]}")

COMMITISH=\${1:?"Provide a tag, commit or branch"}

WORKTREE=\${SCRIPT_LOCATION}/worktrees/\$COMMITISH

if [ ! -d \$WORKTREE ] ; then
  echo "\$WORKTREE does not exist"
else
  # go-git up to v5.4.2 cannot handle git worktree correctly
  #git worktree remove \$WORKTREE
  rm -rf \$WORKTREE
fi
EOF
cat << EOF > scripts/start-cosmovisor.sh
#!/bin/bash
export CHAIN_HOME=\${1:?"chain home required"}

# search the genesis bin folder
# then skip the bin folder name in the search results
# then select the first result as binary
export DAEMON_NAME=\`find \$CHAIN_HOME/cosmovisor/genesis/bin -print |  tail -n +2 | head -1 \`
export DAEMON_NAME=\${DAEMON_NAME##*/}
export DAEMON_HOME=\`realpath \$CHAIN_HOME\`
echo "DAEMON_NAME: \$DAEMON_NAME"
echo "DAEMON_HOME: \$DAEMON_HOME"
cosmovisor run start --home \$DAEMON_HOME
EOF
cat << EOF > scripts/prepare-cosmovisor.sh
#!/bin/bash

SCRIPT_LOCATION=\$( dirname -- "\${BASH_SOURCE[0]}")

if ! command -v cosmovisor &> /dev/null
then
  go install github.com/cosmos/cosmos-sdk/cosmovisor/cmd/cosmovisor@v1.3.0
fi

CHAIN_HOME_DIR=\${1:?"chain home required"}
GENESIS=\${2:?"genesis/ initial version required"}

versions=(\$@)
versions=("\${versions[@]:2}")

if [ -d \$CHAIN_HOME_DIR/cosmovisor ]; then
  echo "\$CHAIN_HOME_DIR/cosmovisor already exists. Exiting prepare-cosmovisor.sh"
  exit
fi

\$SCRIPT_LOCATION/prepare-version.sh \$GENESIS \$CHAIN_HOME_DIR/cosmovisor/genesis/bin/

for version in "\${versions[@]}"
do
  \$SCRIPT_LOCATION/prepare-version.sh \$version \$CHAIN_HOME_DIR/cosmovisor/upgrades/\$version/bin/
done
EOF
echo "scripts/worktrees" >> .gitignore
chmod u+x scripts/*.sh

```

Commit the supplementary files with the following commands.
```shell
git add -A
git commit -m "Add Makefile and scripts."
```

The previously prepared scripts work on git tags. To have the initial version of our chain implementation, we should tag the current commit.
We name our first version the genesis version, to do so, we execute:
```shell
git tag genesis
```
If we now execute `git log`, we should see `(HEAD -> master, tag: genesis)` at our latest commit.

Next we set up [Cosmovisor](https://docs.cosmos.network/v0.47/tooling/cosmovisor). This is a tool, which observes our validator process and restarts it in case of abortion.
Do so by executing:
```shell
scripts/prepare-cosmovisor.sh ~/.energytrade-mig/ genesis
```
This installs the Cosmovisor and builds the chain implementation binary of our genesis version.

### Start the chain with Cosmovisor
Try running the chain with Cosmovisor:
```shell
scripts/start-cosmovisor.sh ~/.energytrade-mig/
```
The process can be stopped by pressing `ctrl+c`. This might take some time.
Cosmovisor should now run the `genesis` version of our `energytraded` binary.
Block creation can be seen by messages like `indexed block height=49 module=txindex`
Concretely Cosmovisor always invokes `~/energytrade-mig/cosmovisor/current`, which is a soft link, maintained by Cosmovisor.
Doing `ls -l ~/energytrade-mig/cosmovisor` will show, to which `current` is currently pointing to.

The foundation to upgrade the chain is done now. Next the module can be implemented

## Add the module
To have a reliable basis for payment, all households track their consumption and production periodically.
We add the module with ignite:
```shell
ignite scaffold module energytrade -y
```
We add a structure to track the consumption and production by time and household
```shell
ignite scaffold map prosumption consumption:int production:int --index time:int --index household:string --module energytrade -y
```
Ignite created a reasonable CRUD implementation for us at `./x/energytrade/msg_serrver_prosumption.go`. 

This generated code will be part of our version `v1.0.0`.
```shell
git add -A
git commit -m "Add energytrade module with tracking"
```
### Add the upgrade handler for v1.0.0
Finally we need to an (upgrade handler)[https://docs.cosmos.network/v0.47/core/upgrade#upgrade-handlers], which will take the necessary steps.
For our upgrade, we do not need to do anything special.
We notify the store loader to initialize our store key by setting `Added: []string{energytradetypes.StoreKey}, `.
```shell
cat << EOF > app/upgrade-v1.0.0.go
package app

import (
    storetypes "github.com/cosmos/cosmos-sdk/store/types"
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/cosmos/cosmos-sdk/types/module"
    upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
    energytradetypes "energytrade/x/energytrade/types"
)

func RegisterV1_0_0(app *App, cfg module.Configurator) {
    const versionName = "v1.0.0"
    upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
    if err != nil {
        // no upgrade file at the disk
        panic(err)
    }
    // if the current height matches the upgrade height
    if upgradeInfo.Name == versionName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
        storeUpgrades := getStoreUpgrades()
        // add store upgrades if necessary
        if !IsEmpty(&storeUpgrades) {
            // configure store loader that checks if version == upgradeHeight and applies store upgrades
            app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
        }
    }
    // ad the actual upgradeHandler
    app.UpgradeKeeper.SetUpgradeHandler(versionName, migrateToV1(app, cfg))
}

func IsEmpty(s *storetypes.StoreUpgrades) bool {
    return len(s.Added) == 0 && len(s.Renamed) == 0 && len(s.Deleted) == 0
}

func getStoreUpgrades() storetypes.StoreUpgrades {
    return storetypes.StoreUpgrades{
        Added: []string{energytradetypes.StoreKey},
    }
}
func migrateToV1(app *App, cfg module.Configurator, skipInitGenesis ...module.AppModule) upgradetypes.UpgradeHandler {
    return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
        for _, m := range skipInitGenesis {
			if withConsensusVer, ok := m.(module.HasConsensusVersion); ok {
				fromVM[m.Name()] = withConsensusVer.ConsensusVersion()
			}
        }
        return app.mm.RunMigrations(ctx, cfg, fromVM)
    }
}
EOF
cat << EOF >> app/app.go

// RegisterUpgradeHandlers register for the different versions the relevant UgradeHandlers
func (app *App) RegisterUpgradeHandlers(cfg module.Configurator) {
    RegisterV1_0_0(app, cfg)
}
EOF
```
We have to make the `app/app.go` invoke the method `RegisterUpgradeHandlers`.
```shell
    sed -i "s/app.mm.RegisterServices(app.configurator)/app.mm.RegisterServices(app.configurator)\n        app.RegisterUpgradeHandlers(app.configurator)/" app/app.go
```

Finally we commit the changes and tag it as version `v1.0.0`.
```shell
git add -A
git commit -m "Register upgrade handler for v1.0.0"
git tag v1.0.0
```

### Start the chain
Open an additional terminal and start Cosmovisor from the project folder with
```shell
scripts/start-cosmovisor.sh ~/.energytrade-mig/
```
The next steps are executed in the previous terminal.

### Prepare the version upgrade
The Cosmovisor currently only knows the `genesis` version.
To provide the `v1.0.0` we execute:
```shell
scripts/prepare-version.sh v1.0.0 ~/.energytrade-mig/cosmovisor/upgrades/v1.0.0/bin
```
The validator is now prepared to upgrade to the new version `v1.0.0`.
### Set variables for the different client versions
If we execute `energytraded`, the OS will look up the path to find the binary and most likely find it at the GOROOT.

Because we want to use more than the currently installed version as client, we point with variables to the versions, we built with `scripts/prepare-version.sh` invocations.
And also export some of the account addresses for convenience.
```shell
CLIENT_GENESIS=(scripts/worktrees/genesis/build/energytraded --home $HOME/.energytrade-mig)
CLIENT_V100=(scripts/worktrees/v1.0.0/build/energytraded --home $HOME/.energytrade-mig)

alice=$($CLIENT_GENESIS keys show alice -a)
bob=$($CLIENT_GENESIS keys show bob -a)
```

### Try invoking the added message type
We use the v1.0.0 client to query for consumption data.
```shell
$CLIENT_V100 query energytrade list-prosumption
```
We should see an error stating `unknown request`

### Propose a version upgrade
Proposal consist of multiple steps (see [gov module](https://docs.cosmos.network/v0.47/modules/gov/#concepts)).
The submission, the voting and the tally.
We submit as `alice` our proposal for a software upgrade with the genesis client.
We create a `tx`. We want to make a transaction with `gov`-module's `submit-proposal` message type. Our proposal is an `software-upgrade` to the version `v1.0.0`.
The `height` value is the block number, at which all validators should stop and switch to the new version. 
For testing we use the height output from the validator log (`indexed block height=### module=txindex`) plus an offset (e.g. `60`).
We `title` and `summary` are for human interpretation and serve no technical purpose.
With `--from alice` we instruct the client to sign the transaction with alice's key.
The final `--yes` state, that we want to broadcast that transaction.
```shell
GOV_MOD_ACC=$($CLIENT_GENESIS query auth module-account gov -o json | jq -r  '.account.base_account.address')
BLOCK_OFFSET=60
cat <<EOF> proposal.json
{
  "messages": [
    {
      "@type": "/cosmos.upgrade.v1beta1.MsgSoftwareUpgrade",
        "authority": "$GOV_MOD_ACC",
        "plan": {
          "name": "v1.0.0",
          "time": "0001-01-01T00:00:00Z",
          "height": "$(( $($CLIENT_GENESIS status | jq -r ".SyncInfo.latest_block_height")+$BLOCK_OFFSET))",
          "info": "{\"linux/amd64\":\"file:///home/wel1wa3/git/catenax/energytrade/scripts/worktrees/v1.0.0/build/energytraded?checksum=af\"}",
          "upgraded_client_state": null
        }
    }
  ],
  "title": "Upgrade to version v1.0.0",
  "summary": "To support energytrading we need this upgrade.",
  "deposit": "10000000stake"
}
EOF
$CLIENT_GENESIS tx gov submit-proposal proposal.json  --from $alice --yes
```

We can see all proposals with
```shell
$CLIENT_GENESIS query gov proposals
```
At the response is the `proposal_id: ##` of interest.
To prevent spam, proposals have to have [deposit](https://docs.cosmos.network/v0.47/modules/gov/#deposit), before they 
can be voted on. This is now already done at proposal submission.

By querying the proposals again, we can not observe, that `voting_start_time` and `voting_end_timed` are now set to actual values.
During this time we can vote for or against the upgrade.
### Accept the version upgrade
To accept a proposal 2/3 of the stake have to vote yes.
Our network consist of only one staker (alice), by voting as alice, we directly set the outcome.
Again we need to set the correct proposal id (the `1` after `vote`).
```shell
$CLIENT_GENESIS tx gov vote 1 yes --from $alice --yes
```
The `yes` after the `1` is the vote on the proposal. The final `--yes` is a shortcut to broadcast that vote transaction.
When the voting time is over the proposal contains the voting result
```yaml
  final_tally_result:
    abstain: "0"
    "no": "0"
    no_with_veto: "0"
    "yes": "100000000"
```

### Application of the software upgrade
The validator exits with the following message on reaching the upgrade height of the proposal
```
3:57PM ERR UPGRADE "v1.0.0" NEEDED at height: ####:
3:57PM ERR CONSENSUS FAILURE!!! err="UPGRADE \"v1.0.0\" NEEDED at height: ####: " module=consensus stack=...
```
The chain generates a .json file, which the Cosmovisor looks into and updates the `~/.energytrade-mig/cosmovisor/current` to the according versioned binary at `~/.energytrade-mig/cosmovisor/upgrades/<version>/bin/`.
Next the validator binary starts and looks for an upgrade handler for the version specified at the `.json` file.
If the binary does not have such a handler, the validator binary will stop again.

### Using the new module
Now that the new implementation is in place, we can track our production and consumption.
Given we are bob, we track `10` production and `5` consumption at time `0` and our household is named `bob`
```shell
$CLIENT_V100 tx energytrade create-prosumption 0 bob 5 10 --from bob --yes
```
This time querying for prosumptions gives us the information we just tracked for bob.
```shell
$CLIENT_V100 query energytrade list-prosumption
```
The output looks like this:
```yaml
pagination:
  next_key: null
  total: "0"
prosumption:
- consumption: 5
  creator: nrg18a5hfm097s7ykpww8qu3v5af8hljmf5s3ng6ed
  household: bob
  production: 10
  time: 0
```
Track 3 further records 
```shell
$CLIENT_V100 tx energytrade create-prosumption 1 bob 8 10 --from bob --yes
$CLIENT_V100 tx energytrade create-prosumption 0 alice 10 8 --from alice --yes
$CLIENT_V100 tx energytrade create-prosumption 1 alice 8 8 --from alice --yes
```
