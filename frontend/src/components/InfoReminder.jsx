import { connect } from "react-redux"
import { Alert } from "@mui/material"
import { useTranslation } from "react-multi-lang"

const mapState = state => ({ lang: state.lang.value })

export default connect(mapState)(function InfoReminder(props) {
    const t = useTranslation()

    return (
        <Alert variant="outlined" severity="info" sx={{ borderRadius: "1.5rem" }}>
            <p>{t("info.pleaseNote")}</p>
            <p>
                {props.lang === "en"
                    ? "The values shown here are costs and savings estimated using inverter’s measurement data and the utility tariff data. Your actual utility bill may look different from the data displayed in this page due to the gaps between the utility meter measurement and the inverter measurement. In any discrepancies, your actual utility bill shall take precedence."
                    : <>
                        本頁面所顯示之能源成本及節省電費均為根據逆變器測量資料和用戶電價費率所計算之結果，
                            <span className="text-red-500 font-bold">僅供參考所用</span>
                            。由於逆變器採集之數據或與電力公司電表實際量測數據略有差異，因此您真實的電費帳單也許會和本頁面所顯示之數字有所出入。最終實際電費以您的電費帳單為準，
                            <span className="text-red-500 font-bold">前述之誤差無法做為電費爭議所用</span>。
                        </>
                }
            </p>
        </Alert>
    )
})