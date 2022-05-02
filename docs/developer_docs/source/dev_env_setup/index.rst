.. _dev_env_setup:

Development environment setup
=============================

Prerequisites
-------------

Install Ignite-cli
^^^^^^^^^^^^^^^^^^

Follow  `these instructions <https://docs.ignite.com/guide/install.html>`_


Install nodejs with npm
^^^^^^^^^^^^^^^^^^^^^^^

Suggestion is to use `Node Version Manager <https://github.com/nvm-sh/nvm#installing-and-updating>`_ for installing nodejs

Serve the development chain
---------------------------

Clone the repository
^^^^^^^^^^^^^^^^^^^^

.. code-block::

   $ git clone git@github.com:catenax/esc-backbone.git

Serve the development chain
^^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block::

   $ cd esc-backbone && ignite chain serve

Serve the vue web app
^^^^^^^^^^^^^^^^^^^^^

Open a new terminal

.. code-block::

   $ cd esc-backbone/vue && npm install && npm run build && npm run serve