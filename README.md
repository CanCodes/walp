# Walp 
Random wallpapers for your desktop environment.

## How Does It Work?
Walp uses the [Reddit API](https://www.reddit.com/dev/api/) to get the top posts of the month from `r/wallpaper`. It then filters out the posts that are not images and then randomly selects one of the images to set as the wallpaper.

## Installation
Check the [releases page]() for the latest version. Download the executable for your platform.

## How Do I Use It?
To use walp, download the latest release from the [releases page]() and run the executable. After running the executable, walp will appear in your system tray. `Right click` on the walp icon and select `Get New Wallpaper` to change your wallpaper. You can also select `Quit` to exit the program. 

## How Do I Build It?
To build walp, you will need to have [Go](https://golang.org/) installed. Once you have Go installed, run the following commands:
```bash
$ git clone https://github.com/cancodes/walp.git
$ cd walp
$ go build
```

## Notes

#### Platform Support
**Walp has only been tested on MacOS X** and theoeretically, it should work on Windows and Linux as well. However, I have not tested it on those platforms. If you would like to help, please submit a pull request.

#### More Wallpaper Sources
Walp currently only uses `r/wallpaper` as the source for wallpapers. in the future, I would like to add more sources. If you have any suggestions, please submit an issue with a link to the source API.

#### Better UI
Walp currently has a very basic UI and it is using the native system tray menu of the platform. I would like to add a better UI in the future. If you have any suggestions, please submit an issue. 

## Disclaimer
This project is not affiliated with Reddit in any way. This project is not endorsed or certified by Reddit. All trademarks are the property of their respective owners. Wallpapers are also the property of their respective owners.
I am not responsible for any damages caused by this program. Use at your own risk.