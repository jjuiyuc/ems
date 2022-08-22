const repeatedTranslationKeys = (json, ignoredList) => {
    let keys = [], repeatedKeys = []

    const validate = key => {
        if (keys.includes(key) && !ignoredList.includes(key))
            repeatedKeys.push(key)
        else keys.push(key)
    }

    Object.keys(json).forEach(mainKey => {
        const section = json[mainKey]

        Object.keys(section).forEach(subKey => {
            const sub = section[subKey]

            if (typeof sub === "string") validate(subKey)
            else Object.keys(sub).forEach(minorKey => validate(minorKey))
        })
    })

    return repeatedKeys
}

export default repeatedTranslationKeys