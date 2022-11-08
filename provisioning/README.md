/provisioning directory structure



chains

    icon

        bmc

            deploy.sh

            upgrade.sh

            helpers

                add_link.sh

                add_relay_to_link.sh

                set_fee_aggregator.sh

        bsh

            bts

                deploy.sh

                upgrade.sh

                helpers

                    add_bts_to_bmc.sh

                    add_to_blacklist.sh

                    deploy_token.sh

                    ..

            xcall

                ...

        docker

            docker-compose.yml

            Dockerfile

    evm

        bmc

        bsh

        docker

    bsc

        bmc

        bsh

        docker

    near

        docker

    ...

config

    input_params.json

deployments

    localnet

        iconbsc

            contracts

            users

                keys

            bmr.config.json

            contract_addresses.json

            e2e.config.json

            state.json

        iconsnow

            contracts...

            ...state.json

        ...

    mainnet

        contracts...

        ...state.json



    testnet

        iconbsc

            contracts...

            ...state.json

        iconnear

            ontracts...

            ...state.json

        ...

provisions

    localnet.sh

    mainnet.sh

    testnet.sh

    i2b_full.sh

    i2s_link.sh

    upgrades

        hotfix_v0.0.9.sh

        upgrade_v0.0.10.sh

        ...

relay

    docker

    build_relay_image.sh

    deploy_remote.sh

    run_relay.sh





----------------------------------------------------------------------------------------------------------------

/provisioning/chains: 

    holds scripts specific to chain

    can reuse scripts if applicable (e.g. bsc snow chains can reuse scripts present inside chains/evm)



/provisioning/config

    holds input config necessary for deployment and configuration



/provisioning/deployments

    holds file specific to deployment

    should be .gitignored

    holds deployment artifacts e.g. relay config as well as state variables inside state.json

    separate directory for separate testnet, localnet deployments



/provisioning/provisions

    holds task-specific scripts (e.g. for full deployment of some specific chain or linking a new chain to existing deployment, etc)

    wraps scripts inside /provisioning/chains

    responsible for updating state variables, generating and placing deployment artifacts



/provisioning/relay

    relay specific scripts



----------------------------------------------------------------------------------------------------------------



With this approach, 

A new chain integration should include 

    1. addition of /provisioning/chain/ scripts (reuse if applicable)

    2. addition of necessary config 

    3. create necessary wrapper script inside /provisioning/provisions 



A new script that does some specific custom task (whether called directly or through github action)

    1. add script inside /provisioning/provisions for specifc task

    2. script should work with the deployment's state.json (read/write)

