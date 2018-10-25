# VLC subtitles

Small go based tool to download subtitles for TV shows from addic7ed using
[matcornic/addic7ed](https://github.com/matcornic/addic7ed).

It will search the subtitles for the specified video file and download them in
the same folder with the same filename and extension `.srt`.

Optionally it can also launch VLC for the specified filename. Examples:

```bash
$ vlc-subtitles /home/user/videos/show.S01E01.720.mkv
```
Downloads `/home/user/videos/show.S01E01.720.mkv.srt`

```bash
$ vlc-subtitles /home/videos/show.S01E01.720.mkv --launch
```
Downloads `/home/user/videos/show.S01E01.720.mkv.srt` and starts playing the
video on VLC with the subtitles.

## Compile

```bash
$ go fetch github.com/matcornic/addic7ed
$ go build .
```

## Usage with other media players

If your video player recognizes `*.srt` files, use this tool inside a bash wrapper and then launch your prefered program:

```
#!/bin/bash
vlc-subtitles $1
myplayer $@
```

### TODO

  - try relative paths
  - use Flags to parse argument flags
