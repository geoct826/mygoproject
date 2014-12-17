
     ,-----.,--.                  ,--. ,---.   ,--.,------.  ,------.
    '  .--./|  | ,---. ,--.,--. ,-|  || o   \  |  ||  .-.  \ |  .---'
    |  |    |  || .-. ||  ||  |' .-. |`..'  |  |  ||  |  \  :|  `--, 
    '  '--'\|  |' '-' ''  ''  '\ `-' | .'  /   |  ||  '--'  /|  `---.
     `-----'`--' `---'  `----'  `---'  `--'    `--'`-------' `------'
    ----------------------------------------------------------------- 


Hi there! Welcome to Cloud9 IDE!

To get you started, create some files, play with the terminal,
or visit http://docs.c9.io for our documentation.
If you want, you can also go watch some training videos at
http://www.youtube.com/user/c9ide.

Happy coding!
The Cloud9 IDE team

# Setting up Google App Engine on Cloud9

## 1. Download and setup the latest Google App Engine SDK for Linux 64 bit.
```
* cd $HOME
* wget https://storage.googleapis.com/appengine-sdks/featured/go_appengine_sdk_linux_amd64-1.9.15.zip
* unzip go_appengine_sdk_linux_amd64-1.9.15.zip
* vim ~/.bashrc
* add "export PATH=$HOME/go_appengine:$PATH"
```

## 2. Go AppEngine SDK's development server is hardcoded to bind to localhost. Need to change that for Cloud9 IDE.
```
* cd go_appengine/google/appengine/tools/devappserver2/
* vim devappserver2.py
* go to the main() section towards the end of file 
* add the following below "options = PARSER.parse_args()"
  * options.host = os.environ['IP']
  * options.api_host = os.environ['IP']
  * options.admin_host = os.environ['IP']"
```

## 3. Create a new runner to start appengine 
```
{
    "cmd" : ["$HOME/go_appengine/dev_appserver.py", "$HOME/workspace/"]
}
Note: Cloud9 only supports one public port, as a result, only :8080 will be available through the preview site. To use the App Engine admin, a port forwarding solution would be required to port :8000 to a handler.
```

# Add in Zurb Foundation response front-end framework

# Add in Backbone


