// file_pipe.go
package main

import (
	"fmt"
	"testing"
)

var nodeTest = TestProc{
	filter:           []string{"STARPORT]", "!", "lava_", "ERR_", "panic"},
	expectedEvents:   []string{"🌍", "lava_spec_add", "lava_provider_stake_new", "lava_client_stake_new", "lava_relay_payment"},
	unexpectedEvents: []string{"exit status", "cannot build app", "connection refused", "ERR_client_entries_pairing", "ERR"},
	tests:            events(),
	strict:           false}
var initTest = TestProc{
	filter:           []string{":::", "raw_log", "Error", "error", "panic"},
	expectedEvents:   []string{"init done"},
	unexpectedEvents: []string{"Error"},
	tests:            events(),
	strict:           true}
var providersTest = TestProc{
	filter:           []string{"sent (new/from cache)", "Server", "updated", "server", "error"},
	expectedEvents:   []string{"listening"},
	unexpectedEvents: []string{"ERROR", "refused", "Missing Payment"},
	tests:            events(),
	strict:           true}
var clientTest = TestProc{
	filter:         []string{":::", "reply", "no pairings available", "update", "connect", "rpc", "pubkey", "signal", "Error", "error", "panic"},
	expectedEvents: []string{"update pairing list!", "Client pubkey"},
	// unexpectedEvents: []string{"no pairings available", "error", "Error", "signal: interrupt"},
	unexpectedEvents: []string{"no pairings available", "Error", "signal: interrupt"},
	tests:            events(),
	strict:           true}

func FullFlowTest(t *testing.T) ([]*TestResult, error) {
	prepTest(t)

	// Test Configs
	resetGenesis := true
	init_chain := true
	run_providers_osmosis := true
	run_providers_eth := true
	run_providers_gth := true
	run_providers_ftm := true
	run_providers_juno := true
	run_providers_coshub := true
	run_client_osmosis := true
	run_client_eth := true
	run_client_gth := true
	run_client_ftm := true
	run_client_juno := true

	run_client_coshub := true
	start_lava := "killall ignite; killall lavad; cd " + homepath + " && ignite chain serve -v -r  "
	if !resetGenesis {
		start_lava = "lavad start "
	}

	// Start Test Processes
	node := TestProcess("ignite", start_lava, nodeTest)
	await(node, "lava node is running", lava_up, "awaiting for node to proceed...")

	if init_chain {
		sleep(2)
		init := TestProcess("init", homepath+"scripts/init.sh", initTest)
		await(init, "get init done", init_done, "awaiting for init to proceed...")
	}

	if run_providers_osmosis {
		fmt.Println(" ::: Starting Providers Processes [OSMOSIS] ::: ")
		prov_osm := TestProcess("providers_osmosis", homepath+"scripts/osmosis.sh", providersTest)
		// debugOn(prov_osm)
		println(" ::: Providers Processes Started ::: ")
		await(prov_osm, "Osmosis providers ready", providers_ready, "awaiting for providers to listen to proceed...")

		if run_client_osmosis {
			sleep(1)
			fmt.Println(" ::: Starting Client Process [OSMOSIS] ::: ")
			clientOsmoRPC := TestProcess("clientOsmoRPC", "lavad test_client COS3 tendermintrpc --from user1", clientTest)
			await(node, "relay payment 1/3 osmosis", found_relay_payment, "awaiting for OSMOSIS payment to proceed... ")
			fmt.Println(" ::: GOT OSMOSIS PAYMENT !!!")
			silent(clientOsmoRPC)
			clientOsmoRest := TestProcess("clientOsmoRest", "lavad test_client COS3 rest --from user1", clientTest)
			await(node, "relay payment 1/3 osmosis", found_relay_payment, "awaiting for OSMOSIS payment to proceed... ")
			fmt.Println(" ::: GOT OSMOSIS PAYMENT !!!")
			silent(clientOsmoRest)
			silent(prov_osm)
		}
	}
	if run_providers_eth {
		fmt.Println(" ::: Starting Providers Processes [ETH] ::: ")
		prov_eth := TestProcess("providers_eth", homepath+"scripts/eth.sh", providersTest)
		fmt.Println(" ::: Providers Processes Started ::: ")
		await(prov_eth, "ETH providers ready", providers_ready_eth, "awaiting for providers to listen to proceed...")

		if run_client_eth {
			sleep(1)
			fmt.Println(" ::: Starting Client Process [ETH] ::: ")
			clientETH := TestProcess("clientETH", "lavad test_client ETH1 jsonrpc --from user1", clientTest)
			await(clientETH, "reply rpc", found_rpc_reply, "awaiting for rpc reply to proceed...")
			await(node, "relay payment 3/3 eth", found_relay_payment, "awaiting for ETH payment to proceed...")
			fmt.Println(" ::: GOT ETH PAYMENT !!!")
			silent(clientETH)
			silent(prov_eth)
		}
	}
	if run_providers_gth {
		fmt.Println(" ::: Starting Providers Processes [GTH] ::: ")
		prov_gth := TestProcess("providers_gth", homepath+"scripts/gth.sh", providersTest)
		fmt.Println(" ::: Providers Processes Started ::: ")
		await(prov_gth, "GTH providers ready", providers_ready_eth, "awaiting for providers to listen to proceed...")

		if run_client_gth {
			sleep(1)
			fmt.Println(" ::: Starting Client Process [GTH] ::: ")
			clientGTH := TestProcess("clientGTH", "lavad test_client GTH1 jsonrpc --from user1", clientTest)
			await(clientGTH, "reply rpc", found_rpc_reply, "awaiting for rpc reply to proceed...")
			await(node, "relay payment 3/3 gth", found_relay_payment, "awaiting for GTH payment to proceed...")
			fmt.Println(" ::: GOT GTH PAYMENT !!!")
			silent(clientGTH)
			silent(prov_gth)
		}
	}
	if run_providers_ftm {
		fmt.Println(" ::: Starting Providers Processes [FTM] ::: ")
		prov_ftm := TestProcess("providers_ftm", homepath+"scripts/ftm.sh", providersTest)
		fmt.Println(" ::: Providers Processes Started ::: ")
		await(prov_ftm, "FTM providers ready", providers_ready_eth, "awaiting for providers to listen to proceed...")

		if run_client_ftm {
			sleep(1)
			fmt.Println(" ::: Starting Client Process [FTM] ::: ")
			clientFTM := TestProcess("clientFTM", "lavad test_client FTM250 jsonrpc --from user1", clientTest)
			await(clientFTM, "reply rpc", found_rpc_reply, "awaiting for rpc reply to proceed...")
			await(node, "relay payment 3/3 ftm", found_relay_payment, "awaiting for FTM payment to proceed...")
			fmt.Println(" ::: GOT FTM PAYMENT !!!")
			silent(clientFTM)
			silent(prov_ftm)
		}
	}

	if run_providers_juno {
		fmt.Println(" ::: Starting Providers Processes [JUN1] ::: ")
		prov_JUN1 := TestProcess("providers_juno", homepath+"scripts/juno.sh", providersTest)
		fmt.Println(" ::: Providers Processes Started ::: ")
		await(prov_JUN1, "JUN1 providers ready", providers_ready_eth, "awaiting for providers to listen to proceed...")

		if run_client_juno {
			sleep(1)
			fmt.Println(" ::: Starting Client Process [JUN1] ::: ")
			clientJUN1Rpc := TestProcess("clientJUN1", "lavad test_client JUN1 tendermintrpc --from user1", clientTest)
			await(clientJUN1Rpc, "reply rpc", found_rpc_reply, "awaiting for rpc reply to proceed...")
			await(node, "relay payment 2/2 juno", found_relay_payment, "awaiting for JUN1 payment to proceed...")
			fmt.Println(" ::: GOT JUN1 PAYMENT !!!")
			silent(clientJUN1Rpc)
			clientOsmoRest := TestProcess("clientOsmoRest", "lavad test_client JUN1 rest --from user1", clientTest)
			await(node, "relay payment 2/2 osmosis", found_relay_payment, "awaiting for OSMOSIS payment to proceed... ")
			fmt.Println(" ::: GOT OSMOSIS PAYMENT !!!")
			silent(clientOsmoRest)
			silent(prov_JUN1)
		}
	}

	if run_providers_coshub {
		fmt.Println(" ::: Starting Providers Processes [COS5] ::: ")
		prov_cos5 := TestProcess("providers_coshub", homepath+"scripts/coshub.sh", providersTest)
		fmt.Println(" ::: Providers Processes Started ::: ")
		await(prov_cos5, "COS5 providers ready", providers_ready_eth, "awaiting for providers to listen to proceed...")

		if run_client_coshub {
			sleep(1)
			fmt.Println(" ::: Starting Client Process [COS5] ::: ")
			clientCos5Rpc := TestProcess("clientCOS5", "lavad test_client COS5 tendermintrpc --from user1", clientTest)
			await(clientCos5Rpc, "reply rpc", found_rpc_reply, "awaiting for rpc reply to proceed...")
			await(node, "relay payment 2/2 coshub", found_relay_payment, "awaiting for COS5 payment to proceed...")
			fmt.Println(" ::: GOT COS5 PAYMENT !!!")
			silent(clientCos5Rpc)
			clientOsmoRest := TestProcess("clientOsmoRest", "lavad test_client COS5 rest --from user1", clientTest)
			await(node, "relay payment 2/2 osmosis", found_relay_payment, "awaiting for OSMOSIS payment to proceed... ")
			fmt.Println(" ::: GOT OSMOSIS PAYMENT !!!")
			silent(clientOsmoRest)
			silent(prov_cos5)
		}
	}

	// FINISHED TEST PROCESSESS
	println("::::::::::::::::::::::::::::::::::::::::::::::")
	awaitErrorsTimeout := 10
	fmt.Println(" ::: wait ", awaitErrorsTimeout, " seconds for potential errors...")
	sleep(awaitErrorsTimeout)
	fmt.Println("::::::::::::::::::::::::::::::::::::::::::::::")
	fmt.Println("::::::::::::::::::::::::::::::::::::::::::::::")
	fmt.Println("::::::::::::::::::::::::::::::::::::::::::::::")

	// Finalize & Display Results
	final := finalizeResults(t)

	return final, nil
}

func main() {
	FullFlowTest(nil)
}