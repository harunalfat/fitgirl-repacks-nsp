# FITGIRL-REPACKS-NSP

This is a tool/executable for Windows to extract [Fitgirl Repack](https://fitgirl-repacks.site/) switch emulated package game into its own NSP. You may need to get standalone NSP file format, maybe to play into your own CFW Switch or another emulator (Ryujinx, etc)

### How to RUN
- Download the latest `extractor.exe` binary from [release page](https://github.com/harunalfat/fitgirl-repacks-nsp/releases)
- Copy the downloaded `extractor.exe` into the root of switch emulated Fitgirl-repacks directory. Example directory is `"C:/Games/Mario Kart 8 Deluxe"`
- Double click or run the `extractor.exe` file, and the default name for output file is `out.nsp`. You may also run it from CMD/Powershell and changing the output file name
```
extractor.exe -o mario-kart-8-deluxe.nsp
```

### How to COMPILE
- You need to have Golang with version >= 1.16
- Run command below
```
GOOS=windows GOARCH=amd64 go build extractor.go
```

### Dependency
- [nspBuild](https://github.com/CVFireDragon/nspBuild/releases) by CVFireDragon

