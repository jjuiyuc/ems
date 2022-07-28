import {ConvertTimeToNumber} from "../utils/utils"

test(
    "Convert Time to Number",
    () => expect(ConvertTimeToNumber("07:30:00", "+0800")).toEqual(7.5),
    () => expect(ConvertTimeToNumber("07:30:00", "+0900")).toEqual(6.5),
    () => expect(ConvertTimeToNumber("20:40", "+0800")).toEqual(20.7),
    () => expect(ConvertTimeToNumber("00:15", "+08:00")).toEqual(0.25)
)