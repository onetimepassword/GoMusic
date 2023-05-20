# GoMusic

https://github.com/DylanMeeus/MediumCode/blob/master/Audio/FirstSound/main.go
```shell
go run jingle.go jingle.dat jingle.bin
ffplay -f f32le -ar 44100 -showmode 1 -autoexit -i jingle.bin
ffmpeg -f f32le -ar 44100 -ac 1 -i jingle.bin -acodec mp3 track01.mp3
```
[![YouTube](https://img.youtube.com/vi/rDn8jC1J7Z4/0.jpg)](https://www.youtube.com/watch?v=rDn8jC1J7Z4)

