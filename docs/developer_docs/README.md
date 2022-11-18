<!--
Copyright (c) 2022 - for information on the respective copyright owner
see the NOTICE file and/or the repository at
https://github.com/catenax-ng/product-esc-backbone-code

SPDX-License-Identifier: Apache-2.0
-->


in case you do not want to install sphinx to your system python
```shell
python3 -m venv sphinx-venv
source sphinx-venv/bin/activate
pip --require-virtualenv  install -r requirements.txt 
```

to generate the html run
```shell
source sphinx-venv/bin/activate
make html
```

When writing documentation follow 
[these conventions](https://www.sphinx-doc.org/en/master/usage/restructuredtext/basics.html#sections) 
for section titles.
