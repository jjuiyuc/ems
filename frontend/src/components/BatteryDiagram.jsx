import styled from "styled-components"

const stopKeyframe = (level, direction) => {
    if (level == 100 && direction == "dischargingTo") return ""
    const percentage = direction == "dischargingTo" ? (100 - level) : level
    return `${percentage}% {
      height: ${level}%;
    }`
}

const Animation = styled.div`
  @keyframes animation {
    from {
      height: 0%;
    }
    80% {
      height: ${(props) => (props.level) + "%"};
    }
    to {
      height: ${(props) => (props.level) + "%"};
    }
  }
  animation-duration: ${(props) => (props.level / 10 * 1.1) + "s"};
  animation-iteration-count: infinite;
  animation-name: ${(props) => props.level == 100 && props.direction == "chargingFrom"
        ? "" : "animation"};
  animation-timing-function: linear;`

export default function BatteryDiagram(props) {

    const { direction, target } = props
    const state = props.state > 100
        ? 100
        : (props.state < 0 ? 0 : props.state)
    const lines = Array.from(Array(10).keys()).map((key) => (
        <div
            className={"border-gray-100" + (key < 9 ? " border-b-2" : "")}
            key={"line-" + key}
        />
    ))
    return (
        <div className="flex flex-col items-center w-24">
            <div className="bg-gray-400 h-3 w-10 rounded-t-md" />
            <div className="bg-gray-400 h-48 p-2 rounded-lg w-full">
                <div className="bg-gray-100 h-full p-1 rounded-md">
                    <div className="h-full overflow-hidden relative rounded">
                        {target == ""
                            ? <>
                                <div className={"absolute w-full bg-primary-500" +
                                    (direction === "chargingFrom" ? "bottom-0" : "")}
                                    style={{
                                        height: (state || 0) + "%",
                                        top:
                                            direction === "dischargingTo"
                                                ? (100 - (state || 0)) + "%"
                                                : ""
                                    }}>
                                </div>
                            </>
                            : <>
                                <div className={"absolute w-full " +
                                    (direction === "chargingFrom" ? "bottom-0 bg-primary-200" : "bg-primary-500")}
                                    style={{
                                        height: (state || 0) + "%",
                                        top:
                                            direction === "dischargingTo"
                                                ? (100 - (state || 0)) + "%"
                                                : ""
                                    }}></div>
                                <Animation
                                    className={
                                        "absolute w-full " +
                                        (direction === "chargingFrom" ? "bottom-0 bg-primary-500" : "bg-primary-200")
                                    }
                                    direction={direction}
                                    level={(state || 0)}
                                    style={{
                                        height: (state || 0) + "%",
                                        top:
                                            direction === "dischargingTo"
                                                ? (100 - (state || 0)) + "%"
                                                : ""
                                    }}
                                />
                            </>}
                        <div className="absolute grid h-full w-full">{lines}</div>
                    </div>
                </div>
            </div>
        </div>
    )
}