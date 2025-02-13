commands used 
// curl -i -G -d "id=UCwFl9Y49sWChrddQTD9QhRA&part=statistics&key=API-KEY" https://www.googleapis.com/youtube/v3/channels

Tip - curl is a fantastic tool for testing API endpoints and I honestly wish I’d invested more time learning it when I just started my career. A simple curl command can save you a lot of time debugging.


So, as we are going to be using an API Key that is from our own personal Google accounts, we want to ensure that we aren’t exposing these to the rest of the world if we commit our project to Git. An excellent way of preventing this from happening is from never hard coding any credentials and to use environment variables.

$ export YOUTUBE_KEY=YOUR-KEY-FROM-DEVELOPER-CONSOLE
$ export CHANNEL_ID=YOUR-YOUTUBE-CHANNEL-ID
