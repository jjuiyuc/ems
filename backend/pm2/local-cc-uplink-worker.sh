#!/bin/bash

KAFKA_SCRIPT=/opt/kafka/bin/kafka-console-producer.sh
TMP_DATA_FILE=/home/derems/tmp/cc-mock-data.json
BROKER="10.5.10.24:9092"
TOPIC="iot.cc.fct.record.0"
GWID="04F1FD6D9C6F64C3352285CCEAF59EE1"

mock_data_substr1="{\"gwID\":\"${GWID}\",\"timestamp\":"
mock_data_substr2=",\"gridIsPeakShaving\":0,\"loadGridAveragePowerAC\":10,\"batteryGridAveragePowerAC\":0,\"gridContractPowerAC\":15,\"loadPvAveragePowerAC\":20,\"loadBatteryAveragePowerAC\":0,\"batterySoC\":80,\"batteryProducedAveragePowerAC\":20,\"batteryConsumedAveragePowerAC\":0,\"batteryChargingFrom\":\"Solar\",\"batteryDischargingTo\":\"\",\"pvAveragePowerAC\":40,\"loadAveragePowerAC\":30,\"loadLinks\": {\"grid\":0,\"battery\":0,\"pv\":0},\"gridLinks\": {\"load\":1,\"battery\":0,\"pv\":0},\"pvLinks\": {\"load\":1,\"battery\":1,\"grid\":0},\"batteryLinks\": {\"load\":0,\"pv\":0,\"grid\":0},\"batteryPvAveragePowerAC\":20,\"gridPvAveragePowerAC\":0,\"gridProducedAveragePowerAC\":10,\"gridConsumedAveragePowerAC\":0,\"batteryLifetimeOperationCycles\":16,\"batteryProducedLifetimeEnergyAC\":500,\"batteryConsumedLifetimeEnergyAC\":500,\"batteryAveragePowerAC\":-3.5,\"batteryVoltage\":28}"

send_cc_mock_data() {
    echo "send cc mock data"
    timestamp=$(date +%s)
    echo "timestamp = ${timestamp}"
    echo ${mock_data_substr1}${timestamp}${mock_data_substr2} > ${TMP_DATA_FILE}
    cat ${TMP_DATA_FILE}
    ${KAFKA_SCRIPT} --broker-list ${BROKER} --topic ${TOPIC} < ${TMP_DATA_FILE}
    rm ${TMP_DATA_FILE}
}

while true
do
    send_cc_mock_data
    sleep 1m
done