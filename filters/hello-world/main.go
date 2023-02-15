package main

import (
  "math/rand"
  "time"

  "github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
  "github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

const tickMilliseconds uint32 = 1000

func main() {
  proxywasm.SetVMContext(&vmContext{})
}

type vmContext struct {
  types.DefaultVMContext
}

func (*vmContext) NewPluginContext(contextID uint32) types.PluginContext {
  return &helloWorld{}
}

type helloWorld struct {
  types.DefaultPluginContext
  contextID uint32
}

func (ctx *helloWorld) OnPluginStart(pluginConfigurationSize int) types.OnPluginStartStatus {
  rand.Seed(time.Now().UnixNano())

  proxywasm.LogInfo("OnPluginStart from Go!")
  if err := proxywasm.SetTickPeriodMilliSeconds(tickMilliseconds); err != nil {
    proxywasm.LogCriticalf("failed to set tick period: %v", err)
  }

  return types.OnPluginStartStatusOK
}

func (ctx *helloWorld) OnTick() {
  t := time.Now().UnixNano()
  proxywasm.LogInfof("It's %d: random value: %d", t, rand.Uint64())
  proxywasm.LogInfof("OnTick called")
}

func (*helloWorld) NewHttpContext(uint32) types.HttpContext {
  return &types.DefaultHttpContext{}
}

