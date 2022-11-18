.. Copyright (c) 2022 - for information on the respective copyright owner
.. see the NOTICE file and/or the repository at
.. https://github.com/catenax-ng/product-esc-backbone-code
..
.. SPDX-License-Identifier: Apache-2.0

.. _wallet_tutorial:

Wallet tutorial
===============

Prerequisites
-------------

Build and launch the chain and frontend
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

See :ref:`dev_env_setup`

Install Keplr extension
^^^^^^^^^^^^^^^^^^^^^^^

Install the `keplr extension <https://chrome.google.com/webstore/search/keplr>`_ to chrome.

.. image:: images/0_install_keplr_extension.png
   :alt: install keplr extension
   :align: center


Use the wallet browser extension
--------------------------------

Import the example/ test wallet
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Click at the Keplr extension at the top right.


.. image:: images/1_click_on_extension_at_the_top_right.png
   :alt: keplr_extension_icon
   :align: center


If there is no wallet configured you are taken to this page. Click on either on `Create new account` for a new account 
or on `Import existing account` to import an account with mnemonic.

.. image:: images/2_import_existing_account.png
   :alt: create_an_account
   :align: center

**NOTE: The password entered at the next page is used to secure the Keplr walled entirely and not just the account. 
All later added accounts are protected by this password too.**

If you chose creating a new one, you have to enter your mnemonic at the last page by clicking at the words in the correct order to prove, you backed up the mnemonic.

This description chose `Import existing account` to import a pre-funded test account. 
Enter the following mnemonic for the pre-funded test account.
```text
old square lecture frog curtain habit bunker casino awesome defy fashion cry wife rain outer scene fork leaf raven twin hen hurt calm bulb
```
Enter a password to unlock the Keplr extension in the future.
Do not modify the derivation path, which would result in a different account otherwise.

.. image:: images/3_enter_mnemonic_of_test_acc.png
   :alt: enter_mnemonic_of_test_acc
   :align: center

Clicking now on the icon should show an account similar to the image below for the Cosmos Hub.

.. image:: images/4_cosmos_hub_balance_of_test_acc.png
   :alt: cosmos_hub_balance_of_test_acc
   :align: center

Clicking on `Cosmos Hub` shows the default chains, which Keplr knows. Our test chain is not contained in it yet.

.. image:: images/5_test_net_not_available_yet.png
   :alt: test_net_not_available_yet
   :align: center

Importing the local chain into Keplr
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Visit the web app at `http://localhost:3000/ <http://localhost:3000/>`_.

Keplr will ask for importing the suggested chain `catenax-1`. Approve it.

.. image:: images/6_open_vue_app_and_approve_chain_addition.png
   :alt: open_vue_app_and_approve_chain_addition
   :align: center

Now select the testnet `catenax testnet`. 

.. image:: images/7_select_catena_testnet.png
   :alt: select_catena_testnet
   :align: center

This will show the accounts balance from the chain, which was locally start in the first terminal by `starport chain serve`.

.. image:: images/8_balance_of_the_testnet.png
   :alt: balance_of_the_testnet
   :align: center

Importing the testnet chain into Keplr
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Visit the web app at `https://validator1.dev.demo.catena-x.net/ <https://validator1.dev.demo.catena-x.net/>`_.
Keplr will ask for importing the suggested chain `catenax-testnet-1`. Approve it.

.. _request_funds_testnet:

Request funds for the testnet chain
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
Visit `https://validator1.dev.demo.catena-x.net/ <https://validator1.dev.demo.catena-x.net/>`_ and put your public address
for the testnet into the text box and press the `Fund me` button.

Regarding chain suggestions
---------------------------

Suggesting a chain in a web app
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

The following snippet at :download:`KeplrChainSuggestion.tsx<../../../../web/src/KeplrChainSuggestion.tsx>` makes the chain suggestion to the Keplr extension

.. code-block:: javascript

    async function SuggestKeplrChain() {
        if (!window.getOfflineSigner || !window.keplr) {
            alert('Keplr extension not found. Please install it.');
        } else {
            if (window.keplr.experimentalSuggestChain) {
                let suggestion = await createChainSuggestion()
                try {
                    await window.keplr.experimentalSuggestChain(suggestion)
                } catch (e){
                    console.log("ERROR", e);
                    console.log('Chain %s added.',suggestion.chainName);
                    alert(`Add the chain ${suggestion.chainName} failed: ${String(e)}`)
                }
            } else {
                alert('Please use the recent version of Keplr extension');
            }
        }
    }

    async function createChainSuggestion() {
        const catenax_suggestion=await fetch("/chain/catenax-testnet-1-suggestion.json")
        return catenax_suggestion.json();
    }

Generation of the chain suggestion json
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

The :download:`catenax-testnet-1-suggestion.json <../../../../web/public/chain/catenax-testnet-1-suggestion.json>` can be generated
with the :download:`keplr-suggestion.go <../../../../cmd/keplr-suggestion/main.go>` command.
The command is currently missing a useful cli and parameters can be changed in the code.

Further information about the chain suggestion json can be found in `Keplr's documentation <https://docs.keplr.app/api/suggest-chain.html>`_ 
and its `example repository <https://github.com/chainapsis/keplr-example/blob/master/src/main.js>`_.
