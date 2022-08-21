# Learn Go with Tests: Blinking LED with TinyGo (WIP)

Going to try and make an exciting TinyGo chapter for [learn go with tests](https://quii.gitbook.io/learn-go-with-tests/), this repo is for me to play around

## Notes

- No excuses!
    - TDD NASA Mars rockets, what's your excuse?
    - Physical hardware, embedded symstems, IOT, etc all have different testing problems **at the edge**, but if we're separating accidental complexity away from actual complexity, we can still unit test a hell of a lot of our work
    - Do the top down thing, "i know i need a thing that turns a light on or off, don't care how right now"
- Feedback loops of TinyGo improved with testing
    - Show the manual steps from the tutorial
        - > How bad would this be for something more complicated than blinking an LED

## Info on arduino nano 33

I bought one. Make sure you have a spare micro-b usb as it doesn't come with one. Finally, the box full of miscellaneous cables was useful for me.

https://tinygo.org/docs/reference/microcontrollers/arduino-nano33/

```shell
brew install bossa
```



To flash the Arduino's ROM

```bash
tinygo flash -target=arduino-nano33
```



## Setting up IDE

https://blog.jetbrains.com/go/2021/06/02/tinygo-for-tiny-applications-discover-a-new-plugin-for-goland/

It's really nice! Flashing via IDE is pretty swish