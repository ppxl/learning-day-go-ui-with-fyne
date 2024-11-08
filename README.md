# learning-day-go-ui-with-fyne
I am learning how to build cross devices applications can be developed with fyne.io

Source: https://fyne.io/

## Making it work as android app

1. install NDK somewhere and create an env var for that
   1. `wget https://dl.google.com/android/repository/android-ndk-r27c-linux.zip`
   2. unzip it f. i. to `/home/YOURHOME/apps`
   3. `export ANDROID_NDK_HOME=/home/YOURHOME/apps/android-ndk-r27c`
2. install Android tooling for `adb`-Support and make the binaries findable
   1. `wget https://dl.google.com/android/repository/commandlinetools-linux-11076708_latest.zip`
   2. unzip it f. i. to 
   3. `export PATH="/home/YOURHOME/apps/android-cmdline-tools/bin/:/home/YOURHOME/apps/android-cmdline-tools/platform-tools/:$PATH"
   4. `sdkmanager --sdk_root=/home/YOURHOME/apps/android-cmdline-tools --install "platform-tools"`
   5. test that `adb` is runnable: `adb --version"
3. Prepare your Android phone - **this is insecure and no fun at all but works with Android 14 as of 2024**
   1. enable USB debugging so that your phone can automatically receive the app from your computer
      1. go to your Phone settings > Phone info > Software info > tap 7 times the Build number
      2. you should see a message of the developer mode being enabled
      3. go back to the settings, enter the new entry "Developer Options"
      4. \o/ SWITCH ALL THE GOOD SWITCHES!!
   2. disable Google Play protection
      1. open the Google Play app
      2. tap on your profile app for more options
      3. tap on the entry "Play Protect"
      4. disable play protect
4. now adb install commands work. Try make `install-android`
5. after development: **DO NOT FORGET TO ENABLE Play Protect AGAIN!** 

Developmental builds can be run on the laptop with `go run . -tags mobile app.go`

## A note on the binary name

the resulting .apk file will be after `Name` property in [FyneApp.toml](FyneApp.toml). The build number increases on each run of `make build-android` or `fyne package` 