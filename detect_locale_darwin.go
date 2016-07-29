package go_locale

func DetectLocale() (string, error) {
    return getCommandOutput("defaults", "read", "-g", "AppleLocale")
}
