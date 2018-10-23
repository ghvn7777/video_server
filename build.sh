#!/bin/bash

set -e 
{
cd /home/kaka/kaka/go_workspace/src/video_server/web
go install
cp /home/kaka/kaka/go_workspace/bin/web /home/kaka/kaka/go_workspace/bin/video_server_web_ui/web
cp -R /home/kaka/kaka/go_workspace/src/video_server/templates /home/kaka/kaka/go_workspace/bin/video_server_web_ui
}
