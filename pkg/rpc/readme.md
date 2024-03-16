# Go Chia RPC

Library for interacting with Chia RPC. Supports both HTTP and Websocket communications.

## Usage

When creating a new client, chia configuration will automatically be read from `CHIA_ROOT`. If chia is installed for the same user go-chia-rpc is running as, the config should be automatically discovered if it is in the default location. If the config is in a non-standard location, ensure `CHIA_ROOT` environment variable is set to the same value that is used for chia-blockchain.

### HTTP Mode

To use HTTP mode, create a new client and specify `ConnectionModeHTTP`:

```go
package main

import (
	"github.com/chia-network/go-chia-libs/pkg/rpc"
)

func main() {
	client, err := rpc.NewClient(rpc.ConnectionModeHTTP, rpc.WithAutoConfig())
	if err != nil {
		// error happened
	}
}
```

#### HTTP Mode w/ Manual Config

To provide a manual config struct instead of using AutoConfig, use something like the following:

```go
package main

import (
	"github.com/chia-network/go-chia-libs/pkg/config"
	"github.com/chia-network/go-chia-libs/pkg/rpc"
)

func main() {
	client, err := rpc.NewClient(rpc.ConnectionModeHTTP, rpc.WithManualConfig(config.ChiaConfig{
		ChiaRoot: "/path/to/.chia",
		FullNode: config.FullNodeConfig{
			PortConfig: config.PortConfig{
				Port:    8444,
				RPCPort: 8555,
			},
			SSL: config.SSLConfig{
				PrivateCRT: "config/ssl/full_node/private_full_node.crt",
				PrivateKey: "config/ssl/full_node/private_full_node.key",
			},
		},
	}))
}
```

### Websocket Mode

To use Websocket mode, specify ConnectionModeWebsocket when creating the client:

```go
package main

import (
	"github.com/chia-network/go-chia-libs/pkg/rpc"
)

func main() {
	client, err := rpc.NewClient(rpc.ConnectionModeWebsocket, rpc.WithAutoConfig())
	if err != nil {
		// error happened
	}
}
```

Websockets function asynchronously and as such, there are a few implementation differences compared to using the simpler HTTP request/response pattern. You must define a handler function to process responses received over the websocket connection, and you must also specifically subscribe to the events the handler should receive.

#### Handler Functions

Handler functions must use the following signature: `func handlerFunc(data *types.WebsocketResponse, err error)`. The function will be passed the data that was received from the websocket and an error.

Initializing a client, and defining and registering a handler function looks like the following:

```go
package main

import (
	"log"
	"github.com/chia-network/go-chia-libs/pkg/rpc"
	"github.com/chia-network/go-chia-libs/pkg/types"
)

func main() {
	client, err := rpc.NewClient(rpc.ConnectionModeWebsocket, rpc.WithAutoConfig())
	if err != nil {
		log.Fatalln(err.Error())
	}

	client.AddHandler(gotResponse)

	// Other application logic here
}

func gotResponse(data *types.WebsocketResponse, err error) {
	log.Printf("Received a `%s` command response\n", data.Command)
}
```

#### Synchronous Mode

If you want websockets to behave more like request/response style calls, you can enable sync mode.

To make all calls sync by default, you can set an option on the client:

```go
package main

import (
	"fmt"

	"github.com/chia-network/go-chia-libs/pkg/rpc"
)

func main() {
	client, err := rpc.NewClient(rpc.ConnectionModeWebsocket, rpc.WithAutoConfig(), rpc.WithSyncWebsocket())
	if err != nil {
		// error happened
	}

	netInfo, _, err := client.DaemonService.GetNetworkInfo(&rpc.GetNetworkInfoOptions{})
	if err != nil {
		// error happened
	}

	// netInfo has the actual network information, since we're running in sync mode
	fmt.Println(netInfo.NetworkName.OrEmpty())
}
```

You can also temporarily enable synchronous mode and then turn it back off

```go
package main

import (
	"fmt"

	"github.com/chia-network/go-chia-libs/pkg/rpc"
	"github.com/chia-network/go-chia-libs/pkg/types"
)

func main() {
	client, err := rpc.NewClient(rpc.ConnectionModeWebsocket, rpc.WithAutoConfig())
	if err != nil {
		// error happened
	}

	client.AddHandler(gotAsyncResponse)

	client.SetSyncMode()

	netInfo, _, err := client.DaemonService.GetNetworkInfo(&rpc.GetNetworkInfoOptions{})
	if err != nil {
		// error happened
	}
	fmt.Println(netInfo.NetworkName.OrEmpty())

	client.SetAsyncMode()
}

func gotAsyncResponse(data *types.WebsocketResponse, err error) {
	log.Printf("Received a `%s` async command response\n", data.Command)
}
```

The output of this program will look something like the following. Note that both the async handler AND the sync response
variables saw the event and were able to handle it.

```shell
Received a `get_network_info` command response
mainnet
```

#### Subscribing to Events

There are two helper functions to subscribe to events that come over the websocket.

`client.SubscribeSelf()` - Calling this method subscribes to response events for any requests made from this client

`client.Subscribe(service)` - Calling this method, with an appropriate service, subscribes to any events that chia may generate that are not necessarily in responses to requests made from this client (for instance, `metrics` events fire when relevant updates are available that may impact metrics services)

### Get Transactions

#### HTTP Mode

```go
client, err := rpc.NewClient(rpc.ConnectionModeHTTP, rpc.WithAutoConfig())
if err != nil {
	log.Fatal(err)
}

transactions, _, err := client.WalletService.GetTransactions(
	&rpc.GetWalletTransactionsOptions{
		WalletID: 1,
	},
)
if err != nil {
	log.Fatal(err)
}

if transactions.Transactions.IsPresent() {
    for _, transaction := range transactions.Transactions.MustGet() {
        log.Println(transaction.Name)
    }
}
```

#### Websocket Mode

```go
func main() {
	client, err := rpc.NewClient(rpc.ConnectionModeWebsocket, rpc.WithAutoConfig())
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = client.SubscribeSelf()
	if err != nil {
		log.Fatalln(err.Error())
	}

	client.AddHandler(gotResponse)

	client.WalletService.GetTransactions(
		&rpc.GetWalletTransactionsOptions{
			WalletID: 1,
		},
	)
}

func gotResponse(data *types.WebsocketResponse, err error) {
	log.Printf("Received a `%s` command response\n", data.Command)

	if data.Command == "get_transactions" {
		txns := &rpc.GetWalletTransactionsResponse{}
		err = json.Unmarshal(data.Data, txns)
		if err != nil {
			log.Fatalln(err.Error())
		}

		log.Printf("%+v", txns)
	}
}
```

### Get Full Node Status

```go
state, _, err := client.FullNodeService.GetBlockchainState()
if err != nil {
	log.Fatal(err)
}

if state.BlockchainState.IsPresent() {
    log.Println(state.BlockchainState.MustGet().Difficulty)	
}
```

### Get Estimated Network Space

Gets the estimated network space and formats it to a readable version using FormatBytes utility function

```go
//import (
//    "log"
//
//    "github.com/chia-network/go-chia-libs/pkg/rpc"
//    "github.com/chia-network/go-chia-libs/pkg/util"
//)

state, _, err := client.FullNodeService.GetBlockchainState()
if err != nil {
	log.Fatal(err)
}

if state.BlockchainState.IsPresent() {
    log.Println(state.BlockchainState.MustGet().Space)
}
```

### Request Cache

When using HTTP mode, there is an optional request cache that can be enabled with a configurable cache duration. To use the cache, initialize the client with the `rpc.WithCache()` option like the following example:

```go
client, err := rpc.NewClient(rpc.ConnectionModeHTTP, rpc.WithAutoConfig(), rpc.WithCache(60 * time.Second))
if err != nil {
	// error happened
}
```

This example sets the cache time to 60 seconds. Any identical requests within the 60 seconds will be served from the local cache rather than making another RPC call.
