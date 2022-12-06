import moment from "moment"

const ConvertTimeToNumber = (string, timezone) => {
    const
        today = moment().format("YYYY-MM-DD"),
        time = moment(today + " " + string + timezone),
        hour = time.hour(),
        minute = time.minute()

    return hour + Number((minute / 60).toFixed(1))
}

const ValidateEmail = email =>
    /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/.test(email)

const ValidatePassword = password =>
    /^(?=.{8,})((?=.*[^a-zA-Z\s])(?=.*[a-z])(?=.*[A-Z])|(?=.*[^a-zA-Z0-9\s])(?=.*\d)(?=.*[a-zA-Z])).*$/.test(password)

const ValidateNum = num =>
    /^(0?|[1-9][0-9]*)$/.test(num)
export { ConvertTimeToNumber, ValidateEmail, ValidatePassword, ValidateNum }