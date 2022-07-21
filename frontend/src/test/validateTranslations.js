const validateTranslations = langs => {
    const allLangs = Object.keys(langs)

    const flatternedMap = allLangs.reduce((map, lang) => {
        Object.keys(langs[lang]).forEach(category => {
            Object.keys(langs[lang][category]).forEach(item => {
                if (typeof(langs[lang][category][item]) === "string") {
                    const mapField = map[category + "." + item]

                    if (mapField) {
                        mapField[lang] = langs[lang][category][item]
                    }
                    else {
                        map[category + "." + item] = {
                            [lang]: langs[lang][category][item]
                        }
                    }
                }
                else {
                    Object.keys(langs[lang][category][item]).forEach(entry => {
                        const mapField = map[`${category}.${item}.${entry}`]

                        if (mapField) {
                            mapField[lang] = langs[lang][category][item][entry]
                        }
                        else {
                            map[`${category}.${item}.${entry}`] = {
                                [lang]: langs[lang][category][item][entry]
                            }
                        }
                    })
                }
            })
        })

        return map
    }, {})

    let missingItems = []

    Object.keys(flatternedMap).forEach(key => {
        allLangs.forEach(lang => {
            if (!flatternedMap[key][lang]) {
                missingItems.push({key, lang})
            }
        })
    })

    return missingItems
}

module.exports = validateTranslations