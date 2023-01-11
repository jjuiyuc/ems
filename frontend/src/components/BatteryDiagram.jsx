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

    const { direction } = props
    const state = props.state > 100
        ? 100
        : (props.state < 0 ? 0 : props.state)
    const lines = Array.from(Array(10).keys()).map((key) => (
        <div
            className={"border-white" + (key < 9 ? " border-b-2" : "")}
            key={"line-" + key}
        />
    ))
    return (
        <div className="flex flex-col items-center w-24">
            <div className="bg-gray-400 h-3 w-10 rounded-t-md" />
            <div className="bg-gray-400 h-48 p-2 rounded-lg w-full">
                <div className="bg-white h-full p-1 rounded-md">
                    <div className="h-full overflow-hidden relative rounded">
                        <div className={"absolute  w-full " +
                            (direction === "chargingFrom" ? "bottom-0" : "")}
                            style={{
                                backgroundColor:
                                    direction === "dischargingTo"
                                        ? "#12c9c9"
                                        : "#c4f2f2",
                                height: (state || 0) + "%",
                                top:
                                    direction === "dischargingTo"
                                        ? (100 - (state || 0)) + "%"
                                        : ""
                            }}></div>
                        <Animation
                            className={
                                "absolute bg-primary-500 w-full " +
                                (direction === "chargingFrom" ? "bottom-0" : "")
                            }
                            direction={direction}
                            level={(state || 0)}
                            style={{
                                backgroundColor:
                                    direction === "dischargingTo"
                                        ? "#c4f2f2"
                                        : "#12c9c9",
                                height: (state || 0) + "%",
                                top:
                                    direction === "dischargingTo"
                                        ? (100 - (state || 0)) + "%"
                                        : ""
                            }}
                        />
                        <div className="absolute grid h-full w-full">{lines}</div>
                    </div>
                </div>
            </div>
        </div>
    )
}