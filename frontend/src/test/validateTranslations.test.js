import validateTranslation from "./validateTranslations"

import en from "../translations/en"
import zhtw from "../translations/zhtw"

const langs = {
    en: en,
    zhtw: zhtw
}

test(
    "validate translations",
    () => expect(validateTranslation(langs)).toEqual([])
)