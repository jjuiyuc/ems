import styled from "styled-components"
import { useState } from "react"

const stopKeyframe = (level) => `
  ${(1 - 1 / (level + 1)) * 100}% {
    height: ${level * 10}%;
  }`

const Animation = styled.div`
  @keyframes animation {
    from {
      height: 0%
    }
    ${(props) => stopKeyframe(props.level)}
    to {
      height: ${(props) => props.level * 10 + "%"}
    }
  }
  animation-duration: ${(props) => props.level + 1 + "s"}
  animation-iteration-count: infinite
  animation-name: animation;
  animation-timing-function: linear`

export default function BatteryDiagram() {

    const [batteryLevel] = useState(6),
        [direction] = useState("charge")

    const lines = Array.from(Array(10).keys()).map((key) => (
        <div
            className={"border-white" + (key < 9 ? " border-b-2" : "")}
            key={"line-" + key}
        />
    ))
    return (
        <div className="p-2">
            <div className="flex flex-col items-center w-24">
                <div className="bg-gray-400 h-3 w-10 rounded-t-md" />
                <div className="bg-gray-400 h-48 p-2 rounded-lg w-full">
                    <div className="bg-white h-full p-1 rounded-md">
                        <div className="h-full overflow-hidden relative rounded">
                            <Animation
                                className={
                                    "absolute bg-green-500 w-full " +
                                    (direction === "charge" ? "bottom-0" : "")
                                }
                                level={batteryLevel}
                                style={{
                                    top:
                                        direction === "discharge"
                                            ? 100 - batteryLevel * 10 + "%"
                                            : ""
                                }}
                            />
                            <div className="absolute grid h-full w-full">{lines}</div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}