

in case you do not want to install sphinx to your system python
```shell
virtualenv .
source bin/activate
pip --require-virtualenv  install -r requirements.txt 
```

to generate the html run
```shell
source bin/activate
make html
```