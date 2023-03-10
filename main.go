package main

import (
	"log"

	"github.com/Shopify/go-lua"
);

var script = `
bestBid = bestBid();
bestAsk = bestAsk();
spread = math.abs(bestBid - bestAsk);
print(spread);
`;

func main() {
	luaState := lua.NewState();
	lua.OpenLibraries(luaState);

	registerBestAsk(luaState);
	registerBestBid(luaState);

	if err := lua.DoString(luaState, script); err != nil {
		log.Fatal(err)
	}

}

func registerBestBid( luaState *lua.State){
	luaState.Register("bestBid", func(luaState *lua.State)  int {
		luaState.PushInteger(10000);
		return 1;
	});
}

func registerBestAsk( luaState *lua.State){
	luaState.Register("bestAsk", func(luaState *lua.State) int {
		luaState.PushInteger(10001);
		return 1;
	});
}
