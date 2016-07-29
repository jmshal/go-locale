package go_locale

func DetectLocale() (string, error) {
    out, err := getCommandOutput("wmic", "os", "get", "locale")
    if err != nil {
        return "", err
    }

    out = strings.Replace(out, "Locale", "", -1)
    out = strings.TrimSpace(out)

    id, err := strconv.Atoi(out)
    if err != nil {
        return "", err
    }

    lcid := LCID()
    locale, err := lcid.ById(id)
    if err != nil {
        return "", err
    }

    return locale, nil
}
