import repeatedTranslationKeys from "./repeatedTranslationKeys"

import en from "../translations/en"
import zhtw from "../translations/zhtw"

const ignoredList = [
    "analysis",
    "batteryDischarge",
    "chargeToBattery",
    "chargingFrom",
    "dashboard",
    "demandCharge",
    "economics",
    "exportToGrid",
    "forgotPassword",
    "importFromGrid",
    "load",
    "resetPassword",
    "thisMonth"
]

test(
    "validate repeated keys",
    () => expect(repeatedTranslationKeys(en, ignoredList)).toEqual([]),
    () => expect(repeatedTranslationKeys(zhtw, ignoredList)).toEqual([]),
)