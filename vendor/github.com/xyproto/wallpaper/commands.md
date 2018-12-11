# Command collection

## Cinnamon

     gsettings set org.cinnamon.desktop.background picture-uri  "file:///filename"

## Gnome 2

    gconftool-2 –type string –set /desktop/gnome/background/picture_filename “/full/path/to/file.png”

## Gnome 3

    gsettings set org.gnome.desktop.background picture-uri file:///$PATH_TO_FILE

## i3, XMonad, others

    feh --bg-scale /path/to/image

## Sway

    swaymsg 'output "*" background $HOME/somewhere.png fill'

## KDE / Plasma

    dbus-send --session --dest=org.kde.plasmashell --type=method_call /PlasmaShell org.kde.PlasmaShell.evaluateScript 'string:
    var Desktops = desktops();
    for (i=0;i<Desktops.length;i++) {
            d = Desktops[i];
            d.wallpaperPlugin = "org.kde.image";
            d.currentConfigGroup = Array("Wallpaper",
                                        "org.kde.image",
                                        "General");
            d.writeConfig("Image", "file:///PATH/TO/IMAGE.png");
    }'
