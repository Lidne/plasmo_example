import React from "react"

export const config = {
  matches: ["https://*/*"],
  world: "MAIN"
}

const WasmComponent = () => {
  return (
    <div>
      <button
        onClick={() => {
          window.hello()
        }}></button>
    </div>
  )
}

export default WasmComponent
